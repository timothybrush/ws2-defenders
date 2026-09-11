# AITF Event Conventions

AITF defines events for security findings, compliance actions, and other discrete occurrences that need to be captured alongside trace spans.

## Security Events

### `security.threat_detected`

Emitted when a security threat is detected in AI input or output.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `security.threat_type` | string | Type of threat | Yes |
| `security.owasp_category` | string | OWASP LLM category | Yes |
| `security.risk_level` | string | Risk level | Yes |
| `security.risk_score` | double | Risk score (0-100) | Yes |
| `security.confidence` | double | Detection confidence (0-1) | Yes |
| `security.detection_method` | string | How detected | Recommended |
| `security.blocked` | boolean | Whether blocked | Recommended |
| `security.details` | string | Threat details (JSON) | Recommended |

#### Threat Types

| Value | OWASP | Description |
|-------|-------|-------------|
| `prompt_injection` | LLM01 | Direct or indirect prompt injection |
| `sensitive_data_exposure` | LLM02 | Sensitive information disclosure |
| `supply_chain` | LLM03 | Compromised training data or models |
| `data_poisoning` | LLM04 | Data poisoning attempts |
| `improper_output` | LLM05 | Improper output handling |
| `excessive_agency` | LLM06 | Excessive autonomy or permissions |
| `system_prompt_leak` | LLM07 | System prompt leakage |
| `vector_data_weakness` | LLM08 | Vector/embedding weaknesses |
| `misinformation` | LLM09 | Generated misinformation |
| `unbounded_consumption` | LLM10 | Resource exhaustion / DoS |
| `jailbreak` | LLM01 | Jailbreak attempt |
| `data_exfiltration` | LLM02 | Data exfiltration attempt |
| `model_theft` | LLM03 | Model extraction attempt |

### `security.pii_detected`

Emitted when PII is detected in content.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `security.pii.types` | string[] | Types of PII found | Yes |
| `security.pii.count` | int | Number of instances | Yes |
| `security.pii.action` | string | Action taken | Yes |
| `security.pii.location` | string | `"input"`, `"output"`, `"tool_result"` | Recommended |

#### PII Types

| Value | Description |
|-------|-------------|
| `email` | Email addresses |
| `phone` | Phone numbers |
| `ssn` | Social Security Numbers |
| `credit_card` | Credit card numbers |
| `api_key` | API keys and tokens |
| `password` | Passwords |
| `address` | Physical addresses |
| `ip_address` | IP addresses |
| `name` | Personal names |
| `dob` | Dates of birth |
| `passport` | Passport numbers |
| `driver_license` | Driver's license numbers |
| `jwt` | JWT tokens |

### `security.guardrail_triggered`

Emitted when a guardrail check produces a result.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `security.guardrail.name` | string | Guardrail name | Yes |
| `security.guardrail.type` | string | `"input"`, `"output"`, `"both"` | Yes |
| `security.guardrail.result` | string | `"pass"`, `"fail"`, `"warn"` | Yes |
| `security.guardrail.provider` | string | Guardrail provider | Recommended |
| `security.guardrail.policy` | string | Policy name | Recommended |
| `security.guardrail.details` | string | Details (JSON) | Recommended |

---

## Compliance Events

### `compliance.control_mapped`

Emitted when an AI event is mapped to compliance controls.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `compliance.frameworks` | string[] | Mapped frameworks | Yes |
| `compliance.event_type` | string | Type of AI event | Yes |
| `compliance.controls` | string | All controls (JSON) | Recommended |

### `compliance.violation_detected`

Emitted when a potential compliance violation is detected.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `compliance.framework` | string | Framework | Yes |
| `compliance.control` | string | Violated control | Yes |
| `compliance.severity` | string | Violation severity | Yes |
| `compliance.details` | string | Violation details | Recommended |
| `compliance.remediation` | string | Suggested remediation | Recommended |

---

## Cost Events

### `cost.budget_warning`

Emitted when budget utilization reaches a threshold.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `cost.budget.limit` | double | Budget limit (USD) | Yes |
| `cost.budget.used` | double | Budget used (USD) | Yes |
| `cost.budget.remaining` | double | Budget remaining (USD) | Yes |
| `cost.budget.utilization_pct` | double | Utilization percentage | Yes |
| `cost.budget.threshold` | string | `"75%"`, `"90%"`, `"100%"` | Yes |
| `cost.attribution.project` | string | Project | Recommended |

### `cost.budget_exceeded`

Emitted when a budget limit is exceeded.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `cost.budget.limit` | double | Budget limit (USD) | Yes |
| `cost.budget.used` | double | Budget used (USD) | Yes |
| `cost.budget.overage` | double | Amount over budget (USD) | Yes |
| `cost.attribution.project` | string | Project | Recommended |
| `cost.action` | string | `"warn"`, `"throttle"`, `"block"` | Recommended |

---

## Quality Events

### `quality.low_confidence`

Emitted when model confidence is below threshold.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `quality.confidence` | double | Confidence score (0-1) | Yes |
| `quality.threshold` | double | Threshold value | Yes |
| `gen_ai.request.model` | string | Model ID | Recommended |

### `quality.hallucination_detected`

Emitted when potential hallucination is detected.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `quality.hallucination_score` | double | Score (0-1) | Yes |
| `quality.threshold` | double | Threshold value | Yes |
| `quality.details` | string | Details (JSON) | Recommended |

### `quality.user_feedback`

Emitted when user provides feedback.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `quality.feedback.rating` | double | Rating (1-5) | Conditional |
| `quality.feedback.thumbs` | string | `"up"`, `"down"` | Conditional |
| `quality.feedback.comment` | string | Free-text comment | Recommended |

---

## Agent Events

### `gen_ai.agent.error_recovery`

Emitted when an agent encounters and recovers from an error.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `gen_ai.agent.name` | string | Agent name | Yes |
| `gen_ai.agent.error.type` | string | Error type | Yes |
| `gen_ai.agent.error.message` | string | Error message | Yes |
| `gen_ai.agent.error.recovery_action` | string | Recovery action taken | Recommended |
| `gen_ai.agent.error.retry_count` | int | Retry count | Recommended |

### `gen_ai.agent.human_approval`

Emitted when agent requests/receives human approval.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `gen_ai.agent.name` | string | Agent name | Yes |
| `gen_ai.agent.approval.action` | string | Action requiring approval | Yes |
| `gen_ai.agent.approval.status` | string | `"requested"`, `"approved"`, `"denied"` | Yes |
| `gen_ai.agent.approval.approver` | string | Approver identity | Recommended |
| `gen_ai.agent.approval.reason` | string | Reason for decision | Recommended |

---

## Model Operations (LLMOps/MLOps) Events

### `model_ops.training_completed`

Emitted when a training or fine-tuning run completes.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `model_ops.training.run_id` | string | Training run ID | Yes |
| `model_ops.training.type` | string | Training type | Yes |
| `model_ops.training.base_model` | string | Base model | Yes |
| `model_ops.training.status` | string | `"completed"`, `"failed"`, `"cancelled"` | Yes |
| `model_ops.training.loss_final` | double | Final training loss | Recommended |
| `model_ops.training.output_model.id` | string | Output model ID | Recommended |
| `model_ops.training.compute.gpu_hours` | double | GPU hours consumed | Recommended |

### `model_ops.evaluation_completed`

Emitted when a model evaluation run completes.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `model_ops.evaluation.run_id` | string | Evaluation run ID | Yes |
| `model_ops.evaluation.model_id` | string | Model evaluated | Yes |
| `model_ops.evaluation.type` | string | Evaluation type | Yes |
| `model_ops.evaluation.pass` | boolean | Passed quality gates | Yes |
| `model_ops.evaluation.metrics` | string | JSON metric results | Recommended |
| `model_ops.evaluation.regression_detected` | boolean | Regression found | Recommended |

### `model_ops.model_promoted`

Emitted when a model is promoted through lifecycle stages.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `model_ops.registry.model_id` | string | Model ID | Yes |
| `model_ops.registry.stage` | string | New stage | Yes |
| `model_ops.registry.previous_stage` | string | Previous stage | Yes |
| `model_ops.registry.model_alias` | string | Alias assigned | Recommended |
| `model_ops.registry.approval.approver` | string | Approver | Recommended |

### `model_ops.deployment_completed`

Emitted when a model deployment completes (success or failure).

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `model_ops.deployment.id` | string | Deployment ID | Yes |
| `model_ops.deployment.model_id` | string | Model deployed | Yes |
| `model_ops.deployment.strategy` | string | Deployment strategy | Yes |
| `model_ops.deployment.status` | string | `"completed"`, `"failed"`, `"rolled_back"` | Yes |
| `model_ops.deployment.environment` | string | Target environment | Recommended |
| `model_ops.deployment.canary_percent` | double | Canary traffic % | Recommended |

### `model_ops.drift_detected`

Emitted when model monitoring detects drift above threshold.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `model_ops.monitoring.model_id` | string | Monitored model | Yes |
| `model_ops.monitoring.drift_type` | string | Type of drift | Yes |
| `model_ops.monitoring.drift_score` | double | Drift magnitude (0-1) | Yes |
| `model_ops.monitoring.result` | string | Alert level | Yes |
| `model_ops.monitoring.baseline_value` | double | Baseline value | Recommended |
| `model_ops.monitoring.metric_value` | double | Current value | Recommended |
| `model_ops.monitoring.action_triggered` | string | Automated action | Recommended |

### `model_ops.fallback_triggered`

Emitted when a model serving fallback occurs.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `model_ops.serving.fallback.original_model` | string | Original model | Yes |
| `model_ops.serving.fallback.final_model` | string | Fallback model used | Yes |
| `model_ops.serving.fallback.trigger` | string | Trigger reason | Yes |
| `model_ops.serving.fallback.depth` | int | Fallback depth | Recommended |
| `model_ops.serving.cache.cost_saved_usd` | double | Cost saved by cache | Recommended |

### `model_ops.prompt_promoted`

Emitted when a prompt version is promoted to a deployment label.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `model_ops.prompt.name` | string | Prompt name | Yes |
| `model_ops.prompt.version` | string | Promoted version | Yes |
| `model_ops.prompt.label` | string | Target label | Yes |
| `model_ops.prompt.previous_version` | string | Previous version at label | Recommended |
| `model_ops.prompt.evaluation.score` | double | Evaluation score | Recommended |

---

## Identity Events

### `identity.created`

Emitted when a new agent identity is created.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `identity.agent_id` | string | Agent identity ID | Yes |
| `identity.agent_name` | string | Agent name | Yes |
| `identity.type` | string | Identity type | Yes |
| `identity.provider` | string | Identity provider | Recommended |
| `identity.owner` | string | Identity owner | Recommended |
| `identity.credential_type` | string | Credential type | Recommended |
| `identity.ttl_seconds` | int | TTL | Recommended |

### `identity.auth_failed`

Emitted when agent authentication fails.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `identity.agent_id` | string | Agent identity ID | Yes |
| `identity.agent_name` | string | Agent name | Yes |
| `identity.auth.method` | string | Auth method used | Yes |
| `identity.auth.result` | string | Failure type | Yes |
| `identity.auth.failure_reason` | string | Failure reason | Yes |
| `identity.auth.target_service` | string | Target service | Recommended |

### `identity.authz_denied`

Emitted when an authorization request is denied.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `identity.agent_id` | string | Agent identity ID | Yes |
| `identity.agent_name` | string | Agent name | Yes |
| `identity.authz.resource` | string | Resource requested | Yes |
| `identity.authz.action` | string | Action requested | Yes |
| `identity.authz.deny_reason` | string | Denial reason | Yes |
| `identity.authz.policy_id` | string | Policy that denied | Recommended |

### `identity.delegation_created`

Emitted when credentials are delegated between agents.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `identity.delegation.delegator` | string | Delegating agent | Yes |
| `identity.delegation.delegatee` | string | Receiving agent | Yes |
| `identity.delegation.type` | string | Delegation type | Yes |
| `identity.delegation.scope_delegated` | string[] | Delegated scopes | Yes |
| `identity.delegation.chain_depth` | int | Chain depth | Recommended |
| `identity.delegation.scope_attenuated` | boolean | Scope was reduced | Recommended |
| `identity.delegation.ttl_seconds` | int | Delegation TTL | Recommended |

### `identity.privilege_escalation`

Emitted when potential privilege escalation is detected — an agent attempts to access resources beyond its delegated scope.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `identity.agent_id` | string | Agent identity ID | Yes |
| `identity.agent_name` | string | Agent name | Yes |
| `identity.authz.resource` | string | Resource attempted | Yes |
| `identity.authz.action` | string | Action attempted | Yes |
| `identity.authz.scope_required` | string[] | Scopes required | Yes |
| `identity.authz.scope_present` | string[] | Scopes present | Yes |
| `identity.delegation.chain` | string[] | Delegation chain | Recommended |

### `identity.credential_rotated`

Emitted when agent credentials are rotated.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `identity.agent_id` | string | Agent identity ID | Yes |
| `identity.credential_type` | string | Credential type | Yes |
| `identity.auto_rotate` | boolean | Was auto-rotation | Yes |
| `identity.expires_at` | string | New expiration | Recommended |

### `identity.revoked`

Emitted when an agent identity is revoked.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `identity.agent_id` | string | Agent identity ID | Yes |
| `identity.agent_name` | string | Agent name | Yes |
| `identity.status` | string | New status (`"revoked"`) | Yes |
| `identity.previous_status` | string | Previous status | Yes |
| `identity.lifecycle.operation` | string | `"revoke"` | Yes |

### `identity.trust_established`

Emitted when trust is established between two agents.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `identity.agent_name` | string | This agent | Yes |
| `identity.trust.peer_agent` | string | Peer agent | Yes |
| `identity.trust.method` | string | Trust method | Yes |
| `identity.trust.result` | string | Trust result | Yes |
| `identity.trust.trust_level` | string | Trust level | Recommended |
| `identity.trust.cross_domain` | boolean | Cross-domain | Recommended |

### `identity.session_hijack_detected`

Emitted when potential session hijacking is detected.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `identity.agent_id` | string | Agent identity ID | Yes |
| `identity.session.id` | string | Session ID | Yes |
| `identity.session.ip_address` | string | Anomalous source IP | Recommended |
| `identity.session.user_agent` | string | Anomalous user agent | Recommended |

### `identity.approval.requested` [RFC v0.4 gap closure]

Emitted when a human is asked to approve an operation. Pairs with `identity.approval.decided` via `identity.approval.id`. See [`identity-spans.md`](identity-spans.md#human-approval--elicitation-rfc-v04-gap-closure).

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `identity.approval.id` | string | Correlation ID joining request to decision | Yes |
| `identity.approval.required` | boolean | Whether approval was required | Yes |
| `identity.approval.status` | string | `"pending"` at request time | Yes |
| `identity.approval.operation` | string | Operation being approved, as shown to the approver | Yes |
| `identity.approval.prompt_hash` | string | SHA-256 of the exact text shown to the approver | Yes |
| `identity.approval.requested_at` | string | RFC 3339 request timestamp | Yes |
| `identity.approval.trigger` | string | `"policy"`, `"risk_score"`, `"scope_escalation"`, `"first_use"`, `"destructive_action"`, `"elicitation"` | Recommended |
| `identity.approval.channel` | string | `"cli"`, `"web_ui"`, `"chat"`, `"email"`, `"mcp_elicitation"`, `"api"` | Recommended |
| `identity.approval.timeout_ms` | int | Configured approval timeout | Optional |
| `identity.approval.timeout_action` | string | `"deny"`, `"allow"`, `"escalate"` | Recommended |
| `identity.approval.prior_denials` | int | Prior denials for this operation in the session | Optional |
| `identity.approval.elicitation.fields` | string[] | Field names solicited — never values | Optional |
| `identity.approval.elicitation.schema_hash` | string | Digest of the requested-input schema | Optional |

### `identity.approval.decided` [RFC v0.4 gap closure]

Emitted when the approval outcome is known. Carries the same `identity.approval.id` as the corresponding request. Absence of this event MUST NOT be read as a denial — see `identity.approval.status`.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `identity.approval.id` | string | Correlation ID from the request | Yes |
| `identity.approval.status` | string | `"resolved"` or `"expired"` | Yes |
| `identity.approval.decision` | string | `"approved"`, `"denied"`, `"timeout"`, `"auto_approved"`, `"bypassed"`, `"cancelled"` | Yes |
| `identity.approval.approver` | string | Resolved approver identity — not a session token | Yes |
| `identity.approval.approver_type` | string | `"human"`, `"policy"`, `"automation"`, `"none"` | Yes |
| `identity.approval.scope_binding.result` | string | `"match"`, `"mismatch"`, `"not_checked"` | Yes |
| `identity.approval.decided_at` | string | RFC 3339 decision timestamp | Yes |
| `identity.approval.latency_ms` | int | Request-to-decision time — the approval-fatigue signal | Yes |
| `identity.approval.approver_verified` | boolean | Approver authenticated at approval time | Recommended |
| `identity.approval.operation_hash` | string | SHA-256 of the operation as executed | Recommended |
| `identity.approval.scope_binding.divergence` | string[] | Arguments differing between approved and executed | Recommended |
| `identity.approval.scope` | string | `"single_use"`, `"session"`, `"always"`, `"time_bounded"` | Recommended |
| `identity.approval.bypass_reason` | string | Present when `decision` is `"bypassed"` | Recommended |
| `identity.approval.auth_method` | string | `"session"`, `"reauth"`, `"mfa"`, `"webauthn"`, `"none"` | Optional |
| `identity.approval.remembered` | boolean | Decision cached and reused without re-prompting | Optional |

`identity.approval.decision`, `.approver_type`, and `.latency_ms` SHOULD also be mirrored onto the span the approval gates, so an enforcement decision is evaluable without a join.

---

## Asset Inventory Events

### `asset.registered`

Emitted when a new AI asset is registered in the inventory.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `asset.id` | string | Asset identifier | Yes |
| `asset.name` | string | Asset name | Yes |
| `asset.type` | string | Asset type | Yes |
| `asset.owner` | string | Asset owner | Yes |
| `asset.deployment_environment` | string | Environment | Recommended |
| `asset.risk_classification` | string | Risk classification | Recommended |

### `asset.shadow_detected`

Emitted when a shadow (unregistered) AI asset is discovered during a scan.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `asset.type` | string | Discovered asset type | Yes |
| `asset.discovery.scope` | string | Discovery scope | Yes |
| `asset.discovery.method` | string | Discovery method | Yes |
| `asset.name` | string | Discovered asset name | Recommended |
| `asset.deployment_environment` | string | Where found | Recommended |

### `asset.audit_failed`

Emitted when an AI asset fails a compliance or integrity audit.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `asset.id` | string | Asset identifier | Yes |
| `asset.audit.type` | string | Audit type | Yes |
| `asset.audit.result` | string | `"fail"` | Yes |
| `asset.audit.framework` | string | Framework | Recommended |
| `asset.audit.findings` | string | JSON findings | Recommended |
| `asset.audit.risk_score` | double | Risk score | Recommended |

### `asset.risk_reclassified`

Emitted when an asset's risk classification changes.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `asset.id` | string | Asset identifier | Yes |
| `asset.risk_classification` | string | New classification | Yes |
| `asset.classification.previous` | string | Previous classification | Yes |
| `asset.classification.framework` | string | Framework | Yes |
| `asset.classification.reason` | string | Reason for change | Recommended |

### `asset.decommissioned`

Emitted when an AI asset is decommissioned.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `asset.id` | string | Asset identifier | Yes |
| `asset.type` | string | Asset type | Yes |
| `asset.decommission.reason` | string | Decommission reason | Yes |
| `asset.decommission.replacement_id` | string | Replacement asset | Recommended |

### `asset.audit_overdue`

Emitted when an asset's audit is overdue.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `asset.id` | string | Asset identifier | Yes |
| `asset.audit.next_audit_due` | string | Overdue audit date | Yes |
| `asset.risk_classification` | string | Risk level | Recommended |
| `asset.deployment_environment` | string | Environment | Recommended |

### `asset.capability.changed` [RFC v0.4 gap closure]

Emitted when an agent's reachable capability set changes — a tool, MCP server, skill, scope, data source, model, or peer agent discovered, added, removed, or modified. MUST NOT be left to be inferred by diffing inventory snapshots; the diff loses the actor, the source, and the approval status. See [`asset-inventory-spans.md`](asset-inventory-spans.md#event-assetcapabilitychanged-rfc-v04-gap-closure).

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `asset.id` | string | Agent or asset whose capability set changed | Yes |
| `asset.capability.change_type` | string | `"added"`, `"removed"`, `"modified"`, `"scope_expanded"`, `"scope_reduced"`, `"reconfigured"` | Yes |
| `asset.capability.set.hash` | string | Digest of the canonicalised current set | Yes |
| `asset.capability.set.previous_hash` | string | Digest of the previously observed set | Yes |
| `asset.capability.added` | string[] | Capabilities gained | Yes |
| `asset.capability.change_source` | string | `"deployment"`, `"config"`, `"server_advertised"`, `"marketplace_update"`, `"runtime_discovery"`, `"agent_self_modification"` | Yes |
| `asset.capability.approved` | boolean | Whether the new set was reviewed and approved | Yes |
| `asset.capability.detected_at` | string | RFC 3339 observation timestamp | Yes |
| `asset.capability.removed` | string[] | Capabilities lost | Recommended |
| `asset.capability.modified` | string[] | Capabilities changed while keeping their name | Recommended |
| `asset.capability.category` | string[] | `"tool"`, `"mcp_server"`, `"skill"`, `"scope"`, `"data_source"`, `"model"`, `"peer_agent"` | Recommended |
| `asset.capability.risk_delta` | string | `"increased"`, `"decreased"`, `"unchanged"` | Recommended |
| `asset.capability.privileged_added` | string[] | Newly gained privileged or destructive capabilities | Recommended |
| `asset.capability.change_actor` | string | Principal responsible for the change | Recommended |
| `asset.capability.set.size` | int | Size of the current set | Recommended |
| `asset.capability.approval_ref` | string | Approval record reference | Optional |
| `asset.capability.detection_method` | string | `"registration"`, `"periodic_scan"`, `"first_use"`, `"handshake_diff"` | Optional |

### `asset.instrumentation.gap_detected` [RFC v0.4 gap closure]

Emitted when declared instrumentation coverage is not matched by active coverage, or when hooks are found to have been disabled, replaced, or downgraded after startup. This is the event that distinguishes "no events" from "not observed". See [`asset-inventory-spans.md`](asset-inventory-spans.md#instrumentation-coverage--hook-attestation-rfc-v04-gap-closure).

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `asset.id` | string | Asset whose coverage is incomplete | Yes |
| `asset.instrumentation.enabled` | boolean | Whether instrumentation is active at all | Yes |
| `asset.instrumentation.version` | string | Instrumentation version in force | Yes |
| `asset.instrumentation.hooks.declared` | string[] | Hook points claimed | Yes |
| `asset.instrumentation.hooks.active` | string[] | Hook points verified active | Yes |
| `asset.instrumentation.hooks.missing` | string[] | The enumerated blind spots | Yes |
| `asset.instrumentation.coverage_ratio` | double | Fraction of declared hooks active (0.0–1.0) | Yes |
| `asset.instrumentation.sdk` | string | SDK providing instrumentation | Recommended |
| `asset.instrumentation.tamper_detected` | boolean | Hooks altered after startup | Recommended |
| `asset.instrumentation.tamper_indicator` | string | `"hook_removed"`, `"exporter_disabled"`, `"processor_replaced"`, `"version_downgrade"`, `"config_override"` | Optional |
| `asset.instrumentation.attestation.method` | string | How the coverage claim was established | Recommended |
| `asset.instrumentation.attestation.verified` | boolean | Whether independently verified | Recommended |
| `asset.instrumentation.uninstrumented_paths` | string[] | Known silent code paths | Recommended |
| `asset.instrumentation.exporter.reachable` | boolean | Whether the collector endpoint is reachable | Optional |
| `asset.instrumentation.dropped_spans` | int | Spans lost since last report | Optional |

---

## A2A Task Lifecycle Events [RFC v0.4 gap closure]

An A2A task is asynchronous and outlives the span that created it, so its transitions cannot be recorded as span attributes alone — the terminal transition will usually belong to a different trace. All events in this section carry `a2a.task.id` and `a2a.task.lifecycle.initiating_trace_id`, which together reconnect an asynchronous completion to its origin. Full attribute definitions are in [`a2a-spans.md`](a2a-spans.md#event-a2atasklifecycle-rfc-v04-gap-closure).

### `a2a.task.lifecycle.transition`

Emitted on every A2A task state transition, including the read-side transitions (`streamed`, `polled`, `resubscribed`) that change no state but reveal who is watching the task.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `a2a.task.id` | string | Task identifier | Yes |
| `a2a.task.lifecycle.event` | string | Transition being recorded | Yes |
| `a2a.task.lifecycle.from_state` | string | State before | Yes |
| `a2a.task.lifecycle.to_state` | string | State after | Yes |
| `a2a.task.lifecycle.actor` | string | Identity that caused the transition | Yes |
| `a2a.task.lifecycle.at` | string | RFC 3339 transition timestamp | Yes |
| `a2a.task.lifecycle.terminal` | boolean | Whether `to_state` is terminal | Yes |
| `a2a.task.lifecycle.initiating_trace_id` | string | Trace that created the task | Yes |
| `a2a.task.lifecycle.actor_type` | string | `"client_agent"`, `"remote_agent"`, `"human"`, `"system"`, `"timeout"`, `"unknown"` | Recommended |
| `a2a.task.lifecycle.transition_valid` | boolean | Permitted by the A2A state machine | Recommended |
| `a2a.task.lifecycle.age_ms` | int | Elapsed time since task creation | Recommended |
| `a2a.task.lifecycle.expected_terminal_by` | string | Deadline for reaching a terminal state | Recommended |
| `a2a.task.lifecycle.delegated_scope` | string[] | Scopes outstanding while the task is live | Recommended |
| `a2a.task.lifecycle.delegation_expires_at` | string | Expiry of the backing authority | Recommended |
| `a2a.task.lifecycle.root_principal` | string | Principal ultimately accountable | Recommended |
| `a2a.task.lifecycle.failure_reason` | string | Reason for `"failed"` or `"rejected"` | Recommended |
| `a2a.task.lifecycle.initiating_span_id` | string | Span that created the task | Recommended |
| `a2a.task.lifecycle.transition_count` | int | Cumulative transitions | Optional |
| `a2a.task.lifecycle.poll_count` | int | `tasks/get` polls observed | Optional |
| `a2a.task.lifecycle.resubscribe_count` | int | Stream re-attachments | Optional |
| `a2a.task.lifecycle.subscriber` | string | Principal attached to the stream | Optional |
| `a2a.task.lifecycle.cancel_requested_by` | string | Principal that requested cancellation | Optional |
| `a2a.task.lifecycle.input_required_reason` | string | Why the task awaits input | Optional |

### `a2a.task.lifecycle.orphaned`

Emitted when a task passes `expected_terminal_by` without reaching a terminal state. This is the only signal for a task that fails silently: no error is raised, the task simply stops being mentioned, and the authority delegated to service it stays outstanding.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `a2a.task.id` | string | Orphaned task | Yes |
| `a2a.task.lifecycle.orphaned` | boolean | `true` | Yes |
| `a2a.task.lifecycle.from_state` | string | Last observed state | Yes |
| `a2a.task.lifecycle.age_ms` | int | Age at detection | Yes |
| `a2a.task.lifecycle.expected_terminal_by` | string | Deadline that was passed | Yes |
| `a2a.task.lifecycle.initiating_trace_id` | string | Trace that created the task | Yes |
| `a2a.task.lifecycle.delegated_scope` | string[] | Authority still outstanding | Recommended |
| `a2a.task.lifecycle.delegation_expires_at` | string | Whether that authority lapses on its own | Recommended |
| `a2a.task.lifecycle.root_principal` | string | Principal accountable for the orphan | Recommended |
| `a2a.task.lifecycle.poll_count` | int | Whether anyone was watching | Optional |

### `a2a.push.config.changed`

Emitted when push notification configuration is set or altered. A change to the callback target on a live task redirects where results are delivered, and SHOULD be treated as an attempted exfiltration until cleared.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `a2a.task.id` | string | Task the configuration applies to | Yes |
| `a2a.push.config.changed` | boolean | Whether the target changed after task creation | Yes |
| `a2a.push.config.url_hash` | string | Digest of the callback URL | Yes |
| `a2a.push.config.in_allowlist` | boolean | Destination is on the egress allowlist | Recommended |
| `a2a.push.config.authenticated` | boolean | Whether the callback is authenticated | Recommended |
| `a2a.push.config.url` | string | Callback URL | Recommended |
| `a2a.push.config.scheme` | string | `"none"`, `"bearer"`, `"hmac"`, `"mtls"` | Optional |
| `a2a.push.config.set_by` | string | Principal that altered the configuration | Optional |

---

## Drift Detection Events

### `drift.detected`

Emitted when model drift is detected above threshold. Provides structured forensic-quality drift analysis beyond the basic `model_ops.drift_detected`.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `drift.model_id` | string | Monitored model | Yes |
| `drift.type` | string | Drift type | Yes |
| `drift.score` | double | Drift magnitude (0–1) | Yes |
| `drift.result` | string | Alert level | Yes |
| `drift.detection_method` | string | Statistical method | Yes |
| `drift.baseline_metric` | double | Baseline value | Recommended |
| `drift.current_metric` | double | Current value | Recommended |
| `drift.p_value` | double | Statistical significance | Recommended |
| `drift.affected_segments` | string[] | Impacted segments | Recommended |
| `drift.reference_dataset` | string | Reference dataset | Recommended |
| `drift.action_triggered` | string | Automated action | Recommended |

### `drift.baseline_updated`

Emitted when a drift baseline is created or refreshed.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `drift.model_id` | string | Model | Yes |
| `drift.baseline.operation` | string | `"create"` or `"refresh"` | Yes |
| `drift.baseline.id` | string | Baseline identifier | Yes |
| `drift.baseline.dataset` | string | Baseline dataset | Recommended |
| `drift.baseline.sample_size` | int | Sample size | Recommended |

### `drift.investigation_completed`

Emitted when a drift investigation completes with root cause analysis.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `drift.model_id` | string | Model investigated | Yes |
| `drift.investigation.root_cause_category` | string | Root cause category | Yes |
| `drift.investigation.severity` | string | Severity | Yes |
| `drift.investigation.blast_radius` | string | Impact scope | Yes |
| `drift.investigation.affected_users_estimate` | int | Affected users | Recommended |
| `drift.investigation.recommendation` | string | Recommendation | Recommended |

### `drift.remediation_completed`

Emitted when a drift remediation action completes.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `drift.model_id` | string | Model remediated | Yes |
| `drift.remediation.action` | string | Action taken | Yes |
| `drift.remediation.status` | string | Outcome status | Yes |
| `drift.remediation.automated` | boolean | Was automated | Yes |
| `drift.remediation.validation_passed` | boolean | Post-validation passed | Recommended |

---

## Memory Security Events

### `memory.poisoning_detected`

Emitted when memory poisoning is detected (unexpected content injection).

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `memory.key` | string | Affected memory key | Yes |
| `memory.store` | string | Memory store | Yes |
| `memory.security.poisoning_score` | double | Poisoning score (0-1) | Yes |
| `memory.provenance` | string | Content provenance | Yes |
| `gen_ai.conversation.id` | string | Session ID | Recommended |

### `memory.integrity_violation`

Emitted when memory content hash does not match expected integrity hash.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `memory.key` | string | Affected memory key | Yes |
| `memory.store` | string | Memory store | Yes |
| `memory.security.integrity_hash` | string | Expected hash | Yes |
| `memory.security.content_hash` | string | Actual hash | Yes |
| `gen_ai.conversation.id` | string | Session ID | Recommended |

### `memory.cross_session_access`

Emitted when a session accesses memory belonging to another session.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `memory.key` | string | Accessed memory key | Yes |
| `memory.store` | string | Memory store | Yes |
| `gen_ai.conversation.id` | string | Accessing session | Yes |

### `memory.growth_anomaly`

Emitted when session memory growth exceeds configured thresholds.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `gen_ai.conversation.id` | string | Session ID | Yes |
| `memory.security.mutation_count` | int | Current entry count | Yes |
| `memory.security.content_size` | int | Current total size | Recommended |

### `memory.untrusted_provenance`

Emitted when memory is written from an untrusted provenance source.

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `memory.key` | string | Memory key | Yes |
| `memory.store` | string | Memory store | Yes |
| `memory.provenance` | string | Untrusted provenance | Yes |
| `gen_ai.conversation.id` | string | Session ID | Recommended |
