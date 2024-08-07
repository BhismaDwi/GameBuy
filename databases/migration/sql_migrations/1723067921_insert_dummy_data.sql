-- +migrate Up
-- +migrate StatementBegin


INSERT INTO public.platform
(name, created_at, created_by, modified_at, modified_by)
VALUES('PS 4', '2024-08-08 04:49:04.246', 'Bhismas-MBP.lan', '2024-08-08 04:49:04.246', 'Bhismas-MBP.lan');

INSERT INTO public.role
(name)
VALUES('Admin');

INSERT INTO public.role
(name)
VALUES('Buyer');

INSERT INTO public.category
(name, created_at, created_by, modified_at, modified_by)
VALUES('Mistery', '2024-08-08 05:13:10.811', 'Bhismas-MBP.lan', '2024-08-08 05:13:57.140', 'Bhismas-MBP.lan');


-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

-- +migrate StatementEnd