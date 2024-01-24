CREATE TABLE IF NOT EXISTS customer_companies (
    company_id SERIAL PRIMARY KEY,
    company_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS customers (
    user_id SERIAL PRIMARY KEY,
    login VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    company_id INTEGER,
    credit_cards JSONB,
    FOREIGN KEY (company_id) REFERENCES customer_companies(company_id)
);

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    created_at timestamp,
    order_name VARCHAR(255),
    customer_id INTEGER,
    FOREIGN KEY (customer_id) REFERENCES customers(user_id)
);

CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER,
    price_per_unit DECIMAL(10,4),
    quantity INTEGER,
    product VARCHAR(255),
    FOREIGN KEY (order_id) REFERENCES orders(id)
);

CREATE TABLE IF NOT EXISTS deliveries (
    id SERIAL PRIMARY KEY,
    order_item_id INTEGER,
    delivered_quantity INTEGER,
    FOREIGN KEY (order_item_id) REFERENCES order_items(id)
);

INSERT INTO customer_companies (company_id, company_name) VALUES 
    (1, 'Roga & Kopyta'),
    (2, 'Pupkin & Co');

INSERT INTO customers (user_id, login, password, name, company_id, credit_cards) VALUES 
    (1, 'ivan', '12345', 'Ivan Ivanovich', 1, '["*****-1234", "*****-5678"]'),
    (2, 'petr', '54321', 'Petr Petrovich', 2, '["*****-4321", "*****-8765"]');

INSERT INTO orders (id, created_at, order_name, customer_id) VALUES 
    (1, '2020-01-02T15:34:12Z', 'PO #001-I', 1),
    (2, '2020-01-15T17:34:12Z', 'PO #002-I', 1),
    (3, '2020-01-05T05:34:12Z', 'PO #003-I', 1),
    (4, '2020-02-02T15:34:12Z', 'PO #004-I', 1),
    (5, '2020-01-03T10:34:12Z', 'PO #005-I', 1),
    (6, '2020-01-02T15:34:12Z', 'PO #001-P', 2),
    (7, '2020-01-15T17:34:12Z', 'PO #002-P', 2),
    (8, '2020-01-05T05:34:12Z', 'PO #003-P', 2),
    (9, '2020-02-02T15:34:12Z', 'PO #004-P', 2),
    (10, '2020-01-03T10:34:12Z', 'PO #005-P', 2);

INSERT INTO order_items (id, order_id, price_per_unit, quantity, product) VALUES 
    (1, 1, 1.3454, 10, 'Corrugated Box'),
    (2, 2, 23.14, 11, 'Corrugated Box'),
    (3, 3, 123.0345, 12, 'Corrugated Box'),
    (4, 4, NULL, 13, 'Corrugated Box'),
    (5, 5, 100, 14, 'Corrugated Box'),
    (6, 6, 1.5454, 15, 'Corrugated Box'),
    (7, 7, 25.14, 16, 'Corrugated Box'),
    (8, 8, 133.0345, 17, 'Corrugated Box'),
    (9, 9, 13.456, 18, 'Corrugated Box'),
    (10, 10, 110, 19, 'Corrugated Box'),
    (11, 1, 45.2334, 20, 'Hand sanitizer'),
    (12, 2, NULL, 21, 'Hand sanitizer'),
    (13, 3, 273.1234, 22, 'Hand sanitizer'),
    (14, 4, 11.45, 23, 'Hand sanitizer'),
    (15, 5, 12.467, 24, 'Hand sanitizer'),
    (16, 6, 11, 25, 'Hand sanitizer'),
    (17, 7, 123, 26, 'Hand sanitizer'),
    (18, 8, 173.1234, 27, 'Hand sanitizer'),
    (19, 9, 23.876, 28, 'Hand sanitizer'),
    (20, 10, 120, 29, 'Hand sanitizer');

INSERT INTO deliveries (id, order_item_id, delivered_quantity) VALUES 
    (1, 1, 5),
    (2, 2, 11),
    (3, 3, 12),
    (4, 4, 3),
    (5, 6, 15),
    (6, 7, 8),
    (7, 8, 3),
    (8, 16, 25),
    (9, 17, 26),
    (10, 18, 27),
    (11, 19, 28),
    (12, 20, 29),
    (13, 4, 5),
    (14, 8, 8),
    (15, 8, 6);