-- +goose Up

CREATE TABLE IF NOT EXISTS stock_headers (
    id SERIAL PRIMARY KEY,
    search_id TEXT,
    groww_company_id TEXT,
    isin TEXT,
    industry_id BIGINT,
    industry_name TEXT,
    display_name TEXT,
    short_name TEXT,
    type TEXT,
    is_fno_enabled BOOLEAN,
    nse_script_code TEXT,
    bse_script_code TEXT,
    nse_trading_symbol TEXT,
    bse_trading_symbol TEXT,
    is_bse_tradable BOOLEAN,
    is_nse_tradable BOOLEAN,
    logo_url TEXT,
    floating_shares BIGINT,
    is_bse_fno_enabled BOOLEAN,
    is_nse_fno_enabled BOOLEAN
);
