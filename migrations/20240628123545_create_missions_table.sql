-- +goose Up
CREATE TABLE missions
(
    id       SERIAL PRIMARY KEY,
    cat_id   INTEGER REFERENCES cats (id),
    complete BOOLEAN DEFAULT FALSE
);

-- +goose Down
DROP TABLE missions;