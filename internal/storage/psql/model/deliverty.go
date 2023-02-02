package psqlmodel

import "test-module/internal/domain"

type Delivery struct {
	Name    string
	Phone   string
	Zip     string
	City    string
	Address string
	Region  string
	Email   string
}

func (modelDelivery Delivery) ToDomain() domain.Delivery {
	return domain.Delivery{
		Name:    modelDelivery.Name,
		Phone:   modelDelivery.Phone,
		Zip:     modelDelivery.Zip,
		City:    modelDelivery.City,
		Address: modelDelivery.Address,
		Region:  modelDelivery.Region,
		Email:   modelDelivery.Email,
	}
}

func NewDelivery(domainDelivery domain.Delivery) (modelDelivery Delivery) {
	return Delivery{
		Name:    domainDelivery.Name,
		Phone:   domainDelivery.Phone,
		Zip:     domainDelivery.Zip,
		City:    domainDelivery.City,
		Address: domainDelivery.Address,
		Region:  domainDelivery.Region,
		Email:   domainDelivery.Email,
	}
}
