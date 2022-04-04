package parser_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/parser"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/registry"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	schemas := []string{
		"../schemas/pain.013.001.07.xsd", "../schemas/pain.014.001.07.xsd",
	}

	p := parser.NewParser(parser.Config{Registry: registry.Config{SimpleTypesInCommonPackage: true}})
	for _, xsdFileName := range schemas {
		msgName := strings.TrimSuffix(filepath.Base(xsdFileName), ".xsd")
		t.Log("Msg: ", msgName)

		b, err := ioutil.ReadFile(xsdFileName)
		require.NoError(t, err)

		msg := registry.ISO20022Message{Name: msgName}
		err = p.Parse(msg, b)
		require.NoError(t, err)
	}

	te := p.TypeRegistry.Types["ActiveCurrencyAndAmount"]
	t.Log("ActiveCurrencyAndAmount", te.Local)

	p.TypeRegistry.ShowInfo(nil)

}
