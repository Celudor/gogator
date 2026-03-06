-- +goose Up
CREATE TABLE users (
    id uuid PRIMARY KEY,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,
    name varchar(50) NOT NULL UNIQUE

);

-- +goose Down
DROP TABLE users;
