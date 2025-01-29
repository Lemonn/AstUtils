package AstSearch

import "go/ast"

// FoundNodes holds the information for each found node
type FoundNodes struct {
	Node    *ast.Node
	Parents []*ast.Node
}

// SearchNodes Searches the Ast-tree. The search function decides what's a match. foundNodes holds all matches including
// their parents. This allows to modify a method call inside a function, then traverses upwards to modify the containing
// function parameters as well.
// Set completed to true inside the search function, if the search should be terminated.
func SearchNodes(decl ast.Node, foundNodes *[]*FoundNodes, parents []*ast.Node, searchFunction func(node *ast.Node, parents []*ast.Node, completed *bool) bool, completed *bool) {
	if completed == nil {
		b := false
		completed = &b
	}
	if searchFunction(&decl, parents, completed) {
		*foundNodes = append(*foundNodes, &FoundNodes{
			Node:    &decl,
			Parents: parents,
		})
	}
	if *completed {
		return
	}
	parents = append([]*ast.Node{&decl}, parents...)
	switch decl.(type) {
	case *ast.Comment:
		if decl.(*ast.Comment) == nil {
			return
		}
		return
	case *ast.CommentGroup:
		if decl.(*ast.CommentGroup) == nil {
			return
		}
		if decl.(*ast.CommentGroup) == nil {
			return
		}
		return
	case *ast.Field:
		if decl.(*ast.Field) == nil {
			return
		}
		parents = append(parents, &decl)
		SearchNodes(decl.(*ast.Field).Doc, foundNodes, parents, searchFunction, completed)
		if decl.(*ast.Field).Names != nil {
			for i, _ := range decl.(*ast.Field).Names {
				SearchNodes(decl.(*ast.Field).Names[i], foundNodes, parents, searchFunction, completed)
			}
		}
		SearchNodes(decl.(*ast.Field).Type, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.Field).Tag, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.Field).Comment, foundNodes, parents, searchFunction, completed)
	case *ast.FieldList:
		if decl.(*ast.FieldList) == nil {
			return
		}
		if decl.(*ast.FieldList).List != nil {
			for i, _ := range decl.(*ast.FieldList).List {
				SearchNodes(decl.(*ast.FieldList).List[i], foundNodes, parents, searchFunction, completed)
			}
		}
	case *ast.BadExpr:
		if decl.(*ast.BadExpr) == nil {
			return
		}
		return
	case *ast.Ident:
		if decl.(*ast.Ident) == nil {
			return
		}
		return
	case *ast.Ellipsis:
		if decl.(*ast.Ellipsis) == nil {
			return
		}
		SearchNodes(decl.(*ast.Ellipsis).Elt, foundNodes, parents, searchFunction, completed)
	case *ast.BasicLit:
		if decl.(*ast.BasicLit) == nil {
			return
		}
		return
	case *ast.FuncLit:
		if decl.(*ast.FuncLit) == nil {
			return
		}
		SearchNodes(decl.(*ast.FuncLit).Type, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.FuncLit).Body, foundNodes, parents, searchFunction, completed)
	case *ast.CompositeLit:
		if decl.(*ast.CompositeLit) == nil {
			return
		}
		SearchNodes(decl.(*ast.CompositeLit).Type, foundNodes, parents, searchFunction, completed)
		if decl.(*ast.CompositeLit).Elts != nil {
			for i, _ := range decl.(*ast.CompositeLit).Elts {
				SearchNodes(decl.(*ast.CompositeLit).Elts[i], foundNodes, parents, searchFunction, completed)
			}
		}
	case *ast.ParenExpr:
		if decl.(*ast.ParenExpr) == nil {
			return
		}
		SearchNodes(decl.(*ast.ParenExpr).X, foundNodes, parents, searchFunction, completed)
	case *ast.SelectorExpr:
		if decl.(*ast.SelectorExpr) == nil {
			return
		}
		SearchNodes(decl.(*ast.SelectorExpr).X, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.SelectorExpr).Sel, foundNodes, parents, searchFunction, completed)
	case *ast.IndexExpr:
		if decl.(*ast.IndexExpr) == nil {
			return
		}
		SearchNodes(decl.(*ast.IndexExpr).X, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.IndexExpr).Index, foundNodes, parents, searchFunction, completed)
	case *ast.IndexListExpr:
		if decl.(*ast.IndexListExpr) == nil {
			return
		}
		SearchNodes(decl.(*ast.IndexListExpr).X, foundNodes, parents, searchFunction, completed)
		if decl.(*ast.IndexListExpr).Indices != nil {
			for i, _ := range decl.(*ast.IndexListExpr).Indices {
				SearchNodes(decl.(*ast.IndexListExpr).Indices[i], foundNodes, parents, searchFunction, completed)
			}
		}
	case *ast.SliceExpr:
		if decl.(*ast.SliceExpr) == nil {
			return
		}
		SearchNodes(decl.(*ast.SliceExpr).X, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.SliceExpr).Low, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.SliceExpr).High, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.SliceExpr).Max, foundNodes, parents, searchFunction, completed)
	case *ast.TypeAssertExpr:
		if decl.(*ast.TypeAssertExpr) == nil {
			return
		}
		SearchNodes(decl.(*ast.TypeAssertExpr).X, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.TypeAssertExpr).Type, foundNodes, parents, searchFunction, completed)
	case *ast.CallExpr:
		if decl.(*ast.CallExpr) == nil {
			return
		}
		SearchNodes(decl.(*ast.CallExpr).Fun, foundNodes, parents, searchFunction, completed)
		if decl.(*ast.CallExpr).Args != nil {
			for i, _ := range decl.(*ast.CallExpr).Args {
				SearchNodes(decl.(*ast.CallExpr).Args[i], foundNodes, parents, searchFunction, completed)
			}
		}
	case *ast.StarExpr:
		if decl.(*ast.StarExpr) == nil {
			return
		}
		SearchNodes(decl.(*ast.StarExpr).X, foundNodes, parents, searchFunction, completed)
	case *ast.UnaryExpr:
		if decl.(*ast.UnaryExpr) == nil {
			return
		}
		SearchNodes(decl.(*ast.UnaryExpr).X, foundNodes, parents, searchFunction, completed)
	case *ast.BinaryExpr:
		if decl.(*ast.BinaryExpr) == nil {
			return
		}
		SearchNodes(decl.(*ast.BinaryExpr).X, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.BinaryExpr).Y, foundNodes, parents, searchFunction, completed)
	case *ast.KeyValueExpr:
		if decl.(*ast.KeyValueExpr) == nil {
			return
		}
		SearchNodes(decl.(*ast.KeyValueExpr).Key, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.KeyValueExpr).Value, foundNodes, parents, searchFunction, completed)
	case *ast.ArrayType:
		if decl.(*ast.ArrayType) == nil {
			return
		}
		SearchNodes(decl.(*ast.ArrayType).Elt, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.ArrayType).Len, foundNodes, parents, searchFunction, completed)
	case *ast.StructType:
		if decl.(*ast.StructType) == nil {
			return
		}
		SearchNodes(decl.(*ast.StructType).Fields, foundNodes, parents, searchFunction, completed)
	case *ast.FuncType:
		if decl.(*ast.FuncType) == nil {
			return
		}
		SearchNodes(decl.(*ast.FuncType).TypeParams, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.FuncType).Params, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.FuncType).Results, foundNodes, parents, searchFunction, completed)
	case *ast.MapType:
		if decl.(*ast.MapType) == nil {
			return
		}
		SearchNodes(decl.(*ast.MapType).Key, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.MapType).Value, foundNodes, parents, searchFunction, completed)
	case *ast.ChanType:
		if decl.(*ast.ChanType) == nil {
			return
		}
		SearchNodes(decl.(*ast.ChanType).Value, foundNodes, parents, searchFunction, completed)
	case *ast.BadStmt:
		if decl.(*ast.BadStmt) == nil {
			return
		}
		return
	case *ast.DeclStmt:
		if decl.(*ast.DeclStmt) == nil {
			return
		}
		SearchNodes(decl.(*ast.DeclStmt).Decl, foundNodes, parents, searchFunction, completed)
	case *ast.EmptyStmt:
		if decl.(*ast.EmptyStmt) == nil {
			return
		}
		return
	case *ast.LabeledStmt:
		if decl.(*ast.LabeledStmt) == nil {
			return
		}
		SearchNodes(decl.(*ast.LabeledStmt).Label, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.LabeledStmt).Stmt, foundNodes, parents, searchFunction, completed)
	case *ast.ExprStmt:
		if decl.(*ast.ExprStmt) == nil {
			return
		}
		SearchNodes(decl.(*ast.ExprStmt).X, foundNodes, parents, searchFunction, completed)
	case *ast.SendStmt:
		if decl.(*ast.SendStmt) == nil {
			return
		}
		SearchNodes(decl.(*ast.SendStmt).Chan, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.SendStmt).Value, foundNodes, parents, searchFunction, completed)
	case *ast.IncDecStmt:
		if decl.(*ast.IncDecStmt) == nil {
			return
		}
		SearchNodes(decl.(*ast.IncDecStmt).X, foundNodes, parents, searchFunction, completed)
	case *ast.AssignStmt:
		if decl.(*ast.AssignStmt) == nil {
			return
		}
		if decl.(*ast.AssignStmt).Rhs != nil {
			for i, _ := range decl.(*ast.AssignStmt).Rhs {
				SearchNodes(decl.(*ast.AssignStmt).Rhs[i], foundNodes, parents, searchFunction, completed)
			}
		}
		if decl.(*ast.AssignStmt).Lhs != nil {
			for i, _ := range decl.(*ast.AssignStmt).Lhs {
				SearchNodes(decl.(*ast.AssignStmt).Lhs[i], foundNodes, parents, searchFunction, completed)
			}
		}
	case *ast.GoStmt:
		if decl.(*ast.GoStmt) == nil {
			return
		}
		SearchNodes(decl.(*ast.GoStmt).Call, foundNodes, parents, searchFunction, completed)
	case *ast.DeferStmt:
		if decl.(*ast.DeferStmt) == nil {
			return
		}
		SearchNodes(decl.(*ast.DeferStmt).Call, foundNodes, parents, searchFunction, completed)
	case *ast.ReturnStmt:
		if decl.(*ast.ReturnStmt) == nil {
			return
		}
		if decl.(*ast.ReturnStmt).Results != nil {
			for i, _ := range decl.(*ast.ReturnStmt).Results {
				SearchNodes(decl.(*ast.ReturnStmt).Results[i], foundNodes, parents, searchFunction, completed)
			}
		}
	case *ast.BranchStmt:
		if decl.(*ast.BranchStmt) == nil {
			return
		}
		SearchNodes(decl.(*ast.BranchStmt).Label, foundNodes, parents, searchFunction, completed)
	case *ast.BlockStmt:
		if decl.(*ast.BlockStmt) == nil {
			return
		}
		if decl.(*ast.BlockStmt).List != nil {
			for i, _ := range decl.(*ast.BlockStmt).List {
				SearchNodes(decl.(*ast.BlockStmt).List[i], foundNodes, parents, searchFunction, completed)
			}
		}
	case *ast.IfStmt:
		SearchNodes(decl.(*ast.IfStmt).Init, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.IfStmt).Cond, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.IfStmt).Body, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.IfStmt).Else, foundNodes, parents, searchFunction, completed)
	case *ast.CaseClause:
		if decl.(*ast.CaseClause) == nil {
			return
		}
		if decl.(*ast.CaseClause).List != nil {
			for i, _ := range decl.(*ast.CaseClause).List {
				SearchNodes(decl.(*ast.CaseClause).List[i], foundNodes, parents, searchFunction, completed)
			}
		}
		if decl.(*ast.CaseClause).Body != nil {
			for i, _ := range decl.(*ast.CaseClause).Body {
				SearchNodes(decl.(*ast.CaseClause).Body[i], foundNodes, parents, searchFunction, completed)
			}
		}
	case *ast.SwitchStmt:
		if decl.(*ast.SwitchStmt) == nil {
			return
		}
		SearchNodes(decl.(*ast.SwitchStmt).Init, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.SwitchStmt).Tag, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.SwitchStmt).Body, foundNodes, parents, searchFunction, completed)
	case *ast.TypeSwitchStmt:
		if decl.(*ast.TypeSwitchStmt) == nil {
			return
		}
		SearchNodes(decl.(*ast.TypeSwitchStmt).Init, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.TypeSwitchStmt).Assign, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.TypeSwitchStmt).Body, foundNodes, parents, searchFunction, completed)
	case *ast.CommClause:
		if decl.(*ast.CommClause) == nil {
			return
		}
		SearchNodes(decl.(*ast.CommClause).Comm, foundNodes, parents, searchFunction, completed)
		if decl.(*ast.CommClause).Body != nil {
			for i, _ := range decl.(*ast.CommClause).Body {
				SearchNodes(decl.(*ast.CommClause).Body[i], foundNodes, parents, searchFunction, completed)
			}
		}
	case *ast.SelectStmt:
		if decl.(*ast.SelectStmt) == nil {
			return
		}
		SearchNodes(decl.(*ast.SelectStmt).Body, foundNodes, parents, searchFunction, completed)
	case *ast.ForStmt:
		if decl.(*ast.ForStmt) == nil {
			return
		}
		SearchNodes(decl.(*ast.ForStmt).Init, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.ForStmt).Cond, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.ForStmt).Post, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.ForStmt).Body, foundNodes, parents, searchFunction, completed)
	case *ast.RangeStmt:
		if decl.(*ast.RangeStmt) == nil {
			return
		}
		SearchNodes(decl.(*ast.RangeStmt).Key, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.RangeStmt).Value, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.RangeStmt).X, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.RangeStmt).Body, foundNodes, parents, searchFunction, completed)
	case *ast.ImportSpec:
		if decl.(*ast.ImportSpec) == nil {
			return
		}
		SearchNodes(decl.(*ast.ImportSpec).Doc, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.ImportSpec).Name, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.ImportSpec).Path, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.ImportSpec).Comment, foundNodes, parents, searchFunction, completed)
	case *ast.ValueSpec:
		if decl.(*ast.ValueSpec) == nil {
			return
		}
		SearchNodes(decl.(*ast.ValueSpec).Doc, foundNodes, parents, searchFunction, completed)
		if decl.(*ast.ValueSpec).Names != nil {
			for i, _ := range decl.(*ast.ValueSpec).Names {
				SearchNodes(decl.(*ast.ValueSpec).Names[i], foundNodes, parents, searchFunction, completed)
			}
		}
		SearchNodes(decl.(*ast.ValueSpec).Type, foundNodes, parents, searchFunction, completed)
		if decl.(*ast.ValueSpec).Values != nil {
			for i, _ := range decl.(*ast.ValueSpec).Values {
				SearchNodes(decl.(*ast.ValueSpec).Values[i], foundNodes, parents, searchFunction, completed)
			}
		}
		SearchNodes(decl.(*ast.ValueSpec).Comment, foundNodes, parents, searchFunction, completed)
	case *ast.TypeSpec:
		if decl.(*ast.TypeSpec) == nil {
			return
		}
		SearchNodes(decl.(*ast.TypeSpec).Doc, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.TypeSpec).Name, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.TypeSpec).TypeParams, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.TypeSpec).Type, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.TypeSpec).Comment, foundNodes, parents, searchFunction, completed)
	case *ast.BadDecl:
		if decl.(*ast.BadDecl) == nil {
			return
		}
		return
	case *ast.GenDecl:
		if decl.(*ast.GenDecl) == nil {
			return
		}
		SearchNodes(decl.(*ast.GenDecl).Doc, foundNodes, parents, searchFunction, completed)
		if decl.(*ast.GenDecl).Specs != nil {
			for i, _ := range decl.(*ast.GenDecl).Specs {
				SearchNodes(decl.(*ast.GenDecl).Specs[i], foundNodes, parents, searchFunction, completed)
			}
		}
	case *ast.FuncDecl:
		if decl.(*ast.FuncDecl) == nil {
			return
		}
		SearchNodes(decl.(*ast.FuncDecl).Doc, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.FuncDecl).Recv, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.FuncDecl).Name, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.FuncDecl).Type, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.FuncDecl).Body, foundNodes, parents, searchFunction, completed)
	case *ast.File:
		if decl.(*ast.File) == nil {
			return
		}
		SearchNodes(decl.(*ast.File).Doc, foundNodes, parents, searchFunction, completed)
		SearchNodes(decl.(*ast.File).Name, foundNodes, parents, searchFunction, completed)
		if decl.(*ast.File).Decls != nil {
			for i, _ := range decl.(*ast.File).Decls {
				SearchNodes(decl.(*ast.File).Decls[i], foundNodes, parents, searchFunction, completed)
			}
		}
		if decl.(*ast.File).Imports != nil {
			for i, _ := range decl.(*ast.File).Imports {
				SearchNodes(decl.(*ast.File).Imports[i], foundNodes, parents, searchFunction, completed)
			}
		}
		if decl.(*ast.File).Unresolved != nil {
			for i, _ := range decl.(*ast.File).Unresolved {
				SearchNodes(decl.(*ast.File).Unresolved[i], foundNodes, parents, searchFunction, completed)
			}
		}
		if decl.(*ast.File).Comments != nil {
			for i, _ := range decl.(*ast.File).Comments {
				SearchNodes(decl.(*ast.File).Comments[i], foundNodes, parents, searchFunction, completed)
			}
		}
	}
}
