# Agentic Identity Span Conventions

AITF defines comprehensive semantic conventions for AI agent identity management — covering the full lifecycle from identity creation through authentication, authorization, delegation, and revocation. These conventions enable secure, auditable, and observable identity operations in single-agent and multi-agent systems.

## Overview

The `identity.*` namespace covers the complete agent identity lifecycle:

| Stage | Span Name | Description |
|-------|-----------|-------------|
| Lifecycle | `identity.lifecycle` | Identity creation, rotation, suspension, revocation |
| Authentication | `identity.authentication` | Agent authentication (OAuth, mTLS, SPIFFE, JWT) |
| Authorization | `identity.authorization` | Permission checks, policy evaluation |
| Delegation | `identity.delegation` | Credential delegation, token exchange, OBO flows |
| Trust | `identity.trust` | Agent-to-agent trust establishment, VC verification |
| Session | `identity.session` | Identity session management |

---

## Span: `identity.lifecycle`

Represents an agent identity lifecycle event (creation, rotation, revocation).

### Span Name

Format: `identity.lifecycle.{identity.lifecycle.operation} {identity.agent_id}`

### Span Kind

`INTERNAL`

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.agent_id` | string | Agent identity identifier |
| `identity.agent_name` | string | Agent name |
| `identity.lifecycle.operation` | string | Lifecycle operation (see below) |

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.type` | string | Identity type (see below) |
| `identity.provider` | string | Identity provider (`"okta"`, `"entra_id"`, `"auth0"`, `"spiffe"`, `"custom"`) |
| `identity.owner` | string | Human or system owner of this identity |
| `identity.owner_type` | string | `"human"`, `"service"`, `"organization"` |
| `identity.credential_type` | string | `"api_key"`, `"oauth_token"`, `"jwt"`, `"mtls_cert"`, `"spiffe_svid"`, `"did"` |
| `identity.credential_id` | string | Credential identifier (not the secret) |
| `identity.expires_at` | string | Credential expiration ISO timestamp |
| `identity.ttl_seconds` | int | Time to live in seconds |
| `identity.auto_rotate` | boolean | Whether credential auto-rotates |
| `identity.rotation_interval_seconds` | int | Rotation interval |
| `identity.scope` | string[] | Granted scopes |
| `identity.tags` | string | JSON-encoded identity metadata tags |
| `identity.status` | string | `"active"`, `"suspended"`, `"revoked"`, `"expired"` |
| `identity.previous_status` | string | Previous status (for transitions) |

### Lifecycle Operations

| Value | Description |
|-------|-------------|
| `create` | Create new agent identity |
| `register` | Register identity with an IdP or service |
| `activate` | Activate a pending identity |
| `rotate` | Rotate credentials (key, cert, token) |
| `suspend` | Temporarily suspend an identity |
| `reactivate` | Reactivate a suspended identity |
| `revoke` | Permanently revoke an identity |
| `expire` | Identity expired automatically |
| `update` | Update identity metadata or scope |

### Identity Types

| Value | Description |
|-------|-------------|
| `persistent` | Long-lived agent identity with stable credentials |
| `ephemeral` | Short-lived identity created for a single task |
| `delegated` | Identity derived from delegation of another identity |
| `federated` | Identity from a federated trust domain |
| `workload` | SPIFFE-style workload identity |

---

## Span: `identity.authentication`

Represents an agent authentication attempt.

### Span Name

Format: `identity.auth {identity.agent_name}`

### Span Kind

`CLIENT`

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.agent_id` | string | Agent identity ID |
| `identity.agent_name` | string | Agent name |
| `identity.auth.method` | string | Authentication method (see below) |
| `identity.auth.result` | string | `"success"`, `"failure"`, `"denied"`, `"expired"`, `"revoked"` |

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.auth.provider` | string | Auth provider service |
| `identity.auth.target_service` | string | Service being authenticated to |
| `identity.auth.failure_reason` | string | Reason for failure (if applicable) |
| `identity.auth.token_type` | string | Token type received (`"bearer"`, `"dpop"`, `"mtls_bound"`) |
| `identity.auth.token_lifetime_seconds` | int | Token lifetime |
| `identity.auth.scope_requested` | string[] | Scopes requested |
| `identity.auth.scope_granted` | string[] | Scopes actually granted |
| `identity.auth.mfa_used` | boolean | Whether MFA was used |
| `identity.auth.continuous` | boolean | Whether this is continuous re-authentication |
| `identity.auth.protocol_version` | string | Auth protocol version |
| `identity.auth.client_id` | string | OAuth client ID |
| `identity.auth.pkce_used` | boolean | Whether PKCE was used |
| `identity.auth.dpop_used` | boolean | Whether DPoP proof was included |

### Authentication Methods

| Value | Description |
|-------|-------------|
| `api_key` | Static API key authentication |
| `oauth2` | OAuth 2.0/2.1 flow |
| `oauth2_pkce` | OAuth 2.1 with PKCE (MCP standard) |
| `jwt_bearer` | JWT Bearer token (RFC 7523) |
| `mtls` | Mutual TLS certificate authentication |
| `spiffe_svid` | SPIFFE Verifiable Identity Document |
| `did_vc` | Decentralized Identifier + Verifiable Credential |
| `http_signature` | HTTP Message Signature (W3C) |
| `token_exchange` | OAuth Token Exchange (RFC 8693) |
| `saml` | SAML assertion |

---

## Span: `identity.authorization`

Represents a permission check or policy evaluation.

### Span Name

Format: `identity.authz {identity.agent_name} -> {identity.authz.resource}`

### Span Kind

`INTERNAL`

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.agent_id` | string | Agent identity ID |
| `identity.agent_name` | string | Agent name |
| `identity.authz.decision` | string | `"allow"`, `"deny"`, `"conditional"` |
| `identity.authz.resource` | string | Resource being accessed |

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.authz.action` | string | Action being performed (`"read"`, `"write"`, `"execute"`, `"delete"`) |
| `identity.authz.policy_engine` | string | Policy engine (`"opa"`, `"cedar"`, `"casbin"`, `"custom"`) |
| `identity.authz.policy_id` | string | Matched policy identifier |
| `identity.authz.policy_version` | string | Policy version |
| `identity.authz.deny_reason` | string | Reason for denial (if denied) |
| `identity.authz.conditions` | string | JSON-encoded conditions (for conditional allow) |
| `identity.authz.scope_required` | string[] | Scopes required for this action |
| `identity.authz.scope_present` | string[] | Scopes present in token |
| `identity.authz.risk_score` | double | Risk-based authorization score (0-100) |
| `identity.authz.context` | string | JSON-encoded authorization context (time, location, etc.) |
| `identity.authz.privilege_level` | string | `"standard"`, `"elevated"`, `"admin"` |
| `identity.authz.jea` | boolean | Whether Just-Enough-Access was applied |
| `identity.authz.time_limited` | boolean | Whether permission is time-limited |
| `identity.authz.expires_at` | string | Permission expiration ISO timestamp |

---

## Span: `identity.delegation`

Represents credential delegation in a multi-agent system — On-Behalf-Of flows, token exchange, and authority attenuation.

### Span Name

Format: `identity.delegate {identity.delegation.delegator} -> {identity.delegation.delegatee}`

### Span Kind

`INTERNAL`

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.delegation.delegator` | string | Agent/user delegating authority |
| `identity.delegation.delegator_id` | string | Delegator identity ID |
| `identity.delegation.delegatee` | string | Agent receiving delegated authority |
| `identity.delegation.delegatee_id` | string | Delegatee identity ID |
| `identity.delegation.type` | string | Delegation type (see below) |

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.delegation.chain` | string[] | Full delegation chain (ordered from origin) |
| `identity.delegation.chain_depth` | int | Depth of delegation chain |
| `identity.delegation.scope_delegated` | string[] | Scopes being delegated |
| `identity.delegation.scope_attenuated` | boolean | Whether scope was reduced (should always be true) |
| `identity.delegation.original_scope` | string[] | Original (broader) scope |
| `identity.delegation.result` | string | `"success"`, `"failure"`, `"denied"`, `"expired"` |
| `identity.delegation.reason` | string | Reason for delegation |
| `identity.delegation.task_id` | string | Task this delegation is for |
| `identity.delegation.expires_at` | string | Delegation expiration ISO timestamp |
| `identity.delegation.ttl_seconds` | int | Delegation time to live |
| `identity.delegation.revocable` | boolean | Whether delegation can be revoked |
| `identity.delegation.proof_type` | string | `"dpop"`, `"mtls_binding"`, `"signed_assertion"` |
| `identity.delegation.obo_token_id` | string | On-Behalf-Of token identifier |
| `identity.delegation.act_claim` | string | JWT `act` claim value |

### Delegation Types

| Value | Description |
|-------|-------------|
| `on_behalf_of` | OAuth On-Behalf-Of (OBO) delegation |
| `token_exchange` | OAuth Token Exchange (RFC 8693) |
| `credential_forwarding` | Forwarding credentials to sub-agent |
| `impersonation` | Agent impersonating delegator (with explicit grant) |
| `capability_grant` | Granting specific capabilities to another agent |
| `scoped_proxy` | Scoped proxy access through the delegator |

### Trust-Domain Crossing & Delegation Depth [RFC v0.4 gap closure]

> **RFC field:** Trust-Domain Crossing & Delegation Depth · **Tier:** SHOULD · **§13 Identity, Delegation & Attribution** · **Grounding attacks:** `AOC-04`, `AOC-09`, `AOC-11`, `TA-11` · **ODIS:** `trust_domain` (6.1), `max_depth` (6.3)

`identity.delegation.chain_depth` records how long the chain is. It does not record whether the chain left the domain that issued the authority, and that is the more consequential fact: authority delegated three hops inside one organisation is a different risk from authority delegated once to a third-party agent. A delegation chain that crosses a trust boundary has handed a credential to a principal governed by someone else's policy, and nothing in the existing attributes says so.

These attributes are attached to `identity.delegation` and to any span representing a call to a counterparty.

#### Boundary Attributes

| Attribute | Type | Requirement | Notes |
|-----------|------|-------------|-------|
| `identity.boundary.crossed` | boolean | **Conditionally Required** | Whether this hop crossed a trust-domain boundary. Emit `false` explicitly — silence is not evidence of containment |
| `identity.boundary.source_domain` | string | **Conditionally Required** | Trust domain the authority came from (ODIS `trust_domain`) |
| `identity.boundary.target_domain` | string | **Conditionally Required** | Trust domain it was handed to |
| `identity.boundary.crossing_type` | string | **Recommended** | `"internal"`, `"cross_team"`, `"cross_tenant"`, `"cross_org"`, `"external_vendor"`, `"public"` |
| `identity.boundary.decision` | string | **Conditionally Required** | `"allow"`, `"deny"`, `"allow_with_attenuation"`, `"not_evaluated"`. `"not_evaluated"` means the crossing was unpoliced |
| `identity.boundary.decision_reason` | string | **Recommended** | Why the crossing was permitted or refused |
| `identity.boundary.policy_ref` | string | **Recommended** | Policy that governed the crossing (ODIS `policy_profile_ref`) |
| `identity.boundary.crossings_count` | int | **Recommended** | Cumulative boundary crossings in this delegation chain |
| `identity.boundary.domains_traversed` | string[] | **Optional** | Ordered list of domains the authority has passed through |

#### Depth Attributes

| Attribute | Type | Requirement | Notes |
|-----------|------|-------------|-------|
| `identity.delegation.depth` | int | **Conditionally Required** | Hops from the originating principal at this point in the chain |
| `identity.delegation.max_depth` | int | **Recommended** | Configured maximum depth (ODIS `max_depth`). Without it, `depth` has no threshold to be judged against |
| `identity.delegation.depth_exceeded` | boolean | **Conditionally Required** | Whether `depth` exceeded `max_depth`. A chain that continued past its limit is an unbounded grant |
| `identity.delegation.root_principal` | string | **Recommended** | The human or system principal the chain originates from (ODIS `originating_principal`) |
| `identity.delegation.root_principal_type` | string | **Recommended** | `"human"`, `"service"`, `"scheduled"`, `"external"`, `"unknown"` |
| `identity.delegation.root_authenticated_at` | string | **Optional** | RFC 3339 timestamp the root principal last authenticated. Authority far from a fresh human authentication is authority worth questioning |

**Propagation constraint (RFC Appendix D.5).** These attributes MUST NOT be carried in W3C `baggage`. Baggage crosses the very trust boundary these fields describe, and a downstream service that reads its own depth or trust domain from an attacker-influenced header is reading a claim, not a fact. Depth and domain are established by the receiving side from its own configuration and from the presented credential.

---

## Span: `identity.credential.mint` [RFC v0.4 gap closure]

> **RFC field:** Credential Minting & Scope-Narrowing Check · **Tier:** SHOULD · **§13 Identity, Delegation & Attribution** · **Grounding attacks:** `AOC-02`, `AOC-08`, `IR-04`, `TA-08` · **ODIS:** `granted_authorizations`, `binding_profile` (6.2/6.3)

`identity.lifecycle.operation` records that a credential was created. `identity.delegation.scope_attenuated` records a boolean claim that scope was reduced. Neither records what the new credential can do *relative to the one it was derived from*, and attenuation is the property that matters: a token exchange returning broader scope, a longer lifetime, or a wider audience than its parent is an escalation. It is invisible unless parent and child are recorded together on the same event.

Represents a credential-minting operation: token exchange, client assertion, client credentials, or role assumption.

### Span Name

Format: `identity.credential.mint {identity.credential.mint.operation}`

### Span Kind

`CLIENT`

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.credential.mint.operation` | string | `"issue"`, `"exchange"`, `"refresh"`, `"impersonate"`, `"assume_role"` |
| `identity.credential.mint.child_id` | string | Identifier of the newly minted credential |
| `identity.credential.mint.issuer` | string | Authority that minted it |
| `identity.credential.mint.subject` | string | Subject the credential was minted for |
| `identity.credential.scope.parent` | string[] | Scopes held by the parent credential |
| `identity.credential.scope.child` | string[] | Scopes granted to the new credential |
| `identity.credential.scope.narrowed` | boolean | Whether child scope is a strict subset of parent scope |
| `identity.credential.scope.escalation` | boolean | Whether the mint produced broader authority by scope, audience, resource, or lifetime |
| `identity.credential.mint.result` | string | `"success"`, `"denied"`, `"failed"` |

`identity.credential.mint.parent_id` is Required where a parent exists; its absence is meaningful, since it means the credential was minted from nothing and is a root grant.

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.credential.mint.grant_type` | string | `"token_exchange"` (RFC 8693), `"client_assertion"`, `"client_credentials"`, `"authorization_code"`, `"refresh_token"`, `"device_code"` |
| `identity.credential.mint.subject_class` | string | Whose identity the token represents: `"end_user"`, `"client_application"`, `"calling_workload"`, `"enforcement_point"` |
| `identity.credential.mint.audience` | string[] | Intended audience(s) of the new credential |
| `identity.credential.scope.requested` | string[] | Scopes asked for, giving the requested-versus-granted delta |
| `identity.credential.scope.added` | string[] | Scopes in the child but not the parent. Non-empty is a privilege escalation |
| `identity.credential.scope.verified` | boolean | Whether `scope.child` was read back from the issued credential after minting rather than assumed from the request |
| `identity.credential.scope.forwarded_unchanged` | boolean | Whether the inbound credential was passed downstream without attenuation — the confused-deputy precondition |
| `identity.credential.mint.parent_ttl_seconds` | int | Remaining lifetime of the parent |
| `identity.credential.mint.child_ttl_seconds` | int | Lifetime granted to the child; exceeding the parent's remaining lifetime is a lifetime escalation |
| `identity.credential.mint.resource_indicators` | string[] | RFC 8707 resource indicators constraining the credential |
| `identity.credential.mint.denial_reason` | string | Why minting was refused |
| `identity.credential.scope.removed` | string[] | Scopes dropped during minting — evidence of intentional attenuation |
| `identity.credential.mint.constraints` | string | JSON of further constraints (time windows, source IP, rate) |
| `identity.credential.mint.on_behalf_of` | string | Principal the credential acts for, where different from the subject |

**Note on `subject_class`.** A hop that mints as `"enforcement_point"` where its parent was `"end_user"` has dropped the human from the accountability chain. The operation may be entirely legitimate, but the resulting credential can no longer be attributed to a person, and downstream telemetry will not show that the substitution happened unless this field records it.

---

## Span: `identity.trust`

Represents trust establishment between agents — agent-to-agent authentication, verifiable credential exchange, and trust boundary crossings.

### Span Name

Format: `identity.trust.{identity.trust.operation} {identity.trust.peer_agent}`

### Span Kind

`CLIENT`

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.agent_id` | string | This agent's identity ID |
| `identity.agent_name` | string | This agent's name |
| `identity.trust.operation` | string | Trust operation (see below) |
| `identity.trust.peer_agent` | string | Peer agent name |
| `identity.trust.peer_agent_id` | string | Peer agent identity ID |

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.trust.result` | string | `"established"`, `"failed"`, `"rejected"`, `"revoked"` |
| `identity.trust.method` | string | `"mtls"`, `"spiffe"`, `"did_vc"`, `"http_signature"`, `"pki"` |
| `identity.trust.trust_domain` | string | Trust domain (e.g., SPIFFE trust domain) |
| `identity.trust.peer_trust_domain` | string | Peer's trust domain |
| `identity.trust.cross_domain` | boolean | Whether this is a cross-domain trust operation |
| `identity.trust.vc_type` | string | Verifiable Credential type presented |
| `identity.trust.vc_issuer` | string | VC issuer |
| `identity.trust.vc_verified` | boolean | Whether VC was successfully verified |
| `identity.trust.trust_level` | string | `"none"`, `"basic"`, `"verified"`, `"high"`, `"full"` |
| `identity.trust.protocol` | string | Trust protocol (`"mcp"`, `"a2a"`, `"custom"`) |
| `identity.trust.federation_id` | string | Federation identifier (for federated trust) |

### Trust Operations

| Value | Description |
|-------|-------------|
| `establish` | Establish trust with a peer agent |
| `verify` | Verify a peer agent's identity |
| `present_credential` | Present a verifiable credential to a peer |
| `verify_credential` | Verify a peer's verifiable credential |
| `federation_join` | Join a trust federation |
| `federation_leave` | Leave a trust federation |
| `boundary_cross` | Cross a trust boundary between domains |
| `revoke_trust` | Revoke trust with a peer |

---

## Span: `identity.session`

Represents an identity session — binding an authenticated identity to a set of operations.

### Span Name

Format: `identity.session {identity.agent_name}`

### Span Kind

`INTERNAL`

### Required Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.agent_id` | string | Agent identity ID |
| `identity.agent_name` | string | Agent name |
| `identity.session.id` | string | Identity session ID |
| `identity.session.operation` | string | Session operation (see below) |

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.session.auth_method` | string | Authentication method used |
| `identity.session.scope` | string[] | Active session scopes |
| `identity.session.expires_at` | string | Session expiration ISO timestamp |
| `identity.session.created_at` | string | Session creation ISO timestamp |
| `identity.session.last_activity` | string | Last activity ISO timestamp |
| `identity.session.actions_count` | int | Number of actions in this session |
| `identity.session.delegations_count` | int | Number of delegations from this session |
| `identity.session.ip_address` | string | Source IP (if network-based) |
| `identity.session.user_agent` | string | Client user agent |
| `identity.session.termination_reason` | string | `"completed"`, `"timeout"`, `"revoked"`, `"error"`, `"manual"` |

### Session Operations

| Value | Description |
|-------|-------------|
| `create` | Create new identity session |
| `refresh` | Refresh session credentials |
| `validate` | Validate an existing session |
| `terminate` | Terminate a session |
| `timeout` | Session timed out |
| `hijack_detected` | Potential session hijacking detected |

---

## Human Approval / Elicitation [RFC v0.4 gap closure]

> **RFC field:** Human Approval / Elicitation Event · **Tier:** SHOULD · **§16 Policy Enforcement & Mediation** · **Grounding attacks:** `AOC-01`, `AOC-02`, `AOC-07`, `AOC-11` · **ODIS:** `originating_principal` (6.3, partial)

Existing conventions record that approval was *required* and that it was *given* — `mcp.tool.approval_required` and `mcp.tool.approved` do exactly this. Neither records who approved, what text they were shown, how long they took to decide, or whether the thing they approved is the thing that then executed. Those four questions are what an approval-fatigue, consent-spoofing, or bypassed-control investigation actually turns on. An approval returned in 200 milliseconds, against a prompt whose text does not describe the action performed, is the signature of a control that has stopped functioning as human oversight, and no attribute in the pre-v0.4 registry makes that visible.

Approval is a lifecycle rather than a point-in-time property, so it is modelled as a **pair of events** joined by `identity.approval.id`.

### Emission Rule (RFC Appendix D.2)

The request and the decision are separated in time and frequently do not share a span — a CLI approval may block for minutes while the agent's span has already ended, and an out-of-band channel such as email or chat may return a decision on an entirely different trace. Implementations therefore:

- MUST emit `identity.approval.requested` when the human is asked, carrying `identity.approval.id`, `requested_at`, `operation`, `prompt_hash`, `trigger`, and `channel`.
- MUST emit `identity.approval.decided` when the outcome is known, carrying the same `identity.approval.id` plus `decision`, `approver`, `approver_type`, `decided_at`, and `latency_ms`.
- SHOULD mirror `identity.approval.decision`, `identity.approval.approver_type`, and `identity.approval.latency_ms` onto the acting span (the `mcp.tool.invoke`, `agent.execute`, or `policy.enforce` span that the approval gates), so that an enforcement decision can be evaluated without performing a join.
- MUST NOT treat the absence of a `decided` event as a denial. Use `identity.approval.status` of `"pending"` or `"expired"`; an approval that lapsed unanswered leaves the requested authority *unresolved*, which is a different security state from *refused*.

### Required Attributes

Required on both events of the pair; `identity.approval.id` is what makes the pair a pair.

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.approval.id` | string | Correlation identifier joining the request to its decision |
| `identity.approval.required` | boolean | Whether approval was required for this operation |
| `identity.approval.status` | string | `"pending"`, `"resolved"`, `"expired"` |
| `identity.approval.operation` | string | The operation being approved, in the same form shown to the approver |
| `identity.approval.prompt_hash` | string | SHA-256 of the exact text shown to the approver |
| `identity.approval.scope_binding.result` | string | `"match"`, `"mismatch"`, `"not_checked"` |
| `identity.approval.decision` | string | `"approved"`, `"denied"`, `"timeout"`, `"auto_approved"`, `"bypassed"`, `"cancelled"` |
| `identity.approval.approver` | string | Resolved identity of the approving principal (ODIS `owner_ref`) — not a session token |
| `identity.approval.approver_type` | string | `"human"`, `"policy"`, `"automation"`, `"none"` |
| `identity.approval.requested_at` | string | RFC 3339 timestamp the approval was requested |
| `identity.approval.decided_at` | string | RFC 3339 timestamp the decision was recorded |
| `identity.approval.latency_ms` | int | Time between request and decision |

`decision`, `approver`, `approver_type`, `decided_at`, and `latency_ms` are Required on `identity.approval.decided` and absent on `identity.approval.requested`.

### Recommended Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.approval.trigger` | string | `"policy"`, `"risk_score"`, `"scope_escalation"`, `"first_use"`, `"destructive_action"`, `"elicitation"` |
| `identity.approval.operation_hash` | string | SHA-256 of the canonicalised operation as executed |
| `identity.approval.scope_binding.divergence` | string[] | Argument names differing between what was approved and what ran |
| `identity.approval.approver_verified` | boolean | Whether the approver was authenticated at approval time, not inherited from a session |
| `identity.approval.timeout_action` | string | `"deny"`, `"allow"`, `"escalate"` |
| `identity.approval.channel` | string | `"cli"`, `"web_ui"`, `"chat"`, `"email"`, `"mcp_elicitation"`, `"api"` |
| `identity.approval.scope` | string | `"single_use"`, `"session"`, `"always"`, `"time_bounded"` |
| `identity.approval.bypass_reason` | string | Present when `decision` is `"bypassed"` |

### Optional Attributes

| Attribute | Type | Notes |
|-----------|------|-------|
| `identity.approval.auth_method` | string | `"session"`, `"reauth"`, `"mfa"`, `"webauthn"`, `"none"` |
| `identity.approval.timeout_ms` | int | Configured approval timeout |
| `identity.approval.elicitation.schema_hash` | string | For MCP `elicitation`, digest of the requested-input schema |
| `identity.approval.elicitation.fields` | string[] | Field *names* solicited from the human — never values |
| `identity.approval.remembered` | boolean | Whether the decision was cached and reused without re-prompting |
| `identity.approval.prior_denials` | int | Prior denials for the same operation in this session |

### Reading the Fields

Four combinations are worth naming, because each is a control failure that the pre-v0.4 attributes would have reported as a clean success:

- `approver_type: "human"` with `latency_ms` consistently under a second across many approvals is **approval fatigue** (`AOC-01`). The control is running; nobody is reading it.
- `scope_binding.result: "not_checked"` means the approval was never bound to the action it authorised. The human approved a description, and something else executed (`AOC-07`).
- `scope_binding.result: "mismatch"` with a populated `divergence` list is the confirmed case of the above — the arguments changed between consent and execution (`AOC-07`, `AOC-11`).
- `decision: "bypassed"`, or `timeout_action: "allow"`, is a **fail-open** oversight control (`AOC-02`). Under the assume-breach provenance rule of RFC §16, `approver_verified` absent must be read as *unverified*: the approver identity is self-asserted by the agent runtime, not attested by the identity provider.

`identity.approval.scope: "always"` deserves separate attention — it converts a per-action control into a standing grant, so a single moment of consent authorises an unbounded number of later actions. Fleet-level counts of `"always"` grants are a better measure of residual oversight than counts of approvals given.

MCP elicitation flows map onto this section directly: set `channel` to `"mcp_elicitation"` and `trigger` to `"elicitation"`, and record the solicited field names in `elicitation.fields`. See [`mcp-spans.md`](mcp-spans.md#human-approval) for the invocation-side attributes this pairs with.

---

## Example: Multi-Agent Delegation Chain

```
Span: identity.lifecycle.create agent-orchestrator
  identity.type: "persistent"
  identity.credential_type: "spiffe_svid"
  identity.provider: "spiffe"
  identity.owner: "platform-team"
  │
  └─ Span: identity.auth agent-orchestrator
       identity.auth.method: "spiffe_svid"
       identity.auth.result: "success"
       identity.auth.scope_granted: ["tools:*", "agents:delegate", "data:read"]
       │
       ├─ Span: identity.authz agent-orchestrator -> customer-db
       │    identity.authz.decision: "allow"
       │    identity.authz.action: "read"
       │    identity.authz.policy_engine: "opa"
       │
       └─ Span: identity.delegate agent-orchestrator -> agent-researcher
            identity.delegation.type: "on_behalf_of"
            identity.delegation.scope_delegated: ["data:read"]
            identity.delegation.scope_attenuated: true
            identity.delegation.chain: ["user-alice", "agent-orchestrator", "agent-researcher"]
            identity.delegation.chain_depth: 2
            │
            └─ Span: identity.trust.establish agent-writer
                 identity.trust.method: "mtls"
                 identity.trust.result: "established"
                 identity.trust.cross_domain: false
```

## Example: Ephemeral Task Identity

```
Span: identity.lifecycle.create task-agent-7f3a
  identity.type: "ephemeral"
  identity.credential_type: "jwt"
  identity.ttl_seconds: 300
  identity.owner: "agent-orchestrator"
  identity.scope: ["tools:web_search", "data:read:public"]
  │
  ├─ Span: identity.session task-agent-7f3a
  │    identity.session.operation: "create"
  │    identity.session.scope: ["tools:web_search", "data:read:public"]
  │    identity.session.expires_at: "2026-02-16T10:05:00Z"
  │
  ├─ Span: identity.authz task-agent-7f3a -> web-search-api
  │    identity.authz.decision: "allow"
  │    identity.authz.jea: true
  │    identity.authz.time_limited: true
  │
  └─ Span: identity.lifecycle.expire task-agent-7f3a
       identity.status: "expired"
       identity.previous_status: "active"
```
