import { jsx, jsxs, Fragment } from "react/jsx-runtime";
import { useQuery } from "@tanstack/react-query";
import { C as Card, a as CardHeader, b as CardTitle, c as CardContent, n as networkService } from "./network.service-Vg4jAWs3.js";
import { Clock, Hourglass, Loader2, CalendarClock, CheckCircle2, XCircle, Ban } from "lucide-react";
import * as React from "react";
import * as RechartsPrimitive from "recharts";
import { BarChart, CartesianGrid, XAxis, YAxis, Bar } from "recharts";
import { c as cn } from "./utils-H80jjgLf.js";
import "axios";
import "clsx";
import "tailwind-merge";
const THEMES = { light: "", dark: ".dark" };
const ChartContext = React.createContext(null);
function useChart() {
  const context = React.useContext(ChartContext);
  if (!context) {
    throw new Error("useChart must be used within a <ChartContainer />");
  }
  return context;
}
const ChartContainer = React.forwardRef(({ id, className, children, config, ...props }, ref) => {
  const uniqueId = React.useId();
  const chartId = `chart-${id || uniqueId.replace(/:/g, "")}`;
  return /* @__PURE__ */ jsx(ChartContext.Provider, { value: { config }, children: /* @__PURE__ */ jsxs(
    "div",
    {
      "data-chart": chartId,
      ref,
      className: cn(
        "flex aspect-video justify-center text-xs [&_.recharts-cartesian-axis-tick_text]:fill-muted-foreground [&_.recharts-cartesian-grid_line[stroke='#ccc']]:stroke-border/50 [&_.recharts-curve.recharts-tooltip-cursor]:stroke-border [&_.recharts-dot[stroke='#fff']]:stroke-transparent [&_.recharts-layer]:outline-none [&_.recharts-polar-grid_[stroke='#ccc']]:stroke-border [&_.recharts-radial-bar-background-sector]:fill-muted [&_.recharts-rectangle.recharts-tooltip-cursor]:fill-muted [&_.recharts-reference-line_[stroke='#ccc']]:stroke-border [&_.recharts-sector[stroke='#fff']]:stroke-transparent [&_.recharts-sector]:outline-none [&_.recharts-surface]:outline-none",
        className
      ),
      ...props,
      children: [
        /* @__PURE__ */ jsx(ChartStyle, { id: chartId, config }),
        /* @__PURE__ */ jsx(RechartsPrimitive.ResponsiveContainer, { children })
      ]
    }
  ) });
});
ChartContainer.displayName = "Chart";
const ChartStyle = ({ id, config }) => {
  const colorConfig = Object.entries(config).filter(([, config2]) => config2.theme || config2.color);
  if (!colorConfig.length) {
    return null;
  }
  return /* @__PURE__ */ jsx(
    "style",
    {
      dangerouslySetInnerHTML: {
        __html: Object.entries(THEMES).map(
          ([theme, prefix]) => `
${prefix} [data-chart=${id}] {
${colorConfig.map(([key, itemConfig]) => {
            const color = itemConfig.theme?.[theme] || itemConfig.color;
            return color ? `  --color-${key}: ${color};` : null;
          }).join("\n")}
}
`
        ).join("\n")
      }
    }
  );
};
const ChartTooltip = RechartsPrimitive.Tooltip;
const ChartTooltipContent = React.forwardRef(
  ({
    active,
    payload,
    className,
    indicator = "dot",
    hideLabel = false,
    hideIndicator = false,
    label,
    labelFormatter,
    labelClassName,
    formatter,
    color,
    nameKey,
    labelKey
  }, ref) => {
    const { config } = useChart();
    const tooltipLabel = React.useMemo(() => {
      if (hideLabel || !payload?.length) {
        return null;
      }
      const [item] = payload;
      const key = `${labelKey || item?.dataKey || item?.name || "value"}`;
      const itemConfig = getPayloadConfigFromPayload(config, item, key);
      const value = !labelKey && typeof label === "string" ? config[label]?.label || label : itemConfig?.label;
      if (labelFormatter) {
        return /* @__PURE__ */ jsx("div", { className: cn("font-medium", labelClassName), children: labelFormatter(value, payload) });
      }
      if (!value) {
        return null;
      }
      return /* @__PURE__ */ jsx("div", { className: cn("font-medium", labelClassName), children: value });
    }, [label, labelFormatter, payload, hideLabel, labelClassName, config, labelKey]);
    if (!active || !payload?.length) {
      return null;
    }
    const nestLabel = payload.length === 1 && indicator !== "dot";
    return /* @__PURE__ */ jsxs(
      "div",
      {
        ref,
        className: cn(
          "grid min-w-[8rem] items-start gap-1.5 rounded-lg border border-border/50 bg-background px-2.5 py-1.5 text-xs shadow-xl",
          className
        ),
        children: [
          !nestLabel ? tooltipLabel : null,
          /* @__PURE__ */ jsx("div", { className: "grid gap-1.5", children: payload.filter((item) => item.type !== "none").map((item, index) => {
            const key = `${nameKey || item.name || item.dataKey || "value"}`;
            const itemConfig = getPayloadConfigFromPayload(config, item, key);
            const indicatorColor = color || item.payload.fill || item.color;
            return /* @__PURE__ */ jsx(
              "div",
              {
                className: cn(
                  "flex w-full flex-wrap items-stretch gap-2 [&>svg]:h-2.5 [&>svg]:w-2.5 [&>svg]:text-muted-foreground",
                  indicator === "dot" && "items-center"
                ),
                children: formatter && item?.value !== void 0 && item.name ? formatter(item.value, item.name, item, index, item.payload) : /* @__PURE__ */ jsxs(Fragment, { children: [
                  itemConfig?.icon ? /* @__PURE__ */ jsx(itemConfig.icon, {}) : !hideIndicator && /* @__PURE__ */ jsx(
                    "div",
                    {
                      className: cn(
                        "shrink-0 rounded-[2px] border-(--color-border) bg-(--color-bg)",
                        {
                          "h-2.5 w-2.5": indicator === "dot",
                          "w-1": indicator === "line",
                          "w-0 border-[1.5px] border-dashed bg-transparent": indicator === "dashed",
                          "my-0.5": nestLabel && indicator === "dashed"
                        }
                      ),
                      style: {
                        "--color-bg": indicatorColor,
                        "--color-border": indicatorColor
                      }
                    }
                  ),
                  /* @__PURE__ */ jsxs(
                    "div",
                    {
                      className: cn(
                        "flex flex-1 justify-between leading-none",
                        nestLabel ? "items-end" : "items-center"
                      ),
                      children: [
                        /* @__PURE__ */ jsxs("div", { className: "grid gap-1.5", children: [
                          nestLabel ? tooltipLabel : null,
                          /* @__PURE__ */ jsx("span", { className: "text-muted-foreground", children: itemConfig?.label || item.name })
                        ] }),
                        item.value && /* @__PURE__ */ jsx("span", { className: "font-mono font-medium tabular-nums text-foreground", children: item.value.toLocaleString() })
                      ]
                    }
                  )
                ] })
              },
              item.dataKey
            );
          }) })
        ]
      }
    );
  }
);
ChartTooltipContent.displayName = "ChartTooltip";
const ChartLegendContent = React.forwardRef(({ className, hideIcon = false, payload, verticalAlign = "bottom", nameKey }, ref) => {
  const { config } = useChart();
  if (!payload?.length) {
    return null;
  }
  return /* @__PURE__ */ jsx(
    "div",
    {
      ref,
      className: cn(
        "flex items-center justify-center gap-4",
        verticalAlign === "top" ? "pb-3" : "pt-3",
        className
      ),
      children: payload.filter((item) => item.type !== "none").map((item) => {
        const key = `${nameKey || item.dataKey || "value"}`;
        const itemConfig = getPayloadConfigFromPayload(config, item, key);
        return /* @__PURE__ */ jsxs(
          "div",
          {
            className: cn(
              "flex items-center gap-1.5 [&>svg]:h-3 [&>svg]:w-3 [&>svg]:text-muted-foreground"
            ),
            children: [
              itemConfig?.icon && !hideIcon ? /* @__PURE__ */ jsx(itemConfig.icon, {}) : /* @__PURE__ */ jsx(
                "div",
                {
                  className: "h-2 w-2 shrink-0 rounded-[2px]",
                  style: {
                    backgroundColor: item.color
                  }
                }
              ),
              itemConfig?.label
            ]
          },
          item.value
        );
      })
    }
  );
});
ChartLegendContent.displayName = "ChartLegend";
function getPayloadConfigFromPayload(config, payload, key) {
  if (typeof payload !== "object" || payload === null) {
    return void 0;
  }
  const payloadPayload = "payload" in payload && typeof payload.payload === "object" && payload.payload !== null ? payload.payload : void 0;
  let configLabelKey = key;
  if (key in payload && typeof payload[key] === "string") {
    configLabelKey = payload[key];
  } else if (payloadPayload && key in payloadPayload && typeof payloadPayload[key] === "string") {
    configLabelKey = payloadPayload[key];
  }
  return configLabelKey in config ? config[configLabelKey] : config[key];
}
const SUMMARY_API_URL = "/api/jobs/summary";
async function fetchSummary() {
  const res = await networkService.get(SUMMARY_API_URL);
  if (res.status === "error") throw new Error(`HTTP ${res.res.message}`);
  return res.data;
}
const meta = [{
  key: "queued",
  label: "Queued",
  Icon: Clock,
  color: "text-slate-300"
}, {
  key: "pending",
  label: "Pending",
  Icon: Hourglass,
  color: "text-slate-300"
}, {
  key: "processing",
  label: "Processing",
  Icon: Loader2,
  color: "text-blue-300"
}, {
  key: "scheduled",
  label: "Scheduled",
  Icon: CalendarClock,
  color: "text-amber-300"
}, {
  key: "completed",
  label: "Completed",
  Icon: CheckCircle2,
  color: "text-emerald-300"
}, {
  key: "failed",
  label: "Failed",
  Icon: XCircle,
  color: "text-red-300"
}, {
  key: "cancelled",
  label: "Cancelled",
  Icon: Ban,
  color: "text-muted-foreground"
}];
const chartConfig = {
  count: {
    label: "Jobs",
    color: "hsl(217 91% 60%)"
  }
};
function DashboardPage() {
  const {
    data,
    isLoading,
    error,
    refetch,
    isFetching
  } = useQuery({
    queryKey: ["job-summary"],
    queryFn: fetchSummary,
    refetchInterval: 5e3
  });
  const total = data ? meta.reduce((sum, m) => sum + (data[m.key] ?? 0), 0) : 0;
  const chartData = meta.map((m) => ({
    status: m.label,
    count: data?.[m.key] ?? 0
  }));
  return /* @__PURE__ */ jsxs("div", { className: "p-8 space-y-6", children: [
    /* @__PURE__ */ jsx("div", { className: "flex items-center justify-between", children: /* @__PURE__ */ jsxs("div", { children: [
      /* @__PURE__ */ jsx("h1", { className: "text-2xl font-semibold tracking-tight", children: "Dashboard" }),
      /* @__PURE__ */ jsx("p", { className: "text-sm text-muted-foreground", children: isLoading ? "Loading…" : `${total} total jobs across all states` })
    ] }) }),
    error ? /* @__PURE__ */ jsxs(Card, { className: "border-red-500/40 bg-red-500/5", children: [
      /* @__PURE__ */ jsx(CardHeader, { children: /* @__PURE__ */ jsx(CardTitle, { className: "text-red-300", children: "Failed to load summary" }) }),
      /* @__PURE__ */ jsxs(CardContent, { className: "text-sm text-muted-foreground space-y-1", children: [
        /* @__PURE__ */ jsx("div", { children: error.message }),
        /* @__PURE__ */ jsx("div", { className: "font-mono text-xs", children: SUMMARY_API_URL })
      ] })
    ] }) : null,
    /* @__PURE__ */ jsx("div", { className: "grid gap-4 grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-7", children: meta.map(({
      key,
      label,
      Icon,
      color
    }) => /* @__PURE__ */ jsxs(Card, { className: "border-border/60", children: [
      /* @__PURE__ */ jsxs(CardHeader, { className: "flex flex-row items-center justify-between space-y-0 pb-2", children: [
        /* @__PURE__ */ jsx(CardTitle, { className: "text-sm font-medium text-muted-foreground", children: label }),
        /* @__PURE__ */ jsx(Icon, { className: `h-4 w-4 ${color}` })
      ] }),
      /* @__PURE__ */ jsx(CardContent, { children: /* @__PURE__ */ jsx("div", { className: "text-3xl font-semibold tabular-nums", children: data?.[key] ?? 0 }) })
    ] }, key)) }),
    /* @__PURE__ */ jsxs(Card, { className: "border-border/60", children: [
      /* @__PURE__ */ jsx(CardHeader, { children: /* @__PURE__ */ jsx(CardTitle, { children: "Jobs by status" }) }),
      /* @__PURE__ */ jsx(CardContent, { children: /* @__PURE__ */ jsx(ChartContainer, { config: chartConfig, className: "h-72 w-full", children: /* @__PURE__ */ jsxs(BarChart, { data: chartData, children: [
        /* @__PURE__ */ jsx(CartesianGrid, { strokeDasharray: "3 3", vertical: false, className: "stroke-border/50" }),
        /* @__PURE__ */ jsx(XAxis, { dataKey: "status", tickLine: false, axisLine: false }),
        /* @__PURE__ */ jsx(YAxis, { tickLine: false, axisLine: false, allowDecimals: false }),
        /* @__PURE__ */ jsx(ChartTooltip, { content: /* @__PURE__ */ jsx(ChartTooltipContent, {}) }),
        /* @__PURE__ */ jsx(Bar, { dataKey: "count", fill: "var(--color-count)", radius: [6, 6, 0, 0] })
      ] }) }) })
    ] })
  ] });
}
export {
  DashboardPage as component
};
