-- +goose Up

CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    feed_id UUID NOT NULL,
    FOREIGN KEY (feed_id) REFERENCES feeds (id) ON DELETE CASCADE,
    user_id UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    UNIQUE (feed_id, user_id) 
);

-- +goose Down

DROP TABLE feed_follows;
