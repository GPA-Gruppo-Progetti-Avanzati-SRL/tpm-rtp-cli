package parser_test

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
	"tpm-rtp-cli/iso-20022/parser"
	"tpm-rtp-cli/iso-20022/registry"
)

func TestParser(t *testing.T) {

	schemas := []string{
		"../schemas/pain.013.001.07.xsd", "../schemas/pain.014.001.07.xsd",
	}

	p := parser.NewParser()
	for _, xsdFileName := range schemas {
		msgName := strings.TrimSuffix(filepath.Base(xsdFileName), ".xsd")
		t.Log("Msg: ", msgName)

		b, err := ioutil.ReadFile(xsdFileName)
		require.NoError(t, err)

		msg := registry.ISO20022Message{Name: msgName, XSDData: b}
		err = p.Parse(msg)
		require.NoError(t, err)
	}

	te := p.TypeRegistry["ActiveCurrencyCode"]
	t.Log("ActiveCurrencyCode", te.Local)

	pkg := registry.PackageNameForMessageName("pain.013.001.07")
	t.Log("complex elements in package -------------", pkg)
	entries := p.TypeRegistry.GetEntriesForPackage(registry.ElementTypeComplex, pkg)
	for i, e := range entries {
		t.Log(i, e.Local, e.RefCount)
	}

	t.Log("simple elements in package -------------", pkg)
	entries = p.TypeRegistry.GetEntriesForPackage(registry.ElementTypeSimple, pkg)
	for i, e := range entries {
		t.Log(i, e.Local, e.RefCount)
	}

	t.Log("builtin elements in package -------------", pkg)
	entries = p.TypeRegistry.GetEntriesForPackage(registry.ElementTypeBuiltin, pkg)
	for i, e := range entries {
		t.Log(i, e.Local, e.RefCount)
	}

	pkg = registry.PackageNameForMessageName("common")
	t.Log("complex elements in package -------------", pkg)
	entries = p.TypeRegistry.GetEntriesForPackage(registry.ElementTypeComplex, pkg)
	for i, e := range entries {
		t.Log(i, e.Local, e.RefCount)
	}

	t.Log("simple elements in package -------------", pkg)
	entries = p.TypeRegistry.GetEntriesForPackage(registry.ElementTypeSimple, pkg)
	for i, e := range entries {
		t.Log(i, e.Local, e.RefCount)
	}

	t.Log("builtin elements in package -------------", pkg)
	entries = p.TypeRegistry.GetEntriesForPackage(registry.ElementTypeBuiltin, pkg)
	for i, e := range entries {
		t.Log(i, e.Local, e.RefCount)
	}
}
