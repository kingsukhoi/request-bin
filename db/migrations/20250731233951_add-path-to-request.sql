-- migrate:up
ALTER TABLE requests ADD COLUMN path TEXT;

-- migrate:down
ALTER TABLE requests DROP COLUMN path;
