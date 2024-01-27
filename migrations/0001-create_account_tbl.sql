-- Table: public.account

-- DROP TABLE IF EXISTS public.account;

CREATE TABLE IF NOT EXISTS public.account
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    username character varying(50) COLLATE pg_catalog."default" NOT NULL,
    password bytea NOT NULL,
    CONSTRAINT account_pkey PRIMARY KEY (id),
    CONSTRAINT account_username_key UNIQUE (username)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.account
    OWNER to alex;