-- +goose Up
CREATE TABLE cats
(
    id                  SERIAL PRIMARY KEY,
    name                VARCHAR(100) NOT NULL,
    years_of_experience INTEGER      NOT NULL,
    breed               VARCHAR(100),
    salary              FLOAT
);

-- +goose Down
DROP TABLE cats;