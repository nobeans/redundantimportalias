package testcases

import (
	. "fmt"
	f "fmt" // want "redundant import alias: f \"fmt\""
)

func testFunction() {
	Println("using dot import")
	f.Printf("using alias")
}