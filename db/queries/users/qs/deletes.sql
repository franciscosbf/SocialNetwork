-- name: RemoveProfile :exec
DELETE FROM users_info.profiles
WHERE aid = @aid;

-- name: RemoveAccount :exec
DELETE FROM users_info.accounts
WHERE aid = @aid;
