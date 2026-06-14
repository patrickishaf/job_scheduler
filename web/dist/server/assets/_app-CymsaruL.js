import { jsxs, jsx } from "react/jsx-runtime";
import { useRouterState, Link, Outlet } from "@tanstack/react-router";
import { Cpu, LayoutDashboard, ListChecks, PlusCircle, AlertOctagon } from "lucide-react";
import { c as cn } from "./utils-H80jjgLf.js";
import { Toaster as Toaster$1 } from "sonner";
import "clsx";
import "tailwind-merge";
const items = [
  { to: "/", label: "Dashboard", icon: LayoutDashboard, exact: true },
  { to: "/jobs", label: "Jobs", icon: ListChecks, exact: false },
  { to: "/jobs/new", label: "Create Job", icon: PlusCircle, exact: true },
  { to: "/dlq", label: "Dead Letter Queue", icon: AlertOctagon, exact: true }
];
function AppSidebar() {
  const pathname = useRouterState({ select: (s) => s.location.pathname });
  return /* @__PURE__ */ jsxs("aside", { className: "hidden md:flex w-60 shrink-0 flex-col border-r border-border bg-card/40 backdrop-blur", children: [
    /* @__PURE__ */ jsxs("div", { className: "flex items-center gap-2 px-5 py-5 border-b border-border", children: [
      /* @__PURE__ */ jsx("div", { className: "flex h-8 w-8 items-center justify-center rounded-md bg-primary/20 text-primary", children: /* @__PURE__ */ jsx(Cpu, { className: "h-4 w-4" }) }),
      /* @__PURE__ */ jsxs("div", { children: [
        /* @__PURE__ */ jsx("div", { className: "text-sm font-semibold tracking-tight", children: "Jobrunner" }),
        /* @__PURE__ */ jsx("div", { className: "text-xs text-muted-foreground", children: "Background workers" })
      ] })
    ] }),
    /* @__PURE__ */ jsx("nav", { className: "flex flex-col gap-1 p-3", children: items.map((it) => {
      const active = it.exact ? pathname === it.to : pathname === it.to || pathname.startsWith(it.to + "/");
      const isActive = it.to === "/jobs" ? pathname === "/jobs" : active;
      return /* @__PURE__ */ jsxs(
        Link,
        {
          to: it.to,
          className: cn(
            "flex items-center gap-3 rounded-md px-3 py-2 text-sm transition-colors",
            isActive ? "bg-primary/10 text-primary" : "text-muted-foreground hover:bg-muted/50 hover:text-foreground"
          ),
          children: [
            /* @__PURE__ */ jsx(it.icon, { className: "h-4 w-4" }),
            it.label
          ]
        },
        it.to
      );
    }) }),
    /* @__PURE__ */ jsx("div", { className: "mt-auto p-4 text-xs text-muted-foreground", children: "Mock data · resets on server restart" })
  ] });
}
const Toaster = ({ ...props }) => {
  return /* @__PURE__ */ jsx(
    Toaster$1,
    {
      className: "toaster group",
      toastOptions: {
        classNames: {
          toast: "group toast group-[.toaster]:bg-background group-[.toaster]:text-foreground group-[.toaster]:border-border group-[.toaster]:shadow-lg",
          description: "group-[.toast]:text-muted-foreground",
          actionButton: "group-[.toast]:bg-primary group-[.toast]:text-primary-foreground",
          cancelButton: "group-[.toast]:bg-muted group-[.toast]:text-muted-foreground"
        }
      },
      ...props
    }
  );
};
function AppLayout() {
  return /* @__PURE__ */ jsxs("div", { className: "min-h-screen bg-background text-foreground", children: [
    /* @__PURE__ */ jsxs("div", { className: "flex min-h-screen w-full", children: [
      /* @__PURE__ */ jsx(AppSidebar, {}),
      /* @__PURE__ */ jsx("main", { className: "flex-1 min-w-0", children: /* @__PURE__ */ jsx(Outlet, {}) })
    ] }),
    /* @__PURE__ */ jsx(Toaster, { theme: "dark", richColors: true, position: "top-right" })
  ] });
}
export {
  AppLayout as component
};
