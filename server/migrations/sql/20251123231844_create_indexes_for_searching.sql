-- +goose Up

CREATE INDEX IF NOT EXISTS idx_stock_headers_search_id ON stock_headers(search_id);
CREATE INDEX IF NOT EXISTS idx_stock_stats_header_id ON stock_stats(stock_header_id);
CREATE INDEX IF NOT EXISTS idx_stock_shp_header_id ON stock_share_holding_patterns(stock_header_id);
CREATE INDEX IF NOT EXISTS idx_stock_price_header_id ON stock_price_data(stock_header_id);
CREATE INDEX IF NOT EXISTS idx_stock_financial_header_id ON stock_financial_statements(stock_header_id);
CREATE INDEX IF NOT EXISTS idx_stock_similar_assets_header_id ON stock_similar_assets(stock_header_id);
