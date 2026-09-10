# OCSF Issue — ready to post

> **Rebase note, 2026-09-05.** Written while #1641 was open. **It merged and shipped in OCSF v1.9.0
> (2026-08-03)**, so this issue now proposes extending a *released* profile rather than a pending one.
> The `ai` category proposed alongside it in #1640 was **not** created — `categories.json` stops at
> `uid 8` — so the two follow-up bullets that assumed it have been rewritten. The core ask is
> unchanged and, if anything, easier to accept now.

**Title:**

```
[Proposal] Extend the ai_operation profile with inference, tool, and finding telemetry
```

---

**Body:**

### Summary

Following the `ai_agent` object (#1641) and `delegation` object (#1665) released
in **v1.9.0**, and the `ai_operation` profile that release extended, I'd like to propose adding the
**substance of an AI operation** to that profile: token usage, cost, latency,
request parameters, tool/MCP invocation, and AI-specific finding context. These
are **optional objects and profile attributes on existing classes** — no new
category, no new required fields, and consistent with the profile-over-classes
approach v1.9.0 chose.

Opening this as a discussion to agree on placement and naming before I submit
the schema PR.

### Gap

v1.9.0 captures *who* acted (`ai_agent`) and *under what authority*
(`delegation`), carried on `ai_operation` across the `system`, `network`,
`application` and `iam` base classes. It does not capture what the operation
consumed or did:

| Gap (not in OCSF today) | Why it matters |
|---|---|
| **Token usage** | quota/abuse detection, cost attribution |
| **Cost** | FinOps, runaway-spend / exfiltration-by-volume detection |
| **Latency** (TTFT, tokens/sec, queue time) | DoS / "Unbounded Consumption" (OWASP LLM10) detection, SLOs |
| **Request parameters** (temperature, top_p, max_tokens) | reproducibility, jailbreak-tuning detection, governance |
| **Tool / function / MCP invocation** | the agent's *action* — where tool-abuse, RCE, and data-access detections attach |
| **AI finding context** (OWASP LLM Top 10, guardrail, PII) | triage of prompt injection / leakage, and distinguishing an AI Detection Finding from a non-AI one on the same class |

### Proposed (high level)

Riding on existing classes (API Activity `6003`, Datastore Activity `6005`,
Detection Finding `2004`):

1. **`ai_token_usage` object** — `input_tokens`, `output_tokens`,
   `cached_tokens`, `reasoning_tokens`, `total_tokens`.
2. **`ai_cost` object** — `input_cost`, `output_cost`, `total_cost`, `currency`.
3. **`ai_request_parameters` object** — `temperature`, `top_p`, `max_tokens`.
4. **`ai_latency` object** — `total_time_ms`, `time_to_first_token_ms`,
   `tokens_per_second`, `queue_time_ms`.
5. **`ai_tool` object** — `name`, `type_id` (Function / MCP Tool / Skill / API
   / Other), `server`, `transport`, `approval_required`, `is_approved`.
6. **`ai_operation` profile additions** — `token_usage`, `cost`, `latency`,
   `request_parameters`, `ai_tool`, `finish_reason`, `is_streaming`, alongside
   the four attributes already released (`ai_agent`, `ai_model`, `delegation`,
   `message_context`).
7. **`ai_finding` profile** — `owasp_llm_id` (OWASP LLM Top 10 2025),
   `detection_method`, `guardrail_name`, `is_blocked`, `pii_types`.

### Out of scope (potential follow-ups)

Each closes a gap where an AI event reuses an OCSF class that lacks the fields:

- **RAG / vector retrieval** (`ai_retrieval` object) on Datastore Activity
  `6005` — top_k, similarity scores, embedding model/dims, reranking,
  retrieved-chunk scores, pipeline stage.
- **AI supply-chain / provenance** (model hash/signature/signer, AI-BOM) on
  Vulnerability Finding `2002`.
- **MLOps lifecycle** — a new `ai_model_activity` class (training / evaluation
  / deployment / drift), which Application Lifecycle `6002` cannot express.
  Absent an `ai` category, the natural home is Application Activity (`6`).
- **Agentic identity enrichment** on Authentication `3002` — SPIFFE/DID/DPoP
  auth methods, trust establishment, scope request/grant/attenuation.
- **MITRE ATLAS** technique mapping on findings (OCSF `attacks` is ATT&CK only).
- **EU AI Act risk classification** (`risk_level` enum) on Compliance Finding
  `2003` / asset inventory.
- **Agent reasoning & multi-agent** objects (`ai_reasoning`, `ai_team`) on the
  `ai_operation` profile; **AI quality metrics** (`ai_quality`) likewise.
- **Agent-to-agent communication** — ONE generic `agent_message` object with a
  `protocol_id` discriminator (A2A / ACP / ANP / MCP / Other), **not** a
  per-protocol object. Same conceptual core for every protocol (peer agents +
  delegation, unit-of-work + canonical lifecycle status, operation, transport,
  trust/DID). Per-protocol objects would fragment cross-protocol detection and
  chase fast schema churn; mirror OCSF's "generic class + protocol id" pattern
  (`network_activity` + `tls`/`dns_query`). Proposed as an object on
  `ai_operation`, carried on API Activity `6003` — complementary to the
  `message_context` attribute v1.9.0 shipped, which correlates a conversation
  but does not describe the peer or the authority.

### Backwards compatibility

All additions are optional; nothing existing changes; no category is added.

### Prior art

Drawn from the AI Telemetry Framework (AITF), which already implements these as
OCSF-mapped output. I have a complete schema PR (object JSON, profile diffs,
`dictionary.json` entries) ready to open once placement is agreed.

Two questions for maintainers. Would you prefer these as `ai_operation`
extensions, or as a separate profile — the finding-context group in particular
may belong in its own `ai_finding` profile since it applies to a different class
family. And now that v1.9.0 has settled the shape for agentic telemetry, is
there a documented position on what would justify a future `ai` category, so
the MLOps-lifecycle follow-up can be filed against it rather than guessing?
