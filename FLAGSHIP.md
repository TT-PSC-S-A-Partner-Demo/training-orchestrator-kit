# Flagship exercise (Session 3) - Build your own SDLC agent-team, then audit its guardrails

**Time:** ~60 min, four milestones · **Tool:** Codex CLI (or Claude Code / Devin) with real
subagents · **Target:** `sample-repo/` (build a feature) or a repo you brought

You are **building** the orchestrator, not installing ours. By the end you have a workflow
you designed - phases, handoffs, a rewind, and guardrails - that you can defend, plus a clean
report from the `guardrail-audit` skill on your own work.

> This kit ships a **reference solution** (`skills/`, `agents/`, `config-snippet.toml`). Do
> not copy it to start. Build yours first; open the reference only to get unstuck or to
> compare afterwards. Copying it defeats the exercise.

## What you start with (that is all)

- A **target**: build an `invoices` service in `sample-repo/` from `sample-repo/spec.md`
  (the spec is deliberately thin), or the same shape on a repo you brought.
- The **`guardrail-audit`** skill, installed, to check your workflow at the end:
  ```bash
  mkdir -p ~/.codex/skills && cp -r skills/guardrail-audit ~/.codex/skills/
  ```
- Nothing else is set up for you. The workflow is yours to build.

## Milestone A - build the pipeline (20 min)

Get one prompt to drive four stages - **spec -> implement -> test -> review** - on the task.
You decide the shape:

- Which stages are **skills** (a repeatable procedure) and which is a **subagent** (its own
  context)? A read-only reviewer subagent is the usual first one to split out - why?
- How does each stage **hand off** to the next - what artifact does it leave behind?
- Write the skills / subagent yourself (`~/.codex/skills/<name>/SKILL.md`, an agent `.toml`
  or `.md`). Keep each one small.

**Done:** one prompt takes the invoices request to a diff plus tests, with the review as its
own step. It does not have to be right yet - it has to *run*.

## Milestone B - make it rewind (15 min)

Run it and a test will fail: the code matched the spec, but the spec forgot the rule the
sibling services already follow. Now build the thing that makes this a *team*, not a chain:

- Have the test stage attach, to each finding, the **phase that caused it** (here: analysis -
  the spec never stated the rule), not the phase that found it.
- Make your orchestrator **route back** to that phase and continue - fix the spec, then
  re-implement - instead of patching the code.

**Done:** you can point at where your workflow decides "this goes back to spec", and it does.

## Milestone C - add the guardrails (15 min)

Make it safe to leave running. Add - and be able to justify - each of:

- a **sandbox** scope (what may it touch?),
- an **approval** point (where does it stop for a human?),
- a **give-up** limit (how does the loop quit - iterations, rewinds, no-progress?),
- a **read-only** reviewer (least privilege),
- a **tool allowlist** (only what the task needs).

You choose the values. There is no per-run cost cap in the tools, so bound cost with these.

## Milestone D - audit your own work (10 min)

Point the `guardrail-audit` skill at what you built:

> audit the guardrails on my orchestrator: config.toml, the agent files, and my driving prompt

It grades seven guardrails PRESENT / MISSING / RISKY and gives the fix for each. **Fix every
gap, then re-audit until the verdict is SAFE.** Then confirm the real gate:

```bash
cd sample-repo    && python -m pytest -q     # green only when the rule reached the spec
cd sample-repo-go && go test ./...           # Go variant
```

## Done when

- The workflow is **yours** - you can explain every stage, the rewind, and each guardrail.
- `guardrail-audit` reports **SAFE**, and the test suite is green.
- If you opened the reference, you can name one thing you did differently and why.

## Then: the same task on BMAD and your own agent (second half)

You built a workflow by hand. Now run the *same* task through BMAD and through your everyday
agent, and see which of them would have caught the spec gap and rewound - and which just
patches. See `COMPARE.md`.
