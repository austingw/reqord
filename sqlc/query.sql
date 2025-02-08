-- name: ListAllRequests :many
SELECT *, (SELECT COUNT(*) FROM requests) AS total_count) 
FROM requests
ORDER BY name;

-- name: ListProjectRequests :many
SELECT *, (SELECT COUNT(*) FROM requests) AS total_count 
FROM requests
WHERE project_id = ?
ORDER BY name;

-- name: GetRequest :one
SELECT * FROM requests
WHERE project_id = ? AND name = ? LIMIT 1;

-- name: CreateRequest :one
INSERT INTO requests (
  project_id,
  name,
  curl,
  method,
  url,
  headers,
  body
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateRequest :exec
UPDATE requests
set name = ?,
WHERE id = ?;

-- name: DeleteRequest :exec
DELETE FROM requests
WHERE id = ?;

-- name: ListAllProjects :many
SELECT *, (SELECT COUNT(*) FROM requests) AS total_count) 
FROM projects
ORDER BY name;

-- name: CreateProject :one
INSERT INTO projects (name) VALUES (?)
RETURNING *;

-- name: UpdateProject :exec
UPDATE projects
set name = ?
WHERE id = ?;

-- name: DeleteProject :exec
DELETE FROM projects
WHERE id = ?;

-- name: SelectProject :exec
INSERT OR REPLACE INTO selected_project (rowid, project_id) VALUES (1, ?);
