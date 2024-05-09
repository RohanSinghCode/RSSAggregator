-- +goose UP

CREATE TABLE Users (
    Id UUID PRIMARY KEY,
    CreatedAt TIMESTAMP NOT NULL,
    UpdatedAt TIMESTAMP NOT NULL,
    NAME TEXT NOT NULL
);

-- +goose Down

DROP TABLE Users