CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT,
    email TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    amount INT,
    description TEXT
);
INSERT INTO users(name, email) VALUES ('VL', 'vl@chammy.info');