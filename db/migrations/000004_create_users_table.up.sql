CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE,
    password VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE
);