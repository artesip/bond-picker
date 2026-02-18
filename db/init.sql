CREATE TABLE IF NOT EXISTS t_user
(
    id       UUID DEFAULT uuidv7() PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT        NOT NULL,
    salt     TEXT        NOT NULL
);

CREATE INDEX username_indx ON t_user USING HASH (username);
