package psqlmodel

import "test-module/internal/domain"

type Payment struct {
	Transaction  string ``
	RequestId    string ``
	Currency     string ``
	Provider     string ``
	Amount       int32  ``
	PaymentDt    int32  ``
	Bank         string ``
	DeliveryCost int32  ``
	GoodsTotal   int32  ``
	CustomFee    int32  ``
}

func (modelPayment Payment) ToDomain() domain.Payment {
	return domain.Payment{
		Transaction:  modelPayment.Transaction,
		RequestId:    modelPayment.RequestId,
		Currency:     modelPayment.Currency,
		Provider:     modelPayment.Provider,
		Amount:       modelPayment.Amount,
		PaymentDt:    modelPayment.PaymentDt,
		Bank:         modelPayment.Bank,
		DeliveryCost: modelPayment.DeliveryCost,
		GoodsTotal:   modelPayment.GoodsTotal,
		CustomFee:    modelPayment.CustomFee,
	}
}

func NewPayment(domainItem domain.Payment) (modelPayment Payment) {
	return Payment{
		Transaction:  domainItem.Transaction,
		RequestId:    domainItem.RequestId,
		Currency:     domainItem.Currency,
		Provider:     domainItem.Provider,
		Amount:       domainItem.Amount,
		PaymentDt:    domainItem.PaymentDt,
		Bank:         domainItem.Bank,
		DeliveryCost: domainItem.DeliveryCost,
		GoodsTotal:   domainItem.GoodsTotal,
		CustomFee:    domainItem.CustomFee,
	}
}
