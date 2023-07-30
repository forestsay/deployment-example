package main

import (
	"net/http"
)

type handler struct{}

func (h *handler) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet || r.Method == http.MethodPost {
		_, err := writer.Write([]byte("Hello, world!"))
		if err != nil {
			return
		}
	}
	return
}

func main() {
	handler := &handler{}
	http.ListenAndServe(":8080", handler)
}
