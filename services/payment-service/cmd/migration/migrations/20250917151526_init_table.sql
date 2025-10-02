-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS "payment_bank" (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"name" VARCHAR(255),
	"code" VARCHAR(255),
	"icon" VARCHAR(255),
	"is_active" BOOLEAN,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMPTZ NULL DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS "idx_payment_bank_name" ON payment_bank USING btree ("name");
CREATE INDEX IF NOT EXISTS "idx_payment_bank_code" ON payment_bank USING btree ("code");

CREATE TABLE IF NOT EXISTS "payment_category" (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"name" VARCHAR(255),
	"code" VARCHAR(255),
	"sequence" INT,
	"is_active" BOOLEAN,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMPTZ NULL DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS "idx_payment_category_name" ON payment_category USING btree ("name");

CREATE TABLE IF NOT EXISTS "payment_provider" (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"name" VARCHAR(255),
	"code" VARCHAR(255),
	"is_active" BOOLEAN,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMPTZ NULL DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS "idx_payment_provider_name" ON payment_provider USING btree ("name");

CREATE TABLE IF NOT EXISTS "payment_method" (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"name" VARCHAR(255),
	"code" VARCHAR(255),
	"description" TEXT,
	"user_id" INT,
	"icon" VARCHAR(255),
	"category_id" INT,
	"bank_id" INT,
	"provider_id" INT,
	"is_active" BOOLEAN,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMPTZ NULL DEFAULT NULL,
	CONSTRAINT "fk_payment_category_payment_method"
		FOREIGN KEY("category_id")
			REFERENCES payment_category(id)
			ON DELETE CASCADE,
	CONSTRAINT "fk_payment_bank_payment_method"
		FOREIGN KEY("bank_id")
			REFERENCES payment_bank(id)
			ON DELETE CASCADE,
	CONSTRAINT "fk_payment_provider_payment_method"
		FOREIGN KEY("provider_id")
			REFERENCES payment_provider(id)
			ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_payment_method_name" ON payment_method USING btree ("name");
CREATE INDEX IF NOT EXISTS "idx_payment_method_code" ON payment_method USING btree ("code");

CREATE TABLE IF NOT EXISTS "payment_order" (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"order_id" VARCHAR(255) UNIQUE,
	"amount" DECIMAL(10, 2),
	"status" VARCHAR(50),
	"fraud_status" VARCHAR(50),
	"transaction_status" VARCHAR(50),
	"order_data" JSONB,
	"method_id" INT,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMPTZ NULL DEFAULT NULL,
	CONSTRAINT "fk_payment_method_payment_order"
		FOREIGN KEY("method_id")
			REFERENCES payment_method(id)
			ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS "idx_payment_order_order_id" ON payment_order USING btree ("order_id");

CREATE TABLE IF NOT EXISTS "payment_log" (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"order_id" INT,
	"provider_id" INT,
	"payment_url" TEXT,
	"status_code" INT,
	"request_data" JSONB,
	"response_data" JSONB,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMPTZ NULL DEFAULT NULL,
	CONSTRAINT "fk_payment_provider_payment_log"
		FOREIGN KEY("provider_id")
			REFERENCES payment_provider(id)
			ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "payment_log";
DROP TABLE IF EXISTS "payment_order";
DROP TABLE IF EXISTS "payment_method";
DROP TABLE IF EXISTS "payment_provider";
DROP TABLE IF EXISTS "payment_category";
DROP TABLE IF EXISTS "payment_bank";
-- +goose StatementEnd
