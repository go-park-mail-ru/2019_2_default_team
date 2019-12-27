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

-- жанры

create table user_genres (
   genre_id          bigserial   primary key,
    user_id          int          not null ,
    genre            text        not null
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

INSERT INTO user_profile(nickname, password, first_name, last_name, email, avatar,  is_deleted)
VALUES('usernick', 'password', 'user', 'surname', 'me@mail.ru', 'https://yt3.ggpht.com/a/AGF-l79J268L8ezHGDGnFh57D0wyFA-ltncqGF3cVA=s900-c-k-c0xffffffff-no-rj-mo', false);

insert into public.user_genres(genre_id, user_id, genre) VALUES(1, 1,'Drama');
insert into public.user_genres(genre_id, user_id, genre) VALUES(2, 1,'Science');

INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (2, 'Форд Против Феррари', 1, 'Драма', 169, 'Грей', 'США', 2019, 'Бейл', 'Противостояние машин... ', 2, 'https://static.karofilm.ru/uploads/film/desktop/03/b9/83/e827e7909888a133b74969ce6f.jpg', 'https://static.karofilm.ru/uploads/film/desktop/a0/dd/98/6ab8284a4d469111457c17e10f.jpg', 'https://www.youtube.com/embed/fAD-D3P-s0I', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (1, 'Джокер', 1, 'Драма', 159, 'Тодд', 'США', 2019, 'Феникс', 'Джокер, начало', 0, 'https://static.karofilm.ru/uploads/film/desktop/e3/b6/4b/5215d5811b80298172dad73fd1.jpg', 'https://static.karofilm.ru/uploads/film/desktop/06/5b/94/d93f295f6ac79d1e2060f657de.jpg', 'https://www.youtube.com/embed/50IJyz7ecqc', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (3, 'Звездные Войны', 1, 'Драма', 200, 'Абрамс', 'США', 2019, 'Фишер', 'Да прибудет с тобой сила', 1, 'https://static.karofilm.ru/uploads/film/desktop/bc/5f/6f/475136826ebdb3cd6319074222.jpg', 'https://static.karofilm.ru/uploads/film/desktop/8a/aa/45/65154177830e35a8d767451f3f.jpg', 'https://www.youtube.com/embed/KVRBfWQgyuY', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (4, 'Черное Рождество', 1, 'Триллер', 200, 'Такал', 'США', 2020, 'Путс', 'Осторожно...', 1, 'https://static.karofilm.ru/uploads/film/desktop/c8/30/8a/107315a28c4807e7ac5fd1798b.jpg', 'https://static.karofilm.ru/uploads/film/desktop/c8/30/8a/107315a28c4807e7ac5fd1798b.jpg', 'https://www.youtube.com/embed/dD_hjAcmnBM', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (5, 'Союз спасения', 1, 'Драма', 115, 'Кравчук', 'Россия', 2020, 'Прилучный', 'Мы вышли. Нам не вернуться.', 1, 'https://static.karofilm.ru/uploads/film/desktop/eb/9e/29/855c4cb7572e6aad56627846ce.jpg', 'https://static.karofilm.ru/uploads/film/desktop/eb/9e/29/855c4cb7572e6aad56627846ce.jpg', 'https://www.youtube.com/embed/U6QVzdXjQiQ', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (6, 'Не время умирать', 1, 'Триллер', 200, 'Фукунага', 'США', 2020, 'Крэйг', 'Не время умирать', 1, 'https://vogue.ua/cache/inline_990x/uploads/article-inline/976/52f/149/5d5d14952f976.jpeg', 'https://vogue.ua/cache/inline_990x/uploads/article-inline/976/52f/149/5d5d14952f976.jpeg', 'https://www.youtube.com/embed/PJX6a06gnEY', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (7, 'Грех', 1, 'Драма', 134, 'Кончаловский', 'Россия', 2019, 'Гуэррини', 'Грех...', 1, 'https://static.karofilm.ru/uploads/film/desktop/5b/d2/6a/01dece66b600e714565fb2f975.jpg', 'https://static.karofilm.ru/uploads/film/desktop/5b/d2/6a/01dece66b600e714565fb2f975.jpg', 'https://www.youtube.com/embed/bgo2k5YFtE0', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (8, 'Джуманджи', 1, 'Приключения', 100, 'Кэздан', 'США', 2019, 'Джонсон', 'Чтобы спасти одного из приятелей, остальным приходится вернуться в игру.', 1, 'https://static.karofilm.ru/uploads/film/desktop/cf/2e/32/a9472597e121fc10166a145b08.jpg', 'https://static.karofilm.ru/uploads/film/desktop/cf/2e/32/a9472597e121fc10166a145b08.jpg', 'https://www.youtube.com/embed/CYfB_wfzZuw', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (9, 'Холодное сердце 2', 1, 'Мультфильм', 100, 'Бак', 'США', 2019, 'Гад', 'Новогодняя сказка', 1, 'https://static.karofilm.ru/uploads/film/desktop/82/be/b1/e0b401256b508f396e379c76bf.jpg', 'https://static.karofilm.ru/uploads/film/desktop/82/be/b1/e0b401256b508f396e379c76bf.jpg', 'https://www.youtube.com/embed/kwG17PIZgr8', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (10, 'большой злой лис', 1, 'Мультфильм', 89, 'Имбер', 'Франция', 2019, 'Сам лис', 'Что делать Лису, если трое цыплят решили, что он — их мама? ', 1, 'https://static.karofilm.ru/uploads/film/desktop/71/48/fd/a01b6bdf1caf513384924f55ad.jpg', 'https://static.karofilm.ru/uploads/film/desktop/71/48/fd/a01b6bdf1caf513384924f55ad.jpg', 'https://www.youtube.com/embed/dowzU0kdp7s', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (11, 'Король лев', 1, 'Приключения', 120, 'Фавро', 'США', 2019, 'Сам лев', 'История того самого льва короля', 1, 'https://static.karofilm.ru/uploads/film/desktop/f7/b2/90/e45404daa4b739f26e19b3d4a5.jpg', 'https://static.karofilm.ru/uploads/film/desktop/f7/b2/90/e45404daa4b739f26e19b3d4a5.jpg', 'https://www.youtube.com/embed/kNhdUVgukVk', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (12, 'Вторжение', 1, 'Фантастика', 120, 'Бондарчук', 'Россия', 2019, 'Петров', 'Падение инопланетного объекта', 1, 'https://static.karofilm.ru/uploads/film/desktop/3b/65/ae/212e0ce1ccb361fc6fbd7f217d.jpg', 'https://static.karofilm.ru/uploads/film/desktop/3b/65/ae/212e0ce1ccb361fc6fbd7f217d.jpg', 'https://www.youtube.com/embed/6qONzflbQuY', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (13, 'Малефисента', 1, 'Фантастика', 115, 'Роннинг', 'США', 2019, 'Джоли', 'Действие происходит через несколько лет после того, как Малефисента наложила злые чары на принцессу Аврору. ', 1, 'https://static.karofilm.ru/uploads/film/desktop/68/81/d2/5a6e8d6275022a27b815c41557.jpg', 'https://static.karofilm.ru/uploads/film/desktop/68/81/d2/5a6e8d6275022a27b815c41557.jpg', 'https://www.youtube.com/embed/L0ttxMz-tyo', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (14, 'Во всё тяжкое', 1, 'Комедия', 90, 'Робертс', 'США', 2019, 'Депп', 'Профессор колледжа решает жить на полную катушку после того, как ему ставят серьезный диагноз.  ', 1, 'https://static.karofilm.ru/uploads/film/desktop/92/15/b3/07a4b425f0ce53aa4831a08a68.jpg', 'https://static.karofilm.ru/uploads/film/desktop/92/15/b3/07a4b425f0ce53aa4831a08a68.jpg', 'https://www.youtube.com/embed/Zy3lV2hHlOY', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (15, 'Полицейский с Рублевки. Новогодний беспредел 2', 1, 'Комедия', 90, 'Куликов', 'Россия', 2019, 'Бурунов', 'Близится новый год, и сотрудники отдела полиции Барвихи планируют праздновать его за городом в тёплой компании старых друзей и коллег. ', 1, 'https://static.karofilm.ru/uploads/film/desktop/db/a0/e1/b3a1c9c5b23666243f10d50157.jpg', 'https://static.karofilm.ru/uploads/film/desktop/db/a0/e1/b3a1c9c5b23666243f10d50157.jpg', 'https://www.youtube.com/embed/eCtlT3xu5Wo', false);



INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (1, 'super-hall2', 2, '2019-12-27 21:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (2, 'super-hall', 1, '2019-12-27 20:50:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (3, 'super-hall', 1, '2019-12-27 21:00:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (4, 'super-hall', 1, '2019-12-27 22:40:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (5, 'super-hall', 1, '2019-12-28 22:50:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (6, 'super-hall', 1, '2019-12-27 23:20:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (7, 'super-hall', 5, '2019-12-27 21:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (8, 'super-hall', 4, '2019-12-27 22:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (9, 'super-hall', 3, '2019-12-27 23:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (10, 'super-hall', 2, '2019-12-28 23:30:16.968122', 'IMAX', false);


INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (1, 1, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (2, 1, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (3, 1, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (4, 1, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (5, 1, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (6, 1, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (7, 1, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (8, 1, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (9, 1, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (10, 1, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (11, 1, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (12, 1, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (13, 1, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (14, 1, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (15, 1, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (16, 1, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (17, 1, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (18, 1, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (19, 1, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (20, 1, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (21, 2, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (22, 2, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (23, 2, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (24, 2, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (25, 2, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (26, 2, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (27, 2, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (28, 2, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (29, 2, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (30, 2, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (31, 2, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (32, 2, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (33, 2, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (34, 2, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (35, 2, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (36, 2, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (37, 2, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (38, 2, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (39, 2, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (40, 2, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (41, 3, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (42, 3, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (43, 3, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (44, 3, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (45, 3, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (46, 3, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (47, 3, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (48, 3, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (49, 3, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (50, 3, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (51, 3, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (52, 3, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (53, 3, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (54, 3, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (55, 3, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (56, 3, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (57, 3, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (58, 3, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (59, 3, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (60, 3, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (61, 4, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (62, 4, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (63, 4, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (64, 4, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (65, 4, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (66, 4, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (67, 4, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (68, 4, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (69, 4, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (70, 4, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (71, 4, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (72, 4, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (73, 4, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (74, 4, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (75, 4, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (76, 4, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (77, 4, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (78, 4, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (79, 4, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (80, 4, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (81, 5, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (82, 5, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (83, 5, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (84, 5, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (85, 5, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (86, 5, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (87, 5, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (88, 5, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (89, 5, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (90, 5, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (91, 5, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (92, 5, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (93, 5, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (94, 5, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (95, 5, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (96, 5, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (97, 5, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (98, 5, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (99, 5, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (100, 5, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (101, 6, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (102, 6, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (103, 6, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (104, 6, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (105, 6, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (106, 6, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (107, 6, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (108, 6, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (109, 6, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (110, 6, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (111, 6, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (112, 6, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (113, 6, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (114, 6, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (115, 6, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (116, 6, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (117, 6, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (118, 7, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (119, 7, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (120, 7, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (121, 7, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (122, 7, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (123, 7, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (124, 7, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (125, 7, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (126, 7, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (127, 7, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (128, 7, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (129, 7, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (130, 7, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (131, 7, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (132, 7, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (133, 7, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (134, 7, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (135, 7, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (136, 7, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (137, 7, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (138, 8, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (139, 8, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (140, 8, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (141, 8, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (142, 8, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (143, 8, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (144, 8, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (145, 8, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (146, 8, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (147, 8, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (148, 8, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (149, 8, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (150, 8, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (151, 8, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (152, 8, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (153, 8, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (154, 8, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (155, 8, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (156, 8, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (157, 8, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (158, 9, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (159, 9, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (160, 9, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (161, 9, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (162, 9, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (163, 9, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (164, 9, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (165, 9, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (166, 9, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (167, 9, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (168, 9, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (169, 9, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (170, 9, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (171, 9, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (172, 9, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (173, 9, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (174, 9, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (175, 9, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (176, 9, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (177, 9, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (178, 10, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (179, 10, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (180, 10, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (181, 10, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (182, 10, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (183, 10, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (184, 10, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (185, 10, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (186, 10, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (187, 10, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (188, 10, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (189, 10, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (190, 10, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (191, 10, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (192, 10, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (193, 10, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (194, 10, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (195, 10, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (196, 10, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (197, 10, false, 6, 20, 500, false);

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
drop table user_genres cascade;

