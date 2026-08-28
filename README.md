# orchestrator-kit - your SDLC as a chain of skills

Workshop kit for Session 2 (Exercise 4): build your own orchestrator. One root agent
drives four SDLC stages - **spec -> implement -> test -> review**. The first three are
skills; review is a read-only subagent. One prompt runs the whole thing.

## Contents

```
skills/spec-change/SKILL.md         # stage 1: request -> spec + acceptance criteria
skills/implement-to-spec/SKILL.md   # stage 2: smallest diff to the spec
skills/test-change/SKILL.md         # stage 3: one test per criterion, green
agents/reviewer.toml                # stage 4: read-only reviewer subagent
config-snippet.toml                 # the [agents] block for .codex/config.toml
PROMPTS.md                          # setup, a concrete task, the driving prompt
```

## Setup

```bash
cp -r skills/spec-change       ~/.codex/skills/spec-change
cp -r skills/implement-to-spec ~/.codex/skills/implement-to-spec
cp -r skills/test-change       ~/.codex/skills/test-change
# paste config-snippet.toml into .codex/config.toml, copy agents/reviewer.toml to .codex/agents/
```

Then follow `PROMPTS.md`: pick a small change on a target repo (calculator or go-jira) and
drive the whole SDLC from one prompt.
