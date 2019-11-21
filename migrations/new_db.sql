-- Профиль
create table user_profile (
    user_id                  bigserial   primary key,
    nickname            text        not null,
    password            text        not null,
    first_name          text        not null,
    last_name           text        not null,
    email               text        not null,
    avatar              text        null,
    is_deleted          boolean     not null default(false)
);

-- Фильм
create table film_profile (
    film_id                  bigserial   primary key,
    title                text        not null,
    admin_id            bigserial   not null references user_profile (user_id),
    genre               text        not null,
    length              int         not null,
    director            text        not null,
    production          text        not null,
    year                int         not null,
    actors              text        not null,
    description         text        not null,
    is_deleted          boolean     not null default(false)
);

--  Место
create table seat (
    seat_id                  bigserial   primary key,
    hall_name           text        not null,
    row                 int         not null,
    seat_number         int         not null,
    is_deleted          boolean     not null default(false)
);

create table movie_session (
    ms_id                  bigserial   primary key,
    hall_name           text        not null,
    movie_id            bigint      not null references film_profile (film_id),
    start_datetime      timestamp   not null,
    type                text        not null,
    is_deleted          boolean     not null default(false)
);

create table ticket_profile (
    ticket_id                  bigserial   primary key,
    movie_session_id    bigint      not null references movie_session (ms_id),
    seat_id             bigint      not null references seat (seat_id),
    profile_id          bigint      null references user_profile (user_id),
    price               int         not null,
    is_deleted          boolean     not null default(false)
);

INSERT INTO user_profile(nickname, password, first_name, last_name, email, avatar, is_deleted)
VALUES('usernick', 'password', 'user', 'surname', 'me@mail.ru', 'urlurl', false);

INSERT INTO film_profile(title, admin_id, genre, length, director, production, year, actors, description, is_deleted)
VALUES('Joker', 1, 'Genre', 159, 'Todd', 'USA', 2019, 'Phoenix', 'Absolutely madness', false);

INSERT INTO movie_session(hall_name, movie_id, start_datetime, type, is_deleted)
VALUES('1', 1, '2019-11-20 15:31:37.000000', '2D', false);

INSERT INTO seat(hall_name, row, seat_number, is_deleted)
VALUES('1', 1, 1, false);

INSERT INTO ticket_profile(movie_session_id, seat_id, profile_id, price, is_deleted)
VALUES(1, 1, 1, 100, false);


