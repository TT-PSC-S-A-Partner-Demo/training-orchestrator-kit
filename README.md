# orchestrator-kit - your SDLC as a team of agents

One root agent drives four SDLC stages - **spec -> implement -> test -> review**. Two ways
to run it:

- **Session 2 (the chain):** the stages go forward once. See `PROMPTS.md`.
- **Session 3 (the flagship):** the team **rewinds** - a defect is routed to the phase that
  *caused* it, not the one that noticed it - on a repo with real convention drift, under
  guardrails, gated by real tests. See `FLAGSHIP.md`.

## Contents

```
skills/spec-change/SKILL.md         # stage 1: request -> spec + acceptance criteria
skills/implement-to-spec/SKILL.md   # stage 2: smallest diff to the spec
skills/test-change/SKILL.md         # stage 3: one test per criterion, green

agents/{tester,reviewer}.toml       # Codex subagents (.toml)
agents-claude/{tester,reviewer}.md  # Claude Code subagents (.claude/agents/)
agents-devin/{tester,reviewer}.md   # Devin subagents (.devin/agents/)
config-snippet.toml                 # the [agents] block for .codex/config.toml

sample-repo/                        # Python: 3 services whose normalize() drifted
sample-repo-go/                     # Go: the same drift (go test ./...)

PROMPTS.md                          # Session 2: the chain
FLAGSHIP.md                         # Session 3: the rewind flagship, four stages
```

The tester + reviewer subagents are the same roles for all three tools - Codex uses `.toml`,
Claude Code and Devin use `.md` with YAML frontmatter. Same target_phase routing rule.

## Setup

```bash
# --- Codex CLI ---
mkdir -p ~/.codex/skills ~/.codex/agents
cp -r skills/spec-change skills/implement-to-spec skills/test-change ~/.codex/skills/
cp agents/*.toml ~/.codex/agents/           # -> ~/.codex/agents/{tester,reviewer}.toml
# paste config-snippet.toml into ~/.codex/config.toml (config_file = ./agents/*.toml)

# --- Claude Code ---
mkdir -p ~/.claude/agents ~/.claude/skills
cp agents-claude/*.md ~/.claude/agents/
cp -r skills/* ~/.claude/skills/

# --- Devin (relative to the repo root) ---
mkdir -p .devin/agents .devin/skills
cp agents-devin/*.md .devin/agents/
cp -r skills/* .devin/skills/
```

## The drift repo (Session 3)

```bash
cd sample-repo    && python -m pytest -q && cd ..    # 4 failed, 5 passed = the drift
cd sample-repo-go && go test ./...        && cd ..    # same drift in Go
```

`orders`, `billing`, `shipping` each handle blank input differently. The failures are one
**analysis** gap (no agreed convention), not three bugs - which is why the flagship rewinds
to spec instead of patching each service. Full walk-through in `FLAGSHIP.md`.
