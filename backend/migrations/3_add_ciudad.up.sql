CREATE TABLE IF NOT EXISTS ciudad(
    id serial PRIMARY KEY,
    nombre text NOT NULL UNIQUE,
    id_provincia integer NOT NULL REFERENCES provincia(id)
);
