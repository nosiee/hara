CREATE TABLE apikeys(
    ID SERIAL PRIMARY KEY,
    uuid VARCHAR(64) UNIQUE NOT NULL,
    key VARCHAR(64) UNIQUE NOT NULL,
    maxquota INT NOT NULL,
    quota INT NOT NULL,
    updatetime INT NOT NULL
);
