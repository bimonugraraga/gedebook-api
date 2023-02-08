CREATE TABLE if NOT EXISTS chapters (
  id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  book_id INT NOT NULL,
  chapter_title VARCHAR(255) NOT NULL,
  chapter_text TEXT NOT NULL,
  chapter_cover TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (book_id)
    REFERENCES books(id)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON chapters FOR EACH ROW EXECUTE
PROCEDURE trigger_update_timestamp();