INSERT INTO payments(
    "transaction",
    "request_id",
    "currency",
    "provider",
    "amount",
    "payment_dt",
    "bank",
    "delivery_cost",
    "goods_total",
    "custom_fee"
) VALUES(
    $1, 
    $2, 
    $3, 
    $4, 
    $5, 
    $6, 
    $7,
    $8,
    $9,
    $10
) RETURNING payment_id;