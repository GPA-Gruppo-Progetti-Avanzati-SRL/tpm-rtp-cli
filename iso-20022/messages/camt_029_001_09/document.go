// Package camt_029_001_09
// Do not Edit. This stuff it's been automatically generated.
package camt_029_001_09

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"reflect"
	"strings"
	"sync"
)

// Document type definition
type Document struct {
	XMLName         xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Document"`
	RsltnOfInvstgtn ResolutionOfInvestigationV09 `xml:"RsltnOfInvstgtn"`
	mapper          *common.Mapper               `xml:"-"`
}

func NewDocument() Document {
	d := Document{
		mapper: mapper(),
	}

	return d
}

func (d *Document) ToXML() ([]byte, error) {
	w := &bytes.Buffer{}
	w.Write([]byte(xml.Header))

	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	err := enc.Encode(d)
	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

func NewDocumentFromXML(b []byte) (*Document, error) {
	d := &Document{}
	err := xml.Unmarshal(b, d)
	return d, err
}

// IsValid checks if Document is valid
func (d Document) IsValid(optional bool) bool {

	valid := true
	valid = valid && d.RsltnOfInvstgtn.IsValid(false)

	return valid
}

/*
 * Document reflection funcs
 */

// NameMapper is used to map column names to struct field names.  By default,
// it uses strings.ToLower to lowercase struct field names.  It can be set
// to whatever you want, but it is encouraged to be set before use
// as name-to-field mappings are cached after first use on a type.
var NameMapper = NameMapperFunc

func NameMapperFunc(s string) string {
	return strings.TrimPrefix(s, "urn:iso:std:iso:20022:tech:xsd:pain.013.001.07")
}

// Rather than creating on init, this is created when necessary so that
// importers have time to customize the NameMapper.
var mpr *common.Mapper

// mprMu protects mpr.
var mprMu sync.Mutex

// mapper returns a valid mapper using the configured NameMapper func.
func mapper() *common.Mapper {
	mprMu.Lock()
	defer mprMu.Unlock()

	if mpr == nil {
		mpr = common.NewMapperFunc("xml", NameMapper)
	}
	return mpr
}

// fieldsByName fills a values interface with fields from the passed value based
// on the traversals in int.  If ptrs is true, return addresses instead of values.
// We write this instead of using FieldsByName to save allocations and map lookups
// when iterating over many rows.  Empty traversals will get an interface pointer.
// Because of the necessity of requesting ptrs or values, it's considered a bit too
// specialized for inclusion in reflectx itself.
func fieldsByTraversal(v reflect.Value, traversals [][]int, values []interface{}, ptrs bool) error {
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return errors.New("argument not a struct")
	}

	for i, traversal := range traversals {
		if len(traversal) == 0 {
			values[i] = new(interface{})
			continue
		}
		f := common.FieldByIndexes(v, traversal)
		if ptrs {
			values[i] = f.Addr().Interface()
		} else {
			values[i] = f.Interface()
		}
	}
	return nil
}
