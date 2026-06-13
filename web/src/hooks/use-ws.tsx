import { WsContext } from "@/lib/ws";
import { useContext } from "react";

export function useWs() {
  const ctx = useContext(WsContext);
  if (!ctx) throw new Error("useWs must be used within a WsProvider");
  return ctx;
}

// Convenience hook: subscribe and react to messages in one call
