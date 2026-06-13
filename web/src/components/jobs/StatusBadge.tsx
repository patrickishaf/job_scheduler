import { Badge } from "@/components/ui/badge";
import type { JobStatus } from "@/lib/jobs.types";
import { cn } from "@/lib/utils";

const styles: Record<JobStatus, string> = {
  queued: "bg-slate-500/15 text-slate-300 border-slate-500/30",
  processing: "bg-blue-500/15 text-blue-300 border-blue-500/30",
  completed: "bg-emerald-500/15 text-emerald-300 border-emerald-500/30",
  failed: "bg-red-500/15 text-red-300 border-red-500/30",
  pending: "bg-amber-500/15 text-amber-300 border-amber-500/30",
  cancelled: "",
};

export function StatusBadge({ status }: { status: JobStatus }) {
  return (
    <Badge variant="outline" className={cn("font-medium capitalize", styles[status])}>
      {status}
    </Badge>
  );
}
