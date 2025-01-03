-- +goose Up
CREATE TABLE posts (
   id UUID PRIMARY KEY, 
   created_at TIMESTAMP NOT NULL, 
   updated_at TIMESTAMP NOT NULL, 
   user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE, 
   body TEXT NOT NULL,
   likes INT NOT NULL
);

-- +goose Down
DROP TABLE posts;
