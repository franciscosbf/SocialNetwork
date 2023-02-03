-- Accounts related queries.

-- name: GetAccountInternals :one
SELECT aid, email
FROM users_info.accounts
WHERE username = @username;

-- name: GetAccountInternalsAuth :one
SELECT aid, email
FROM users_info.accounts
WHERE username = @username AND
      password = @password;

-- name: GetAccountPhone :one
SELECT phone_prefix, phone_number
FROM users_info.accounts
WHERE aid = @aid;

-- Profiles related queries.

-- name: GetProfileName :one
SELECT first_name, middle_name, surname
FROM users_info.profiles
WHERE aid = @aid;

-- name: GetProfileLocation :one
SELECT location
FROM users_info.profiles
WHERE aid = @aid;

-- name: GetProfileDescription :one
SELECT description
FROM users_info.profiles
WHERE aid = @aid;
