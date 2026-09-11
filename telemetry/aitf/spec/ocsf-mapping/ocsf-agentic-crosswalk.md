# AITF ↔ OCSF Agentic AI Alignment

> **Status:** Aligned to **OCSF v1.9.0** (released 2026-08-03), which landed
> the objects across four merged PRs — `ai_agent`
> ([#1641](https://github.com/ocsf/ocsf-schema/pull/1641)), `delegation`
> ([#1665](https://github.com/ocsf/ocsf-schema/pull/1665)), `attestation` /
> `prev_event` / `record_integrity`
> ([#1661](https://github.com/ocsf/ocsf-schema/pull/1661)) and `prompt_text` /
> `response_text` on `message_context`
> ([#1674](https://github.com/ocsf/ocsf-schema/pull/1674)) — and **extended**
> the pre-existing `ai_operation` profile rather than introducing it.
> [Issue #1640](https://github.com/ocsf/ocsf-schema/issues/1640) is **still
> open**: its object and profile asks are satisfied, but the placement of an
> agent-layer summary event class is recorded there as TBD. Sections marked
> *AITF extension* remain the open upstream ask.

This document describes how AITF maps onto OCSF's agentic-AI model. The
guiding principle, from OCSF itself, is **reuse existing objects and profiles**
rather than minting bespoke AI event classes:

- **All AI activity** — inference, tool/function calls, retrieval, model
  lifecycle, findings, identity, inventory, *and* agent/delegation/agent-comm
  lifecycle — is emitted under the **existing OCSF event class** it naturally
  belongs to, **enriched with the `ai_operation` profile** (the `ai_agent`
  object + `ai_model`) and, where relevant, the `delegation` object.
- There is **no AITF-specific category or class**. Every `category_uid` and
  `class_uid` AITF emits is a released OCSF one.

**What v1.9.0 ratified.** OCSF v1.9.0 landed the `ai_agent` and `delegation`
objects and the `ai_operation` profile, and — per the release notes — attached
that profile to the `system`, `network`, `application` and `iam` base event
classes, so all System Activity, Network Activity, Application Activity and
IAM events inherit agent attribution. It also added `ai_agent` and
`hosted_ai_agent_list` to the `process` object.

**Two rounds of AITF retirement.** AITF first defined a bespoke **Category 7**
(classes 7001–7010); that collided with released OCSF (`uid 7` =
*Remediation*, `uid 8` = *Unmanned Systems*) and ran against OCSF's reuse
philosophy, so Category 7 was dropped. AITF then proposed an **`ai` category
(`uid 9`)** with provisional classes `9001`–`9003` for control-plane
lifecycle. That category was **never ratified** — v1.9.0 `categories.json`
still stops at `uid 8` — so those three classes are **retired** too, and their
lifecycle events now ride released classes (§1). Telemetry recorded by
AITF ≤ 0.4 stays decodable through the `LEGACY_AI_CLASS_UIDS` table each SDK
ships.

## 1. Class mapping (what AITF emits)

| AITF event | OCSF category | OCSF class | `class_uid` |
|---|---|---|---|
| Model Inference | 6 Application Activity | API Activity | 6003 |
| Tool / MCP / Function Execution | 6 Application Activity | API Activity | 6003 |
| Data Retrieval (RAG / vector) | 6 Application Activity | Datastore Activity | 6005 |
| Model Operations (train/eval/deploy) | 6 Application Activity | Application Lifecycle | 6002 |
| Security Finding | 2 Findings | Detection Finding | 2004 |
| Supply Chain | 2 Findings | Vulnerability Finding | 2002 |
| Governance / Compliance | 2 Findings | Compliance Finding | 2003 |
| Identity / Authentication | 3 IAM | Authentication | 3002 |
| Asset Inventory | 5 Discovery | Inventory Info | 5001 |
| **Agent lifecycle** | **6 Application Activity** | **API Activity** | **6003\*** |
| **Delegation lifecycle** | **3 IAM** | **Authorize Session** | **3003\*** |
| **Agent-to-agent comms (A2A/ACP/ANP)** | **6 Application Activity** | **API Activity** | **6003\*** |

\* Retired from the unratified `ai` category (was `9001`/`9002`/`9003`). Agent
lifecycle and agent-to-agent communication are request/response operations
between software components, so **API Activity** is their released home;
agent session start/stop MAY instead use **Application Lifecycle (6002)** where
that reads more naturally. Delegation lifecycle assigns privileges to a new
principal, which is exactly what OCSF **Authorize Session (3003)** describes
("privileges, groups or roles assigned to a new user session"), and IAM is one
of the four categories v1.9.0 attached `ai_operation` to.

Inference, tool execution, agent lifecycle and agent-to-agent communication all
share API Activity (6003); they are distinguished by `activity_id`, by
`ai_agent.type_id`, and by the AITF attributes on the event.

> **Implication for detection content.** Because AI events reuse shared OCSF
> classes, downstream rules/dashboards must filter on **`class_uid` *plus* the
> presence of the `ai_operation` profile** (e.g. `ai_agent` is set), not on a
> dedicated AI class UID. A non-AI Detection Finding and an AI Detection
> Finding share `class_uid 2004`; the `ai_agent`/`ai_model`/`delegation`
> fields are what mark the latter as AI.

## 2. The `ai_agent` object (released, OCSF v1.9.0)

Carried on the `ai_operation` profile of every agent-attributable event.
Distinct from the OCSF `agent` object (security sensors). All eight attributes
below shipped in v1.9.0 — AITF has **no outstanding ask** on this object.

| OCSF `ai_agent` field | Req. | AITF source attribute | Notes |
|---|---|---|---|
| `uid` | Required | `gen_ai.agent.id` ∥ `identity.agent_id` ∥ `agent.workflow_id` | Stable logical agent id |
| `instance_uid` | Recommended | `gen_ai.conversation.id` | Restart-sensitive running instance |
| `name` | Recommended | `gen_ai.agent.name` ∥ `identity.agent_name` | |
| `type` | Optional | caption of `type_id` | e.g. `LangChain` |
| `type_id` | Recommended | normalized from `agent.framework` | enum below |
| `ai_model` | Recommended | `gen_ai.request.model` ∥ `gen_ai.response.model` | Backing model |
| `version` | Recommended | `gen_ai.agent.version` | Agent code/config revision |
| `charter` | Optional | `gen_ai.agent.description` | Role / operating boundaries |

### `type_id` enum (framework normalization)

Matches the released v1.9.0 enum exactly. Frameworks without a dedicated OCSF
member normalize to `Other (99)`.

| `type_id` | Caption | AITF `agent.framework` values |
|---|---|---|
| 0 | Unknown | *(absent)* |
| 1 | Native | `native` |
| 2 | LangChain | `langchain`, `langgraph` |
| 3 | AutoGen | `autogen` |
| 4 | CrewAI | `crewai` |
| 99 | Other | `semantic_kernel`, `custom`, … |

> **`MCP` and `A2A` members did not ship.** They were discussed on PR #1641 but
> are absent from released `ai_agent.type_id`. AITF normalizes MCP-hosted and
> A2A-participating agents to `Other (99)` and preserves the precise framework
> in `agent.framework`. Adding those two members is part of AITF's remaining
> upstream ask.

## 3. The `ai_operation` profile (released, OCSF v1.9.0)

The profile carries four attributes — `ai_agent`, `ai_model`, `delegation` and
`message_context` — and v1.9.0 attached it to the `system`, `network`,
`application` and `iam` base classes. AITF applies it to **every** event it
emits. Profile fields on the AITF base event:

| Field | Type | Status | Source |
|---|---|---|---|
| `ai_agent` | `ai_agent` object | released | §2 |
| `ai_model` | string | released | `gen_ai.request.model` |
| `delegation` | `delegation` object | released | §4 |
| `message_context` | object | released | A2A/ACP message envelope (see [a2a-spans](../semantic-conventions/a2a-spans.md)) |
| `delegation_lineage` | `delegation_lineage` object | **AITF extension** | §4 |

## 4. The `delegation` object & lineage

### 4.1 Released attributes (OCSF v1.9.0)

The released `delegation` object carries four attributes. AITF populates all of
them:

| OCSF `delegation` field | Req. | AITF source attribute |
|---|---|---|
| `uid` | Required | `identity.delegation.delegatee_id` ∥ `agent.delegation.target_agent_id` ∥ last node of chain |
| `created_time` | Optional | start time of the delegation-creation span; or derived as `identity.delegation.expires_at − identity.delegation.ttl_seconds` |
| `parent_uid` | Optional | `identity.delegation.delegator_id` |
| `issuer_uid` | Optional | `identity.provider` |

### 4.2 AITF extension (remaining upstream ask)

The released object records *that* a delegation exists and its lineage, but not
*what was delegated* or *for how long*. AITF keeps these fields on the
`delegation` object as an extension; they are the substance of the outstanding
OCSF request, because scope attenuation and TTL are what make confused-deputy
and privilege-escalation detection possible:

| AITF field on `delegation` | AITF source attribute | Why it matters |
|---|---|---|
| `delegator` | `identity.delegation.delegator` | Human-readable principal, not just a uid |
| `delegatee` | `identity.delegation.delegatee` | Human-readable target principal |
| `type` | `identity.delegation.type` | `oauth_token_exchange`, `capability_grant`, … |
| `scope` | `identity.delegation.scope_delegated` | Permissions actually passed on |
| `proof_type` | `identity.delegation.proof_type` | `jwt`, `dpop`, `mtls`, … — how the grant is provable |
| `ttl_seconds` | `identity.delegation.ttl_seconds` | Bounds the blast radius of a leaked grant |

`delegation_lineage` (also an AITF extension) is materialized from
`identity.delegation.chain` (ordered origin→current); each entry becomes a
`delegation_node` with `uid`, `parent_uid`, `agent_uid`, and `depth`. The
released `delegation.parent_uid` expresses one hop of the same graph, so a
consumer that only understands released OCSF still gets correct single-hop
ancestry; `delegation_lineage` is what makes *depth-bounded* chain queries and
gap detection possible in one pass.

## 5. Control-plane activities

The lifecycle events now ride released classes (§1). The class envelope
changed; the activity-name mapping below is unchanged and remains the
authoritative semantics.

### Agent lifecycle → API Activity (6003)

| AITF `activity_id` | AITF activity | Lifecycle semantics |
|---|---|---|
| 1 | Session Start | Spawn |
| 2 | Session End | Terminate |
| 3 | Step Execute | Update |
| 4 | Delegation | Register |
| 5 | Memory Access | Resume |
| 6 | Error Recovery | Resume |
| 7 | Human Approval | Suspend |
| 99 | Other | Unknown |

Because agent lifecycle now rides API Activity, these activity names are AITF
semantics carried in the event body; the OCSF `activity_id` on the emitted
event follows API Activity's own enum.

### Delegation lifecycle → Authorize Session (3003)

| AITF delegation op | Lifecycle semantics |
|---|---|
| `create` / `grant` | Create |
| `revoke` | Revoke |
| `expire` | Expire |
| `complete` | Complete |

## 6. `hosted_ai_agent_list` on `process` (released, OCSF v1.9.0)

v1.9.0 added both `ai_agent` and `hosted_ai_agent_list` to the OCSF `process`
object, covering the case where a single OS process hosts several agents and
per-agent attribution is otherwise impossible. When AITF emits an event with
host/process context **and** an `ai_agent`, exporters that populate the OCSF
`process` object SHOULD append the `ai_agent` to
`process.hosted_ai_agent_list`.

## 7. Event-sequence continuity → the `record_integrity` profile

OCSF v1.9.0 also added the **`record_integrity`** profile, carrying an
`attestation_list` of `attestation` objects, and the **`prev_event`** object.
That is a direct home for most of AITF's `observability.sequence.*` group
([security-cross-cutting.md § Event Sequence Continuity](../semantic-conventions/security-cross-cutting.md#event-sequence-continuity)),
which builds a hash-chained, tamper-evident event log. AITF emits the profile
whenever `observability.sequence.hash` is present:

| AITF attribute | OCSF `attestation` field |
|---|---|
| `observability.sequence.hash` | `fingerprint` — canonical serialization of this event |
| `observability.sequence.prev_hash` | `prev_event.fingerprint` |
| *(previous event's `metadata.uid`)* | `prev_event.uid` (Required) |
| *(previous event's `type_uid`)* | `prev_event.type_uid` |
| `observability.sequence.scope_id` | `chain_uid` — the append-only chain this event belongs to |
| `observability.sequence.signature` + `signer_key_id` | one entry in `signatures` |
| `identity.provider` ∥ signing authority | `authority_uid` |

`attestation` requires at least one of `fingerprint` or `signatures`, which
AITF satisfies whenever the sequence group is emitted at all.

### Remaining gap (AITF extension)

Five attributes have no released OCSF home and stay on the AITF event:

| AITF attribute | Why OCSF has no equivalent |
|---|---|
| `observability.sequence.number` | `attestation` chains by *hash*, not by ordinal. A hash chain proves *tampering*; a monotonic counter is what proves *absence* — you cannot detect a dropped event from hashes alone if the dropper also re-links the chain |
| `observability.sequence.scope_type` | Disambiguates whether the counter is per session, run, agent instance or emitter |
| `observability.sequence.gap_detected` | Consumer-side observation that a number was skipped |
| `observability.sequence.gap_size` | How many events are missing |
| `observability.sequence.reordered` | Out-of-order arrival, distinct from loss |

Adding a sequence ordinal and gap markers to `attestation` (or to `prev_event`)
is the second half of AITF's remaining upstream ask, alongside the delegation
scope/TTL fields in §4.2. EU AI Act Art. 12 record-keeping requires
demonstrating that a log is *complete*, not merely unaltered.

## 8. Reference implementation

| SDK | Objects, profile & class mapping | Builders & crosswalk tables |
|---|---|---|
| Python | `src/aitf/ocsf/schema.py`, `event_classes.py` | `src/aitf/ocsf/crosswalk.py` |
| TypeScript | `src/ocsf/schema.ts`, `event-classes.ts` | `src/ocsf/crosswalk.ts` |
| Go | `ocsf/schema.go`, `events.go` | `ocsf/crosswalk.go` |
| Rust | `src/ocsf/schema.rs`, `mapper.rs` | `src/ocsf/crosswalk.rs` |

Each event class declares its reused OCSF `category_uid` + `class_uid`, and the
`OCSFMapper` enriches every event with the `ai_operation` profile
(`ai_agent`, `ai_model`, `delegation`, `delegation_lineage`) automatically.
The authoritative AITF→OCSF class table lives in `OCSF_CLASS_CROSSWALK`; every
SDK's test suite asserts that no row targets `category_uid ≥ 9` or
`class_uid ≥ 9000`, so the retirement cannot silently regress. The legacy
`9001`/`9002`/`9003` → released-class decode table is exported as
`LEGACY_AI_CLASS_UIDS` in all four SDKs.
