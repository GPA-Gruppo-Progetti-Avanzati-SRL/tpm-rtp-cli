package docmapper

import (
	"encoding/json"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/vars"
	"github.com/rs/zerolog/log"
	"strings"
)

var builtins = FuncMap{
	"nop":       func(s string) string { return s },
	"trimSpace": func(s string) string { return strings.TrimSpace(s) },
}

func NewMapperClass(opts ...Option) (MappingClass, error) {

	mc := MappingClass{FuncMap: builtins}
	for _, o := range opts {
		_ = o(&mc)
	}

	return mc, nil
}

func NewMapperClassFromYAML(data []byte, opts ...Option) (MappingClass, error) {

	mc := MappingClass{FuncMap: builtins}
	for _, o := range opts {
		_ = o(&mc)
	}

	err := json.Unmarshal(data, &mc)
	if err != nil {
		return mc, err
	}

	return mc, nil
}

func WithFuncMap(funcMap FuncMap) Option {
	return func(mc *MappingClass) error {
		if funcMap != nil {
			m := make(map[string]MappingFunction)
			for n, v := range funcMap {
				m[n] = v
			}

			for n, v := range builtins {
				m[n] = v
			}

			mc.FuncMap = m
		} else {
			mc.FuncMap = builtins
		}

		return nil
	}
}

func (mc *MappingClass) GetResolver(opts ...ResolverOption) (*Resolver, error) {

	/*	opts := []ResolverOption{
		WithMappableDoc(sourceMap),
	}*/

	r := &Resolver{}
	for _, o := range opts {
		err := o(r)
		if err != nil {
			return r, err
		}
	}

	return r, nil
}

func (mc *MappingClass) Map(sourceDoc MappableDocument, targetDoc MappableDocument) error {

	resolver, err := mc.GetResolver(WithMappableDoc(sourceDoc))
	if err != nil {
		return err
	}

	for _, r := range mc.Rules {
		err := r.apply(targetDoc, mc.FuncMap, resolver)
		if err != nil {
			return err
		}
	}

	return nil
}

func (mr *MappingRule) apply(targetDoc MappableDocument, funcs FuncMap, resolver *Resolver) error {
	s, err := varResolver.ResolveVariables(mr.SourceValue, varResolver.SimpleVariableReference, resolver.ResolveVar, true)
	if err != nil {
		return err
	}

	if mr.MapFunc != "" {
		f, ok := funcs[mr.MapFunc]
		if ok {
			s = f(s)
		} else {
			log.Warn().Str("func-name", mr.MapFunc).Msg("cannot find function map")
		}
	}

	if s != "" {
		return targetDoc.Set(mr.TargetPath, s)
	}

	return nil
}
