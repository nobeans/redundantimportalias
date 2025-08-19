// Different packages with same base name
package testcases

import (
	http1 "net/http"
	http2 "github.com/example/http"
)

func testFunction() {
	http1.Get("url1")
	http2.Get("url2")
}