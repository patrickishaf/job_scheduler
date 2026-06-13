import { createFileRoute, useRouter } from "@tanstack/react-router";
import { useSuspenseQuery, queryOptions } from "@tanstack/react-query";
import { getJobCounts } from "@/lib/jobs.functions";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { RefreshCw, Clock, Loader2, CheckCircle2, XCircle, CalendarClock } from "lucide-react";
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
  type ChartConfig,
} from "@/components/ui/chart";
import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from "recharts";

const countsQueryOptions = queryOptions({
  queryKey: ["job-counts"],
  queryFn: () => getJobCounts(),
});

export const Route = createFileRoute("/_app/")({
  head: () => ({ meta: [{ title: "Dashboard · Jobrunner" }] }),
  loader: ({ context }) => context.queryClient.ensureQueryData(countsQueryOptions),
  component: DashboardPage,
});

const meta = [
  { key: "queued", label: "Queued", Icon: Clock, color: "text-slate-300" },
  { key: "running", label: "Running", Icon: Loader2, color: "text-blue-300" },
  { key: "completed", label: "Completed", Icon: CheckCircle2, color: "text-emerald-300" },
  { key: "failed", label: "Failed", Icon: XCircle, color: "text-red-300" },
  { key: "scheduled", label: "Scheduled", Icon: CalendarClock, color: "text-amber-300" },
] as const;

const chartConfig: ChartConfig = {
  count: { label: "Jobs", color: "hsl(217 91% 60%)" },
};

function DashboardPage() {
  const { data } = useSuspenseQuery(countsQueryOptions);
  const router = useRouter();

  const chartData = meta.map((m) => ({
    status: m.label,
    count: data.counts[m.key],
  }));

  return (
    <div className="p-8 space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-semibold tracking-tight">Dashboard</h1>
          <p className="text-sm text-muted-foreground">{data.total} total jobs across all states</p>
        </div>
        <Button variant="outline" size="sm" onClick={() => router.invalidate()}>
          <RefreshCw className="h-4 w-4 mr-2" /> Refresh
        </Button>
      </div>

      <div className="grid gap-4 grid-cols-2 md:grid-cols-3 lg:grid-cols-5">
        {meta.map(({ key, label, Icon, color }) => (
          <Card key={key} className="border-border/60">
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium text-muted-foreground">{label}</CardTitle>
              <Icon className={`h-4 w-4 ${color}`} />
            </CardHeader>
            <CardContent>
              <div className="text-3xl font-semibold tabular-nums">{data.counts[key]}</div>
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
