-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    user_name       TEXT PRIMARY KEY,
    user_password   TEXT NOT NULL,
    login_status    TEXT NOT NULL,
    email           TEXT NOT NULL,
    created_on      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_on      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE tasks (
    task_id         SERIAL PRIMARY KEY,
    user_name       TEXT REFERENCES users(user_name),
    task_name       TEXT NOT NULL,
    task_details    TEXT NOT NULL,
    task_status     TEXT NOT NULL,
    created_on      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_on      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_task_status ON tasks(task_status);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_user_name ON users(user_name);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
DROP TABLE users;
-- +goose StatementEnd