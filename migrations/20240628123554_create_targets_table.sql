-- +goose Up
CREATE TABLE targets
(
    id         SERIAL PRIMARY KEY,
    mission_id INTEGER REFERENCES missions (id),
    name       VARCHAR(100) NOT NULL,
    country    VARCHAR(100),
    notes      TEXT,
    complete   BOOLEAN DEFAULT FALSE
);

-- +goose Down
DROP TABLE targets;