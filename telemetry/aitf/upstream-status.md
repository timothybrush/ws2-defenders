# Upstream status: OpenTelemetry GenAI and OCSF

What the two upstream standards cover today, what is still missing, and which AITF namespace carries
each missing field in the interim.

Companion to **Secure by Design: Telemetry Field Selection for AI Systems**, Working Draft v0.4
(<https://github.com/cosai-oasis/ws2-defenders/blob/main/telemetry/CoSAI-AI-Telemetry-RFC.md>), and to
the [Appendix C resolution log](AITF_gaps.md).

> **Source constraint, revised 2026-09-05.** Coverage claims here originate in **RFC v0.4 Appendix D**
> (OpenTelemetry, the instrumentation bridge) and **Appendix E** (OCSF, the standardization bridge).
> That origin still holds for **Part 1**: the OpenTelemetry registries were **not** re-read, and
> Appendix D's own verification note (reproduced below) applies with full force — **re-verify against
> the live registry before filing anything upstream.**
>
> **Part 2 is different, and the difference is load-bearing.** The live `ocsf/ocsf-schema` repository
> **has** now been re-read. **OCSF v1.9.0 shipped on 2026-08-03**, delivering — per its own
> `CHANGELOG.md` — the `ai_agent` object
> ([#1641](https://github.com/ocsf/ocsf-schema/pull/1641), merged 2026-06-29), the `delegation`
> object ([#1665](https://github.com/ocsf/ocsf-schema/pull/1665), merged 2026-07-24), the
> `attestation` and `prev_event` objects and the `record_integrity` profile
> ([#1661](https://github.com/ocsf/ocsf-schema/pull/1661), merged 2026-07-17), and `prompt_text` /
> `response_text` on `message_context`
> ([#1674](https://github.com/ocsf/ocsf-schema/pull/1674), merged 2026-06-22). Note that v1.9.0
> **extended** the `ai_operation` profile rather than introducing it — the profile pre-dates the
> release; v1.9.0 widened its reach to the system, network, application, IAM and email activity
> classes and added the `delegation` attribute to it. It did **not** create an `ai` category —
> `categories.json` still stops at `uid 8`. Every OCSF row below has been re-marked against what
> actually shipped, so Part 2 is now evidence rather than restatement. Where a row still says
> *Missing*, that absence was checked, not assumed.
>
> **Reading caveat.** The object and profile files were read from the `main` branch, whose
> `version.json` reports `1.10.0-dev`. Attribute-level claims are therefore accurate as of
> 2026-09-05 against `main`; where `main` has drifted ahead of the v1.9.0 tag, an attribute could be
> newer than the release it is attributed to. Version attribution above is taken from `CHANGELOG.md`
> rather than inferred from file contents.

> **Appendix D verification note, carried forward.** The GenAI conventions **moved** out of the main
> `open-telemetry/semantic-conventions` repository into a dedicated
> **`open-telemetry/semantic-conventions-genai`** repository. Attributes still listed in the main
> registry are marked *Deprecated* to reflect that relocation, not abandonment. Every convention cited
> below carries **Status: Development** — none is stable, which is precisely why this is the moment to
> contribute.

## How to read this document

The AITF position is not that both standards are deficient. It is narrower, and it differs for each.

For **OpenTelemetry**: OTel is not a security telemetry framework and should not become one. Its GenAI
conventions are shaped by observability concerns — latency, cost, token accounting, evaluation quality —
and that is the right centre of gravity. The ask is only for the subset of security-relevant fields that
map cleanly onto the model OTel already has, so that deployments already running OTel get them by
default rather than by bespoke effort. Fields that do not map cleanly are named as out of scope and
left to OCSF. **This document asks OTel to cover what OTel is already shaped to cover, and no more.**

For **OCSF**: OCSF is the canonical at-rest schema and the correct home for the authorization-domain
fields OTel is being asked not to take. The asks there are larger, because the AI-specific object and
class surface is genuinely absent rather than merely incomplete.

Status labels used in the matrices:

| Label | Meaning |
|:---|:---|
| **Covered** | The upstream standard already carries the field. AITF should align to the upstream name rather than define its own. |
| **Partial** | Structurally present but missing the security semantics, or present in a neighbouring signal that requires a documented correlation pattern. |
| **Missing** | No upstream analogue. AITF carries it in the interim. |
| **Out of scope** | Deliberately not asked of this standard. Named as such rather than pushed. |

---

# Part 1 — OpenTelemetry GenAI semantic conventions

## 1.1 What the conventions already cover

The conventions are richer than is commonly assumed, and several fields have a natural OTel home
already. Per Appendix D.1:

**Spans.** `chat`, `text_completion`, `embeddings`, `generate_content`, `execute_tool`, `create_agent`,
`invoke_agent` (client and internal variants), `invoke_workflow`, and a `plan` span, discriminated by
`gen_ai.operation.name`.

**Core attributes.** `gen_ai.provider.name`; `gen_ai.agent.id` / `.name` / `.description` / `.version`;
`gen_ai.conversation.id`; `gen_ai.workflow.name`; `gen_ai.request.model` and the full decoding set
(`temperature`, `top_p`, `top_k`, `max_tokens`, `stop_sequences`, `seed`, `frequency_penalty`,
`presence_penalty`, `choice.count`, `stream`); `gen_ai.response.id` / `.model` / `.finish_reasons`;
`gen_ai.usage.input_tokens` / `.output_tokens` / `.reasoning.output_tokens` / `.cache_read.input_tokens` /
`.cache_creation.input_tokens`.

**Content and definitions.** `gen_ai.input.messages`, `gen_ai.output.messages`,
`gen_ai.system_instructions`, `gen_ai.tool.definitions`, `gen_ai.tool.name` / `.description` / `.type` /
`.call.id` / `.call.arguments` / `.call.result`, `gen_ai.output.type`, `gen_ai.prompt.name`.

**Retrieval.** `gen_ai.retrieval.query.text`, `gen_ai.retrieval.documents`, `gen_ai.data_source.id`,
`gen_ai.embeddings.dimension.count`.

**Memory.** A `gen_ai.memory.*` namespace: `store.id`, `record.id`, `record.count`, `query.text`,
`records`, with a `MemoryRecord` schema (`content`, `id`, `metadata`, `score`); seven memory operations
on `gen_ai.operation.name` (`create_memory`, `create_memory_store`, `delete_memory`,
`delete_memory_store`, `search_memory`, `update_memory`, `upsert_memory`); and a `gen_ai.memory.client`
span, already implemented by `aws-bedrock-agentcore` and `google-adk`.

**Evaluation.** A `gen_ai.evaluation.result` event with `gen_ai.evaluation.name`, `.score.value`,
`.score.label`, `.explanation`.

**Metrics.** `gen_ai.client.operation.duration`, `gen_ai.client.token.usage`,
`gen_ai.client.operation.time_to_first_chunk` / `.time_per_output_chunk`, `gen_ai.execute_tool.duration`,
`gen_ai.invoke_agent.duration` / `.inference_calls` / `.tool_calls`, `gen_ai.invoke_workflow.duration`,
`gen_ai.server.request.duration` / `.time_to_first_token` / `.time_per_output_token`.

**MCP.** A distinct `mcp.*` namespace: `mcp.method.name`, `mcp.protocol.version`, `mcp.request.id`,
`mcp.resource.uri`, `mcp.session.id`, plus `mcp.client.operation.duration`,
`mcp.server.operation.duration`, and client/server `session.duration` metrics.

Three of these close gaps AITF had treated as open. `gen_ai.system_instructions` gives **System Prompt**
(§5) a home. `gen_ai.tool.definitions` gives **Tool Definition Digest** (§9) one — the raw material for a
digest is already in scope, and only the *hash-and-compare* is missing. And
`gen_ai.invoke_agent.tool_calls` / `.inference_calls` are already the shape of **Loop / Step-Count
Signal** (§12) and part of **Resource-Consumption Aggregate** (§12), as metrics rather than attributes.

## 1.2 Coverage matrix by field cluster

From Appendix D.2. The **AITF carrier** column names the namespace that carries the field today; a
namespace marked **new in v0.4 closure** did not exist before the [Appendix C resolution](AITF_gaps.md)
and is defined in the [attributes registry](spec/semantic-conventions/attributes-registry.md).

The tier tag records which tiers the *identified gap* spans, not the full tier mix of the underlying
RFC section.

| Field cluster (tier) | Status | OTel today | What is missing | AITF carrier |
|:---|:---:|:---|:---|:---|
| Execution context & agent identity (MUST / SHOULD) | Partial | `gen_ai.agent.*`, `gen_ai.conversation.id`, `gen_ai.workflow.name`; span hierarchy | No turn/step identifier; no trigger type; no tenant; no surface | [`gen_ai.turn.*` / `gen_ai.step.*`](spec/semantic-conventions/attributes-registry.md#gen_aiturn--gen_aistep-rfc-v04-gap-closure), [`gen_ai.trigger.*`](spec/semantic-conventions/attributes-registry.md#gen_aitrigger-rfc-v04-gap-closure), [`asset.tenant.*`](spec/semantic-conventions/attributes-registry.md#assettenant-and-tenancy-scoping-rfc-v04-gap-closure) — all new in v0.4 closure |
| Prompt / response / system prompt (MUST) | **Covered** | `gen_ai.input.messages`, `gen_ai.output.messages`, `gen_ai.system_instructions` | Nothing structural. Gated behind content-capture opt-in; a **hash-only** capture mode is the remaining ask (§1.5) | `gen_ai.system_prompt.*` — align to upstream names |
| Content modality & attachments (MUST) | Partial | Message parts carry types | No attachment name, size, or **hash** | [`gen_ai.content.*`](spec/semantic-conventions/attributes-registry.md#gen_aicontent-rfc-v04-gap-closure) — new in v0.4 closure |
| Input trust classification (MUST) | **Missing** | none | **No trust-provenance concept anywhere in the conventions** | `security.*` (trust level), [`security.attribute_source.*`](spec/semantic-conventions/attributes-registry.md#securityattribute_source-rfc-v04-gap-closure) — the latter new in v0.4 closure |
| Guardrail verdicts (MUST) | Partial | `gen_ai.evaluation.*`, quality-oriented | No security guardrail verdict, blocked flag, or threat classification | `security.guardrail.*`, [`security.guardrail.modification.*`](spec/semantic-conventions/attributes-registry.md#securityguardrailmodification-rfc-v04-gap-closure) |
| Citations (MUST) | Partial | `gen_ai.retrieval.documents` exists | Output-side citations not linked to retrieved items | [`rag.citation.*`](spec/semantic-conventions/attributes-registry.md#ragcitation-rfc-v04-gap-closure) — new in v0.4 closure |
| Model & serving (SHOULD) | Partial | `gen_ai.request.*` decoding set, `gen_ai.provider.name`, `gen_ai.response.*` | **Inference parameters fully covered.** No model provenance or signing | `supply_chain.*`, `model_ops.registry.*` |
| Token counts & resource aggregates (MUST) | **Covered** | `gen_ai.client.token.usage`, `gen_ai.invoke_agent.*` metrics | Per-run budget/threshold semantics; the security use of the loop metrics is undocumented | `gen_ai.usage.*`, `cost.*` |
| Tools & MCP (MUST / SHOULD) | Partial | `gen_ai.tool.*` incl. `definitions` and `call.id`; full `mcp.*` namespace | No tool-definition **digest**; no MCP **primitive** discriminator; no sandbox/isolation attributes | [`mcp.tool.definition.*`](spec/semantic-conventions/attributes-registry.md#mcptooldefinition-rfc-v04-gap-closure), [`mcp.primitive`](spec/semantic-conventions/attributes-registry.md#mcpprimitive-and-server-identity-rfc-v04-gap-closure), [`supply_chain.runtime.*`](spec/semantic-conventions/attributes-registry.md#supply_chainruntime-rfc-v04-gap-closure) — all new in v0.4 closure |
| Memory (MUST) | Partial | `gen_ai.memory.*`: `store.id`, `record.id`, `record.count`, `query.text`, `records`; seven memory operations; a `gen_ai.memory.client` span | Operation and item identity are **covered**. No **provenance** and no **footprint** attribute exists anywhere in the `gen_ai` registry | `memory.*`, `memory.security.*`, [`memory.config.*`](spec/semantic-conventions/attributes-registry.md#memoryconfig-rfc-v04-gap-closure) |
| Retrieval / RAG (MUST / SHOULD) | Partial | `gen_ai.retrieval.query.text`, `.documents`, `gen_ai.data_source.id` | No per-item **source/provenance**, freshness, or integrity signal | `rag.doc.*`, [`rag.source.*`](spec/semantic-conventions/attributes-registry.md#ragsource-rfc-v04-gap-closure) |
| Output egress (MUST) | **Missing** | none | No link from model output to destination | Correlate via existing HTTP/network semconv on the child span — a documented pattern, not a new attribute |
| Orchestration & multi-agent (MUST) | Partial | `invoke_agent`, `invoke_workflow`, `plan` spans; agent metrics | No inter-agent message attributes; no background-task or termination-condition signal | [`gen_ai.agent.peer.*`](spec/semantic-conventions/attributes-registry.md#gen_aiagentpeer-rfc-v04-gap-closure), [`a2a.task.lifecycle.*`](spec/semantic-conventions/attributes-registry.md#a2atasklifecycle-rfc-v04-gap-closure), [`mcp.envelope.*`](spec/semantic-conventions/attributes-registry.md#mcpenvelope-rfc-v04-gap-closure) |
| Identity & delegation (SHOULD) | **Out of scope** | none | No principal, delegation chain, scope, or attestation | `identity.*` — carry to OCSF Authentication / Delegation instead |
| Asset inventory & AgBOM (MUST / SHOULD) | Partial | Resource attributes, partially | No capability-change event; no BOM reference | [`asset.capability.*`](spec/semantic-conventions/attributes-registry.md#assetcapability-rfc-v04-gap-closure), `supply_chain.ai_bom.*` |
| Policy enforcement & mediation (MUST / SHOULD) | **Out of scope** | none | No authorization decision, taint, approval, or attribute-provenance concept | [`security.authorization.*`](spec/semantic-conventions/attributes-registry.md#securityauthorization-rfc-v04-gap-closure), [`security.taint.*`](spec/semantic-conventions/attributes-registry.md#securitytaint-rfc-v04-gap-closure), [`security.mediation.*`](spec/semantic-conventions/attributes-registry.md#securitymediation-rfc-v04-gap-closure), [`identity.approval.*`](spec/semantic-conventions/attributes-registry.md#identityapproval-rfc-v04-gap-closure) — carry to OCSF |
| Observability-plane integrity (SHOULD) | Partial | none | No enforcement-availability or instrumentation-coverage representation | [`asset.instrumentation.*`](spec/semantic-conventions/attributes-registry.md#assetinstrumentation-rfc-v04-gap-closure) is *partially* natural for OTel as a resource attribute; [`security.enforcement.*`](spec/semantic-conventions/attributes-registry.md#securityenforcement-rfc-v04-gap-closure) outcomes belong to OCSF |

## 1.3 The nine asks, in priority order

From Appendix D.3, ordered by ratio of security value to specification cost. Every item is scoped to
something OTel already models. The **AITF shape** column names the candidate that exists in this
repository today and can be offered as a starting point rather than a blank page.

| # | Ask | Grounding | AITF shape |
|---:|:---|:---|:---|
| 1 | **A trust-provenance attribute on message parts** — `gen_ai.input.trust_level`, distinguishing trusted instruction from untrusted environmental data | The single highest-value addition and the one with no current analogue anywhere in the conventions. Cheap (an enum on an existing structure), and the field the CoSAI Risk Map treats as the core agentic control (§6) | `security.*` trust level; [`security.attribute_source.*`](spec/semantic-conventions/security-cross-cutting.md#attribute-source--trusted-provenance-marking) is the companion that marks *who asserted* the level |
| 2 | **A security-guardrail signal, ideally by extending `gen_ai.evaluation.*`** — adding a blocked/allowed outcome, a guardrail type, and a threat-technique reference | The evaluation event already carries name, score, label and explanation: the right shape for a classifier verdict. It lacks only the security semantics. This is what makes classifier *bypass* detectable (`TA-01`) | `security.guardrail.*`; the alternative dedicated `gen_ai.guardrail.*` namespace is the fallback if reuse is rejected |
| 3 | **Memory provenance and footprint on the existing `gen_ai.memory.*` namespace** | Operation and item identity are covered; provenance and footprint are not, and neither appears anywhere in the `gen_ai` registry. `IR-02` and `IR-05` are direct grounding, `AOC-05` for footprint | `memory.*` provenance, [`memory.config.*`](spec/semantic-conventions/agent-spans.md#declared-memory-configuration-rfc-v04-gap-closure) for the declared footprint bound |
| 4 | **Retrieval provenance** — per-document source, owner, trust level, last-modified | `gen_ai.retrieval.documents` exists; per-document provenance does not, and `TA-09` turns specifically on recently-modified retrievable content | `rag.doc.*`, [`rag.source.*`](spec/semantic-conventions/rag-spans.md) |
| 5 | **Turn and step identifiers** | `gen_ai.conversation.id` and `gen_ai.workflow.name` exist; the intermediate levels do not. `TA-08` and `IR-01` are across-turn patterns (§5) | [`gen_ai.turn.*` / `gen_ai.step.*`](spec/semantic-conventions/gen-ai-spans.md#identifier-hierarchy-rfc-v04-gap-closure) |
| 6 | **`gen_ai.trigger.type` and `gen_ai.trigger.event`** — whether a run was user-initiated or autonomous, and what started it | `TA-01` is zero-click; this is the first filter of any injection hunt | [`gen_ai.trigger.*`](spec/semantic-conventions/gen-ai-spans.md) |
| 7 | **Tool-definition digest and MCP primitive discriminator** | `gen_ai.tool.definitions` and `mcp.method.name` already carry the raw material; a stable hash attribute and an explicit primitive value set make definition drift and non-tool MCP surfaces queryable (§9) | [`mcp.tool.definition.*`](spec/semantic-conventions/mcp-spans.md), [`mcp.primitive`](spec/semantic-conventions/mcp-spans.md) — the value set (tool / resource / prompt / sampling / elicitation / roots) is already enumerated |
| 8 | **Attachment identity on content parts** — name, size, hash | `AOC-12` is an image/OCR injection; `AOC-05` is attachment flooding | [`gen_ai.content.*`](spec/semantic-conventions/gen-ai-spans.md) |
| 9 | **Model provenance attributes** — hash, signature, source | Completes the supply-chain story that `gen_ai.request.model` starts | `supply_chain.*` |

### Two clusters deliberately *not* asked of OTel

**Identity and delegation** (§13) and **policy enforcement** (§16) are authorization-domain concerns
with mature homes elsewhere: OCSF Authentication, and the proposed `ai_authorization` / `ai_taint` /
`ai_approval` objects in [Part 2](#22-what-cosai-asks-ocsf-to-include). Pushing them into OTel would
duplicate schema and invite drift.

The one exception worth raising is **correlation**: an OTel span should be able to reference an
authorization decision by ID so the two layers join at query time. That is a single attribute rather
than a namespace, and AITF already emits it —
[`security.authorization.decision_id`](spec/semantic-conventions/security-cross-cutting.md#authorization-decision-record)
is the join key.

## 1.4 Signal placement

OTel has three signal types and OCSF has one event model, so this guidance has no counterpart in
Part 2 — but getting it wrong is the most common way security telemetry becomes unusable or
unaffordable. From Appendix D.4:

| Signal | Use for | Fields |
|:---|:---|:---|
| **Span attributes** | Low-cardinality identifiers and the execution skeleton | Agent / instance / run / session / turn / step IDs, action type, execution status, model and provider, tool name and type, trigger type, trust level, decision references |
| **Events / logs** | Content and anything high-cardinality, large, or privacy-bearing | Prompts, responses, system instructions, tool arguments and results, memory operations, retrieved documents, citations, guardrail verdicts, capability-change events |
| **Metrics** | Aggregates, budgets, and rate-based detections | Token usage, loop and step counts, tool-call rates, resource aggregates, guardrail block rates, deny rates |

Three rules correct the mistakes that show up most often when GenAI instrumentation is reused for
security. **Never put content in span attributes** — prompts, responses and tool results are unbounded
and often contain PII; OTel already models them as event bodies, and a span attribute carrying a full
prompt breaks cardinality limits and leaks into every trace backend that samples the span. **Emit
detection-relevant aggregates as metrics, not as derived queries** — loop counts and token budgets
(§12) are cheap as metrics and expensive as trace aggregations, and metrics survive sampling where
traces may not. **Cross-reference rather than duplicate** — a guardrail verdict event should carry the
span and trace IDs, not a copy of the prompt it evaluated.

These rules were applied when placing the 27 closed gaps: content-bearing fields went to
[`events.md`](spec/semantic-conventions/events.md) and only bounded hashes and counts were promoted to
spans.

## 1.5 Three operational traps

From Appendix D.5. These are not schema gaps — they are configuration failures that silently defeat the
entire field set even when every attribute above is present.

**Propagation is solved for MCP, and should be adopted deliberately.** The MCP conventions specify that
instrumentations SHOULD inject context into the MCP request `params._meta` property bag, with
`traceparent`, `tracestate` and `baggage` written unprefixed per **SEP-414**, and that the receiver uses
the extracted context as the remote parent. That is exactly the mechanism **Trace Context (propagated)**
(§5) requires, so cross-hop correlation over MCP is a matter of configuration rather than invention.

Two cautions. HTTP-level propagation covers the HTTP request but **not** individual messages within a
streaming request/response, which matters for long-lived agent sessions. And, more important for
security, **`baggage` crosses the trust boundary**. It is attacker-influenceable in exactly the way
§16's provenance rule describes, so baggage may carry correlation identifiers but **must never carry
trust levels, authorization decisions, taint labels, or identity claims**. Those come from the
enforcement point, not from the wire. AITF states this prohibition normatively in
[`security-cross-cutting.md`](spec/semantic-conventions/security-cross-cutting.md).

**Sampling is the trap most likely to silently defeat this entire field set.** OTel's default is
head-based sampling at some fraction of traces. Applied to security telemetry that means *most attacks
are simply not recorded*, and the sample is drawn without regard to whether an event is
security-relevant — a guardrail block has the same chance of being discarded as a routine completion.
Three requirements follow, and the RFC treats them as **normative for any deployment relying on OTel as
its security-telemetry carrier**:

1. **Security-relevant events must not be head-sampled.** Guardrail verdicts, refusals, tool errors,
   authorization denials, capability changes, and any event carrying a fired detection are recorded at
   **100%**.
2. **Where tail sampling is used, security relevance must be a retention predicate.** A trace containing
   a block, a denial, an error, or a flagged classification is always kept.
3. **The sampling configuration in force is itself telemetry.** A detection that never fires because its
   input was sampled away is indistinguishable from a clean environment. This is the same failure mode
   as fail-open enforcement (§15) and deserves the same treatment.

AITF carries requirement 3 as
[`observability.sequence.*`](spec/semantic-conventions/security-cross-cutting.md#event-sequence-continuity),
including `observability.sampling.security_relevant` as the marker samplers must honour.

**Privacy: content capture is opt-in, and that default is correct.** GenAI instrumentations gate message
content behind an explicit capture setting, which aligns with the RFC's privacy-preserving logging
position (§18). The gap is that capture is close to binary — on or off — where security work needs a
**middle setting**: hashes and classifications without raw content, so that correlation (same payload
across many sessions, same attachment hash) survives even where raw capture is prohibited. That is the
companion ask to item 1 in §1.3 and the OTel expression of the RFC's **hash-first** principle. The OTel
Collector is also the correct place to run redaction, since it applies uniformly across every
instrumented service rather than per-library.

---

# Part 2 — OCSF

## 2.1 Coverage matrix by field cluster

From Appendix E.1. For each cluster: where it lands in OCSF today, the gap, the AITF interim carrier,
and the recommended OCSF change.

| Field cluster (tier) | Status | OCSF today | What is missing | AITF interim carrier | Recommended OCSF change |
|:---|:---:|:---|:---|:---|:---|
| Execution context & agent identity (MUST / SHOULD) | Partial | API Activity (6003) base attributes | No agent instance, workflow, action-type or autonomy fields | `gen_ai.agent.*`; emitted on **API Activity (6003)** with the `ai_operation` profile | **Partly landed in v1.9.0**: the `ai_agent` object carries `uid`, `instance_uid`, `name`, `type`/`type_id`, `version`, `charter`, `ai_model`. Still to ask for: **workflow, action-type and autonomy** attributes on the profile |
| Prompt / response / system prompt (MUST) | Partial | `message_context.prompt_text` / `.response_text` (both optional, v1.9.0) | **System prompt** has no field; and the released pair is *raw text only* — no content hash, no redaction or PII marking, no modality or attachment identity, so there is no privacy-preserving way to record that content existed without storing it | `gen_ai.prompt` / `gen_ai.completion` / system message | Narrowed by v1.9.0. Ask is no longer "a place to put content" but **an `ai_content` object beside the raw fields**: content hash + optional raw + redaction/PII flags + modality, so deployments that cannot retain prompt text can still attest to it. Plus a `system_prompt` field |
| Input trust classification (MUST) | Partial | Detection Finding (2004), partial | No trust-provenance enum | `security.*` (trust / threat) | New **`trust_level`** enum (trusted-instruction / trusted-data / untrusted-data / adversarial-suspected), usable on content and message objects |
| Guardrail verdicts (MUST) | Partial | Detection Finding (2004), partial | No guardrail-verdict object | `security.guardrail.*`, `security.blocked`, `security.threat_type` | New **`ai_guardrail`** object: type, verdict, score, blocked flag, threat reference |
| Threat classification / **ATLAS technique tag** (MUST) | Partial | Detection Finding (2004) carries MITRE **ATT&CK** technique | ATLAS `AML.Txxxx` is not a first-class technique value | `security.threat_type` + `compliance.framework=mitre_atlas` / `compliance.control_id` | Extend the finding technique/attack object to accept **ATLAS `AML.Txxxx`** natively, not only ATT&CK |
| Output egress (MUST) | Partial | Network / HTTP Activity, partial | No link from model output → egress channel or recipient | `mcp.tool.call.arguments`, `security.pii.*` | Add an **egress correlation** attribute on Agent Activity linking output → destination / recipient / URL |
| Model & serving (SHOULD) | Partial | API Activity (6003) model attrs; App Lifecycle (6002); Vulnerability Finding (2002) | Provenance and signing not standardized in the profile | `gen_ai.request.model`, `gen_ai.provider.name`, `supply_chain.*` | Standardize model name / version / provider plus **provenance and signing** attributes in `ai_operation` |
| Tools & MCP (MUST / SHOULD) | **Missing** | API Activity (6003) | No MCP object, tool trust-boundary, or ACL/scope | `mcp.*`, `identity.auth.scope_granted` | New **`ai_tool` / `mcp`** object, a **`tool_trust_boundary`** enum (mcp / internal / direct-storage), and a scope attribute |
| Memory (MUST / SHOULD) | **Missing** | Datastore Activity (6005), loosely | No memory-operation object, provenance, or poisoning/isolation signals | `memory.*`, `memory.security.*`, [`memory.config.*`](spec/semantic-conventions/attributes-registry.md#memoryconfig-rfc-v04-gap-closure) | New **`ai_memory`** object: op, provenance, footprint, poisoning score, isolation-verified |
| Retrieval / RAG (MUST / SHOULD) | **Missing** | Datastore Activity (6005) | No retrieved-content source, provenance or integrity | `rag.*`, [`rag.source.*`](spec/semantic-conventions/attributes-registry.md#ragsource-rfc-v04-gap-closure) | New **`ai_retrieval`** object: query, items, source/provenance, integrity signal |
| Orchestration & multi-agent (MUST) | Partial | `message_context.ai_role_id` distinguishes User / Assistant / Tool / **Agent** / **Orchestrator** / **Retriever** / Other; `message_context.uid` correlates a conversation | No inter-agent *message* object, no background task, no loop/step counters, no resource aggregate. The role enum labels a participant's function; it does not describe the exchange between two of them | **API Activity (6003)** with the `ai_operation` profile; [`a2a.task.lifecycle.*`](spec/semantic-conventions/attributes-registry.md#a2atasklifecycle-rfc-v04-gap-closure) | **Narrowed, not closed, by v1.9.0.** `ai_role_id` supplies the vocabulary for *who is playing what part*, which is more than this row previously credited. Ask: an **`agent_message`** object for the exchange itself, plus background-task, loop and resource-aggregate attributes on the profile |
| Identity & delegation (SHOULD) | Partial | Authentication (3002) | No multi-hop delegation chain, granted scope, runtime attestation, or on-behalf-of | `identity.*`; [`identity.boundary.*`](spec/semantic-conventions/attributes-registry.md#identityboundary-rfc-v04-gap-closure), [`identity.credential.mint.*`](spec/semantic-conventions/attributes-registry.md#identitycredentialmint-rfc-v04-gap-closure); emitted on **Authorize Session (3003)** | **Partly landed in v1.9.0**: the `delegation` object shipped with `uid`, `created_time`, `parent_uid`, `issuer_uid` — the *chain*, but not the *authority*. Still to ask for: **delegator, delegatee, type, scope, proof_type, ttl_seconds**; align to ODIS §6.2 / §6.3 |
| Asset inventory & fleet (SHOULD / MAY) | Partial | Inventory Info, partial | No AI-asset object; fleet aggregates are derived | `asset.*` | New **`ai_asset`** object: version, software ref, ownership, status |
| Capability-set change & AgBOM (MUST / SHOULD) | **Missing** | Inventory Info; App Lifecycle (6002) | No agent-composition BOM, no dependency graph, no capability-change event | `supply_chain.ai_bom.*`, [`asset.capability.*`](spec/semantic-conventions/attributes-registry.md#assetcapability-rfc-v04-gap-closure) | New **`ai_bom`** object (BOM ref, format, signature, dependency edges) plus a **capability-change** activity semantics on the `ai_operation` profile |
| Observability-plane integrity (SHOULD) | Partial | **`record_integrity` profile** + `attestation` / `prev_event` objects (v1.9.0) | Event continuity is now representable as a **hash chain**, which detects *alteration*. No **ordinal**, so *omission* stays undetectable. No enforcement-availability, fail-open or hook-coverage representation | `security.guardrail.*` (partial), [`security.enforcement.*`](spec/semantic-conventions/attributes-registry.md#securityenforcement-rfc-v04-gap-closure), [`asset.instrumentation.*`](spec/semantic-conventions/attributes-registry.md#assetinstrumentation-rfc-v04-gap-closure), [`observability.sequence.*`](spec/semantic-conventions/attributes-registry.md#observabilitysequence-rfc-v04-gap-closure) mapped onto `record_integrity` (see [crosswalk §7](spec/ocsf-mapping/ocsf-agentic-crosswalk.md)) | Add **`sequence_number`**, **`gap_detected`** and **`gap_count`** to the released `attestation` object; add **enforcement availability / failure-mode** attributes to `ai_guardrail`; add **instrumentation coverage** to `ai_asset` |
| Policy enforcement & mediation (MUST / SHOULD) | **Missing** | Authorization (3003) partially; Detection Finding (2004) for classifier verdicts | No authorization-decision object for AI operations; no information-flow or taint labels; no human-approval lifecycle; no attribute-provenance marking; no mediation-coverage representation | `security.*` (partial), [`security.authorization.*`](spec/semantic-conventions/attributes-registry.md#securityauthorization-rfc-v04-gap-closure), [`security.taint.*`](spec/semantic-conventions/attributes-registry.md#securitytaint-rfc-v04-gap-closure), [`security.mediation.*`](spec/semantic-conventions/attributes-registry.md#securitymediation-rfc-v04-gap-closure), [`security.attribute_source.*`](spec/semantic-conventions/attributes-registry.md#securityattribute_source-rfc-v04-gap-closure), [`identity.approval.*`](spec/semantic-conventions/attributes-registry.md#identityapproval-rfc-v04-gap-closure) | New **`ai_authorization`** object (decision, reason, code, deciding authority, rule id, obligations); new **`ai_taint`** object (labels, scope, origin, taint-caused denial); new **`ai_approval`** object (correlation id, status, IdP-verified approver, channel, scope-binding result); an **`attribute_source`** enum (idp / pdp / enforcement-state / platform / **self-asserted**) usable on identity, authorization and agent-state attributes |

## 2.2 What CoSAI asks OCSF to include

From Appendix E.2.

1. ~~**Promote the two proposed AI event classes to ratified:** **AI Agent Activity (9001)** and
   **AI Delegation Activity (9002)**.~~ **Answered, in the negative, by OCSF v1.9.0.** OCSF chose to
   represent agentic activity as a *profile over existing classes* rather than as new classes in a new
   category. AITF has followed: agent lifecycle and agent-to-agent messaging emit on **API Activity
   (6003)**, delegation lifecycle on **Authorize Session (3003)**, all carrying `ai_operation`. See
   [`event-classes.md`](spec/ocsf-mapping/event-classes.md).
2. **Standardize an `ai_operation` profile** representing the RFC's classification summary using
   **class- and activity-specific applicability**: common correlation attributes required at profile
   level, the remaining MUST fields required when their defining operation or event applies, and
   SHOULD/MAY fields recommended or optional within the same applicable scope.
3. **Add AI-specific objects:** `ai_content`, `ai_guardrail`, `ai_tool` / `mcp`, `ai_memory`,
   `ai_retrieval`, ~~`ai_delegation`~~, ~~`ai_attestation`~~, `ai_asset`, `ai_bom`, `ai_authorization`,
   `ai_taint`, `ai_approval`. **Two of the twelve shipped in v1.9.0** under different names:
   `ai_delegation` as **`delegation`** (four attributes; the authority fields are still the ask) and
   `ai_attestation` as **`attestation`** on the `record_integrity` profile (the sequence-ordinal fields
   are still the ask). A thirteenth object AITF did not propose, **`ai_agent`**, also shipped.
4. **Add AI-specific enums:** `trust_level`, `autonomy_level` (L1–L5), `tool_trust_boundary`,
   `memory_provenance`, `enforcement_decision` (allow / deny / modify), `enforcement_failure_mode`
   (fail-open / fail-closed), `mcp_primitive` (tool / resource / prompt / sampling / elicitation /
   roots), `trigger_type` (user-initiated / autonomous), `attribute_source` (idp / pdp /
   enforcement-state / platform / self-asserted), `taint_scope` (session / message), `approval_status`
   (pending / resolved / expired / bypassed).
5. **Make MITRE ATLAS a first-class technique reference** on Detection Finding (2004), alongside ATT&CK,
   so the ATLAS technique tag is portable.
6. **Type AI agents as first-class actors.** OWASP AOS's OCSF binding currently represents the agent as
   `actor.type_id: 99` ("Other") with `type: "AI Agent"` — a documented workaround. Ratifying a proper
   actor type removes the dependence on a free-text override. Small change, broad benefit, and AOS's
   need for it is independent corroboration of ask 1. **v1.9.0 answered this, but not the way the ask
   was framed.** OCSF declined to add an actor *type*; instead the released `actor` object carries an
   explicit instruction — *"When the initiator of the activity is an autonomous AI agent, apply the
   `ai_operation` profile and populate `ai_operation.ai_agent` rather than describing the agent here"*
   — and `process` gained `ai_agent` and `hosted_ai_agent_list`. The agent is therefore first-class,
   but as a *sibling of* the actor rather than a *kind of* actor. The practical consequence for AOS and
   for AITF is the same and it is good: `actor.user.type_id: 99` with `type: "AI Agent"` is no longer
   the recommended shape, so the free-text workaround can be retired. **This ask is closed.**

**AITF status against these asks.** Ask 4 is where the v0.4 closure moved furthest, and it is worth being
exact about how far. **Eight of the eleven** proposed enums now have a concrete AITF value set that could
be offered as a starting point. Three have none. Of those eight, **none matches the Appendix E.2 wording
value-for-value**: four are supersets, three diverge on structure, and one pre-dates the closure. Those
differences are the substance of the negotiation, not an untidiness to be smoothed over before filing.

| OCSF proposed enum | AITF carrier | AITF values | Alignment |
|:---|:---|:---|:---|
| `enforcement_decision` | [`security.authorization.decision`](spec/semantic-conventions/security-cross-cutting.md#authorization-decision-record) | allow / deny / challenge / not_applicable / error | **Differs.** E.2 proposes allow / deny / **modify**. AITF has no `modify` and adds `challenge`, `not_applicable`, `error` |
| `enforcement_failure_mode` | [`security.enforcement.failure_mode`](spec/semantic-conventions/security-cross-cutting.md#enforcement-point-availability--failure-mode) | none / fail_open / fail_closed / timeout / unreachable / degraded | **Superset.** E.2's two values plus four operational states |
| `attribute_source` | [`security.attribute_source.*`](spec/semantic-conventions/security-cross-cutting.md#attribute-source--trusted-provenance-marking) | self_asserted / verified / derived / unknown | **Different axis.** E.2 names the *producing system* (idp / pdp / enforcement-state / platform / self-asserted); AITF names the *assurance level*. Both are useful; they are not the same field |
| `taint_scope` | [`security.taint.scope`](spec/semantic-conventions/security-cross-cutting.md#session-taint-labels--information-flow-decisions) | session / turn / message / run | **Superset.** E.2's session / message plus turn and run |
| `mcp_primitive` | [`mcp.primitive`](spec/semantic-conventions/mcp-spans.md) | tool / resource / prompt / sampling / completion / elicitation / root / logging | **Superset**, with one spelling difference: AITF uses singular `root`, E.2 writes `roots` |
| `approval_status` | [`identity.approval.status`](spec/semantic-conventions/identity-spans.md) | pending / resolved / expired | **Differs.** E.2 folds `bypassed` into this enum; AITF keeps lifecycle state separate from outcome and carries `bypassed` on `identity.approval.decision` |
| `trigger_type` | [`gen_ai.trigger.type`](spec/semantic-conventions/gen-ai-spans.md) | user_initiated / scheduled / webhook / event_driven / agent_initiated / system / retry / unknown | **Superset.** E.2's user-initiated / autonomous split, refined into eight |
| `memory_provenance` | [`memory.provenance`](spec/semantic-conventions/attributes-registry.md) | conversation / tool_result / imported | **Matches** in spirit; pre-dates the v0.4 closure |
| `trust_level` | — | — | **No AITF value set.** See below |
| `autonomy_level` (L1–L5) | — | — | **No AITF value set** |
| `tool_trust_boundary` | — | — | **No AITF value set** |

All eight carriers are emitted as enum value sets in all four SDKs — Python, Go, TypeScript and Rust —
so an adopter can key off the constant today and re-map once OCSF settles the wording.

The `trust_level` absence is the one worth flagging rather than burying. It is simultaneously OCSF ask
4 and the highest-value OpenTelemetry ask ([ask 1](#13-the-nine-asks-in-priority-order)), and AITF
defines no equivalent of the trusted-instruction / trusted-data / untrusted-data / adversarial-suspected
value set anywhere. Two attributes carry the *name* — `rag.source.trust_level` and
`identity.trust.trust_level` — but neither is that enum. This was not one of the 27 Appendix C gaps, so
the closure did not touch it; it is new work, not a regression.

Ask 3 is partially covered: `ai_authorization`, `ai_taint` and `ai_approval` now have complete attribute
sets in AITF that can serve as the object field lists.

## 2.3 The incremental path, and where AITF sits on it

From Appendix E.3. AITF exists so adopters can emit this telemetry **before** OCSF ratifies it, and so
each upstream proposal is staged to be backward-compatible.

| Phase | Content | AITF status |
|:---|:---|:---|
| **Phase 0 — today** | AITF carries every MUST field as OTel attributes and emits OCSF via the **released** `ai_operation` profile on existing classes (6003 / 6005 / 2004 / 3002 / 3003); the ATLAS tag rides on `compliance.control_id` with framework `mitre_atlas` | **Complete.** All 27 Appendix C gaps closed; see the [resolution log](AITF_gaps.md) |
| **Phase 1 — profile extension** (backward-compatible) | Contribute the MUST attribute set plus the `ai_content`, `ai_guardrail` and `ai_tool` objects to the OCSF `ai_operation` profile; no new classes required | **Ready to file, and now easier than when this was written.** The profile is no longer hypothetical — v1.9.0 released it, so every Phase 1 item is an *addition to an existing profile* rather than a proposal for a new one. Attribute sets exist; see [`spec/ocsf-mapping/`](spec/ocsf-mapping/) |
| **Phase 2 — agentic objects** (was "agentic classes") | The class half is **done and was answered differently than proposed**: no new class, agent lifecycle rides **API Activity (6003)** with `ai_operation`. What remains of Phase 2 is the object half — **`ai_memory`** and **`ai_retrieval`**, both still absent from OCSF and both MUST here | **Unblocked, and half-delivered.** The category-UID divergence described in [`COSAI-OCSF-discussions.md`](COSAI-OCSF-discussions.md) is resolved in practice by OCSF declining, so far, to create an `ai` category at all — though issue #1640 remains open with an agent-layer class placement still TBD; AITF has now retired *both* its bespoke Category 7 *and* its provisional `uid 9`, so it is unaffected either way. See [`event-classes.md`](spec/ocsf-mapping/event-classes.md) and the [OCSF PR plan §1](upstream-pr-plan-ocsf.md#1-the-category-question-is-answered-in-practice-but-not-formally-closed) |
| **Phase 3 — delegated authority** | The class and the two objects all landed in v1.9.0 in reduced form — delegation lifecycle on **Authorize Session (3003)**, the **`delegation`** object (4 attributes), the **`attestation`** object on `record_integrity`. What remains is the **authority** fields on `delegation` (delegator, delegatee, type, scope, proof_type, ttl) and the **ordinal** fields on `attestation`, plus the ATLAS technique reference on findings | **No longer depends on Phase 2** — the structures it needed exist. Attribute sets exist ([`identity.*`](spec/semantic-conventions/identity-spans.md), ODIS §6.1–6.3 mapped) and are mapped onto the released objects in the [crosswalk](spec/ocsf-mapping/ocsf-agentic-crosswalk.md) §4 and §7 |
| **Phase 4 — inventory & analytics** | SHOULD/MAY clusters — `ai_asset`, drift, quality, fleet aggregates — as they stabilize | Attribute sets exist ([`asset.*`](spec/semantic-conventions/asset-inventory-spans.md), [`drift.*`](spec/semantic-conventions/drift-detection-spans.md)) |

**Governance rule for promotion.** A field graduates from AITF-proposed to an OCSF standardization ask
when **≥ 2 independent attacks in Appendix A** require it — the same attack-grounding rule that sets the
MUST tier. This keeps the OCSF surface minimal and evidence-driven rather than speculative. Every one of
the 27 closed gaps records its grounding attacks in the [resolution log](AITF_gaps.md), so the rule can
be applied per-field rather than argued in the abstract.

---

# Part 3 — Consolidated status

The 27 fields closed in AITF v0.4, and whether each has an upstream home. This is the delta a reader
should care about: AITF defining an attribute is a workaround, and the workaround is only retired when
one of the two columns below turns green.

| # | AITF field | Tier | OTel status | OCSF status |
|---:|:---|:---:|:---|:---|
| 1 | Session / Turn / Step IDs | MUST | **Ask D.3-5** — `conversation.id` and `workflow.name` exist, intermediate levels do not | **Still open.** v1.9.0 gave a home (6003 + `ai_operation`) but no turn/step attributes |
| 2 | Trigger Type & Source Event | MUST | **Ask D.3-6** | Enum ask E.2-4 (`trigger_type`) |
| 3 | Content Modality & Attachment Identity | MUST ‡ | **Ask D.3-8** — parts carry types, identity is missing | Object ask E.2-3 (`ai_content`) |
| 4 | Peer Agent Card / Descriptor | SHOULD | Missing — no inter-agent message attributes | **Partly landed.** `ai_agent` (v1.9.0) describes *an* agent; no object describes the *peer* in a message. Ask: `agent_message` |
| 5 | Backend / Route Restriction Decision | SHOULD | Out of scope — enforcement domain | Object ask E.2-3 (`ai_authorization`) |
| 6 | Guardrail Modification Record | SHOULD ‡ | Part of ask D.3-2 | Object ask E.2-3 (`ai_guardrail`) |
| 7 | Encoded / Obfuscated Payload Indicator | MAY | Missing — no upstream ask filed; low tier | Missing |
| 8 | Authorization Decision Record | MUST | **Out of scope by design** — correlation attribute only | Object ask E.2-3 (`ai_authorization`) |
| 9 | Session Taint Labels & Information-Flow Decisions | SHOULD | **Out of scope by design** | Object ask E.2-3 (`ai_taint`) + enum `taint_scope` |
| 10 | Enforcement-Point Availability & Failure Mode | SHOULD | Partial — resource attribute is natural | E.1 observability-plane row + enum `enforcement_failure_mode` |
| 11 | Tool Definition Digest | SHOULD | **Ask D.3-7** — `tool.definitions` exists, hash does not | Object ask E.2-3 (`ai_tool`) |
| 12 | MCP Server Identity & Primitive | SHOULD | **Ask D.3-7** — `mcp.method.name` partially serves | Object ask E.2-3 (`mcp`) + enum `mcp_primitive` |
| 13 | Protocol Envelope Capture | MAY | Missing — no upstream ask filed; low tier | Missing |
| 14 | Declared Memory Configuration | MAY | Adjacent to ask D.3-3 (footprint) | Object ask E.2-3 (`ai_memory`) |
| 15 | Citations / Source Attribution | MUST | Missing — `retrieval.documents` exists, output-side link does not | Object ask E.2-3 (`ai_retrieval`) |
| 16 | Declared Knowledge-Source Configuration | MAY | Adjacent to ask D.3-4 | Object ask E.2-3 (`ai_retrieval`) |
| 17 | Trust-Domain Crossing & Delegation Depth | SHOULD | **Out of scope by design** | **Still open.** `delegation.parent_uid` gives the chain; depth and trust-domain crossing must be *derived* by walking it. ODIS §6.3 |
| 18 | Credential Minting & Scope-Narrowing Check | SHOULD | **Out of scope by design** | **Still open, and the sharpest remaining ask.** The released `delegation` object has no `scope`, so scope *attenuation* — the confused-deputy signal — is not representable. ODIS §6.2/§6.3 |
| 19 | Human Approval / Elicitation Event | SHOULD | **Out of scope by design** | Object ask E.2-3 (`ai_approval`) + enum `approval_status` |
| 20 | Organization / Tenant ID | SHOULD | Missing — "adopt existing tenant/resource attrs" per D.2 | Not separately asked; rides on profile E.2-2 |
| 21 | Capability-Set Change Event | MUST | Missing — capability-change event is a D.2 recommendation | Object ask E.2-3 (`ai_bom`) + a capability-change activity on the `ai_operation` profile |
| 22 | Instrumentation Coverage / Hook Attestation | SHOULD | Partial — "OTel can record hook coverage as a resource attribute" | E.1 observability-plane row (`ai_asset`) |
| 23 | Execution Environment / Sandbox | MUST | Missing — part of the tools/MCP D.2 recommendation | Object ask E.2-3 (`ai_tool`) |
| 24 | A2A Task Lifecycle Event | SHOULD | Missing — inter-agent messaging attributes | **Still open.** Emits on 6003 today; `message_context` (v1.9.0) is a correlation handle, not a task lifecycle |
| 25 | Attribute Source / Trusted-Provenance Marking | MUST ‡ | Companion to **ask D.3-1**; carries the D.5 baggage prohibition | **Enum ask E.2-4** (`attribute_source`) |
| 26 | Mediation Coverage & Bypass Path | SHOULD | **Out of scope by design** | E.1 policy-enforcement row |
| 27 | Event Sequence Continuity | SHOULD | Partial — bound up with the D.5 sampling requirements | **Partly landed.** `record_integrity` / `attestation` / `prev_event` (v1.9.0) carry the hash chain; the ordinal and gap markers are the residual ask |

Read as a whole: of the 27, **nine** map to a named OpenTelemetry ask, **six** are deliberately out of
OTel's scope, and **all but two** have a named OCSF landing place. The two without one — Encoded /
Obfuscated Payload Indicator and Protocol Envelope Capture — are both MAY-tier, and both fall below the
Appendix E.3 promotion threshold of two independent grounding attacks. They stay AITF-local until the
attack evidence changes.

**What OCSF v1.9.0 moved.** Four of the 27 rows changed status: #4 (peer agent card), #17 (delegation
depth), #18 (credential minting) and #27 (event-sequence continuity) all went from *no upstream
structure* to *upstream structure exists, carrying part of the field*. None went fully green, and the
pattern in what shipped versus what did not is consistent enough to be worth naming: OCSF took the
**identity and integrity plumbing** — who the agent is, which delegation this descends from, what this
record hashes to — and left the **authority and completeness semantics** — what scope was granted,
whether it narrowed, whether any event is missing. That is the residual AITF surface, and it is a
coherent one to argue for rather than a scattering of leftovers.

The concentration is worth stating plainly. The OTel asks cluster on **content and execution structure**
(trust level, guardrail semantics, memory and retrieval provenance, turn/step IDs, trigger type) — all
extensions to structures OTel already has. The OCSF asks cluster on **authority** (authorization
decisions, delegation, approval, taint) — all genuinely new objects. That split is the whole argument of
Appendices D and E, and it is why the two proposals should be sequenced together rather than filed as
one.

---

# Part 4 — Filing checklist

Before any of this is taken upstream:

1. **Re-verify against the live registries.** Appendix D's names were read at the time of writing.
   `open-telemetry/semantic-conventions-genai` is the current home; names in the main
   `semantic-conventions` registry are marked *Deprecated* to reflect relocation, and everything is
   **Status: Development**.
2. ~~**Confirm the OCSF category UID is ratified, not merely proposed.**~~ **Done — the answer was
   "there is no category."** `categories.json` in the released v1.9.0 schema stops at `uid 8`
   (`unmanned_systems`). OCSF represented agentic activity as a *profile over existing classes*
   instead. AITF has retired the provisional `uid 9` and the three `9xxx` classes that lived in it; the
   SDKs keep a `LEGACY_AI_CLASS_UIDS` map so telemetry emitted before v0.4 stays decodable. See
   [`event-classes.md`](spec/ocsf-mapping/event-classes.md),
   [`COSAI-OCSF-discussions.md`](COSAI-OCSF-discussions.md) §1 and §3, and the
   [OCSF PR plan §1](upstream-pr-plan-ocsf.md#1-the-category-question-is-answered-in-practice-but-not-formally-closed).
3. **Sequence D and E together.** AITF is the interim binding between them; filing one without the other
   leaves the binding pointing at a moving target.
4. **Apply the E.3 promotion rule per field.** Two independent Appendix A attacks, or the field stays
   AITF-local. The grounding attacks for all 27 are recorded in the [resolution log](AITF_gaps.md).
5. **Work from the filing plans, not from this document.** Part 2 says what is missing; the
   [OpenTelemetry](upstream-pr-plan-opentelemetry.md) and [OCSF](upstream-pr-plan-ocsf.md) PR plans stage
   it into ordered pull requests with per-filing acceptance criteria and their own verification gates.

## What this document does not claim

It does not claim that the upstream state described here is current as of the date you are reading it.

The two Parts differ in what backs them, and the difference matters when deciding how much weight to
put on a row. **Part 1 (OpenTelemetry)** is a faithful restatement of RFC v0.4 Appendix D plus a mapping
onto the AITF namespaces that now exist; nothing in it was verified against a live registry, by explicit
constraint. Treat Part 1 as the filing brief, not as evidence. **Part 2 (OCSF)** was re-verified on
2026-09-05 — objects, profiles, `categories.json`, `CHANGELOG.md` and the merged PRs were read
directly — so its *Covered* and *Partial* marks are evidence. Its *Missing* marks are evidence of
absence at that date and nothing later.

Two caveats on Part 2 specifically. The schema files were read from the **`main` branch**, whose
`version.json` reports `1.10.0-dev`; version attribution therefore comes from `CHANGELOG.md` rather
than from the file contents, and an attribute seen on `main` could in principle post-date the v1.9.0
tag. And Part 2 has already been wrong once in a way worth recording: an earlier revision asserted
that `actor.type_id` had no AI-agent member and that the free-text workaround survived. In fact
`actor` has no `type_id` attribute at all, and the object's own description directs producers to
populate `ai_operation.ai_agent` instead. The error came from writing a plausible claim from memory
rather than reading the file. Rows here should be trusted to the extent they cite something
re-readable.

It also does not claim that any of the AITF attributes referenced here are *emitted* by a given
deployment. They are *defined*: present in the registry, documented in a convention document, and
available as SDK constants. Emission is an instrumentation question, and
[`asset.instrumentation.*`](spec/semantic-conventions/security-cross-cutting.md) exists precisely because
the difference between defined and emitted is where security telemetry usually fails.

## Related documents

- [Upstream PR plan: OpenTelemetry](upstream-pr-plan-opentelemetry.md) — the nine asks staged as filable pull requests
- [Upstream PR plan: OCSF](upstream-pr-plan-ocsf.md) — existing drafts plus the v0.4 authority-cluster delta, sequenced
- [AITF gap closure: RFC v0.4 Appendix C resolution log](AITF_gaps.md)
- [Attributes registry](spec/semantic-conventions/attributes-registry.md)
- [Cross-cutting security & observability-plane conventions](spec/semantic-conventions/security-cross-cutting.md)
- [OCSF ↔ CoSAI AITF alignment discussion](COSAI-OCSF-discussions.md)
- [OCSF mapping](spec/ocsf-mapping/)
- [Specification overview](spec/overview.md)
