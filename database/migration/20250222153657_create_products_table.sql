-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
                       id bigserial PRIMARY KEY,
                       name VARCHAR(255),
                       category_id bigint REFERENCES categories(id) ON DELETE CASCADE , 
                       description text , 
                       image text,
                       price int,
                       stock int , 
                       created_at timestamp DEFAULT current_timestamp,
                       updated_at timestamp DEFAULT current_timestamp,
                       deleted_at timestamp DEFAULT NULL
);
CREATE INDEX idx_products_category_id ON products(category_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
