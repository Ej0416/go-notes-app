-- name: AddUsers :exec
INSERT INTO users (email, first_name, last_name, password_hash)
VALUES ($1, $2, $3, $4);

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: GetUserByID :one
SELECT email, first_name, last_name, created_at FROM users 
WHERE id = $1;

-- name: GetUserAuth :one
SELECT id,email, password_hash, updated_at FROM users
WHERE email = $1;

-- name: UpdateUserInfo :one
UPDATE users
SET first_name = $1,
    last_name = $2,
    updated_at = now()
WHERE id = $3
RETURNING *;

-- name: ChangeUserEmail :one
UPDATE users
SET email = $1
WHERE id = $2
RETURNING *; 

-- name: DeleteUser :one
UPDATE users
SET is_active = FALSE
WHERE id = $1
RETURNING *;

-- name: CreateNote :exec
INSERT INTO notes (user_id, title, body)
VALUES ($1, $2, $3);

-- name: ListUserNotes :many
SELECT * FROM notes 
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: ListAllNotes :many
SELECT * FROM notes
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: GetNotesByID :one
SELECT * FROM notes
WHERE id = $1;

-- name: EditNotes :one
UPDATE notes
SET title = $1, 
    body = $2, 
    updated_at = now()
WHERE id = $3
RETURNING *; 

-- name: DeleteNotes :one
UPDATE notes
SET is_deleted = TRUE
RETURNING *;
