-- +goose Up

CREATE TABLE IF NOT EXISTS live_price_dtos (
    id SERIAL PRIMARY KEY,
    type TEXT,
    symbol TEXT,
    ts_in_millis INTEGER,
    open REAL,
    high REAL,
    low REAL,
    close REAL,
    ltp REAL,
    day_change REAL,
    day_change_perc REAL,
    low_price_range REAL,
    high_price_range REAL,
    volume BIGINT,
    total_buy_qty REAL,
    total_sell_qty REAL,
    oi_day_change REAL,
    oi_day_change_perc REAL,
    last_trade_qty INTEGER,
    last_trade_time INTEGER
);

CREATE TABLE IF NOT EXISTS stocks (
    id SERIAL PRIMARY KEY,
    isin TEXT,
    groww_contract_id TEXT,
    company_name TEXT,
    company_short_name TEXT,
    search_id TEXT,
    industry_code INTEGER,
    bse_script_code INTEGER,
    nse_script_code TEXT,
    yearly_high_price REAL,
    yearly_low_price REAL,
    close_price REAL,
    market_cap BIGINT,

    live_price_dto_id INTEGER,
    CONSTRAINT fk_livepricedto FOREIGN KEY (live_price_dto_id)
        REFERENCES live_price_dtos(id) ON DELETE SET NULL
);
