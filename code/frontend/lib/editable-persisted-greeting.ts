export type GreetingResponse = {
  greeting: string;
};

function apiBase() {
  if (typeof window === "undefined") {
    return process.env.API_ORIGIN ?? "http://backend:8080";
  }
  return process.env.NEXT_PUBLIC_API_URL ?? "/api";
}

async function requestGreeting(init?: RequestInit): Promise<GreetingResponse> {
  const response = await fetch(`${apiBase()}/v1/greeting`, init);
  if (!response.ok) {
    throw new Error("Greeting API request failed");
  }
  return response.json() as Promise<GreetingResponse>;
}

export function readGreeting(): Promise<GreetingResponse> {
  return requestGreeting({ cache: "no-store" });
}

export function saveGreeting(greeting: string): Promise<GreetingResponse> {
  return requestGreeting({
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ greeting }),
  });
}
