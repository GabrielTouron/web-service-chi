-- +goose Up
-- +goose StatementBegin
CREATE TABLE commands (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    command text NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE commands;
-- +goose StatementEnd
