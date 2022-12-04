BEGIN;

DROP INDEX tags__name;
DROP INDEX tags__text_uuid;
DROP TABLE tags;

COMMIT;
