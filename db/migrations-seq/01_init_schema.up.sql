CREATE TABLE IF NOT EXISTS customer(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email_id VARCHAR(100) NOT NULL,
    access_key VARCHAR(100) NOT NULL,
    plan_id INTEGER NOT NULL,
    secret_key VARCHAR(100) NOT NULL,
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,   
	updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS plan(
    id SERIAL PRIMARY KEY,
    plan_name VARCHAR(100) NOT NULL,
    daily_base_cost INTEGER NOT NULL,
	api_pricing FLOAT(10) NOT NULL,
	storage_price FLOAT(10) NOT NULL
);

INSERT INTO plan (plan_name, daily_base_cost, api_pricing, storage_price) VALUES ('basic', 10, 0.1, 0.1);
INSERT INTO plan (plan_name, daily_base_cost, api_pricing, storage_price) VALUES ('advanced', 15, 0.05, 0.05);
INSERT INTO plan (plan_name, daily_base_cost, api_pricing, storage_price) VALUES ('enterprise', 20, 0.1, 0.01);