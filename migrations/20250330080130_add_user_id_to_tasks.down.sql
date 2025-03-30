DO $$
BEGIN
    ALTER TABLE tasks
        DROP CONSTRAINT IF EXISTS fk_tasks_users,
        DROP COLUMN IF EXISTS user_id;
END $$;