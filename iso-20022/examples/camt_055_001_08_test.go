package examples_test

import (
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
)

const ()

func TestDocumentCamt_055_001_08(t *testing.T) {

	b, err := camt_055_001_08_Document.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_Camt_055_001_08_1, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_Camt_055_001_08_1)

}
