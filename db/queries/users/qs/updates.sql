-- Account related queries.

-- name: UpdateAccountUsername :exec
UPDATE users_info.accounts
SET username = @username
WHERE aid = @aid;

-- name: UpdateAccountPassword :exec
UPDATE users_info.accounts
SET password = @password
WHERE aid = @aid;

-- Profile related queries.

-- name: UpdateProfileDescription :exec
UPDATE users_info.profiles
SET description = @description
WHERE aid = @aid;

-- name: UpdateProfileFirstName :exec
UPDATE users_info.profiles
SET first_name = @firstName
WHERE aid = @aid;

-- name: UpdateProfileMiddleName :exec
UPDATE users_info.profiles
SET middle_name = @middleName
WHERE aid = @aid;

-- name: UpdateProfileSurname :exec
UPDATE users_info.profiles
SET surname = @surname
WHERE aid = @aid;

-- name: UpdateProfileLocation :exec
UPDATE users_info.profiles
SET location = @location
WHERE aid = @aid;
