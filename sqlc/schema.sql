CREATE TABLE IF NOT EXISTS profiles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    age smallint NOT NULL CHECK(age BETWEEN 18 AND 100)
);

CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS login_info (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL REFERENCES accounts(id),
    password_hesh VARCHAR(60) NOT NULL
);