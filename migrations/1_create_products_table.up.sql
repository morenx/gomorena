CREATE TABLE IF NOT EXISTS products
(
    id BINARY(16) NOT NULL DEFAULT unhex(replace(uuid(), '-', '')),
    name VARCHAR(20) CHARACTER SET utf8 NOT NULL,
    PRIMARY KEY(id)
);
