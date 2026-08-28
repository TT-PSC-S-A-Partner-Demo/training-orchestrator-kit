---
name: implement-to-spec
description: Implements a change to match an existing spec with the smallest diff that satisfies its acceptance criteria. Use when asked to implement a spec, build to acceptance criteria, or make the change after a spec exists in an SDLC pipeline.
---

# Implement to spec

Build exactly what the spec asks, nothing more.

## Instructions

1. Read the spec from the previous stage - Goal, Acceptance criteria, Files.
2. Make the **smallest diff** that meets every acceptance criterion.
3. Touch only the files the spec named. Do not refactor unrelated code, rename things, or add features not in scope.
4. If the spec is ambiguous, note the assumption in one line rather than expanding scope.
5. Show the diff. Stop - hand off to the test stage. Do not write tests here.

## Output

- A diff limited to the spec's files.
- One line per assumption made, if any.
