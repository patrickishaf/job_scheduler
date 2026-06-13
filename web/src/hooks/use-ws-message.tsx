import { MessageListener } from "@/lib/ws";
import { useWs } from "./use-ws";
import { useEffect } from "react";

export function useWsMessage(listener: MessageListener) {
  const { subscribe } = useWs();
  useEffect(() => subscribe(listener), [subscribe, listener]);
}
