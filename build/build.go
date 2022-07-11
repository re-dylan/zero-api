package build

import (
	"github.com/zeromicro/zero-api/ast"
	"github.com/zeromicro/zero-api/parser"
	"github.com/zeromicro/zero-api/token"
	"path/filepath"
	"strconv"
)

type context struct {
	fset *token.FileSet
	parsedFile map[string]*ast.File
	src string
}

type tree struct {
	root *ast.File
	nodes []*ast.File // for import api
}

type builder struct {
	fSet *token.FileSet
}

func (b *builder) Build(path string) error {

	f, err := parser.ParseFile(b.fSet, path,, "", nil)
	if err != nil {
		panic(err)
	}

}

func (b *builder) invoke(ctx context, path string) (*tree, error) {
	f, err := parser.ParseFile(ctx.fset, path, "", 0)
	if err != nil {
		return nil, err
	}

	if len(f.ImportDecls) > 0 {

	}
}

func (b *builder) invokeImportApi(ctx context, imports []*ast.GenDecl) ( error {
	for _, each := range imports {
		dir := filepath.Dir(ctx.src)
		for _, spec := range each.Specs {
			impPath, err := strconv.Unquote(spec.(*ast.ImportSpec).Path.Value)
			if err != nil {
				return err
			}
			if !filepath.IsAbs(impPath) {
				impPath = filepath.Join(dir, impPath)
			}

			if _, ok := ctx.parsedFile[impPath]; ok {
				continue
			}


		}
	}
}