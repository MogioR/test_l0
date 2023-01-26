package core

type Item struct {
	Chrt_id      int32  `json:"chrt_id"`
	Track_number string `json:"track_number"`
	Price        int32  `json:"price"`
	Rid          string `json:"rid"`
	Name         string `json:"name"`
	Sale         int32  `json:"sale"`
	Size         string `json:"size"`
	Total_price  int32  `json:"total_price"`
	Nm_id        int32  `json:"nm_id"`
	Brand        string `json:"brand"`
	Status       int32  `json:"status"`
}
