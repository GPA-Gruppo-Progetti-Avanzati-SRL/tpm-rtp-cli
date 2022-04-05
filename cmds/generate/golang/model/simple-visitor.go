package model

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

type SimpleVisitedItem struct {
	Name     string
	Type     string
	IsStruct bool
	IsPtr    bool
	IsArray  bool
}

// TypeWithPackageTrimmed trim the type if the package matches the parameter. Used to avoid self references to package.
func (si SimpleVisitedItem) TypeWithPackageTrimmed(currentPackage string) string {

	n := si.Type
	if currentPackage != "" && strings.HasPrefix(si.Type, currentPackage) {
		n = si.Type[len(currentPackage)+1:]
	}

	return n
}

func (si SimpleVisitedItem) String() string {
	modifier := ""
	if si.IsPtr {
		modifier = "*"
	}

	if si.IsArray {
		modifier = "[]"
	}

	path := fmt.Sprintf("%s %s%s", si.Name, modifier, si.Type)
	return path
}

type SimpleVisitorPath []SimpleVisitedItem

func (p SimpleVisitorPath) Identifier() string {
	v := p.Value()
	return strings.Replace(v, ".", "_", -1)
}

func (p SimpleVisitorPath) ItemPathReference(itemNdx int) string {
	var sb strings.Builder

	if itemNdx == 0 || itemNdx >= len(p) {
		return fmt.Sprintf("out of index %d (len = %d)", itemNdx, len(p))
	}
	// Take the intermediate nodes skipping the first.
	for i := 1; i <= (itemNdx - 1); i++ {
		el := p[i]
		if i > 1 {
			sb.WriteString(".")
		}
		sb.WriteString(el.Name)
		if el.IsArray {
			sb.WriteString("[0]")
		}
	}

	if itemNdx > 1 {
		sb.WriteString(".")
	}
	sb.WriteString(p[itemNdx].Name)
	return sb.String()
}

func (p SimpleVisitorPath) Value() string {
	var sb strings.Builder
	for ndx, item := range p {
		if ndx > 0 {
			sb.WriteString(".")
		}

		n := item.Name
		if n == "_self" {
			n = "Doc"
		}
		sb.WriteString(n)
	}

	return sb.String()
}

func (p SimpleVisitorPath) LastItem() SimpleVisitedItem {
	return p[len(p)-1]
}

func (p SimpleVisitorPath) LastItemIndex() int {
	return len(p) - 1
}

type SimpleVisitor struct {
	NumberOfLeaves int
	currentPath    SimpleVisitorPath
	Paths          []SimpleVisitorPath
}

func (sv SimpleVisitor) currentPathToString() string {
	return sv.pathToString(sv.currentPath)
}

func (sv SimpleVisitor) pathToString(p []SimpleVisitedItem) string {

	var sb strings.Builder
	for _, item := range p {
		sb.WriteString("/")
		sb.WriteString(item.String())
	}

	return sb.String()
}

func (sv *SimpleVisitor) Visit(name, aType string, isStruct, isPrt, isArray bool) (string, error) {
	item := SimpleVisitedItem{Name: name, Type: aType, IsStruct: isStruct, IsPtr: isPrt, IsArray: isArray}
	if !isStruct {
		sv.NumberOfLeaves++
	}
	sv.currentPath = append(sv.currentPath, item)

	var arr SimpleVisitorPath
	arr = append(arr, sv.currentPath...)
	sv.Paths = append(sv.Paths, arr)
	return sv.currentPathToString(), nil
}

func (sv *SimpleVisitor) Pop() (string, error) {

	if len(sv.currentPath) > 0 {
		sv.currentPath = sv.currentPath[0 : len(sv.currentPath)-1]
	} else {
		return "", fmt.Errorf("cannot pop from path: %s", sv.currentPathToString())
	}

	return sv.currentPathToString(), nil
}

func (sv *SimpleVisitor) ShowInfo(currentPackage string) error {
	for i, p := range sv.Paths {
		log.Trace().Int("path-num", i).Msg(sv.pathToString(p))
	}

	return nil
}
