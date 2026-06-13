import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { ArrowLeft } from "lucide-react";
import { CreateJobForm } from "@/components/jobs/CreateJobForm";

export const Route = createFileRoute("/_app/jobs/new")({
  head: () => ({ meta: [{ title: "Create Job · Jobrunner" }] }),
  component: NewJobPage,
});

function NewJobPage() {
  const navigate = useNavigate();
  return (
    <div className="p-8 max-w-2xl">
      <Button variant="ghost" size="sm" onClick={() => navigate({ to: "/jobs" })} className="mb-4">
        <ArrowLeft className="h-4 w-4 mr-2" /> Back to jobs
      </Button>

      <Card className="border-border/60">
        <CardHeader>
          <CardTitle>Create new job</CardTitle>
        </CardHeader>
        <CardContent>
          <CreateJobForm
            onCreated={() => navigate({ to: "/jobs" })}
            onCancel={() => navigate({ to: "/jobs" })}
          />
        </CardContent>
      </Card>
    </div>
  );
}
