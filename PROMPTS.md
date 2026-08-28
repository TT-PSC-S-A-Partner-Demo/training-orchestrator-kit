# Build your own orchestrator - your SDLC as a chain of skills

The finale of Session 2. One orchestrator (the root agent) drives four SDLC stages:
**spec -> implement -> test -> review**. The first three are skills; review is a
read-only subagent. One prompt runs the whole thing.

## What's in this kit

```
demo-kit-orchestrator/
  skills/spec-change/SKILL.md         # stage 1: request -> spec + acceptance criteria
  skills/implement-to-spec/SKILL.md   # stage 2: smallest diff to the spec
  skills/test-change/SKILL.md         # stage 3: one test per criterion, green
  agents/reviewer.toml                # stage 4: read-only reviewer subagent
  config-snippet.toml                 # the [agents] block for .codex/config.toml
```

These are a working reference. In the exercise you can use them as-is, or write your own
three skills from the Exercise 2 shape and keep only the reviewer.

## Setup (once)

```bash
# install the three skills (any agent's skills dir; Codex shown)
cp -r skills/spec-change       ~/.codex/skills/spec-change
cp -r skills/implement-to-spec ~/.codex/skills/implement-to-spec
cp -r skills/test-change       ~/.codex/skills/test-change
/skills                         # confirm all three loaded

# register the reviewer subagent: paste config-snippet.toml into
# .codex/config.toml, and copy agents/reviewer.toml next to it (.codex/agents/)
```

## Pick a concrete change to drive the pipeline

Use one of the other demo kits as the target repo, with a real, small change:

- **Calculator** (`demo-kit-calc/`): "Add a divide-by-zero guard - dividing by zero shows
  'Cannot divide by zero', not Infinity."
- **go-jira** (`demo-kit-gojira/`): "Harden DateFormat - a blank input returns a clear
  error instead of the opaque time.Parse error."

## Run the whole SDLC from one prompt

In the target repo, prompt the root agent once:

```
Use the spec-change, implement-to-spec, and test-change skills in order, then the
reviewer subagent. The change: add a divide-by-zero guard to the calculator so
dividing by zero shows "Cannot divide by zero" instead of Infinity.
Spec it, implement it, test it, then have the reviewer check the diff.
Stop and show me between each stage.
```

Watch each stage's output feed the next: the spec's acceptance criteria drive the diff,
the diff drives the tests, the reviewer checks the lot - read-only, in its own context.

## Done when

- One prompt runs all four stages **in order**.
- Tests end **green**, one per acceptance criterion.
- The reviewer runs as a **separate read-only** agent (you see it as its own thread).
- You can point to the spec line that drove the code.

## Best practice being taught

- Chain of **3-4 stages**, not more - past that the handoffs blur and context fills.
- `max_depth = 1` - the reviewer is a child, it cannot spawn its own children.
- Subagents **read-only** unless a stage truly needs to write.
- This is the hand-built version of what Session 3's named methods (BMAD, Spec Kit) do.
