-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS custom;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA IF EXISTS custom CASCADE;
-- +goose StatementEnd