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
    profile_id          bigint      null references user_profile (user_id),
    price               int         not null,
    is_deleted          boolean     not null default(false)
);

INSERT INTO user_profile(nickname, password, first_name, last_name, email, avatar, is_deleted)
VALUES('usernick', 'password', 'user', 'surname', 'me@mail.ru', 'urlurl', false);

INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (2, 'ToStars', 1, 'Science, Fantastic', 169, 'Gray', 'USA', 2019, 'Pitt', 'Future, stars... ', 2, 'https://static.karofilm.ru/uploads/film/desktop/e3/b6/4b/5215d5811b80298172dad73fd1.jpg', 'https://static.karofilm.ru/uploads/film/desktop/06/5b/94/d93f295f6ac79d1e2060f657de.jpg', 'https://www.youtube.com/watch?v=50IJyz7ecqc', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (1, 'Joker', 1, 'Comics', 159, 'Todd', 'USA', 2019, 'Phoenix', 'Absolutely madness', 0, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT92gzaTkJ040i9HBmLktkcNESMEYS5mm4gFLPhNOVbt0_MSk1E', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT92gzaTkJ040i9HBmLktkcNESMEYS5mm4gFLPhNOVbt0_MSk1E', 'https://www.youtube.com/embed/m6vhNci6RHc', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, rating, poster, poster_popup, trailer, is_deleted) VALUES (3, 'StarWars', 1, 'Fantastic', 200, 'Abrams', 'USA', 2019, 'Fisher', 'May the force be with you', 1, 'https://static.karofilm.ru/uploads/film/desktop/bc/5f/6f/475136826ebdb3cd6319074222.jpg', 'https://static.karofilm.ru/uploads/film/desktop/8a/aa/45/65154177830e35a8d767451f3f.jpg', 'https://www.youtube.com/embed/KVRBfWQgyuY', false);


INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (3, 'super-hall', 3, '2019-12-16 19:28:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (4, 'super-hall', 2, '2019-12-16 18:28:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (5, 'super-hall', 2, '2019-12-16 20:28:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (6, 'super-hall', 1, '2019-12-16 20:28:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (7, 'super-hall', 1, '2019-12-16 21:28:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (34, 'super-hall2', 2, '2019-12-16 14:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (35, 'super-hall2', 3, '2019-12-16 14:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (36, 'super-hall2', 3, '2019-12-17 14:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (37, 'super-hall2', 2, '2019-12-17 14:30:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (38, 'super-hall2', 2, '2019-12-17 17:30:16.968122', 'IMAX', false);


INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (140, 36, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (141, 36, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (142, 36, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (143, 36, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (144, 36, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (145, 36, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (146, 36, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (147, 36, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (148, 36, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (149, 36, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (150, 36, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (151, 36, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (152, 36, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (153, 36, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (154, 36, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (155, 36, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (156, 36, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (157, 36, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (158, 36, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (159, 36, false, 6, 20, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (160, 37, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (161, 37, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (162, 37, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (163, 37, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (164, 37, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (165, 37, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (166, 37, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (167, 37, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (168, 37, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (169, 37, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (170, 37, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (171, 37, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (172, 37, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (173, 37, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (174, 37, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (175, 37, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (176, 37, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (177, 37, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (178, 37, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (179, 37, false, 6, 20, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (180, 38, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (181, 38, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (182, 38, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (183, 38, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (184, 38, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (185, 38, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (186, 38, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (187, 38, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (188, 38, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (189, 38, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (190, 38, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (191, 38, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (192, 38, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (193, 38, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (194, 38, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (195, 38, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (196, 38, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (197, 38, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (198, 38, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (100, 34, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (101, 34, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (102, 34, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (103, 34, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (104, 34, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (105, 34, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (106, 34, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (107, 34, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (108, 34, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (109, 34, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (110, 34, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (111, 34, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (112, 34, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (113, 34, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (114, 34, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (115, 34, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (116, 34, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (117, 34, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (118, 34, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (119, 34, false, 6, 20, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (120, 35, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (121, 35, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (122, 35, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (123, 35, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (124, 35, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (125, 35, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (126, 35, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (127, 35, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (96, 7, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (128, 35, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (129, 35, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (130, 35, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (131, 35, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (132, 35, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (133, 35, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (134, 35, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (135, 35, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (136, 35, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (1, 3, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (2, 3, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (3, 3, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (4, 3, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (5, 3, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (6, 3, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (7, 3, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (8, 3, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (9, 3, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (10, 3, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (11, 3, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (12, 3, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (13, 3, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (14, 3, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (15, 3, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (16, 3, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (17, 3, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (18, 3, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (19, 3, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (20, 3, false, 6, 20, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (21, 4, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (22, 4, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (23, 4, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (24, 4, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (25, 4, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (26, 4, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (27, 4, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (28, 4, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (29, 4, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (30, 4, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (31, 4, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (32, 4, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (33, 4, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (34, 4, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (35, 4, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (36, 4, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (37, 4, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (38, 4, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (39, 4, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (40, 4, false, 6, 20, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (41, 5, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (42, 5, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (43, 5, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (44, 5, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (45, 5, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (46, 5, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (47, 5, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (48, 5, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (49, 5, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (50, 5, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (51, 5, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (52, 5, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (53, 5, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (54, 5, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (55, 5, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (56, 5, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (57, 5, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (58, 5, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (59, 5, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (60, 5, false, 6, 20, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (61, 6, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (62, 6, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (63, 6, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (64, 6, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (65, 6, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (66, 6, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (67, 6, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (68, 6, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (69, 6, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (70, 6, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (71, 6, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (72, 6, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (73, 6, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (74, 6, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (75, 6, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (76, 6, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (77, 6, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (78, 6, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (79, 6, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (80, 6, false, 6, 20, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (81, 7, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (82, 7, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (83, 7, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (84, 7, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (85, 7, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (86, 7, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (87, 7, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (88, 7, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (89, 7, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (90, 7, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (91, 7, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (92, 7, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (93, 7, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (94, 7, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (95, 7, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (137, 35, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (138, 35, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (139, 35, false, 6, 20, false);


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

