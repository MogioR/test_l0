package domain

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int32  `json:"amount"`
	PaymentDt    int32  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int32  `json:"delivery_cost"`
	GoodsTotal   int32  `json:"goods_total"`
	CustomFee    int32  `json:"custom_fee"`
}
