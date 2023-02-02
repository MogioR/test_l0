package psqlmodel

import (
	"test-module/internal/domain"
	"time"
)

type Order struct {
	Uuid              string
	TrackNumber       string
	Entry             string
	Delivery          Delivery
	Payment           Payment
	Item              []Item
	Locale            string
	InternalSignature string
	CustomerId        string
	DeliveryService   string
	Shardkey          string
	SmId              int32
	DateCreated       time.Time
	OofShard          string
}

func (modelOrder Order) ToDomain() domain.Order {
	return domain.Order{
		Uuid:              modelOrder.Uuid,
		TrackNumber:       modelOrder.TrackNumber,
		Entry:             modelOrder.Entry,
		Delivery:          modelOrder.Delivery.ToDomain(),
		Payment:           modelOrder.Payment.ToDomain(),
		Item:              itemToDomain(modelOrder.Item),
		Locale:            modelOrder.Locale,
		InternalSignature: modelOrder.InternalSignature,
		CustomerId:        modelOrder.CustomerId,
		DeliveryService:   modelOrder.DeliveryService,
		Shardkey:          modelOrder.Shardkey,
		SmId:              modelOrder.SmId,
		DateCreated:       modelOrder.DateCreated,
		OofShard:          modelOrder.OofShard,
	}
}

func NewOrder(domainOrder domain.Order) (modelOrder Order) {
	return Order{
		Uuid:              domainOrder.Uuid,
		TrackNumber:       domainOrder.TrackNumber,
		Entry:             domainOrder.Entry,
		Delivery:          NewDelivery(domainOrder.Delivery),
		Payment:           NewPayment(domainOrder.Payment),
		Item:              domainToItem(domainOrder.Item),
		Locale:            domainOrder.Locale,
		InternalSignature: domainOrder.InternalSignature,
		CustomerId:        domainOrder.CustomerId,
		DeliveryService:   domainOrder.DeliveryService,
		Shardkey:          domainOrder.Shardkey,
		SmId:              domainOrder.SmId,
		DateCreated:       domainOrder.DateCreated,
		OofShard:          domainOrder.OofShard,
	}
}

func itemToDomain(moduleItems []Item) []domain.Item {
	var domainItems []domain.Item

	for _, moduleItem := range moduleItems {
		domainItems = append(domainItems, domain.Item(moduleItem))
	}
	return domainItems
}

func domainToItem(domainItems []domain.Item) []Item {
	var moduleItems []Item

	for _, domainItem := range domainItems {
		moduleItems = append(moduleItems, NewlItem(domainItem))
	}
	return moduleItems
}
