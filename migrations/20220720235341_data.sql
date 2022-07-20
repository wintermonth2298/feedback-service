-- +goose Up
CREATE TABLE feedback (
    id SERIAL PRIMARY KEY,
    rate int NOT NULL,
    text TEXT NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS feedback
