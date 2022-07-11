package types

import "github.com/zeromicro/zero-api/ast"

type Checker struct {
	Syntax *ast.SyntaxDecl
	Import []*ast.ImportSpec
}
