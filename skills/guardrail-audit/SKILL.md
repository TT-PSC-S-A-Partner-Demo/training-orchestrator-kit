---
name: guardrail-audit
description: Audits an agent workflow or config for missing guardrails - sandbox scope, approvals, a give-up limit, a tool allowlist, read-only review, and network egress. Use when asked to check guardrails, review an agent or orchestrator config for safety, audit a workflow before running it unattended, or verify a config.toml / .claude / .devin setup is safe.
---

# Guardrail audit

Check whether an agent workflow is safe to run - especially unattended. You are auditing
**guardrails**, not code quality. Report what is present, what is missing, and the exact fix
for each gap. Do not change anything unless asked; produce a report.

## What to read

Look at whatever the workflow is built from:
- Codex: `config.toml` (`sandbox_mode`, `approval_policy`, `[agents]`, `[mcp_servers.*]`),
  the agent `.toml` files, and the orchestrator prompt.
- Claude Code: `.claude/agents/*.md` frontmatter (`tools`, `permissionMode`, `maxTurns`),
  settings, and the driving prompt.
- Devin: `.devin/agents/*.md` (`allowed-tools`, `max-nesting`) and the playbook/prompt.
- Any orchestrator prompt or script that drives the loop.

## The seven guardrails

Grade each **PRESENT / MISSING / RISKY** and give the one-line fix.

1. **Sandbox scope** - is filesystem/exec scope limited? `sandbox_mode` should be
   `workspace-write` or `read-only` for normal work. `danger-full-access` as the default is
   **RISKY**. Claude: a restricted `tools`/`permissionMode`. Devin: `allowed-tools`.
2. **Approvals** - is there a human stop point? `approval_policy` set (`on-request`/`untrusted`).
   `never` **with** a wide sandbox and no other gate is **RISKY**.
3. **Give-up / stop condition** - does the loop bound itself? An iteration or rewind cap, a
   no-progress rule, `maxTurns`. A loop that cannot quit is **MISSING** a guardrail even if
   nothing else is wrong.
4. **Tool allowlist** - are tools scoped to what the task needs? MCP `enabled_tools` /
   `disabled_tools` / `default_tools_approval_mode = "writes"`; Claude `tools:`; Devin
   `allowed-tools`. A subagent that inherits every tool is **RISKY** for a write-capable role.
5. **Least privilege on read-only roles** - is a reviewer/analyst actually read-only?
   A reviewer with write or exec access is **RISKY** - it should be `sandbox_mode = "read-only"`
   (Codex) or have no edit/exec tools (Claude/Devin).
6. **Network egress** - is the network off unless the task needs it? In `workspace-write`
   it is off by default; `network_access = true` without a reason is **RISKY**.
7. **Cost bound** - there is no per-run cost cap flag, so cost must be bounded indirectly:
   a turn/rewind cap plus a tight sandbox plus approvals. If none of those exist, cost is
   unbounded - **MISSING**.

## Hard flags (override the verdict)

- `danger-full-access` **and** `approval_policy = never` **and** no allowlist -> **RISKY**,
  regardless of anything else. This is the "make it just work" config that does real damage.
- No give-up condition on an unattended loop -> at best **NEEDS WORK**.

## Output

Produce a report, in this shape:

```
GUARDRAIL AUDIT - <what was audited>
Verdict: SAFE TO RUN UNATTENDED | NEEDS WORK | RISKY

1. Sandbox scope       PRESENT  - workspace-write
2. Approvals           MISSING  - fix: set approval_policy = "on-request"
3. Give-up             MISSING  - fix: add "stop after N rewinds" to the prompt
4. Tool allowlist      PRESENT  - MCP default_tools_approval_mode = "writes"
5. Read-only review    RISKY    - reviewer can write; fix: sandbox_mode = "read-only"
6. Network egress      PRESENT  - off (workspace-write default)
7. Cost bound          MISSING  - fix: add a turn/rewind cap

Top fix: <the single most important thing to change first>
```

Rank the fixes; lead with the one that closes the biggest hole. If everything passes, say
so in one line and name what an attacker or a runaway would have to get past.

## What this is not

Not a code review, not a test runner. It only answers: *if you left this running, what
would stop it doing something dumb or expensive?*
