# Cross-Cutting Security & Observability-Plane Conventions

**Status:** Stable · **Version:** 1.0 · **Closes RFC v0.4 Appendix C gaps**

These conventions are **cross-cutting**: unlike `gen-ai-spans.md` or `mcp-spans.md`, they do not
describe a span of their own. They are attribute groups attached to whichever span or event
represents the mediated operation — an inference call, a tool invocation, a delegation hop, a
retrieval. This document exists because eight fields in RFC Appendix C either landed in
`security.*` without a home document or were recorded as *"none identified — a new attribute or
namespace is required"*.

Canonical definitions for every attribute below live in
[`attributes-registry.md`](attributes-registry.md). This document adds the emission rules and the
interpretation guidance that the registry's row-per-attribute format cannot carry.

## Contents

| Convention | Tier | RFC § | Registry |
|---|---|---|---|
| [Attribute Source / Trusted-Provenance Marking](#attribute-source--trusted-provenance-marking) | MUST ‡ | 16 | [`security.attribute_source.*`](attributes-registry.md#securityattribute_source-rfc-v04-gap-closure) |
| [Authorization Decision Record](#authorization-decision-record) | MUST | 16 | [`security.authorization.*`](attributes-registry.md#securityauthorization-rfc-v04-gap-closure) |
| [Execution Environment / Sandbox](#execution-environment--sandbox) | MUST | 9 | [`supply_chain.runtime.*`](attributes-registry.md#supply_chainruntime-rfc-v04-gap-closure) |
| [Mediation Coverage & Bypass Path](#mediation-coverage--bypass-path) | SHOULD | 16 | [`security.mediation.*`](attributes-registry.md#securitymediation-rfc-v04-gap-closure) |
| [Enforcement-Point Availability & Failure Mode](#enforcement-point-availability--failure-mode) | SHOULD | 15 | [`security.enforcement.*`](attributes-registry.md#securityenforcement-rfc-v04-gap-closure) |
| [Guardrail Modification Record](#guardrail-modification-record) | SHOULD ‡ | 6 | [`security.guardrail.modification.*`](attributes-registry.md#securityguardrailmodification-rfc-v04-gap-closure) |
| [Session Taint Labels & Information-Flow Decisions](#session-taint-labels--information-flow-decisions) | SHOULD | 16 | [`security.taint.*`](attributes-registry.md#securitytaint-rfc-v04-gap-closure) |
| [Event Sequence Continuity](#event-sequence-continuity) | SHOULD | 15 | [`observability.sequence.*`](attributes-registry.md#observabilitysequence-rfc-v04-gap-closure) |
| [Tenant-Crossing Detection](#tenant-crossing-detection) | SHOULD | 5 | [`security.tenant.*`](attributes-registry.md#securitytenant-rfc-v04-gap-closure) |
| [Encoded / Obfuscated Payload Indicator](#encoded--obfuscated-payload-indicator) | MAY | 6 | [`security.obfuscation.*`](attributes-registry.md#securityobfuscation-rfc-v04-gap-closure) |

---

## Attribute Source / Trusted-Provenance Marking

> **RFC field:** Attribute Source / Trusted-Provenance Marking · **Tier:** MUST ‡ · **§16 Policy Enforcement & Mediation** · **Grounding attacks:** `AOC-01`, `AOC-08`, `AOC-10`, `AOC-15` · **ODIS:** `attestation_method` (6.2, partial)

Appendix C recorded no namespace for this field. AITF lands it in `security.*` because it is a
property of a security *assertion* rather than of an identity, and because it is the direct AITF
expression of the `attribute_source` enum the RFC asks OCSF to add (Appendix E.2 item 4).

This is the convention on which the others depend. A `security.authorization.decision` of `"allow"`
means one thing when a policy decision point produced it and something else entirely when the model
reported it. Every other group in this document is only as trustworthy as its provenance marking.

### The Rule

For every security-relevant attribute, record the authority that supplied it. Per RFC §16, an
**unmarked value is an unverified value** — consumers MUST read an unmarked security attribute as
`"self_asserted"`. Producers therefore do not need to mark self-asserted values exhaustively, but
they MUST NOT leave a *verified* value unmarked, because that under-states its trustworthiness in the
opposite direction and defeats the point of the marking.

The most operationally useful form is the inverse list. `security.attribute_source.self_asserted`
answers *"which of the claims on this span came from the model rather than from an enforcement
point?"* in a single query, without a consumer having to enumerate every attribute and check it.

### Emission Rule (RFC Appendix D.5)

W3C `baggage` crosses the trust boundary and is attacker-influenceable.

- Baggage MAY carry correlation identifiers.
- Baggage **MUST NOT** carry `security.taint.*`, `security.authorization.*`,
  `security.attribute_source.*`, trust levels, or identity claims.
- Any such value received over the wire **MUST** be recorded as `"self_asserted"` regardless of what
  the sender labelled it.

| Attribute | Type | Requirement | Description | Compliance |
|-----------|------|-------------|-------------|------------|
| `security.attribute_source.default` | string | **Required** | Source that applies to every security-relevant attribute on this span unless overridden in the map: `"idp"`, `"pdp"`, `"enforcement_state"`, `"platform"`, `"runtime"`, `"self_asserted"` | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 |
| `security.attribute_source.map` | string | **Recommended** | JSON object mapping attribute name → source enum, for attributes whose source differs from the default | NIST AI RMF GOVERN-1.2 |
| `security.attribute_source.self_asserted` | string[] | **Required** | Names of attributes on this span asserted by the agent or model rather than by an authority. The primary detection query | MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051), NIST AI RMF GOVERN-1.2 |
| `security.attribute_source.verified` | string[] | **Recommended** | Names of attributes supplied by a verified authority | NIST AI RMF GOVERN-1.2 |
| `security.attribute_source.authority.id` | string | **Recommended** | Identifier of the authority behind `default` (IdP issuer, PDP instance, enforcement-point ID) | NIST AI RMF MAP-1.5 |
| `security.attribute_source.verification_time` | string | **Optional** | ISO 8601 timestamp at which the verified attributes were last confirmed | EU AI Act Art.12 |

---

## Authorization Decision Record

> **RFC field:** Authorization Decision Record · **Tier:** MUST · **§16 Policy Enforcement & Mediation** · **Grounding attacks:** `AOC-02`, `AOC-08`, `AOC-10`, `TA-08`, `TA-12`, `TA-13` · **ODIS:** `policy_profile_ref` (6.1)

Per mediated operation: the decision, the deny reason and its machine-readable code, the deciding
authority, the rule that produced it, and any obligations attached.

**Distinct from a content-guardrail verdict.** `security.guardrail.*` carries a classifier score;
this carries the authorization outcome. An operation can pass every guardrail and still be denied,
and the two records answer different questions during an investigation.

Six independent attacks in RFC Appendix A require this field, so under the Appendix E.3 promotion
rule it is an OCSF standardization ask (`ai_authorization`), not merely an AITF-local attribute.

**Relationship to `identity.authz.*`.** `identity.authz.*` remains the identity-plane view, mapped to
OCSF Authorization (3003), for principal-to-resource decisions. `security.authorization.*` is the
AI-operation view, emitted on the mediated agent operation itself. Where both apply,
`security.authorization.decision_id` joins them.

### Emission Rule

- Every mediated operation **MUST** carry `security.authorization.decision`. A missing decision is
  a coverage gap, not an implicit allow — see [Mediation Coverage](#mediation-coverage--bypass-path).
- A `deny` **MUST** carry `security.authorization.reason_code`. Free-text `reason` is for humans;
  the code is what detections match on.
- Where obligations are attached, `obligations_fulfilled` **MUST** be recorded. An allow whose
  obligations were not fulfilled is a failed enforcement, not a successful one.

| Attribute | Type | Requirement | Description | Compliance |
|-----------|------|-------------|-------------|------------|
| `security.authorization.decision` | string | **Required** | `"allow"`, `"deny"`, `"modify"` | OWASP LLM06, NIST AI RMF GOVERN-1.2, EU AI Act Art.14 |
| `security.authorization.decision_id` | string | **Required** | Correlation identifier for this decision. Referenced from the OTel span so trace and OCSF event join at query time (RFC Appendix D.3) | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 |
| `security.authorization.operation` | string | **Required** | The mediated operation: `"tool_call"`, `"model_invoke"`, `"memory_write"`, `"retrieval"`, `"delegation"`, `"egress"`, `"capability_change"` | OWASP LLM06 |
| `security.authorization.reason` | string | **Recommended** | Human-readable reason, required on `"deny"` and `"modify"` | NIST AI RMF MEASURE-2.5 |
| `security.authorization.reason_code` | string | **Conditionally Required** (on `"deny"`) | Machine-readable deny code (`"SCOPE_INSUFFICIENT"`, `"TAINTED_SESSION"`, `"DEPTH_EXCEEDED"`, `"UNAPPROVED_TOOL"`, `"RESIDENCY_VIOLATION"`) | NIST AI RMF GOVERN-1.2 |
| `security.authorization.authority.type` | string | **Required** | `"inline_rule"` or `"external_pdp"` | NIST AI RMF GOVERN-1.2 |
| `security.authorization.authority.id` | string | **Recommended** | Identifier of the deciding authority instance | NIST AI RMF MAP-1.5 |
| `security.authorization.authority.engine` | string | **Recommended** | Policy engine: `"cedar"`, `"cel"`, `"opa"`, `"xacml"`, `"casbin"`, `"custom"` | NIST AI RMF MAP-1.5 |
| `security.authorization.rule_id` | string | **Required** | Rule or policy identifier that produced the decision | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 |
| `security.authorization.policy_version` | string | **Recommended** | Version or digest of the policy bundle in force | NIST AI RMF GOVERN-1.2 |
| `security.authorization.obligations` | string[] | **Recommended** | Obligations attached to an `"allow"` or `"modify"` (`"redact_pii"`, `"require_approval"`, `"log_only"`, `"rate_limit"`) | EU AI Act Art.14, NIST AI RMF GOVERN-1.2 |
| `security.authorization.obligations_fulfilled` | boolean | **Recommended** | Whether every attached obligation was carried out. `false` is a control failure | EU AI Act Art.14 |
| `security.authorization.principal` | string | **Recommended** | Principal the decision was made for. Joins to `identity.*` | NIST AI RMF GOVERN-1.2 |
| `security.authorization.resource` | string | **Recommended** | Resource the decision was made about | OWASP LLM06 |
| `security.authorization.latency_ms` | double | **Optional** | Time to obtain the decision | NIST AI RMF MEASURE-2.5 |
| `compliance.framework` | string | **Recommended** | Single framework reference for this decision, e.g. `"mitre_atlas"`, `"nist_ai_rmf"`, `"iso_42001"`, `"eu_ai_act"`, `"csa_aicm"` | NIST AI RMF GOVERN-1.2 |
| `compliance.control_id` | string | **Recommended** | Single control identifier within `compliance.framework`. Carries the ATLAS `AML.Txxxx` tag in Phase 0 until OCSF accepts ATLAS natively (RFC Appendix E.3) | MITRE ATLAS, NIST AI RMF GOVERN-1.2 |

### Reading the Fields

`decision: "allow"` with `authority.type` absent or `"none"` is an operation that no policy engine
evaluated — it was permitted by default, which is a materially weaker statement than being permitted
by rule.

`obligations` non-empty with `obligations_fulfilled: false` is the fail-open case that most resembles
a clean allow in aggregate dashboards. It should be alerted on, not counted as an allow.

`decision: "error"` is the honest value when the authority could not be consulted. Recording it as
`"allow"` because the operation proceeded destroys the distinction between a permitted action and an
unevaluated one.

---

## Execution Environment / Sandbox

> **RFC field:** Execution Environment / Sandbox · **Tier:** MUST · **§9 Tools & External Services** · **Grounding attacks:** `TA-06`, `AOC-02`, `AOC-04`, `AOC-14` · **ODIS:** `binding_profile` (6.2, partial)

An agent action's blast radius is a property of the environment it executes in, not of the action
itself. The same `exec` in a network-isolated read-only container and on a privileged host with
unrestricted egress are different events entirely, and telemetry that does not record the containment
posture cannot distinguish them. These attributes describe the environment **in force at execution
time**, and are the precondition for judging whether a sandbox escape occurred.

### Emission Rule

`supply_chain.runtime.*` describes the environment, not the action.

- Where the environment is stable for the process lifetime, these SHOULD be attached as OTel
  **Resource** attributes.
- Where execution is delegated to a per-call environment — a fresh code-interpreter container, a
  remote sandbox service — they **MUST** be attached to the span, because the environment then
  differs between spans within a single process.

| Attribute | Type | Requirement | Description | Compliance |
|-----------|------|-------------|-------------|------------|
| `supply_chain.runtime.sandbox.mode` | string | **Required** | Containment posture: `"none"`, `"process"`, `"container"`, `"microvm"`, `"gvisor"`, `"wasm"`, `"vm"`, `"remote_service"` | MITRE ATLAS [AML.T0011](https://atlas.mitre.org/techniques/AML.T0011), OWASP LLM06 (Excessive Agency) |
| `supply_chain.runtime.sandbox.provider` | string | **Recommended** | Sandbox implementation: `"docker"`, `"firecracker"`, `"kata"`, `"e2b"`, `"seccomp"`, `"nsjail"` | NIST AI RMF MAP-1.5 |
| `supply_chain.runtime.language.name` | string | **Recommended** | Language runtime executing the action: `"python"`, `"node"`, `"go"`, `"java"`, `"shell"`, `"wasm"` | NIST AI RMF MAP-1.5, MITRE ATLAS [AML.T0011](https://atlas.mitre.org/techniques/AML.T0011) |
| `supply_chain.runtime.language.version` | string | **Recommended** | Version of that runtime. Pins which known interpreter escapes apply to this execution | MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010), NIST AI RMF GOVERN-1.2 |
| `supply_chain.runtime.os.type` | string | **Recommended** | Operating system of the execution environment: `"linux"`, `"darwin"`, `"windows"` | NIST AI RMF MAP-1.5 |
| `supply_chain.runtime.os.version` | string | **Optional** | OS or kernel version | MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010) |
| `supply_chain.runtime.architecture` | string | **Recommended** | CPU architecture: `"amd64"`, `"arm64"`, `"wasm32"` | NIST AI RMF MAP-1.5 |
| `supply_chain.runtime.instance.id` | string | **Required** | Identifier of the running instance (container ID, VM ID, pod UID) — the join key for correlating an action to its environment | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 (Record-Keeping) |
| `supply_chain.runtime.image.digest` | string | **Recommended** | Digest of the execution image (`sha256:…`) | MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010), NIST AI RMF GOVERN-1.2 |
| `supply_chain.runtime.image.ref` | string | **Optional** | Human-readable image reference (registry/repo:tag) | NIST AI RMF MAP-1.5 |
| `supply_chain.runtime.privileged` | boolean | **Required** | Whether execution runs with elevated privilege. Emit `false` explicitly | OWASP LLM06 (Excessive Agency), MITRE ATLAS [AML.T0011](https://atlas.mitre.org/techniques/AML.T0011) |
| `supply_chain.runtime.user` | string | **Recommended** | OS user/UID the workload runs as | OWASP LLM06 |
| `supply_chain.runtime.capabilities` | string[] | **Optional** | Linux capabilities or equivalent privileges retained | OWASP LLM06, MITRE ATLAS [AML.T0011](https://atlas.mitre.org/techniques/AML.T0011) |
| `supply_chain.runtime.network.egress_policy` | string | **Required** | Egress posture: `"none"`, `"allowlist"`, `"proxied"`, `"unrestricted"` | OWASP LLM02 (Sensitive Information Disclosure), MITRE ATLAS [AML.T0025](https://atlas.mitre.org/techniques/AML.T0025) (Exfiltration) |
| `supply_chain.runtime.network.egress_allowlist` | string[] | **Recommended** | Permitted egress destinations when `egress_policy` is `"allowlist"` | MITRE ATLAS [AML.T0025](https://atlas.mitre.org/techniques/AML.T0025), OWASP LLM02 |
| `supply_chain.runtime.network.namespace` | string | **Optional** | Network namespace or VPC/subnet identifier | NIST AI RMF MAP-1.5 |
| `supply_chain.runtime.filesystem.mode` | string | **Required** | Filesystem posture: `"read_only"`, `"ephemeral"`, `"scoped_write"`, `"host_mount"`, `"unrestricted"` | OWASP LLM06, MITRE ATLAS [AML.T0011](https://atlas.mitre.org/techniques/AML.T0011) |
| `supply_chain.runtime.filesystem.mounts` | string[] | **Optional** | Host paths mounted into the environment — the principal escape surface | MITRE ATLAS [AML.T0011](https://atlas.mitre.org/techniques/AML.T0011), OWASP LLM06 |
| `supply_chain.runtime.resource.cpu_limit` | double | **Optional** | CPU limit in cores | NIST AI RMF MEASURE-2.5 |
| `supply_chain.runtime.resource.memory_limit_bytes` | int | **Optional** | Memory limit in bytes | NIST AI RMF MEASURE-2.5 |
| `supply_chain.runtime.resource.timeout_ms` | int | **Optional** | Wall-clock execution limit | NIST AI RMF MEASURE-2.5 |
| `supply_chain.runtime.secrets.exposed` | string[] | **Recommended** | Names (never values) of secrets or credentials reachable from the environment | OWASP LLM02, MITRE ATLAS [AML.T0055](https://atlas.mitre.org/techniques/AML.T0055) (Unsecured Credentials) |
| `supply_chain.runtime.secrets.count` | int | **Optional** | Count of reachable secrets | OWASP LLM02 |
| `supply_chain.runtime.attestation.method` | string | **Recommended** | How the environment was attested: `"none"`, `"self_declared"`, `"orchestrator"`, `"tpm"`, `"sev_snp"`, `"tdx"`, `"nitro"` | NIST AI RMF GOVERN-1.2, MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010) |
| `supply_chain.runtime.attestation.verified` | boolean | **Recommended** | Whether attestation was verified. Absent or `false` means self-asserted per the provenance rule | NIST AI RMF GOVERN-1.2 |
| `supply_chain.runtime.escape_detected` | boolean | **Optional** | Whether an attempt to act outside the declared boundary was observed | MITRE ATLAS [AML.T0011](https://atlas.mitre.org/techniques/AML.T0011), OWASP LLM06 |
| `supply_chain.runtime.escape_indicator` | string | **Optional** | What was observed: `"host_path_access"`, `"undeclared_egress"`, `"privilege_gain"`, `"namespace_break"` | MITRE ATLAS [AML.T0011](https://atlas.mitre.org/techniques/AML.T0011) |

### Reading the Fields

`sandbox.mode: "none"` on a span that also carries tool arguments derived from untrusted input is the
`TA-06` precondition stated in one place.

`privileged: true`, a permissive `network.egress_policy`, or a non-empty `secrets.exposed` each widen
the blast radius independently of what the action did. They are the difference between an incident
scoped to one container and one scoped to the host.

`attestation.verified: false` means the posture above is *declared*, not measured. Per §16, treat the
whole group as self-asserted in that case.

---

## Mediation Coverage & Bypass Path

> **RFC field:** Mediation Coverage & Bypass Path · **Tier:** SHOULD · **§16 Policy Enforcement & Mediation** · **Grounding attacks:** `AOC-02`, `AOC-14`, `TA-06`

Appendix C recorded no namespace for this field. AITF lands it in `security.*` alongside the other
reference-monitor attributes.

It answers whether the operation traversed a reference monitor **at all**, at which placement, and
whether unmediated paths to the same capability exist. A capability with a bypass path has an
effective control strength equal to its weakest route, not its instrumented one — so a dashboard
showing 100% policy coverage over mediated calls is reporting on the numerator only.

| Attribute | Type | Requirement | Description | Compliance |
|-----------|------|-------------|-------------|------------|
| `security.mediation.mediated` | boolean | **Required** | Whether this operation traversed a reference monitor | NIST AI RMF GOVERN-1.2, EU AI Act Art.14 |
| `security.mediation.placement` | string | **Conditionally Required** (when `mediated` = `true`) | Where mediation occurred: `"inbound_gateway"`, `"egress_sidecar"`, `"in_process_framework"`, `"proxy"`, `"host_agent"`, `"none"` | NIST AI RMF MAP-1.5 |
| `security.mediation.reference_monitor.id` | string | **Recommended** | Identifier of the monitor that mediated the operation | NIST AI RMF GOVERN-1.2 |
| `security.mediation.bypass_available` | boolean | **Required** | Whether an unmediated path to the same capability is known to exist | OWASP LLM06, NIST AI RMF GOVERN-1.2 |
| `security.mediation.bypass_paths` | string[] | **Recommended** (when `bypass_available` = `true`) | Known unmediated routes (`"direct_sdk_call"`, `"sidecar_disabled"`, `"local_exec"`, `"alternate_endpoint"`) | OWASP LLM06 |
| `security.mediation.coverage_ratio` | double | **Optional** | Fraction of known invocation paths to this capability that are mediated (0.0–1.0) | NIST AI RMF MEASURE-2.5 |
| `security.mediation.capability` | string | **Recommended** | The capability being mediated, so coverage can be aggregated per capability rather than per span | NIST AI RMF MAP-1.5 |
| `security.mediation.assessed_at` | string | **Optional** | ISO 8601 timestamp of the last bypass-path assessment. A stale assessment is weak evidence | EU AI Act Art.12 |

### Reading the Fields

`mediated: false` is the value that makes the absence of a `security.authorization.decision`
interpretable. Without it, an unmediated operation and a mediated-and-allowed operation are the same
empty record.

`bypass_available: true` bounds every claim made about the control. `bypass_paths` names what an
assessment must cover before the control can be called effective; `coverage_ratio` quantifies how
much of the capability's traffic the reference monitor actually sees.

`placement: "in_process"` is the weakest placement against a compromised agent, since the monitor and
the thing it monitors share a trust boundary. It is not a finding on its own, but it is the context
in which `AOC-14` becomes plausible.

---

## Enforcement-Point Availability & Failure Mode

> **RFC field:** Enforcement-Point Availability & Failure Mode · **Tier:** SHOULD · **§15 Observability-Plane Integrity** · **Grounding attacks:** `TA-01`, `TA-10`, `IR-01`, `AOC-12`

For each enforcement callout — guardrail, policy engine, external guardian — whether it was reached,
its latency, and on failure whether the system **failed open or failed closed**, plus the action taken
anyway.

A guardrail that times out and is skipped produces no verdict, and *"no verdict"* is indistinguishable
from *"clean"* unless this is recorded. This is the same failure mode the RFC identifies for sampling
in Appendix D.5.

### Emission Rule

- When an enforcement callout is attempted, `security.enforcement.reached` **MUST** be recorded.
- When `reached` is `false`, `failure_mode` and `action_taken` **MUST** be recorded. The absence of a
  guardrail verdict **MUST NOT** be rendered as a passing verdict.

| Attribute | Type | Requirement | Description | Compliance |
|-----------|------|-------------|-------------|------------|
| `security.enforcement.point.name` | string | **Required** | Name of the enforcement point called | NIST AI RMF GOVERN-1.2 |
| `security.enforcement.point.type` | string | **Required** | `"guardrail"`, `"policy_engine"`, `"external_guardian"`, `"pdp"`, `"dlp"`, `"classifier"`, `"human_approval"` | NIST AI RMF MAP-1.5 |
| `security.enforcement.point.version` | string | **Optional** | Version of the enforcement point or its ruleset | NIST AI RMF GOVERN-1.2 |
| `security.enforcement.reached` | boolean | **Required** | Whether the callout completed and returned a verdict | EU AI Act Art.14, NIST AI RMF GOVERN-1.2 |
| `security.enforcement.latency_ms` | double | **Recommended** | Round-trip latency of the callout | NIST AI RMF MEASURE-2.5 |
| `security.enforcement.timeout_ms` | double | **Optional** | Configured timeout, for interpreting the latency | NIST AI RMF MEASURE-2.5 |
| `security.enforcement.failure_mode` | string | **Conditionally Required** (when `reached` = `false`) | `"fail_open"`, `"fail_closed"`, `"retry"`, `"degraded"` | EU AI Act Art.14, NIST AI RMF GOVERN-1.2 |
| `security.enforcement.failure_reason` | string | **Recommended** | `"timeout"`, `"connection_error"`, `"rate_limited"`, `"auth_error"`, `"disabled"`, `"not_configured"` | NIST AI RMF MEASURE-2.5 |
| `security.enforcement.action_taken` | string | **Required** | What happened to the operation regardless: `"allowed"`, `"blocked"`, `"degraded"`, `"queued"`, `"error"` | EU AI Act Art.14 |
| `security.enforcement.degraded_mode` | boolean | **Optional** | Whether the point returned a reduced-confidence verdict rather than failing outright | NIST AI RMF MEASURE-2.5 |

### Reading the Fields

`failure_mode: "fail_open"` is the single most consequential value in this document. It converts a
silent gap into a known one: the operation proceeded *without* the control, and any downstream
statistic that counts it as clean is wrong.

A rising rate of `timeout` or `unreachable` against one `point.name` is the `TA-10` pattern — an
attacker does not need to defeat a guardrail that can be made unavailable.

`degraded_mode: true` means the control ran but not at full fidelity; verdicts from it carry less
weight than verdicts from the same point in normal operation.

---

## Guardrail Modification Record

> **RFC field:** Guardrail Modification Record · **Tier:** SHOULD ‡ · **§6 Input Handling & Trust Provenance** · **Grounding attacks:** `TA-01`, `AOC-12`, `AOC-03`, `IR-01`

**Cross-cutting.** `security.guardrail.result` records pass, fail or warn, which is sufficient only
when an enforcement point's two options are allow and block. In practice enforcement points
**rewrite** — masking, redacting, stripping or normalizing a payload and letting it through.

Without this record the payload the model saw is not the payload that was logged, and a downstream
investigation reconstructs the wrong input. Applies on both the input and the output side.

### Emission Rule

- When an enforcement point alters a payload, `security.guardrail.modified` **MUST** be set `true`
  and `modification.side` recorded. A silent rewrite is a logging integrity defect.
- `hash_before` and `hash_after` SHOULD both be recorded even where the content itself is not
  retained; they are what allow the substitution to be detected later without storing the payload.

| Attribute | Type | Requirement | Description | Compliance |
|-----------|------|-------------|-------------|------------|
| `security.guardrail.action` | string | **Required** | What the enforcement point did: `"allow"`, `"block"`, `"modify"`, `"annotate"`. Supersedes the pass/fail reading of `security.guardrail.result` where the point can rewrite | EU AI Act Art.14, NIST AI RMF GOVERN-1.2 |
| `security.guardrail.modified` | boolean | **Required** | Whether the payload was altered in flight | MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051), OWASP LLM01 |
| `security.guardrail.modification.type` | string | **Recommended** | How it was altered: `"mask"`, `"redact"`, `"strip"`, `"rewrite"`, `"truncate"`, `"normalize"`, `"reorder"` | OWASP LLM01, EU AI Act Art.10 |
| `security.guardrail.modification.side` | string | **Recommended** | Which side was altered: `"input"`, `"output"` | OWASP LLM01, OWASP LLM02 |
| `security.guardrail.modification.enforcement_point` | string | **Required** | Which enforcement point made the change. Joins to `security.enforcement.point.name` | NIST AI RMF GOVERN-1.2 |
| `security.guardrail.modification.hash_before` | string | **Required** | SHA-256 of the payload as received | MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `security.guardrail.modification.hash_after` | string | **Required** | SHA-256 of the payload as forwarded | MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `security.guardrail.modification.count` | int | **Recommended** | Number of discrete edits applied | NIST AI RMF MEASURE-2.5 |
| `security.guardrail.modification.bytes_removed` | int | **Optional** | Net byte delta between before and after | NIST AI RMF MEASURE-2.5 |
| `security.guardrail.modification.redaction_map` | string | **Optional** | JSON map of offset/length/replacement-token for each edit. **Policy-gated**: emit only where retention policy permits, since it partially reconstructs the original | EU AI Act Art.10, GDPR Art.5 |
| `security.guardrail.modification.reason` | string | **Recommended** | Why the modification was applied | NIST AI RMF MEASURE-2.5 |

### Reading the Fields

`modified: true` with `action: "pass"` is the case this convention exists for — an operation recorded
as clean whose payload was materially changed on the way through.

`hash_before` ≠ `hash_after` on the **output** side (`modification.side: "output"`) means the logged
completion is not the completion the user received, which bounds what any post-hoc content analysis
can conclude.

`bytes_removed` large relative to the payload indicates stripping rather than masking, and the removed
span is exactly the region an investigation needs and no longer has.

---

## Session Taint Labels & Information-Flow Decisions

> **RFC field:** Session Taint Labels & Information-Flow Decisions · **Tier:** SHOULD · **§16 Policy Enforcement & Mediation** · **Grounding attacks:** `AOC-03`, `TA-01`, `TA-02`, `TA-05` · **ODIS:** `constraints.data_classification` (6.3)

Information-flow labels in force for the session or message: which labels are set, at what scope,
what operation applied each, and — critically — when an operation is **denied because of accumulated
taint rather than anything in its own payload**.

This is the write-down record, and it is the only field in the set that explains a denial whose own
payload looks clean. An analyst reviewing such a denial without taint labels sees a false positive.

### Emission Rule

- Where a flow decision is made, `security.taint.flow.decision` **MUST** be recorded.
- Where `denied` is `true`, `denied_labels` **MUST** be recorded — the labels that caused it, not
  merely that labels existed.
- Declassification **MUST** record `declassified_by` and `declassification_reason`. Unattributed
  declassification is indistinguishable from label loss.
- Per Appendix D.5, taint labels **MUST NOT** be propagated in W3C `baggage`.

| Attribute | Type | Requirement | Description | Compliance |
|-----------|------|-------------|-------------|------------|
| `security.taint.labels` | string[] | **Required** | Labels currently in force (`"untrusted_web"`, `"pii"`, `"confidential"`, `"external_email"`, `"customer_data"`) | EU AI Act Art.10, NIST AI RMF GOVERN-1.2 |
| `security.taint.scope` | string | **Required** | Scope at which the labels apply: `"session"`, `"turn"`, `"message"`, `"run"` | NIST AI RMF GOVERN-1.2 |
| `security.taint.label_count` | int | **Recommended** | Number of labels in force. Monotonic growth across a session is itself a signal | NIST AI RMF MEASURE-2.5 |
| `security.taint.applied_by` | string | **Recommended** | Operation that applied the most recent label (`"rag.retrieve"`, `"mcp.tool.invoke:fetch_url"`) | OWASP LLM01 |
| `security.taint.origin` | string | **Recommended** | Where the taint entered: `"user_input"`, `"retrieval"`, `"tool_result"`, `"memory"`, `"inbound_agent_message"`, `"attachment"` | OWASP LLM01, MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `security.taint.origin_span_id` | string | **Optional** | Span in which the taint was first applied, for back-tracing | NIST AI RMF GOVERN-1.2 |
| `security.taint.flow.decision` | string | **Recommended** | Information-flow outcome for this operation: `"allow"`, `"deny"`, `"declassify"`, `"downgrade"` | NIST AI RMF GOVERN-1.2, EU AI Act Art.14 |
| `security.taint.denied` | boolean | **Required** | Whether this operation was denied **because of accumulated taint** rather than its own payload | OWASP LLM02 (Sensitive Information Disclosure), NIST AI RMF GOVERN-1.2 |
| `security.taint.denied_labels` | string[] | **Recommended** (when `denied` = `true`) | Which labels caused the denial | NIST AI RMF GOVERN-1.2 |
| `security.taint.declassified_by` | string | **Optional** | Authority that removed or downgraded a label. Declassification is privileged and MUST NOT be self-asserted | NIST AI RMF GOVERN-1.2, EU AI Act Art.14 |
| `security.taint.declassification_reason` | string | **Optional** | Why a label was removed | NIST AI RMF MEASURE-2.5 |

### Reading the Fields

`denied: true` with `denied_labels` populated and a benign payload is the intended reading: the
operation was refused for what the session had already touched, not for what it asked.

`declassified_by` naming the agent itself, rather than a policy engine or a human, is self-service
declassification — the label was removed by the component the label existed to constrain.

`scope: "message"` where the deployment intends session-level tracking means labels are not
accumulating across turns, so `AOC-03` write-down across a conversation will not be caught.

---

## Event Sequence Continuity

> **RFC field:** Event Sequence Continuity · **Tier:** SHOULD · **§15 Observability-Plane Integrity** · **Grounding attacks:** `AOC-01`, `AOC-10`, `AOC-07`

Appendix C recorded no namespace for this field. AITF introduces `observability.*` for it, rather
than overloading `security.*`, because it describes the integrity of the telemetry channel itself
rather than the behaviour of the system under observation.

A per-session monotonic sequence number, optionally chained or signed, is what makes a **gap** in the
event stream distinguishable from a quiet period. It is the assume-breach counterpart to
`asset.instrumentation.*`: that group establishes whether a signal would ever have been produced,
this one establishes whether a produced signal survived the journey to the backend intact.

### Emission Rule

- `observability.sequence.number` SHOULD be monotonic within `sequence.scope_id` and assigned at the
  producer, not the collector — a collector-assigned number cannot detect loss in transit.
- Where chaining is used, `prev_hash` **MUST** reference the immediately preceding event in the same
  scope, so that a removed event breaks the chain rather than shortening it.
- `sampling.security_relevant` **MUST** be honoured by samplers: per Appendix D.5, security-relevant
  events are not eligible for probabilistic drop.

| Attribute | Type | Requirement | Description | Compliance |
|-----------|------|-------------|-------------|------------|
| `observability.sequence.number` | int | **Required** | Monotonically increasing sequence number, scoped to `observability.sequence.scope_id` | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 |
| `observability.sequence.scope_id` | string | **Required** | Scope the counter is monotonic within. Normally `gen_ai.conversation.id` | NIST AI RMF GOVERN-1.2 |
| `observability.sequence.scope_type` | string | **Recommended** | `"session"`, `"run"`, `"agent_instance"`, `"emitter"` | — |
| `observability.sequence.prev_hash` | string | **Optional** | SHA-256 of the previous event in the chain, forming a tamper-evident log | MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051), NIST AI RMF GOVERN-1.2 |
| `observability.sequence.hash` | string | **Optional** | SHA-256 of this event including `prev_hash` | NIST AI RMF GOVERN-1.2 |
| `observability.sequence.signature` | string | **Optional** | Detached signature over `hash`, where the emitter holds a signing key | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 |
| `observability.sequence.signer_key_id` | string | **Optional** | Key identifier used for the signature | NIST AI RMF GOVERN-1.2 |
| `observability.sequence.gap_detected` | boolean | **Recommended** | Whether the consumer observed a missing sequence number before this event | NIST AI RMF MEASURE-2.5 |
| `observability.sequence.gap_size` | int | **Optional** | Number of sequence numbers missing | NIST AI RMF MEASURE-2.5 |
| `observability.sequence.reordered` | boolean | **Optional** | Whether this event arrived out of sequence order | NIST AI RMF MEASURE-2.5 |
| `observability.sampling.decision` | string | **Recommended** | Sampling outcome for this event: `"recorded"`, `"head_sampled_out"`, `"tail_retained"`, `"forced_retain"`. Per RFC Appendix D.5, security-relevant events are recorded at 100% and the sampling configuration in force is itself telemetry | NIST AI RMF GOVERN-1.2 |
| `observability.sampling.security_relevant` | boolean | **Recommended** | Whether this event carries a security predicate that forces retention | NIST AI RMF GOVERN-1.2 |

### Reading the Fields

`gap_detected: true` with `gap_size` is the difference between an agent that did nothing and an agent
whose activity was not recorded. Only the second is an incident, and without this field both look
alike.

`reordered: true` bounds any conclusion drawn from event ordering — including causality claims in an
incident timeline.

`sampling.decision: "drop"` on an event carrying `security.authorization.*` or `security.taint.*` is a
sampler misconfiguration under D.5, and is itself worth alerting on.

---

## Tenant-Crossing Detection

> Supports **Organization / Tenant ID** · **Tier:** SHOULD · **§5 Application & Agent Reasoning Core** · **Grounding attacks:** `AOC-02`, `AOC-05`, `TA-05`, `TA-01`, `TA-11`

The identity half of the Organization / Tenant ID field lands in `asset.*` and on the session and
user — see [Cross-Cutting: Tenancy Scoping](asset-inventory-spans.md#cross-cutting-tenancy-scoping-rfc-v04-gap-closure).
The *detection* half lands here, because a cross-tenant crossing is a security verdict rather than an
inventory fact.

| Attribute | Type | Requirement | Description | Compliance |
|-----------|------|-------------|-------------|------------|
| `security.tenant.crossing_detected` | boolean | **Recommended** | Whether the agent, session and invoking-user tenants disagree on this operation | OWASP LLM02, NIST AI RMF GOVERN-1.2, GDPR Art.5 |
| `security.tenant.crossing_type` | string | **Optional** | Which pair disagreed: `"agent_session"`, `"session_user"`, `"agent_user"`, `"resource_tenant"` | NIST AI RMF GOVERN-1.2 |
| `security.tenant.expected` | string | **Optional** | Tenant the operation was expected to run under | NIST AI RMF GOVERN-1.2 |

### Reading the Fields

`crossing_detected: true` is emitted when the agent, session and invoking-user tenants disagree on one
operation. In most deployments this is never legitimate, which makes it a high-precision signal;
`crossing_type` distinguishes the benign platform-service case from the rest, and `expected` records
what the tenant should have been so the deviation is self-describing.

---

## Encoded / Obfuscated Payload Indicator

> **RFC field:** Encoded / Obfuscated Payload Indicator · **Tier:** MAY · **§6 Input Handling & Trust Provenance** · **Grounding attacks:** `AOC-12`, `TA-03`

A flag plus the decoded form when input contains base64, image-embedded (OCR) or markup "authority"
tags. The decoded hash is the more durable signal, since it lets the same payload be correlated
across sessions and across encodings without retaining the content.

### Emission Rule

Per Appendix D.4, decoded *content* is emitted on events rather than spans. `detected`, `encodings`,
`depth`, `decoded_hash`, `decoded_length`, `source_field` and `inspected_after_decode` are bounded and
MAY be promoted to the span; `decoded_form` MUST NOT be, unless a retention policy explicitly permits
it.

| Attribute | Type | Requirement | Description | Compliance |
|-----------|------|-------------|-------------|------------|
| `security.obfuscation.detected` | boolean | **Recommended** | Whether an encoded or obfuscated construct was found | OWASP LLM01, MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `security.obfuscation.encodings` | string[] | **Recommended** | Encodings identified: `"base64"`, `"hex"`, `"url"`, `"rot13"`, `"unicode_escape"`, `"homoglyph"`, `"zero_width"`, `"bidi_override"`, `"image_ocr"`, `"markup_authority_tag"`, `"nested"` | OWASP LLM01, MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `security.obfuscation.depth` | int | **Optional** | Number of decode passes needed to reach plaintext. Depth > 1 is itself a signal | OWASP LLM01 |
| `security.obfuscation.decoded_form` | string | **Optional** | The decoded payload. **Policy-gated**; prefer the hash where raw capture is restricted | OWASP LLM01 |
| `security.obfuscation.decoded_hash` | string | **Recommended** | SHA-256 of the decoded payload. Survives redaction and enables cross-session correlation | MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `security.obfuscation.decoded_length` | int | **Optional** | Byte length of the decoded payload | — |
| `security.obfuscation.source_field` | string | **Recommended** | Which field carried the construct (`"gen_ai.input.messages"`, `"gen_ai.tool.call.result"`, `"rag.doc"`) | OWASP LLM01 |
| `security.obfuscation.inspected_after_decode` | boolean | **Recommended** | Whether guardrails re-ran against the decoded form. `false` is the `TA-03` bypass condition | NIST AI RMF GOVERN-1.2, EU AI Act Art.14 |

### Reading the Fields

`inspected_after_decode: false` is the finding. A guardrail that ran against the encoded form and
passed it has not evaluated the payload the model will act on — the verdict applies to the wrapper,
not the contents.

`depth` greater than one indicates nested encoding, which is rarely accidental in user-supplied input.

`decoded_hash` correlates the same payload across sessions and across encodings, which is what makes a
campaign visible when each individual instance looks like a one-off.

---

## Related Documents

- [Attributes Registry](attributes-registry.md) — canonical definition of every attribute above
- [GenAI Spans](gen-ai-spans.md) — security enrichment on inference spans
- [Identity Spans](identity-spans.md) — `identity.authz.*`, delegation, approval
- [Asset Inventory Spans](asset-inventory-spans.md) — `asset.instrumentation.*`, tenancy scoping
- [MCP Spans](mcp-spans.md) — where `supply_chain.runtime.*` attaches for delegated tool execution
- [Events](events.md) — event-shaped counterparts to these span attributes
