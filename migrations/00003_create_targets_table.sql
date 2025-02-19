-- +goose Up
CREATE TABLE targets
(
    id           SERIAL PRIMARY KEY,
    mission_id   INT REFERENCES missions (id) ON DELETE CASCADE,
    name         VARCHAR(100) NOT NULL,
    country      VARCHAR(100) NOT NULL,
    notes        TEXT,
    is_completed BOOLEAN   DEFAULT FALSE,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

-- +goose Down
DROP TABLE targets;