import { jsxs, jsx } from "react/jsx-runtime";
import { useNavigate } from "@tanstack/react-router";
import { C as Card, a as CardHeader, b as CardTitle, c as CardContent } from "./network.service-Vg4jAWs3.js";
import { B as Button } from "./button-BC9oXVxV.js";
import { ArrowLeft } from "lucide-react";
import { C as CreateJobForm } from "./CreateJobForm-JskyILMC.js";
import "react";
import "./utils-H80jjgLf.js";
import "clsx";
import "tailwind-merge";
import "axios";
import "@radix-ui/react-slot";
import "class-variance-authority";
import "@tanstack/react-query";
import "react-hook-form";
import "@hookform/resolvers/zod";
import "zod";
import "@radix-ui/react-label";
import "@radix-ui/react-select";
import "sonner";
function NewJobPage() {
  const navigate = useNavigate();
  return /* @__PURE__ */ jsxs("div", { className: "p-8 max-w-2xl", children: [
    /* @__PURE__ */ jsxs(Button, { variant: "ghost", size: "sm", onClick: () => navigate({
      to: "/jobs"
    }), className: "mb-4", children: [
      /* @__PURE__ */ jsx(ArrowLeft, { className: "h-4 w-4 mr-2" }),
      " Back to jobs"
    ] }),
    /* @__PURE__ */ jsxs(Card, { className: "border-border/60", children: [
      /* @__PURE__ */ jsx(CardHeader, { children: /* @__PURE__ */ jsx(CardTitle, { children: "Create new job" }) }),
      /* @__PURE__ */ jsx(CardContent, { children: /* @__PURE__ */ jsx(CreateJobForm, { onCreated: () => navigate({
        to: "/jobs"
      }), onCancel: () => navigate({
        to: "/jobs"
      }) }) })
    ] })
  ] });
}
export {
  NewJobPage as component
};
