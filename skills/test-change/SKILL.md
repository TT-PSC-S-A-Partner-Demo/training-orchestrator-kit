---
name: test-change
description: Writes and runs tests for a change until they pass green, one test per acceptance criterion. Use when asked to test a change, add tests, verify acceptance criteria, or run the test stage of an SDLC pipeline after implementing.
---

# Test a change

Turn the spec's acceptance criteria into tests, then make them green.

## Instructions

1. Read the spec's acceptance criteria and the diff from the implement stage.
2. Write **one test per acceptance criterion**, plus the original behavior so a regression shows up.
3. Run the tests. Report the command and the result.
4. If red, fix the smallest thing and re-run. Repeat until green. Do not weaken a test to pass it.
5. Report: the test command, pass/fail count, and which acceptance criterion each test covers. Stop - hand off to the reviewer.

## Output

- The test command and its green result.
- A one-line map: acceptance criterion -> test that proves it.
