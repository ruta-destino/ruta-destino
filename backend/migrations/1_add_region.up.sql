CREATE TABLE IF NOT EXISTS region(
    id serial PRIMARY KEY,
    numero int NOT NULL UNIQUE CHECK (numero > 0),
    nombre text NOT NULL UNIQUE
);
