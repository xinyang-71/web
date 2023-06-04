//go:build e2e

package frameWork

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var h Server = &HttpServer{}

	h.AddRoute(http.MethodGet, "/user", func(ctx Context) {

	})
	//
	http.ListenAndServe(":8081", h)
	http.ListenAndServeTLS(":443", "", "", h)

	//
	h.Start(":8081")
}
