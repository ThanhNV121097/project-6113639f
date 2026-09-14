export type GreetingResponse = {
  greeting: string;
};

const STORAGE_KEY = "hello-world-acceptance-3:greeting";
const DEFAULT_GREETING = "Hello, World!";

export function readGreeting(): GreetingResponse {
  if (typeof window === "undefined") {
    return { greeting: DEFAULT_GREETING };
  }

  return { greeting: window.localStorage.getItem(STORAGE_KEY) || DEFAULT_GREETING };
}

export function saveGreeting(greeting: string): GreetingResponse {
  const nextGreeting = greeting.trim();

  if (!nextGreeting) {
    return readGreeting();
  }

  window.localStorage.setItem(STORAGE_KEY, nextGreeting);
  return { greeting: nextGreeting };
}
