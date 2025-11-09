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
    SET name = COALESCE(sqlc.narg('name'), name),
    last_name = COALESCE(sqlc.narg('last_name'), last_name),
    age = COALESCE(sqlc.narg('age'), age)
WHERE id = $1
RETURNING *;

-- name: DeleteProfile :one
DELETE FROM profiles
WHERE id = $1
RETURNING id;

-- name: GetAccountById :one
SELECT * FROM accounts
WHERE id = $1;

-- name: GetAccountByUsername :one
SELECT * FROM accounts
where username = $1;

-- name: GetLoginInfo :one
SELECT * FROM login_info
WHERE id = $1;

-- name: GetLoginInfoByUsername :one
SELECT * FROM login_info
JOIN accounts ON account_id = accounts.id
WHERE username = $1;

-- name: CreateAccount :one
INSERT INTO accounts (username)
VALUES ($1)
RETURNING *;

-- name: CreateLoginInfo :one
INSERT INTO login_info (account_id, password_hesh)
VALUES ($1, $2)
RETURNING *;