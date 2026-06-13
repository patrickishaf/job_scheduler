import { createServerFn } from "@tanstack/react-start";
import { z } from "zod";
import { createJobSchema, type Job, type JobStatus, type JsonValue, jobTypes } from "./jobs.types";

// In-memory mock store. Resets on worker restart — acceptable for the mock.
// Use globalThis to survive HMR in dev.
const g = globalThis as unknown as { __jobs?: Job[] };

function seed(): Job[] {
  const now = Date.now();
  const statuses: JobStatus[] = ["queued", "running", "completed", "failed", "scheduled"];
  const jobs: Job[] = [];
  for (let i = 0; i < 32; i++) {
    const status = statuses[i % statuses.length];
    const type = jobTypes[i % jobTypes.length];
    const createdAt = new Date(now - i * 1000 * 60 * 17).toISOString();
    const scheduledAt = new Date(now - i * 1000 * 60 * 10 + 1000 * 60 * 5).toISOString();
    const recurring = i % 4 === 0;
    const failed = status === "failed";
    jobs.push({
      id: `job_${(1000 + i).toString(36)}${Math.random().toString(36).slice(2, 6)}`,
      type,
      priority: ((i * 3) % 10) + 1,
      status,
      retry_count: failed ? (i % 3) + 1 : i % 2,
      maxRetries: 3,
      scheduled_time: scheduledAt,
      interval: recurring ? 3600 : null,
      created_at: createdAt,
      payload: { sample: true, index: i },
      ...(failed
        ? {
            error:
              i % 2 === 0
                ? "ECONNREFUSED: upstream service unreachable"
                : "TypeError: cannot read property 'id' of undefined",
            stackTrace: `Error: ${i % 2 === 0 ? "ECONNREFUSED" : "TypeError"}\n    at handler (/app/jobs/${type}.ts:42:11)\n    at processTicksAndRejections (node:internal/process/task_queues:96:5)\n    at async Worker.run (/app/worker.ts:88:7)`,
            last_attempt_at: new Date(now - i * 1000 * 60 * 3).toISOString(),
          }
        : {}),
    });
  }
  return jobs;
}

function store(): Job[] {
  if (!g.__jobs) g.__jobs = seed();
  return g.__jobs;
}

export const getJobCounts = createServerFn({ method: "GET" }).handler(async () => {
  const jobs = store();
  const counts: Record<JobStatus, number> = {
    queued: 0,
    running: 0,
    completed: 0,
    failed: 0,
    scheduled: 0,
  };
  for (const j of jobs) counts[j.status]++;
  return { counts, total: jobs.length };
});
