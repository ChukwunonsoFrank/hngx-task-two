-- +goose Up

CREATE TABLE persons (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE persons;