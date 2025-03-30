ALTER TABLE tasks
    ADD COLUMN IF NOT EXISTS user_id INTEGER,
    ADD CONSTRAINT fk_tasks_users
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE SET NULL;