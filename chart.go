package chart

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// Name of go routine called is other node
// Edge is the "calls" Relation
// So ParentFn "calls" gostmt
type GoStmt struct {
	gostmt   *ast.GoStmt
	parentFn string
}

func extractGoStmnt(src string) []GoStmt {
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}
	var currFn string = "global"
	gostmts := make([]GoStmt, 0)

	//TODO: use channels to extract info out of Inspect into main goroutine
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GoStmt:
			gostmts = append(gostmts, GoStmt{gostmt: x, parentFn: currFn})
		case *ast.FuncDecl:
			currFn = x.Name.String()
		}
		return true
	})
	return gostmts
}
