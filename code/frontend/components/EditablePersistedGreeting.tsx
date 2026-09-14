"use client";

import { FormEvent, useRef, useState } from "react";
import { readGreeting, saveGreeting } from "../lib/mock/editable-persisted-greeting";
import styles from "./EditablePersistedGreeting.module.css";

export function EditablePersistedGreeting() {
  const [greeting, setGreeting] = useState(() => readGreeting().greeting);
  const [draft, setDraft] = useState(greeting);
  const inputRef = useRef<HTMLInputElement>(null);

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const nextGreeting = draft.trim();

    if (!nextGreeting) {
      inputRef.current?.focus();
      return;
    }

    const response = saveGreeting(nextGreeting);
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
