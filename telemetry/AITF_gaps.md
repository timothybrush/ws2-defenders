# AITF gap closure: RFC v0.4 Appendix C resolution log

Companion to **Secure by Design: Telemetry Field Selection for AI Systems**, Working Draft v0.4
(<https://github.com/cosai-oasis/ws2-defenders/blob/main/telemetry/CoSAI-AI-Telemetry-RFC.md>).

This file previously listed **27 open gaps** extracted from RFC Appendix C — conceptual fields in the
RFC with no corresponding attribute in [AITF](https://github.com/cosai-oasis/ws2-defenders/tree/main/telemetry).
**All 27 are now closed.** This is the resolution log: what was added, where it landed, and how to
verify it.

| | |
|---|---|
| Gaps closed | **27 of 27** |
| Attributes added | **390** |
| Registry sections added | **29** (tagged `[RFC v0.4 gap closure]`) |
| Convention-document sections added | **23** across 8 documents, plus 1 new document |
| New events defined | **7** |
| New spans defined | **1** (`identity.credential.mint`) |
| SDK constants added | **390** across 12 attribute groups, plus **18** enum value sets, in **all four SDKs** |
| SDK version | **0.4.0** (Python, Go, TypeScript, Rust), matching RFC Working Draft v0.4 |

Three gaps were recorded in Appendix C as *"none identified — a new attribute or namespace is
required"*. Those namespace decisions are recorded under [New namespace decisions](#new-namespace-decisions)
below.

## Verification

Every claim in this log is checkable from the repository:

```bash
# 29 registry sections, one per closed gap group
grep -c '^### .*\[RFC v0.4 gap closure\]' spec/semantic-conventions/attributes-registry.md

# 23 convention-document sections
grep -rn '^#\{2,4\} .*\[RFC v0.4 gap closure\]' spec/semantic-conventions/*.md \
  | grep -v attributes-registry | wc -l
```

Every one of the 390 attribute names defined in a gap-closure registry section is reachable as an SDK
constant. This asserts it, and exits non-zero if any name is missing:

```bash
PYTHONPATH=sdk/python/src python3 - <<'PY'
import re, pathlib, aitf.semantic_conventions as s

names, in_gap = [], False
for line in pathlib.Path(
        "spec/semantic-conventions/attributes-registry.md").read_text().splitlines():
    m = re.match(r"^#{2,4}\s+(.*)$", line)
    if m:
        in_gap = "[RFC v0.4 gap closure]" in m.group(1)
    elif in_gap and line.startswith("|"):
        cell = line.strip("|").split("|")[0].strip()
        if cell.startswith("`") and re.fullmatch(r"`[a-z][\w.]*`", cell):
            names.append(cell.strip("`"))

values, stack = set(), [getattr(s, n) for n in s.__all__]
while stack:
    for k, v in vars(stack.pop()).items():
        if k.startswith("_"):
            continue
        values.add(v) if isinstance(v, str) else \
            stack.append(v) if isinstance(v, type) else None

missing = sorted(set(names) - values)
print(f"registry gap attributes: {len(set(names))}   missing from SDK: {len(missing)}")
assert not missing, missing
PY
```

The same 390 names are carried by the Go, TypeScript, and Rust SDKs. This asserts parity across all
four, and exits non-zero if any SDK is short:

```bash
PYTHONPATH=sdk/python/src python3 - <<'PY'
import re, pathlib, aitf.semantic_conventions as s

names, in_gap = set(), False
for line in pathlib.Path(
        "spec/semantic-conventions/attributes-registry.md").read_text().splitlines():
    m = re.match(r"^#{2,4}\s+(.*)$", line)
    if m:
        in_gap = "[RFC v0.4 gap closure]" in m.group(1)
    elif in_gap and line.startswith("|"):
        cell = line.strip("|").split("|")[0].strip()
        if re.fullmatch(r"`[a-z][\w.]*`", cell):
            names.add(cell.strip("`"))

def read(p):
    return pathlib.Path(p).read_text()

values, stack = set(), [getattr(s, n) for n in s.__all__]
while stack:
    for k, v in vars(stack.pop()).items():
        if k.startswith("_"):
            continue
        values.add(v) if isinstance(v, str) else \
            stack.append(v) if isinstance(v, type) else None

have = {
    "python": values,
    "go": set(re.findall(r'attribute\.Key\("([^"]+)"\)',
                         read("sdk/go/semconv/attributes.go"))),
    "typescript": set(re.findall(r'"([^"]+)"',
                      read("sdk/typescript/src/semantic-conventions/attributes.ts"))),
    "rust": set(re.findall(r'pub const \w+: &str = "([^"]+)";',
                           read("sdk/rust/src/semconv.rs"))),
}

bad = 0
for lang in ("python", "go", "typescript", "rust"):
    miss = sorted(names - have[lang])
    bad += len(miss)
    print(f"{lang:11} missing {len(miss):3} of {len(names)}  {miss[:3]}")
assert not bad
PY
```

## Resolution index

| # | RFC field | Tier | § | Namespace delivered | Attrs |
|---:|:---|:---:|:---:|:---|---:|
| 1 | [Session / Turn / Step IDs](#1-session--turn--step-ids) | MUST | 5 | [`gen_ai.turn.*`](spec/semantic-conventions/attributes-registry.md#gen_aiturn--gen_aistep-rfc-v04-gap-closure) | 6 |
| 2 | [Trigger Type & Source Event](#2-trigger-type--source-event) | MUST | 5 | [`gen_ai.trigger.*`](spec/semantic-conventions/attributes-registry.md#gen_aitrigger-rfc-v04-gap-closure) | 7 |
| 3 | [Content Modality & Attachment Identity](#3-content-modality--attachment-identity) | MUST ‡ | 6 | [`gen_ai.content.*`](spec/semantic-conventions/attributes-registry.md#gen_aicontent-rfc-v04-gap-closure) | 13 |
| 4 | [Peer Agent Card / Descriptor](#4-peer-agent-card--descriptor) | SHOULD | 12 | [`gen_ai.agent.peer.*`](spec/semantic-conventions/attributes-registry.md#gen_aiagentpeer-rfc-v04-gap-closure) | 16 |
| 5 | [Backend / Route Restriction Decision](#5-backend--route-restriction-decision) | SHOULD | 16 | [`gen_ai.route.*`](spec/semantic-conventions/attributes-registry.md#gen_airoute-rfc-v04-gap-closure) | 10 |
| 6 | [Guardrail Modification Record](#6-guardrail-modification-record) | SHOULD ‡ | 6 | [`security.guardrail.modification.*`](spec/semantic-conventions/attributes-registry.md#securityguardrailmodification-rfc-v04-gap-closure) | 11 |
| 7 | [Encoded / Obfuscated Payload Indicator](#7-encoded--obfuscated-payload-indicator) | MAY | 6 | [`security.obfuscation.*`](spec/semantic-conventions/attributes-registry.md#securityobfuscation-rfc-v04-gap-closure) | 8 |
| 8 | [Authorization Decision Record](#8-authorization-decision-record) | MUST | 16 | [`security.authorization.*`](spec/semantic-conventions/attributes-registry.md#securityauthorization-rfc-v04-gap-closure) | 17 |
| 9 | [Session Taint Labels & Information-Flow Decisions](#9-session-taint-labels--information-flow-decisions) | SHOULD | 16 | [`security.taint.*`](spec/semantic-conventions/attributes-registry.md#securitytaint-rfc-v04-gap-closure) | 11 |
| 10 | [Enforcement-Point Availability & Failure Mode](#10-enforcement-point-availability--failure-mode) | SHOULD | 15 | [`security.enforcement.*`](spec/semantic-conventions/attributes-registry.md#securityenforcement-rfc-v04-gap-closure) | 10 |
| 11 | [Tool Definition Digest](#11-tool-definition-digest) | SHOULD | 9 | [`mcp.tool.definition.*`](spec/semantic-conventions/attributes-registry.md#mcptooldefinition-rfc-v04-gap-closure) | 10 |
| 12 | [MCP Server Identity & Primitive](#12-mcp-server-identity--primitive) | SHOULD | 9 | [`mcp.primitive`](spec/semantic-conventions/attributes-registry.md#mcpprimitive-and-server-identity-rfc-v04-gap-closure) | 11 |
| 13 | [Protocol Envelope Capture](#13-protocol-envelope-capture) | MAY | 12 | [`mcp.envelope.*`](spec/semantic-conventions/attributes-registry.md#mcpenvelope-rfc-v04-gap-closure) | 11 |
| 14 | [Declared Memory Configuration](#14-declared-memory-configuration) | MAY | 10 | [`memory.config.*`](spec/semantic-conventions/attributes-registry.md#memoryconfig-rfc-v04-gap-closure) | 21 |
| 15 | [Citations / Source Attribution](#15-citations--source-attribution) | MUST | 7 | [`rag.citation.*`](spec/semantic-conventions/attributes-registry.md#ragcitation-rfc-v04-gap-closure) | 12 |
| 16 | [Declared Knowledge-Source Configuration](#16-declared-knowledge-source-configuration) | MAY | 11 | [`rag.source.*`](spec/semantic-conventions/attributes-registry.md#ragsource-rfc-v04-gap-closure) | 22 |
| 17 | [Trust-Domain Crossing & Delegation Depth](#17-trust-domain-crossing--delegation-depth) | SHOULD | 13 | [`identity.boundary.*`](spec/semantic-conventions/attributes-registry.md#identityboundary-rfc-v04-gap-closure) | 15 |
| 18 | [Credential Minting & Scope-Narrowing Check](#18-credential-minting--scope-narrowing-check) | SHOULD | 13 | [`identity.credential.mint.*`](spec/semantic-conventions/attributes-registry.md#identitycredentialmint-rfc-v04-gap-closure) | 24 |
| 19 | [Human Approval / Elicitation Event](#19-human-approval--elicitation-event) | SHOULD | 16 | [`identity.approval.*`](spec/semantic-conventions/attributes-registry.md#identityapproval-rfc-v04-gap-closure) | 26 |
| 20 | [Organization / Tenant ID](#20-organization--tenant-id) | SHOULD | 5 | [`asset.tenant.*`](spec/semantic-conventions/attributes-registry.md#assettenant-and-tenancy-scoping-rfc-v04-gap-closure) · [`security.tenant.*`](spec/semantic-conventions/attributes-registry.md#securitytenant-rfc-v04-gap-closure) | 13 |
| 21 | [Capability-Set Change Event](#21-capability-set-change-event) | MUST | 14 | [`asset.capability.*`](spec/semantic-conventions/attributes-registry.md#assetcapability-rfc-v04-gap-closure) | 16 |
| 22 | [Instrumentation Coverage / Hook Attestation](#22-instrumentation-coverage--hook-attestation) | SHOULD | 15 | [`asset.instrumentation.*`](spec/semantic-conventions/attributes-registry.md#assetinstrumentation-rfc-v04-gap-closure) | 17 |
| 23 | [Execution Environment / Sandbox](#23-execution-environment--sandbox) | MUST | 9 | [`supply_chain.runtime.*`](spec/semantic-conventions/attributes-registry.md#supply_chainruntime-rfc-v04-gap-closure) | 27 |
| 24 | [A2A Task Lifecycle Event](#24-a2a-task-lifecycle-event) | SHOULD | 12 | [`a2a.task.lifecycle.*`](spec/semantic-conventions/attributes-registry.md#a2atasklifecycle-rfc-v04-gap-closure) · [`a2a.push.config.*`](spec/semantic-conventions/attributes-registry.md#a2apushconfig-rfc-v04-gap-closure) | 30 |
| 25 | [Attribute Source / Trusted-Provenance Marking](#25-attribute-source--trusted-provenance-marking) | MUST ‡ | 16 | [`security.attribute_source.*`](spec/semantic-conventions/attributes-registry.md#securityattribute_source-rfc-v04-gap-closure) | 6 |
| 26 | [Mediation Coverage & Bypass Path](#26-mediation-coverage--bypass-path) | SHOULD | 16 | [`security.mediation.*`](spec/semantic-conventions/attributes-registry.md#securitymediation-rfc-v04-gap-closure) | 8 |
| 27 | [Event Sequence Continuity](#27-event-sequence-continuity) | SHOULD | 15 | [`observability.sequence.*`](spec/semantic-conventions/attributes-registry.md#observabilitysequence-rfc-v04-gap-closure) | 12 |

## New namespace decisions

Appendix C recorded no AITF landing space for three fields. The namespaces chosen, and why:

**`security.attribute_source.*`** for Attribute Source / Trusted-Provenance Marking. This is a
property of a security *assertion* rather than of an identity, so it belongs beside the other
enforcement attributes rather than in `identity.*`. It is also the direct AITF expression of the
`attribute_source` enum the RFC asks OCSF to add in Appendix E.2 item 4, so keeping the name aligned
keeps the upstream ask legible.

**`security.mediation.*`** for Mediation Coverage & Bypass Path. Placed alongside the other
reference-monitor attributes, since it describes the monitor's coverage rather than any single
decision the monitor made.

**`observability.*`** — a new top-level namespace — for Event Sequence Continuity. This describes the
integrity of the telemetry channel itself rather than the behaviour of the system under observation,
and overloading `security.*` with it would conflate "the agent did something suspicious" with "the
record of what the agent did may be incomplete." Those are different findings with different
responses.

## Cross-cutting reading rule

Per RFC §16, an **unmarked security attribute is a self-asserted attribute**. Every group below that
carries a security verdict is only as trustworthy as its
[`security.attribute_source.*`](spec/semantic-conventions/security-cross-cutting.md#attribute-source--trusted-provenance-marking)
marking. Per Appendix D.5, none of these values may be propagated in W3C `baggage`; a value received
over the wire is recorded as `"self_asserted"` regardless of how the sender labelled it.

---

## Closed gaps

### 1. Session / Turn / Step IDs

- **Tier:** MUST · **RFC §5** · **Grounding attacks:** `AOC-03`, `AOC-07`, `TA-07`, `TA-08`
- **Namespace:** `gen_ai.turn.*` — **6 attributes**
- **Registry:** [`gen_ai.turn.*`](spec/semantic-conventions/attributes-registry.md#gen_aiturn--gen_aistep-rfc-v04-gap-closure)
- **Conventions:** [`gen-ai-spans.md` § Identifier Hierarchy](spec/semantic-conventions/gen-ai-spans.md#identifier-hierarchy-rfc-v04-gap-closure)
- **Python SDK:** `GenAIAttributes.TURN_*` / `STEP_*` / `RUN_ID`

Adds the three-level execution hierarchy beneath the run. `gen_ai.turn.id` is **Required**; `turn.index` orders turns within a conversation and `step.id`/`step.parent_id` resolve one action within a turn. A detection can now name *which turn* behaviour changed rather than only which run, which is what `AOC-03` and `TA-07` require to be scoped.

### 2. Trigger Type & Source Event

- **Tier:** MUST · **RFC §5** · **Grounding attacks:** `TA-01`, `AOC-04`, `AOC-10`, `AOC-12`
- **Namespace:** `gen_ai.trigger.*` — **7 attributes**
- **Registry:** [`gen_ai.trigger.*`](spec/semantic-conventions/attributes-registry.md#gen_aitrigger-rfc-v04-gap-closure)
- **Conventions:** [`gen-ai-spans.md` § Trigger Provenance](spec/semantic-conventions/gen-ai-spans.md#trigger-provenance-rfc-v04-gap-closure)
- **Python SDK:** `GenAIAttributes.TRIGGER_*`, `.TriggerType`

Separates *who or what started the run* from the channel it arrived on. `trigger.type` is **Required**; `trigger.event`, `event.id`, `source` and `source.principal` are **Conditionally Required** whenever the type is not `user_initiated`. `human_in_loop` records whether a person was in the path at all. An autonomous run whose originating event is an inbound email is now distinguishable from a user typing the same request.

### 3. Content Modality & Attachment Identity

- **Tier:** MUST ‡ · **RFC §6** · **Grounding attacks:** `AOC-12`, `AOC-05`, `TA-05`, `TA-01`
- **Namespace:** `gen_ai.content.*` — **13 attributes**
- **Registry:** [`gen_ai.content.*`](spec/semantic-conventions/attributes-registry.md#gen_aicontent-rfc-v04-gap-closure)
- **Conventions:** [`gen-ai-spans.md` § Content Modality & Attachment Identity](spec/semantic-conventions/gen-ai-spans.md#content-modality--attachment-identity-rfc-v04-gap-closure)
- **Python SDK:** `GenAIAttributes.CONTENT_*`

Records, for every content-bearing field, the part type, MIME type, and for files the name, size and content hash. Instructions arriving as an image, PDF or structured blob are invisible to text-only inspection and text-only logging; `content.attachment.extracted_text_hash` pairs with `security.obfuscation.*` so that OCR-recovered text is correlatable across sessions without retaining it. Per Appendix D.4 the content itself is emitted on events — only the bounded hash and count attributes are promoted to spans.

### 4. Peer Agent Card / Descriptor

- **Tier:** SHOULD · **RFC §12** · **Grounding attacks:** `AOC-09`, `AOC-11`, `AOC-08`, `AOC-16`
- **ODIS field:** `agent_id`, `approved_software_refs` (6.1)
- **Namespace:** `gen_ai.agent.peer.*` — **16 attributes**
- **Registry:** [`gen_ai.agent.peer.*`](spec/semantic-conventions/attributes-registry.md#gen_aiagentpeer-rfc-v04-gap-closure)
- **Conventions:** [`agent-spans.md` § Peer Agent Card / Descriptor](spec/semantic-conventions/agent-spans.md#peer-agent-card--descriptor-rfc-v04-gap-closure)
- **Python SDK:** `AgentAttributes.PEER_*`, `.PeerVerificationResult`

Captures the counterparty's declared descriptor as presented at contact — name, URL, version, provider, advertised skills — with change detection against prior contacts and, critically, the outcome of any verification attempted. The `verification.result` enum distinguishes `not_attempted` (a configuration finding), `unverified` (verification ran but produced no evaluable answer) and `refused` (the peer declined inspection, which is a positive act rather than silence). `card.changed: true` with `approved: true` is the peer-agent rug-pull.

### 5. Backend / Route Restriction Decision

- **Tier:** SHOULD · **RFC §16** · **Grounding attacks:** `AOC-05`, `AOC-06`, `TA-01`
- **ODIS field:** `resource_indicators`, `constraints` (6.3)
- **Namespace:** `gen_ai.route.*` — **10 attributes**
- **Registry:** [`gen_ai.route.*`](spec/semantic-conventions/attributes-registry.md#gen_airoute-rfc-v04-gap-closure)
- **Conventions:** [`gen-ai-spans.md` § Backend / Route Restriction Decision](spec/semantic-conventions/gen-ai-spans.md#backend--route-restriction-decision-rfc-v04-gap-closure)
- **Python SDK:** `GenAIAttributes.ROUTE_*`, `.RouteDecision`

Records where an operation was allowed to execute: the candidate set, the constraint that narrowed it, the selection made, and `no_candidate_action` — the behaviour when nothing qualified. That last field is the fail-open detector. Distinct from `model_ops.serving.route.*`, which is the operations view of the same routing; this is the policy view.

### 6. Guardrail Modification Record

- **Tier:** SHOULD ‡ · **RFC §6** · **Grounding attacks:** `TA-01`, `AOC-12`, `AOC-03`, `IR-01`
- **Namespace:** `security.guardrail.modification.*` — **11 attributes**
- **Registry:** [`security.guardrail.modification.*`](spec/semantic-conventions/attributes-registry.md#securityguardrailmodification-rfc-v04-gap-closure)
- **Conventions:** [`security-cross-cutting.md` § Guardrail Modification Record](spec/semantic-conventions/security-cross-cutting.md#guardrail-modification-record)
- **Python SDK:** `SecurityAttributes.GUARDRAIL_MODIFICATION_*`

Closes the case where an enforcement point **rewrites rather than blocks**. `security.guardrail.result` alone is sufficient only when the options are allow and block; in practice payloads are masked, redacted, stripped or normalised and let through. Recording `modified`, the enforcement point, the side, and a before/after digest means the payload the model saw can be reconciled with the payload that was logged. `modified: true` with `action: "pass"` is the specific case this exists for.

### 7. Encoded / Obfuscated Payload Indicator

- **Tier:** MAY · **RFC §6** · **Grounding attacks:** `AOC-12`, `TA-03`
- **Namespace:** `security.obfuscation.*` — **8 attributes**
- **Registry:** [`security.obfuscation.*`](spec/semantic-conventions/attributes-registry.md#securityobfuscation-rfc-v04-gap-closure)
- **Conventions:** [`security-cross-cutting.md` § Encoded / Obfuscated Payload Indicator](spec/semantic-conventions/security-cross-cutting.md#encoded--obfuscated-payload-indicator)
- **Python SDK:** `SecurityAttributes.OBFUSCATION_*`

Adds a flag plus the decoded form when input contains base64, image-embedded (OCR) or markup "authority" tags. `inspected_after_decode: false` is the finding — a guardrail that evaluated the wrapper rather than the contents. `decoded_hash` is the durable signal, correlating the same payload across sessions and encodings without retaining content, which is what makes a campaign visible when each instance looks like a one-off.

### 8. Authorization Decision Record

- **Tier:** MUST · **RFC §16** · **Grounding attacks:** `AOC-02`, `AOC-08`, `AOC-10`, `TA-08`, `TA-12`, `TA-13`
- **ODIS field:** `policy_profile_ref` (6.1)
- **Namespace:** `security.authorization.*` — **17 attributes**
- **Registry:** [`security.authorization.*`](spec/semantic-conventions/attributes-registry.md#securityauthorization-rfc-v04-gap-closure)
- **Conventions:** [`security-cross-cutting.md` § Authorization Decision Record](spec/semantic-conventions/security-cross-cutting.md#authorization-decision-record)
- **Python SDK:** `SecurityAttributes.AUTHORIZATION_*`, `.AuthorizationDecision`

Records, per mediated operation, the decision, the deny reason and its machine-readable code, the deciding authority (inline rule vs external PDP, and which engine), the rule identifier, and any obligations attached. Explicitly distinct from a guardrail verdict: that is a classifier score, this is the authorization outcome. Six independent Appendix A attacks require it, so under the Appendix E.3 promotion rule this is an OCSF standardization ask (`ai_authorization`), not merely an AITF-local attribute.

### 9. Session Taint Labels & Information-Flow Decisions

- **Tier:** SHOULD · **RFC §16** · **Grounding attacks:** `AOC-03`, `TA-01`, `TA-02`, `TA-05`
- **ODIS field:** `constraints.data_classification` (6.3)
- **Namespace:** `security.taint.*` — **11 attributes**
- **Registry:** [`security.taint.*`](spec/semantic-conventions/attributes-registry.md#securitytaint-rfc-v04-gap-closure)
- **Conventions:** [`security-cross-cutting.md` § Session Taint Labels](spec/semantic-conventions/security-cross-cutting.md#session-taint-labels--information-flow-decisions)
- **Python SDK:** `SecurityAttributes.TAINT_*`, `.TaintFlowDecision`

Records the information-flow labels in force, their scope, what applied each, and — the field that matters most — when an operation is denied **because of accumulated taint rather than anything in its own payload**. Without `denied_labels`, an analyst reviewing such a denial sees a false positive. Declassification must name `declassified_by`; the agent declassifying its own session is self-service removal of the label that existed to constrain it.

### 10. Enforcement-Point Availability & Failure Mode

- **Tier:** SHOULD · **RFC §15** · **Grounding attacks:** `TA-01`, `TA-10`, `IR-01`, `AOC-12`
- **Namespace:** `security.enforcement.*` — **10 attributes**
- **Registry:** [`security.enforcement.*`](spec/semantic-conventions/attributes-registry.md#securityenforcement-rfc-v04-gap-closure)
- **Conventions:** [`security-cross-cutting.md` § Enforcement-Point Availability](spec/semantic-conventions/security-cross-cutting.md#enforcement-point-availability--failure-mode)
- **Python SDK:** `SecurityAttributes.ENFORCEMENT_*`, `.EnforcementFailureMode`

For each enforcement callout: whether it was reached, its latency, and on failure whether the system failed open or failed closed, plus the action taken anyway. `failure_mode: "fail_open"` converts a silent gap into a known one — the operation proceeded without the control, and any statistic counting it as clean is wrong. A rising rate of `timeout` against one enforcement point is the `TA-10` pattern: an attacker need not defeat a guardrail that can be made unavailable.

### 11. Tool Definition Digest

- **Tier:** SHOULD · **RFC §9** · **Grounding attacks:** `AOC-10`, `AOC-14`, `TA-06`, `AOC-09`
- **ODIS field:** `approved_software_refs` (6.1)
- **Namespace:** `mcp.tool.definition.*` — **10 attributes**
- **Registry:** [`mcp.tool.definition.*`](spec/semantic-conventions/attributes-registry.md#mcptooldefinition-rfc-v04-gap-closure)
- **Conventions:** [`mcp-spans.md` § Tool Definition Digest](spec/semantic-conventions/mcp-spans.md#tool-definition-digest-rfc-v04-gap-closure)
- **Python SDK:** `MCPAttributes.TOOL_DEFINITION_*`

Hashes the tool's declared contract **as presented at invocation time** — name, description, argument and output schema — canonicalised per RFC 8785 (JCS) and compared against the approved baseline. `definition.changed: true` with `approved: true` is the rug-pull: a tool whose definition mutated after approval. `change_type` distinguishes a description edit (the prompt-injection vector) from a schema change.

### 12. MCP Server Identity & Primitive

- **Tier:** SHOULD · **RFC §9** · **Grounding attacks:** `TA-06`, `AOC-10`, `TA-01`, `AOC-09`, `TA-12`, `TA-13`
- **Namespace:** `mcp.primitive` — **11 attributes**
- **Registry:** [`mcp.primitive`](spec/semantic-conventions/attributes-registry.md#mcpprimitive-and-server-identity-rfc-v04-gap-closure)
- **Conventions:** [`mcp-spans.md` § Primitive & Server Identity](spec/semantic-conventions/mcp-spans.md#cross-cutting-primitive--server-identity-rfc-v04-gap-closure)
- **Python SDK:** `MCPAttributes.PRIMITIVE`, `.SERVER_*`

Adds server name, version, transport, endpoint and instance identity, plus `mcp.primitive` — which of `tool`, `resource`, `prompt`, `sampling`, `elicitation` or `roots` was exercised. Sampling and elicitation are the primitives that let a server drive the client, and without this attribute they were indistinguishable from ordinary tool traffic. `server.identity.verified: false` marks the whole server record as self-asserted under §16.

### 13. Protocol Envelope Capture

- **Tier:** MAY · **RFC §12** · **Grounding attacks:** `AOC-09`, `AOC-12`, `TA-06`
- **Namespace:** `mcp.envelope.*` — **11 attributes**
- **Registry:** [`mcp.envelope.*`](spec/semantic-conventions/attributes-registry.md#mcpenvelope-rfc-v04-gap-closure)
- **Conventions:** [`mcp-spans.md` § Protocol Envelope Capture](spec/semantic-conventions/mcp-spans.md#protocol-envelope-capture-rfc-v04-gap-closure)
- **Python SDK:** `MCPAttributes.ENVELOPE_*`

Preserves the raw MCP / A2A JSON-RPC envelope alongside the interpreted fields, retaining protocol-level detail that framework-level abstraction discards. Tiered **MAY** with explicit size, truncation and redaction attributes, because capture is a retention decision rather than a default. `envelope.captured: false` states that the raw form was not retained, so an investigation knows the interpreted fields are all there is.

### 14. Declared Memory Configuration

- **Tier:** MAY · **RFC §10** · **Grounding attacks:** `AOC-05`, `AOC-07`, `IR-02`
- **Namespace:** `memory.config.*` — **21 attributes**
- **Registry:** [`memory.config.*`](spec/semantic-conventions/attributes-registry.md#memoryconfig-rfc-v04-gap-closure)
- **Conventions:** [`agent-spans.md` § Declared Memory Configuration](spec/semantic-conventions/agent-spans.md#declared-memory-configuration-rfc-v04-gap-closure)
- **Python SDK:** `MemoryAttributes.CONFIG_*`, `.ConfigScope`

Records the memory store's declared identity and limits: backend, scope, persistence, retention, size caps, retrieval spec, write principals, and whether writes cross session or tenant boundaries. Memory poisoning is dangerous because it *persists*, and durability of impact cannot be read from a `memory.store` span alone. `agent_writable: true` with `write_review: "none"` is the configuration in which a once-injected agent can persist its own injection.

### 15. Citations / Source Attribution

- **Tier:** MUST · **RFC §7** · **Grounding attacks:** `TA-01`, `TA-02`, `IR-03`, `AOC-03`
- **Namespace:** `rag.citation.*` — **12 attributes**
- **Registry:** [`rag.citation.*`](spec/semantic-conventions/attributes-registry.md#ragcitation-rfc-v04-gap-closure)
- **Conventions:** [`rag-spans.md` § Citation & Source Attribution](spec/semantic-conventions/rag-spans.md#citation--source-attribution-rfc-v04-gap-closure)
- **Python SDK:** `RAGAttributes.CITATION_*`

Records the sources the agent *claims* it drew on, per output, **plus whether each resolves to an item actually returned by a logged retrieval event**. `citation.fabricated`, `unresolved_count` and `unresolved_ids` are the signal; `citation.retrieval_span_id` is the join that makes resolution checkable rather than assumed. `uncited_content_ratio` bounds how much of an output claims no source at all.

### 16. Declared Knowledge-Source Configuration

- **Tier:** MAY · **RFC §11** · **Grounding attacks:** `TA-09`, `IR-03`, `TA-02`
- **Namespace:** `rag.source.*` — **22 attributes**
- **Registry:** [`rag.source.*`](spec/semantic-conventions/attributes-registry.md#ragsource-rfc-v04-gap-closure)
- **Conventions:** [`rag-spans.md` § Declared Knowledge-Source Configuration](spec/semantic-conventions/rag-spans.md#declared-knowledge-source-configuration-rfc-v04-gap-closure)
- **Python SDK:** `RAGAttributes.SOURCE_*`

Records each knowledge source's declared identity and contract — name, description, index/collection identity, schema and hash, trust level, classification, write access, and search parameters including top-k, filters, scoring and reranker. `source.undeclared_access: true` records retrieval from a source that was never declared, which is the config-drift case the RFC's §11 asks about.

### 17. Trust-Domain Crossing & Delegation Depth

- **Tier:** SHOULD · **RFC §13** · **Grounding attacks:** `AOC-04`, `AOC-09`, `AOC-11`, `TA-11`
- **ODIS field:** `trust_domain` (6.1), `max_depth` (6.3)
- **Namespace:** `identity.boundary.*` — **15 attributes**
- **Registry:** [`identity.boundary.*`](spec/semantic-conventions/attributes-registry.md#identityboundary-rfc-v04-gap-closure)
- **Conventions:** [`identity-spans.md` § Trust-Domain Crossing & Delegation Depth](spec/semantic-conventions/identity-spans.md#trust-domain-crossing--delegation-depth-rfc-v04-gap-closure)
- **Python SDK:** `IdentityAttributes.BOUNDARY_*`, `.DELEGATION_*`

Records the counterparty's trust domain and the delegation depth at this hop, plus whether either crossed a configured limit. Authority leaving the domain that issued it, and the number of hops from the originating principal, were both previously unrecoverable from AITF telemetry. `delegation.root_principal` anchors the chain to the human who started it, so `TA-11` attribution survives multiple hops.

### 18. Credential Minting & Scope-Narrowing Check

- **Tier:** SHOULD · **RFC §13** · **Grounding attacks:** `AOC-02`, `AOC-08`, `IR-04`, `TA-08`
- **ODIS field:** `granted_authorizations`, `binding_profile` (6.2/6.3)
- **Namespace:** `identity.credential.mint.*` — **24 attributes**
- **Registry:** [`identity.credential.mint.*`](spec/semantic-conventions/attributes-registry.md#identitycredentialmint-rfc-v04-gap-closure)
- **Conventions:** [`identity-spans.md` § Span `identity.credential.mint`](spec/semantic-conventions/identity-spans.md#span-identitycredentialmint-rfc-v04-gap-closure)
- **Python SDK:** `IdentityAttributes.CREDENTIAL_MINT_*`, `.CREDENTIAL_SCOPE_*`

Adds a dedicated `identity.credential.mint` span for the credential-exchange event at each hop: grant type, whose identity the minted token represents, audience, issuer, lifetimes, and the **requested-vs-granted scope delta verified after minting**. `scope.forwarded_unchanged: true` detects an inbound token passed on without narrowing; `scope.escalation: true` detects a child credential broader than its parent. Both were previously invisible.

### 19. Human Approval / Elicitation Event

- **Tier:** SHOULD · **RFC §16** · **Grounding attacks:** `AOC-01`, `AOC-02`, `AOC-07`, `AOC-11`
- **ODIS field:** `originating_principal` (6.3, partial)
- **Namespace:** `identity.approval.*` — **26 attributes**
- **Registry:** [`identity.approval.*`](spec/semantic-conventions/attributes-registry.md#identityapproval-rfc-v04-gap-closure)
- **Conventions:** [`identity-spans.md` § Human Approval / Elicitation](spec/semantic-conventions/identity-spans.md#human-approval--elicitation-rfc-v04-gap-closure) · [`events.md` § `identity.approval.requested` / `.decided`](spec/semantic-conventions/events.md#identityapprovalrequested-rfc-v04-gap-closure)
- **Python SDK:** `IdentityAttributes.APPROVAL_*`, `.ApprovalDecision`

Adds the out-of-band approval lifecycle as a **pair of events** — `identity.approval.requested` and `identity.approval.decided` — joined by `identity.approval.id`, plus span attributes mirroring the outcome. The emission rule states that the **absence of a `decided` event MUST NOT be read as a denial**. `scope_binding.result: "mismatch"` is the consent-substitution finding: what was approved is not what was executed. Approver identity is recorded as verified by the identity provider, not as reported by the agent.

### 20. Organization / Tenant ID

- **Tier:** SHOULD · **RFC §5** · **Grounding attacks:** `AOC-02`, `AOC-05`, `TA-05`, `TA-01`, `TA-11`
- **ODIS field:** `trust_domain` (6.1)
- **Namespace:** `asset.tenant.*` · `security.tenant.*` — **13 attributes**
- **Registry:** [`asset.tenant.*`](spec/semantic-conventions/attributes-registry.md#assettenant-and-tenancy-scoping-rfc-v04-gap-closure) · [`security.tenant.*`](spec/semantic-conventions/attributes-registry.md#securitytenant-rfc-v04-gap-closure)
- **Conventions:** [`asset-inventory-spans.md` § Tenancy Scoping](spec/semantic-conventions/asset-inventory-spans.md#cross-cutting-tenancy-scoping-rfc-v04-gap-closure) · [`security-cross-cutting.md` § Tenant-Crossing Detection](spec/semantic-conventions/security-cross-cutting.md#tenant-crossing-detection)
- **Python SDK:** `AssetInventoryAttributes.TENANT_*`, `SecurityAttributes.TENANT_*`

Records the tenant and organization owning the agent, the session and the invoking user, on each. The identity half lands in `asset.*` and on the session and user; the detection half lands in `security.tenant.*`, because a crossing is a security verdict rather than an inventory fact. The emission rule requires these as span attributes rather than Resource attributes in any multi-tenant process, since a Resource attribute cannot vary per request.

### 21. Capability-Set Change Event

- **Tier:** MUST · **RFC §14** · **Grounding attacks:** `AOC-09`, `AOC-10`, `TA-06`, `AOC-02`
- **ODIS field:** `approved_software_refs` (6.1)
- **Namespace:** `asset.capability.*` — **16 attributes**
- **Registry:** [`asset.capability.*`](spec/semantic-conventions/attributes-registry.md#assetcapability-rfc-v04-gap-closure)
- **Conventions:** [`asset-inventory-spans.md` § Event `asset.capability.changed`](spec/semantic-conventions/asset-inventory-spans.md#event-assetcapabilitychanged-rfc-v04-gap-closure) · [`events.md` § `asset.capability.changed`](spec/semantic-conventions/events.md#assetcapabilitychanged-rfc-v04-gap-closure)
- **Python SDK:** `AssetInventoryAttributes.CAPABILITY_*`, `.CapabilityChangeType`

Adds an event emitted whenever the agent's usable capability set changes at runtime — a tool, MCP server, model, knowledge source or memory store discovered, added, removed or modified — with before/after set hashes and what triggered the change. `change_source` separates governed paths (`deployment`, `configuration`, `operator`) from ungoverned ones (`runtime_discovery`, `server_push`, `unknown`). The emission rule requires an event; a consumer must not be expected to infer capability change by diffing inventory snapshots.

### 22. Instrumentation Coverage / Hook Attestation

- **Tier:** SHOULD · **RFC §15** · **Grounding attacks:** `AOC-10`, `AOC-01`, `AOC-09`
- **Namespace:** `asset.instrumentation.*` — **17 attributes**
- **Registry:** [`asset.instrumentation.*`](spec/semantic-conventions/attributes-registry.md#assetinstrumentation-rfc-v04-gap-closure)
- **Conventions:** [`asset-inventory-spans.md` § Instrumentation Coverage & Hook Attestation](spec/semantic-conventions/asset-inventory-spans.md#instrumentation-coverage--hook-attestation-rfc-v04-gap-closure) · [`events.md` § `asset.instrumentation.gap_detected`](spec/semantic-conventions/events.md#assetinstrumentationgap_detected-rfc-v04-gap-closure)
- **Python SDK:** `AssetInventoryAttributes.INSTRUMENTATION_*`

Records which lifecycle hooks are instrumented and active, the instrumentation version and SDK, declared-vs-active hook sets, and a coverage ratio — the difference between "no events" and "not observed." `coverage_ratio` bounds the confidence of any negative finding. `attestation.method: "self_declared"` is unverifiable under §16, and `exporter.configured: true` with `exporter.reachable: false` describes an agent that is functionally uninstrumented while appearing configured. A new `asset.instrumentation.gap_detected` event gives the group an emission trigger.

### 23. Execution Environment / Sandbox

- **Tier:** MUST · **RFC §9** · **Grounding attacks:** `TA-06`, `AOC-02`, `AOC-04`, `AOC-14`
- **ODIS field:** `binding_profile` (6.2, partial)
- **Namespace:** `supply_chain.runtime.*` — **27 attributes**
- **Registry:** [`supply_chain.runtime.*`](spec/semantic-conventions/attributes-registry.md#supply_chainruntime-rfc-v04-gap-closure)
- **Conventions:** [`security-cross-cutting.md` § Execution Environment / Sandbox](spec/semantic-conventions/security-cross-cutting.md#execution-environment--sandbox)
- **Python SDK:** `SupplyChainAttributes.RUNTIME_*`

Records the isolation posture in force at execution time: sandbox mode, language runtime, OS and architecture, image digest, privilege, capabilities, network egress policy, filesystem mode, resource limits and exposed secrets. The same `exec` in a network-isolated read-only container and on a privileged host are different events, and telemetry that omits containment cannot distinguish them. Attached as Resource attributes where the environment is stable, and **required on the span** where execution is delegated to a per-call sandbox.

### 24. A2A Task Lifecycle Event

- **Tier:** SHOULD · **RFC §12** · **Grounding attacks:** `AOC-04`, `AOC-11`, `AOC-09`, `TA-01`
- **ODIS field:** `delegation_id`, `parent_delegation_id` (6.3)
- **Namespace:** `a2a.task.lifecycle.*` · `a2a.push.config.*` — **30 attributes**
- **Registry:** [`a2a.task.lifecycle.*`](spec/semantic-conventions/attributes-registry.md#a2atasklifecycle-rfc-v04-gap-closure) · [`a2a.push.config.*`](spec/semantic-conventions/attributes-registry.md#a2apushconfig-rfc-v04-gap-closure)
- **Conventions:** [`a2a-spans.md` § Event `a2a.task.lifecycle.*`](spec/semantic-conventions/a2a-spans.md#event-a2atasklifecycle-rfc-v04-gap-closure) · [`events.md` § A2A Task Lifecycle Events](spec/semantic-conventions/events.md#a2a-task-lifecycle-events-rfc-v04-gap-closure)
- **Python SDK:** `A2AAttributes.TASK_LIFECYCLE_*`, `.PUSH_CONFIG_*`

Adds delegated-task state transitions across the A2A surface as events — submitted, working, input-required, cancelled, completed, expired, orphaned — with `transition_valid` flagging illegal transitions, plus push-notification configuration changes that register an outbound callback destination. `initiating_trace_id` is **Required** because the terminal event usually belongs to a different trace than the one that started the task. `orphaned: true` identifies a task holding outstanding delegated authority that no one is waiting on.

### 25. Attribute Source / Trusted-Provenance Marking

- **Tier:** MUST ‡ · **RFC §16** · **Grounding attacks:** `AOC-01`, `AOC-08`, `AOC-10`, `AOC-15`
- **ODIS field:** `attestation_method` (6.2, partial)
- **Namespace:** `security.attribute_source.*` — **6 attributes**
- **Registry:** [`security.attribute_source.*`](spec/semantic-conventions/attributes-registry.md#securityattribute_source-rfc-v04-gap-closure)
- **Conventions:** [`security-cross-cutting.md` § Attribute Source / Trusted-Provenance Marking](spec/semantic-conventions/security-cross-cutting.md#attribute-source--trusted-provenance-marking)
- **Python SDK:** `SecurityAttributes.ATTRIBUTE_SOURCE_*`, `.AttributeSource`

Records, for every security-relevant attribute, the authority that supplied it: verified IdP token, policy decision point, enforcement-owned store, platform runtime, or self-asserted by the agent. The most useful form is the inverse list — `self_asserted` answers *"which claims on this span came from the model rather than an enforcement point?"* in one query. This is the convention every other security group in the framework depends on, and it carries the Appendix D.5 baggage prohibition.

### 26. Mediation Coverage & Bypass Path

- **Tier:** SHOULD · **RFC §16** · **Grounding attacks:** `AOC-02`, `AOC-14`, `TA-06`
- **Namespace:** `security.mediation.*` — **8 attributes**
- **Registry:** [`security.mediation.*`](spec/semantic-conventions/attributes-registry.md#securitymediation-rfc-v04-gap-closure)
- **Conventions:** [`security-cross-cutting.md` § Mediation Coverage & Bypass Path](spec/semantic-conventions/security-cross-cutting.md#mediation-coverage--bypass-path)
- **Python SDK:** `SecurityAttributes.MEDIATION_*`

Records whether the operation traversed a reference monitor **at all**, at which placement, and whether unmediated paths to the same capability exist. `mediated: false` is what makes a missing authorization decision interpretable rather than empty. `bypass_available: true` bounds every claim made about the control: a capability's effective strength equals its weakest route, so a dashboard reporting full policy coverage over mediated calls is reporting the numerator only.

### 27. Event Sequence Continuity

- **Tier:** SHOULD · **RFC §15** · **Grounding attacks:** `AOC-01`, `AOC-10`, `AOC-07`
- **Namespace:** `observability.sequence.*` — **12 attributes**
- **Registry:** [`observability.sequence.*`](spec/semantic-conventions/attributes-registry.md#observabilitysequence-rfc-v04-gap-closure)
- **Conventions:** [`security-cross-cutting.md` § Event Sequence Continuity](spec/semantic-conventions/security-cross-cutting.md#event-sequence-continuity)
- **Python SDK:** `ObservabilityAttributes.SEQUENCE_*`, `.SAMPLING_*`

Adds a per-session monotonic sequence number, optionally chained via `prev_hash` or signed, assigned at the producer rather than the collector. This is what makes a **gap** in the event stream distinguishable from a quiet period — an agent that did nothing versus an agent whose activity was not recorded. `sampling.security_relevant` marks events that samplers must not probabilistically drop, per Appendix D.5.

---

## What this log does not claim

The 390 attributes are **defined**, not **emitted**. Closing an Appendix C gap means AITF now has a
place to put the field and a rule for how to put it there; whether any given deployment populates it
is a separate question, and is precisely what
[`asset.instrumentation.*`](spec/semantic-conventions/security-cross-cutting.md) and
[`observability.sequence.*`](spec/semantic-conventions/security-cross-cutting.md#event-sequence-continuity)
exist to answer.

Upstream coverage is tracked separately in [`upstream-status.md`](upstream-status.md): several of
these fields have a home in the OpenTelemetry GenAI conventions or in OCSF today, several are
proposed, and several are net-new asks. Appendix D.1 and D.2 of the RFC should be read before filing
any of them upstream as new.
