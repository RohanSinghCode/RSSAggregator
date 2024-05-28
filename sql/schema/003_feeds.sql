-- +goose UP

CREATE TABLE Feeds (
    Id UUID PRIMARY KEY,
    CreatedAt TIMESTAMP NOT NULL,
    UpdatedAt TIMESTAMP,
    NAME TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID NOT NULL REFERENCES Users(Id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE Feeds;