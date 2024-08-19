CREATE SCHEMA api;

CREATE TABLE IF NOT EXISTS api.Users (
    Id          SERIAL PRIMARY KEY,
    UserName    VARCHAR(50) NOT NULL,
    DisplayName VARCHAR(50) NOT NULL,
    Password    VARCHAR(50) NOT NULL -- NOTE: May need to adjust this based on the size of the hash
);

CREATE TABLE IF NOT EXISTS api.Messages (
    Id SERIAL PRIMARY KEY NOT NULL,
    Receiver INTEGER NOT NULL REFERENCES api.Users,
    Sender INTEGER NOT NULL REFERENCES api.Users,
    Context TEXT NOT NULL,
    Delivered TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC')
);
