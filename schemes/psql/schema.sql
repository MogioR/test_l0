CREATE TABLE deliveries(
    "delivery_id" SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL,
    "phone" VARCHAR(20) NOT NULL,
    "zip" VARCHAR(10) NOT NULL,
    "city" VARCHAR(50) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "region" VARCHAR(255) NOT NULL,
    "email" VARCHAR(50) NOT NULL
);

CREATE TABLE payments(
    "payment_id" SERIAL PRIMARY KEY,
    "transaction" VARCHAR(36) NOT NULL,
    "request_id" VARCHAR(36) NOT NULL,
    "currency" VARCHAR(3) NOT NULL,
    "provider" VARCHAR(50) NOT NULL,
    "amount" INTEGER NOT NULL,
    "payment_dt" INTEGER NOT NULL,
    "bank" VARCHAR(50) NOT NULL,
    "delivery_cost" INTEGER NOT NULL,
    "goods_total" INTEGER NOT NULL,
    "custom_fee" INTEGER NOT NULL
);

CREATE TABLE items(
    "item_id" SERIAL PRIMARY KEY,
    "chrt_id" INT NOT NULL,
    "track_number" VARCHAR(36) NOT NULL,
    "price" INTEGER NOT NULL,
    "rid" VARCHAR(36) NOT NULL,
    "name" VARCHAR(50) NOT NULL,
    "sale" INTEGER CHECK(sale >=0 AND sale <= 100) NOT NULL,
    "size" VARCHAR(10) NOT NULL,
    "total_price" INTEGER NOT NULL,
    "nm_id" INTEGER NOT NULL,
    "brand" VARCHAR(50) NOT NULL,
    "status" INTEGER NOT NULL
);

CREATE TABLE orders(
    "order_uid"  VARCHAR(36) PRIMARY KEY,
    "track_number"  VARCHAR(36) NOT NULL,
    "entry"  VARCHAR(36) NOT NULL,
    "delivery" INTEGER REFERENCES deliveries(delivery_id) ON DELETE CASCADE ON UPDATE CASCADE,
    "payment" INTEGER REFERENCES payments(payment_id) ON DELETE CASCADE ON UPDATE CASCADE,
    "locale"  VARCHAR(2) NOT NULL,
    "internal_signature"  VARCHAR(50) NOT NULL,
    "customer_id"  VARCHAR(36) NOT NULL,
    "delivery_service"  VARCHAR(50) NOT NULL,
    "shardkey"  VARCHAR(10) NOT NULL,
    "sm_id" INTEGER NOT NULL,
    "date_created" TIMESTAMP NOT NULL,
    "oof_shard"  VARCHAR(10) NOT NULL
);

CREATE TABLE orders_items(
    "order_id" VARCHAR(36) REFERENCES orders(order_uid) ON DELETE CASCADE ON UPDATE CASCADE,
    "item_id" INTEGER REFERENCES items(item_id) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (order_id, item_id)
);