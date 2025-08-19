package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestRedundantImportAlias(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer, "testcases")
}