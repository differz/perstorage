CREATE TABLE orders_temp
(
  id          INTEGER PRIMARY KEY                 NOT NULL,
  customer_id INTEGER                             NOT NULL,
  category    INTEGER   DEFAULT 0,
  size        INTEGER   DEFAULT 0,
  description TEXT      DEFAULT '',
  order_date  TIMESTAMP DEFAULT current_timestamp NOT NULL,
  expire_date TIMESTAMP DEFAULT '2222-01-01'      NOT NULL
);
INSERT INTO orders_temp(id, customer_id, description, order_date, expire_date)
SELECT id, customer_id, description, order_date, expire_date
FROM orders;
DROP TABLE orders;
ALTER TABLE orders_temp RENAME TO orders;
