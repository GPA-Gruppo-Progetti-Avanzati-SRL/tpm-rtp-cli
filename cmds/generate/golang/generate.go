package golang

import (
	"embed"
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/cmds/generate/golang/model"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/cmds/generate/registry"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

//go:embed templates/*
var templates embed.FS

const (
	TmplComplexTypes           = "templates/%s/complex-types.tmpl"
	TmplComplexTypesOps        = "templates/%s/complex-types-ops.tmpl"
	TmplSimpleTypes            = "templates/%s/simple-types.tmpl"
	TmplSimpleTypesOps         = "templates/%s/simple-types-ops.tmpl"
	TmplSimpleTypesStringOps   = "templates/%s/simple-types-string-ops.tmpl"
	TmplSimpleTypesB64Ops      = "templates/%s/simple-types-b64-ops.tmpl"
	TmplSimpleTypesDateOps     = "templates/%s/simple-types-date-ops.tmpl"
	TmplSimpleTypesDateTimeOps = "templates/%s/simple-types-datetime-ops.tmpl"
	RestrictionUtil            = "templates/%s/restriction-util.tmpl"
	XsDtTypes                  = "templates/%s/xsdt-types.tmpl"
	XsDtTypesMethods           = "templates/%s/xsdt-types-methods.tmpl"
	DocumentType               = "templates/%s/document-type.tmpl"
	DocumentPaths              = "templates/%s/document-paths.tmpl"
	DocumentSetComplexOps      = "templates/%s/document-complex-set-ops.tmpl"
	DocumentSetSimpleOps       = "templates/%s/document-simple-set-ops.tmpl"
	DocumentGetOps             = "templates/%s/document-get-ops.tmpl"
	DocumentExample            = "templates/%s/document-example.tmpl"
	DocumentExampleNode        = "templates/%s/document-example-node.tmpl"
)

func complexTypesTmplList(version string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(TmplComplexTypes, version))
	return s
}

func complexTypesOpsTmplList(version string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(TmplComplexTypesOps, version))
	return s
}

func simpleTypesTmplList(version string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(TmplSimpleTypes, version))
	return s
}

func simpleTypesOpsTmplList(version string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(TmplSimpleTypesOps, version))
	s = append(s, fmt.Sprintf(TmplSimpleTypesStringOps, version))
	s = append(s, fmt.Sprintf(TmplSimpleTypesDateOps, version))
	s = append(s, fmt.Sprintf(TmplSimpleTypesDateTimeOps, version))
	s = append(s, fmt.Sprintf(TmplSimpleTypesB64Ops, version))
	return s
}

func restrictionUtilTmplList(version string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(RestrictionUtil, version))
	return s
}

func documentExampleTmplList(version string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(DocumentExample, version))
	s = append(s, fmt.Sprintf(DocumentExampleNode, version))

	return s
}

func documentTypeTmplList(version string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(DocumentType, version))

	return s
}

func documentSetComplexOpsTmplList(version string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(DocumentSetComplexOps, version))

	return s
}

func documentSetSimpleOpsTmplList(version string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(DocumentSetSimpleOps, version))

	return s
}

func documentGetOpsTmplList(version string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(DocumentGetOps, version))

	return s
}

func documentPathsTmplList(version string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(DocumentPaths, version))

	return s
}

func xsdtTypesTmplList(version string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(XsDtTypes, version))
	return s
}

func xsdtTypesMethodsTmplList(version string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(XsDtTypesMethods, version))
	return s
}

type GenerationContext struct {
	MsgName     string
	PackageName string
	ProducedAt  time.Time
	Model       *model.GoModel
}

func Generate(cfg *Config, gm *model.GoModel) error {

	err := setupOutputFolders(cfg, gm.Msgs)
	if err != nil {
		return err
	}

	genCtx := GenerationContext{Model: gm, ProducedAt: time.Now()}
	for _, msg := range gm.Msgs {

		genCtx.MsgName = msg.Name
		if err := emitPackage(msg.PackageName(), genCtx, cfg.OutFolder, cfg.FormatCode); err != nil {
			return err
		}
	}

	if err := emitPackage("common", genCtx, cfg.OutFolder, cfg.FormatCode); err != nil {
		return err
	}

	if err := emitXSDT(genCtx, cfg.OutFolder, cfg.FormatCode); err != nil {
		return err
	}

	return nil
}

func emitXSDT(genCtx GenerationContext, outFolder string, formatCode bool) error {
	log.Info().Msg("generation of xsdt types")
	if err := emit(genCtx, filepath.Join(outFolder, "xsdt"), strings.Join([]string{"types", "go"}, "."), xsdtTypesTmplList("v2"), formatCode); err != nil {
		return err
	}

	log.Info().Msg("generation of xsdt types methods")
	if err := emit(genCtx, filepath.Join(outFolder, "xsdt"), strings.Join([]string{"types-methods", "go"}, "."), xsdtTypesMethodsTmplList("v2"), formatCode); err != nil {
		return err
	}

	return nil
}

func emitPackage(pkgName string, genCtx GenerationContext, outFolder string, formatCode bool) error {
	/*
	 * complex-types for message
	 */
	genCtx.PackageName = pkgName

	countComplexTypes := genCtx.Model.CountTypes(pkgName, true)
	if countComplexTypes > 0 {
		log.Info().Int("num-types", countComplexTypes).Str("pkg", pkgName).Msg("generation of complex types")
		if err := emit(genCtx, filepath.Join(outFolder, pkgName), strings.Join([]string{"complex-types", "go"}, "."), complexTypesTmplList("v2"), formatCode); err != nil {
			return err
		}

		/*
		 * complex-types ops for message
		 */
		if err := emit(genCtx, filepath.Join(outFolder, pkgName), strings.Join([]string{"complex-types-ops", "go"}, "."), complexTypesOpsTmplList("v2"), formatCode); err != nil {
			return err
		}

		/*
		 * document example
		 */
		if pkgName != "common" {
			if err := emit(genCtx, filepath.Join(outFolder, pkgName), strings.Join([]string{"document", "go"}, "."), documentTypeTmplList("v2"), formatCode); err != nil {
				return err
			}

			if err := emit(genCtx, filepath.Join(outFolder, pkgName), strings.Join([]string{"document-paths", "go"}, "."), documentPathsTmplList("v2"), formatCode); err != nil {
				return err
			}

			if err := emit(genCtx, filepath.Join(outFolder, pkgName), strings.Join([]string{"document-complex-set-ops", "go"}, "."), documentSetComplexOpsTmplList("v2"), formatCode); err != nil {
				return err
			}

			if err := emit(genCtx, filepath.Join(outFolder, pkgName), strings.Join([]string{"document-simple-set-ops", "go"}, "."), documentSetSimpleOpsTmplList("v2"), formatCode); err != nil {
				return err
			}

			if err := emit(genCtx, filepath.Join(outFolder, pkgName), strings.Join([]string{pkgName + "_test", "go"}, "."), documentExampleTmplList("v2"), formatCode); err != nil {
				return err
			}
		}

	} else {
		fileName := filepath.Join(outFolder, pkgName, strings.Join([]string{"complex-types", "go"}, "."))
		log.Info().Str("file", fileName).Str("pkg", pkgName).Msg("no complex types found, clearing output files")
		_ = os.Remove(fileName)

		fileName = filepath.Join(outFolder, pkgName, strings.Join([]string{"complex-types-ops", "go"}, "."))
		log.Info().Str("file", fileName).Str("pkg", pkgName).Msg("no complex types found, clearing output file")
		_ = os.Remove(fileName)
	}

	countSimpleTypes := genCtx.Model.CountTypes(pkgName, false)
	if countSimpleTypes > 0 {
		log.Info().Int("num-types", countSimpleTypes).Str("pkg", pkgName).Msg("generation of simple types")
		/*
		 * simple-types for message
		 */
		if err := emit(genCtx, filepath.Join(outFolder, pkgName), strings.Join([]string{"simple-types", "go"}, "."), simpleTypesTmplList("v2"), formatCode); err != nil {
			return err
		}

		/*
		 * simple-types ops for message
		 */
		if err := emit(genCtx, filepath.Join(outFolder, pkgName), strings.Join([]string{"simple-types-ops", "go"}, "."), simpleTypesOpsTmplList("v2"), formatCode); err != nil {
			return err
		}

		/*
		 * restriction util for simple types
		 */
		if err := emit(genCtx, filepath.Join(outFolder, pkgName), strings.Join([]string{"restriction-util", "go"}, "."), restrictionUtilTmplList("v2"), formatCode); err != nil {
			return err
		}

	} else {
		fileName := filepath.Join(outFolder, pkgName, strings.Join([]string{"simple-types", "go"}, "."))
		log.Info().Str("file", fileName).Str("pkg", pkgName).Msg("no simple types found, clearing output files")
		_ = os.Remove(fileName)

		fileName = filepath.Join(outFolder, pkgName, strings.Join([]string{"simple-types-ops", "go"}, "."))
		log.Info().Str("file", fileName).Str("pkg", pkgName).Msg("no simple types found, clearing output file")
		_ = os.Remove(fileName)
	}

	return nil
}

func setupOutputFolders(cfg *Config, msgs []registry.ISO20022Message) error {
	for _, msg := range msgs {
		_, err := cfg.SetupOutFolder(msg.PackageName())
		if err != nil {
			return err
		}
	}

	_, err := cfg.SetupOutFolder("common")
	if err != nil {
		return err
	}

	return nil
}

func emit(genCtx GenerationContext, outFolder string, generatedFileName string, templates []string, formatCode bool) error {
	if t, err := loadTemplate(templates...); err == nil {
		destinationFile := filepath.Join(outFolder, generatedFileName)
		log.Info().Str("file", destinationFile).Msg("generating text from template")

		if pkgTemplate, err := util.ParseTemplates(t, getTemplateUtilityFunctions(genCtx.Model)); err != nil {
			return err
		} else {
			if err := util.ProcessTemplateWrite2File(pkgTemplate, genCtx, destinationFile, formatCode); err != nil {
				return err
			}
		}
	} else {
		return errors.New("unable to load template ...skipping")
	}

	return nil
}

func getTemplateUtilityFunctions(gm *model.GoModel) template.FuncMap {

	fMap := template.FuncMap{
		"getImportForPackage": func(pkgName string) string {
			return gm.Cfg.PackageImport(pkgName)
		},
		"getComplexTypesImports": func(pkgName string) []string {
			return gm.GetImports(pkgName, true)
		},
		"getSimpleTypesImports": func(pkgName string) []string {
			return gm.GetImports(pkgName, false)
		},
		"getComplexTypes": func(pkgName string) []model.GoTypeDefinition {
			return gm.GetTypes(pkgName, true)
		},
		"getSimpleTypes": func(pkgName string) []model.GoTypeDefinition {
			return gm.GetTypes(pkgName, false)
		},
		"treeVisit": func(pkgName string) *model.TreeVisitor {
			v := &model.TreeVisitor{}
			_ = gm.VisitDocument(pkgName, v)
			return v
		},
		"simpleVisit": func(pkgName string) *model.SimpleVisitor {
			v := &model.SimpleVisitor{}
			_ = gm.VisitDocument(pkgName, v)
			return v
		},
		"mustToFunctionSignature": func(typeName string) string {
			ndx := strings.Index(typeName, ".")
			if ndx > 0 {
				var sb strings.Builder
				sb.WriteString(typeName[0:ndx])
				sb.WriteString(".MustTo")
				sb.WriteString(typeName[ndx+1:])
				return sb.String()
			}
			return "MustTo" + typeName
		},
		// dict is to put together a number of params to be passed to a child template.
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
		"camelize": func(s string) string {
			return util.Camelize(s)
		},
	}

	return fMap
}

func loadTemplate(templatePath ...string) ([]util.TemplateInfo, error) {

	res := make([]util.TemplateInfo, 0)
	for _, tpath := range templatePath {

		b, err := templates.ReadFile(tpath)
		if err != nil {
			return nil, err
		}
		res = append(res, util.TemplateInfo{Name: filepath.Base(tpath), Content: string(b)})
	}

	return res, nil
}
