import { jsx, jsxs } from "react/jsx-runtime";
import { useQueryClient, useMutation } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { B as Button } from "./button-BC9oXVxV.js";
import * as React from "react";
import { c as cn } from "./utils-H80jjgLf.js";
import * as LabelPrimitive from "@radix-ui/react-label";
import { cva } from "class-variance-authority";
import * as SelectPrimitive from "@radix-ui/react-select";
import { ChevronDown, Check, ChevronUp } from "lucide-react";
import { toast } from "sonner";
import { n as networkService } from "./network.service-Vg4jAWs3.js";
const jobStatuses = [
  "queued",
  "processing",
  "completed",
  "failed",
  "cancelled",
  "pending"
];
const jobTypes = [
  "email.send",
  "report.generate",
  "cleanup.tmp",
  "image.resize",
  "webhook.deliver",
  "data.export"
];
const createJobSchema = z.object({
  type: z.string().min(1).max(100),
  priority: z.number().int().min(1).max(10),
  scheduled_time: z.string().min(1),
  interval: z.number().int().min(0).max(86400 * 30).nullable(),
  maxRetries: z.number().int().min(0).max(20),
  payload: z.string()
  // JSON string, parsed server-side
});
const Input = React.forwardRef(
  ({ className, type, ...props }, ref) => {
    return /* @__PURE__ */ jsx(
      "input",
      {
        type,
        className: cn(
          "flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
          className
        ),
        ref,
        ...props
      }
    );
  }
);
Input.displayName = "Input";
const Select = SelectPrimitive.Root;
const SelectValue = SelectPrimitive.Value;
const SelectTrigger = React.forwardRef(({ className, children, ...props }, ref) => /* @__PURE__ */ jsxs(
  SelectPrimitive.Trigger,
  {
    ref,
    className: cn(
      "flex h-9 w-full items-center justify-between whitespace-nowrap rounded-md border border-input bg-transparent px-3 py-2 text-sm shadow-sm ring-offset-background cursor-pointer data-[placeholder]:text-muted-foreground focus:outline-none focus:ring-1 focus:ring-ring disabled:cursor-not-allowed disabled:opacity-50 [&>span]:line-clamp-1",
      className
    ),
    ...props,
    children: [
      children,
      /* @__PURE__ */ jsx(SelectPrimitive.Icon, { asChild: true, children: /* @__PURE__ */ jsx(ChevronDown, { className: "h-4 w-4 opacity-50" }) })
    ]
  }
));
SelectTrigger.displayName = SelectPrimitive.Trigger.displayName;
const SelectScrollUpButton = React.forwardRef(({ className, ...props }, ref) => /* @__PURE__ */ jsx(
  SelectPrimitive.ScrollUpButton,
  {
    ref,
    className: cn("flex cursor-default items-center justify-center py-1", className),
    ...props,
    children: /* @__PURE__ */ jsx(ChevronUp, { className: "h-4 w-4" })
  }
));
SelectScrollUpButton.displayName = SelectPrimitive.ScrollUpButton.displayName;
const SelectScrollDownButton = React.forwardRef(({ className, ...props }, ref) => /* @__PURE__ */ jsx(
  SelectPrimitive.ScrollDownButton,
  {
    ref,
    className: cn("flex cursor-default items-center justify-center py-1", className),
    ...props,
    children: /* @__PURE__ */ jsx(ChevronDown, { className: "h-4 w-4" })
  }
));
SelectScrollDownButton.displayName = SelectPrimitive.ScrollDownButton.displayName;
const SelectContent = React.forwardRef(({ className, children, position = "popper", ...props }, ref) => /* @__PURE__ */ jsx(SelectPrimitive.Portal, { children: /* @__PURE__ */ jsxs(
  SelectPrimitive.Content,
  {
    ref,
    className: cn(
      "relative z-50 max-h-(--radix-select-content-available-height) min-w-[8rem] overflow-y-auto overflow-x-hidden rounded-md border bg-popover text-popover-foreground shadow-md data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 origin-(--radix-select-content-transform-origin)",
      position === "popper" && "data-[side=bottom]:translate-y-1 data-[side=left]:-translate-x-1 data-[side=right]:translate-x-1 data-[side=top]:-translate-y-1",
      className
    ),
    position,
    ...props,
    children: [
      /* @__PURE__ */ jsx(SelectScrollUpButton, {}),
      /* @__PURE__ */ jsx(
        SelectPrimitive.Viewport,
        {
          className: cn(
            "p-1",
            position === "popper" && "h-[var(--radix-select-trigger-height)] w-full min-w-[var(--radix-select-trigger-width)]"
          ),
          children
        }
      ),
      /* @__PURE__ */ jsx(SelectScrollDownButton, {})
    ]
  }
) }));
SelectContent.displayName = SelectPrimitive.Content.displayName;
const SelectLabel = React.forwardRef(({ className, ...props }, ref) => /* @__PURE__ */ jsx(
  SelectPrimitive.Label,
  {
    ref,
    className: cn("px-2 py-1.5 text-sm font-semibold", className),
    ...props
  }
));
SelectLabel.displayName = SelectPrimitive.Label.displayName;
const SelectItem = React.forwardRef(({ className, children, ...props }, ref) => /* @__PURE__ */ jsxs(
  SelectPrimitive.Item,
  {
    ref,
    className: cn(
      "relative flex w-full cursor-default select-none items-center rounded-sm py-1.5 pl-2 pr-8 text-sm outline-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
      className
    ),
    ...props,
    children: [
      /* @__PURE__ */ jsx("span", { className: "absolute right-2 flex h-3.5 w-3.5 items-center justify-center", children: /* @__PURE__ */ jsx(SelectPrimitive.ItemIndicator, { children: /* @__PURE__ */ jsx(Check, { className: "h-4 w-4" }) }) }),
      /* @__PURE__ */ jsx(SelectPrimitive.ItemText, { children })
    ]
  }
));
SelectItem.displayName = SelectPrimitive.Item.displayName;
const SelectSeparator = React.forwardRef(({ className, ...props }, ref) => /* @__PURE__ */ jsx(
  SelectPrimitive.Separator,
  {
    ref,
    className: cn("-mx-1 my-1 h-px bg-muted", className),
    ...props
  }
));
SelectSeparator.displayName = SelectPrimitive.Separator.displayName;
const labelVariants = cva(
  "text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
);
const Label = React.forwardRef(({ className, ...props }, ref) => /* @__PURE__ */ jsx(LabelPrimitive.Root, { ref, className: cn(labelVariants(), className), ...props }));
Label.displayName = LabelPrimitive.Root.displayName;
const Textarea = React.forwardRef(
  ({ className, ...props }, ref) => {
    return /* @__PURE__ */ jsx(
      "textarea",
      {
        className: cn(
          "flex min-h-[60px] w-full rounded-md border border-input bg-transparent px-3 py-2 text-base shadow-sm placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
          className
        ),
        ref,
        ...props
      }
    );
  }
);
Textarea.displayName = "Textarea";
const CREATE_JOB_API_URL = "/api/jobs";
function defaultScheduledAt() {
  const d = /* @__PURE__ */ new Date();
  d.setMinutes(d.getMinutes() - d.getTimezoneOffset());
  return d.toISOString().slice(0, 16);
}
function CreateJobForm({ onCreated, onCancel }) {
  const qc = useQueryClient();
  const form = useForm({
    resolver: zodResolver(createJobSchema),
    defaultValues: {
      type: jobTypes[0],
      priority: 2,
      scheduled_time: defaultScheduledAt(),
      interval: null,
      maxRetries: 3,
      payload: '{\n  "example": true\n}'
    }
  });
  const mutation = useMutation({
    mutationFn: async (data) => {
      let parsedPayload;
      try {
        parsedPayload = JSON.parse(data.payload);
      } catch {
        throw new Error("Payload is not valid JSON");
      }
      if (!parsedPayload || typeof parsedPayload !== "object" || Array.isArray(parsedPayload)) {
        throw new Error("Payload must be a JSON object");
      }
      const body = {
        type: data.type,
        priority: data.priority,
        payload: parsedPayload,
        scheduled_time: new Date(data.scheduled_time).toISOString()
      };
      if (data.interval && data.interval > 0) {
        body.interval_seconds = data.interval;
      }
      const res = await networkService.post(CREATE_JOB_API_URL, body);
      if (res.status === "error") {
        throw new Error(res.message);
      }
      return res.data;
    },
    onSuccess: (job) => {
      toast.success("Job created", { description: job.id });
      qc.invalidateQueries({ queryKey: ["jobs"] });
      qc.invalidateQueries({ queryKey: ["job-summary"] });
      form.reset();
      onCreated?.(job);
    },
    onError: (err) => toast.error("Failed to create job", { description: err.message })
  });
  return /* @__PURE__ */ jsxs("form", { onSubmit: form.handleSubmit((d) => mutation.mutate(d)), className: "space-y-5", children: [
    /* @__PURE__ */ jsxs("div", { className: "space-y-2", children: [
      /* @__PURE__ */ jsx(Label, { htmlFor: "type", children: "Type" }),
      /* @__PURE__ */ jsxs(Select, { value: form.watch("type"), onValueChange: (v) => form.setValue("type", v), children: [
        /* @__PURE__ */ jsx(SelectTrigger, { id: "type", children: /* @__PURE__ */ jsx(SelectValue, {}) }),
        /* @__PURE__ */ jsx(SelectContent, { children: jobTypes.map((t) => /* @__PURE__ */ jsx(SelectItem, { value: t, children: t }, t)) })
      ] }),
      /* @__PURE__ */ jsx("p", { className: "text-xs text-muted-foreground", children: "The worker handler registered for this job type." })
    ] }),
    /* @__PURE__ */ jsxs("div", { className: "grid grid-cols-2 gap-4", children: [
      /* @__PURE__ */ jsxs("div", { className: "space-y-2", children: [
        /* @__PURE__ */ jsx(Label, { htmlFor: "priority", children: "Priority (1–3)" }),
        /* @__PURE__ */ jsx(
          Input,
          {
            id: "priority",
            type: "number",
            min: 1,
            max: 3,
            ...form.register("priority", { valueAsNumber: true })
          }
        ),
        /* @__PURE__ */ jsx("p", { className: "text-xs text-muted-foreground", children: "1 = lowest, 3 = highest." })
      ] }),
      /* @__PURE__ */ jsxs("div", { className: "space-y-2", children: [
        /* @__PURE__ */ jsx(Label, { htmlFor: "maxRetries", children: "Max retries" }),
        /* @__PURE__ */ jsx(
          Input,
          {
            id: "maxRetries",
            type: "number",
            min: 0,
            max: 20,
            ...form.register("maxRetries", { valueAsNumber: true })
          }
        )
      ] })
    ] }),
    /* @__PURE__ */ jsxs("div", { className: "grid grid-cols-2 gap-4", children: [
      /* @__PURE__ */ jsxs("div", { className: "space-y-2", children: [
        /* @__PURE__ */ jsx(Label, { htmlFor: "scheduledAt", children: "Scheduled at" }),
        /* @__PURE__ */ jsx(Input, { id: "scheduledAt", type: "datetime-local", ...form.register("scheduled_time") }),
        /* @__PURE__ */ jsx("p", { className: "text-xs text-muted-foreground", children: "Future = scheduled, now/past = queued." })
      ] }),
      /* @__PURE__ */ jsxs("div", { className: "space-y-2", children: [
        /* @__PURE__ */ jsx(Label, { htmlFor: "interval", children: "Interval (seconds)" }),
        /* @__PURE__ */ jsx(
          Input,
          {
            id: "interval",
            type: "number",
            min: 0,
            placeholder: "0 = one-off",
            onChange: (e) => {
              const n = e.target.valueAsNumber;
              form.setValue("interval", Number.isFinite(n) && n > 0 ? n : null);
            }
          }
        ),
        /* @__PURE__ */ jsx("p", { className: "text-xs text-muted-foreground", children: "Re-run every N seconds." })
      ] })
    ] }),
    /* @__PURE__ */ jsxs("div", { className: "space-y-2", children: [
      /* @__PURE__ */ jsx(Label, { htmlFor: "payload", children: "Payload (JSON)" }),
      /* @__PURE__ */ jsx(
        Textarea,
        {
          id: "payload",
          rows: 6,
          className: "font-mono text-sm",
          ...form.register("payload")
        }
      )
    ] }),
    /* @__PURE__ */ jsxs("div", { className: "flex justify-end gap-2 pt-2", children: [
      onCancel && /* @__PURE__ */ jsx(Button, { type: "button", variant: "outline", onClick: onCancel, children: "Cancel" }),
      /* @__PURE__ */ jsx(Button, { type: "submit", disabled: mutation.isPending, children: mutation.isPending ? "Creating…" : "Create job" })
    ] })
  ] });
}
export {
  CreateJobForm as C,
  Input as I,
  Select as S,
  SelectTrigger as a,
  SelectValue as b,
  SelectContent as c,
  SelectItem as d,
  jobStatuses as j
};
