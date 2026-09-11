# AI Asset Inventory Span Conventions

AITF defines semantic conventions for AI asset inventory management — covering the registration, discovery, audit, and risk classification of all AI system components. These conventions align with CoSAI's AI Incident Response requirement for maintaining complete inventories of models, datasets, prompts, and infrastructure dependencies.

## Overview

The `asset.*` namespace covers the complete AI asset lifecycle:

| Stage | Span Name | Description |
|-------|-----------|-------------|
| Registration | `asset.register` | Asset registration into inventory |
| Discovery | `asset.discover` | Automated asset discovery and scanning |
| Audit | `asset.audit` | Periodic audit and compliance verification |
| Risk Classification | `asset.classify` | Risk classification (EU AI Act, internal policy) |
| Dependency Mapping | `asset.dependency` | Dependency graph resolution |
| Decommission | `asset.decommission` | Asset retirement and decommissioning |

---

## Span: `asset.register`

Represents the registration of an AI asset (model, dataset, prompt template, vector DB, MCP server, agent) into the organization's AI asset inventory.

### Span Name

Format: `asset.register {asset.type} {asset.name}`

### Span Kind

`INTERNAL`

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.id` | string | Unique asset identifier |
| `asset.name` | string | Human-readable asset name |
| `asset.type` | string | Asset type (see below) |

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.version` | string | Asset version (SemVer or hash) |
| `asset.hash` | string | Content hash for integrity verification |
| `asset.owner` | string | Asset owner (team or individual) |
| `asset.owner_type` | string | `"team"`, `"individual"`, `"organization"` |
| `asset.deployment_environment` | string | Deployment environment |
| `asset.risk_classification` | string | Risk classification |
| `asset.description` | string | Asset description |
| `asset.tags` | string[] | Searchable tags |
| `asset.source_repository` | string | Source code/model repository URL |
| `asset.created_at` | string | Creation timestamp (ISO 8601) |

### Asset Types

| Value | Description |
|-------|-------------|
| `model` | ML/LLM model artifact |
| `dataset` | Training, evaluation, or reference dataset |
| `prompt_template` | Prompt template or system prompt |
| `vector_db` | Vector database or embedding index |
| `mcp_server` | MCP server or tool provider |
| `agent` | Autonomous AI agent |
| `pipeline` | ML/data pipeline definition |
| `guardrail` | Safety guardrail configuration |
| `embedding_model` | Embedding model (distinct from inference model) |
| `knowledge_base` | RAG knowledge base or document collection |

### Deployment Environments

| Value | Description |
|-------|-------------|
| `production` | Production environment |
| `staging` | Pre-production staging |
| `development` | Development/testing |
| `shadow` | Shadow/dark-launch environment |

### Risk Classifications (EU AI Act aligned)

| Value | Description |
|-------|-------------|
| `unacceptable` | Unacceptable risk — prohibited under EU AI Act |
| `high_risk` | High risk — subject to conformity assessment |
| `limited_risk` | Limited risk — transparency obligations |
| `minimal_risk` | Minimal risk — no specific obligations |
| `systemic` | Systemic risk — GPAI models with systemic risk |
| `not_classified` | Not yet classified |

---

## Span: `asset.discover`

Represents automated asset discovery — scanning infrastructure for unregistered or shadow AI assets.

### Span Name

Format: `asset.discover {asset.discovery.scope}`

### Span Kind

`INTERNAL`

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.discovery.scope` | string | Discovery scope (`"cluster"`, `"namespace"`, `"environment"`, `"organization"`) |
| `asset.discovery.method` | string | `"api_scan"`, `"network_scan"`, `"registry_sync"`, `"log_analysis"`, `"manual"` |

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.discovery.assets_found` | int | Number of assets discovered |
| `asset.discovery.new_assets` | int | Number of previously unknown assets |
| `asset.discovery.shadow_assets` | int | Number of unregistered (shadow) AI assets |
| `asset.discovery.status` | string | `"completed"`, `"partial"`, `"failed"` |

---

## Span: `asset.audit`

Represents a periodic audit of an AI asset — verifying integrity, compliance, and operational status.

### Span Name

Format: `asset.audit {asset.type} {asset.id}`

### Span Kind

`INTERNAL`

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.id` | string | Asset being audited |
| `asset.audit.type` | string | Audit type (see below) |
| `asset.audit.result` | string | `"pass"`, `"fail"`, `"warning"`, `"not_applicable"` |

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.audit.auditor` | string | Auditor identity (human or automated) |
| `asset.audit.framework` | string | Compliance framework (`"eu_ai_act"`, `"nist_ai_rmf"`, `"iso_42001"`, `"soc2"`) |
| `asset.audit.findings` | string | JSON array of audit findings |
| `asset.audit.last_audit_time` | string | Previous audit timestamp |
| `asset.audit.next_audit_due` | string | Next scheduled audit timestamp |
| `asset.audit.risk_score` | double | Calculated risk score (0-100) |
| `asset.audit.integrity_verified` | boolean | Whether integrity hash matches |
| `asset.audit.compliance_status` | string | `"compliant"`, `"non_compliant"`, `"partially_compliant"` |

### Audit Types

| Value | Description |
|-------|-------------|
| `integrity` | Hash/signature verification |
| `compliance` | Regulatory compliance check |
| `access_review` | Access control and permission audit |
| `drift` | Model drift and performance audit |
| `security` | Security vulnerability assessment |
| `lineage` | Data/model lineage verification |
| `full` | Comprehensive full audit |

---

## Span: `asset.classify`

Represents risk classification of an AI asset — assigning regulatory risk levels per EU AI Act, internal policy, or other frameworks.

### Span Name

Format: `asset.classify {asset.id}`

### Span Kind

`INTERNAL`

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.id` | string | Asset being classified |
| `asset.risk_classification` | string | Assigned risk level (see Risk Classifications above) |
| `asset.classification.framework` | string | Framework used for classification |

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.classification.previous` | string | Previous risk classification |
| `asset.classification.reason` | string | Reason for classification |
| `asset.classification.assessor` | string | Person/system performing classification |
| `asset.classification.use_case` | string | Intended use case (affects risk level) |
| `asset.classification.affected_persons` | string | `"employees"`, `"consumers"`, `"public"`, `"children"` |
| `asset.classification.sector` | string | Deployment sector (affects risk — healthcare, finance, etc.) |
| `asset.classification.biometric` | boolean | Uses biometric data |
| `asset.classification.autonomous_decision` | boolean | Makes autonomous decisions affecting rights |

---

## Span: `asset.dependency`

Represents dependency mapping — resolving the dependency graph of an AI asset.

### Span Name

Format: `asset.dependency {asset.id}`

### Span Kind

`INTERNAL`

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.id` | string | Asset whose dependencies are being resolved |
| `asset.dependency.operation` | string | `"resolve"`, `"update"`, `"validate"` |

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.dependency.count` | int | Total dependency count |
| `asset.dependency.direct_count` | int | Direct dependency count |
| `asset.dependency.transitive_count` | int | Transitive dependency count |
| `asset.dependency.vulnerable_count` | int | Dependencies with known vulnerabilities |
| `asset.dependency.graph` | string | JSON dependency graph |

---

## Span: `asset.decommission`

Represents the decommissioning/retirement of an AI asset.

### Span Name

Format: `asset.decommission {asset.type} {asset.id}`

### Span Kind

`INTERNAL`

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.id` | string | Asset being decommissioned |
| `asset.decommission.reason` | string | `"replaced"`, `"deprecated"`, `"security_risk"`, `"compliance"`, `"end_of_life"` |

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.decommission.replacement_id` | string | Replacement asset ID (if applicable) |
| `asset.decommission.data_retention` | string | `"purge"`, `"archive"`, `"retain"` |
| `asset.decommission.approved_by` | string | Approver identity |
| `asset.decommission.effective_date` | string | Effective decommission date |

---

## Cross-Cutting: Tenancy Scoping [RFC v0.4 gap closure]

> **RFC field:** Organization / Tenant ID · **Tier:** SHOULD · **§5 Application & Agent Reasoning Core** · **Grounding attacks:** `AOC-02`, `AOC-05`, `TA-05`, `TA-01`, `TA-11` · **ODIS:** `trust_domain` (6.1)

Tenancy is the boundary that decides whether an observed event is routine or a breach. The same retrieval, the same memory write, the same peer-agent call is unremarkable within a tenant and a serious isolation failure across one — and a span carrying no tenant identifier cannot be assessed either way. The RFC marks this field **‡ cross-cutting** for that reason: it belongs on every content-bearing and every security-relevant signal, not on one span type.

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.tenant.id` | string | Tenant the asset belongs to — the primary isolation-boundary identifier |
| `asset.organization.id` | string | Organization owning the asset (ODIS `owner_ref`). May equal `tenant.id`; MUST still be emitted |
| `asset.organization.name` | string | Human-readable organization name |
| `asset.tenant.tier` | string | `"dedicated"`, `"pooled"`, `"shared_infrastructure"`, `"shared_model"` |
| `asset.tenant.isolation_boundary` | string | `"process"`, `"container"`, `"namespace"`, `"database"`, `"row_level"`, `"logical_only"` |
| `asset.tenant.data_residency` | string | Required data-residency region |

Three further identifiers carry the same boundary onto the acting signals, and are defined in the registry rather than here because they live in other namespaces: `gen_ai.agent.tenant.id` (the tenant the agent acts for), `gen_ai.conversation.tenant.id` (the tenant owning the session), and `gen_ai.data_source.tenant.id` (the tenant owning what was read). `user.tenant.id` records the human principal's tenant. The comparisons between them are where the findings are — an agent tenant differing from a conversation tenant is a cross-tenant execution; an agent tenant differing from a data-source tenant on a retrieval is a cross-tenant read.

**Emission rule.** Where a process serves exactly one tenant, these SHOULD be OTel **Resource** attributes. In any multi-tenant process, where the tenant varies per request, they MUST be span attributes. Emitting them only at Resource level in a pooled deployment records the deployment's tenancy rather than the request's, which is precisely the case the field exists to cover.

The detection counterpart — whether a crossing was *observed*, and whether it was permitted — lives at `security.tenant.*` in the [attributes registry](attributes-registry.md#securitytenant-rfc-v04-gap-closure). `asset.tenant.*` supplies the identifiers; `security.tenant.*` supplies the verdict.

---

## Event: `asset.capability.changed` [RFC v0.4 gap closure]

> **RFC field:** Capability-Set Change Event · **Tier:** MUST · **§14 Asset Inventory & Fleet Aggregates** · **Grounding attacks:** `AOC-09`, `AOC-10`, `TA-06`, `AOC-02` · **ODIS:** `approved_software_refs` (6.1)

An agent's capability set — the tools, servers, skills, scopes, data sources, models, and peer agents it can reach — is its blast radius. That set changes without any code change and without any deployment: an MCP server advertises a new tool on the next handshake, a marketplace pushes a skill update, a scope is broadened in a config file. `asset.register` and `asset.discover` capture point-in-time inventory, which means a change is only recoverable by an analyst diffing two snapshots — and a snapshot diff loses the actor, the source, and the approval status, which are the three fields that determine whether an expansion was legitimate.

### Emission Rule (RFC Appendix D.2)

This is a state transition and MUST be emitted as an event carrying both the before and after sets, so the delta is recoverable from a single record. It MUST NOT be left to be inferred from inventory snapshots.

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.id` | string | Agent or asset whose capability set changed |
| `asset.capability.change_type` | string | `"added"`, `"removed"`, `"modified"`, `"scope_expanded"`, `"scope_reduced"`, `"reconfigured"` |
| `asset.capability.set.hash` | string | Digest of the canonicalised current capability set — the single value to alert on |
| `asset.capability.set.previous_hash` | string | Digest of the previously observed set |
| `asset.capability.added` | string[] | Capabilities gained (tool names, server IDs, skill IDs, scopes) |
| `asset.capability.change_source` | string | `"deployment"`, `"config"`, `"server_advertised"`, `"marketplace_update"`, `"runtime_discovery"`, `"agent_self_modification"` |
| `asset.capability.approved` | boolean | Whether the new set was reviewed and approved |
| `asset.capability.detected_at` | string | RFC 3339 timestamp the change was observed |

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.capability.set.size` | int | Number of capabilities in the current set |
| `asset.capability.removed` | string[] | Capabilities lost |
| `asset.capability.modified` | string[] | Capabilities whose definition changed while keeping their name |
| `asset.capability.category` | string[] | `"tool"`, `"mcp_server"`, `"skill"`, `"scope"`, `"data_source"`, `"model"`, `"peer_agent"` |
| `asset.capability.risk_delta` | string | `"increased"`, `"decreased"`, `"unchanged"` |
| `asset.capability.privileged_added` | string[] | Newly gained capabilities classed as privileged or destructive |
| `asset.capability.change_actor` | string | Principal responsible for the change |

### Optional Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.capability.approval_ref` | string | Approval record reference (`identity.approval.id` where applicable) |
| `asset.capability.detection_method` | string | `"registration"`, `"periodic_scan"`, `"first_use"`, `"handshake_diff"` |

### Reading the Fields

`change_source` is the field that separates governed expansion from ungoverned. `"deployment"` and `"config"` are reviewed paths with a change record somewhere outside the agent. `"server_advertised"`, `"marketplace_update"`, and `"runtime_discovery"` are not — the capability set grew because a remote party said so. `"agent_self_modification"` means the agent expanded its own reach, which under `AOC-09` should be treated as requiring justification in every case.

The combination to alert on is `change_type: "added"` or `"scope_expanded"` with `approved: false` and a `change_source` outside the reviewed set. `privileged_added` being non-empty raises that from a finding to an incident. Note that `asset.capability.modified` — a capability keeping its name while changing its definition — is the fleet-level expression of the same rug-pull that `mcp.tool.definition.hash` detects per-invocation; the two should agree, and a disagreement means one of the two observation points is not seeing the whole system.

---

## Instrumentation Coverage & Hook Attestation [RFC v0.4 gap closure]

> **RFC field:** Instrumentation Coverage / Hook Attestation · **Tier:** SHOULD · **§15 Observability-Plane Integrity** · **Grounding attacks:** `AOC-10`, `AOC-01`, `AOC-09`

Every other attribute in AITF answers "what happened?". These answer "would we have seen it?". Under the RFC's assume-breach posture, an adversary's first move against an instrumented system is to reduce what the instrumentation sees: disable a hook, downgrade the SDK, replace a span processor, point the exporter at nothing. Absence of telemetry then becomes indistinguishable from absence of activity, and every negative finding drawn from that asset silently loses its basis.

Declaring coverage converts silence into a checkable claim. If coverage is attested at 1.0 and a code path emits nothing, that is a contradiction an analyst can act on rather than an assumption they have to make.

### Attributes

These attach to `asset.register` and `asset.audit` spans, and SHOULD additionally be carried as OTel Resource attributes so that coverage is knowable for a process even when it emits no spans at all.

| Attribute | Type | Notes |
|-----------|------|-------|
| `asset.instrumentation.enabled` | boolean | Whether AITF instrumentation is active |
| `asset.instrumentation.version` | string | Version in force — a downgrade is a coverage-reduction signal |
| `asset.instrumentation.sdk` | string | `"aitf-python"`, `"aitf-go"`, `"otel-auto"`, … |
| `asset.instrumentation.hooks.declared` | string[] | Hook points claimed: `"gen_ai.inference"`, `"tool.call"`, `"mcp.request"`, `"rag.retrieve"`, `"memory.write"`, `"agent.delegate"`, `"guardrail"`, `"auth"` |
| `asset.instrumentation.hooks.active` | string[] | Hook points verified active at runtime |
| `asset.instrumentation.hooks.missing` | string[] | Declared but not active — the enumerated blind spots |
| `asset.instrumentation.coverage_ratio` | double | Fraction of declared hooks active (0.0–1.0) |
| `asset.instrumentation.uninstrumented_paths` | string[] | Known code paths that emit no telemetry |
| `asset.instrumentation.attestation.method` | string | `"self_declared"`, `"startup_probe"`, `"runtime_verified"`, `"signed_manifest"`, `"tee_attested"` (ODIS `attestation_method`) |
| `asset.instrumentation.attestation.verified` | boolean | Whether the claim was independently verified |
| `asset.instrumentation.attestation.at` | string | RFC 3339 timestamp of the assessment |
| `asset.instrumentation.attestation.signature` | string | Signature over the coverage manifest |
| `asset.instrumentation.tamper_detected` | boolean | Hooks disabled, replaced, or downgraded after startup |
| `asset.instrumentation.tamper_indicator` | string | `"hook_removed"`, `"exporter_disabled"`, `"processor_replaced"`, `"version_downgrade"`, `"config_override"` |
| `asset.instrumentation.exporter.configured` | boolean | Whether an exporter is configured |
| `asset.instrumentation.exporter.reachable` | boolean | Whether the collector endpoint is reachable |
| `asset.instrumentation.dropped_spans` | int | Spans lost to queue overflow or export failure since last report |

### Reading the Fields

`coverage_ratio` bounds the confidence of every negative finding from the asset. At 0.75, "we saw no memory writes" means "we saw no memory writes through the three hooks that were running" — a materially weaker statement, and `hooks.missing` says exactly which one was not.

`attestation.method: "self_declared"` with `attestation.verified` absent is the default state, and under the RFC §16 provenance rule it must be read as unverified: the process is asserting its own coverage, which an adversary with code execution in that process can assert equally well. `"signed_manifest"` and `"tee_attested"` are the methods that survive that assumption.

`exporter.configured: true` with `exporter.reachable: false`, or a rising `dropped_spans`, describes instrumentation that is running and emitting into a void. Functionally this is uninstrumented, but it looks healthy from inside the process, so it must be reported rather than inferred.

**Relationship to `observability.sequence.*`.** These two groups are complementary and neither substitutes for the other: `asset.instrumentation.*` establishes whether a signal would ever have been produced, and `observability.sequence.*` establishes whether a signal that was produced survived the journey to the backend intact.

---

## Examples

### Example 1: Registering a production model

```
Span: asset.register model customer-support-llama-70b
  asset.id: "model-cs-llama70b-v3"
  asset.name: "customer-support-llama-70b"
  asset.type: "model"
  asset.version: "3.1.0"
  asset.hash: "sha256:abc123def456"
  asset.owner: "ml-platform-team"
  asset.owner_type: "team"
  asset.deployment_environment: "production"
  asset.risk_classification: "high_risk"
  asset.source_repository: "https://registry.internal/models/cs-llama70b"
  asset.tags: ["customer-support", "llama", "fine-tuned"]
```

### Example 2: Shadow AI discovery scan

```
Span: asset.discover organization
  asset.discovery.scope: "organization"
  asset.discovery.method: "api_scan"
  asset.discovery.assets_found: 47
  asset.discovery.new_assets: 3
  asset.discovery.shadow_assets: 2
  asset.discovery.status: "completed"
```

### Example 3: EU AI Act risk classification

```
Span: asset.classify model-hiring-screener-v1
  asset.id: "model-hiring-screener-v1"
  asset.risk_classification: "high_risk"
  asset.classification.framework: "eu_ai_act"
  asset.classification.reason: "Employment context — CV screening affects natural persons"
  asset.classification.use_case: "automated resume screening"
  asset.classification.affected_persons: "consumers"
  asset.classification.sector: "hr_recruitment"
  asset.classification.autonomous_decision: true
```

### Example 4: Periodic compliance audit

```
Span: asset.audit model model-cs-llama70b-v3
  asset.id: "model-cs-llama70b-v3"
  asset.audit.type: "compliance"
  asset.audit.result: "warning"
  asset.audit.framework: "eu_ai_act"
  asset.audit.risk_score: 72.5
  asset.audit.integrity_verified: true
  asset.audit.compliance_status: "partially_compliant"
  asset.audit.last_audit_time: "2026-01-15T10:00:00Z"
  asset.audit.next_audit_due: "2026-04-15T10:00:00Z"
  asset.audit.findings: "[{\"finding\": \"Missing bias evaluation for protected groups\", \"severity\": \"high\"}]"
```
