import { jsxs, jsx } from "react/jsx-runtime";
import { useQuery, useQueryClient, useMutation } from "@tanstack/react-query";
import { useState, useCallback } from "react";
import { C as Card, n as networkService } from "./network.service-Vg4jAWs3.js";
import { B as Button } from "./button-BC9oXVxV.js";
import { u as useWsMessage, T as Table, e as TableHeader, f as TableRow, g as TableHead, h as TableBody, i as TableCell, D as Dialog, a as DialogContent, b as DialogHeader, c as DialogTitle, d as DialogDescription } from "./use-ws-message-DzgHz9Ac.js";
import { AlertOctagon, Eye, RotateCw } from "lucide-react";
import { toast } from "sonner";
import "./utils-H80jjgLf.js";
import "clsx";
import "tailwind-merge";
import "axios";
import "@radix-ui/react-slot";
import "class-variance-authority";
import "@radix-ui/react-dialog";
import "./router-CpxjGD0j.js";
import "@tanstack/react-router";
const DLQ_API_URL = "/api/dlq";
async function fetchDLQJobs() {
  const res = await networkService.get(DLQ_API_URL);
  if (res.status === "error") throw new Error(`Request failed: ${res.message}`);
  const data = res.data;
  return data ?? [];
}
async function requeueDLQJob(jobID) {
  const res = await networkService.patch(`/api/jobs/${jobID}/requeue`);
  if (res.status === "error") throw new Error(`Request failed: ${res.message}`);
  const data = res.data;
  return data;
}
function fmt(iso) {
  return iso ? new Date(iso).toLocaleString() : "—";
}
function DlqPage() {
  const {
    data,
    isLoading,
    error,
    refetch: refetchDLQ,
    isFetching
  } = useQuery({
    queryKey: ["jobs"],
    queryFn: fetchDLQJobs,
    refetchOnWindowFocus: false
  });
  const qc = useQueryClient();
  const [selected, setSelected] = useState(null);
  const jobs = data ?? [];
  const retryMutation = useMutation({
    mutationFn: (id) => requeueDLQJob(id),
    onSuccess: (job) => {
      toast.success("Job re-queued", {
        description: job.id
      });
      qc.invalidateQueries({
        queryKey: ["dlq"]
      });
      qc.invalidateQueries({
        queryKey: ["jobs"]
      });
      qc.invalidateQueries({
        queryKey: ["job-counts"]
      });
    },
    onError: (err) => toast.error("Retry failed", {
      description: err.message
    })
  });
  useWsMessage(useCallback((event) => {
    const msg = JSON.parse(event.data);
    switch (msg.event) {
      case "dlq_updated":
        refetchDLQ();
        break;
    }
  }, []));
  return /* @__PURE__ */ jsxs("div", { className: "p-8 space-y-6", children: [
    /* @__PURE__ */ jsx("div", { className: "flex items-center justify-between", children: /* @__PURE__ */ jsxs("div", { className: "flex items-center gap-3", children: [
      /* @__PURE__ */ jsx("div", { className: "flex h-10 w-10 items-center justify-center rounded-md bg-red-500/15 text-red-300", children: /* @__PURE__ */ jsx(AlertOctagon, { className: "h-5 w-5" }) }),
      /* @__PURE__ */ jsxs("div", { children: [
        /* @__PURE__ */ jsx("h1", { className: "text-2xl font-semibold tracking-tight", children: "Dead Letter Queue" }),
        /* @__PURE__ */ jsxs("p", { className: "text-sm text-muted-foreground", children: [
          jobs.length,
          " jobs exhausted their retries"
        ] })
      ] })
    ] }) }),
    /* @__PURE__ */ jsx(Card, { className: "border-border/60 overflow-hidden", children: /* @__PURE__ */ jsxs(Table, { children: [
      /* @__PURE__ */ jsx(TableHeader, { children: /* @__PURE__ */ jsxs(TableRow, { className: "hover:bg-transparent", children: [
        /* @__PURE__ */ jsx(TableHead, { children: "ID" }),
        /* @__PURE__ */ jsx(TableHead, { children: "Type" }),
        /* @__PURE__ */ jsx(TableHead, { children: "Error" }),
        /* @__PURE__ */ jsx(TableHead, { children: "Retries" }),
        /* @__PURE__ */ jsx(TableHead, { children: "Last attempt" }),
        /* @__PURE__ */ jsx(TableHead, { className: "text-right", children: "Actions" })
      ] }) }),
      /* @__PURE__ */ jsx(TableBody, { children: isLoading || isFetching ? /* @__PURE__ */ jsx(TableRow, { children: /* @__PURE__ */ jsx(TableCell, { colSpan: 8, className: "text-center text-muted-foreground py-10", children: "Loading jobs…" }) }) : error ? /* @__PURE__ */ jsx(TableRow, { children: /* @__PURE__ */ jsx(TableCell, { colSpan: 8, className: "text-center text-muted-foreground py-10", children: /* @__PURE__ */ jsxs("div", { className: "text-muted-foreground", children: [
        error.message,
        ". Make sure the API server is running at",
        " ",
        /* @__PURE__ */ jsx("code", { className: "font-mono", children: DLQ_API_URL }),
        "."
      ] }) }) }) : jobs.length === 0 ? /* @__PURE__ */ jsx(TableRow, { children: /* @__PURE__ */ jsx(TableCell, { colSpan: 6, className: "text-center text-muted-foreground py-10", children: "No failed jobs. 🎉" }) }) : jobs.map((j) => /* @__PURE__ */ jsxs(TableRow, { children: [
        /* @__PURE__ */ jsx(TableCell, { className: "font-mono text-xs", children: j.id }),
        /* @__PURE__ */ jsx(TableCell, { children: j.type }),
        /* @__PURE__ */ jsx(TableCell, { className: "max-w-[28rem] truncate text-red-300", children: j.error }),
        /* @__PURE__ */ jsxs(TableCell, { className: "tabular-nums", children: [
          j.retry_count,
          "/3"
        ] }),
        /* @__PURE__ */ jsx(TableCell, { className: "text-muted-foreground", children: fmt(j.last_attempt_at ?? void 0) }),
        /* @__PURE__ */ jsx(TableCell, { className: "text-right", children: /* @__PURE__ */ jsxs("div", { className: "flex justify-end gap-2", children: [
          /* @__PURE__ */ jsxs(Button, { variant: "outline", size: "sm", onClick: () => setSelected(j), children: [
            /* @__PURE__ */ jsx(Eye, { className: "h-4 w-4 mr-1" }),
            " Details"
          ] }),
          /* @__PURE__ */ jsxs(Button, { size: "sm", onClick: () => retryMutation.mutate(j.id), disabled: retryMutation.isPending, children: [
            /* @__PURE__ */ jsx(RotateCw, { className: "h-4 w-4 mr-1" }),
            " Retry"
          ] })
        ] }) })
      ] }, j.id)) })
    ] }) }),
    /* @__PURE__ */ jsx(Dialog, { open: !!selected, onOpenChange: (o) => !o && setSelected(null), children: /* @__PURE__ */ jsxs(DialogContent, { className: "max-w-2xl", children: [
      /* @__PURE__ */ jsxs(DialogHeader, { children: [
        /* @__PURE__ */ jsx(DialogTitle, { className: "font-mono text-sm", children: selected?.id }),
        /* @__PURE__ */ jsxs(DialogDescription, { children: [
          selected?.type,
          " · last attempt ",
          fmt(selected?.last_attempt_at ?? void 0)
        ] })
      ] }),
      selected && /* @__PURE__ */ jsxs("div", { className: "space-y-4", children: [
        /* @__PURE__ */ jsxs("div", { children: [
          /* @__PURE__ */ jsx("div", { className: "text-xs uppercase text-muted-foreground mb-1", children: "Error" }),
          /* @__PURE__ */ jsx("div", { className: "text-sm text-red-300", children: selected.error })
        ] }),
        /* @__PURE__ */ jsxs("div", { children: [
          /* @__PURE__ */ jsx("div", { className: "text-xs uppercase text-muted-foreground mb-1", children: "Stack trace" }),
          /* @__PURE__ */ jsx("pre", { className: "rounded-md bg-muted/40 border border-border/60 p-3 text-xs overflow-auto max-h-64 whitespace-pre", children: "stacktrace not available" })
        ] }),
        /* @__PURE__ */ jsxs("div", { children: [
          /* @__PURE__ */ jsx("div", { className: "text-xs uppercase text-muted-foreground mb-1", children: "Payload" }),
          /* @__PURE__ */ jsx("pre", { className: "rounded-md bg-muted/40 border border-border/60 p-3 text-xs overflow-auto max-h-40", children: JSON.stringify(selected.payload, null, 2) })
        ] }),
        /* @__PURE__ */ jsx("div", { className: "flex justify-end", children: /* @__PURE__ */ jsxs(Button, { onClick: () => {
          retryMutation.mutate(selected.id);
          setSelected(null);
        }, children: [
          /* @__PURE__ */ jsx(RotateCw, { className: "h-4 w-4 mr-1" }),
          " Retry job"
        ] }) })
      ] })
    ] }) })
  ] });
}
export {
  DlqPage as component
};
