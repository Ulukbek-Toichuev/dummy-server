package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", handlerGet)
	http.ListenAndServe(":8080", nil)
}

func handlerGet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Server", "Dummy server")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
