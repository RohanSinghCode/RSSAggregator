-- +goose UP
ALTER TABLE Feeds 
ALTER COLUMN UpdatedAt DROP NOT NULL;


-- +goose Down
ALTER TABLE feeds
ALTER COLUMN updatedat SET NOT NULL;