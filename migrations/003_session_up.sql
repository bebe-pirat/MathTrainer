CREATE TABLE sessions (
    id SERIAL PRIMARY KEY, 
    user_id INT, 
    expires_at TIMESTAMP, 
    FOREIGN KEY user_id REFERENCES users(id)
);