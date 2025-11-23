-- +goose Up

CREATE TABLE IF NOT EXISTS stock_price_data (
    id SERIAL PRIMARY KEY,
    stock_header_id INT REFERENCES stock_headers(id) ON DELETE CASCADE,

    nse_year_low_price DOUBLE PRECISION,
    nse_year_high_price DOUBLE PRECISION,
    bse_year_low_price DOUBLE PRECISION,
    bse_year_high_price DOUBLE PRECISION
);
