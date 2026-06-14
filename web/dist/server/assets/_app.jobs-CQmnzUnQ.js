import { jsx, jsxs } from "react/jsx-runtime";
import { useQuery } from "@tanstack/react-query";
import { useState, useMemo, useCallback } from "react";
import { cva } from "class-variance-authority";
import { c as cn } from "./utils-H80jjgLf.js";
import { C as CreateJobForm, I as Input, S as Select, a as SelectTrigger, b as SelectValue, c as SelectContent, d as SelectItem, j as jobStatuses } from "./CreateJobForm-JskyILMC.js";
import { u as useWsMessage, D as Dialog, a as DialogContent, b as DialogHeader, c as DialogTitle, d as DialogDescription, T as Table, e as TableHeader, f as TableRow, g as TableHead, h as TableBody, i as TableCell } from "./use-ws-message-DzgHz9Ac.js";
import { C as Card, n as networkService } from "./network.service-Vg4jAWs3.js";
import { B as Button } from "./button-BC9oXVxV.js";
import { Plus, AlertTriangle } from "lucide-react";
import "clsx";
import "tailwind-merge";
import "react-hook-form";
import "@hookform/resolvers/zod";
import "zod";
import "@radix-ui/react-label";
import "@radix-ui/react-select";
import "sonner";
import "@radix-ui/react-dialog";
import "./router-CpxjGD0j.js";
import "@tanstack/react-router";
import "axios";
import "@radix-ui/react-slot";
const badgeVariants = cva(
  "inline-flex items-center rounded-md border px-2.5 py-0.5 text-xs font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2",
  {
    variants: {
      variant: {
        default: "border-transparent bg-primary text-primary-foreground shadow hover:bg-primary/80",
        secondary: "border-transparent bg-secondary text-secondary-foreground hover:bg-secondary/80",
        destructive: "border-transparent bg-destructive text-destructive-foreground shadow hover:bg-destructive/80",
        outline: "text-foreground"
      }
    },
    defaultVariants: {
      variant: "default"
    }
  }
);
function Badge({ className, variant, ...props }) {
  return /* @__PURE__ */ jsx("div", { className: cn(badgeVariants({ variant }), className), ...props });
}
const styles = {
  queued: "bg-slate-500/15 text-slate-300 border-slate-500/30",
  processing: "bg-blue-500/15 text-blue-300 border-blue-500/30",
  completed: "bg-emerald-500/15 text-emerald-300 border-emerald-500/30",
  failed: "bg-red-500/15 text-red-300 border-red-500/30",
  pending: "bg-amber-500/15 text-amber-300 border-amber-500/30",
  cancelled: ""
};
function StatusBadge({ status }) {
  return /* @__PURE__ */ jsx(Badge, { variant: "outline", className: cn("font-medium capitalize", styles[status]), children: status });
}
const JOBS_API_URL = "/api/jobs";
async function fetchJobs() {
  const res = await networkService.get(JOBS_API_URL);
  if (res.status === "error") throw new Error(`Request failed: ${res.message}`);
  const data = res.data;
  return data ?? [];
}
function fmt(iso) {
  return iso ? new Date(iso).toLocaleString() : "—";
}
function JobsPage() {
  const {
    data,
    isLoading,
    error,
    refetch,
    isFetching
  } = useQuery({
    queryKey: ["jobs"],
    queryFn: fetchJobs,
    refetchOnWindowFocus: false
  });
  const [status, setStatus] = useState("all");
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
  useWsMessage(useCallback((event) => {
    const msg = JSON.parse(event.data);
    switch (msg.event) {
      case "job_created":
      case "job_completed":
      case "job_failed":
      case "job_queued":
      case "job_running":
      case "job_scheduled":
        refetch();
        break;
    }
  }, []));
  return /* @__PURE__ */ jsxs("div", { className: "p-8 space-y-6", children: [
    /* @__PURE__ */ jsxs("div", { className: "flex items-center justify-between flex-wrap gap-3", children: [
      /* @__PURE__ */ jsxs("div", { children: [
        /* @__PURE__ */ jsx("h1", { className: "text-2xl font-semibold tracking-tight", children: "Jobs" }),
        /* @__PURE__ */ jsxs("p", { className: "text-sm text-muted-foreground", children: [
          rows.length,
          " of ",
          jobs.length,
          " jobs"
        ] })
      ] }),
      /* @__PURE__ */ jsx("div", { className: "flex items-center gap-2", children: /* @__PURE__ */ jsxs(Button, { size: "sm", onClick: () => setCreateOpen(true), children: [
        /* @__PURE__ */ jsx(Plus, { className: "h-4 w-4 mr-2" }),
        " New job"
      ] }) })
    ] }),
    /* @__PURE__ */ jsx(Dialog, { open: createOpen, onOpenChange: setCreateOpen, children: /* @__PURE__ */ jsxs(DialogContent, { className: "max-w-2xl", children: [
      /* @__PURE__ */ jsxs(DialogHeader, { children: [
        /* @__PURE__ */ jsx(DialogTitle, { children: "Create new job" }),
        /* @__PURE__ */ jsx(DialogDescription, { children: "Configure and enqueue a new background job." })
      ] }),
      /* @__PURE__ */ jsx(CreateJobForm, { onCreated: () => {
        setCreateOpen(false);
        refetch();
      }, onCancel: () => setCreateOpen(false) })
    ] }) }),
    /* @__PURE__ */ jsxs("div", { className: "flex items-center gap-3 flex-wrap", children: [
      /* @__PURE__ */ jsx(Input, { placeholder: "Search by ID or type…", value: search, onChange: (e) => setSearch(e.target.value), className: "max-w-xs" }),
      /* @__PURE__ */ jsxs(Select, { value: status, onValueChange: (v) => setStatus(v), children: [
        /* @__PURE__ */ jsx(SelectTrigger, { className: "w-40", children: /* @__PURE__ */ jsx(SelectValue, {}) }),
        /* @__PURE__ */ jsxs(SelectContent, { children: [
          /* @__PURE__ */ jsx(SelectItem, { value: "all", children: "All statuses" }),
          jobStatuses.map((s) => /* @__PURE__ */ jsx(SelectItem, { value: s, className: "capitalize", children: s }, s))
        ] })
      ] })
    ] }),
    error && /* @__PURE__ */ jsxs(Card, { className: "border-red-500/40 bg-red-500/5 p-4 flex items-start gap-3", children: [
      /* @__PURE__ */ jsx(AlertTriangle, { className: "h-4 w-4 text-red-300 mt-0.5" }),
      /* @__PURE__ */ jsxs("div", { className: "text-sm", children: [
        /* @__PURE__ */ jsx("div", { className: "font-medium text-red-200", children: "Failed to load jobs" }),
        /* @__PURE__ */ jsxs("div", { className: "text-muted-foreground", children: [
          error.message,
          ". Make sure the API server is running at",
          " ",
          /* @__PURE__ */ jsx("code", { className: "font-mono", children: JOBS_API_URL }),
          "."
        ] })
      ] })
    ] }),
    /* @__PURE__ */ jsx(Card, { className: "border-border/60 overflow-hidden", children: /* @__PURE__ */ jsxs(Table, { children: [
      /* @__PURE__ */ jsx(TableHeader, { children: /* @__PURE__ */ jsxs(TableRow, { className: "hover:bg-transparent", children: [
        /* @__PURE__ */ jsx(TableHead, { children: "ID" }),
        /* @__PURE__ */ jsx(TableHead, { children: "Type" }),
        /* @__PURE__ */ jsx(TableHead, { children: "Priority" }),
        /* @__PURE__ */ jsx(TableHead, { children: "Status" }),
        /* @__PURE__ */ jsx(TableHead, { children: "Retries" }),
        /* @__PURE__ */ jsx(TableHead, { children: "Scheduled" }),
        /* @__PURE__ */ jsx(TableHead, { children: "Interval" }),
        /* @__PURE__ */ jsx(TableHead, { children: "Created" })
      ] }) }),
      /* @__PURE__ */ jsx(TableBody, { children: isLoading || isFetching ? /* @__PURE__ */ jsx(TableRow, { children: /* @__PURE__ */ jsx(TableCell, { colSpan: 8, className: "text-center text-muted-foreground py-10", children: "Loading jobs…" }) }) : rows.length === 0 ? /* @__PURE__ */ jsx(TableRow, { children: /* @__PURE__ */ jsx(TableCell, { colSpan: 8, className: "text-center text-muted-foreground py-10", children: "No jobs match the current filters." }) }) : rows.map((j) => /* @__PURE__ */ jsxs(TableRow, { children: [
        /* @__PURE__ */ jsx(TableCell, { className: "font-mono text-xs", children: j.id }),
        /* @__PURE__ */ jsx(TableCell, { children: j.type }),
        /* @__PURE__ */ jsx(TableCell, { className: "tabular-nums", children: j.priority }),
        /* @__PURE__ */ jsx(TableCell, { children: /* @__PURE__ */ jsx(StatusBadge, { status: j.status }) }),
        /* @__PURE__ */ jsx(TableCell, { className: "tabular-nums", children: `${j.retry_count}/3` }),
        /* @__PURE__ */ jsx(TableCell, { className: "text-muted-foreground", children: fmt(j.scheduled_time) }),
        /* @__PURE__ */ jsx(TableCell, { className: "text-muted-foreground", children: j.interval ? `${j.interval}s` : "—" }),
        /* @__PURE__ */ jsx(TableCell, { className: "text-muted-foreground", children: fmt(j.created_at) })
      ] }, j.id)) })
    ] }) })
  ] });
}
export {
  JobsPage as component
};
