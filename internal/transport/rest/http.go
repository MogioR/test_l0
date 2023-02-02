package rest

import (
	"fmt"
	"net/http"
	resthendler "test-module/internal/transport/rest/v1"
)

var (
	mux *http.ServeMux
)

func RegisterHendlers() {
	mux = http.NewServeMux()
	fs := http.FileServer(http.Dir("web/static"))

	mux.HandleFunc("/", resthendler.IndexHendler)
	mux.HandleFunc("/order/", resthendler.OrderHendler)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}

func StartServer(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
