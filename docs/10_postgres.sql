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
INSERT INTO users(name, email) 
VALUES ('VL', 'vl@chammy.info')
RETURNING id;
INSERT INTO orders (user_id, amount, description) 
VALUES ($1, $2, $3);
SELECT id, amount, description FROM orders
WHERE id=$1;
