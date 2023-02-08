CREATE TABLE IF NOT EXISTS books(
  id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  user_id INT NOT NULL,
  title VARCHAR (255) NOT NULL,
  book_cover TEXT,
  type VARCHAR(255) NOT NULL,
  main_category_id INT NOT NULL,
  status VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id)
    REFERENCES users(id),
  FOREIGN KEY (main_category_id)
    REFERENCES categories(id)
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON books FOR EACH ROW EXECUTE
PROCEDURE trigger_update_timestamp();