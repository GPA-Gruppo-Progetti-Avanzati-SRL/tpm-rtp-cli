package model

import (
	"fmt"
	"strings"
)

type SimpleVisitedItem struct {
	Name     string
	Type     string
	IsStruct bool
	IsPtr    bool
	IsArray  bool
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

type SimpleVisitor struct {
	NumberOfLeaves int
	currentPath    []SimpleVisitedItem
	paths          [][]SimpleVisitedItem
}

func (sv SimpleVisitor) currentPathToString() string {

	var sb strings.Builder
	for _, item := range sv.currentPath {
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
	sv.paths = append(sv.paths, sv.currentPath)
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
