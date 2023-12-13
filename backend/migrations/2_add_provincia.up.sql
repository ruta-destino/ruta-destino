CREATE TABLE IF NOT EXISTS provincia(
    id serial PRIMARY KEY,
    nombre text NOT NULL UNIQUE,
    id_region integer NOT NULL REFERENCES region(id)
);
