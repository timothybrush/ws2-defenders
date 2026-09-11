# OCSF ↔ CoSAI AITF: Agentic AI Telemetry Alignment

**Purpose:** Map the OCSF agentic-observability proposal (issue [#1640](https://github.com/ocsf/ocsf-schema/issues/1640) / PR [#1641](https://github.com/ocsf/ocsf-schema/pull/1641)) against the CoSAI AITF AI Telemetry Framework ([cosai-oasis/ws2-defenders/telemetry](https://github.com/cosai-oasis/ws2-defenders/tree/main/telemetry)), call out where they already agree, where they diverge, and propose a concrete path to a single shared foundation.

**Audience:** OCSF schema maintainers/code owners and CoSAI WS2 (Defenders) AITF contributors.

> **Status update, revised 2026-09-05 — this discussion is effectively over, and OCSF won the
> modelling argument.**
>
> **OCSF v1.9.0 shipped on 2026-08-03**, delivering the asks across four merged PRs rather than the
> single one this document tracked: `ai_agent` in
> [#1641](https://github.com/ocsf/ocsf-schema/pull/1641), `delegation` in
> [#1665](https://github.com/ocsf/ocsf-schema/pull/1665), `attestation` / `prev_event` /
> `record_integrity` in [#1661](https://github.com/ocsf/ocsf-schema/pull/1661), and `prompt_text` /
> `response_text` on `message_context` in
> [#1674](https://github.com/ocsf/ocsf-schema/pull/1674). The `ai_operation` profile was **extended**
> by the release, not introduced by it. **No `ai` category was created** — `categories.json` still
> stops at `uid 8` (*Unmanned Systems*) — and no `agent_activity` or `delegation_activity` class
> exists. OCSF answered the modelling question by declining it: agentic activity is a **profile over
> existing classes**.
>
> Two qualifications keep this honest. First, **issue #1640 is still open**: its object and profile
> asks are satisfied, but the placement of an *agent-layer summary event class* remains recorded as
> TBD between `ai`, `findings` and elsewhere — so OCSF has declined an `ai` category *so far*, rather
> than ruled one out. Second, "full stop" was too strong in the previous revision of this note, and
> the correction below on `message_context` shows why: parts of this discussion that were marked
> divergent have since been settled quietly by OCSF in ways this document did not anticipate.
>
> AITF has now retired its AI classes **twice**, and both retirements moved in the same direction.
> First it dropped its bespoke **Category 7** (`7001`–`7010`) and adopted OCSF's reuse-first model for
> data-plane activity, keeping a provisional `ai` category (`uid 9`, classes `9001`–`9003`) for
> control-plane lifecycle. v1.9.0 showed that reservation was itself a bet on a category that never
> came, so **`uid 9` and its three classes are retired too**: agent lifecycle and agent-to-agent
> messaging emit on **API Activity (6003)**, delegation lifecycle on **Authorize Session (3003)**,
> all carrying `ai_operation`. Every SDK exports a `LEGACY_AI_CLASS_UIDS` map so telemetry emitted
> against either retired numbering stays decodable.
>
> **Divergences D1 and D2 are therefore settled, and §4.1's "blocker" no longer exists.** The honest
> reading is that AITF's instinct on D1 — align to whatever number OCSF assigns — was right, but
> insufficiently sceptical: the correct answer was that there would be no number. §4.1's own advice to
> *"reference the category symbolically rather than by number"* would have avoided one of the two
> migrations, which is a lesson worth keeping rather than editing away.
>
> What remains genuinely open is §4.2, and it is now the whole of the AITF→OCSF ask: the released
> `delegation` object carries `uid`, `created_time`, `parent_uid` and `issuer_uid` — the *chain* — but
> not the *authority*. Scope, scope attenuation, proof type and TTL did not ship. See
> [`spec/ocsf-mapping/event-classes.md`](spec/ocsf-mapping/event-classes.md),
> [`ocsf-agentic-crosswalk.md`](spec/ocsf-mapping/ocsf-agentic-crosswalk.md), and the
> [OCSF PR plan](upstream-pr-plan-ocsf.md). §3's divergence table and §4's proposals are retained
> unedited as the record of the discussion; read them as history.

**TL;DR:** The two efforts are strongly complementary, not competing. OCSF #1640 is a *narrow, schema-first* proposal for agent identity + delegation + cross-event attribution. AITF is a *broad, producer-first* framework (OTel-native instrumentation → OCSF normalization) covering the full AI lifecycle. They overlap almost exactly on the identity/delegation/attribution core, and the main thing that must be reconciled now — before either ships — is the **AI category UID** (#1640 proposes `9`; AITF README documents `7`) and the **modeling philosophy** (profile-on-existing-classes vs. dedicated AI event classes).

---

## 1. What each side is proposing

### OCSF issue #1640 + PR #1641 (author: @Aniak5)

A schema-only extension to make autonomous agents and the authority they act under first-class in OCSF.

- **New** `ai_agent` **object** (landed in PR #1641): stable logical `uid`, restart-sensitive `instance_uid`, `name`, `type`/`type_id` (`Native`, `MCP`, `A2A`, `LangChain`, `AutoGen`, `CrewAI`, `Other`), backing `ai_model`, `version`. Deliberately **separate from the existing** `agent` **object** (which models EDR/DLP/APM sensors).
- **New** `delegation` **object** (in #1640, deferred from #1641): durable authorization context with stable `uid`, parent `uid` for lineage, trusted-issuer `uid`; persists independently of any trace/session/workflow. IDs minted at a **broker** rather than self-asserted, to enable binding-integrity checks.
- `delegation_lineage` **/** `delegation_node`: directed-graph representation of delegation ancestry for multi-agent compositions.
- `ai_operation` **profile extension**: lets data-plane classes carry `ai_agent` + `delegation` context. PR #1641 already adds the profile to `file_activity`, `network_activity`, `web_resources_activity`, `email_activity`, `script_activity`, `scheduled_job_activity` (joining `process_activity`, `datastore_activity`, `api_activity`). It also wires `ai_agent` into `actor.json` (added to the `at_least_one` constraint) and `process.json`.
- **New category** `ai` **(uid: 9)** with control-plane classes `delegation_activity` (Create/Revoke/Expire/Complete), `agent_activity` (Spawn/Terminate/Suspend/Resume/Update/Register), and a generic `ai` parent — all deferred to follow-up PRs.
- **Agent-layer summary class** (open question): one record for an atomic agent *outcome* that decomposes into the underlying atomic events sharing a delegation UID. Explicitly **not** proposed as a `finding` subclass; placement TBD.

**Posture:** minimal, additive, reuse-first. Attribution is threaded through *existing* event classes; only control-plane lifecycle gets net-new classes.

> **Outcome, verified 2026-09-05.** The six bullets above are the *proposal as written*, retained for
> the record. What actually shipped in v1.9.0 differs on four points. **(1)** `ai_agent.type_id`
> released as `0 Unknown / 1 Native / 2 LangChain / 3 AutoGen / 4 CrewAI / 99 Other` — the **`MCP`
> and `A2A` members did not ship**. **(2)** `delegation` was not deferred indefinitely; it landed
> via [#1665](https://github.com/ocsf/ocsf-schema/pull/1665) on 2026-07-24 with four attributes
> (`uid` required, `created_time`, `issuer_uid`, `parent_uid`) — the lineage, but none of the
> authority fields. **(3)** `delegation_lineage` / `delegation_node` **did not ship in any form**;
> ancestry is expressed only by `parent_uid` chaining. **(4)** The `ai` category and its
> control-plane classes **were not created**, and the agent-layer summary class remains the open
> question it was — still marked TBD on #1640, which is still open. So of the six bullets, two
> landed in reduced form, two did not land, and the last is unchanged after a year.

### CoSAI AITF — AI Telemetry Framework 

A security-first telemetry framework built on **OpenTelemetry GenAI + OCSF**, spanning instrumentation → collection → normalization → analytics.

- **Dual-pipeline**: single instrumentation pass emits OTLP (Jaeger/Tempo/Datadog) *and/or* **OCSF Category 7** events (Splunk/Security Lake/QRadar/Sentinel).
- **Ten OCSF Category-7 AI event classes**: `7001` Model Inference, `7002` Agent Activity (lifecycle, reasoning, delegation), `7003` Tool Execution (incl. MCP), `7004` Data Retrieval (RAG), `7005` Security Finding, `7006` Supply Chain, `7007` Governance, `7008` Identity (identity/auth/authz/delegation/trust), `7009` Model Operations, `7010` Asset Inventory.
- **Semantic-convention namespaces**: `gen_ai.*` (OTel), plus AITF-defined `mcp.*`, `skill.*`, `rag.*`, `security.*`, `compliance.*`, `cost.*`, `quality.*`, `supply_chain.*`, `identity.*`, `model_ops.*`, and CoSAI-aligned `asset.*`, `drift.*`, `memory.security.*`.
- **Identity model**: `identity.*` with OAuth 2.1, SPIFFE, DID-VC trust methods; `scope_requested` vs `scope_granted`; credential lifecycle (create/rotate/revoke); TTL.
- **Enrichment + detection**: OWASP LLM/Agentic/MCP Top-10 mapping, 14 detection rules (e.g., `AITF-DET-006` Unauthorized Agent Delegation, `AITF-DET-007` Agent Session Hijack), Sigma/SPL, anomaly engine, 8 compliance-framework mappings.

**Posture:** comprehensive, producer-first. OCSF is one normalization target among several; AITF owns instrumentation and enrichment, with dedicated AI-native event classes.

---

## 2. What we have in common

| Theme | OCSF #1640 / #1641 | CoSAI AITF | Verdict |
| --- | --- | --- | --- |
| **Agent as a first-class entity, distinct from sensors** | New `ai_agent` object, explicitly *not* the `agent` (EDR/DLP) object | `gen_ai.agent.*` + `7002 AI Agent Activity` + `identity.*` | ✅ Strong agreement |
| **Framework/protocol typing** | `ai_agent.type_id`: Native/MCP/A2A/LangChain/AutoGen/CrewAI/Other | Per-framework vendor mappings (LangChain, CrewAI, …) + provider detection | ✅ Same intent |
| **Delegation as core concept** | `delegation` object + `delegation_activity` lifecycle + lineage graph | `identity.*` delegation chains, `7008 AI Identity`, `AITF-DET-006` | ✅ Both treat delegation as central |
| **Cross-cutting attribution of agent actions to atomic events** | `ai_operation` profile on file/network/email/process/api/etc. | OTel spans → OCSF mapping; `gen_ai.tool.*`, action logging | ✅ Same goal, different mechanism |
| **Control-plane vs data-plane separation** | `delegation_activity` / `agent_activity` (control) vs profiled data-plane events | `7002`/`7008` (lifecycle/identity) vs `7001/7003/7004` (actions) | ✅ Aligned conceptually |
| **Summary-over-atomic-events** | Proposed agent-layer summary class, linked by shared delegation UID | Agent session/span tree (`gen_ai.agent.session`, steps) | ✅ Same idea, two encodings |
| **OCSF as canonical at-rest schema** | OCSF *is* the deliverable | OCSF export is a primary target; now built on **OCSF v1.9.0** (was v1.1.0 when this was written, with a bespoke Category 7) | ✅ Both anchor on OCSF |
| **MCP / A2A as named primitives** | `type_id` enum includes MCP, A2A | `mcp.*` namespace, MCP Top-10 coverage, A2A trust | ✅ Both name them explicitly |

The identity/delegation/attribution **core is essentially the same design** arrived at independently — that's the strongest possible signal for convergence.

---

## 3. Where we diverge

| \# | Dimension | OCSF #1640 / #1641 | CoSAI AITF | Why it matters |
| --- | --- | --- | --- | --- |
| D1 | **AI category UID** | New category `ai` = **uid 9** | Documents AI classes under **Category 7** (`7001`–`7010`) | Direct numbering conflict. Must be resolved before either is referenced in production mappings, or the two will produce incompatible `class_uid`s. |
| D2 | **Modeling philosophy** | *Profile-on-existing-classes* + a few control-plane classes. Reuse-first, minimal new taxonomy. | *Dedicated AI-native event classes* (10 of them) covering the whole domain. | Determines whether an agent file-write is a `file_activity` w/ `ai_operation` profile (OCSF) or a `7003`/`7002` AI class (AITF). Affects every consumer query. |
| D3 | **Scope** | Tightly scoped: identity + delegation + lineage + attribution. | Full lifecycle: inference, tool/MCP, RAG, security, supply chain, governance, identity, model-ops, asset inventory, cost, quality. | OCSF lands a small core; AITF needs homes for \~7 domains OCSF #1640 doesn't touch yet. |
| D4 | **Production model** | Schema-only; agnostic to how events are emitted. | OTel GenAI-native, dual OTLP+OCSF pipeline; OCSF is one of several exporters. | Defines the seam: who produces vs. who standardizes the at-rest record. |
| D5 | **Delegation representation** | Durable standalone `delegation` object + explicit `delegation_lineage` graph, broker-minted IDs, decoupled from trace. | Delegation rides OTel trace/span + identity spans + chains; detection-oriented. | OCSF wants a queryable durable authority object; AITF currently expresses it as trace context. Need a canonical durable form. |
| D6 | **Identity/trust depth** | `ai_agent` + `delegation` (uid, parent, issuer). | Richer auth/trust: OAuth 2.1, SPIFFE, DID-VC, scope_requested/granted, credential lifecycle, TTL. | AITF has fields OCSF's objects don't model yet — a contribution opportunity, not a conflict. |
| D7 | **Governance / detection / compliance** | Out of scope (schema only). | First-class: OWASP Top-10 mapping, 14 detection rules, Sigma/SPL, 8 compliance frameworks. | Need to decide what belongs in the OCSF schema vs. the AITF producing/enrichment layer. |
| D8 | **Agent-layer summary placement** | Open question (under `ai`? `findings`? new?), not a `finding` subclass. | Implicit as the agent-session span; surfaces as `7002`. | Both must agree on one canonical "agent outcome" record and where it lives. |

---

## 4. Proposal: strong common ground

A layered split that lets each effort do what it does best, with OCSF as the single normalized schema and AITF as the reference producer/enrichment layer.

### 4.1 Resolve the category UID first (blocker — D1)

> **Outcome:** ~~blocker~~ — **resolved by OCSF v1.9.0, in a way neither side proposed.** No `ai`
> category was created. The paragraph below is retained as written; its closing recommendation
> ("reference the category symbolically rather than by number") was the durable part.

Pick **one** AI category UID and use it everywhere. Recommendation: **AITF aligns to the OCSF-assigned UID** (currently proposed as `9` in #1640), treating its documented "Category 7" as a pre-standardization placeholder. Action: OCSF confirms/reserves the AI category UID in #1640; AITF updates its `OCSFMapper`, `event-classes.md`, and schema constants to match. Until confirmed, both sides should reference the category symbolically (`ocsf.category.ai`) rather than by number.

### 4.2 Adopt OCSF objects as the canonical identity/authority vocabulary (D5, D6)

- Treat OCSF `ai_agent` and `delegation` (+ `delegation_lineage`) as the **canonical at-rest representation**. AITF maps `gen_ai.agent.*` / `identity.*` onto them.
- Contribute AITF's richer fields **back into OCSF** so nothing is lost: add to `ai_agent`/`delegation` (or a `trust`/`auth` sub-object) the trust method (`mTLS`/`SPIFFE`/`DID-VC`/`OAuth2.1`), `scope_requested` vs `scope_granted`, credential lifecycle state, and `ttl`. This is the single highest-value cross-pollination: OCSF gains depth, AITF gains canonical objects.

### 4.3 Adopt a hybrid event model (D2, D3) — the key compromise

> **Outcome:** the hybrid was not adopted. **v1.9.0 chose the profile side wholesale**, including for
> control-plane lifecycle, so the third bullet below no longer has a referent — there is no
> `delegation_activity` or `agent_activity` class. AITF follows: 6003 for agent lifecycle and
> agent-to-agent messaging, 3003 for delegation lifecycle. The second bullet's "principled rule for
> what earns a new class" survives as the open question, and the strongest remaining candidate for it
> is **model lifecycle (train / evaluate / deploy / drift)**, which no existing class can express.

Stop framing it as "profile vs. dedicated classes." Use both, by rule:

- **Use the** `ai_operation` **profile** (OCSF approach) for actions that already have a natural OCSF home — file, network, email, process, scheduled-job, script, api, datastore. An agent-driven file write stays a `file_activity` *plus* `ai_agent`/`delegation` context.
- **Use dedicated AI classes** (AITF approach) only for AI-native concepts with **no existing OCSF home** — model inference, tool/MCP execution semantics, retrieval/RAG, model-ops, asset inventory. These become net-new classes under the agreed `ai` category.
- **Control-plane lifecycle** stays as #1640's `delegation_activity` / `agent_activity`.

This collapses AITF's 10 classes into "map to existing class + profile" vs. "promote to new `ai` class," and gives OCSF a principled rule for what earns a new class.

### 4.4 Joint mapping table (proposed)

| AITF Category-7 class | Proposed OCSF home |
| --- | --- |
| 7001 AI Model Inference | New `ai` class (no existing home) |
| 7002 AI Agent Activity | `agent_activity` (control-plane) + agent-layer summary class |
| 7003 AI Tool Execution (MCP) | New `ai` class for tool/MCP; underlying side-effects map to existing classes + `ai_operation` profile |
| 7004 AI Data Retrieval (RAG) | New `ai` class (or `datastore_activity` + profile where it fits) |
| 7005 AI Security Finding | Existing OCSF `finding` / `detection_finding` (Category Findings) |
| 7006 AI Supply Chain | New `ai` class or extend existing supply-chain/inventory modeling |
| 7007 AI Governance | Enrichment/profile, not necessarily a new class |
| 7008 AI Identity | `delegation` + `ai_agent` objects + `delegation_activity` |
| 7009 AI Model Operations | New `ai` class |
| 7010 AI Asset Inventory | Align with OCSF Discovery/Inventory classes |

(Table is a starting point for the working session, not a final assignment.)

### 4.5 Agent-layer summary record (D8)

> **Outcome: still open, and now the cleanest joint deliverable left.** v1.9.0 shipped no summary
> class. With no `ai` category, the recommendation below to place it there is void; the natural home
> is Application Activity (`6`). The decomposition mechanism the paragraph proposes — a shared
> `delegation.uid` — did ship, so the hard part is available.

Agree on one canonical "agent outcome" record that decomposes into constituent atomic events via a **shared** `delegation.uid` (and/or agent session id). AITF's agent-session span is the producer; OCSF's proposed summary class is the at-rest form. Recommend it live under the `ai` category (not under `findings`), consistent with #1640's instinct.

### 4.6 Define the layer seam (D4, D7)

- **CoSAI AITF = producer + enrichment layer**: OTel-native instrumentation, detection rules, OWASP/compliance mapping, dual-pipeline export. Owns "how telemetry is generated and enriched."
- **OCSF = canonical schema layer**: the normalized at-rest record. Owns "what the event looks like once stored."
- Governance/compliance/OWASP tags ride as OCSF **enrichment/profile attributes** (e.g., a `compliance` or `security_control` profile) rather than bloating core classes — so SIEMs get the mappings without OCSF having to standardize every framework.

### 4.7 Process / next steps

1. ~~**Joint working session** between #1640 authors and AITF/WS2 to ratify the category UID and the §4.3 hybrid rule.~~ **Moot** — no category, and the hybrid was decided in favour of the profile.
2. ~~**#1641 proceeds as-is**~~ — **done.** It merged and shipped in v1.9.0 with the `ai_agent` object; `delegation` followed in #1665. AITF consumes both directly.
3. **AITF submits a follow-up OCSF PR** contributing the trust/auth/scope/TTL fields (§4.2) into the released `delegation` / `ai_agent` objects. **This is now the single most important item on the list**, because v1.9.0 shipped the delegation *chain* without the delegation *authority*, and scope attenuation is what makes a confused-deputy attack visible. Drafted in [`spec/ocsf-mapping/`](spec/ocsf-mapping/) and sequenced in the [OCSF PR plan](upstream-pr-plan-ocsf.md).
4. **Publish the mapping table (§4.4)** as a shared doc in both repos so consumers have one source of truth. Superseded in practice by [`ocsf-agentic-crosswalk.md`](spec/ocsf-mapping/ocsf-agentic-crosswalk.md), which maps AITF onto *released* v1.9.0 objects rather than onto proposals.
5. **Co-author the agent-layer summary class** (§4.5) — still open, and still worth doing.
6. **New: ask for the `sequence_number` and gap markers on `attestation`.** v1.9.0's `record_integrity` profile proves a log is *unaltered*; nothing in it proves the log is *complete*. That distinction is what EU AI Act Art. 12 record-keeping turns on.

---

## 5. One-paragraph summary for the thread

> **Post this version (2026-09-05).** OCSF v1.9.0 settled the two things this thread was convened to
> reconcile, so the original paragraph below is no longer postable. Replacement:
>
> *OCSF v1.9.0 released the `ai_agent` and `delegation` objects and the `ai_operation` profile, and
> chose to model agentic activity as a profile over existing classes rather than as a new category.
> CoSAI AITF has aligned to that decision completely: it has retired both its bespoke Category 7 and
> the provisional `ai` category it had reserved, and now emits agent lifecycle and agent-to-agent
> messaging on API Activity (6003), delegation lifecycle on Authorize Session (3003), and data-plane
> AI activity on the classes that already fit — all carrying `ai_operation`. Legacy class UIDs remain
> decodable through an exported map in each SDK. Two asks remain, and both are narrow additions to
> objects v1.9.0 already shipped rather than new taxonomy. First, the released `delegation` object
> carries the delegation chain but not the delegated authority: adding `scope`, `proof_type` and a TTL
> is what makes scope attenuation — and therefore confused-deputy abuse — visible at all. Second, the
> `record_integrity` profile's hash chain proves a log has not been altered but cannot prove it is
> complete; a monotonic `sequence_number` and gap markers on `attestation` close that, which matters
> for anyone relying on OCSF for EU AI Act Art. 12 record-keeping. AITF remains the OTel-native
> producer and enrichment layer above OCSF's at-rest schema, and both asks are drafted as filable
> PRs.*

**Superseded original (retained as the record):**

OCSF #1640/#1641 and CoSAI AITF independently converged on the same core: agents as first-class entities distinct from security sensors, delegation as the unit of authority, and cross-event attribution of agent-driven actions. They diverge mainly on (1) the AI category UID — #1640 says `9`, AITF documents `7` — which must be reconciled immediately, and (2) modeling philosophy — OCSF threads attribution through existing classes via the `ai_operation` profile while AITF defines dedicated AI classes. The recommended common ground is a hybrid: keep #1640's `ai_agent`/`delegation`/profile core as the canonical schema, promote only genuinely AI-native concepts (inference, tool/MCP, RAG, model-ops) to new classes under one agreed `ai` category, contribute AITF's richer trust/auth/scope fields back into OCSF's delegation objects, and treat AITF as the OTel-native producer/enrichment layer above OCSF's at-rest schema.

---

### Sources

- OCSF — Add agentic AI observability to the schema (issue #1640): https://github.com/ocsf/ocsf-schema/issues/1640
- OCSF — Add `ai_agent` object and extend `ai_operation` profile coverage (PR #1641): https://github.com/ocsf/ocsf-schema/pull/1641
- CoSAI AITF — AI Telemetry Framework: https://github.com/cosai-oasis/ws2-defenders/tree/main/telemetry