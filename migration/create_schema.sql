
CREATE TABLE IF NOT EXISTS customer(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email_id VARCHAR(100) NOT NULL,
    plan VARCHAR(100) NOT NULL,
    access_key VARCHAR(100) NOT NULL,
    secret_key VARCHAR(100) NOT NULL,
    created_time TIMESTAMP NOT NULL,
    updated_time TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS plan(
    id INT PRIMARY KEY AUTO_INCREMENT,
    plan_name VARCHAR(100) NOT NULL,
    daily_base_cost FLOAT(10) NOT NULL,
	match_score FLOAT(10) NOT NULL,
	storage_price FLOAT(10) NOT NULL,
);