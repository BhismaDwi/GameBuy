-- +migrate Up
-- +migrate StatementBegin


INSERT INTO public.platform
(name, created_at, created_by, modified_at, modified_by)
VALUES('PS 4', '2024-08-08 04:49:04.246', 'Bhismas-MBP.lan', '2024-08-08 04:49:04.246', 'Bhismas-MBP.lan');

INSERT INTO public.role
(name)
VALUES('admin');

INSERT INTO public.role
(name)
VALUES('buyer');

INSERT INTO public.category
(name, created_at, created_by, modified_at, modified_by)
VALUES('Mistery', '2024-08-08 05:13:10.811', 'Bhismas-MBP.lan', '2024-08-08 05:13:57.140', 'Bhismas-MBP.lan');

INSERT INTO public.game
(title, harga, category_id, platform_id, created_at, created_by, modified_at, modified_by)
VALUES('God Of War', 50000, 1, 1, '2024-08-08 20:28:31.000', 'Bhismas-MBP.lan', '2024-08-08 20:28:31.000', 'Bhismas-MBP.lan');

INSERT INTO public.game
(title, harga, category_id, platform_id, created_at, created_by, modified_at, modified_by)
VALUES('God Of War 2', 75000, 1, 1, '2024-08-08 20:28:31.000', 'Bhismas-MBP.lan', '2024-08-08 20:28:31.000', 'Bhismas-MBP.lan');

insert into users(username, password, role_id) Values('admin','$2a$10$CkX4OPGpKbfQAMqYOCd3SeeM/UGdStSq3iIS29sLSn/yh7vy8HEdu', 1)
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

-- +migrate StatementEnd