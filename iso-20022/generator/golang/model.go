package golang

import (
	"aqwari.net/xml/xsd"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/registry"
	"github.com/rs/zerolog/log"
	"strings"
)

func NewModel(cfg *ModelConfig, msgs []registry.ISO20022Message, tr registry.TypeRegistry) (GoModel, error) {
	gm := GoModel{cfg: cfg, Msgs: msgs, TypeRegistry: tr}

	for _, msg := range msgs {
		goMsgTypeDefs, err := gm.newTypeDefinitions(msg.PackageName(), tr)
		if err != nil {
			return gm, err
		}

		gm.TypeDefs = append(gm.TypeDefs, goMsgTypeDefs...)
	}

	goCommonTypedefs, err := gm.newTypeDefinitions("common", tr)
	if err != nil {
		return gm, err
	}
	gm.TypeDefs = append(gm.TypeDefs, goCommonTypedefs...)
	return gm, nil
}

func (gm *GoModel) newTypeDefinitions(pkgName string, tr registry.TypeRegistry) ([]GoTypeDefinition, error) {

	goTypeDefs, err := gm.newComplexTypeDefinitions(pkgName, tr)
	if err != nil {
		return nil, err
	}

	goSimpleTypeDefs, err := gm.newSimpleTypeDefinitions(pkgName, tr)
	if err != nil {
		return nil, err
	}

	return append(goTypeDefs, goSimpleTypeDefs...), nil
}

func (gm *GoModel) newComplexTypeDefinitions(pkgName string, tr registry.TypeRegistry) ([]GoTypeDefinition, error) {

	var goStructs []GoTypeDefinition
	entries := tr.GetEntriesForPackage(registry.ElementTypeComplex, pkgName)
	for _, e := range entries {
		tpn := e.Local
		if strings.HasPrefix(e.Local, "Document#") {
			tpn = "Document"
		}

		gs := GoTypeDefinition{PackageName: pkgName, Name: tpn, StructDef: true, Type: GoType{Name: tpn}, Doc: e.Complex.Doc}
		for _, attr := range e.Complex.Attributes {
			var err error
			var gt GoType
			ga := GoAttr{Name: attr.Name.Local, XMLAttr: true, Optional: attr.Optional}
			switch tattr := attr.Type.(type) {
			case *xsd.SimpleType:
				gt, err = gm.goTypeOfElement(tr, pkgName, tattr.Name.Local)
				if err != nil {
					return nil, err
				}

				ga.Type = gt
			case xsd.Builtin:
				ga.Type, err = MapBuiltinToGoType(tattr.String(), gm.cfg.XsDtPackageImport)
				if err != nil {
					return nil, fmt.Errorf("%s - %s", attr.Name.Local, err.Error())
				}
			}

			gs.Attrs = append(gs.Attrs, ga)
		}

		// Handle cases such as: ActiveCurrencyAndAmount
		if len(e.Complex.Elements) == 0 {
			if tbase, ok := e.Complex.Base.(xsd.Builtin); ok {
				ga := GoAttr{Name: "Value", Chardata: true}
				var err error
				ga.Type, err = MapBuiltinToGoType(tbase.String(), gm.cfg.XsDtPackageImport)
				if err != nil {
					return nil, fmt.Errorf("%s - %s", e.Local, err.Error())
				}
				gs.Attrs = append(gs.Attrs, ga)
			}
		}

		for _, el := range e.Complex.Elements {

			var err error
			var gt GoType
			ga := GoAttr{Name: el.Name.Local, Array: el.Plural, Optional: el.Optional}

			switch tel := el.Type.(type) {
			case *xsd.ComplexType:
				gt, err = gm.goTypeOfElement(tr, pkgName, tel.Name.Local)
				if err != nil {
					return nil, err
				}

				ga.StructType = true
				ga.Type = gt
			case *xsd.SimpleType:
				gt, err = gm.goTypeOfElement(tr, pkgName, tel.Name.Local)
				if err != nil {
					return nil, err
				}

				ga.Type = gt
			case xsd.Builtin:
				ga.Type, err = MapBuiltinToGoType(tel.String(), gm.cfg.XsDtPackageImport)
				if err != nil {
					return nil, fmt.Errorf("(%s) - %s", el.Name.Local, err.Error())
				}

				if ga.Name == "" {
					ga.Name = "Item"
				}
			}

			gs.Attrs = append(gs.Attrs, ga)
		}

		goStructs = append(goStructs, gs)
	}
	return goStructs, nil
}

func (gm *GoModel) newSimpleTypeDefinitions(pkgName string, tr registry.TypeRegistry) ([]GoTypeDefinition, error) {

	var goStructs []GoTypeDefinition
	entries := tr.GetEntriesForPackage(registry.ElementTypeSimple, pkgName)
	for _, e := range entries {
		if bt, ok := e.Simple.Base.(xsd.Builtin); ok {
			gt, err := MapBuiltinToGoType(bt.String(), gm.cfg.XsDtPackageImport)
			if err != nil {
				return nil, err
			}
			r := e.Restrictions
			gs := GoTypeDefinition{PackageName: pkgName, Name: e.Local, StructDef: false, Type: gt, Doc: e.Simple.Doc, Restrictions: r}
			goStructs = append(goStructs, gs)
		} else {
			return nil, fmt.Errorf("the Base type is not builtin for %s", e.Local)
		}
	}

	return goStructs, nil
}

func (gm *GoModel) goTypeOfElement(tr registry.TypeRegistry, currentPkg string, localTypeName string) (GoType, error) {
	te, ok := tr.GetEntryForType(localTypeName)
	if !ok {
		return GoType{}, fmt.Errorf("this is weird.... not found type %s in registry", localTypeName)
	}

	gtn := localTypeName
	reqImport := ""
	pkgName := te.ComputePackageName(tr.Cfg.SimpleTypesInCommonPackage)
	if pkgName != currentPkg {
		gtn = strings.Join([]string{pkgName, gtn}, ".")
		reqImport = gm.resolveGoImport(pkgName)
	}

	return GoType{Name: gtn, Import: reqImport}, nil
}

func (gm *GoModel) resolveGoImport(importName string) string {
	if importName == "common" {
		return gm.cfg.CommonPackageImport
	}

	return importName
}

// GetTypes Method
func (gm *GoModel) GetTypes(pkgName string, structDefs bool) []GoTypeDefinition {

	var types []GoTypeDefinition
	for _, td := range gm.TypeDefs {
		if td.PackageName == pkgName && td.StructDef == structDefs {
			types = append(types, td)
		}
	}

	return types
}

// CountTypes Method
func (gm *GoModel) CountTypes(pkgName string, structDefs bool) int {

	count := 0
	for _, td := range gm.TypeDefs {
		if td.PackageName == pkgName && td.StructDef == structDefs {
			count++
		}
	}

	return count
}

// GetImports Method
func (gm *GoModel) GetImports(pkgName string, structDef bool) []string {

	importsMap := make(map[string]struct{})
	for _, td := range gm.TypeDefs {

		if td.StructDef != structDef {
			continue
		}

		if td.PackageName == pkgName {
			if td.Type.Import != "" {
				if _, ok := importsMap[td.Type.Import]; !ok {
					importsMap[td.Type.Import] = struct{}{}
				}
			}

			for _, a := range td.Attrs {
				if a.Type.Import != "" {
					if _, ok := importsMap[a.Type.Import]; !ok {
						importsMap[a.Type.Import] = struct{}{}
					}
				}
			}
		}

	}

	var imports []string
	for n, _ := range importsMap {
		imports = append(imports, n)
	}
	return imports
}

// ShowInfo Method
func (gm *GoModel) ShowInfo() {

	// It's kind of three loops
	for _, msg := range gm.Msgs {
		log.Trace().Str("msg", msg.Name).Str("pkg", msg.PackageName()).Msg("package information")
		log.Trace().Msg("composite defs")
		types := gm.GetTypes(msg.PackageName(), true)
		showTypeDefinitions(types)

		log.Trace().Msg("simple information")
		types = gm.GetTypes(msg.PackageName(), false)
		showTypeDefinitions(types)
	}

	log.Trace().Msg("common composite defs")
	types := gm.GetTypes("common", true)
	showTypeDefinitions(types)

	log.Trace().Msg("common simple information")
	types = gm.GetTypes("common", false)
	showTypeDefinitions(types)
}

func showTypeDefinitions(types []GoTypeDefinition) {
	for _, td := range types {
		log.Trace().Str("name", td.Name).Str("pkg", td.PackageName).Str("type", td.Type.Name).Str("import", td.Type.Import).Msg("### type def")
		for _, a := range td.Attrs {
			log.Trace().Str("name", a.Name).Bool("built-in", a.Type.IsBuiltin).Str("type", a.Type.Name).Interface("struct,array,opt,xml-attr", []bool{a.StructType, a.Array, a.Optional, a.XMLAttr}).Str("import", a.Type.Import).Msg("--- attr def")
		}
	}
}
