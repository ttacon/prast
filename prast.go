package prast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"

	"github.com/ttacon/pretty"
)

type Prast interface {
	// FilterByType filters all AST nodes to only include those with the
	// given nodeType (and those nodes' descendents).
	FilterByType(nodeType reflect.Type)
	// FilterByName filters all AST nodes to only include those
	// with the given name (and those nodes' descendents).
	FilterByName(name string)
	// Print prints out the current state that Prast has inspected.
	Print()
	String() string
}

func NewPrast(name string) (Prast, error) {
	fInfo, err := os.Stat(name)
	if err != nil {
		return nil, err
	}

	if fInfo.IsDir() {
		return newPkgPrast(name)
	}
	return newFilePrast(name)
}

func newFilePrast(name string) (Prast, error) {
	fSet := token.NewFileSet()
	node, err := parser.ParseFile(fSet, name, nil, parser.AllErrors)
	return &filePrast{
		fileSet: fSet,
		node:    node,
	}, err
}

func goFilter(fInfo os.FileInfo) bool {
	return filepath.Ext(fInfo.Name()) == "go"
}

func newPkgPrast(name string) (Prast, error) {
	fSet := token.NewFileSet()
	nodes, err := parser.ParseDir(fSet, name, goFilter, parser.AllErrors)
	return &pkgPrast{
		fileSet: fSet,
		nodes:   nodes,
	}, err
}

type filePrast struct {
	node          *ast.File
	fileSet       *token.FileSet
	filteredNodes []ast.Node
}

func (f *filePrast) FilterByType(nodeType reflect.Type) {
	tF := &typeFilter{
		nodeType: nodeType,
	}
	ast.Walk(tF, f.node)
	f.filteredNodes = tF.nodes
}

func (f *filePrast) FilterByName(name string) {
	// TODO(ttacon): ✔
}

func (f *filePrast) Print() {
	if f.filteredNodes == nil {
		pretty.Println(f.node)
		return
	}
	pretty.Println(f.filteredNodes)
}

func (f *filePrast) String() string {
	if f.filteredNodes == nil {
		return pretty.Sprintf("%#v", f.node)
	}
	return pretty.Sprintf("%#v", f.filteredNodes)
}

type pkgPrast struct {
	fileSet       *token.FileSet
	nodes         map[string]*ast.Package
	filteredNodes []ast.Node
}

func (p *pkgPrast) FilterByType(nodeType reflect.Type) {
	p.filteredNodes = nil
	for _, v := range p.nodes {
		tF := &typeFilter{
			nodeType: nodeType,
		}
		ast.Walk(tF, v)
		p.filteredNodes = append(p.filteredNodes, tF.nodes...)
	}
}

func (p *pkgPrast) FilterByName(name string) {
	// TODO(ttacon): ✔
}

func (p *pkgPrast) Print() {
	if p.filteredNodes == nil {
		pretty.Println(p.nodes)
		return
	}
	pretty.Println(p.filteredNodes)
}

func (p *pkgPrast) String() string {
	if p.filteredNodes == nil {
		return pretty.Sprintf("%#v", p.nodes)
	}
	return pretty.Sprintf("%#v", p.filteredNodes)
}

type typeFilter struct {
	nodes    []ast.Node
	nodeType reflect.Type
}

func (t *typeFilter) Visit(node ast.Node) (w ast.Visitor) {
	rt := reflect.TypeOf(node)
	if rt == t.nodeType {
		t.nodes = append(t.nodes, node)
		return nil
	}
	return t
}
