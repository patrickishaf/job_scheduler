import { Link, useRouterState } from "@tanstack/react-router";
import { LayoutDashboard, ListChecks, PlusCircle, AlertOctagon, Cpu } from "lucide-react";
import { cn } from "@/lib/utils";

const items = [
  { to: "/", label: "Dashboard", icon: LayoutDashboard, exact: true },
  { to: "/jobs", label: "Jobs", icon: ListChecks, exact: false },
  { to: "/jobs/new", label: "Create Job", icon: PlusCircle, exact: true },
  { to: "/dlq", label: "Dead Letter Queue", icon: AlertOctagon, exact: true },
] as const;

export function AppSidebar() {
  const pathname = useRouterState({ select: (s) => s.location.pathname });
  return (
    <aside className="hidden md:flex w-60 shrink-0 flex-col border-r border-border bg-card/40 backdrop-blur">
      <div className="flex items-center gap-2 px-5 py-5 border-b border-border">
        <div className="flex h-8 w-8 items-center justify-center rounded-md bg-primary/20 text-primary">
          <Cpu className="h-4 w-4" />
        </div>
        <div>
          <div className="text-sm font-semibold tracking-tight">Jobrunner</div>
          <div className="text-xs text-muted-foreground">Background workers</div>
        </div>
      </div>
      <nav className="flex flex-col gap-1 p-3">
        {items.map((it) => {
          const active = it.exact
            ? pathname === it.to
            : pathname === it.to || pathname.startsWith(it.to + "/");
          // Special case: /jobs/new shouldn't also activate /jobs
          const isActive = it.to === "/jobs" ? pathname === "/jobs" : active;
          return (
            <Link
              key={it.to}
              to={it.to}
              className={cn(
                "flex items-center gap-3 rounded-md px-3 py-2 text-sm transition-colors",
                isActive
                  ? "bg-primary/10 text-primary"
                  : "text-muted-foreground hover:bg-muted/50 hover:text-foreground",
              )}
            >
              <it.icon className="h-4 w-4" />
              {it.label}
            </Link>
          );
        })}
      </nav>
      <div className="mt-auto p-4 text-xs text-muted-foreground">
        Mock data · resets on server restart
      </div>
    </aside>
  );
}
