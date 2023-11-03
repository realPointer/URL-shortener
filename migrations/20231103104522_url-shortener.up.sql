CREATE TABLE IF NOT EXISTS links (
    shortURL varchar(20) PRIMARY KEY,
    originalURL varchar(255) NOT NULL UNIQUE
);