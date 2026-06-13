import { createFileRoute } from "@tanstack/react-router";
import { useServerFn } from "@tanstack/react-start";
import { useMutation, useQueryClient, useQuery } from "@tanstack/react-query";
import { useState } from "react";
import { retryJob } from "@/lib/jobs.functions";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog";
import { Eye, RotateCw, AlertOctagon } from "lucide-react";
import { toast } from "sonner";
import type { Job } from "@/lib/jobs.types";
import { networkService } from "@/lib/api/network.service";

const DLQ_API_URL = "/api/dlq";

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

async function fetchDLQJobs(): Promise<ApiJob[]> {
  const res = await networkService.get(DLQ_API_URL);
  if (res.status === "error") throw new Error(`Request failed: ${res.message}`);
  const data = res.data as ApiJob[] | null;
  return data ?? [];
}

export const Route = createFileRoute("/_app/dlq")({
  head: () => ({ meta: [{ title: "Dead Letter Queue · Jobrunner" }] }),
  component: DlqPage,
});

function fmt(iso?: string) {
  return iso ? new Date(iso).toLocaleString() : "—";
}

function DlqPage() {
  const { data, isLoading, error, isFetching } = useQuery({
    queryKey: ["jobs"],
    queryFn: fetchDLQJobs,
    refetchOnWindowFocus: false,
  });
  const qc = useQueryClient();
  const retryFn = useServerFn(retryJob);
  const [selected, setSelected] = useState<Job | null>(null);

  const jobs = data ?? [];

  const retryMutation = useMutation({
    mutationFn: (id: string) => retryFn({ data: { id } }),
    onSuccess: (job) => {
      toast.success("Job re-queued", { description: job.id });
      qc.invalidateQueries({ queryKey: ["dlq"] });
      qc.invalidateQueries({ queryKey: ["jobs"] });
      qc.invalidateQueries({ queryKey: ["job-counts"] });
    },
    onError: (err: Error) => toast.error("Retry failed", { description: err.message }),
  });

  return (
    <div className="p-8 space-y-6">
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-3">
          <div className="flex h-10 w-10 items-center justify-center rounded-md bg-red-500/15 text-red-300">
            <AlertOctagon className="h-5 w-5" />
          </div>
          <div>
            <h1 className="text-2xl font-semibold tracking-tight">Dead Letter Queue</h1>
            <p className="text-sm text-muted-foreground">
              {jobs.length} jobs exhausted their retries
            </p>
          </div>
        </div>
      </div>

      <Card className="border-border/60 overflow-hidden">
        <Table>
          <TableHeader>
            <TableRow className="hover:bg-transparent">
              <TableHead>ID</TableHead>
              <TableHead>Type</TableHead>
              <TableHead>Error</TableHead>
              <TableHead>Retries</TableHead>
              <TableHead>Last attempt</TableHead>
              <TableHead className="text-right">Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {isLoading || isFetching ? (
              <TableRow>
                <TableCell colSpan={8} className="text-center text-muted-foreground py-10">
                  Loading jobs…
                </TableCell>
              </TableRow>
            ) : error ? (
              <TableRow>
                <TableCell colSpan={8} className="text-center text-muted-foreground py-10">
                  <div className="text-muted-foreground">
                    {(error as Error).message}. Make sure the API server is running at{" "}
                    <code className="font-mono">{DLQ_API_URL}</code>.
                  </div>
                </TableCell>
              </TableRow>
            ) : jobs.length === 0 ? (
              <TableRow>
                <TableCell colSpan={6} className="text-center text-muted-foreground py-10">
                  No failed jobs. 🎉
                </TableCell>
              </TableRow>
            ) : (
              jobs.map((j) => (
                <TableRow key={j.id}>
                  <TableCell className="font-mono text-xs">{j.id}</TableCell>
                  <TableCell>{j.type}</TableCell>
                  <TableCell className="max-w-[28rem] truncate text-red-300">{j.error}</TableCell>
                  <TableCell className="tabular-nums">`${j.retry_count}/0`</TableCell>
                  <TableCell className="text-muted-foreground">
                    {fmt(j.last_attempt_at ?? undefined)}
                  </TableCell>
                  <TableCell className="text-right">
                    <div className="flex justify-end gap-2">
                      <Button variant="outline" size="sm" onClick={() => setSelected(j)}>
                        <Eye className="h-4 w-4 mr-1" /> Details
                      </Button>
                      <Button
                        size="sm"
                        onClick={() => retryMutation.mutate(j.id)}
                        disabled={retryMutation.isPending}
                      >
                        <RotateCw className="h-4 w-4 mr-1" /> Retry
                      </Button>
                    </div>
                  </TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </Card>

      <Dialog open={!!selected} onOpenChange={(o) => !o && setSelected(null)}>
        <DialogContent className="max-w-2xl">
          <DialogHeader>
            <DialogTitle className="font-mono text-sm">{selected?.id}</DialogTitle>
            <DialogDescription>
              {selected?.type} · last attempt {fmt(selected?.last_attempt_at)}
            </DialogDescription>
          </DialogHeader>
          {selected && (
            <div className="space-y-4">
              <div>
                <div className="text-xs uppercase text-muted-foreground mb-1">Error</div>
                <div className="text-sm text-red-300">{selected.error}</div>
              </div>
              <div>
                <div className="text-xs uppercase text-muted-foreground mb-1">Stack trace</div>
                <pre className="rounded-md bg-muted/40 border border-border/60 p-3 text-xs overflow-auto max-h-64 whitespace-pre">
                  {selected.stackTrace}
                </pre>
              </div>
              <div>
                <div className="text-xs uppercase text-muted-foreground mb-1">Payload</div>
                <pre className="rounded-md bg-muted/40 border border-border/60 p-3 text-xs overflow-auto max-h-40">
                  {JSON.stringify(selected.payload, null, 2)}
                </pre>
              </div>
              <div className="flex justify-end">
                <Button
                  onClick={() => {
                    retryMutation.mutate(selected.id);
                    setSelected(null);
                  }}
                >
                  <RotateCw className="h-4 w-4 mr-1" /> Retry job
                </Button>
              </div>
            </div>
          )}
        </DialogContent>
      </Dialog>
    </div>
  );
}
