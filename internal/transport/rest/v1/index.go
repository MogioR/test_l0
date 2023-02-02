package resthendler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"net/url"
	orderservice "test-module/internal/service/order"
)

var (
	index_template = template.Must(template.ParseFiles("web/index.html"))
)

func IndexHendler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	u, err := url.Parse(r.URL.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	params := u.Query()
	if orderUid := params.Get("orderUid"); orderUid != "" {
		order, err := orderservice.GetOrder(orderUid)
		if err != nil {
			index_template.Execute(w, 404)
		} else {
			orderJsonBytes, _ := json.Marshal(order)
			index_template.Execute(w, string(orderJsonBytes))
		}
	} else {
		index_template.Execute(w, "")
	}
}
