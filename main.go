package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

func main(){
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset,"testdata/a.go",nil,0)
	if err != nil{
		panic(err)
	}

	for _, d := range f.Decls {
		switch d := d.(type){
		case *ast.FuncDecl:
			fmt.Println(d.Name)
			vars := map[string]string{}
			for _, b := range d.Body.List {
				switch b := b.(type){
				case *ast.AssignStmt:
					for i, l := range b.Lhs{
						switch c := l.(type) {
						case *ast.Ident:
							r := b.Rhs[i]
							switch d := r.(type) {
							case *ast.BasicLit:
								v, _ := strconv.Unquote(d.Value)
								if v == "" {
									vars = map[string]string{
										c.Name: v,
									}
								}
							case *ast.FuncLit:
								fmt.Println(d.Body)
							}
						}
					}
				case *ast.ReturnStmt:
					for _, e := range b.Results {
						switch r := e.(type) {
						case *ast.BasicLit:
							v, err := strconv.Unquote(r.Value)
							if err == nil && v == "" {
								fmt.Println("call empty")
							}
						case *ast.Ident:
							if vars[r.Name] == "" {
								fmt.Println("call empty")
							}
						}
					}
				}
			}
		}
	}
}
