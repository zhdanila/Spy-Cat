-- +goose Up
CREATE TABLE missions
(
    id           SERIAL PRIMARY KEY,
    cat_id       INT UNIQUE REFERENCES spy_cats (id) ON DELETE SET NULL,
    is_completed BOOLEAN   DEFAULT FALSE,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE missions;