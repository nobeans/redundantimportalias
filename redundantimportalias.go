package main

import (
	"go/ast"
	"path"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var analyzer = &analysis.Analyzer{
	Name:             "redundantimportalias",
	Doc:              "checks redundant import alias",
	Run:              run,
	RunDespiteErrors: true,
}

func main() {
	singlechecker.Main(analyzer)
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		// パッケージのベース名の出現回数と、別名付きインポートの情報を収集する。
		baseNameCounts := make(map[string]int)
		var impsWithName []*ast.ImportSpec
		for _, imp := range f.Imports {
			if imp.Path == nil {
				continue
			}

			if imp.Name != nil {
				if shouldIgnore(imp.Name.String()) {
					continue
				}
				impsWithName = append(impsWithName, imp)
			}

			baseNameCounts[baseName(imp)] += 1
		}

		// 1つのファイルの中で、同一パッケージ名の回避以外にimport aliasを利用している箇所を検出する。
		for _, imp := range impsWithName {
			if baseNameCounts[baseName(imp)] <= 1 {
				pass.Reportf(imp.Pos(), "redundant import alias: %s %s", imp.Name.String(), imp.Path.Value)
			}
		}
	}
	return nil, nil
}

func shouldIgnore(name string) bool {
	// ドットインポートやブランクインポートは、チェックの対象外とする。
	// また、接頭辞としてのアンダースコアをチェック対象外とする理由は、
	// 同一ファイル内における変数名とのコンフリクトを解消したい場合などに、
	// 元のパッケージ名の先頭にアンダースコアを付与して回避するために利用することを想定している。
	return name == "." || strings.HasPrefix(name, "_")
}

func baseName(imp *ast.ImportSpec) string {
	// パスから引用符を除去してベース名を返す。
	return path.Base(strings.Trim(imp.Path.Value, "\""))
}
