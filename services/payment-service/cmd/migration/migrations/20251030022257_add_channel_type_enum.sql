-- +goose Up
-- +goose StatementBegin
CREATE TYPE channel_type AS ENUM ('material', 'sparepart', 'rent_heavy_equipment');
ALTER TABLE payment_order ADD COLUMN channel channel_type;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE payment_order DROP COLUMN IF EXISTS channel;
DROP TYPE IF EXISTS channel_type;
-- +goose StatementEnd
