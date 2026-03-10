-- +goose Up
CREATE TABLE feeds (
    id uuid PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    name VARCHAR(100) NOT NULL,
    url VARCHAR(255) NOT NULL UNIQUE,
    user_id uuid NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
