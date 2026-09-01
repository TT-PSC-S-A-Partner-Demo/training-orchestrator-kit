# Flagship exercise (Session 3) - An agent team that rewinds to root cause

**Time:** ~60 min, four stages · **Tool:** Codex CLI with real subagents · **Repo:** `sample-repo/`
(three services that drifted apart)

This is the Session 2 orchestrator, grown up. There, the chain only went forward:
spec &rarr; implement &rarr; test &rarr; review. Here it goes **backward when it has to** - a defect
is routed to the phase that *caused* it, not the one that noticed it. You drive real Codex
subagents on a repo with real convention drift, under real guardrails, gated by real tests.

> **Coming from Session 2?** If you did not finish the orchestrator chain there, Stage A is
> where you land it. If you did, Stage A is a 5-minute warm-up.

## The repo: real convention drift

`sample-repo/services/` has three services - `orders`, `billing`, `shipping` - each with a
`normalize()` that handles blank/`None` input **differently**. `shipping` rejects it loudly;
`orders` and `billing` drifted. `test_conventions.py` demands they behave the same:

```bash
cd sample-repo
python -m pytest -q          # 4 failed, 5 passed  <- the drift, made visible
cd ..
```

The failures are not three separate bugs. Their root cause is one **analysis** gap: the
team never wrote down the rule, so three authors each guessed. That is what makes rewind
matter here.

**Go instead of Python?** `sample-repo-go/` is the same drift in Go - three packages
(`orders`, `billing`, `shipping`) whose `Normalize` disagrees, and a test that demands they
agree. The gate is `go test ./...` instead of pytest; everything else in this walk-through
is identical.

```bash
cd sample-repo-go
go test ./...                 # FAIL: orders + billing drift; shipping is right
cd ..
```

## Pre-warm (already done on the workshop image)

The kit ships the same tester + reviewer subagents for **all three tools** - only the
folder and format differ. Install the pair for whichever you use:

```bash
# --- Codex CLI ---
mkdir -p ~/.codex/skills ~/.codex/agents
cp -r skills/spec-change skills/implement-to-spec skills/test-change ~/.codex/skills/
cp agents/*.toml ~/.codex/agents/           # -> ~/.codex/agents/{tester,reviewer}.toml
# paste config-snippet.toml into ~/.codex/config.toml; its config_file paths are
# "./agents/<name>.toml", i.e. relative to that config.toml (so ~/.codex/agents/)
/plugins                                     # confirm the skills loaded

# --- Claude Code ---
mkdir -p ~/.claude/agents ~/.claude/skills
cp agents-claude/*.md ~/.claude/agents/      # -> ~/.claude/agents/{tester,reviewer}.md
cp -r skills/* ~/.claude/skills/             # same SKILL.md files, portable

# --- Devin (paths are relative to the repo root) ---
mkdir -p .devin/agents .devin/skills
cp agents-devin/*.md .devin/agents/          # -> .devin/agents/{tester,reviewer}.md
cp -r skills/* .devin/skills/
```

The subagent instructions are identical across the three - Codex uses `.toml`, Claude Code
and Devin use `.md` with YAML frontmatter. Same roles, same target_phase routing rule.

---

## Stage A - run it, watch the rewind (10 min)

Run the whole loop from one prompt in the repo. Paste this to the root agent:

```
You are the orchestrator for an SDLC loop on this repo. Drive these phases, pausing
to show me each:
  1. spec-change: agree ONE convention for normalize() across all services, with
     acceptance criteria (blank input is rejected with a clear error; valid input
     is trimmed). [Python: raise ValueError; Go: return an error]
  2. implement-to-spec: apply that convention to every service that drifts.
  3. tester subagent: run the suite; each failure carries a target_phase.
  4. reviewer subagent: read-only check of the diff.

Rewind rule: if a finding's target_phase is earlier than the current phase, go BACK
to that phase and continue - do NOT patch forward. A convention that N modules
implement N ways is an analysis gap: rewind to spec, decide it once, reapply.

Stop after 4 rewinds (give up and report) or when the suite is green and review is clean.
```

Watch the tester tag the drift as **target_phase: analysis**, and the orchestrator rewind to
`spec-change` to decide the convention - then apply it everywhere. Suite goes green.

> Half the value is seeing the tester refuse to call this "three bugs". It names the cause.

## Stage B - break the routing (10 min)

Run it again, but tell the orchestrator to route naively:

```
Same loop, but ignore target_phase: send every failure straight to implement-to-spec
and patch the service that failed. Do not touch the spec.
```

Now it patches `orders`, re-tests, finds `billing` still drifts, patches `billing`, and if
you had a fourth service it would never end. The convention is still unwritten, so every new
service reopens the drift. This is whack-a-mole across services - the exact failure the
target_phase rule exists to prevent.

## Stage C - harden it: guardrails and give-up (20 min)

An orchestrator you would leave running needs limits on **what it may touch**, **when it
stops for you**, and **when it gives up**. These are real Codex controls, not prose:

- **Sandbox + approvals** on the run:
  ```bash
  codex exec --sandbox workspace-write --ask-for-approval on-request "<the Stage A prompt>"
  ```
  `workspace-write` keeps every subagent inside the repo; network is off by default;
  `on-request` pauses before the riskier steps. The reviewer subagent is already
  `sandbox_mode = "read-only"` in `agents/reviewer.toml` - it can never write.
- **Give-up** is in the prompt: *stop after 4 rewinds*. Prove it - point the loop at a
  spec you keep contradicting and watch it stop and report instead of looping.
- **Write lane:** the sandbox is the allowlist. Try `--sandbox read-only` and watch the
  implement stage refuse to edit - useful for a dry run before you let it write.

There is no per-run **cost** cap in Codex - you bound cost with scope, approvals, and the
rewind limit, and you watch `/status`. Know that going in.

## Stage D - the real gate (20 min)

The test suite is the referee - nothing merges on opinion. With root-cause routing on:

```bash
cd sample-repo    && python -m pytest -q && cd ..    # Python
cd sample-repo-go && go test ./...        && cd ..    # Go variant
```

The loop is done only when this is green **for all three services** - which happens only
when the convention was decided once (in spec) and applied everywhere, not patched per
service. Re-run from the drifted state to repeat the stage cleanly (git restore the
services, or re-clone).

> Force a rewind on a real failure: revert one service by hand, re-run the loop, and watch
> the tester re-open the analysis finding and the orchestrator rewind to spec again.

## Done when

- One prompt drives spec &rarr; implement &rarr; test &rarr; review with **real subagents** (you see
  the tester and reviewer as their own threads).
- You have watched the loop **rewind to analysis** to fix the convention once, not patch
  three services.
- You have seen naive routing whack-a-mole, and the give-up rule stop a bad loop.
- The suite is green for all three services, under `workspace-write` + `on-request`.

## Then: bring your own (second half)

Take this shape to a repo you own - typically your own convention drift across services.
See the deck's "bring your own" slide and the three variants there.
