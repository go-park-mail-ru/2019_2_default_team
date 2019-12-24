-- +migrate Up
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
    rating              int         not null default(0),
    poster              text        not null,
    poster_popup        text        not null,
    trailer              text       not null,
    is_deleted          boolean     not null default(false)
);

create table rating (
    vote_id             bigserial  primary key,
    user_id             int        not null,
    film_id             int        not null
);

--  Место
create table seat (
    seat_id                  bigserial   primary key,
    movie_session_id    int         not null,
--     hall_name           text        not null,
    is_taken            boolean     not null default(false),
    row                 int         not null,
    seat_number         int         not null,
    price               int         not null,
    is_deleted          boolean     not null default(false)
);

create table movie_session (
    ms_id               bigserial   primary key,
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
    start_datetime      timestamp   not null,
    profile_id          bigint      null references user_profile (user_id),
    price               int         not null,
    is_deleted          boolean     not null default(false)
);

INSERT INTO user_profile(nickname, password, first_name, last_name, email, avatar, is_deleted)
VALUES('usernick', 'password', 'user', 'surname', 'me@mail.ru', 'https://yt3.ggpht.com/a/AGF-l79J268L8ezHGDGnFh57D0wyFA-ltncqGF3cVA=s900-c-k-c0xffffffff-no-rj-mo', false);

INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (2, 'FordVSFerrari', 1, 'Drama', 169, 'Gray', 'USA', 2019, 'Bale', 'Future, cars... ', 2, 'https://static.karofilm.ru/uploads/film/desktop/03/b9/83/e827e7909888a133b74969ce6f.jpg', 'https://static.karofilm.ru/uploads/film/desktop/a0/dd/98/6ab8284a4d469111457c17e10f.jpg', 'https://www.youtube.com/embed/fAD-D3P-s0I', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (1, 'Joker', 1, 'Drama', 159, 'Todd', 'USA', 2019, 'Phoenix', 'Absolutely madness', 0, 'https://static.karofilm.ru/uploads/film/desktop/e3/b6/4b/5215d5811b80298172dad73fd1.jpg', 'https://static.karofilm.ru/uploads/film/desktop/06/5b/94/d93f295f6ac79d1e2060f657de.jpg', 'https://www.youtube.com/embed/50IJyz7ecqc', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (3, 'StarWars', 1, 'Drama', 200, 'Abrams', 'USA', 2019, 'Fisher', 'May the force be with you', 1, 'https://static.karofilm.ru/uploads/film/desktop/bc/5f/6f/475136826ebdb3cd6319074222.jpg', 'https://static.karofilm.ru/uploads/film/desktop/8a/aa/45/65154177830e35a8d767451f3f.jpg', 'https://www.youtube.com/embed/KVRBfWQgyuY', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (4, 'Черное Рождество', 1, 'Thriller', 200, 'Такал', 'USA', 2020, 'Путс', 'Осторожно...', 1, 'https://static.karofilm.ru/uploads/film/desktop/c8/30/8a/107315a28c4807e7ac5fd1798b.jpg', 'https://static.karofilm.ru/uploads/film/desktop/c8/30/8a/107315a28c4807e7ac5fd1798b.jpg', 'https://www.youtube.com/embed/dD_hjAcmnBM', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (5, 'Союз спасения', 1, 'Drama', 115, 'Кравчук', 'Россия', 2020, 'Прилучный', 'Мы вышли. Нам не вернуться.', 1, 'https://static.karofilm.ru/uploads/film/desktop/eb/9e/29/855c4cb7572e6aad56627846ce.jpg', 'https://static.karofilm.ru/uploads/film/desktop/eb/9e/29/855c4cb7572e6aad56627846ce.jpg', 'https://www.youtube.com/embed/U6QVzdXjQiQ', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (6, 'No Time to Die', 1, 'Thriller', 200, 'Фукунага', 'USA', 2020, 'Крэйг', 'Не время умирать', 1, 'https://vogue.ua/cache/inline_990x/uploads/article-inline/976/52f/149/5d5d14952f976.jpeg', 'https://vogue.ua/cache/inline_990x/uploads/article-inline/976/52f/149/5d5d14952f976.jpeg', 'https://www.youtube.com/embed/PJX6a06gnEY', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (7, 'Грех', 1, 'Drama', 134, 'Кончаловский', 'Россия', 2019, 'Гуэррини', 'Грех...', 1, 'https://static.karofilm.ru/uploads/film/desktop/5b/d2/6a/01dece66b600e714565fb2f975.jpg', 'https://static.karofilm.ru/uploads/film/desktop/5b/d2/6a/01dece66b600e714565fb2f975.jpg', 'https://www.youtube.com/embed/bgo2k5YFtE0', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (8, 'Текст', 1, 'Drama', 100, 'Шипенко', 'Россия', 2019, 'Петров', 'Спасти только одного...', 1, 'https://static.karofilm.ru/uploads/film/desktop/16/bd/fa/022b677f71fd227ab5d9414b7a.jpg', 'https://static.karofilm.ru/uploads/film/desktop/16/bd/fa/022b677f71fd227ab5d9414b7a.jpg', 'https://www.youtube.com/embed/PzsQd18Lu8k', false);



INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (3, 'super-hall', 3, '2019-12-24 19:28:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (4, 'super-hall', 2, '2019-12-24 23:50:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (5, 'super-hall', 2, '2019-12-24 23:55:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (6, 'super-hall', 1, '2019-12-24 17:00:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (7, 'super-hall', 1, '2019-12-24 18:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (34, 'super-hall2', 2, '2019-12-25 14:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (35, 'super-hall2', 3, '2019-12-25 14:45:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (36, 'super-hall2', 3, '2019-12-25 14:50:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (37, 'super-hall2', 5, '2019-12-25 15:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (38, 'super-hall2', 4, '2019-12-25 17:30:16.968122', 'IMAX', false);


INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (140, 36, false, 1, 1, 500, false);



CREATE TABLE message (
    message_id serial PRIMARY KEY,
    author_id integer references user_profile,
    to_user integer references user_profile,
    time timestamp with time zone DEFAULT now() NOT NULL,
    is_edited boolean DEFAULT FALSE NOT NULL,
    message text NOT NULL
);

CREATE table user_to_sup(
    uts_id serial PRIMARY key,
    user_id integer references user_profile,
    sup_id integer references user_profile
);

create table support(
    id serial primary key ,
    user_id integer references user_profile,
    status integer NOT NULL default 0
);

create table comments(
    id serial primary key ,
    film_title text,
    username text,
    comment text
);

-- +migrate Down
drop table comments cascade;
drop table film_profile cascade;
drop table message cascade;
drop table movie_session cascade;
drop table rating cascade;
drop table seat cascade;
drop table support cascade;
drop table ticket_profile cascade;
drop table user_profile cascade;
drop table user_to_sup cascade;

