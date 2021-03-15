CREATE USER short_line;

CREATE DATABASE short_line;
ALTER USER short_line WITH PASSWORD 'QWERTY';
GRANT ALL PRIVILEGES ON DATABASE short_line TO short_line;


CREATE TABLE public.urls (
    id integer NOT NULL,
    full_address_name TEXT,
    short_key TEXT NULL
);


CREATE SEQUENCE public.urls_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.urls_id_seq OWNER TO short_line;


ALTER SEQUENCE public.urls_id_seq OWNED BY public.urls.id;
