CREATE TABLE IF NOT EXISTS t_user
(
    id       UUID DEFAULT uuidv7() PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT        NOT NULL
);
