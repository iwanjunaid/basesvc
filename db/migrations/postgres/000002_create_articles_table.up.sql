CREATE TABLE IF NOT EXISTS articles(
  id          UUID PRIMARY KEY,
  author_id   UUID,
  title       VARCHAR(255) UNIQUE NOT NULL, 
  content     TEXT NOT NULL,
  created_at  TIMESTAMP NOT NULL,
  updated_at  TIMESTAMP DEFAULT NULL,
  deleted_at  TIMESTAMP DEFAULT NULL,
  CONSTRAINT FK_AAI_AID FOREIGN KEY(author_id) REFERENCES authors(id)
);