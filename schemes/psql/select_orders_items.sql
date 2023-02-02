SELECT
orders_items.order_id,
items.chrt_id,
items.track_number,
items.price,
items.rid,
items.name,
items.sale,
items.size,
items.total_price,
items.nm_id,
items.brand,
items.status
FROM orders_items
inner join items on orders_items.item_id = items.item_id
ORDER BY orders_items.order_id;