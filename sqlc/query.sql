-- name: GetProfile :one    
SELECT * from profiles
WHERE id = $1
LIMIT 1;

-- name: CreateProfile :one
INSERT INTO profiles (name, last_name, age)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateProfile :one
UPDATE profiles
    SET name = $2,
    last_name = $3,
    age = $4
WHERE id = $1
RETURNING *;

-- name: PatchProfile :one
UPDATE profiles
    SET name = COALESCE($2, name),
    last_name = COALESCE($3, last_name),
    age = COALESCE($4, age)
WHERE id = $1
RETURNING *;

-- name: DeleteProfile :one
DELETE FROM profiles
WHERE id = $1
RETURNING id;