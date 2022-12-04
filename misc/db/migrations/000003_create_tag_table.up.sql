BEGIN;

CREATE TABLE IF NOT EXISTS tags (
    id serial NOT NULL PRIMARY KEY,
    text_uuid uuid NOT NULL REFERENCES texts(uuid) ON DELETE CASCADE,
    name VARCHAR(25) NOT NULL
);

CREATE INDEX IF NOT EXISTS tags__text_uuid ON tags (
    text_uuid
);

CREATE INDEX IF NOT EXISTS tags__name ON "tags" (
    name
);

COMMIT;
