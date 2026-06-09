CREATE TABLE dead_letter_queue(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	attempt_count INTEGER NOT NULL DEFAULT 0,
	error TEXT,
	interval_minutes INTEGER,
	last_attempt_at TIMESTAMPTZ,
	priority INTEGER NOT NULL DEFAULT 3,
	retry_count INTEGER NOT NULL DEFAULT 0,
	scheduled_time TIMESTAMPTZ NOT NULL,
	status VARCHAR(100),
	type VARCHAR(255)
);
