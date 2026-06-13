import { createContext, useEffect, useRef, useCallback, useState, type ReactNode } from "react";

export type MessageListener = (event: MessageEvent) => void;

interface WsContextValue {
  send: (data: string | object) => void;
  subscribe: (listener: MessageListener) => () => void;
  status: "connecting" | "open" | "closed" | "error";
}

export const WsContext = createContext<WsContextValue | null>(null);

export function WsProvider({ url, children }: { url: string; children: ReactNode }) {
  const wsRef = useRef<WebSocket | null>(null);
  const listenersRef = useRef<Set<MessageListener>>(new Set());
  const [status, setStatus] = useState<WsContextValue["status"]>("connecting");

  useEffect(() => {
    const ws = new WebSocket(url);
    wsRef.current = ws;

    ws.onopen = () => setStatus("open");
    ws.onerror = () => setStatus("error");
    ws.onclose = () => setStatus("closed");
    ws.onmessage = (event) => {
      listenersRef.current.forEach((fn) => fn(event));
    };

    return () => ws.close();
  }, [url]);

  const send = useCallback((data: string | object) => {
    const ws = wsRef.current;
    if (ws?.readyState === WebSocket.OPEN) {
      ws.send(typeof data === "string" ? data : JSON.stringify(data));
    }
  }, []);

  const subscribe = useCallback((listener: MessageListener) => {
    listenersRef.current.add(listener);
    return () => listenersRef.current.delete(listener);
  }, []);

  return <WsContext.Provider value={{ send, subscribe, status }}>{children}</WsContext.Provider>;
}
