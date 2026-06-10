ALTER TABLE dead_letter_queue RENAME COLUMN interval TO interval_minutes;

ALTER TABLE jobs RENAME COLUMN interval TO interval_minutes;
