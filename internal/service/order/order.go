package orderservice

import (
	"errors"
	"test-module/internal/domain"

	"github.com/jackc/pgx/v5/pgconn"
)

type Cache interface {
	Get(orderUID string) (item interface{}, err error)
	Set(orderUID string, item interface{}) error
}

type Repository interface {
	SaveOrder(order domain.Order) (err error)
	GetOrder(orderUID string) (domainOrder domain.Order, err error)
	GetAllOrders() (domainOrders []domain.Order, err error)
	CreateTables() (err error)
}

type OrderService struct {
	cache      Cache
	repository Repository
}

var (
	Instanse *OrderService
)

func Create(cache Cache, repository Repository) (err error) {
	Instanse = &OrderService{
		cache:      cache,
		repository: repository,
	}
	err = Instanse.initCache()
	return err
}

func GetOrder(orderUID string) (order domain.Order, err error) {
	if Instanse.cache == nil || Instanse.repository == nil {
		return order, errors.New("OrderService not created!")
	}

	order_raw, err := Instanse.cache.Get(orderUID)
	order, ok := order_raw.(domain.Order)
	if err == nil && ok == true {
		return order, nil
	}

	order, err = Instanse.repository.GetOrder(orderUID)
	if err != nil {
		return order, err
	}
	err = Instanse.cache.Set(order.Uuid, order)
	return order, err
}

func SetOrder(order domain.Order) (err error) {
	if Instanse.cache == nil || Instanse.repository == nil {
		return errors.New("OrderService not created!")
	}
	err = Instanse.cache.Set(order.Uuid, order)
	if err != nil {
		return err
	}
	err = Instanse.repository.SaveOrder(order)
	return err
}

func (service *OrderService) initCache() (err error) {
	orders, err := service.repository.GetAllOrders()
	if err != nil {
		pgError, ok := err.(*pgconn.PgError)
		if ok && pgError.Code == "42P01" {
			service.repository.CreateTables()
			if orders, err = service.repository.GetAllOrders(); err != nil {
				return err
			}
		} else {
			return err
		}

	}

	for _, order := range orders {
		err = service.cache.Set(order.Uuid, order)
	}
	return err
}
