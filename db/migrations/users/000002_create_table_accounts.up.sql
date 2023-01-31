CREATE TABLE IF NOT EXISTS users_info.accounts (
    aid BIGSERIAL,
    username VARCHAR(30) NOT NULL,
    email VARCHAR(320) NOT NULL,
    password TEXT NOT NULL,
    phone_prefix INTEGER,
    phone_number INTEGER,

    PRIMARY KEY (aid),
    UNIQUE (username), UNIQUE (email),
    UNIQUE (username, email)
);
