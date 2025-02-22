-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS categories (
                       id bigserial PRIMARY KEY,
                       name VARCHAR(255),
                       code VARCHAR(255) NOT NULL,
                       description VARCHAR(255),
                       created_at timestamp DEFAULT current_timestamp,
                       updated_at timestamp DEFAULT current_timestamp,
                       deleted_at timestamp DEFAULT NULL
);
CREATE UNIQUE INDEX idx_categories_unique_code ON categories(code) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd
