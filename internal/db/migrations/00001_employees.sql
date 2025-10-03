-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    deleted_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

insert into employees (name, email, phone) values ('従業員1', 'test1@example.com', '09012345678');
insert into employees (name, email, phone) values ('従業員2', 'test2@example.com', '09012345679');
insert into employees (name, email, phone) values ('従業員3', 'test3@example.com', '09012345680');
-- 検索用インデックス
CREATE INDEX IF NOT EXISTS idx_employees_name ON employees (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_employees_name;
DROP TABLE IF EXISTS employees;
-- +goose StatementEnd
