package testcases

import (
	_ "unsafe"
	_test "testing"
	__internal "internal/reflectlite"
	normal "os" // want "redundant import alias: normal \"os\""
)

func testFunction() {
	normal.Exit(0)
}