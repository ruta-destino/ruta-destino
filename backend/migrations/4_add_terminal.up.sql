CREATE TABLE IF NOT EXISTS terminal(
    id serial PRIMARY KEY,
    nombre text NOT NULL UNIQUE,
    longitud double precision NOT NULL,
    latitud double precision NOT NULL,
    direccion text NOT NULL,
    id_ciudad integer NOT NULL REFERENCES ciudad(id)
);
