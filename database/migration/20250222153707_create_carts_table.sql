-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS carts (
                       id bigserial PRIMARY KEY,
                       user_id bigint REFERENCES users(id) ON DELETE CASCADE,
                       created_at timestamp DEFAULT current_timestamp,
                       updated_at timestamp DEFAULT current_timestamp,
                       deleted_at timestamp DEFAULT NULL
);
CREATE INDEX idx_carts_user_id ON carts(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS carts;
-- +goose StatementEnd
