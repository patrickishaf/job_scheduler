import { Outlet, createFileRoute } from "@tanstack/react-router";
import { AppSidebar } from "@/components/jobs/Sidebar";
import { Toaster } from "@/components/ui/sonner";

export const Route = createFileRoute("/_app")({
  component: AppLayout,
});

function AppLayout() {
  return (
    <div className="min-h-screen bg-background text-foreground">
      <div className="flex min-h-screen w-full">
        <AppSidebar />
        <main className="flex-1 min-w-0">
          <Outlet />
        </main>
      </div>
      <Toaster theme="dark" richColors position="top-right" />
    </div>
  );
}
