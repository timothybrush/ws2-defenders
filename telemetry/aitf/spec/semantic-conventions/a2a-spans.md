# A2A Span Conventions (Agent-to-Agent Protocol)

Status: **Normative** | OCSF Class: **API Activity (6003)** (`ai_operation` profile, OCSF v1.9.0)

AITF defines semantic conventions for the [Google A2A (Agent-to-Agent) protocol](https://a2a-protocol.org/), covering agent discovery via Agent Cards, task lifecycle (create, poll, cancel), message exchange (synchronous and streaming), and push notifications. A2A uses JSON-RPC 2.0 over HTTP(S) and enables cross-platform agent interoperability.

Key words "MUST", "SHOULD", "MAY" follow [RFC 2119](https://www.rfc-editor.org/rfc/rfc2119).

---

## Overview

```
A2A Protocol Flow:
  discover (Agent Card) -> message/send or message/stream -> tasks/get (poll) -> tasks/cancel

Spans:
  a2a.agent.discover       (fetch Agent Card)
  a2a.message.send         (synchronous message)
  a2a.message.stream       (SSE streaming)
  a2a.task.get             (poll task status)
  a2a.task.cancel          (cancel running task)
```

---

## Span: `a2a.agent.discover`

Represents fetching and parsing an A2A Agent Card from `/.well-known/agent.json`.

### Span Kind

`CLIENT`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `a2a.agent.url` | string | **Required** | Agent service endpoint URL | MITRE ATLAS [AML.T0040](https://atlas.mitre.org/techniques/AML.T0040), NIST AI RMF MAP-1.5 |
| `a2a.agent.name` | string | **Recommended** | Discovered agent name | OWASP LLM06 (Excessive Agency) |
| `a2a.agent.version` | string | **Recommended** | Agent version | NIST AI RMF MAP-1.1 |
| `a2a.agent.provider.organization` | string | **Optional** | Provider organization | EU AI Act Art.13 (Transparency) |
| `a2a.agent.skills` | string[] | **Recommended** | Skill IDs available | OWASP LLM06, NIST AI RMF MAP-1.1 |
| `a2a.agent.capabilities.streaming` | boolean | **Recommended** | Whether streaming is supported | — |
| `a2a.agent.capabilities.push_notifications` | boolean | **Optional** | Whether push notifications are supported | — |
| `a2a.protocol.version` | string | **Recommended** | A2A protocol version | NIST AI RMF MAP-1.1 |
| `a2a.transport` | string | **Recommended** | Transport: `"jsonrpc"`, `"grpc"`, `"http_json"` | MITRE ATLAS AML.T0040 |

---

## Span: `a2a.message.send`

Represents a synchronous `message/send` JSON-RPC call to a remote agent.

### Span Kind

`CLIENT`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `a2a.agent.name` | string | **Required** | Target agent name | OWASP LLM06, MITRE ATLAS [AML.T0048](https://atlas.mitre.org/techniques/AML.T0048) |
| `a2a.method` | string | **Required** | `"message/send"` | NIST AI RMF GOVERN-1.2 |
| `a2a.interaction_mode` | string | **Required** | `"sync"` | — |
| `a2a.task.id` | string | **Recommended** | Server-assigned task ID | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 |
| `a2a.task.context_id` | string | **Recommended** | Context ID grouping related tasks | NIST AI RMF GOVERN-1.2 |
| `a2a.task.state` | string | **Recommended** | Final task state (see Task States below) | OWASP LLM06, NIST AI RMF MEASURE-2.5 |
| `a2a.agent.url` | string | **Optional** | Agent endpoint URL | MITRE ATLAS AML.T0040 |
| `a2a.message.id` | string | **Recommended** | Outgoing message ID | NIST AI RMF GOVERN-1.2 |
| `a2a.message.role` | string | **Recommended** | `"user"` (for outgoing) | — |
| `a2a.message.parts_count` | int | **Optional** | Number of message parts | — |
| `a2a.task.artifacts_count` | int | **Optional** | Number of artifacts produced | — |
| `a2a.jsonrpc.error_code` | int | **Recommended** | JSON-RPC error code (if error) | NIST AI RMF MEASURE-2.5 |
| `a2a.jsonrpc.error_message` | string | **Recommended** | JSON-RPC error message (if error) | NIST AI RMF MEASURE-2.5 |

---

## Span: `a2a.message.stream`

Represents a streaming `message/stream` JSON-RPC call with SSE response.

### Span Kind

`CLIENT`

### Normative Field Table

Same as `a2a.message.send` plus:

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `a2a.interaction_mode` | string | **Required** | `"stream"` | — |
| `a2a.stream.events_count` | int | **Recommended** | Total SSE events received | NIST AI RMF MEASURE-2.5 |

### Events

#### `a2a.stream.event`

Emitted for each SSE event received.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `a2a.stream.event_type` | string | **Required** | `"status-update"` or `"artifact-update"` | — |
| `a2a.stream.is_final` | boolean | **Required** | Whether this is the terminal event | — |

---

## Span: `a2a.task.get`

Represents polling a task's status via `tasks/get`.

### Span Kind

`CLIENT`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `a2a.task.id` | string | **Required** | Task ID | NIST AI RMF GOVERN-1.2 |
| `a2a.method` | string | **Required** | `"tasks/get"` | — |
| `a2a.task.state` | string | **Recommended** | Current task state | NIST AI RMF MEASURE-2.5 |

---

## Span: `a2a.task.cancel`

Represents canceling a running task via `tasks/cancel`.

### Span Kind

`CLIENT`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `a2a.task.id` | string | **Required** | Task ID | NIST AI RMF GOVERN-1.2 |
| `a2a.method` | string | **Required** | `"tasks/cancel"` | — |
| `a2a.task.state` | string | **Recommended** | Task state after cancel | OWASP LLM06 |

---

## Task States

| Value | Description | Compliance |
|---|---|---|
| `submitted` | Task received, not yet started | — |
| `working` | Agent is actively processing | — |
| `input-required` | Agent needs additional input from client | EU AI Act Art.14 (Human Oversight) |
| `completed` | Successfully finished | — |
| `canceled` | Canceled by request | — |
| `failed` | Processing failed | NIST AI RMF MEASURE-2.5 |
| `rejected` | Agent refused the task | OWASP LLM06 |
| `auth-required` | Agent needs additional authentication | — |

---

## Event: `a2a.task.lifecycle.*` [RFC v0.4 gap closure]

> **RFC field:** A2A Task Lifecycle Event · **Tier:** SHOULD · **§12 Orchestration, Multi-Agent & Background Execution** · **Grounding attacks:** `AOC-04`, `AOC-11`, `AOC-09`, `TA-01` · **ODIS:** `delegation_id`, `parent_delegation_id` (6.3)

The spans above record what a client did: it sent a message, it polled, it cancelled. Each carries `a2a.task.state`, which is the state that span happened to observe. None of them records the *transition* — what the state was before, who caused the change, or how long the task has been alive.

That distinction matters because an A2A task is asynchronous and outlives the span that created it. A task submitted at 09:00 that never reaches a terminal state produces no error, fails no assertion, and triggers no alert. It simply stops being mentioned. Meanwhile the authority delegated to service it stays outstanding, so an orphaned task is an unbounded grant that nothing in the existing conventions makes visible. Equally, a task whose state changes with no identifiable actor is a control-plane compromise indicator, and `a2a.task.state` alone cannot express it.

### Emission Rule (RFC Appendix D.2)

Transitions MUST be emitted as events rather than only as attributes on whichever span observes them. Because a task outlives its creating span, its terminal event will usually belong to a *different trace*; `a2a.task.lifecycle.initiating_trace_id` is what makes that completion attributable, which is why it is Required rather than Recommended.

Every task MUST eventually emit an event with `a2a.task.lifecycle.terminal: true`. A task that passes `expected_terminal_by` without one is an orphan, and SHOULD be reported as such by emitting a lifecycle event with `event: "orphaned"`.

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `a2a.task.id` | string | Task the transition applies to |
| `a2a.task.lifecycle.event` | string | See Lifecycle Events below |
| `a2a.task.lifecycle.from_state` | string | State before the transition |
| `a2a.task.lifecycle.to_state` | string | State after the transition |
| `a2a.task.lifecycle.actor` | string | Identity that caused the transition |
| `a2a.task.lifecycle.at` | string | RFC 3339 timestamp of the transition |
| `a2a.task.lifecycle.terminal` | boolean | Whether `to_state` is terminal |
| `a2a.task.lifecycle.initiating_trace_id` | string | Trace that created the task (ODIS `request_trace_id`) |

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `a2a.task.lifecycle.transition_valid` | boolean | Whether the transition is permitted by the A2A state machine |
| `a2a.task.lifecycle.actor_type` | string | `"client_agent"`, `"remote_agent"`, `"human"`, `"system"`, `"timeout"`, `"unknown"` |
| `a2a.task.lifecycle.age_ms` | int | Elapsed time since task creation — the primary orphan signal |
| `a2a.task.lifecycle.expected_terminal_by` | string | RFC 3339 deadline for reaching a terminal state |
| `a2a.task.lifecycle.orphaned` | boolean | Passed `expected_terminal_by` without terminating |
| `a2a.task.lifecycle.delegated_scope` | string[] | Scopes outstanding while the task is live |
| `a2a.task.lifecycle.delegation_expires_at` | string | Expiry of the delegated authority backing the task |
| `a2a.task.lifecycle.initiating_span_id` | string | Span that created the task |
| `a2a.task.lifecycle.root_principal` | string | Principal ultimately accountable (ODIS `owner_ref`) |
| `a2a.task.lifecycle.failure_reason` | string | Reason for `"failed"` or `"rejected"` |

### Optional Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `a2a.task.lifecycle.cancel_requested_by` | string | Principal that requested cancellation |
| `a2a.task.lifecycle.input_required_reason` | string | Why the task entered `"input-required"` |
| `a2a.task.lifecycle.transition_count` | int | Cumulative transitions for this task |
| `a2a.task.lifecycle.poll_count` | int | `tasks/get` polls observed |
| `a2a.task.lifecycle.resubscribe_count` | int | Times a client re-attached to the task stream |
| `a2a.task.lifecycle.subscriber` | string | Principal attached to the stream, where it differs from `actor` |

### Lifecycle Events

| Value | Description |
|-------|-------------|
| `submitted` | Task created and accepted by the remote agent |
| `state_changed` | Ordinary state machine transition |
| `streamed` | An SSE event was delivered to a subscriber (read-side) |
| `polled` | A `tasks/get` observed the task (read-side) |
| `resubscribed` | A client re-attached to the task stream (read-side) |
| `input_required` | Task paused awaiting client input |
| `artifact_produced` | Task emitted an artifact |
| `cancelled` | Task cancelled by request |
| `push_config_changed` | Push notification configuration was set or altered |
| `terminal` | Task reached a terminal state |
| `orphaned` | Task passed its expected terminal deadline |
| `resumed` | Task resumed after `input-required` or `auth-required` |

The three read-side events — `streamed`, `polled`, `resubscribed` — change no task state. They are recorded because they reveal who is watching the task and from where, which is the only way to distinguish a task nobody is monitoring from one being monitored by an unexpected principal.

### Push Notification Configuration

> **RFC field:** A2A Task Lifecycle Event (push-notification surface) · **Tier:** SHOULD · **§12** · **Grounding attacks:** `AOC-04`, `AOC-11`, `AOC-09`, `TA-01`

Push notification configuration is the callback surface of an asynchronous task: it determines where the remote agent will deliver results, and therefore where results can be redirected. It belongs with the lifecycle because a configuration change *mid-task* is an exfiltration primitive — the work was authorised to one destination and delivered to another.

| Attribute | Type | Notes |
|-----------|------|-------|
| `a2a.push.config.url` | string | Registered callback URL |
| `a2a.push.config.url_hash` | string | Digest of the callback URL, for change detection without exposing the endpoint |
| `a2a.push.config.changed` | boolean | Callback target changed after task creation |
| `a2a.push.config.authenticated` | boolean | Whether the callback is authenticated |
| `a2a.push.config.scheme` | string | `"none"`, `"bearer"`, `"hmac"`, `"mtls"` |
| `a2a.push.config.in_allowlist` | boolean | Destination is on the configured egress allowlist |
| `a2a.push.config.set_by` | string | Principal that registered or altered the configuration |

`a2a.push.config.changed: true` on a live task SHOULD be treated as an attempted exfiltration until cleared, and `in_allowlist: false` alongside it removes the ambiguity. `authenticated: false` means anything that can reach the endpoint can impersonate the remote agent's results.

### Reading the Fields

The orphan case is the one that existing conventions cannot express at all. A task with `age_ms` well past `expected_terminal_by`, no terminal event, and a non-empty `delegated_scope` is authority left standing with no work attached to it — and `delegation_expires_at` says whether the grant will lapse on its own or persist indefinitely. The inverse, a task still running past `delegation_expires_at`, is work continuing without the authority that justified it. Both are findings; neither produces an error.

`transition_valid: false` means the remote agent moved the task through a transition the A2A state machine does not permit. That is either a broken peer or a peer whose task state is being manipulated by something other than the protocol.

`actor_type: "unknown"` on a state change is worth alerting on in its own right. Under the RFC §16 provenance rule an unmarked actor is self-asserted, and a transition that cannot be attributed to *any* principal means the control plane governing the task is not fully observed.

Where `subscriber` differs from `actor`, or `resubscribe_count` climbs with differing principals, the task stream is being read by someone other than the client that created it.

---

## Example: A2A Agent Discovery + Message Send

```
Span: a2a.agent.discover
  a2a.agent.url: "https://research-agent.example.com"
  a2a.agent.name: "research-assistant"
  a2a.agent.version: "1.2.0"
  a2a.agent.skills: ["web_search", "summarize", "translate"]
  a2a.agent.capabilities.streaming: true
  a2a.protocol.version: "0.2.5"
  a2a.transport: "jsonrpc"

Span: a2a.message.send research-assistant
  a2a.method: "message/send"
  a2a.interaction_mode: "sync"
  a2a.task.id: "task-abc123"
  a2a.task.context_id: "ctx-xyz789"
  a2a.task.state: "completed"
  a2a.task.artifacts_count: 1
  Events:
    a2a.task.state_change: {a2a.task.state: "submitted"}
    a2a.message: {message_id: "msg-001", role: "user", parts_count: 1}
    a2a.task.state_change: {a2a.task.state: "completed"}
```

---

## References

- [A2A Protocol Specification](https://a2a-protocol.org/latest/specification/)
- [A2A GitHub Repository](https://github.com/a2aproject/A2A)
- [A2A Protocol Definitions](https://a2a-protocol.org/latest/definitions/)
