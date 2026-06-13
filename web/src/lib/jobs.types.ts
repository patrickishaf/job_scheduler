import { z } from "zod";

export const jobStatuses = ["queued", "running", "completed", "failed", "scheduled"] as const;
export type JobStatus = (typeof jobStatuses)[number];

export const jobTypes = [
  "email.send",
  "report.generate",
  "cleanup.tmp",
  "image.resize",
  "webhook.deliver",
  "data.export",
] as const;

export type JsonValue =
  | string
  | number
  | boolean
  | null
  | JsonValue[]
  | { [key: string]: JsonValue };

export type Job = {
  id: string;
  type: string;
  priority: number;
  status: JobStatus;
  retry_count: number;
  maxRetries: number;
  scheduled_time: string;
  interval: number | null;
  created_at: string;
  payload: JsonValue;
  error?: string;
  stackTrace?: string;
  last_attempt_at?: string;
};

export const createJobSchema = z.object({
  type: z.string().min(1).max(100),
  priority: z.number().int().min(1).max(10),
  scheduled_time: z.string().min(1),
  interval: z
    .number()
    .int()
    .min(0)
    .max(86400 * 30)
    .nullable(),
  maxRetries: z.number().int().min(0).max(20),
  payload: z.string(), // JSON string, parsed server-side
});

export type CreateJobInput = z.infer<typeof createJobSchema>;
