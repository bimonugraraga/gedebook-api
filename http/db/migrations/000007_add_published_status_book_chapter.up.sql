ALTER TABLE books
ADD COLUMN published_status VARCHAR(255) DEFAULT 'Draft' NOT NULL;
ALTER TABLE chapters
ADD COLUMN published_status VARCHAR(255) DEFAULT 'Draft' NOT NULL;