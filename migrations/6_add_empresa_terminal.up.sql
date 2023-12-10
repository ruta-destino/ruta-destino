CREATE TABLE IF NOT EXISTS empresa_terminal(
    id serial PRIMARY KEY,
    id_empresa integer NOT NULL REFERENCES empresa(id),
    id_terminal integer NOT NULL REFERENCES terminal(id)
);
