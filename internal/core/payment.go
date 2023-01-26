package core

type Payment struct {
	Transaction   string `json:"transaction"`
	Request_id    string `json:"request_id"`
	Currency      string `json:"currency"`
	Provider      string `json:"provider"`
	Amount        int32  `json:"amount"`
	Payment_dt    int32  `json:"payment_dt"`
	Bank          string `json:"bank"`
	Delivery_cost int32  `json:"delivery_cost"`
	Goods_total   int32  `json:"goods_total"`
	Custom_fee    int32  `json:"custom_fee"`
}
