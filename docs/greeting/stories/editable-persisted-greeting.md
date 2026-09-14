# Story — Editable persisted greeting

Module: `greeting`
Plan item: Editable persisted greeting

## User story

As a Visitor, I want to view and change the shared persisted greeting, so that the page and future reloads show the current stored greeting.

## In scope

- Show the shared current greeting as one large centered heading on the only app page.
- Seed or default the stored greeting to `Hello, World!` when no saved greeting exists.
- Show the same current greeting in the text input on page load.
- Let any visitor enter a non-blank greeting and submit it with the `Save` button or native form submission.
- Trim saved greeting input before persistence and display.
- Persist the saved greeting in PostgreSQL through the Go API so reloads show the saved value.
- Keep the approved one-section minimal design: white background, black text, blue `#2563EB` Save button, no animation.
- Support keyboard access, programmatic `Greeting` label, visible focus, and responsive stacking at 520px and below.

## Out of scope

- Sign-in, identity, roles, or per-visitor greetings — SRS defines one shared greeting and no sign-in.
- Navigation, secondary pages, or extra page sections — approved design has one centered greeting section only.
- External services — stakeholder specified none.
- Visible loading, empty, or error screens — approved design and SRS do not include those states.
- Localization — default content is English and visitor-entered text displays as entered after trimming.
- Greeting history, audit log, undo, conflict UI, or revision tracking — latest completed save wins.
- Custom active or disabled button states — approved design only requires default, hover, and focus-visible states.

## UI scope

This story touches the single Greeting page only: `main > section.greeting-section` with one `h1#greeting-heading`, a programmatically labelled `input#greeting-input` labelled `Greeting`, and a `Save` submit button.

Required screen state: `default` only. The UI must match the approved design system tokens: centered section, white `#FFFFFF` background, black `#000000` text, large bold centered heading, 48px input and button, blue `#2563EB` button, darker blue `#1D4ED8` hover, 3px blue focus outline with 3px offset, no animation or transition, stacked full-width controls at viewport widths of 520px and below.

## Acceptance criteria

- SC-1 [GREETING-001 AC-1]: Given stored greeting is `Hello, World!`, when Visitor opens the app page, then page shows `Hello, World!` as one large centered heading.
- SC-2 [GREETING-001 AC-2]: Given stored greeting is `Good morning`, when Visitor opens the app page, then page shows `Good morning` as the heading.
- SC-3 [GREETING-001 AC-3]: Given stored greeting is `Good morning`, when Visitor opens the app page, then text input value is `Good morning`.
- SC-4 [GREETING-002 AC-1]: Given page shows `Hello, World!`, when Visitor types `Pipeline accepted` and clicks `Save`, then heading changes to `Pipeline accepted`.
- SC-5 [GREETING-002 AC-2]: Given Visitor saved `Pipeline accepted`, when Visitor reloads the page, then heading shows `Pipeline accepted`.
- SC-6 [GREETING-002 AC-3]: Given page shows `Hello, World!`, when Visitor types only spaces and submits, then heading remains `Hello, World!`.
- SC-7 [GREETING-002 AC-4]: Given page shows `Hello, World!`, when Visitor types `  Hello again  ` and submits, then heading shows `Hello again`.
- SC-8 [GREETING-003 AC-1]: Given Visitor opens the page, when visual inspection checks page structure, then page contains one centered greeting section with no navigation and no other sections.
- SC-9 [GREETING-003 AC-2]: Given Visitor opens the page, when visual inspection checks colors, then background is `#FFFFFF`, text is `#000000`, and Save button default fill is `#2563EB`.
- SC-10 [GREETING-003 AC-3]: Given viewport width is 390px, when Visitor opens the page, then text input and Save button are both 48px tall.
- SC-11 [GREETING-003 AC-4]: Given Visitor opens the page, when visual inspection checks motion, then no animation or transition occurs.
- SC-12 [GREETING-003 AC-5]: Given viewport width is 520px or below, when Visitor opens the page, then text input and Save button stack vertically and each uses full available width.

## Dependencies

- PostgreSQL stores one shared greeting value and preserves it across reloads.
- Go API reads and updates the shared greeting using `DATABASE_URL` and parameterized SQL.
- Next.js frontend renders the approved Greeting page and calls the API using project environment conventions.
- Architecture scaffold exists for `code/backend`, `code/frontend`, Docker Compose, and CI.
- No external accounts, credentials, or provider setup required.
