# GenAI Span Conventions (AI_INTERACTION)

Status: **Normative** | CoSAI WS2 Alignment: **AI_INTERACTION** | OCSF Class: **API Activity (6003)** (`ai_operation` profile)

AITF preserves and extends OpenTelemetry GenAI span conventions for LLM inference operations. This specification defines the normative field requirements for AI interaction telemetry, aligned with CoSAI Working Stream 2 (Telemetry for AI) and mapped to applicable compliance and threat frameworks.

> **OTel Alignment Note:** This spec aligns with OpenTelemetry GenAI Semantic Conventions v1.38+.
> The following attributes are deprecated and replaced: `gen_ai.system` → `gen_ai.provider.name`,
> `gen_ai.prompt` → `gen_ai.input.messages`, `gen_ai.completion` → `gen_ai.output.messages`.
> Extension attributes use short namespaces (e.g., `cost.*`, `latency.*`, `security.*`) without the `aitf.` prefix.

Key words "MUST", "SHOULD", "MAY" follow [RFC 2119](https://www.rfc-editor.org/rfc/rfc2119).

---

## Span: `gen_ai.inference`

Represents a single LLM inference request (chat completion, text completion, embedding).

### Span Name

Format: `{gen_ai.operation.name} {gen_ai.request.model}`

Examples:
- `chat gpt-4o`
- `chat claude-sonnet-4-5-20250929`
- `embeddings text-embedding-3-small`

### Span Kind

`CLIENT`

---

## Normative Field Table

Instrumentors MUST emit all Required fields. Instrumentors SHOULD emit Recommended fields when the data is available. Optional fields MAY be emitted for enhanced observability.

### Core Identification

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.provider.name` | string | **Required** | AI system provider identifier (e.g. `"openai"`, `"anthropic"`, `"bedrock"`) | MITRE ATLAS [AML.T0044](https://atlas.mitre.org/techniques/AML.T0044), NIST AI RMF MAP-1.1 |
| `gen_ai.operation.name` | string | **Required** | Operation type: `"chat"`, `"text_completion"`, `"embeddings"` | NIST AI RMF MAP-1.1 |
| `gen_ai.request.model` | string | **Required** | Requested model identifier (e.g. `"gpt-4o"`, `"claude-sonnet-4-5-20250929"`) | MITRE ATLAS [AML.T0044](https://atlas.mitre.org/techniques/AML.T0044), EU AI Act Art.13 |
| `gen_ai.conversation.id` | string | **Recommended** | Conversation or session identifier | NIST AI RMF GOVERN-1.2 |
| `server.address` | string | **Recommended** | API endpoint hostname | MITRE ATLAS [AML.T0044](https://atlas.mitre.org/techniques/AML.T0044), NIST AI RMF MAP-1.5 |
| `server.port` | int | **Optional** | API endpoint port | NIST AI RMF MAP-1.5 |

### Identifier Hierarchy [RFC v0.4 gap closure]

> **RFC field:** Session / Turn / Step IDs · **Tier:** MUST · **§5 Application & Agent Reasoning Core** · **Grounding attacks:** `AOC-03`, `AOC-07`, `TA-07`, `TA-08`

The RFC defines a five-level identifier hierarchy — instance → run → session → turn → step. AITF already carried instance (`gen_ai.agent.id`), run (the OTel `trace_id`), and session (`gen_ai.conversation.id`). The **turn** level had no attribute at all, and the **step** level had only an ordinal (`gen_ai.agent.step.index`) with no identifier independent of the OTel `span_id`. Without the turn level, a detection can say *which run* behaviour changed in but not *which turn*, and multi-turn attacks — where a benign turn establishes context that a later turn exploits — cannot be localised.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.turn.id` | string | **Required** | Identifier for one request→response cycle within a session, stable across every span emitted for that turn | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 |
| `gen_ai.turn.index` | int | **Recommended** | Zero-based ordinal of this turn within `gen_ai.conversation.id` | NIST AI RMF MEASURE-2.5 |
| `gen_ai.turn.parent_id` | string | **Optional** | Preceding turn identifier, where turns branch or are retried | NIST AI RMF GOVERN-1.2 |
| `gen_ai.step.id` | string | **Recommended** | Identifier for one action within a turn. Defaults to the OTel `span_id`; set explicitly where the framework owns step identity | NIST AI RMF GOVERN-1.2 |
| `gen_ai.step.parent_id` | string | **Optional** | Preceding step identifier within the turn | NIST AI RMF GOVERN-1.2 |
| `gen_ai.run.id` | string | **Recommended** | Workflow/run identifier where the deployment does not use the OTel `trace_id` as run identity | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 |

**Emission rule.** These are span attributes, never event bodies, and all are low-cardinality relative to content. A span carrying `gen_ai.turn.id` MUST also carry `gen_ai.conversation.id` — a turn identifier with no session to belong to cannot be ordered against anything.

### Trigger Provenance [RFC v0.4 gap closure]

> **RFC field:** Trigger Type & Source Event · **Tier:** MUST · **§5 Application & Agent Reasoning Core** · **Grounding attacks:** `TA-01`, `AOC-04`, `AOC-10`, `AOC-12`

Records whether a run was started by a human or by the environment, and for autonomous runs, what started it. This is distinct from the surface or application, which records the *channel* rather than the *initiator*. `TA-01` is the zero-click case: content arrives, the agent acts, and no human ever asked for anything. "Show me every autonomous run triggered by inbound email" is therefore the first filter of any injection hunt, and it is unanswerable without these fields.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.trigger.type` | string | **Required** | How the run started: `"user_initiated"`, `"autonomous"`, `"scheduled"`, `"system"`, `"agent_delegated"` | OWASP LLM06, MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `gen_ai.trigger.event` | string | **Conditionally Required** (when `type` ≠ `user_initiated`) | Originating event class: `"inbound_email"`, `"chat_message"`, `"webhook"`, `"schedule"`, `"file_change"`, `"queue_message"`, `"alert"`, `"agent_message"` | OWASP LLM01 (Prompt Injection) |
| `gen_ai.trigger.event.id` | string | **Recommended** | Identifier of the specific originating event | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 |
| `gen_ai.trigger.source` | string | **Recommended** | System or channel that emitted the trigger (`"o365:mailbox"`, `"slack:C0123"`, `"github:webhook"`) | OWASP LLM01, MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `gen_ai.trigger.source.principal` | string | **Recommended** | Principal that caused the trigger, where identifiable. Mark provenance with `security.attribute_source.*` | NIST AI RMF GOVERN-1.2 |
| `gen_ai.trigger.received_at` | string | **Optional** | ISO 8601 timestamp the trigger was received, distinct from run start | EU AI Act Art.12 |
| `gen_ai.trigger.human_in_loop` | boolean | **Recommended** | Whether a human confirmed the run before it started | EU AI Act Art.14 (Human Oversight) |

### Request Configuration

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.request.max_tokens` | int | **Recommended** | Maximum tokens to generate | OWASP LLM10 (Unbounded Consumption) |
| `gen_ai.request.temperature` | double | **Recommended** | Sampling temperature (0.0–2.0) | NIST AI RMF MEASURE-2.5 |
| `gen_ai.request.top_p` | double | **Recommended** | Nucleus sampling parameter (0.0–1.0) | NIST AI RMF MEASURE-2.5 |
| `gen_ai.request.top_k` | int | **Optional** | Top-k sampling parameter | NIST AI RMF MEASURE-2.5 |
| `gen_ai.request.stop_sequences` | string[] | **Optional** | Stop sequences | — |
| `gen_ai.request.frequency_penalty` | double | **Optional** | Frequency penalty | NIST AI RMF MEASURE-2.5 |
| `gen_ai.request.presence_penalty` | double | **Optional** | Presence penalty | NIST AI RMF MEASURE-2.5 |
| `gen_ai.request.seed` | int | **Optional** | Random seed for reproducibility | NIST AI RMF MEASURE-2.5, EU AI Act Art.12 |
| `gen_ai.request.stream` | boolean | **Recommended** | Whether streaming is enabled | — |
| `gen_ai.tool.definitions` | string | **Recommended** | Tool/function definitions (JSON) | OWASP LLM06 (Excessive Agency) |
| `gen_ai.request.tool_choice` | string | **Optional** | Tool selection mode (`"auto"`, `"required"`, `"none"`) | OWASP LLM06 (Excessive Agency) |
| `gen_ai.request.response_format` | string | **Optional** | Expected response format (`"json_object"`, `"text"`) | — |
| `gen_ai.output.type` | string | **Optional** | Output type (e.g. `"text"`, `"json"`, `"tool_calls"`) | — |

### Prompt & Completion Content

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.input.messages` | string | **Recommended** | Input messages content (emitted as event `gen_ai.content.prompt`) | OWASP LLM01 (Prompt Injection), MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `gen_ai.system_prompt.hash` | string | **Recommended** | SHA-256 hash of system prompt (enables leak detection without storing content) | OWASP LLM07 (System Prompt Leakage), MITRE ATLAS [AML.T0051.001](https://atlas.mitre.org/techniques/AML.T0051) |
| `gen_ai.output.messages` | string | **Recommended** | Output completion content (emitted as event `gen_ai.content.completion`) | OWASP LLM05 (Improper Output), OWASP LLM02 (Sensitive Info Disclosure) |

### Content Modality & Attachment Identity [RFC v0.4 gap closure]

> **RFC field:** Content Modality & Attachment Identity · **Tier:** MUST ‡ · **§6 Input Handling & Trust Provenance** · **Grounding attacks:** `AOC-12`, `AOC-05`, `TA-05`, `TA-01`

**Cross-cutting (‡).** These fields apply to every content-bearing field in AITF, not to model input alone: model output, tool arguments and results, retrieved documents, memory records, and inter-agent messages all carry parts, and all can carry parts that are not text.

That is the point of the group. Instructions arriving as an image, a PDF, or a structured blob are invisible to text-only inspection *and* to text-only logging, which means a system can pass every content check it runs and still have been instructed by something nobody looked at. `gen_ai.content.part.hash` matters most of the four Required fields because it survives redaction: correlating the same attachment across sessions, tenants, and users works even where raw capture is prohibited by policy.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.content.part.index` | int | **Recommended** | Zero-based index of this part within its message | — |
| `gen_ai.content.part.type` | string | **Required** | `"text"`, `"file"`, `"image"`, `"audio"`, `"video"`, `"structured_data"`, `"tool_call"`, `"tool_result"` | OWASP LLM01, MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `gen_ai.content.part.mime_type` | string | **Conditionally Required** (for non-text parts) | IANA media type of the part | OWASP LLM01 |
| `gen_ai.content.part.size_bytes` | int | **Recommended** | Size of the part in bytes | OWASP LLM10 (Unbounded Consumption) |
| `gen_ai.content.part.hash` | string | **Required** | SHA-256 of the part's raw bytes. Emitted even when raw capture is disabled | MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051), NIST AI RMF GOVERN-1.2 |
| `gen_ai.content.modalities` | string[] | **Recommended** | Distinct modalities present across all parts | OWASP LLM01 |
| `gen_ai.content.attachment.count` | int | **Recommended** | Number of non-text parts on the message | OWASP LLM10 |
| `gen_ai.content.attachment.name` | string | **Recommended** | Declared filename of a file part | OWASP LLM01, MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010) |
| `gen_ai.content.attachment.hash` | string | **Conditionally Required** (for file parts) | SHA-256 of the attachment bytes | MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010) |
| `gen_ai.content.attachment.size_bytes` | int | **Recommended** | Attachment size in bytes | OWASP LLM10 |
| `gen_ai.content.attachment.source` | string | **Recommended** | `"user_upload"`, `"tool_result"`, `"retrieval"`, `"inbound_message"`, `"memory"` | OWASP LLM01, EU AI Act Art.13 |
| `gen_ai.content.attachment.extracted_text_hash` | string | **Optional** | SHA-256 of text extracted by OCR or parsing. Pairs with `security.obfuscation.*` for `AOC-12` | MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `gen_ai.content.total_size_bytes` | int | **Optional** | Total byte size of all parts, for flooding detection (`AOC-05`) | OWASP LLM10 |

**Signal placement (RFC Appendix D.4).** `gen_ai.content.*` is emitted on **events**, not span attributes — the per-part fan-out is unbounded and belongs in the event body. The sole exception is the hash and count attributes, which are bounded and MAY be promoted onto the span so they stay queryable under content-capture-off configurations. That exception is what keeps attachment correlation possible in the deployments most likely to need it.

### Response Metadata

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.response.id` | string | **Recommended** | Provider-assigned response identifier | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 |
| `gen_ai.response.model` | string | **Recommended** | Actual model used (may differ from requested) | MITRE ATLAS [AML.T0044](https://atlas.mitre.org/techniques/AML.T0044), EU AI Act Art.13 |
| `gen_ai.response.finish_reasons` | string[] | **Recommended** | Finish reasons (`"stop"`, `"length"`, `"tool_calls"`, `"content_filter"`) | NIST AI RMF MEASURE-2.5 |

### Token Usage

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.usage.input_tokens` | int | **Required** | Input/prompt token count | OWASP LLM10 (Unbounded Consumption), NIST AI RMF MEASURE-2.5 |
| `gen_ai.usage.output_tokens` | int | **Required** | Output/completion token count | OWASP LLM10 (Unbounded Consumption), NIST AI RMF MEASURE-2.5 |
| `gen_ai.usage.cache_read.input_tokens` | int | **Optional** | Cached/prefix input tokens read | — |
| `gen_ai.usage.cache_creation.input_tokens` | int | **Optional** | Input tokens used to create cache entries | — |

### Latency & Performance

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `latency.total_ms` | double | **Required** | Total request latency in milliseconds | NIST AI RMF MEASURE-2.5, OWASP LLM10 |
| `latency.time_to_first_token_ms` | double | **Recommended** | Time to first token (streaming) in milliseconds | NIST AI RMF MEASURE-2.5 |
| `latency.tokens_per_second` | double | **Optional** | Token generation throughput | NIST AI RMF MEASURE-2.5 |
| `latency.queue_time_ms` | double | **Optional** | Time spent in request queue | — |
| `latency.inference_time_ms` | double | **Optional** | Pure inference time (excluding queue) | — |

### Cost Attribution

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `cost.total_cost` | double | **Recommended** | Total request cost in USD | OWASP LLM10 (Unbounded Consumption), NIST AI RMF GOVERN-1.5 |
| `cost.input_cost` | double | **Optional** | Input token cost in USD | OWASP LLM10 |
| `cost.output_cost` | double | **Optional** | Output token cost in USD | OWASP LLM10 |

### Security Enrichment

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `security.risk_score` | double | **Optional** | Security risk score (0–100) | OWASP LLM01–LLM10 |
| `quality.confidence` | double | **Optional** | Response confidence score (0.0–1.0) | NIST AI RMF MEASURE-2.5 |

### Backend / Route Restriction Decision [RFC v0.4 gap closure]

> **RFC field:** Backend / Route Restriction Decision · **Tier:** SHOULD · **§16 Policy Enforcement & Mediation** · **Grounding attacks:** `AOC-05`, `AOC-06`, `TA-01` · **ODIS:** `resource_indicators`, `constraints` (6.3)

`gen_ai.provider.name` and `gen_ai.response.model` record where an operation *did* execute. These fields record where it was *allowed* to execute, which is a different question and the one that data-residency and model-restriction policy actually turns on. The security value concentrates in two places: the constraint that narrowed the candidate set, and the behaviour when **no candidate qualified**. An unrecorded fail-open at that second point silently defeats the whole policy — the request completes, the response looks normal, and nothing anywhere records that the restriction was abandoned.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.route.selected` | string | **Required** | Backend/model actually selected | NIST AI RMF MAP-1.1, EU AI Act Art.13 |
| `gen_ai.route.decision` | string | **Required** | `"allowed"`, `"denied"`, `"fallback"`, `"no_candidate"` | NIST AI RMF GOVERN-1.2 |
| `gen_ai.route.candidates` | string[] | **Recommended** | Candidates considered before constraint evaluation | NIST AI RMF MAP-1.1 |
| `gen_ai.route.candidates.count` | int | **Optional** | Size of the candidate set before narrowing | — |
| `gen_ai.route.constraint.type` | string | **Recommended** | `"region"`, `"data_residency"`, `"model"`, `"site"`, `"cost_tier"`, `"tenant"`, `"custom_label"` | EU AI Act Art.13, NIST AI RMF GOVERN-1.2 |
| `gen_ai.route.constraint.value` | string | **Recommended** | The constraint value applied (`"eu-west-1"`, `"gpt-4o-only"`) | NIST AI RMF GOVERN-1.2 |
| `gen_ai.route.constraint.source` | string | **Recommended** | `"policy"`, `"tenant_config"`, `"request"`, `"default"`. Request-sourced constraints are attacker-influenceable; mark with `security.attribute_source.*` | NIST AI RMF GOVERN-1.2 |
| `gen_ai.route.excluded` | string[] | **Optional** | Candidates removed by the constraint | NIST AI RMF GOVERN-1.2 |
| `gen_ai.route.no_candidate_action` | string | **Recommended** (when `decision` = `no_candidate`) | `"fail_closed"`, `"fail_open"`, `"fallback_default"`, `"queued"`, `"error"`. `"fail_open"` here is a reportable control failure | NIST AI RMF GOVERN-1.2, EU AI Act Art.14 |
| `gen_ai.route.policy_id` | string | **Optional** | Routing policy in force. Joins to `security.authorization.rule_id` | NIST AI RMF GOVERN-1.2 |

This group is deliberately distinct from `model_ops.serving.route.*`, which is the *operations* view of routing — cost, latency, fallback health. `gen_ai.route.*` is the *policy* view: what was permitted and what happened when nothing was. Implementations may populate both; they answer to different readers.

---

## Tool Call Events

### Event: `gen_ai.tool.call`

Emitted when the model requests a tool/function call.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.tool.name` | string | **Required** | Tool/function name | OWASP LLM06 (Excessive Agency) |
| `gen_ai.tool.call.id` | string | **Required** | Tool call identifier | NIST AI RMF GOVERN-1.2 |
| `gen_ai.tool.call.arguments` | string | **Recommended** | Tool arguments (JSON) | OWASP LLM01 (Prompt Injection) |
| `gen_ai.tool.type` | string | **Optional** | Tool type (`"function"`, `"extension"`, `"datastore"`) | OWASP LLM06 (Excessive Agency) |
| `gen_ai.tool.description` | string | **Optional** | Human-readable tool description | — |

### Event: `gen_ai.tool.result`

Emitted when a tool/function returns its result.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.tool.name` | string | **Required** | Tool/function name | OWASP LLM06 |
| `gen_ai.tool.call.id` | string | **Required** | Tool call identifier | NIST AI RMF GOVERN-1.2 |
| `gen_ai.tool.call.result` | string | **Recommended** | Tool result content | OWASP LLM05 (Improper Output) |
| `gen_ai.operation.name` | string | **Recommended** | `"execute_tool"` for tool execution spans | NIST AI RMF MAP-1.1 |

---

## Span Status

- `OK` — Inference completed successfully
- `ERROR` — Inference failed (with error description in span status)

---

## CoSAI WS2 Field Mapping

Cross-reference between CoSAI WS2 `AI_INTERACTION` field names and AITF attribute keys:

| CoSAI WS2 Field | AITF Attribute | Notes |
|---|---|---|
| `ai.model.vendor` | `gen_ai.provider.name` | OTel GenAI convention |
| `ai.model.name` | `gen_ai.request.model` | OTel GenAI convention |
| `ai.model.endpoint` | `server.address` | OTel standard |
| `ai.input.prompt` | `gen_ai.input.messages` | Emitted as span event |
| `ai.system_prompt.hash` | `gen_ai.system_prompt.hash` | SHA-256 hash |
| `ai.output.completion` | `gen_ai.output.messages` | Emitted as span event |
| `ai.config.temperature` | `gen_ai.request.temperature` | OTel GenAI convention |
| `ai.config.top_p` | `gen_ai.request.top_p` | OTel GenAI convention |
| `ai.usage.prompt_tokens` | `gen_ai.usage.input_tokens` | OTel GenAI convention |
| `ai.usage.completion_tokens` | `gen_ai.usage.output_tokens` | OTel GenAI convention |
| `ai.latency_ms` | `latency.total_ms` | Extension attribute |
| `ai.finish_reason` | `gen_ai.response.finish_reasons` | Array of reasons |

---

## Example

```
Span: chat claude-sonnet-4-5-20250929
  Kind: CLIENT
  Status: OK
  Attributes:
    gen_ai.provider.name: "anthropic"
    gen_ai.operation.name: "chat"
    gen_ai.request.model: "claude-sonnet-4-5-20250929"
    gen_ai.request.max_tokens: 4096
    gen_ai.request.temperature: 0.7
    gen_ai.system_prompt.hash: "sha256:a3f2b8..."
    gen_ai.response.id: "msg_abc123"
    gen_ai.response.model: "claude-sonnet-4-5-20250929"
    gen_ai.response.finish_reasons: ["end_turn"]
    gen_ai.usage.input_tokens: 150
    gen_ai.usage.output_tokens: 500
    latency.total_ms: 1250.0
    cost.total_cost: 0.0075
  Events:
    gen_ai.content.prompt: {gen_ai.input.messages: "Explain AITF"}
    gen_ai.content.completion: {gen_ai.output.messages: "AITF is..."}
```

## Span: `gen_ai.embeddings`

Represents an embedding generation request.

### Span Name

Format: `embeddings {gen_ai.request.model}`

### Span Kind

`CLIENT`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.provider.name` | string | **Required** | Provider identifier | MITRE ATLAS AML.T0044 |
| `gen_ai.operation.name` | string | **Required** | `"embeddings"` | NIST AI RMF MAP-1.1 |
| `gen_ai.request.model` | string | **Required** | Embedding model identifier | MITRE ATLAS AML.T0044, EU AI Act Art.13 |
| `gen_ai.request.encoding_formats` | string[] | **Optional** | Encoding formats (e.g. `["float"]`, `["base64"]`) | — |
| `gen_ai.embeddings.dimension.count` | int | **Optional** | Embedding dimensions | — |
| `gen_ai.usage.input_tokens` | int | **Required** | Tokens processed | OWASP LLM10, NIST AI RMF MEASURE-2.5 |
| `latency.total_ms` | double | **Required** | Total latency in milliseconds | NIST AI RMF MEASURE-2.5 |
| `cost.total_cost` | double | **Recommended** | Cost in USD | OWASP LLM10 |
