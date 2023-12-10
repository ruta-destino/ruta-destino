CREATE TABLE IF NOT EXISTS empresa(
    id serial PRIMARY KEY,
    nombre text NOT NULL UNIQUE
);
