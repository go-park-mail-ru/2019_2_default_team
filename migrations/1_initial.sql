-- +migrate Up

-- DROP TABLE user_profile;
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE user_profile (
    user_id serial PRIMARY KEY,
    email varchar(400) UNIQUE NOT NULL,
    password varchar(64) NOT NULL,

    nickname varchar(400) UNIQUE NOT NULL,
    avatar text
);

CREATE TABLE film_profile (
    film_id serial PRIMARY KEY,
    title varchar(400) UNIQUE NOT NULL,
    description varchar(64) NOT NULL,
    director varchar(400) NOT NULL,
    mainactor varchar(40) NOT NULL,
    admin_id varchar(40) NOT NULL,
    avatar text
);

CREATE TABLE ticket_profile (
    ticket_id serial PRIMARY KEY,
    user_id varchar(40) NOT NULL,
    film_id varchar(40) NOT NULL
);

INSERT INTO ticket_profile(user_id, film_id)
VALUES('1', '1');

INSERT INTO user_profile(email, password, nickname)
VALUES('me@mail.ru', 'password', 'usernick');

INSERT INTO film_profile(title, description, director, mainactor, admin_id)
VALUES('Joker', 'Absolutely madness', 'Todd', 'Phoenix', 1);

-- +migrate Down
DROP TABLE IF EXISTS user_profile;
DROP TABLE IF EXISTS film_profile;
DROP TABLE IF EXISTS ticket_profile;

DROP EXTENSION IF EXISTS citext;