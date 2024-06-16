# Postgres exercises


```sql
CREATE TABLE IF NOT EXISTS tusers (
    id SERIAL PRIMARY KEY,
    name TEXT,
    email TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS tweets (
    id SERIAL PRIMARY KEY,
    text TEXT,
    user_id INT NOT NULL,
    likes INT
);
CREATE TABLE IF NOT EXISTS likes (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    tweet_id INT NOT NULL
);
```

INSERT INTO tusers(name, email) VALUES ('VL', 'vl@chammy.info') RETURNING id;
INSERT INTO tweets(text, user_id) VALUES ('VL', 'vl@chammy.info') RETURNING id;
INSERT INTO likes(user_id, tweet_id) VALUES ($1, $2) RETURNING id;
