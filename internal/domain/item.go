package domain

type Item struct {
	ChrtId      int32  `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int32  `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int32  `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int32  `json:"total_price"`
	NmId        int32  `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int32  `json:"status"`
}
