package AstUtils

import (
	"go/ast"
	"go/token"
)

// UnnestStruct Unnest structs that are contained inside other structs. If a name is given, only structs that are
// embedded in the named one are considered otherwise all structs inside the file.
func UnnestStruct(structName *string, file *ast.File) {
	var foundNodes []*FoundNodes
	var completed = false
	// Find all structs that are embedded inside another struct. This includes structs that are inside another struct
	//and part of map, channels etc. For example chan Example struct{}, is externalized as well
	SearchNodes(file, &foundNodes, []*ast.Node{}, func(n *ast.Node, parents []*ast.Node, completed *bool) bool {
		if _, ok := (*n).(*ast.StructType); ok && len(parents) > 0 {
			for _, parent := range parents {
				if structName == nil {
					return true
				} else if par, ok := (*parent).(*ast.TypeSpec); ok && len(parents) > 0 && par.Name.Name == *structName {
					return true
				}
			}

		}
		return false
	}, &completed)

	func() {
		for _, node := range foundNodes {
			func() {
				//Search for parent struct. If found, replace inline struct whit newly generated struct type
				for _, parent := range node.Parents {
					if _, ok := (*parent).(*ast.StructType); ok {
						var name string
						for _, parent1 := range node.Parents {
							if g, ok := (*parent1).(*ast.Field); ok {
								name = g.Names[0].Name
							}
						}
						v := &ast.GenDecl{
							Tok: token.TYPE,
							Specs: []ast.Spec{
								&ast.TypeSpec{
									Name: &ast.Ident{
										Name: name,
									},
									Type: (*node.Node).(*ast.StructType),
								},
							},
						}
						t := &ast.StarExpr{
							X: &ast.Ident{
								Name: name,
							},
						}

						file.Decls = append(file.Decls, v)
						ReplaceExprChild(node.Parents[0], t)
						return
					}
				}
			}()
		}
	}()
}

func ReplaceExprChild(decl *ast.Node, n ast.Expr) {
	switch (*decl).(type) {
	case *ast.BadExpr:
		return
	case *ast.Ident:
		return
	case *ast.Ellipsis:
		return
	case *ast.BasicLit:
		return
	case *ast.FuncLit:
		return
	case *ast.CompositeLit:
		return
	case *ast.ParenExpr:
		return
	case *ast.SelectorExpr:
		return
	case *ast.IndexExpr:
		return
	case *ast.IndexListExpr:
		return
	case *ast.SliceExpr:
		return
	case *ast.TypeAssertExpr:
		return
	case *ast.CallExpr:
		return
	case *ast.StarExpr:
		if (*decl).(*ast.StarExpr) == nil {
			return
		}
		(*decl).(*ast.StarExpr).X = n
	case *ast.UnaryExpr:
		if (*decl).(*ast.UnaryExpr) == nil {
			return
		}
		(*decl).(*ast.UnaryExpr).X = n
	case *ast.BinaryExpr:
		if (*decl).(*ast.BinaryExpr) == nil {
			return
		}
		(*decl).(*ast.BinaryExpr).Y = n
	case *ast.KeyValueExpr:
		if (*decl).(*ast.KeyValueExpr) == nil {
			return
		}
		(*decl).(*ast.KeyValueExpr).Value = n
	case *ast.ArrayType:
		if (*decl).(*ast.ArrayType) == nil {
			return
		}
		(*decl).(*ast.ArrayType).Elt = n
	case *ast.StructType:
		return
	case *ast.FuncType:
		return
	case *ast.InterfaceType:
		return
	case *ast.MapType:
		if (*decl).(*ast.MapType) == nil {
			return
		}
		(*decl).(*ast.MapType).Value = n
	case *ast.ChanType:
		if (*decl).(*ast.ChanType) == nil {
			return
		}
		(*decl).(*ast.ChanType).Value = n
	case *ast.Field:
		if (*decl).(*ast.Field) == nil {
			return
		}
		(*decl).(*ast.Field).Type = n
	}
}
