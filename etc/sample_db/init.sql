CREATE USER short_line;

CREATE DATABASE short_line;
ALTER USER short_line WITH PASSWORD 'QWERTY';
GRANT ALL PRIVILEGES ON DATABASE short_line TO short_line;


CREATE TABLE public.urls (
    id  SERIAL PRIMARY KEY,
    url TEXT,
    short_key TEXT NULL
);
