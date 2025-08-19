// Same package imported multiple times
package testcases

import (
	f1 "fmt"
	f2 "fmt"
)

func testFunction() {
	f1.Println("hello")
	f2.Printf("world")
}