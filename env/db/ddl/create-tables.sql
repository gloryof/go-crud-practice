CREATE TABLE users (
    id bigint,
    name varchar(40),
    birthDay date,
    PRIMARY KEY(id)
);

CREATE SEQUENCE user_id_seq;