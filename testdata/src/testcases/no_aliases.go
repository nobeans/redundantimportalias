// Imports without aliases
// Should not report anything
package testcases

import (
	"fmt"
	"os"
	"strings"
)

func testFunction() {
	fmt.Println("hello")
	os.Exit(0)
	strings.ToUpper("test")
}