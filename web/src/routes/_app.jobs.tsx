import { createFileRoute } from "@tanstack/react-router";
import { useQuery } from "@tanstack/react-query";
import { useCallback, useMemo, useState } from "react";
import { StatusBadge } from "@/components/jobs/StatusBadge";
import type { JobStatus, SocketEvent } from "@/lib/jobs.types";
import { jobStatuses } from "@/lib/jobs.types";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Card } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Button } from "@/components/ui/button";
import { Plus, RefreshCw, AlertTriangle } from "lucide-react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { CreateJobForm } from "@/components/jobs/CreateJobForm";
import { networkService } from "@/lib/api/network.service";
import { useWsMessage } from "@/hooks/use-ws-message";

const JOBS_API_URL = "/api/jobs";

type ApiJob = {
  id: string;
  created_at: string;
  updated_at: string;
  error: string | null;
  interval: number | null;
  last_attempt_at: string | null;
  payload: Record<string, unknown>;
  priority: number;
  retry_count: number;
  scheduled_time: string;
  status: string;
  type: string;
};

async function fetchJobs(): Promise<ApiJob[]> {
  const res = await networkService.get(JOBS_API_URL);
  if (res.status === "error") throw new Error(`Request failed: ${res.message}`);
  const data = res.data as ApiJob[] | null;
  return data ?? [];
}

export const Route = createFileRoute("/_app/jobs")({
  head: () => ({ meta: [{ title: "Jobs · Jobrunner" }] }),
  component: JobsPage,
});

function fmt(iso?: string | null) {
  return iso ? new Date(iso).toLocaleString() : "—";
}

function JobsPage() {
  const { data, isLoading, error, refetch, isFetching } = useQuery({
    queryKey: ["jobs"],
    queryFn: fetchJobs,
    refetchOnWindowFocus: false,
  });
  const [status, setStatus] = useState<JobStatus | "all">("all");
  const [search, setSearch] = useState("");
  const [createOpen, setCreateOpen] = useState(false);

  const jobs = data ?? [];

  const rows = useMemo(() => {
    return jobs.filter((j) => {
      if (status !== "all" && j.status !== status) return false;
      if (search) {
        const s = search.toLowerCase();
        if (!j.id.toLowerCase().includes(s) && !j.type.toLowerCase().includes(s)) return false;
      }
      return true;
    });
  }, [jobs, status, search]);

  useWsMessage(
    useCallback((event) => {
      const msg: { event: SocketEvent } = JSON.parse(event.data);
      switch (msg.event) {
        case "job_created":
        case "job_completed":
        case "job_failed":
        case "job_queued":
        case "job_running":
        case "job_scheduled":
          refetch();
          break;
        default:
          break;
      }
    }, []),
  );

  return (
    <div className="p-8 space-y-6">
      <div className="flex items-center justify-between flex-wrap gap-3">
        <div>
          <h1 className="text-2xl font-semibold tracking-tight">Jobs</h1>
          <p className="text-sm text-muted-foreground">
            {rows.length} of {jobs.length} jobs
          </p>
        </div>
        <div className="flex items-center gap-2">
          <Button size="sm" onClick={() => setCreateOpen(true)}>
            <Plus className="h-4 w-4 mr-2" /> New job
          </Button>
        </div>
      </div>

      <Dialog open={createOpen} onOpenChange={setCreateOpen}>
        <DialogContent className="max-w-2xl">
          <DialogHeader>
            <DialogTitle>Create new job</DialogTitle>
            <DialogDescription>Configure and enqueue a new background job.</DialogDescription>
          </DialogHeader>
          <CreateJobForm
            onCreated={() => {
              setCreateOpen(false);
              refetch();
            }}
            onCancel={() => setCreateOpen(false)}
          />
        </DialogContent>
      </Dialog>

      <div className="flex items-center gap-3 flex-wrap">
        <Input
          placeholder="Search by ID or type…"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="max-w-xs"
        />
        <Select value={status} onValueChange={(v) => setStatus(v as JobStatus | "all")}>
          <SelectTrigger className="w-40">
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">All statuses</SelectItem>
            {jobStatuses.map((s) => (
              <SelectItem key={s} value={s} className="capitalize">
                {s}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>

      {error && (
        <Card className="border-red-500/40 bg-red-500/5 p-4 flex items-start gap-3">
          <AlertTriangle className="h-4 w-4 text-red-300 mt-0.5" />
          <div className="text-sm">
            <div className="font-medium text-red-200">Failed to load jobs</div>
            <div className="text-muted-foreground">
              {(error as Error).message}. Make sure the API server is running at{" "}
              <code className="font-mono">{JOBS_API_URL}</code>.
            </div>
          </div>
        </Card>
      )}

      <Card className="border-border/60 overflow-hidden">
        <Table>
          <TableHeader>
            <TableRow className="hover:bg-transparent">
              <TableHead>ID</TableHead>
              <TableHead>Type</TableHead>
              <TableHead>Priority</TableHead>
              <TableHead>Status</TableHead>
              <TableHead>Retries</TableHead>
              <TableHead>Scheduled</TableHead>
              <TableHead>Interval</TableHead>
              <TableHead>Created</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {isLoading || isFetching ? (
              <TableRow>
                <TableCell colSpan={8} className="text-center text-muted-foreground py-10">
                  Loading jobs…
                </TableCell>
              </TableRow>
            ) : rows.length === 0 ? (
              <TableRow>
                <TableCell colSpan={8} className="text-center text-muted-foreground py-10">
                  No jobs match the current filters.
                </TableCell>
              </TableRow>
            ) : (
              rows.map((j) => (
                <TableRow key={j.id}>
                  <TableCell className="font-mono text-xs">{j.id}</TableCell>
                  <TableCell>{j.type}</TableCell>
                  <TableCell className="tabular-nums">{j.priority}</TableCell>
                  <TableCell>
                    <StatusBadge status={j.status as JobStatus} />
                  </TableCell>
                  <TableCell className="tabular-nums">{`${j.retry_count}/3`}</TableCell>
                  <TableCell className="text-muted-foreground">{fmt(j.scheduled_time)}</TableCell>
                  <TableCell className="text-muted-foreground">
                    {j.interval ? `${j.interval}s` : "—"}
                  </TableCell>
                  <TableCell className="text-muted-foreground">{fmt(j.created_at)}</TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </Card>
    </div>
  );
}
