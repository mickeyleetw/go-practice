CREATE TABLE IF NOT EXISTS senao.account (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    name TEXT NOT NULL DEFAULT '' CHECK (name <> ''),
    password TEXT NOT NULL DEFAULT '' CHECK (password <> ''),
    retries INTEGER NOT NULL DEFAULT 0 CHECK (retries >= 0),
    CONSTRAINT account_name_uq UNIQUE (name)
);
CREATE INDEX IF NOT EXISTS account_name_idx ON account (name);
CREATE INDEX IF NOT EXISTS account_id_idx ON account (id);