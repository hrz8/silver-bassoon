CREATE TABLE order_items (
  id SERIAL PRIMARY KEY,
  order_id INTEGER REFERENCES orders(id),
  price_per_unit NUMERIC,
  quantity INTEGER NOT NULL,
  product VARCHAR(255) NOT NULL
);
