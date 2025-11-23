-- +goose Up

CREATE TABLE IF NOT EXISTS stock_financial_statements (
    id SERIAL PRIMARY KEY,
    stock_header_id INT REFERENCES stock_headers(id) ON DELETE CASCADE,

    revenue_yearly JSONB,
    revenue_quarterly JSONB,
    profit_yearly JSONB,
    profit_quarterly JSONB,
    networth_yearly JSONB
);

