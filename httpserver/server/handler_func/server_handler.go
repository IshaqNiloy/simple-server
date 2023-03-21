package handler_func

import (
	"io"
	"net/http"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
	return
}
