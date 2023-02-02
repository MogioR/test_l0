SELECT
orders.order_uid,
orders.track_number,
orders.entry,

deliveries.name,
deliveries.phone,
deliveries.zip,
deliveries.city,
deliveries.address,
deliveries.region,
deliveries.email,

payments.transaction,
payments.request_id,
payments.currency,
payments.provider,
payments.amount,
payments.payment_dt,
payments.bank,
payments.delivery_cost,
payments.goods_total,
payments.custom_fee,

orders.locale,
orders.internal_signature,
orders.customer_id,
orders.delivery_service,
orders.shardkey,
orders.sm_id,
orders.date_created,
orders.oof_shard

FROM orders 
inner join deliveries on orders.delivery = deliveries.delivery_id
inner join payments on orders.payment = payments.payment_id
WHERE orders.order_uid = $1;