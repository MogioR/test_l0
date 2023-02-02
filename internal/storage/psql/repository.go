package psqlrepo

import (
	"io/ioutil"
	"log"
	config "test-module/internal/config"
	"test-module/internal/domain"
	psqlbd "test-module/internal/storage/psql/bd"
	psqlmodel "test-module/internal/storage/psql/model"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db                      psqlbd.DB
	schema_sql              string
	insert_delivery_sql     string
	insert_payment_sql      string
	insert_item_sql         string
	insert_order_sql        string
	insert_order_item_sql   string
	select_order_sql        string
	select_order_items_sql  string
	select_orders_sql       string
	select_orders_items_sql string
}

func (r *Repository) Init(config config.Config) error {
	db, err := psqlbd.Connect(config)
	if err != nil {
		log.Printf("DB connecting error   #%v ", err)
	}
	r.db = *db

	r.schema_sql = load_file("schemes/psql/schema.sql")
	r.insert_delivery_sql = load_file("schemes/psql/insert_delivery.sql")
	r.insert_payment_sql = load_file("schemes/psql/insert_payment.sql")
	r.insert_item_sql = load_file("schemes/psql/insert_item.sql")
	r.insert_order_sql = load_file("schemes/psql/insert_order.sql")
	r.insert_order_item_sql = load_file("schemes/psql/insert_order_item.sql")
	r.select_order_sql = load_file("schemes/psql/select_order.sql")
	r.select_order_items_sql = load_file("schemes/psql/select_order_items.sql")
	r.select_orders_sql = load_file("schemes/psql/select_orders.sql")
	r.select_orders_items_sql = load_file("schemes/psql/select_orders_items.sql")

	return err
}

func (r *Repository) SaveOrder(domainOrder domain.Order) error {
	tx, err := r.db.Beginx()
	if err != nil {
		log.Printf("PSQL can't start transaction:%v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	order := psqlmodel.NewOrder(domainOrder)

	deliveryID, err := r.insert_delivery(tx, order.Delivery)
	if err != nil {
		log.Printf("bd.psql.insert.order.delivery.error %v\n", err)
		return err
	}

	paymentID, err := r.insert_payment(tx, order.Payment)
	if err != nil {
		log.Printf("bd.psql.insert.order.payment.error %v\n", err)
		return err
	}

	orderUID, err := r.insert_order(tx, order, deliveryID, paymentID)
	if err != nil {
		log.Printf("bd.psql.insert.order.order.error %v\n", err)
		return err
	}

	for _, item := range order.Item {
		itemID, err := r.insert_item(tx, item)
		if err != nil {
			log.Printf("bd.psql.insert.order.item.error %v\n", err)
			return err
		}
		err = r.insert_order_item(tx, orderUID, itemID)
		if err != nil {
			log.Printf("bd.psql.insert.order.ortder_item.error %v\n", err)
			return err
		}
	}

	return err
}

func (r *Repository) GetOrder(orderUID string) (domainOrder domain.Order, err error) {
	var order psqlmodel.Order
	err = r.db.QueryRow(
		r.select_order_sql,
		orderUID,
	).Scan(
		&order.Uuid,
		&order.TrackNumber,
		&order.Entry,

		&order.Delivery.Name,
		&order.Delivery.Phone,
		&order.Delivery.Zip,
		&order.Delivery.City,
		&order.Delivery.Address,
		&order.Delivery.Region,
		&order.Delivery.Email,

		&order.Payment.Transaction,
		&order.Payment.RequestId,
		&order.Payment.Currency,
		&order.Payment.Provider,
		&order.Payment.Amount,
		&order.Payment.PaymentDt,
		&order.Payment.Bank,
		&order.Payment.DeliveryCost,
		&order.Payment.GoodsTotal,
		&order.Payment.CustomFee,

		&order.Locale,
		&order.InternalSignature,
		&order.CustomerId,
		&order.DeliveryService,
		&order.Shardkey,
		&order.SmId,
		&order.DateCreated,
		&order.OofShard,
	)
	if err != nil {
		return domain.Order{}, err
	}

	rows, err := r.db.Query(
		r.select_order_items_sql,
		orderUID,
	)
	if err != nil {
		return order.ToDomain(), err
	}

	for rows.Next() {
		var item psqlmodel.Item
		err = rows.Scan(
			&item.ChrtId,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmId,
			&item.Brand,
			&item.Status,
		)
		if err != nil {
			return order.ToDomain(), err
		}
		order.Item = append(order.Item, item)
	}

	return order.ToDomain(), err
}

func (r *Repository) GetAllOrders() (domainOrders []domain.Order, err error) {
	rows, err := r.db.Query(
		r.select_orders_sql,
	)
	if err != nil {
		return domainOrders, err
	}

	var orders map[string]*psqlmodel.Order
	orders = make(map[string]*psqlmodel.Order)

	for rows.Next() {
		var order psqlmodel.Order
		rows.Scan(
			&order.Uuid,
			&order.TrackNumber,
			&order.Entry,

			&order.Delivery.Name,
			&order.Delivery.Phone,
			&order.Delivery.Zip,
			&order.Delivery.City,
			&order.Delivery.Address,
			&order.Delivery.Region,
			&order.Delivery.Email,

			&order.Payment.Transaction,
			&order.Payment.RequestId,
			&order.Payment.Currency,
			&order.Payment.Provider,
			&order.Payment.Amount,
			&order.Payment.PaymentDt,
			&order.Payment.Bank,
			&order.Payment.DeliveryCost,
			&order.Payment.GoodsTotal,
			&order.Payment.CustomFee,

			&order.Locale,
			&order.InternalSignature,
			&order.CustomerId,
			&order.DeliveryService,
			&order.Shardkey,
			&order.SmId,
			&order.DateCreated,
			&order.OofShard,
		)
		if err != nil {
			return domainOrders, err
		}
		orders[order.Uuid] = &order
	}

	rows, err = r.db.Query(
		r.select_orders_items_sql,
	)
	if err != nil {
		return domainOrders, err
	}

	for rows.Next() {
		var item psqlmodel.Item
		var orderUid string

		err = rows.Scan(
			&orderUid,
			&item.ChrtId,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmId,
			&item.Brand,
			&item.Status,
		)
		if err != nil {
			return domainOrders, err
		}
		orders[orderUid].Item = append(orders[orderUid].Item, item)
	}

	for k := range orders {
		domainOrders = append(domainOrders, orders[k].ToDomain())
	}

	return domainOrders, nil
}

func (r *Repository) CreateTables() (err error) {
	_, err = r.db.Exec(r.schema_sql)
	return err
}

func (r *Repository) insert_payment(tx *sqlx.Tx, payment psqlmodel.Payment) (paymentID int32, err error) {
	err = tx.QueryRow(
		r.insert_payment_sql,
		payment.Transaction,
		payment.RequestId,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.PaymentDt,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee,
	).Scan(&paymentID)

	return paymentID, err
}

func (r *Repository) insert_delivery(tx *sqlx.Tx, delivery psqlmodel.Delivery) (deliveryID int32, err error) {
	err = tx.QueryRow(
		r.insert_delivery_sql,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email,
	).Scan(&deliveryID)

	return deliveryID, err
}

func (r *Repository) insert_order(tx *sqlx.Tx, order psqlmodel.Order, deliveryID int32, paymentID int32) (orderUid string, err error) {
	err = tx.QueryRow(
		r.insert_order_sql,
		order.Uuid,
		order.TrackNumber,
		order.Entry,
		deliveryID,
		paymentID,
		order.Locale,
		order.InternalSignature,
		order.CustomerId,
		order.DeliveryService,
		order.Shardkey,
		order.SmId,
		order.DateCreated,
		order.OofShard,
	).Scan(&orderUid)

	return orderUid, err
}

func (r *Repository) insert_item(tx *sqlx.Tx, item psqlmodel.Item) (itemID int32, err error) {
	err = tx.QueryRow(
		r.insert_item_sql,
		item.ChrtId,
		item.TrackNumber,
		item.Price,
		item.Rid,
		item.Name,
		item.Sale,
		item.Size,
		item.TotalPrice,
		item.NmId,
		item.Brand,
		item.Status,
	).Scan(&itemID)

	return itemID, err
}

func (r *Repository) insert_order_item(tx *sqlx.Tx, orderUID string, itemID int32) (err error) {
	_, err = tx.Exec(
		r.insert_order_item_sql,
		orderUID,
		itemID,
	)

	return err
}

func load_file(source string) string {
	file_bytes, err := ioutil.ReadFile(source)
	if err != nil {
		log.Printf("repository load_file error:   #%v ", err)
	}
	return string(file_bytes)
}
