import { jsx } from "react/jsx-runtime";
import * as React from "react";
import { c as cn } from "./utils-H80jjgLf.js";
import axios from "axios";
const Card = React.forwardRef(
  ({ className, ...props }, ref) => /* @__PURE__ */ jsx(
    "div",
    {
      ref,
      className: cn("rounded-xl border bg-card text-card-foreground shadow", className),
      ...props
    }
  )
);
Card.displayName = "Card";
const CardHeader = React.forwardRef(
  ({ className, ...props }, ref) => /* @__PURE__ */ jsx("div", { ref, className: cn("flex flex-col space-y-1.5 p-6", className), ...props })
);
CardHeader.displayName = "CardHeader";
const CardTitle = React.forwardRef(
  ({ className, ...props }, ref) => /* @__PURE__ */ jsx(
    "div",
    {
      ref,
      className: cn("font-semibold leading-none tracking-tight", className),
      ...props
    }
  )
);
CardTitle.displayName = "CardTitle";
const CardDescription = React.forwardRef(
  ({ className, ...props }, ref) => /* @__PURE__ */ jsx("div", { ref, className: cn("text-sm text-muted-foreground", className), ...props })
);
CardDescription.displayName = "CardDescription";
const CardContent = React.forwardRef(
  ({ className, ...props }, ref) => /* @__PURE__ */ jsx("div", { ref, className: cn("p-6 pt-0", className), ...props })
);
CardContent.displayName = "CardContent";
const CardFooter = React.forwardRef(
  ({ className, ...props }, ref) => /* @__PURE__ */ jsx("div", { ref, className: cn("flex items-center p-6 pt-0", className), ...props })
);
CardFooter.displayName = "CardFooter";
const API_BASE_URL = "http://localhost:8000";
const client = axios.create({
  baseURL: API_BASE_URL
});
const networkService = {
  get: async (url, query) => {
    const res = await client.get(url, query && query);
    return res.data;
  },
  patch: async (url, body = {}) => {
    const res = await client.patch(url, body);
    return res.data;
  },
  post: async (url, body) => {
    const res = await client.post(url, body);
    return res.data;
  }
};
export {
  Card as C,
  CardHeader as a,
  CardTitle as b,
  CardContent as c,
  networkService as n
};
