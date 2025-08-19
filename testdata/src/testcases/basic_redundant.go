// Basic case: redundant alias with no naming conflicts
package testcases

import (
	f "fmt" // want "redundant import alias: f \"fmt\""
	"os"
)

func basicRedundant() {
	f.Println("hello")
}