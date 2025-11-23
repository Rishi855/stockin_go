-- +goose Up

CREATE TABLE IF NOT EXISTS stock_news (
    id SERIAL PRIMARY KEY,
    stock_news_id TEXT,
    stock_id INTEGER NOT NULL,

    title TEXT,
    summary TEXT,
    url TEXT,
    image_url TEXT,
    pub_date TIMESTAMP WITH TIME ZONE,
    source TEXT,

    CONSTRAINT fk_stocknews_stock
        FOREIGN KEY (stock_id)
        REFERENCES stocks(id)
        ON DELETE CASCADE
);
