# SRS — Greeting

Module: `greeting`
Design: [View the approved design](http://localhost:8080/design/6113639f-e15d-4545-9f5b-e5ca40d9b57c)
Design system: `design/design-system.md`

> One file per module, at `docs/{module}/SRS.md`. It covers only the functions
> that belong to this module. Never write `docs/SRS.md`.

## 1. Purpose

The greeting module lets any visitor view and change one persisted greeting for "Hello World Acceptance 3". Without this module, the product cannot prove the required end-to-end path from PostgreSQL storage through Go API to Next.js UI.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Visitor | Any person opening the page; no sign-in exists | View the current greeting, edit the greeting text, and save it |

## 3. Scope

**In scope** — the functions specified below, by their plan titles:

- Editable persisted greeting

**Out of scope** — name what a reader would reasonably expect here and say where it lives instead. This section prevents the same argument twice.

- Sign-in and role-based permissions — deliberately not built; stakeholder specified no sign-in.
- Navigation, extra sections, and secondary pages — deliberately not built; approved design contains one centered greeting section only.
- External services — deliberately not built; stakeholder specified none.

## 4. Functional requirements

### 4.1 Editable persisted greeting

**Requirement GREETING-001 — Show current greeting**

*As a* Visitor, *I want to* see the current stored greeting as the main heading, *so that* I know the persisted value.

Behaviour:

1. Visitor opens the app page.
2. The page displays the current stored greeting in one large centered `h1`.
3. When no greeting has been changed before, the displayed greeting is `Hello, World!`.
4. The text input initially contains the same greeting value shown in the heading.

**Acceptance criteria** — each is proved by at least one test case in `docs/greeting/test-cases/editable-persisted-greeting.md`, through the story plan that cites it (`SC-1 [GREETING-001 AC-1]`). Given/When/Then, no compound conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Stored greeting is `Hello, World!` | Visitor opens the app page | Page shows `Hello, World!` as one large centered heading |
| AC-2 | Stored greeting is `Good morning` | Visitor opens the app page | Page shows `Good morning` as the heading |
| AC-3 | Stored greeting is `Good morning` | Visitor opens the app page | Text input value is `Good morning` |

**Failure, boundary and permission behaviour** — the part most often skipped and most often the source of bugs. Every case this function actually has needs a defined outcome; "should not happen" is not an outcome.

| Case | Condition | Expected behaviour |
|---|---|---|
| Missing stored value | No greeting exists yet | Visitor sees `Hello, World!` as the initial greeting |
| Permission | Any visitor opens the page | Allowed; no sign-in or role check exists |
| Upstream failure | Stored greeting cannot be read | No error state is part of the approved design; API error handling belongs in TL service contract |
| Empty state | Greeting value absent from storage | Not applicable as a visible state; default greeting is shown |

**Data touched** — the fields this function reads and writes, in product terms. The physical schema is TL's job in `docs/architecture/erd.md`; this is the list that document has to satisfy.

| Field | Type | Required | Rule |
|---|---|---|---|
| Greeting text | text | yes | Initial value is `Hello, World!`; displayed exactly as stored after trimming save input |

**Requirement GREETING-002 — Save changed greeting**

*As a* Visitor, *I want to* enter a new greeting and save it, *so that* the page and future reloads show my new value.

Behaviour:

1. Visitor edits the accessible text input labelled `Greeting`.
2. Visitor submits the form with the `Save` button or native form submission from the input.
3. Blank input after trimming whitespace is not saved; focus returns to the input and the current heading remains unchanged.
4. Non-blank input is trimmed, saved, shown in the heading, and reflected back in the input value.
5. After reload, the page shows the saved greeting.

**Acceptance criteria** — each is proved by at least one test case in `docs/greeting/test-cases/editable-persisted-greeting.md`, through the story plan that cites it (`SC-1 [GREETING-002 AC-1]`). Given/When/Then, no compound conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Page shows `Hello, World!` | Visitor types `Pipeline accepted` and clicks `Save` | Heading changes to `Pipeline accepted` |
| AC-2 | Visitor saved `Pipeline accepted` | Visitor reloads the page | Heading shows `Pipeline accepted` |
| AC-3 | Page shows `Hello, World!` | Visitor types only spaces and submits | Heading remains `Hello, World!` |
| AC-4 | Page shows `Hello, World!` | Visitor types `  Hello again  ` and submits | Heading shows `Hello again` |

**Failure, boundary and permission behaviour**

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | Greeting input is blank after trimming whitespace | Nothing is saved, heading stays unchanged, input receives focus; no error message is part of the approved design |
| Boundary | Greeting text contains leading or trailing whitespace | Saved value is trimmed before display and persistence |
| Permission | Any visitor submits a non-blank greeting | Allowed; no sign-in or role check exists |
| Conflict | Two visitors save different non-blank greetings | Latest completed save is the value shown after reload |
| Upstream failure | Greeting cannot be saved | No error state is part of the approved design; API error handling belongs in TL service contract |

**Data touched**

| Field | Type | Required | Rule |
|---|---|---|---|
| Greeting text | text | yes | Non-blank after trimming; saved value persists after reload |

## 5. Screens

The design is the source of truth for appearance; this section maps functions onto it so nothing in the design is unaccounted for and nothing specified here is missing from the design.

List only the states the approved design actually shows. A screen the design draws once, with no variant for waiting, for no data, or for a failure, has exactly **one** state and its name is `default`. That is not a placeholder and not an invented state: it is what "this screen has one appearance" is called, and it is the correct and complete answer for a static screen. Writing `loading`, `empty` or `error` for a screen whose design has no such variant invents work, and the reviewer will reject it.

| Screen | Section in the design | Functions it serves | States that must exist |
|---|---|---|---|
| Greeting page | `main > section.greeting-section` with `h1#greeting-heading`, hidden label `Greeting`, `input#greeting-input`, and `button` text `Save` | GREETING-001, GREETING-002 | default |

Approved design elements covered by acceptance criteria:

- Large centered greeting heading — GREETING-001 AC-1, GREETING-001 AC-2, GREETING-002 AC-1.
- Text input labelled `Greeting` — GREETING-001 AC-3, GREETING-002 AC-3, GREETING-002 AC-4.
- `Save` button — GREETING-002 AC-1.
- Text input and Save button are both 48px tall at every screen width — GREETING-003 AC-3.
- Plain white page, black text, blue `#2563EB` button, no animation — GREETING-003 AC-1, GREETING-003 AC-2, GREETING-003 AC-4.

### 5.1 Visual and interaction requirement

**Requirement GREETING-003 — Match approved one-section design**

*As a* Visitor, *I want to* use the greeting page in the approved minimal layout, *so that* the experience matches the stakeholder-approved design.

Behaviour:

1. The app page contains one centered section and no navigation or extra sections.
2. The page background is white `#FFFFFF`, primary text is black `#000000`, and the Save button uses blue `#2563EB` in default state.
3. The greeting appears as one large, bold, centered heading.
4. The form appears below the heading with an accessible text input labelled `Greeting` and a `Save` submit button.
5. The text input and Save button are both 48px tall at every screen width, including 390px.
6. At widths of 520px and below, the form stacks vertically and both controls use full available width.
7. No animation or transition appears.

**Acceptance criteria** — each is proved by at least one test case in `docs/greeting/test-cases/editable-persisted-greeting.md`, through the story plan that cites it (`SC-1 [GREETING-003 AC-1]`).

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Visitor opens the page | Visual inspection checks page structure | Page contains one centered greeting section with no navigation and no other sections |
| AC-2 | Visitor opens the page | Visual inspection checks colors | Background is `#FFFFFF`, text is `#000000`, and Save button default fill is `#2563EB` |
| AC-3 | Viewport width is 390px | Visitor opens the page | Text input and Save button are both 48px tall |
| AC-4 | Visitor opens the page | Visual inspection checks motion | No animation or transition occurs |
| AC-5 | Viewport width is 520px or below | Visitor opens the page | Text input and Save button stack vertically and each uses full available width |

**Failure, boundary and permission behaviour**

| Case | Condition | Expected behaviour |
|---|---|---|
| Responsive boundary | Viewport width is 390px | Input remains 48px tall; no horizontal page scroll appears |
| Focus state | Input or Save button receives keyboard focus | 3px blue focus outline with 3px offset is visible |
| Hover state | Pointer hovers Save button | Save button changes to darker blue `#1D4ED8` |
| Loading, empty, or error state | API is waiting, empty, or failed | Not part of approved design; no visible loading, empty, or error screen is required by this SRS |

**Data touched**

| Field | Type | Required | Rule |
|---|---|---|---|
| Greeting text | text | yes | Rendered in heading and input using approved typography and control layout |

## 6. Non-functional requirements

Only what is real for this module. Delete rows that do not apply rather than inventing a number nobody will check.

| Area | Requirement |
|---|---|
| Accessibility | Text input has programmatic label `Greeting`; input and button are keyboard reachable; focus-visible outline appears on both controls; text contrast is at least 4.5:1 |
| Responsive | Page works from 320px viewport width upward with no horizontal page scroll; at 390px width the input and Save button are each 48px tall |
| Persistence | Saved non-blank greeting survives a browser reload and remains visible as heading and input value |
| Privacy | No personal data is required; only visitor-entered greeting text is stored |

## 7. Dependencies and assumptions

- **Depends on:** PostgreSQL, for storing the greeting across reloads.
- **Depends on:** Go API, for reading and updating the stored greeting.
- **Depends on:** Next.js frontend, for rendering the approved greeting screen.
- **Assumption:** One shared greeting exists for the app, not per visitor. If false, sign-in or visitor identity scope would be needed and is out of scope.
- **Assumption:** Greeting copy is English by default, but visitor-entered text is displayed as entered after trimming. If false, localization requirements would be needed and are out of scope.

| Open question | Proposed default | Who decides |
|---|---|---|
| — | None; stakeholder brief and approved design decide current scope | — |

## 8. Traceability

Every plan item in this module appears exactly once, and every requirement id traces to a test case. A gap in this table is a gap in the build.

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Editable persisted greeting | GREETING-001, GREETING-002, GREETING-003 | `test-cases/editable-persisted-greeting.md` |
