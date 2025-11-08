-- +goose Up
-- +goose StatementBegin
ALTER TABLE payment_method ADD COLUMN admin_fee BIGINT DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE payment_method DROP COLUMN IF EXISTS admin_fee;
-- +goose StatementEnd
