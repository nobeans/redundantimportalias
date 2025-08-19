// Multiple redundant aliases with no conflicts
package testcases

import (
	ctx "context" // want "redundant import alias: ctx \"context\""
	str "strings" // want "redundant import alias: str \"strings\""
	t "time"      // want "redundant import alias: t \"time\""
)

func testFunction() {
	ctx.Background()
	str.ToUpper("test")
	t.Now()
}