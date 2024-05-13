-- name: CreateUser :one
Insert Into users (id, created_at, updated_at, name, api_key)
values ($1, $2, $3, $4, $5)
returning *;

-- name: GetUser :one
Select * from users where api_key = $1;
