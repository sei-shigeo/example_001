-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS vehicles (
    id SERIAL PRIMARY KEY,
    number VARCHAR(100) NOT NULL UNIQUE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    deleted_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- 検索用インデックス
CREATE INDEX IF NOT EXISTS idx_vehicles_number ON vehicles (number);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS vehicles;
-- +goose StatementEnd
