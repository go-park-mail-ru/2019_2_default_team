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

INSERT INTO film_profile(title, admin_id, genre, length, director, production, year, actors, description, is_deleted)
VALUES('Joker', 1, 'Genre', 159, 'Todd', 'USA', 2019, 'Phoenix', 'Absolutely madness', false);
INSERT INTO public.film_profile (film_id, title, admin_id, genre, length, director, production, year, actors, description, is_deleted, rating) VALUES (2, 'ToStars', 1, 'g', 169, 'D', 'USA', 2019, 'A', 'D', false, 2);

INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (1, '1', 1, '2019-11-20 15:31:37.000000', '2D', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (2, 'super-hall', 1, '2019-12-05 16:28:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (5, 'super-hall2', 2, '2019-12-05 20:28:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (6, 'super-hall2', 2, '2019-12-06 20:28:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (7, 'super-hall2', 2, '2019-12-12 20:28:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (3, 'super-hall2', 1, '2019-12-16 16:28:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (4, 'super-hall2', 1, '2019-12-15 18:28:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (9, 'super-hall', 2, '2019-12-15 19:28:16.968122', 'IMAX', false);
INSERT INTO public.movie_session (ms_id, hall_name, movie_id, start_datetime, type, is_deleted) VALUES (8, 'super-hall', 2, '2019-12-15 19:28:16.968122', 'IMAX', false);

INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (1, 2, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (2, 2, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (3, 2, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (4, 2, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (5, 2, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (6, 2, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (7, 2, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (8, 2, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (9, 2, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (10, 2, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (11, 2, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (12, 2, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (13, 2, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (14, 2, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (15, 2, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (16, 2, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (17, 2, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (18, 2, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (19, 2, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (20, 2, false, 6, 20, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (21, 3, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (22, 3, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (23, 3, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (24, 3, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (25, 3, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (26, 3, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (27, 3, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (28, 3, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (29, 3, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (30, 3, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (31, 3, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (32, 3, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (33, 3, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (34, 3, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (35, 3, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (36, 3, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (37, 3, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (38, 3, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (39, 3, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (40, 3, false, 6, 20, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (41, 4, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (42, 4, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (43, 4, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (44, 4, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (45, 4, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (46, 4, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (47, 4, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (48, 4, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (49, 4, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (50, 4, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (51, 4, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (52, 4, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (53, 4, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (54, 4, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (55, 4, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (56, 4, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (57, 4, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (58, 4, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (59, 4, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (60, 4, false, 6, 20, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (61, 5, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (64, 5, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (65, 5, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (66, 5, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (67, 5, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (68, 5, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (69, 5, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (70, 5, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (71, 5, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (72, 5, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (73, 5, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (74, 5, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (75, 5, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (76, 5, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (77, 5, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (78, 5, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (79, 5, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (80, 5, false, 6, 20, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (81, 6, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (82, 6, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (83, 6, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (84, 6, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (85, 6, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (86, 6, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (87, 6, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (88, 6, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (89, 6, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (90, 6, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (91, 6, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (92, 6, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (93, 6, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (94, 6, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (95, 6, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (96, 6, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (97, 6, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (98, 6, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (99, 6, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (100, 6, false, 6, 20, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (101, 7, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (102, 7, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (103, 7, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (104, 7, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (105, 7, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (106, 7, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (107, 7, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (108, 7, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (109, 7, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (110, 7, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (111, 7, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (112, 7, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (113, 7, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (114, 7, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (115, 7, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (116, 7, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (117, 7, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (118, 7, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (119, 7, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (120, 7, false, 6, 20, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (121, 8, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (122, 8, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (123, 8, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (124, 8, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (125, 8, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (126, 8, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (127, 8, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (128, 8, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (129, 8, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (130, 8, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (131, 8, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (132, 8, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (133, 8, false, 4, 13, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (134, 8, false, 4, 14, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (135, 8, false, 4, 15, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (136, 8, false, 5, 16, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (63, 5, true, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (137, 8, false, 5, 17, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (138, 8, false, 5, 18, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (139, 8, false, 5, 19, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (140, 8, false, 6, 20, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (141, 9, false, 1, 1, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (142, 9, false, 1, 2, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (143, 9, false, 1, 3, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (144, 9, false, 2, 4, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (145, 9, false, 2, 5, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (146, 9, false, 2, 6, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (147, 9, false, 2, 7, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (148, 9, false, 3, 8, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (149, 9, false, 3, 9, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (150, 9, false, 3, 10, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (151, 9, false, 3, 11, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (152, 9, false, 4, 12, false);
INSERT INTO public.seat (seat_id, movie_session_id, is_taken, row, seat_number, is_deleted) VALUES (62, 5, true, 1, 2, false);



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
