CREATE TABLE IF NOT EXISTS users(
  id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  email VARCHAR (255) NOT NULL UNIQUE,
  password VARCHAR (255) NOT NULL,
  name VARCHAR (255) NOT NULL,
  profile TEXT,
  profile_picture TEXT,
  life_point BIGINT DEFAULT 500000,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE TRIGGER set_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE
PROCEDURE trigger_update_timestamp();