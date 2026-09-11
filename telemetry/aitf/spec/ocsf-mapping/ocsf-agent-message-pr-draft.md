# OCSF PR (standalone) — `agent_message`: one generic object for agent-to-agent communication

> **Rescoped 2026-09-05.** This draft asked for **one object and one class**, the class living in an
> `ai` category that was expected to land via
> [issue #1640](https://github.com/ocsf/ocsf-schema/issues/1640). **OCSF v1.9.0 (2026-08-03) shipped
> the objects and the profile but created no category** — `categories.json` still stops at `uid 8`.
> **The class half of this PR is therefore withdrawn.** What remains, and what is now fileable, is the
> `agent_message` **object**, proposed as an addition to the released `ai_operation` profile, sitting
> beside the `message_context` attribute v1.9.0 already shipped. Agent-to-agent exchanges are carried
> on **API Activity (`6003`)**.
>
> Two consequences, both good for this proposal. It **no longer depends on anything** — the `ai_agent`
> and `delegation` objects it reuses are released, so it can be filed immediately rather than queued
> behind a category. And the "one generic object, not one per protocol" argument below is *reinforced*
> by v1.9.0's own choice to model agentic activity as a profile over existing classes rather than as a
> bespoke class hierarchy. That is the same instinct applied one level up.

**Title:**

```
Add a generic agent_message object to the ai_operation profile for A2A/ACP/ANP/MCP
```

---

## Description

Adds **one** object — `agent_message` — to the released `ai_operation` profile,
to represent agent-to-agent communication across protocols (A2A, ACP, ANP,
MCP, …) with a `protocol_id` **discriminator**, rather than a dedicated
object per protocol. No new class and no new category: agent-to-agent exchanges
are API calls, and they are carried on **API Activity (`6003`)** with
`ai_operation` applied, exactly as v1.9.0 intends.

## Motivation — one generic object, not one per protocol

OCSF gives SMTP/SMB/DNS/SSH/TLS dedicated objects/classes because those
protocols are mature, ubiquitous, semantically distinct, and have decades-old
detection ecosystems. The new agentic protocols are the opposite:

| Per-protocol modeling pays off when… | A2A / ACP / ANP are… |
|---|---|
| protocol is stable for years | revving monthly |
| protocols are semantically distinct | converging on one model |
| large existing detection corpus | greenfield |

Their conceptual core is identical: **a source agent talks to a peer agent,
under some delegated authority, about a unit of work (task / run / message)
that has a lifecycle status, via a transport, carrying parts/artifacts, and may
error.** A per-protocol object would (a) fragment cross-protocol detection —
a SOC wants "agent contacted an *untrusted* peer" regardless of wire protocol —
and (b) force schema churn on every protocol revision.

This mirrors OCSF's own deeper pattern: a generic class carrying a protocol
object + id (e.g. `network_activity` with `tls` / `dns_query`). Per-protocol
detail that doesn't generalize stays in a `metadata` map.

The same argument is why this is now an object rather than a class. v1.9.0
declined to give agents their own category, choosing instead to describe them
with a profile applied to the classes that already exist. An agent-to-agent
call *is* an API call; what makes it interesting to a defender is the peer, the
authority, and the unit of work — all of which are object-shaped. Adding a
class would have asserted that the transport is what's novel. It isn't.

## Proposed changes

### New object — `objects/agent_message.json`

```json
{
  "caption": "Agent Message",
  "name": "agent_message",
  "description": "A communication between two AI agents over an agentic protocol (A2A, ACP, ANP, MCP, ...). One generic object discriminated by protocol_id rather than a per-protocol object.",
  "extends": "object",
  "attributes": {
    "protocol_id": {
      "requirement": "recommended",
      "enum": {
        "0": { "caption": "Unknown" },
        "1": { "caption": "A2A" },
        "2": { "caption": "ACP" },
        "3": { "caption": "ANP" },
        "4": { "caption": "MCP" },
        "99": { "caption": "Other" }
      }
    },
    "protocol":        { "requirement": "optional", "description": "Protocol name; caption of protocol_id." },
    "protocol_version":{ "requirement": "optional" },
    "direction_id": {
      "requirement": "recommended",
      "enum": {
        "0": { "caption": "Unknown" },
        "1": { "caption": "Request" },
        "2": { "caption": "Response" },
        "3": { "caption": "Stream" },
        "4": { "caption": "Notification" },
        "99": { "caption": "Other" }
      }
    },
    "role_id": {
      "requirement": "optional",
      "enum": {
        "0": { "caption": "Unknown" }, "1": { "caption": "Client" },
        "2": { "caption": "Server" }, "99": { "caption": "Other" }
      }
    },
    "operation":   { "requirement": "optional", "description": "Method/capability invoked (e.g. A2A method, ACP run mode, ANP meta-protocol)." },
    "unit_uid":    { "requirement": "recommended", "description": "Normalized id of the unit of work (task/run/message)." },
    "unit_type":   { "requirement": "optional", "description": "task | run | message." },
    "status_id": {
      "requirement": "recommended",
      "description": "Canonical lifecycle status (normalized across protocols).",
      "enum": {
        "0":  { "caption": "Unknown" },
        "1":  { "caption": "Submitted" },
        "2":  { "caption": "Working" },
        "3":  { "caption": "Input Required" },
        "4":  { "caption": "Completed" },
        "5":  { "caption": "Failed" },
        "6":  { "caption": "Canceling" },
        "7":  { "caption": "Canceled" },
        "99": { "caption": "Other" }
      }
    },
    "previous_status_id": { "requirement": "optional", "description": "Prior status_id (same enum)." },
    "src_agent":   { "requirement": "recommended", "description": "Initiating agent (ai_agent)." },
    "dst_agent":   { "requirement": "recommended", "description": "Peer/target agent (ai_agent)." },
    "delegation":  { "requirement": "optional", "description": "Authorization context for the call (delegation object, #1640)." },
    "parts_count":     { "requirement": "optional" },
    "part_types":      { "requirement": "optional" },
    "artifacts_count": { "requirement": "optional" },
    "transport":   { "requirement": "optional", "description": "jsonrpc | grpc | http | sse | ws." },
    "src_endpoint":{ "requirement": "optional", "description": "Initiator endpoint (network_endpoint)." },
    "dst_endpoint":{ "requirement": "optional", "description": "Peer endpoint (network_endpoint)." },
    "trust_domain":     { "requirement": "optional" },
    "peer_trust_domain":{ "requirement": "optional" },
    "is_cross_domain":  { "requirement": "optional" },
    "peer_did":    { "requirement": "optional", "description": "Peer decentralized identifier (ANP)." },
    "error_code":  { "requirement": "optional" },
    "error_message":{ "requirement": "optional" },
    "duration":    { "requirement": "optional", "description": "Duration in milliseconds." },
    "metadata":    { "requirement": "optional", "description": "Protocol-specific fields that do not generalize." }
  }
}
```

`src_agent` / `dst_agent` reuse the **`ai_agent`** object released in v1.9.0;
`src_endpoint` / `dst_endpoint` reuse **`network_endpoint`**; `delegation`
reuses the **`delegation`** object released in v1.9.0.

### Profile update — `profiles/ai_operation.json`

Add one optional attribute alongside the four v1.9.0 shipped (`ai_agent`,
`ai_model`, `delegation`, `message_context`):

```json
{
  "attributes": {
    "agent_message": { "requirement": "optional", "description": "An agent-to-agent communication carried by this operation." }
  }
}
```

> **Relationship to `message_context`.** The two are complementary rather than
> overlapping. `message_context` correlates an operation to a conversation;
> `agent_message` describes the *exchange itself* — who the peer is, under what
> authority, which unit of work, what lifecycle state. A deployment emitting
> both can answer "which conversation" and "which peer, and was it trusted"
> from one event.

**Withdrawn: the `agent_communication` class.** An earlier revision of this
draft proposed a dedicated class in an `ai` category. Since v1.9.0 created no
such category and modelled agentic activity as a profile instead, the class is
withdrawn and the `activity_id` enum it carried (Send / Receive / Stream /
Notify) is expressed as `agent_message.operation` on API Activity (`6003`).

### Dictionary additions — `dictionary.json`

```json
{
  "attributes": {
    "agent_message":    { "caption": "Agent Message", "description": "An agent-to-agent communication.", "type": "agent_message" },
    "protocol_version": { "caption": "Protocol Version", "description": "Version of the agentic protocol.", "type": "string_t" },
    "unit_uid":         { "caption": "Unit UID", "description": "Normalized unit-of-work id (task/run/message).", "type": "string_t" },
    "unit_type":        { "caption": "Unit Type", "description": "task | run | message.", "type": "string_t" },
    "peer_did":         { "caption": "Peer DID", "description": "Peer decentralized identifier.", "type": "string_t" },
    "peer_trust_domain":{ "caption": "Peer Trust Domain", "description": "Trust domain of the peer agent.", "type": "string_t" },
    "part_types":       { "caption": "Part Types", "description": "Content part media types.", "type": "string_t", "is_array": true },
    "parts_count":      { "caption": "Parts Count", "type": "integer_t" },
    "artifacts_count":  { "caption": "Artifacts Count", "type": "integer_t" }
  }
}
```

> Reuse existing dictionary attributes where present: `protocol`/`protocol_id`,
> `direction`/`direction_id`, `role_id`, `operation`, `status_id`,
> `previous_status_id`, `src_agent`/`dst_agent`, `delegation`, `transport`,
> `src_endpoint`/`dst_endpoint`, `trust_domain`, `is_cross_domain`,
> `error_code`/`error_message`, `duration`, `metadata`. Only add the entries
> above that are not already defined on `main`.

## Cross-protocol status normalization (informative)

Producers map protocol-native lifecycle states onto the canonical `status_id`:

| canonical `status_id` | A2A `task.state` | ACP `run.status` |
|---|---|---|
| Submitted (1) | `submitted` | `created` |
| Working (2) | `working` | `in-progress` |
| Input Required (3) | `input-required`, `auth-required` | `awaiting` |
| Completed (4) | `completed` | `completed` |
| Failed (5) | `failed`, `rejected` | `failed` |
| Canceling (6) | — | `cancelling` |
| Canceled (7) | `canceled` | `cancelled` |

## Backwards compatibility

Additive: one new object plus one optional attribute on an existing profile.
Every attribute is optional. No existing object, class, profile attribute or
category changes, and nothing that consumes v1.9.0 today is affected.

## Testing

- [ ] Schema validation passes with the new object and the profile addition.
- [ ] Server renders `agent_message` and its enums on classes carrying `ai_operation`.
- [ ] `dictionary.json` references resolve.
- [ ] A round-trip sample validates on API Activity (`6003`) with `ai_operation` applied.

## Checklist

- [ ] `objects/agent_message.json` added.
- [ ] `profiles/ai_operation.json` updated with the optional `agent_message` attribute.
- [ ] `dictionary.json` updated.
- [ ] `CHANGELOG.md` updated.
- [ ] Linked to #1640 and #1641, noted as building on released v1.9.0 objects — **no dependency**.
