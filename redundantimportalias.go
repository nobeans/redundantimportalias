package main

import (
	"go/ast"
	"path"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var a = &analysis.Analyzer{
	Name:             "redundantimportalias",
	Doc:              "checks redundant import alias",
	Run:              run,
	RunDespiteErrors: true,
}

func main() {
	singlechecker.Main(a)
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			for _, f := range pass.Files {
				// 1つのファイルの中で、同一パッケージ名の回避以外にimport aliasを利用している箇所を検出する。
				foundBases := make(map[string]int)
				var impsWithName []*ast.ImportSpec
				for _, imp := range f.Imports {
					if imp.Name != nil {
						if strings.HasPrefix(imp.Name.String(), "_") { // アンダースコアで始まる名前(単独も含む)の場合は無視する。
							continue
						}
						impsWithName = append(impsWithName, imp)
					}
					foundBases[path.Base(imp.Path.Value)] += 1
				}
				for _, imp := range impsWithName {
					if foundBases[path.Base(imp.Path.Value)] <= 1 {
						pass.Reportf(imp.Pos(), "redundant import alias: %s %s", imp.Name.String(), imp.Path.Value)
					}
				}
			}
			return true
		})
	}
	return nil, nil
}
