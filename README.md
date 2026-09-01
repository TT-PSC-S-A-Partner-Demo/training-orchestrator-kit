# orchestrator-kit - build an SDLC agent-team, and audit its guardrails

The Session 3 flagship: you **build** your own orchestrator - four SDLC stages
(spec &rarr; implement &rarr; test &rarr; review) with a rewind to root cause and real guardrails -
then run the `guardrail-audit` skill on your own work. Full build guide in `FLAGSHIP.md`.

> **Not copy-paste.** This kit ships a **reference solution** (`skills/spec-change`,
> `implement-to-spec`, `test-change`, `agents/`, `config-snippet.toml`). Build yours first;
> open the reference only to get unstuck or to compare afterwards.

## Contents

```
guardrail-audit/  ->  skills/guardrail-audit/SKILL.md   # the checker you RUN on your build
sample-repo/                        # Python target: 3 consistent services + spec.md (build a 4th)
sample-repo-go/                     # Go target: the same, go test ./...

# --- reference solution (one way to build it - do not copy to start) ---
skills/spec-change  implement-to-spec  test-change       # the three stage skills
agents/{tester,reviewer}.toml       # Codex subagents
agents-claude/{tester,reviewer}.md  # Claude Code subagents (.claude/agents/)
agents-devin/{tester,reviewer}.md   # Devin subagents (.devin/agents/)
config-snippet.toml                 # the [agents] block

FLAGSHIP.md                         # Session 3: build it yourself, four milestones
COMPARE.md                          # Session 3 second half: same task on BMAD + your own agent
PROMPTS.md                          # Session 2: the forward-only chain
```

## The one thing you install up front

The guardrail checker, so you can audit your workflow at the end:

```bash
mkdir -p ~/.codex/skills && cp -r skills/guardrail-audit ~/.codex/skills/
# Claude: cp -r skills/guardrail-audit ~/.claude/skills/   |   Devin: .devin/skills/
```

Then run it on your build: *"audit the guardrails on my orchestrator - config.toml, the agent
files, and my driving prompt."* It grades seven guardrails and gives the fix for each.

## The target

```bash
cd sample-repo    && python -m pytest -q     # 9 passed, 3 failed = invoices to build
cd sample-repo-go && go test ./...           # Go variant
```

`orders`, `billing`, `shipping` share one convention (reject blank input) that lives only in
the code. The task is to build a fourth service, `invoices`, from `sample-repo/spec.md` - a
spec that forgets to state that rule. A good workflow catches the gap and rewinds to the
spec; a shallow one patches the code and lets the next service reopen it.

## Reference solution

`skills/` and `agents/` are one working way to build the pipeline (same roles for Codex
`.toml`, Claude Code and Devin `.md`). Use them to compare after you have built your own -
not before.
