BEGIN;

CREATE TABLE IF NOT EXISTS texts (
    id serial NOT NULL PRIMARY KEY,
    uuid uuid NOT NULL UNIQUE,
    title varchar(255) NOT NULL,
    body text NOT NULL
);

COMMIT;
