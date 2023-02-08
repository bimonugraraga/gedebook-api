ALTER TABLE books 
DROP COLUMN IF EXISTS published_status;
ALTER TABLE chapters
DROP COLUMN IF EXSITS published_status;