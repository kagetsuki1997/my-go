CREATE TYPE doge_type AS ENUM(
    'Doge',
    'DogeCry',
    'DogeBuffed',
    'DogeKachitoritai'
);

CREATE SCHEMA IF NOT EXISTS doge;

CREATE TABLE IF NOT EXISTS
    doge.doge (
        id UUID DEFAULT gen_random_uuid (),
        "name" TEXT NOT NULL,
        magic_number INTEGER NOT NULL,
        "type" doge_type NOT NULL,
        create_timestamp TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
        PRIMARY KEY (id)
    );

CREATE INDEX doge_doge_idx ON doge.doge (id);
