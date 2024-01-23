CREATE TABLE customers (
  user_id VARCHAR(255) PRIMARY KEY,
  login VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  company_id INTEGER REFERENCES customer_companies(company_id),
  credit_cards TEXT[] NOT NULL
);
