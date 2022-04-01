package parser

import (
	"aqwari.net/xml/xsd"
	"aqwari.net/xml/xsdgen"
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
	"tpm-rtp-cli/iso-20022/registry"
)

type Parser struct {
	TypeRegistry registry.TypeRegistry
}

func NewParser() Parser {
	p := Parser{TypeRegistry: registry.NewTypeRegistry()}
	return p
}

func (p *Parser) Parse(msg registry.ISO20022Message) error {

	var cfg xsdgen.Config
	cfg.Option(
		xsdgen.PackageName(msg.PackageName()),
	)

	msgCode, err := cfg.GenCode(msg.XSDData)
	if err != nil {
		return err
	}

	doc, ok := msgCode.DocType(string(msg.Namespace()))
	if !ok {
		return fmt.Errorf("not a root document")
	}

	p.visit(doc)
	return nil
}

func (p *Parser) visit(cplx *xsd.ComplexType) {

	for _, attr := range cplx.Attributes {

		switch tattr := attr.Type.(type) {
		case *xsd.ComplexType:
		case *xsd.SimpleType:
			p.addSimpleType(tattr)
		case xsd.Builtin:
			p.addBuiltin(cplx.Name.Local, tattr)
		default:
			log.Error().Interface("t", attr.Type).Msg("attr type not matched")
		}
	}

	for _, el := range cplx.Elements {
		switch tel := el.Type.(type) {
		case *xsd.ComplexType:
			p.addComplexType(tel)
			p.visit(tel)
		case *xsd.SimpleType:
			p.addSimpleType(tel)
		case xsd.Builtin:
			p.addBuiltin(el.Name.Space, tel)
		default:
			log.Error().Interface("t", el.Type).Msg("element type not matched")
		}
	}
}

func (p *Parser) addComplexType(t *xsd.ComplexType) {

	ns := registry.ISO20022Namespace(t.Name.Space)
	// log.Trace().Str("space", t.Name.Space).Str("local", t.Name.Local).Msg("found complex type")

	n := t.Name.Local
	if n == "Document" {
		n = strings.Join([]string{n, ns.Message()}, "#")
	}

	te := registry.TypeEntry{
		ElementType: registry.ElementTypeComplex,
		Spaces:      []registry.ISO20022Namespace{ns},
		Local:       n,
		Complex:     t,
	}

	p.addEntry(te)
}

func (p *Parser) addSimpleType(t *xsd.SimpleType) {

	ns := registry.ISO20022Namespace(t.Name.Space)
	// log.Trace().Str("space", t.Name.Space).Str("local", t.Name.Local).Msg("found simple type")

	te := registry.TypeEntry{
		ElementType: registry.ElementTypeSimple,
		Spaces:      []registry.ISO20022Namespace{ns},
		Local:       t.Name.Local,
		Simple:      t,
	}

	p.addEntry(te)
}

func (p *Parser) addBuiltin(s string, t xsd.Builtin) {

	// instead of taking the type namespace (i.e. http://www.w3.org/2001/XMLSchema) i take the element namespace that is I now in what space it has been referred
	ns := registry.ISO20022Namespace(s)
	// log.Trace().Str("space", t.Name().Space).Str("local", t.Name().Local).Msg("found builtin type")

	te := registry.TypeEntry{
		ElementType: registry.ElementTypeBuiltin,
		Spaces:      []registry.ISO20022Namespace{ns},
		Local:       t.Name().Local,
		Builtin:     t,
	}

	p.addEntry(te)
}

func (p *Parser) addEntry(te registry.TypeEntry) {

	// TODO
	if te.Local == "Max35Text--" {
		fmt.Println()
	}

	found, ok := p.TypeRegistry[te.Local]
	if ok {
		if !found.ContainsSpace(te.Spaces[0]) {
			found.Spaces = append(found.Spaces, te.Spaces[0])
		}
		found.RefCount = 1 + found.RefCount
		p.TypeRegistry[found.Local] = found
	} else {
		te.RefCount = 1
		p.TypeRegistry[te.Local] = te
	}

}
