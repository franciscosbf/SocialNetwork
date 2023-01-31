CREATE TABLE IF NOT EXISTS users_info.profiles (
    pid BIGSERIAL,
    aid BIGINT,
    description VARCHAR(300),
    first_name TEXT NOT NULL,
    middle_name TEXT,
    surname TEXT NOT NULL,
    location geography(POINT),

    PRIMARY KEY (pid),
    FOREIGN KEY (aid) REFERENCES users_info.accounts (aid)
);
