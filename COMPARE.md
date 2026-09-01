# Second half - run the same task on BMAD and on your own agent

**Time:** ~30-40 min · **Format:** solo or pairs, help on tap

You just built `invoices` and watched our kit rewind to the spec. Now run the **exact same
task** on at least one other setup and compare. The point is not which tool wins - it is to
see how differently each one handles the handoff between phases and, above all, **what it
does when the spec turns out to be wrong**.

## The task (identical to the flagship)

On a clean checkout of `sample-repo/` (or `sample-repo-go/`):

> Build the `invoices` service from `spec.md`. Get the test suite green.

The spec omits the blank rule the sibling services assume. A good run notices the code and
the tests disagree and **fixes the spec**; a shallow run just patches `invoices` and moves
on, leaving the cause open.

## Run it on two of these

- **Our kit (baseline)** - you already did this; keep the result to compare against.
- **BMAD** - install BMAD-METHOD per its README, then run its planning-to-dev flow on the
  same repo and task. Watch how its role personas (analyst / PM / architect / dev / QA) pass
  the work down, and whether the QA/dev step routes the blank-rule gap back to the analyst or
  just patches the code.
- **Your own agent** - whatever you actually use: Codex without our skills, Claude Code with
  your own subagents, Cursor, a custom loop. Give it the same one-line task and nothing else.

## Compare on these axes

Fill this in for each setup you ran:

| Axis | What to watch |
|------|---------------|
| **Caught the spec gap?** | Did anything notice the code satisfied the spec but broke the convention - or did it only see "a failing test"? |
| **Rewind or patch?** | Did it fix the **spec** (cause) or just edit `invoices` (symptom)? The whole game. |
| **Handoffs** | How does work pass between phases/roles - explicit artifacts, or one long context? Where did it blur? |
| **Ceremony vs speed** | How much setup and process before the first line of code? Worth it for this size of task? |
| **Guardrails** | Could you scope what it may touch and when it stops? Or was it all-or-nothing? |

## Leave with

- One sentence per setup: **when would you actually reach for it?**
- The honest answer to: does your **own** everyday setup catch a root cause, or does it patch
  symptoms and let the next feature reopen the gap?
- If BMAD (or your setup) never rewound, that is a finding, not a failure - name what it would
  take to make it route to the cause.

## Note for the room

BMAD is heavy by design - full role personas, planning artifacts. On a task this small it
will feel like overkill, and that is part of the lesson: match the ceremony to the size of
the work. You are not adopting BMAD today; you are feeling where its weight pays off and
where it does not.
