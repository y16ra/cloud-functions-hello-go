package hello

import (
	"fmt"
	"net/http"
)

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
