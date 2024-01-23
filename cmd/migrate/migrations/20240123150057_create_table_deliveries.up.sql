CREATE TABLE deliveries (
  id SERIAL PRIMARY KEY,
  order_item_id INTEGER REFERENCES order_items(id),
  delivered_quantity INTEGER NOT NULL
);
