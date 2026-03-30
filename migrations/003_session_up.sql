CREATE TABLE sessions (
    id SERIAL PRIMARY KEY, 
    user_id INT NOT NULL, 
    expires_at TIMESTAMP NOT NULL, 
    CONSTRAINT fk_sessions_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);