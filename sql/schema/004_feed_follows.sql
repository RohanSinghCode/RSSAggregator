-- +goose UP

CREATE TABLE feed_follows (
    Id UUID PRIMARY KEY,
    CreatedAt TIMESTAMP NOT NULL,
    UpdatedAt TIMESTAMP,
    user_id UUID NOT NULL REFERENCES Users(Id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES Feeds(Id) ON DELETE CASCADE,
    UNIQUE(user_id, feed_id)
);

-- +goose Down

DROP TABLE feed_follows;