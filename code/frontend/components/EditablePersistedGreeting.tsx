"use client";

import { FormEvent, useRef, useState } from "react";
import { saveGreeting } from "../lib/editable-persisted-greeting";
import styles from "./EditablePersistedGreeting.module.css";

const DEFAULT_GREETING = "Hello, World!";

type EditablePersistedGreetingProps = {
  initialGreeting?: string;
};

export function EditablePersistedGreeting({
  initialGreeting = DEFAULT_GREETING,
}: EditablePersistedGreetingProps) {
  const [greeting, setGreeting] = useState(initialGreeting);
  const [draft, setDraft] = useState(initialGreeting);
  const inputRef = useRef<HTMLInputElement>(null);

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const nextGreeting = draft.trim();

    if (!nextGreeting) {
      inputRef.current?.focus();
      return;
    }

    const response = await saveGreeting(nextGreeting);
    setGreeting(response.greeting);
    setDraft(response.greeting);
  }

  return (
    <main className={styles.shell}>
      <section className={styles.section} aria-labelledby="greeting-heading">
        <h1 id="greeting-heading" className={styles.heading}>
          {greeting}
        </h1>
        <form className={styles.form} onSubmit={handleSubmit}>
          <label className={styles.label} htmlFor="greeting-input">
            Greeting
          </label>
          <input
            ref={inputRef}
            id="greeting-input"
            name="greeting"
            type="text"
            value={draft}
            autoComplete="off"
            required
            className={styles.input}
            onChange={(event) => setDraft(event.target.value)}
          />
          <button className={styles.button} type="submit">
            Save
          </button>
        </form>
      </section>
    </main>
  );
}
