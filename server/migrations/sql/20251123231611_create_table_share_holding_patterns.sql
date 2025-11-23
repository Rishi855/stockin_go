-- +goose Up

CREATE TABLE IF NOT EXISTS stock_share_holding_patterns (
    id SERIAL PRIMARY KEY,
    stock_header_id INT REFERENCES stock_headers(id) ON DELETE CASCADE,

    period TEXT,
    promoters_individual DOUBLE PRECISION,
    promoters_government DOUBLE PRECISION,
    promoters_corporation DOUBLE PRECISION,
    mutual_funds DOUBLE PRECISION,
    other_domestic_institutions_insurance DOUBLE PRECISION,
    other_domestic_institutions_other_firms DOUBLE PRECISION,
    foreign_institutions DOUBLE PRECISION,
    retail_and_others DOUBLE PRECISION
);
