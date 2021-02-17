package httphandler

import (
	"fmt"
    "net/http"
)

func defaultHttpHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world!\n")
}