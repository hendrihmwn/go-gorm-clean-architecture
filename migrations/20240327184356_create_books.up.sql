CREATE TABLE IF NOT EXISTS books(
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR,
    author VARCHAR,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);