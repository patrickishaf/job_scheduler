import { createFileRoute } from "@tanstack/react-router";
import { useQuery } from "@tanstack/react-query";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import {
  RefreshCw,
  Clock,
  Loader2,
  CheckCircle2,
  XCircle,
  CalendarClock,
  Ban,
  Hourglass,
} from "lucide-react";
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
  type ChartConfig,
} from "@/components/ui/chart";
import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from "recharts";
import { networkService } from "@/lib/api/network.service";

const SUMMARY_API_URL = "/api/jobs/summary";

type JobsSummary = {
  queued: number;
  processing: number;
  cancelled: number;
  completed: number;
  failed: number;
  pending: number;
  scheduled: number;
};

async function fetchSummary(): Promise<JobsSummary> {
  const res = await networkService.get(SUMMARY_API_URL);
  if (res.status === "error") throw new Error(`HTTP ${res.res.message}`);
  return res.data as JobsSummary;
}

export const Route = createFileRoute("/_app/")({
  head: () => ({ meta: [{ title: "Dashboard · Jobrunner" }] }),
  component: DashboardPage,
});

const meta = [
  { key: "queued", label: "Queued", Icon: Clock, color: "text-slate-300" },
  { key: "pending", label: "Pending", Icon: Hourglass, color: "text-slate-300" },
  { key: "processing", label: "Processing", Icon: Loader2, color: "text-blue-300" },
  { key: "scheduled", label: "Scheduled", Icon: CalendarClock, color: "text-amber-300" },
  { key: "completed", label: "Completed", Icon: CheckCircle2, color: "text-emerald-300" },
  { key: "failed", label: "Failed", Icon: XCircle, color: "text-red-300" },
  { key: "cancelled", label: "Cancelled", Icon: Ban, color: "text-muted-foreground" },
] as const;

const chartConfig: ChartConfig = {
  count: { label: "Jobs", color: "hsl(217 91% 60%)" },
};

function DashboardPage() {
  const { data, isLoading, error, refetch, isFetching } = useQuery({
    queryKey: ["job-summary"],
    queryFn: fetchSummary,
    refetchInterval: 5000,
  });

  const total = data ? meta.reduce((sum, m) => sum + (data[m.key] ?? 0), 0) : 0;

  const chartData = meta.map((m) => ({
    status: m.label,
    count: data?.[m.key] ?? 0,
  }));

  return (
    <div className="p-8 space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-semibold tracking-tight">Dashboard</h1>
          <p className="text-sm text-muted-foreground">
            {isLoading ? "Loading…" : `${total} total jobs across all states`}
          </p>
        </div>
      </div>

      {error ? (
        <Card className="border-red-500/40 bg-red-500/5">
          <CardHeader>
            <CardTitle className="text-red-300">Failed to load summary</CardTitle>
          </CardHeader>
          <CardContent className="text-sm text-muted-foreground space-y-1">
            <div>{(error as Error).message}</div>
            <div className="font-mono text-xs">{SUMMARY_API_URL}</div>
          </CardContent>
        </Card>
      ) : null}

      <div className="grid gap-4 grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-7">
        {meta.map(({ key, label, Icon, color }) => (
          <Card key={key} className="border-border/60">
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium text-muted-foreground">{label}</CardTitle>
              <Icon className={`h-4 w-4 ${color}`} />
            </CardHeader>
            <CardContent>
              <div className="text-3xl font-semibold tabular-nums">{data?.[key] ?? 0}</div>
            </CardContent>
          </Card>
        ))}
      </div>

      <Card className="border-border/60">
        <CardHeader>
          <CardTitle>Jobs by status</CardTitle>
        </CardHeader>
        <CardContent>
          <ChartContainer config={chartConfig} className="h-72 w-full">
            <BarChart data={chartData}>
              <CartesianGrid strokeDasharray="3 3" vertical={false} className="stroke-border/50" />
              <XAxis dataKey="status" tickLine={false} axisLine={false} />
              <YAxis tickLine={false} axisLine={false} allowDecimals={false} />
              <ChartTooltip content={<ChartTooltipContent />} />
              <Bar dataKey="count" fill="var(--color-count)" radius={[6, 6, 0, 0]} />
            </BarChart>
          </ChartContainer>
        </CardContent>
      </Card>
    </div>
  );
}
