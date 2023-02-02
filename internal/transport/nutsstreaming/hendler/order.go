package nutshendler

import (
	"encoding/json"
	"log"
	"test-module/internal/domain"
	orderservice "test-module/internal/service/order"

	"github.com/nats-io/stan.go"
)

func HandleOrder(orderMSG *stan.Msg) {
	order := domain.Order{}

	err := json.Unmarshal(orderMSG.Data, &order)
	if err != nil {
		log.Println("Cannot unmarshal data from nats-streaming-server" + err.Error())
		return
	}
	if !order.Valid() {
		log.Println("Cannot validate incoming json\n" + err.Error())
		return
	}
	log.Println("Set order with Uid: ", order.Uuid)
	orderservice.SetOrder(order)
}
