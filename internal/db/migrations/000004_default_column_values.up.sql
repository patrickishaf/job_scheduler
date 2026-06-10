ALTER TABLE jobs ALTER COLUMN status SET DEFAULT 'pending';

UPDATE jobs SET status='pending' WHERE status IS NULL;
