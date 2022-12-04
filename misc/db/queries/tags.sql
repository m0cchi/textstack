
-- name: ListTagsByUUID :many
SELECT * FROM tags WHERE text_uuid = $1;

-- name: ListTagsByName :many
SELECT * FROM tags WHERE name = $1;

-- name: CreateTag :exec
INSERT INTO tags(text_uuid, name) VALUES ($1, $2);


