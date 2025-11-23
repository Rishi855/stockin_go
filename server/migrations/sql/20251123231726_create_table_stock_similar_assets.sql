-- +goose Up

CREATE TABLE IF NOT EXISTS stock_similar_assets (
    id SERIAL PRIMARY KEY,
    stock_header_id INT REFERENCES stock_headers(id) ON DELETE CASCADE,

    similar_assets TEXT
);
