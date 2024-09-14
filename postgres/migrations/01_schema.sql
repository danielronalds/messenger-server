CREATE SCHEMA api;

CREATE TABLE IF NOT EXISTS api.Users (
    UserName    VARCHAR(50) NOT NULL PRIMARY KEY,
    DisplayName VARCHAR(50) NOT NULL,
    Password    bytea NOT NULL,
    Salt        bytea NOT NULL
);

CREATE TABLE IF NOT EXISTS api.Messages (
    Id SERIAL PRIMARY KEY NOT NULL,
    Receiver INTEGER NOT NULL REFERENCES api.Users,
    Sender INTEGER NOT NULL REFERENCES api.Users,
    Context TEXT NOT NULL,
    Delivered TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC')
);
