# Design System — Hello World Acceptance 3

> Source of truth: the approved `index.html`.
> Every value below is extracted from it. Changing a value here without changing the approved design is a defect.

Last updated: 2026-09-14

## 1. Foundations

### 1.1 Color

Semantic tokens. Name by job, never by hue.

| Token | Value | Used for |
|---|---|---|
| `--color-bg` | `#FFFFFF` | Page background, input background, button text |
| `--color-text` | `#000000` | Heading text, input text, input border |
| `--color-primary-action` | `#2563EB` | Save button background and border |
| `--color-primary-action-hover` | `#1D4ED8` | Save button hover background and border |
| `--color-focus` | `#2563EB` | Input and button focus-visible outline |

#### Contrast audit

Every text-on-background pair actually used. Body text ≥ 4.5:1, large text (≥ 18.66px bold or ≥ 24px) ≥ 3:1, UI borders ≥ 3:1.

| Foreground | Background | Ratio | Passes |
|---|---|---|---|
| `--color-text` | `--color-bg` | `21.0:1` | AA |
| `--color-bg` | `--color-primary-action` | `5.2:1` | AA |
| `--color-bg` | `--color-primary-action-hover` | `6.7:1` | AA |
| `--color-text` input border | `--color-bg` | `21.0:1` | AA |
| `--color-focus` outline | `--color-bg` | `5.2:1` | AA |

### 1.2 Spacing

Base unit: `2px`. Every margin, padding, and gap in approved design uses one of these values.

| Token | Value |
|---|---|
| `--space-0` | `0` |
| `--space-6` | `12px` |
| `--space-7` | `14px` |
| `--space-8` | `16px` |
| `--space-11` | `22px` |
| `--space-12` | `24px` |
| `--space-16` | `32px` |

Layout dimensions extracted from approved design:

| Token | Value | Used for |
|---|---|---|
| `--size-control` | `48px` | Text input and Save button height |
| `--size-container-max` | `560px` | Main content max width |
| `--size-screen-padding-desktop` | `24px` | Body padding above 520px |
| `--size-screen-padding-phone` | `16px` | Body padding at 520px and below |

### 1.3 Typography

Font families:

- Body: `Arial, Helvetica, sans-serif`, loaded from system fonts.
- Headings: `Arial, Helvetica, sans-serif`, inherited from body.

| Token | Size | Line height | Weight | Used for |
|---|---|---|---|---|
| `--text-body` | `18px` | `48px` on input, normal on button via `font: inherit` then explicit size | `400` inherited | Text input value |
| `--text-button` | `18px` | normal browser line-height inside 48px control | `700` | Save button |
| `--text-heading` | `clamp(48px, 10vw, 88px)` | `1` | `700` | h1 greeting |

Heading levels are used in order: only `h1` appears.

| Token | Value | Used for |
|---|---|---|
| `--font-weight-body` | `400` | Input text and inherited body text |
| `--font-weight-action` | `700` | Save button label |
| `--font-weight-heading` | `700` | Greeting h1 |
| `--tracking-tight-heading` | `-0.04em` | Greeting h1 |

### 1.4 Radius, border, shadow, motion

| Token | Value | Used for |
|---|---|---|
| `--radius-none` | `0` | Input and Save button corners |
| `--border-width-control` | `1px` | Input and Save button border |
| `--focus-outline-width` | `3px` | Focus-visible outline |
| `--focus-outline-offset` | `3px` | Focus-visible outline offset |

No shadows appear in approved design. No CSS transitions or animation appear in approved design.

### 1.5 Layout and breakpoints

| Name | Rule | Container | Columns | Gutter |
|---|---|---|---|---|
| `base` | all widths | `max-width: 560px` | one centered section | form gap `12px` |
| `phone` | `@media (max-width: 520px)` | `width: 100%` | form stacks vertically | form gap `12px` |

No z-index values appear in approved design.

## 2. Components

### 2.1 PageShell

**Purpose** — Center one greeting section on plain page background.

**Anatomy** — `[main] [section.greeting-section]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--color-bg`, `--color-text`, `--size-container-max`, `--size-screen-padding-desktop`, `--size-screen-padding-phone` | Only app page |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Desktop | `min-height: calc(100vh - 48px)` | `24px` body padding | inherited body |
| Phone | `min-height: calc(100vh - 32px)` | `16px` body padding | inherited body |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | White centered page, no decoration | `--color-bg`, `--color-text` |

**Accessibility** — Use semantic `main`. Section is labelled by greeting h1.

### 2.2 GreetingHeading

**Purpose** — Display current persisted greeting as primary page content.

**Anatomy** — `[h1#greeting-heading]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--text-heading`, `--font-weight-heading`, `--tracking-tight-heading`, `--space-16` | Current greeting |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Responsive | line-height `1` | margin-bottom `32px` | `--text-heading` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Black large centered heading | `--color-text`, `--text-heading` |

**Accessibility** — Use one `h1`; do not skip heading level. Heading must receive updated greeting text after save.

### 2.3 GreetingForm

**Purpose** — Let visitor submit new greeting.

**Anatomy** — `[visually hidden label] [text input] [Save button]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Row | `--space-6` | Width above 520px |
| Stacked | `--space-6` | Width at 520px and below |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Full width | controls `48px` | gap `12px` | inherited body |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Controls centered in one row, stacked on phone | `--space-6` |

**Accessibility** — Form has visible controls and programmatic label. Submit by Enter or Save button.

### 2.4 TextInput

**Purpose** — Edit greeting text.

**Anatomy** — `[input#greeting-input]` with hidden text label `Greeting`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--color-bg`, `--color-text`, `--radius-none`, `--border-width-control`, `--size-control`, `--text-body` | Greeting text entry |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | `48px` | `0 14px` | `--text-body` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | White field, black text, black 1px border | `--color-bg`, `--color-text`, `--border-width-control` |
| Focus visible | 3px blue outline with 3px offset | `--color-focus`, `--focus-outline-width`, `--focus-outline-offset` |

**Accessibility** — Native text input. Hidden label remains accessible. Minimum hit target is 48px high. Required field keeps empty greeting from saving.

### 2.5 PrimaryButton

**Purpose** — Submit changed greeting.

**Anatomy** — `[button type="submit"] [label]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Primary action | `--color-primary-action`, `--color-bg`, `--radius-none`, `--border-width-control`, `--size-control`, `--text-button` | Save greeting |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | `48px` | `0 22px` | `--text-button` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Blue fill, blue border, white bold label | `--color-primary-action`, `--color-bg`, `--text-button` |
| Hover | Darker blue fill and border | `--color-primary-action-hover`, `--color-bg` |
| Focus visible | 3px blue outline with 3px offset | `--color-focus`, `--focus-outline-width`, `--focus-outline-offset` |

**Accessibility** — Native button. Minimum hit target is 48px high. Keyboard activation uses native Enter/Space behavior. Visible focus indicator required.

## 3. Content and formatting

- Voice and tone: plain, direct, minimal.
- Locale: English content; no date, time, number, or currency formats appear.
- Capitalization: heading uses stakeholder-provided greeting text; button uses title case `Save`; label uses title case `Greeting`.
- Empty-state and error-message wording pattern: no empty or error messages are drawn; empty input returns focus to input.

## 4. Known deviations

| Where | Deviation | Why it stands | Follow-up |
|---|---|---|---|
| Text input and Save button horizontal padding | Uses `14px` and `22px`, outside common 4px spacing scale | Approved design uses these exact control paddings | Keep unless stakeholder asks for tighter token scale |
| Save button | Hover state exists but no active or disabled state is drawn | Approved design only draws default, hover, and focus-visible | Add active or disabled only when product behavior requires it |
| Data-loading view | No loading, empty, or error state is drawn | Approved design shows success path only for one acceptance page | Add states only if story requirements need visible API states |

AI default checks: approved design avoids purple/indigo defaults, decorative gradients, maximum rounding, heavy shadows, generic multi-section layout, emoji iconography, filler copy, removed focus states, text over images, and hover-only affordances.

## 5. Change log

| Date | Change | Design PR |
|---|---|---|
| 2026-09-14 | Initial design system extracted from approved `index.html` | This PR |
