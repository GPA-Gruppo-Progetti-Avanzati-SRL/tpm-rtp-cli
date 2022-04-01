package registry

import (
	"aqwari.net/xml/xsd"
	"fmt"
	"strings"
)

type ISO20022Namespace string

func ISO20022NamespaceForMessage(n string) ISO20022Namespace {
	return ISO20022Namespace(fmt.Sprintf("urn:iso:std:iso:20022:tech:xsd:%s", n))
}

func (n ISO20022Namespace) Message() string {
	return strings.TrimPrefix(string(n), "urn:iso:std:iso:20022:tech:xsd:")
}

func PackageNameForMessageName(n string) string {
	return strings.ReplaceAll(n, ".", "_")
}

type ISO20022Message struct {
	Name    string
	XSDData []byte
}

func (i *ISO20022Message) PackageName() string {
	return PackageNameForMessageName(i.Name)
}

func (i *ISO20022Message) Namespace() ISO20022Namespace {
	return ISO20022NamespaceForMessage(i.Name)
}

type ElementType string

const (
	ElementTypeComplex ElementType = "complex"
	ElementTypeSimple              = "simple"
	ElementTypeBuiltin             = "builtin"
)

type TypeEntry struct {
	RefCount    int
	ElementType ElementType
	Spaces      []ISO20022Namespace
	Local       string
	Complex     *xsd.ComplexType
	Simple      *xsd.SimpleType
	Builtin     xsd.Builtin
}

func (te *TypeEntry) ContainsSpace(ns ISO20022Namespace) bool {
	for _, s := range te.Spaces {
		if s == ns {
			return true
		}
	}

	return false
}

func (te *TypeEntry) ComputePackageName() string {
	pkg := strings.ReplaceAll(te.Spaces[0].Message(), ".", "_")
	if len(te.Spaces) > 1 {
		pkg = "common"
	}

	return pkg
}

type TypeRegistry map[string]TypeEntry

func NewTypeRegistry() TypeRegistry {
	return make(map[string]TypeEntry)
}

func (tr TypeRegistry) GetEntriesForPackage(elType ElementType, pkgName string) []TypeEntry {

	var tes []TypeEntry
	for _, te := range tr {

		// TODO
		if te.Local == "Max35Text--" {
			fmt.Println()
		}
		if te.ElementType == elType {
			if pkgName == te.ComputePackageName() {
				tes = append(tes, te)
			}
		}
	}

	return tes
}
