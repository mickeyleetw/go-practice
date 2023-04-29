CREATE TABLE IF NOT EXISTS senao.account (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    name TEXT NOT NULL DEFAULT '' CHECK (name <> ''),
    password TEXT NOT NULL DEFAULT '' CHECK (password <> '')
);
-- CREATE INDEX IF NOT EXISTS user_name_idx ON user (name);