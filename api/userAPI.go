package api

import (
	"fmt"
	"net/http"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is user page")
}
