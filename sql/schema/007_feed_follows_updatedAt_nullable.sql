-- +goose UP
ALTER TABLE Feed_Follows 
ALTER COLUMN UpdatedAt DROP NOT NULL;


-- +goose Down
ALTER TABLE Feed_Follows
ALTER COLUMN updatedat SET NOT NULL;