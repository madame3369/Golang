package myapp

import "net/http"

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	return mux
}

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)
}
