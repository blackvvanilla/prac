package handlers

import (
	"fmt"
	"net/http"
)

func rootHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "This is the Homepage, Welcome %s!", r.URL.Path[1:])
}
