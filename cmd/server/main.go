package main

import (
	"net/http"

	"github.com/hddskull/practicus/internal/handler"
)

func start(mux *http.ServeMux) {
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/", handler.UpdateMetrics)

	start(mux)
}
