-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id TEXT NOT NULL UNIQUE PRIMARY KEY,

    username VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,

    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE tasks(
    id TEXT NOT NULL UNIQUE PRIMARY KEY,
    user_id TEXT NOT NULL,

    title TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,

    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS tasks CASCADE;
-- +goose StatementEnd
