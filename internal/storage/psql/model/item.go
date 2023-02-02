package psqlmodel

import "test-module/internal/domain"

type Item struct {
	ChrtId      int32
	TrackNumber string
	Price       int32
	Rid         string
	Name        string
	Sale        int32
	Size        string
	TotalPrice  int32
	NmId        int32
	Brand       string
	Status      int32
}

func (modelItem Item) ToDomain() domain.Item {
	return domain.Item{
		ChrtId:      modelItem.ChrtId,
		TrackNumber: modelItem.TrackNumber,
		Price:       modelItem.Price,
		Rid:         modelItem.Rid,
		Name:        modelItem.Name,
		Sale:        modelItem.Sale,
		Size:        modelItem.Size,
		TotalPrice:  modelItem.TotalPrice,
		NmId:        modelItem.NmId,
		Brand:       modelItem.Brand,
		Status:      modelItem.Status,
	}
}

func NewlItem(domainItem domain.Item) (modelItem Item) {
	return Item{
		ChrtId:      domainItem.ChrtId,
		TrackNumber: domainItem.TrackNumber,
		Price:       domainItem.Price,
		Rid:         domainItem.Rid,
		Name:        domainItem.Name,
		Sale:        domainItem.Sale,
		Size:        domainItem.Size,
		TotalPrice:  domainItem.TotalPrice,
		NmId:        domainItem.NmId,
		Brand:       domainItem.Brand,
		Status:      domainItem.Status,
	}
}
