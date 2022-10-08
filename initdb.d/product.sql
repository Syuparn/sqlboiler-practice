USE practice;

DROP TABLE IF EXISTS product;

CREATE TABLE product
(
    id          VARCHAR(26) PRIMARY KEY,
    name        VARCHAR(40) NOT NULL UNIQUE,
    category_id VARCHAR(26),
    CONSTRAINT fk_category_id
    FOREIGN KEY (category_id) 
    REFERENCES category (id)
    ON DELETE RESTRICT ON UPDATE RESTRICT
);
