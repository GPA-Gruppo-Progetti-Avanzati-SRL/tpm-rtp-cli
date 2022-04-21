package docmapper

import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"

type MappingRule struct {
	Name        string `mapstructure:"name" yaml:"name" json:"name"`
	SourceValue string `mapstructure:"source-value" yaml:"source-value" json:"source-value"`
	MapFunc     string `mapstructure:"map-func" yaml:"map-func" json:"map-func"`
	TargetPath  string `mapstructure:"target-path" yaml:"target-path" json:"target-path"`
}

type MappingFunction func(s string) string
type FuncMap map[string]MappingFunction

type Option func(mc *MappingClass) error

type MappingClass struct {
	Name    string        `mapstructure:"name" yaml:"name"  json:"name"`
	MsgName string        `mapstructure:"msg-name" yaml:"msg-name"  json:"msg-name"`
	Rules   []MappingRule `mapstructure:"index" yaml:"index" json:"index"`
	FuncMap FuncMap       `mapstructure:"-" yaml:"-" json:"-"`
}

type MappableDocument interface {
	Set(docPath string, src interface{}) error
	Get(docPath string) (interface{}, error)
}

type MappableDocMap map[string]interface{}

func (m MappableDocMap) Set(docPath string, src interface{}) error {
	return util.SetProperty(m, docPath, src)
}

func (m MappableDocMap) Get(docPath string) (interface{}, error) {
	return util.GetProperty(m, docPath), nil
}
