CREATE TABLE IF NOT EXISTS recorrido(
    id serial PRIMARY KEY,
    dias bit(7) NOT NULL,
    hora smallint NOT NULL CHECK (hora >= 0 AND hora < 24),
    minuto smallint NOT NULL CHECK (minuto >= 0 AND minuto < 60),
    id_terminal_origen integer NOT NULL REFERENCES terminal(id),
    id_terminal_destino integer NOT NULL REFERENCES terminal(id),
    id_empresa integer NOT NULL REFERENCES empresa(id),
    CHECK (id_terminal_origen <> id_terminal_destino)
)
