CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP,
  order_name VARCHAR(255) UNIQUE NOT NULL,
  customer_id VARCHAR(255) REFERENCES customers(user_id)
);
