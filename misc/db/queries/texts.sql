
-- name: GetText :one
SELECT * FROM texts WHERE uuid = $1;

-- name: CreateText :exec
INSERT INTO texts(uuid, title, body) VALUES (gen_random_uuid(), $1, $2);

