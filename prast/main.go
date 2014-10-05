package main

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"os"
	"reflect"
	"strings"

	"github.com/ttacon/chalk"
	"github.com/ttacon/prast"
)

var (
	nodeType = flag.String("node", "", "node type to filter for")
	name     = flag.String("name", "", "name to filter for")
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("err: must supply a file to inspect")
		return
	}

	p, err := prast.NewPrast(args[0])
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	if len(*nodeType) != 0 {
		astType, err := astNodeTypeFromName(*nodeType)
		if err != nil {
			errorLog(fmt.Sprint("err: ", err))
			os.Exit(1)
		}
		p.FilterByType(astType)
	} else if len(*name) != 0 {
		p.FilterByName(*name)
	}
	p.Print()
}

func errorLog(val string) {
	fmt.Println(chalk.Red.Color(val))
}

func astNodeTypeFromName(name string) (t reflect.Type, err error) {
	if !strings.HasPrefix(name, "*ast.") && !strings.HasPrefix(name, "ast.") {
		name = "*ast." + name
	} else if !strings.HasPrefix(name, "*ast.") {
		name = "*" + name
	}
	errorLog(name)
	switch name {
	case "*ast.BadDecl":
		t = reflect.TypeOf(&ast.BadDecl{})
	case "*ast.GenDecl":
		t = reflect.TypeOf(&ast.GenDecl{})
	case "*ast.FuncDecl":
		t = reflect.TypeOf(&ast.FuncDecl{})
	case "*ast.ImportSpec":
		t = reflect.TypeOf(&ast.ImportSpec{})
	case "*ast.ValueSpec":
		t = reflect.TypeOf(&ast.ValueSpec{})
	case "*ast.TypeSpec":
		t = reflect.TypeOf(&ast.TypeSpec{})
	case "*ast.BadStmt":
		t = reflect.TypeOf(&ast.BadStmt{})
	case "*ast.DeclStmt":
		t = reflect.TypeOf(&ast.DeclStmt{})
	case "*ast.EmptyStmt":
		t = reflect.TypeOf(&ast.EmptyStmt{})
	case "*ast.LabeledStmt":
		t = reflect.TypeOf(&ast.LabeledStmt{})
	case "*ast.ExprStmt":
		t = reflect.TypeOf(&ast.ExprStmt{})
	case "*ast.SendStmt":
		t = reflect.TypeOf(&ast.SendStmt{})
	case "*ast.IncDecStmt":
		t = reflect.TypeOf(&ast.IncDecStmt{})
	case "*ast.AssignStmt":
		t = reflect.TypeOf(&ast.AssignStmt{})
	case "*ast.GoStmt":
		t = reflect.TypeOf(&ast.GoStmt{})
	case "*ast.DeferStmt":
		t = reflect.TypeOf(&ast.DeferStmt{})
	case "*ast.ReturnStmt":
		t = reflect.TypeOf(&ast.ReturnStmt{})
	case "*ast.BranchStmt":
		t = reflect.TypeOf(&ast.BranchStmt{})
	case "*ast.BlockStmt":
		t = reflect.TypeOf(&ast.BlockStmt{})
	case "*ast.IfStmt":
		t = reflect.TypeOf(&ast.IfStmt{})
	case "*ast.CaseClause":
		t = reflect.TypeOf(&ast.CaseClause{})
	case "*ast.SwitchStmt":
		t = reflect.TypeOf(&ast.SwitchStmt{})
	case "*ast.TypeSwitchStmt":
		t = reflect.TypeOf(&ast.TypeSwitchStmt{})
	case "*ast.CommClause":
		t = reflect.TypeOf(&ast.CommClause{})
	case "*ast.SelectStmt":
		t = reflect.TypeOf(&ast.SelectStmt{})
	case "*ast.ForStmt":
		t = reflect.TypeOf(&ast.ForStmt{})
	case "*ast.RangeStmt":
		t = reflect.TypeOf(&ast.RangeStmt{})
	case "*ast.BadExpr":
		t = reflect.TypeOf(&ast.BadExpr{})
	case "*ast.Ident":
		t = reflect.TypeOf(&ast.Ident{})
	case "*ast.Ellipsis":
		t = reflect.TypeOf(&ast.Ellipsis{})
	case "*ast.BasicLit":
		t = reflect.TypeOf(&ast.BasicLit{})
	case "*ast.FuncLit":
		t = reflect.TypeOf(&ast.FuncLit{})
	case "*ast.CompositeLit":
		t = reflect.TypeOf(&ast.CompositeLit{})
	case "*ast.ParenExpr":
		t = reflect.TypeOf(&ast.ParenExpr{})
	case "*ast.SelectorExpr":
		t = reflect.TypeOf(&ast.SelectorExpr{})
	case "*ast.IndexExpr":
		t = reflect.TypeOf(&ast.IndexExpr{})
	case "*ast.SliceExpr":
		t = reflect.TypeOf(&ast.SliceExpr{})
	case "*ast.TypeAssertExpr":
		t = reflect.TypeOf(&ast.TypeAssertExpr{})
	case "*ast.CallExpr":
		t = reflect.TypeOf(&ast.CallExpr{})
	case "*ast.StarExpr":
		t = reflect.TypeOf(&ast.StarExpr{})
	case "*ast.UnaryExpr":
		t = reflect.TypeOf(&ast.UnaryExpr{})
	case "*ast.BinaryExpr":
		t = reflect.TypeOf(&ast.BinaryExpr{})
	case "*ast.KeyValueExpr":
		t = reflect.TypeOf(&ast.KeyValueExpr{})
	case "*ast.ArrayType":
		t = reflect.TypeOf(&ast.ArrayType{})
	case "*ast.StructType":
		t = reflect.TypeOf(&ast.StructType{})
	case "*ast.FuncType":
		t = reflect.TypeOf(&ast.FuncType{})
	case "*ast.InterfaceType":
		t = reflect.TypeOf(&ast.InterfaceType{})
	case "*ast.MapType":
		t = reflect.TypeOf(&ast.MapType{})
	case "*ast.ChanType":
		t = reflect.TypeOf(&ast.ChanType{})
	default:
		err = errors.New("not a valid ast.Node type")
	}
	return
}
