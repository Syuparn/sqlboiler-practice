USE practice;

DROP TABLE IF EXISTS category;

CREATE TABLE category
(
    id          VARCHAR(26) PRIMARY KEY,
    name        VARCHAR(40) NOT NULL UNIQUE
);
