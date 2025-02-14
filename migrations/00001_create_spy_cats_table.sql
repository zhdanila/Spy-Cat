-- +goose Up
CREATE TABLE spy_cats
(
    id                  SERIAL PRIMARY KEY,
    name                VARCHAR(100) NOT NULL,
    years_of_experience INT CHECK (years_of_experience >= 0),
    breed               VARCHAR(100) NOT NULL,
    salary              DECIMAL(10, 2) CHECK (salary >= 0),
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE spy_cats;