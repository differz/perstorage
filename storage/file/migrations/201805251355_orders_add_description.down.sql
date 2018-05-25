CREATE TABLE orders_temp
(
  id INTEGER PRIMARY KEY NOT NULL,
  customer_id INTEGER NOT NULL,
  order_date TIMESTAMP,
  expire_date TIMESTAMP
);
INSERT INTO orders_temp(id, customer_id, order_date, expire_date) SELECT id, customer_id, order_date, expire_date FROM orders;
DROP TABLE orders;
ALTER TABLE orders_temp RENAME TO orders;
