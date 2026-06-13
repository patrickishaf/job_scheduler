import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { createJobSchema, jobTypes, type CreateJobInput } from "@/lib/jobs.types";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { toast } from "sonner";
import { networkService } from "@/lib/api/network.service";

const CREATE_JOB_API_URL = "/api/jobs";

type ApiJob = {
  id: string;
  type: string;
  status: string;
  priority: number;
  retry_count: number;
  scheduled_time: string;
  interval: number | null;
  created_at: string;
  payload: Record<string, unknown>;
};

function defaultScheduledAt() {
  const d = new Date();
  d.setMinutes(d.getMinutes() - d.getTimezoneOffset());
  return d.toISOString().slice(0, 16);
}

export interface CreateJobFormProps {
  onCreated?: (job: ApiJob) => void;
  onCancel?: () => void;
}

export function CreateJobForm({ onCreated, onCancel }: CreateJobFormProps) {
  const qc = useQueryClient();

  const form = useForm<CreateJobInput>({
    resolver: zodResolver(createJobSchema),
    defaultValues: {
      type: jobTypes[0],
      priority: 2,
      scheduled_time: defaultScheduledAt(),
      interval: null,
      maxRetries: 3,
      payload: '{\n  "example": true\n}',
    },
  });

  const mutation = useMutation({
    mutationFn: async (data: CreateJobInput): Promise<ApiJob> => {
      let parsedPayload: unknown;
      try {
        parsedPayload = JSON.parse(data.payload);
      } catch {
        throw new Error("Payload is not valid JSON");
      }
      if (!parsedPayload || typeof parsedPayload !== "object" || Array.isArray(parsedPayload)) {
        throw new Error("Payload must be a JSON object");
      }

      const body: Record<string, unknown> = {
        type: data.type,
        priority: data.priority,
        payload: parsedPayload,
        scheduled_time: new Date(data.scheduled_time).toISOString(),
      };
      if (data.interval && data.interval > 0) {
        body.interval_seconds = data.interval;
      }

      const res = await networkService.post(CREATE_JOB_API_URL, body);
      if (res.status === "error") {
        throw new Error(res.message);
      }
      return res.data as ApiJob;
    },
    onSuccess: (job) => {
      toast.success("Job created", { description: job.id });
      qc.invalidateQueries({ queryKey: ["jobs"] });
      qc.invalidateQueries({ queryKey: ["job-summary"] });
      form.reset();
      onCreated?.(job);
    },
    onError: (err: Error) => toast.error("Failed to create job", { description: err.message }),
  });

  return (
    <form onSubmit={form.handleSubmit((d) => mutation.mutate(d))} className="space-y-5">
      <div className="space-y-2">
        <Label htmlFor="type">Type</Label>
        <Select value={form.watch("type")} onValueChange={(v) => form.setValue("type", v)}>
          <SelectTrigger id="type">
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            {jobTypes.map((t) => (
              <SelectItem key={t} value={t}>
                {t}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
        <p className="text-xs text-muted-foreground">
          The worker handler registered for this job type.
        </p>
      </div>

      <div className="grid grid-cols-2 gap-4">
        <div className="space-y-2">
          <Label htmlFor="priority">Priority (1–3)</Label>
          <Input
            id="priority"
            type="number"
            min={1}
            max={3}
            {...form.register("priority", { valueAsNumber: true })}
          />
          <p className="text-xs text-muted-foreground">1 = lowest, 3 = highest.</p>
        </div>
        <div className="space-y-2">
          <Label htmlFor="maxRetries">Max retries</Label>
          <Input
            id="maxRetries"
            type="number"
            min={0}
            max={20}
            {...form.register("maxRetries", { valueAsNumber: true })}
          />
        </div>
      </div>

      <div className="grid grid-cols-2 gap-4">
        <div className="space-y-2">
          <Label htmlFor="scheduledAt">Scheduled at</Label>
          <Input id="scheduledAt" type="datetime-local" {...form.register("scheduled_time")} />
          <p className="text-xs text-muted-foreground">Future = scheduled, now/past = queued.</p>
        </div>
        <div className="space-y-2">
          <Label htmlFor="interval">Interval (seconds)</Label>
          <Input
            id="interval"
            type="number"
            min={0}
            placeholder="0 = one-off"
            onChange={(e) => {
              const n = e.target.valueAsNumber;
              form.setValue("interval", Number.isFinite(n) && n > 0 ? n : null);
            }}
          />
          <p className="text-xs text-muted-foreground">Re-run every N seconds.</p>
        </div>
      </div>

      <div className="space-y-2">
        <Label htmlFor="payload">Payload (JSON)</Label>
        <Textarea
          id="payload"
          rows={6}
          className="font-mono text-sm"
          {...form.register("payload")}
        />
      </div>

      <div className="flex justify-end gap-2 pt-2">
        {onCancel && (
          <Button type="button" variant="outline" onClick={onCancel}>
            Cancel
          </Button>
        )}
        <Button type="submit" disabled={mutation.isPending}>
          {mutation.isPending ? "Creating…" : "Create job"}
        </Button>
      </div>
    </form>
  );
}
