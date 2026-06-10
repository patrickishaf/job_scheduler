ALTER TABLE jobs RENAME COLUMN interval_minutes TO interval;

ALTER TABLE dead_letter_queue RENAME COLUMN interval_minutes TO interval;
