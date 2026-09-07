# Agent Span Conventions (AGENT_TRACE)

> **OTel Alignment Note:** This spec adopts OTel GenAI agent attributes (`gen_ai.agent.{name,id,description,version}`,
> `gen_ai.conversation.id`) where OTel now defines them. Extension attributes for step reasoning, delegation,
> team orchestration, and memory use short namespaces (e.g., `agent.step.*`, `agent.team.*`, `memory.*`).

Status: **Normative** | CoSAI WS2 Alignment: **AGENT_TRACE** | OCSF Class: **API Activity (6003)** (`ai_operation` profile, OCSF v1.9.0)

AITF defines comprehensive semantic conventions for AI agent telemetry, supporting single agents, multi-agent orchestration, delegation, and agent memory. This specification defines the normative field requirements aligned with CoSAI Working Stream 2 (Telemetry for AI) and mapped to applicable compliance and threat frameworks.

Key words "MUST", "SHOULD", "MAY" follow [RFC 2119](https://www.rfc-editor.org/rfc/rfc2119).

---

## Span: `agent.session`

Represents a complete agent session from start to finish.

### Span Name

Format: `agent.session {gen_ai.agent.name}`

### Span Kind

`INTERNAL`

### Normative Field Table

Instrumentors MUST emit all Required fields. Instrumentors SHOULD emit Recommended fields when the data is available. Optional fields MAY be emitted for enhanced observability.

#### Agent Identity

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.agent.name` | string | **Required** | Human-readable agent name | OWASP LLM06 (Excessive Agency), MITRE ATLAS [AML.T0048](https://atlas.mitre.org/techniques/AML.T0048) |
| `gen_ai.agent.id` | string | **Required** | Unique agent instance identifier | OWASP LLM06, MITRE ATLAS [AML.T0048](https://atlas.mitre.org/techniques/AML.T0048), EU AI Act Art.12 |
| `gen_ai.conversation.id` | string | **Required** | Session identifier | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 |
| `agent.workflow_id` | string | **Recommended** | Workflow/DAG identifier linking related agent sessions | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 |

#### Agent Configuration

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `agent.type` | string | **Recommended** | Agent type: `"conversational"`, `"autonomous"`, `"reactive"`, `"proactive"` | EU AI Act Art.13 (Transparency) |
| `agent.framework` | string | **Recommended** | Agent framework: `"langchain"`, `"crewai"`, `"autogen"`, `"semantic_kernel"`, `"custom"` | NIST AI RMF MAP-1.1 |
| `gen_ai.agent.version` | string | **Optional** | Agent version string | NIST AI RMF MAP-1.1 |
| `gen_ai.agent.description` | string | **Optional** | Agent role/purpose description | EU AI Act Art.13 |

#### Session State (CoSAI WS2)

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `agent.state` | string | **Recommended** | Current agent lifecycle state: `"initializing"`, `"planning"`, `"executing"`, `"waiting"`, `"completed"`, `"failed"`, `"suspended"` | OWASP LLM06 (Excessive Agency), MITRE ATLAS [AML.T0048](https://atlas.mitre.org/techniques/AML.T0048) |
| `agent.session.turn_count` | int | **Recommended** | Total turns completed in session | OWASP LLM10 (Unbounded Consumption) |
| `agent.session.start_time` | string | **Optional** | Session start timestamp (ISO 8601) | EU AI Act Art.12 |

#### Team Affiliation

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `agent.team.name` | string | **Optional** | Team name (if part of a multi-agent team) | NIST AI RMF MAP-1.1 |
| `agent.team.id` | string | **Optional** | Team identifier | NIST AI RMF GOVERN-1.2 |

---

## Span: `agent.step`

Represents a single step in the agent's execution loop (think-act-observe).

### Span Name

Format: `agent.step.{agent.step.type} {gen_ai.agent.name}`

### Span Kind

`INTERNAL`

### Normative Field Table

#### Step Identification

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.agent.name` | string | **Required** | Agent name | OWASP LLM06, MITRE ATLAS AML.T0048 |
| `agent.step.type` | string | **Required** | Step type (see Step Types table below) | OWASP LLM06, NIST AI RMF MEASURE-2.5 |
| `agent.step.index` | int | **Required** | Step sequence number (0-indexed) | NIST AI RMF GOVERN-1.2 |

#### ReAct / Chain-of-Thought (CoSAI WS2)

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `agent.step.thought` | string | **Recommended** | Agent's internal reasoning (ReAct "thought") | NIST AI RMF MAP-1.1, EU AI Act Art.13 (Transparency), MITRE ATLAS [AML.T0048](https://atlas.mitre.org/techniques/AML.T0048) |
| `agent.step.action` | string | **Recommended** | Planned/executed action | OWASP LLM06 (Excessive Agency) |
| `agent.step.observation` | string | **Recommended** | Result/observation from action | NIST AI RMF MEASURE-2.5 |
| `agent.scratchpad` | string | **Optional** | Accumulated agent scratchpad / working memory state (JSON) | OWASP LLM02 (Sensitive Info Disclosure), MITRE ATLAS [AML.T0048](https://atlas.mitre.org/techniques/AML.T0048) |
| `agent.next_action` | string | **Recommended** | Next planned action (forward-looking intent) | OWASP LLM06 (Excessive Agency), MITRE ATLAS [AML.T0048](https://atlas.mitre.org/techniques/AML.T0048) |
| `agent.step.status` | string | **Recommended** | Step outcome: `"success"`, `"error"`, `"retry"`, `"skipped"` | NIST AI RMF MEASURE-2.5 |

### Step Types

| Value | Description | Compliance |
|---|---|---|
| `planning` | Agent is planning next actions | EU AI Act Art.13 |
| `reasoning` | Agent is reasoning about observations | EU AI Act Art.13, NIST AI RMF MAP-1.1 |
| `tool_use` | Agent is calling a tool/function | OWASP LLM06, MITRE ATLAS AML.T0048 |
| `delegation` | Agent is delegating to another agent | OWASP LLM06, NIST AI RMF GOVERN-1.7 |
| `response` | Agent is generating final response | OWASP LLM05 |
| `reflection` | Agent is reflecting on its performance | NIST AI RMF MEASURE-2.5 |
| `memory_access` | Agent is accessing memory | OWASP LLM02 |
| `guardrail_check` | Agent is checking guardrails | OWASP LLM01–LLM10 |
| `human_in_loop` | Agent is waiting for human input | EU AI Act Art.14 (Human Oversight) |
| `error_recovery` | Agent is recovering from an error | NIST AI RMF MEASURE-2.5 |

---

## Span: `agent.delegation`

Represents agent-to-agent delegation within a multi-agent system.

### Span Name

Format: `agent.delegate {gen_ai.agent.name} -> {agent.delegation.target_agent}`

### Span Kind

`INTERNAL`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.agent.name` | string | **Required** | Delegating agent name | OWASP LLM06, MITRE ATLAS AML.T0048 |
| `agent.delegation.target_agent` | string | **Required** | Target agent name | OWASP LLM06, MITRE ATLAS AML.T0048 |
| `agent.delegation.target_agent_id` | string | **Required** | Target agent instance ID | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 |
| `agent.delegation.reason` | string | **Recommended** | Why delegation occurred | EU AI Act Art.13 (Transparency) |
| `agent.delegation.strategy` | string | **Recommended** | Strategy: `"round_robin"`, `"capability"`, `"hierarchical"`, `"vote"` | NIST AI RMF GOVERN-1.7 |
| `agent.delegation.task` | string | **Recommended** | Delegated task description | OWASP LLM06 |
| `agent.delegation.result` | string | **Optional** | Result from delegated agent | NIST AI RMF MEASURE-2.5 |
| `agent.delegation.timeout_ms` | double | **Optional** | Delegation timeout in milliseconds | OWASP LLM10 |

### Peer Agent Card / Descriptor [RFC v0.4 gap closure]

> **RFC field:** Peer Agent Card / Descriptor · **Tier:** SHOULD · **§12 Orchestration, Multi-Agent & Background Execution** · **Grounding attacks:** `AOC-09`, `AOC-11`, `AOC-08`, `AOC-16` · **ODIS:** `agent_id`, `approved_software_refs` (6.1)

`agent.delegation.target_agent` and `.target_agent_id` record *who the delegating agent believes it is talking to*. They are strings the peer supplied, and they are trivially stable while everything behind them changes. These fields record the counterparty's **declared descriptor as presented at this contact**, change detection against prior contacts, and — the field that carries most of the weight — the outcome of any verification attempted against it.

This group is protocol-neutral. `a2a.agent.*` in [`a2a-spans.md`](a2a-spans.md) carries the same material for A2A specifically, and implementations SHOULD populate both when the peer is reached over A2A, so that cross-protocol peer analysis remains possible in a fleet that speaks more than one.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.agent.peer.id` | string | **Required** | Peer identifier as declared by the peer | NIST AI RMF GOVERN-1.2 |
| `gen_ai.agent.peer.name` | string | **Required** | Peer name as declared | NIST AI RMF MAP-1.1 |
| `gen_ai.agent.peer.verification.result` | string | **Required** | `"verified"`, `"failed"`, `"refused"`, `"unverified"`, `"not_attempted"` | MITRE ATLAS [AML.T0048](https://atlas.mitre.org/techniques/AML.T0048), NIST AI RMF GOVERN-1.2 |
| `gen_ai.agent.peer.url` | string | **Recommended** | Endpoint at which the peer was contacted | MITRE ATLAS [AML.T0040](https://atlas.mitre.org/techniques/AML.T0040) |
| `gen_ai.agent.peer.version` | string | **Recommended** | Declared peer version | NIST AI RMF MAP-1.1 |
| `gen_ai.agent.peer.provider` | string | **Recommended** | Declared provider/operator organization | NIST AI RMF MAP-1.5 |
| `gen_ai.agent.peer.skills` | string[] | **Recommended** | Advertised skills/capabilities | OWASP LLM06 (Excessive Agency) |
| `gen_ai.agent.peer.protocol` | string | **Recommended** | `"a2a"`, `"mcp"`, `"acp"`, `"anp"`, `"http"`, `"custom"` | — |
| `gen_ai.agent.peer.card.hash` | string | **Recommended** | SHA-256 of the canonicalised descriptor as presented at this contact | MITRE ATLAS [AML.T0040](https://atlas.mitre.org/techniques/AML.T0040) |
| `gen_ai.agent.peer.card.baseline_hash` | string | **Recommended** | Descriptor hash recorded at the previous contact or at approval | MITRE ATLAS [AML.T0040](https://atlas.mitre.org/techniques/AML.T0040) |
| `gen_ai.agent.peer.card.changed` | boolean | **Recommended** | Whether the descriptor differs from the baseline | MITRE ATLAS [AML.T0040](https://atlas.mitre.org/techniques/AML.T0040), OWASP LLM06 |
| `gen_ai.agent.peer.verification.method` | string | **Recommended** | `"signature"`, `"registry_lookup"`, `"inspection_request"`, `"mtls"`, `"spiffe"`, `"none"` | MITRE ATLAS [AML.T0048](https://atlas.mitre.org/techniques/AML.T0048) |
| `gen_ai.agent.peer.approved` | boolean | **Recommended** | Whether this peer identity is in the approved-counterparty set | OWASP LLM06, NIST AI RMF GOVERN-1.2 |
| `gen_ai.agent.peer.card.change_fields` | string[] | **Optional** | Which fields changed: `"name"`, `"url"`, `"version"`, `"provider"`, `"skills"` | OWASP LLM06 |
| `gen_ai.agent.peer.card.first_seen` | string | **Optional** | ISO 8601 timestamp of first contact with this peer identity | EU AI Act Art.12 |
| `gen_ai.agent.peer.verification.authority` | string | **Optional** | Which authority performed the verification | NIST AI RMF MAP-1.5 |

**Reading the fields.** `verification.result` carries a five-value enum rather than a boolean because the distinctions matter. `"not_attempted"` means the system never asked — a configuration finding, not a peer finding. `"unverified"` means it asked and got no answer it could evaluate. `"refused"` is the interesting one: the peer *declined* an inspection request, which is a positive act by a counterparty and should be treated very differently from silence. Under the RFC §16 provenance rule, everything in this table except `verification.result` and `verification.authority` is self-asserted by the peer, so an unmarked descriptor is the peer's claim about itself and nothing more.

`card.changed: true` with `approved: true` is the peer-agent form of a rug-pull: the counterparty was approved on the strength of a descriptor it no longer presents. `card.change_fields` says whether that was a version bump or a change of `url` and `skills`, which are the two fields whose change means the peer is now a different thing.

---

## Span: `agent.team.orchestrate`

Represents a multi-agent team orchestration operation.

### Span Name

Format: `agent.team.orchestrate {agent.team.name}`

### Span Kind

`INTERNAL`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `agent.team.name` | string | **Required** | Team name | NIST AI RMF MAP-1.1 |
| `agent.team.id` | string | **Required** | Team identifier | NIST AI RMF GOVERN-1.2 |
| `agent.team.topology` | string | **Required** | Team topology (see table below) | NIST AI RMF GOVERN-1.7 |
| `agent.team.members` | string[] | **Recommended** | Member agent names | OWASP LLM06 |
| `agent.team.coordinator` | string | **Recommended** | Coordinator agent name | OWASP LLM06 |
| `agent.team.task` | string | **Optional** | Team task description | EU AI Act Art.13 |
| `agent.team.consensus_method` | string | **Optional** | `"majority"`, `"unanimous"`, `"coordinator"` | NIST AI RMF GOVERN-1.7 |
| `agent.team.rounds` | int | **Optional** | Number of interaction rounds | OWASP LLM10 |

### Team Topologies

| Value | Description | Compliance |
|---|---|---|
| `hierarchical` | Manager/supervisor delegates to workers | NIST AI RMF GOVERN-1.7 |
| `peer` | Agents collaborate as equals | NIST AI RMF GOVERN-1.7 |
| `pipeline` | Sequential processing chain | — |
| `consensus` | Agents vote or reach consensus | EU AI Act Art.14 |
| `debate` | Agents debate to reach conclusion | NIST AI RMF MEASURE-2.5 |
| `swarm` | Dynamic self-organizing agents | OWASP LLM06 |

---

## Span: `agent.memory`

Represents an agent memory operation.

### Span Name

Format: `agent.memory.{memory.operation} {gen_ai.agent.name}`

### Span Kind

`INTERNAL`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.agent.name` | string | **Required** | Agent name | OWASP LLM06 |
| `memory.operation` | string | **Required** | Operation: `"store"`, `"retrieve"`, `"update"`, `"delete"`, `"search"` | OWASP LLM02 (Sensitive Info) |
| `memory.store` | string | **Required** | Store type: `"short_term"`, `"long_term"`, `"episodic"`, `"semantic"`, `"procedural"` | NIST AI RMF MAP-1.5 |
| `memory.key` | string | **Recommended** | Memory key/identifier | OWASP LLM02 |
| `memory.hit` | boolean | **Recommended** | Whether memory was found (for retrieve) | NIST AI RMF MEASURE-2.5 |
| `memory.ttl_seconds` | int | **Optional** | Time to live in seconds | — |
| `memory.provenance` | string | **Optional** | Origin of memory entry | NIST AI RMF MAP-1.5 |

### Declared Memory Configuration [RFC v0.4 gap closure]

> **RFC field:** Declared Memory Configuration · **Tier:** MAY · **§10 Memory** · **Grounding attacks:** `AOC-05`, `AOC-07`, `IR-02`

Memory poisoning is dangerous because it *persists*: an injected instruction written into long-term memory affects sessions the attacker never touches. How long it persists, who may write, and whether writes cross session or tenant boundaries are properties of the declared configuration, not of any single memory operation — a `memory.store` span alone cannot be assessed for durability of impact. These attributes are configuration-time and belong on the agent-registration record (`agent.registration` / `asset.registered`); they SHOULD NOT be repeated on every memory span. `memory.config.name` is the join key back to the `memory.store` value carried on operation spans.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `memory.config.enabled` | boolean | **Optional** | Whether persistent memory is enabled for this agent | OWASP LLM04, NIST AI RMF MAP-1.5 |
| `memory.config.name` | string | **Optional** | Declared name of the memory store — join key to `memory.store` on operation spans | NIST AI RMF MAP-1.1, EU AI Act Art.12 |
| `memory.config.types` | string[] | **Optional** | Types configured: `"short_term"`, `"long_term"`, `"episodic"`, `"semantic"`, `"procedural"` | NIST AI RMF MAP-1.1 |
| `memory.config.backend` | string | **Optional** | Backing store: `"in_process"`, `"redis"`, `"postgres"`, `"vector_db"`, `"file"`, `"managed_service"` | NIST AI RMF MAP-1.5 |
| `memory.config.scope` | string | **Optional** | Sharing scope: `"turn"`, `"session"`, `"user"`, `"agent"`, `"tenant"`, `"global"`. `"global"` means one user's injected content can reach another's | OWASP LLM04, GDPR Art.5, MITRE ATLAS AML.T0020 |
| `memory.config.persistence` | string | **Optional** | Durability: `"ephemeral"`, `"session"`, `"durable"` | OWASP LLM04, EU AI Act Art.12 |
| `memory.config.retention_seconds` | int | **Optional** | Configured retention period — how long a poisoned entry survives | GDPR Art.5, OWASP LLM04 |
| `memory.config.max_entries` | int | **Optional** | Cap on stored entries | NIST AI RMF MEASURE-2.5 |
| `memory.config.max_bytes` | int | **Optional** | Declared size cap in bytes; baseline for observed footprint | NIST AI RMF MEASURE-2.5 |
| `memory.config.retrieval.top_k` | int | **Optional** | Entries returned per read. Widening it increases the reach of any single poisoned entry | OWASP LLM04, NIST AI RMF MEASURE-2.5 |
| `memory.config.retrieval.scoring` | string | **Optional** | Selection function: `"cosine"`, `"dot"`, `"euclidean"`, `"bm25"`, `"recency"`, `"hybrid"` | NIST AI RMF MEASURE-2.5, MITRE ATLAS AML.T0043 |
| `memory.config.retrieval.min_score` | double | **Optional** | Relevance floor. A floor of zero admits arbitrary entries into context | OWASP LLM04, NIST AI RMF MEASURE-2.5 |
| `memory.config.write_principals` | string[] | **Optional** | Principals permitted to write — the poisoning surface | MITRE ATLAS AML.T0020, OWASP LLM04 |
| `memory.config.agent_writable` | boolean | **Optional** | Whether the agent may write its own memory without human or policy review | OWASP LLM06, EU AI Act Art.14 |
| `memory.config.write_review` | string | **Optional** | Review applied to writes: `"none"`, `"policy"`, `"guardrail"`, `"human"` | EU AI Act Art.14, OWASP LLM04 |
| `memory.config.cross_session` | boolean | **Optional** | Whether memory written in one session is readable in another | OWASP LLM04, GDPR Art.5 |
| `memory.config.cross_tenant` | boolean | **Optional** | Whether memory crosses tenant boundaries. Expected `false`; `true` is a finding in multi-tenant deployments | GDPR Art.32, EU AI Act Art.10, OWASP LLM02 |
| `memory.config.classification` | string | **Optional** | Classification permitted in memory: `"public"`, `"internal"`, `"confidential"`, `"restricted"` | GDPR Art.32, EU AI Act Art.10 |
| `memory.config.encryption_at_rest` | boolean | **Optional** | Whether memory contents are encrypted at rest | GDPR Art.32, SOC 2 CC6.1 |
| `memory.config.profile_ref` | string | **Optional** | Reference to the governing memory policy profile (ODIS `policy_profile_ref`) | NIST AI RMF GOVERN-1.2 |
| `memory.config.hash` | string | **Optional** | Digest of the declared configuration, for detecting silent change | MITRE ATLAS AML.T0010 |

#### Reading the Fields

The blast radius of a single poisoned write is read from the configuration, not the write. `scope: "global"` or `cross_tenant: true` means content one principal injected is reachable by principals who never interacted with the attacker. `persistence: "durable"` with a large or absent `retention_seconds` means the entry outlives the incident and the remediation window; `IR-02` scoping depends on it.

`agent_writable: true` combined with `write_review: "none"` is the configuration in which an agent that has been prompt-injected once can persist that injection itself, with no mediating control — the fields state plainly whether a self-reinforcing loop is possible.

`retrieval.min_score: 0` with a high `top_k` means low-relevance entries are admitted into context on most reads, so a poisoned entry does not need to be topically relevant to be retrieved.

Per §16, these are declarations. A change in `memory.config.hash` between registration records is the signal that the declaration itself moved.

---

## CoSAI WS2 Field Mapping

Cross-reference between CoSAI WS2 `AGENT_TRACE` field names and AITF attribute keys:

| CoSAI WS2 Field | AITF Attribute | Notes |
|---|---|---|
| `agent.id` | `gen_ai.agent.id` | Direct match |
| `agent.workflow_id` | `agent.workflow_id` | New in CoSAI WS2 alignment |
| `agent.state` | `agent.state` | New in CoSAI WS2 alignment |
| `agent.thought` | `agent.step.thought` | Set on step spans |
| `agent.scratchpad` | `agent.scratchpad` | New in CoSAI WS2 alignment |
| `agent.next_action` | `agent.next_action` | New in CoSAI WS2 alignment |

---

## Example: Multi-Agent Research System

```
Span: agent.team.orchestrate research-team
  agent.team.topology: "hierarchical"
  agent.team.members: ["manager", "researcher", "writer"]
  |
  +- Span: agent.session manager
  |    gen_ai.agent.id: "agent-mgr-001"
  |    agent.type: "autonomous"
  |    agent.framework: "crewai"
  |    agent.workflow_id: "wf-research-abc123"
  |    agent.state: "executing"
  |    |
  |    +- Span: agent.step.planning manager
  |    |    agent.step.thought: "Need to research AI telemetry"
  |    |    agent.next_action: "delegate to researcher"
  |    |    +- Span: chat gpt-4o
  |    |
  |    +- Span: agent.step.delegation manager
  |    |    agent.delegation.target_agent: "researcher"
  |    |    agent.delegation.reason: "Research expertise needed"
  |    |    |
  |    |    +- Span: agent.session researcher
  |    |         agent.workflow_id: "wf-research-abc123"
  |    |         +- Span: agent.step.tool_use researcher
  |    |         |    +- Span: mcp.tool.invoke read_file
  |    |         +- Span: agent.step.reasoning researcher
  |    |              agent.scratchpad: "{\"findings\": [...]}"
  |    |              +- Span: chat claude-sonnet-4-5-20250929
  |    |
  |    +- Span: agent.step.delegation manager
  |         agent.delegation.target_agent: "writer"
  |         |
  |         +- Span: agent.session writer
  |              agent.workflow_id: "wf-research-abc123"
  |              +- Span: agent.step.response writer
  |                   +- Span: chat gpt-4o
```
