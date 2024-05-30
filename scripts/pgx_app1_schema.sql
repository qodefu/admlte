
--@block
CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);
--@block
DESCRIBE authors

--@block
SELECT * FROM information_schema.columns
where table_name = 'authors'   

