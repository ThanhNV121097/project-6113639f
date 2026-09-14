# Test cases — Editable persisted greeting

Module: `greeting`
Function: Editable persisted greeting
Story: `docs/greeting/stories/editable-persisted-greeting.md`
SRS: `docs/greeting/SRS.md`
Service contract: `docs/architecture/services.md`

Risk level: Medium. This story writes shared persisted state through browser, API, and PostgreSQL, but has one actor, one resource, no authentication, and no external services.

## UI and persistence scenarios

**Scenario**: Default stored greeting appears as large centered heading
**Given**: Stored greeting is `Hello, World!`.
**When**: Visitor opens app page.
**Then**: Page contains exactly one `h1` greeting heading with text `Hello, World!`; heading is large, bold, and centered in one greeting section.
Traces: SC-1 (GREETING-001 AC-1)
Check: measure_styles

**Scenario**: Existing stored greeting appears as heading
**Given**: Stored greeting is `Good morning`.
**When**: Visitor opens app page.
**Then**: Page heading text is exactly `Good morning`.
Traces: SC-2 (GREETING-001 AC-2)
Check: render_url

**Scenario**: Existing stored greeting appears in input
**Given**: Stored greeting is `Good morning`.
**When**: Visitor opens app page.
**Then**: Text input labelled `Greeting` has value `Good morning`.
Traces: SC-3 (GREETING-001 AC-3)
Check: render_url

**Scenario**: Visitor saves non-blank greeting with button
**Given**: Page shows `Hello, World!` as heading and input value.
**When**: Visitor types `Pipeline accepted` in input labelled `Greeting` and clicks `Save`.
**Then**: Heading changes to `Pipeline accepted`; input value is `Pipeline accepted`.
Traces: SC-4 (GREETING-002 AC-1)
Check: interact_page

**Scenario**: Saved greeting persists after reload
**Given**: Visitor saved `Pipeline accepted` and page heading shows `Pipeline accepted`.
**When**: Visitor reloads app page.
**Then**: Heading text is exactly `Pipeline accepted`; input value is exactly `Pipeline accepted`.
Traces: SC-5 (GREETING-002 AC-2)
Check: interact_page

**Scenario**: Blank trimmed input is not saved
**Given**: Page shows `Hello, World!` as heading and input value.
**When**: Visitor replaces input value with three spaces and submits form.
**Then**: Heading remains `Hello, World!`; stored greeting remains `Hello, World!`; focus returns to input labelled `Greeting`; no error message is required.
Traces: SC-6 (GREETING-002 AC-3)
Check: interact_page

**Scenario**: Saved greeting is trimmed before display
**Given**: Page shows `Hello, World!` as heading and input value.
**When**: Visitor types `  Hello again  ` in input labelled `Greeting` and submits form.
**Then**: Heading changes to `Hello again`; input value is `Hello again`; stored greeting is `Hello again`.
Traces: SC-7 (GREETING-002 AC-4)
Check: interact_page

**Scenario**: Page has one centered section and no navigation
**Given**: Visitor opens app page.
**When**: Page structure is inspected.
**Then**: Page contains one `main` with one `section.greeting-section`; section contains one `h1#greeting-heading`, one input `#greeting-input` programmatically labelled `Greeting`, and one `Save` submit button; page contains no `nav` and no other sections.
Traces: SC-8 (GREETING-003 AC-1)
Check: render_url

**Scenario**: Page uses approved colors
**Given**: Visitor opens app page.
**When**: Computed colors are measured.
**Then**: Page background is `#FFFFFF`; primary text is `#000000`; Save button default background is `#2563EB`.
Traces: SC-9 (GREETING-003 AC-2)
Check: measure_styles

**Scenario**: Controls are 48px tall at 390px viewport
**Given**: Viewport width is 390px.
**When**: Visitor opens app page.
**Then**: Text input height is 48px; Save button height is 48px; page has no horizontal scroll.
Traces: SC-10 (GREETING-003 AC-3)
Check: measure_styles

**Scenario**: Page has no animation or transition
**Given**: Visitor opens app page.
**When**: Computed motion styles are measured.
**Then**: Greeting section, heading, input, and Save button have no animation and no transition.
Traces: SC-11 (GREETING-003 AC-4)
Check: measure_styles

**Scenario**: Form stacks at 520px viewport
**Given**: Viewport width is 520px.
**When**: Visitor opens app page.
**Then**: Text input and Save button are stacked vertically; each control uses full available form width.
Traces: SC-12 (GREETING-003 AC-5)
Check: measure_styles

**Scenario**: Any visitor may save greeting without sign-in
**Given**: Visitor has no session, token, or credentials, and page shows `Hello, World!`.
**When**: Visitor types `No sign in needed` and clicks `Save`.
**Then**: Heading changes to `No sign in needed`; no sign-in prompt appears.
Traces: GREETING-002 permission behaviour
Check: interact_page

**Scenario**: Latest completed save wins for shared greeting
**Given**: Visitor A and Visitor B both open page while stored greeting is `Hello, World!`.
**When**: Visitor A saves `First save`, then Visitor B saves `Second save`, and page is reloaded.
**Then**: Heading text after reload is exactly `Second save`.
Traces: GREETING-002 conflict behaviour
Check: interact_page

**Scenario**: Keyboard focus outline appears on input and button
**Given**: Visitor opens app page.
**When**: Visitor tabs to input and then tabs to Save button.
**Then**: Focused input and focused Save button each show 3px blue focus outline with 3px offset.
Traces: GREETING-003 focus behaviour; Accessibility NFR
Check: interact_page

**Scenario**: Save button hover uses darker blue
**Given**: Visitor opens app page.
**When**: Pointer hovers Save button.
**Then**: Save button background is `#1D4ED8`.
Traces: GREETING-003 hover behaviour
Check: measure_styles

## API contract scenarios

**Scenario**: GET greeting returns stored greeting JSON
**Given**: Stored greeting is `Hello, World!`.
**When**: Client sends `GET /v1/greeting`.
**Then**: Response status is `200`; `Content-Type` is `application/json; charset=utf-8`; response body is exactly `{"greeting":"Hello, World!"}`.
Traces: contract (GET /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting saves and returns trimmed greeting JSON
**Given**: Stored greeting is `Hello, World!`.
**When**: Client sends `PUT /v1/greeting` with JSON body `{"greeting":"  Contract saved  "}`.
**Then**: Response status is `200`; `Content-Type` is `application/json; charset=utf-8`; response body is exactly `{"greeting":"Contract saved"}`; later `GET /v1/greeting` returns `{"greeting":"Contract saved"}`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT malformed JSON returns invalid request envelope
**Given**: Stored greeting is `Hello, World!`.
**When**: Client sends `PUT /v1/greeting` with malformed JSON body `{`.
**Then**: Response status is `400`; `Content-Type` is `application/json; charset=utf-8`; response body has `error.code` equal to `invalid_request` and safe plain-text `error.message`; later `GET /v1/greeting` still returns `{"greeting":"Hello, World!"}`.
Traces: contract (PUT /v1/greeting error envelope)
Check: fetch_url

**Scenario**: PUT missing greeting returns invalid request envelope
**Given**: Stored greeting is `Hello, World!`.
**When**: Client sends `PUT /v1/greeting` with JSON body `{}`.
**Then**: Response status is `400`; response body has `error.code` equal to `invalid_request`; stored greeting remains `Hello, World!`.
Traces: contract (PUT /v1/greeting missing greeting)
Check: fetch_url

**Scenario**: PUT non-string greeting returns invalid request envelope
**Given**: Stored greeting is `Hello, World!`.
**When**: Client sends `PUT /v1/greeting` with JSON body `{"greeting":123}`.
**Then**: Response status is `400`; response body has `error.code` equal to `invalid_request`; stored greeting remains `Hello, World!`.
Traces: contract (PUT /v1/greeting non-string greeting)
Check: fetch_url

**Scenario**: PUT blank trimmed greeting returns invalid request envelope
**Given**: Stored greeting is `Hello, World!`.
**When**: Client sends `PUT /v1/greeting` with JSON body `{"greeting":"   "}`.
**Then**: Response status is `400`; response body has `error.code` equal to `invalid_request`; response body `error.message` is `Greeting must not be blank.`; stored greeting remains `Hello, World!`.
Traces: contract (PUT /v1/greeting blank greeting)
Check: fetch_url

**Scenario**: PUT ignores undefined extra fields
**Given**: Stored greeting is `Hello, World!`.
**When**: Client sends `PUT /v1/greeting` with JSON body `{"greeting":"Extra ignored","unused":"value"}`.
**Then**: Response status is `200`; response body is exactly `{"greeting":"Extra ignored"}`; later `GET /v1/greeting` returns `{"greeting":"Extra ignored"}` and no `unused` field.
Traces: contract (PUT /v1/greeting undefined request field)
Check: fetch_url

**Scenario**: API persistence failure returns internal error envelope
**Given**: Database persistence is unavailable.
**When**: Client sends `GET /v1/greeting` or `PUT /v1/greeting` with valid JSON body `{"greeting":"Unavailable"}`.
**Then**: Response status is `500`; response body has `error.code` equal to `internal_error`; response body has generic safe plain-text `error.message`; no stack trace, SQL text, secret, or connection string appears in response.
Traces: contract (API persistence failure)
Check: fetch_url
