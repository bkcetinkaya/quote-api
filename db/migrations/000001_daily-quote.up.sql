CREATE TABLE quotes (
  "id" bigserial PRIMARY KEY NOT NULL,
  "author" varchar DEFAULT 'Anonymous',
  "quote" varchar NOT NULL
);