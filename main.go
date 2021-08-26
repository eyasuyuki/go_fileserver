package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir("."))
	addr := r.RemoteAddr
	method := r.Method
	path := r.URL.Path
	agent := r.UserAgent()
	fmt.Printf("Remote: %s [%s] Path=%s Agent=%s\n", addr, method, path, agent)
	fileServer.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
