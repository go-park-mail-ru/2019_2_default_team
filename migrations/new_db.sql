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
    admin_id            bigserial   not null,
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



INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (2, 'Форд Против Феррари', 1, 'Драма', 169, 'Грей', 'США', 2019, 'Бейл', 'Противостояние машин... ', 2, 'https://static.karofilm.ru/uploads/film/desktop/03/b9/83/e827e7909888a133b74969ce6f.jpg', 'https://static.karofilm.ru/uploads/film/desktop/a0/dd/98/6ab8284a4d469111457c17e10f.jpg', 'https://www.youtube.com/embed/fAD-D3P-s0I', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (1, 'Джокер', 1, 'Драма', 159, 'Тодд', 'США', 2019, 'Феникс', 'Джокер, начало', 0, 'https://static.karofilm.ru/uploads/film/desktop/e3/b6/4b/5215d5811b80298172dad73fd1.jpg', 'https://static.karofilm.ru/uploads/film/desktop/06/5b/94/d93f295f6ac79d1e2060f657de.jpg', 'https://www.youtube.com/embed/50IJyz7ecqc', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (3, 'Звездные Войны', 1, 'Драма', 200, 'Абрамс', 'США', 2019, 'Фишер', 'Да прибудет с тобой сила', 1, 'https://static.karofilm.ru/uploads/film/desktop/bc/5f/6f/475136826ebdb3cd6319074222.jpg', 'https://static.karofilm.ru/uploads/film/desktop/8a/aa/45/65154177830e35a8d767451f3f.jpg', 'https://www.youtube.com/embed/KVRBfWQgyuY', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (4, 'Черное Рождество', 1, 'Триллер', 200, 'Такал', 'США', 2020, 'Путс', 'Осторожно...', 1, 'https://static.karofilm.ru/uploads/film/desktop/c8/30/8a/107315a28c4807e7ac5fd1798b.jpg', 'https://static.karofilm.ru/uploads/film/desktop/c8/30/8a/107315a28c4807e7ac5fd1798b.jpg', 'https://www.youtube.com/embed/dD_hjAcmnBM', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (5, 'Союз спасения', 1, 'Драма', 115, 'Кравчук', 'Россия', 2020, 'Прилучный', 'Мы вышли. Нам не вернуться.', 1, 'https://static.karofilm.ru/uploads/film/desktop/eb/9e/29/855c4cb7572e6aad56627846ce.jpg', 'https://static.karofilm.ru/uploads/film/desktop/eb/9e/29/855c4cb7572e6aad56627846ce.jpg', 'https://www.youtube.com/embed/U6QVzdXjQiQ', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (6, 'Не время умирать', 1, 'Триллер', 200, 'Фукунага', 'США', 2020, 'Крэйг', 'Не время умирать', 1, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR3FIXs2TB2e7fToTo8_rSMygr7-OjR0DhfLpj7QVk1TH2RGaUc', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR3FIXs2TB2e7fToTo8_rSMygr7-OjR0DhfLpj7QVk1TH2RGaUc', 'https://www.youtube.com/embed/PJX6a06gnEY', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (7, 'Грех', 1, 'Драма', 134, 'Кончаловский', 'Россия', 2019, 'Гуэррини', 'Грех...', 1, 'https://static.karofilm.ru/uploads/film/desktop/5b/d2/6a/01dece66b600e714565fb2f975.jpg', 'https://static.karofilm.ru/uploads/film/desktop/5b/d2/6a/01dece66b600e714565fb2f975.jpg', 'https://www.youtube.com/embed/bgo2k5YFtE0', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (8, 'Джуманджи', 1, 'Приключения', 100, 'Кэздан', 'США', 2019, 'Джонсон', 'Чтобы спасти одного из приятелей, остальным приходится вернуться в игру.', 1, 'https://static.karofilm.ru/uploads/film/desktop/cf/2e/32/a9472597e121fc10166a145b08.jpg', 'https://static.karofilm.ru/uploads/film/desktop/cf/2e/32/a9472597e121fc10166a145b08.jpg', 'https://www.youtube.com/embed/CYfB_wfzZuw', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (9, 'Холодное сердце 2', 1, 'Мультфильм', 100, 'Бак', 'США', 2019, 'Гад', 'Новогодняя сказка', 1, 'https://static.karofilm.ru/uploads/film/desktop/82/be/b1/e0b401256b508f396e379c76bf.jpg', 'https://static.karofilm.ru/uploads/film/desktop/82/be/b1/e0b401256b508f396e379c76bf.jpg', 'https://www.youtube.com/embed/kwG17PIZgr8', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (10, 'большой злой лис', 1, 'Мультфильм', 89, 'Имбер', 'Франция', 2019, 'Сам лис', 'Что делать Лису, если трое цыплят решили, что он — их мама? ', 1, 'https://static.karofilm.ru/uploads/film/desktop/71/48/fd/a01b6bdf1caf513384924f55ad.jpg', 'https://static.karofilm.ru/uploads/film/desktop/71/48/fd/a01b6bdf1caf513384924f55ad.jpg', 'https://www.youtube.com/embed/dowzU0kdp7s', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (11, 'Король лев', 1, 'Приключения', 120, 'Фавро', 'США', 2019, 'Сам лев', 'История того самого льва короля', 1, 'https://static.karofilm.ru/uploads/film/desktop/f7/b2/90/e45404daa4b739f26e19b3d4a5.jpg', 'https://static.karofilm.ru/uploads/film/desktop/f7/b2/90/e45404daa4b739f26e19b3d4a5.jpg', 'https://www.youtube.com/embed/kNhdUVgukVk', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (12, 'Вторжение', 1, 'Фантастика', 120, 'Бондарчук', 'Россия', 2019, 'Петров', 'Падение инопланетного объекта', 1, 'https://static.karofilm.ru/uploads/film/desktop/3b/65/ae/212e0ce1ccb361fc6fbd7f217d.jpg', 'https://static.karofilm.ru/uploads/film/desktop/3b/65/ae/212e0ce1ccb361fc6fbd7f217d.jpg', 'https://www.youtube.com/embed/6qONzflbQuY', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (13, 'Малефисента', 1, 'Фантастика', 115, 'Роннинг', 'США', 2019, 'Джоли', 'Действие происходит через несколько лет после того, как Малефисента наложила злые чары на принцессу Аврору. ', 1, 'https://static.karofilm.ru/uploads/film/desktop/68/81/d2/5a6e8d6275022a27b815c41557.jpg', 'https://static.karofilm.ru/uploads/film/desktop/68/81/d2/5a6e8d6275022a27b815c41557.jpg', 'https://www.youtube.com/embed/L0ttxMz-tyo', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (14, 'Во всё тяжкое', 1, 'Комедия', 90, 'Робертс', 'США', 2019, 'Депп', 'Профессор колледжа решает жить на полную катушку после того, как ему ставят серьезный диагноз.  ', 1, 'https://static.karofilm.ru/uploads/film/desktop/92/15/b3/07a4b425f0ce53aa4831a08a68.jpg', 'https://static.karofilm.ru/uploads/film/desktop/92/15/b3/07a4b425f0ce53aa4831a08a68.jpg', 'https://www.youtube.com/embed/Zy3lV2hHlOY', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (15, 'Полицейский с Рублевки. Новогодний беспредел 2', 1, 'Комедия', 90, 'Куликов', 'Россия', 2019, 'Бурунов', 'Близится новый год, и сотрудники отдела полиции Барвихи планируют праздновать его за городом в тёплой компании старых друзей и коллег. ', 1, 'https://static.karofilm.ru/uploads/film/desktop/db/a0/e1/b3a1c9c5b23666243f10d50157.jpg', 'https://static.karofilm.ru/uploads/film/desktop/db/a0/e1/b3a1c9c5b23666243f10d50157.jpg', 'https://www.youtube.com/embed/eCtlT3xu5Wo', false);

INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (1, 'super-hall', 1, '2019-12-30 17:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (2, 'super-hall', 2, '2019-12-30 17:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (3, 'super-hall', 3, '2019-12-30 17:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (4, 'super-hall', 4, '2019-12-30 17:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (5, 'super-hall', 5, '2019-12-30 17:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (6, 'super-hall', 6, '2019-12-30 17:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (7, 'super-hall', 7, '2019-12-30 17:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (8, 'super-hall', 8, '2019-12-30 17:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (9, 'super-hall', 9, '2019-12-30 17:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (10, 'super-hall', 10, '2019-12-30 19:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (11, 'super-hall', 11, '2019-12-30 19:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (12, 'super-hall', 12, '2019-12-30 19:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (13, 'super-hall', 13, '2019-12-30 19:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (14, 'super-hall', 14, '2019-12-30 19:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (15, 'super-hall', 15, '2019-12-30 19:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (16, 'super-hall', 1, '2019-12-30 21:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (17, 'super-hall', 3, '2019-12-30 21:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (18, 'super-hall', 5, '2019-12-30 21:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (19, 'super-hall', 7, '2019-12-30 21:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (20, 'super-hall', 9, '2019-12-30 21:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (21, 'super-hall', 2, '2019-12-30 23:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (22, 'super-hall', 4, '2019-12-31 13:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (23, 'super-hall', 6, '2019-12-31 13:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (24, 'super-hall', 8, '2019-12-31 13:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (25, 'super-hall', 10, '2019-12-31 13:30:16.968122', 'IMAX', false);

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
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (99, 6, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (100, 6, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (101, 6, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (102, 6, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (103, 6, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (104, 6, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (105, 6, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (106, 6, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (107, 6, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (108, 6, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (109, 6, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (110, 6, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (111, 6, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (112, 6, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (113, 6, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (114, 6, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (115, 6, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (116, 6, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (117, 6, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (118, 6, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (119, 7, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (120, 7, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (121, 7, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (122, 7, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (123, 7, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (124, 7, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (125, 7, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (126, 7, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (127, 7, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (128, 7, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (129, 7, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (130, 7, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (131, 7, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (132, 7, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (133, 7, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (134, 7, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (135, 7, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (136, 7, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (137, 7, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (138, 7, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (139, 8, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (140, 8, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (141, 8, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (142, 8, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (143, 8, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (144, 8, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (145, 8, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (146, 8, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (147, 8, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (148, 8, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (149, 8, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (150, 8, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (151, 8, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (152, 8, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (153, 8, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (154, 8, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (155, 8, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (156, 8, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (157, 8, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (158, 8, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (159, 9, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (160, 9, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (161, 9, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (162, 9, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (163, 9, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (164, 9, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (165, 9, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (166, 9, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (167, 9, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (168, 9, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (169, 9, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (170, 9, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (171, 9, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (172, 9, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (173, 9, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (174, 9, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (175, 9, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (176, 9, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (177, 9, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (178, 9, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (179, 10, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (180, 10, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (181, 10, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (182, 10, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (183, 10, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (184, 10, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (185, 10, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (186, 10, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (187, 10, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (188, 10, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (189, 10, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (190, 10, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (191, 10, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (192, 10, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (193, 10, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (194, 10, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (195, 10, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (196, 10, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (197, 11, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (198, 11, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (199, 11, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (200, 11, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (201, 11, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (202, 11, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (203, 11, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (204, 11, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (205, 11, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (206, 11, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (207, 11, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (208, 11, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (209, 11, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (210, 11, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (211, 11, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (212, 11, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (213, 11, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (214, 11, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (215, 11, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (216, 11, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (217, 12, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (218, 12, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (219, 12, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (220, 12, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (221, 12, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (222, 12, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (223, 12, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (224, 12, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (225, 12, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (226, 12, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (227, 12, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (228, 12, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (229, 12, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (230, 12, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (231, 12, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (232, 12, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (233, 12, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (234, 12, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (235, 12, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (236, 12, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (237, 13, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (238, 13, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (239, 13, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (240, 13, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (241, 13, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (242, 13, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (243, 13, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (244, 13, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (245, 13, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (246, 13, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (247, 13, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (248, 13, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (249, 13, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (250, 13, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (251, 13, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (252, 13, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (253, 13, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (254, 13, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (255, 13, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (256, 13, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (257, 14, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (258, 14, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (259, 14, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (260, 14, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (261, 14, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (262, 14, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (263, 14, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (264, 14, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (265, 14, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (266, 14, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (267, 14, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (268, 14, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (269, 14, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (270, 14, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (271, 14, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (272, 14, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (273, 14, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (274, 14, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (275, 14, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (276, 14, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (277, 15, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (278, 15, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (279, 15, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (280, 15, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (281, 15, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (282, 15, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (283, 15, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (284, 15, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (285, 15, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (286, 15, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (287, 15, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (288, 15, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (289, 15, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (290, 15, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (291, 15, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (292, 15, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (293, 15, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (294, 15, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (295, 16, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (296, 16, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (297, 16, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (298, 16, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (299, 16, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (300, 16, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (301, 16, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (302, 16, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (303, 16, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (304, 16, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (305, 16, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (306, 16, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (307, 16, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (308, 16, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (309, 16, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (310, 16, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (311, 16, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (312, 16, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (313, 16, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (314, 16, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (315, 17, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (316, 17, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (317, 17, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (318, 17, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (319, 17, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (320, 17, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (321, 17, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (322, 17, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (323, 17, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (324, 17, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (325, 17, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (326, 17, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (327, 17, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (328, 17, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (329, 17, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (330, 17, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (331, 17, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (332, 17, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (333, 17, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (334, 17, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (335, 18, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (336, 18, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (337, 18, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (338, 18, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (339, 18, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (340, 18, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (341, 18, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (342, 18, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (343, 18, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (344, 18, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (345, 18, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (346, 18, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (347, 18, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (348, 18, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (349, 18, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (350, 18, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (351, 18, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (352, 18, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (353, 18, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (354, 18, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (355, 19, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (356, 19, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (357, 19, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (358, 19, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (359, 19, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (360, 19, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (361, 19, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (362, 19, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (363, 19, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (364, 19, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (365, 19, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (366, 19, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (367, 19, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (368, 19, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (369, 19, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (370, 19, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (371, 19, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (372, 19, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (373, 19, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (374, 19, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (375, 20, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (376, 20, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (377, 20, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (378, 20, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (379, 20, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (380, 20, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (381, 20, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (382, 20, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (383, 20, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (384, 20, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (385, 20, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (386, 20, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (387, 20, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (388, 20, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (389, 20, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (390, 20, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (391, 20, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (392, 20, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (393, 21, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (394, 21, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (395, 21, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (396, 21, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (397, 21, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (398, 21, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (399, 21, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (400, 21, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (401, 21, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (402, 21, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (403, 21, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (404, 21, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (405, 21, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (406, 21, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (407, 21, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (408, 21, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (409, 21, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (410, 21, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (411, 21, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (412, 21, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (413, 22, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (414, 22, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (415, 22, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (416, 22, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (417, 22, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (418, 22, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (419, 22, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (420, 22, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (421, 22, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (422, 22, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (423, 22, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (424, 22, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (425, 22, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (426, 22, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (427, 22, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (428, 22, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (429, 22, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (430, 22, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (431, 22, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (432, 22, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (433, 23, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (434, 23, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (435, 23, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (436, 23, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (437, 23, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (438, 23, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (439, 23, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (440, 23, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (441, 23, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (442, 23, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (443, 23, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (444, 23, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (445, 23, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (446, 23, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (447, 23, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (448, 23, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (449, 23, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (450, 23, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (451, 23, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (452, 23, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (453, 24, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (454, 24, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (455, 24, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (456, 24, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (457, 24, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (458, 24, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (459, 24, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (460, 24, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (461, 24, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (462, 24, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (463, 24, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (464, 24, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (465, 24, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (466, 24, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (467, 24, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (468, 24, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (469, 24, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (470, 24, false, 5, 18, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (471, 24, false, 5, 19, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (472, 24, false, 6, 20, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (473, 25, false, 1, 1, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (474, 25, false, 1, 2, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (475, 25, false, 1, 3, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (476, 25, false, 2, 4, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (477, 25, false, 2, 5, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (478, 25, false, 2, 6, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (479, 25, false, 2, 7, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (480, 25, false, 3, 8, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (481, 25, false, 3, 9, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (482, 25, false, 3, 10, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (483, 25, false, 3, 11, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (484, 25, false, 4, 12, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (485, 25, false, 4, 13, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (486, 25, false, 4, 14, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (487, 25, false, 4, 15, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (488, 25, false, 5, 16, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (489, 25, false, 5, 17, 500, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, price, is_deleted) VALUES (490, 25, false, 5, 18, 500, false);



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

