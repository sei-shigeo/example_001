-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    deleted_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

insert into customers (name, email, phone) values ('取引先1', '取引先1@example.com', '09012345678');
insert into customers (name, email, phone) values ('取引先2', '取引先2@example.com', '09012345679');
insert into customers (name, email, phone) values ('取引先3', '取引先3@example.com', '09012345680');

-- 検索用インデックス
CREATE INDEX IF NOT EXISTS idx_customers_name ON customers (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_customers_name;
DROP TABLE IF EXISTS customers;
-- +goose StatementEnd
