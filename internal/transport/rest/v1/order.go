package resthendler

import (
	"encoding/json"
	"net/http"
	"path"
	orderservice "test-module/internal/service/order"
)

func OrderHendler(w http.ResponseWriter, r *http.Request) {
	orderUID := path.Base(r.URL.Path)
	order, err := orderservice.GetOrder(orderUID)
	if err != nil {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	orderJsonBytes, _ := json.Marshal(order)
	w.Write(orderJsonBytes)
}
