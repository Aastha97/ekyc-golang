CREATE TABLE IF NOT EXISTS customer(
    id VARCHAR(100) PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email_id VARCHAR(100) NOT NULL,
    access_key VARCHAR(100) NOT NULL,
    secret_key TEXT NOT NULL,
    created_time TIMESTAMP NOT NULL,
    updated_time TIMESTAMP NOT NULL
);