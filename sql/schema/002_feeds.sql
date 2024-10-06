-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR NOT NULL,
    url VARCHAR NOt NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE feeds;

