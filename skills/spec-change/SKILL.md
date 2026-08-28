---
name: spec-change
description: Turns a change request into a short written spec with testable acceptance criteria, before any code is written. Use when asked to spec a change, write acceptance criteria, plan a feature, or when starting the first stage of an SDLC pipeline.
---

# Spec a change

Turn a one-line request into a spec small enough to build against and test against. No code in this stage.

## Instructions

1. Restate the goal in one sentence, in your own words.
2. List 2-4 **acceptance criteria** as observable, testable statements ("dividing by zero shows a message, not Infinity"), not implementation notes.
3. Note what is explicitly **out of scope** for this change.
4. Name the files you expect to touch. Do not edit them yet.
5. Output the spec as markdown under headings: Goal, Acceptance criteria, Out of scope, Files. Stop - hand off to the implement stage.

## Output

- A short markdown spec. No code, no diff.
- Acceptance criteria phrased so a test can pass or fail on each.
