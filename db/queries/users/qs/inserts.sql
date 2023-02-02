-- name: InsertNewAccount :one
INSERT INTO users_info.accounts (username, email, password, phone_prefix, phone_number)
VALUES (@username, @email, @password, @phone_prefix, @phone_number)
RETURNING aid;

-- name: InsertNewProfile :exec
INSERT INTO users_info.profiles (aid, first_name, middle_name, surname)
VALUES (@aid, @first_name, @middle_name, @surname);
