# Upstream PR plan: OCSF

How AITF v0.4 gets proposed to OCSF — what is already drafted and ready to post, what the v0.4 gap
closure newly makes fileable, and the order the two sets should go in.

Companion to the [upstream status report](upstream-status.md), which establishes *what* is missing, and
to [`COSAI-OCSF-discussions.md`](COSAI-OCSF-discussions.md), which establishes *where the two efforts
agree and diverge*. This document covers *how it gets filed*.

> **Source note, revised 2026-09-05.** Coverage claims come from **RFC v0.4 Appendix E**. The live
> `ocsf/ocsf-schema` repository **has now been re-read**, and the answer changed the plan materially:
> **OCSF v1.9.0 shipped on 2026-08-03** and the asks landed across four merged PRs, not one:
> `ai_agent` in [#1641](https://github.com/ocsf/ocsf-schema/pull/1641), `delegation` in
> [#1665](https://github.com/ocsf/ocsf-schema/pull/1665), `attestation` / `prev_event` /
> `record_integrity` in [#1661](https://github.com/ocsf/ocsf-schema/pull/1661), and
> `prompt_text` / `response_text` on `message_context` in
> [#1674](https://github.com/ocsf/ocsf-schema/pull/1674). The `ai_operation` profile was
> **extended**, not introduced, by this release: v1.9.0 widened it to the `system`, `network`,
> `application`, `iam` and `email_activity` classes and added the `delegation` attribute.
> The **`ai` category (`uid 9`) was never ratified** — `categories.json` still stops at `uid 8` —
> though see §1 on why that question is not yet formally closed. Sections 1, 2, 4 and 7 below are
> rewritten accordingly; asks that have landed are struck from the portfolio rather than re-filed.

---

## 1. The category question is answered in practice, but not formally closed

Every prior planning document in this repository treats the AI category UID as the thing that must be
settled before anything else can move. [`COSAI-OCSF-discussions.md`](COSAI-OCSF-discussions.md) names it
**divergence D1** and calls it a blocker; [§2.3](upstream-status.md#23-the-incremental-path-and-where-aitf-sits-on-it)
listed Phase 2 as "blocked on the category-UID reconciliation."

**OCSF v1.9.0 answered it in practice, and the answer was "no new category."** The release landed the
objects and extended the profile — `ai_agent`, `delegation`, `ai_operation` widened to the `system`,
`network`, `application`, `iam` and `email_activity` classes — and left `categories.json` stopping at
`uid 8`. There is no `ai` category, and attaching the profile to existing base classes is the
deliberate alternative to creating one.

> **But it is not formally closed, and the plan should not pretend otherwise.**
> [Issue #1640](https://github.com/ocsf/ocsf-schema/issues/1640) — the issue that proposed all of
> this — was **still open** when re-read on 2026-09-05. Its object and profile asks are satisfied by
> the merged PRs, but one item is explicitly unresolved: the placement of an **agent-layer summary
> event class**, recorded as *TBD* between `ai`, `findings`, and elsewhere. So the accurate statement
> is that OCSF has *declined to create an `ai` category so far*, not that it has ruled one out. AITF's
> position is unaffected either way — it emits onto released classes and holds no reserved UID — but
> an upstream filing that asserts the question is closed would be overstating the record, and would
> be contradicted by the open issue it cites.

AITF has followed that all the way down. Its own Category 7 (7001–7010) was dropped earlier; the
provisional `ai` category classes `9001`/`9002`/`9003` are now **retired** too, and agent, delegation
and agent-to-agent lifecycle ride released classes:

| Where | What it shows |
|:---|:---|
| [`spec/ocsf-mapping/event-classes.md`](spec/ocsf-mapping/event-classes.md) | Every AITF event maps onto a released OCSF class — API Activity `6003`, Datastore Activity `6005`, Application Lifecycle `6002`, Findings `2002`/`2003`/`2004`, Authentication `3002`, Authorize Session `3003`, Inventory Info `5001` — carrying AI context via the `ai_operation` profile |
| [`spec/ocsf-mapping/ocsf-agentic-crosswalk.md`](spec/ocsf-mapping/ocsf-agentic-crosswalk.md) | The full field-level mapping, marking each field released-in-v1.9.0 or *AITF extension* |
| [`sdk/python/src/aitf/ocsf/schema.py`](sdk/python/src/aitf/ocsf/schema.py) | `OCSF_AI_CATEGORY_UID = None` so callers fail loudly; `LEGACY_AI_CLASS_UIDS` keeps `9001`/`9002`/`9003` telemetry decodable; the same treatment is mirrored in the TypeScript, Go and Rust SDKs |

So AITF did not merely concede the number — it conceded the *idea*, adopting OCSF's reuse-first
modeling philosophy (divergence **D2**) outright. Nothing in the portfolio below asks for a category or
a class. **AITF now files as a downstream adopter of released OCSF**, which is a materially stronger
position to contribute from than the one the discussions document anticipated.

---

## 2. What is already drafted

Unlike the [OpenTelemetry side](upstream-pr-plan-opentelemetry.md), OCSF contribution material already
exists in this repository and is substantially complete. This plan **does not restate it**. Two of these
are ready to post as-is:

| Document | Kind | Status |
|:---|:---|:---|
| [`ocsf-contribution-proposal.md`](spec/ocsf-mapping/ocsf-contribution-proposal.md) | Rationale and scope for the lead PR, plus a four-item roadmap | Background — **needs rebasing onto v1.9.0** |
| [`ocsf-issue-draft.md`](spec/ocsf-mapping/ocsf-issue-draft.md) | Discussion issue proposing the `ai_operation` profile extension | **Rebase then post** — the profile now exists; the issue must ask to *extend* it, not create it |
| [`ocsf-pr-draft.md`](spec/ocsf-mapping/ocsf-pr-draft.md) | Full schema PR: `ai_token_usage`, `ai_cost`, `ai_request_parameters`, `ai_latency`, `ai_tool`, the `ai_operation` profile update, the `ai_finding` profile, dictionary entries | **Rebase then post** — base has shifted; the profile-creation half is now redundant |
| [`ocsf-agent-message-pr-draft.md`](spec/ocsf-mapping/ocsf-agent-message-pr-draft.md) | Standalone PR: one generic `agent_message` object plus an `agent_communication` class, with a `protocol_id` discriminator rather than an object per protocol | **Rescope** — the `agent_communication` *class* half is dead (no `ai` category); the `agent_message` *object* half is still wanted, hung off `ai_operation` alongside `message_context` |
| [`ocsf-agentic-crosswalk.md`](spec/ocsf-mapping/ocsf-agentic-crosswalk.md) | AITF → OCSF field-level alignment, already rebased onto v1.9.0 | Background — **current** |

The lead PR draft also carries a **nine-item deferred list** — RAG retrieval, model provenance, MLOps
lifecycle, agentic identity enrichment, MITRE ATLAS, EU AI Act risk levels, agent reasoning and
multi-agent, quality metrics, and agent-to-agent messaging. Those are real and remain the roadmap. The
job of this plan is to add what that list does not cover.

---

## 3. What v0.4 newly makes fileable

The 27 closed gaps added 390 attributes to AITF, and one cluster among them has **no home in any
existing draft**: the authority and enforcement domain. This is not an oversight in the earlier
drafts — those were written before the closure, and this material did not exist yet.

It is also precisely the cluster that RFC Appendix D declares **out of scope for OpenTelemetry by
design**. Four of the 27 gaps are marked "out of scope by design" on the OTel side with a single
correlation-attribute exception. OCSF is not a fallback for them; it is the correct and only home, which
makes this the most important new contribution AITF has to offer and the one with the least risk of
duplicating someone else's work.

| Proposed OCSF object | AITF donor namespace | Gap attrs | Covered by an existing draft? |
|:---|:---|---:|:---|
| `ai_authorization` | [`security.authorization.*`](spec/semantic-conventions/security-cross-cutting.md#authorization-decision-record) | 15 | **No** |
| `ai_taint` | [`security.taint.*`](spec/semantic-conventions/security-cross-cutting.md#session-taint-labels--information-flow-decisions) | 11 | **No** |
| `ai_approval` | [`identity.approval.*`](spec/semantic-conventions/attributes-registry.md#identityapproval-rfc-v04-gap-closure) | 26 | **No** |
| Enforcement availability on `ai_guardrail` | [`security.enforcement.*`](spec/semantic-conventions/security-cross-cutting.md#enforcement-point-availability--failure-mode) | 10 | **No** |
| Mediation coverage | [`security.mediation.*`](spec/semantic-conventions/security-cross-cutting.md) | 8 | **No** |
| `attribute_source` marking | [`security.attribute_source.*`](spec/semantic-conventions/security-cross-cutting.md#attribute-source--trusted-provenance-marking) | 6 | **No** |
| `ai_content` | `gen_ai.content.*` | 13 | **No** |
| `ai_memory` | [`memory.config.*`](spec/semantic-conventions/attributes-registry.md#memoryconfig-rfc-v04-gap-closure) | 21 | **No** |
| `ai_asset` instrumentation coverage | [`asset.instrumentation.*`](spec/semantic-conventions/attributes-registry.md#assetinstrumentation-rfc-v04-gap-closure) | 17 | **No** |
| `ai_bom` / capability change | [`asset.capability.*`](spec/semantic-conventions/attributes-registry.md#assetcapability-rfc-v04-gap-closure) | 16 | **No** |
| Event continuity | [`observability.sequence.*`](spec/semantic-conventions/attributes-registry.md#observabilitysequence-rfc-v04-gap-closure) | 10 | Partly — v1.9.0's `record_integrity` / `attestation` absorbs 5 of 10 |
| Delegation depth and credential minting | [`identity.boundary.*`](spec/semantic-conventions/attributes-registry.md#identityboundary-rfc-v04-gap-closure), [`identity.credential.mint.*`](spec/semantic-conventions/attributes-registry.md#identitycredentialmint-rfc-v04-gap-closure) | 33 | Partly — deferred item 4; v1.9.0's `delegation` object covers lineage but not scope, TTL or proof |
| `ai_retrieval` | [`rag.source.*`](spec/semantic-conventions/attributes-registry.md#ragsource-rfc-v04-gap-closure), `rag.citation.*` | 34 | Yes — deferred item 1 |
| A2A task lifecycle | [`a2a.task.*`](spec/semantic-conventions/attributes-registry.md#a2atasklifecycle-rfc-v04-gap-closure), `a2a.push.*` | 30 | Yes — the `agent_message` draft, rescoped to an object on the released profile |
| Runtime supply chain | `supply_chain.runtime.*` | 27 | Partly — deferred item 2 |

> **What v1.9.0 already absorbed.** Three items that appeared in earlier versions of this table are
> **struck**: the `ai_agent` object, the `ai_operation` profile itself, and the `delegation` object's
> `uid` / `created_time` / `parent_uid` / `issuer_uid`. All four shipped. What survives from the
> delegation ask is the *authority* half — scope, TTL, proof type, and the lineage graph — which is
> genuinely unaddressed upstream and is the stronger argument anyway.

Beyond the objects, two of the six Appendix E asks have no draft at all: **ask 5** (MITRE ATLAS as a
first-class technique reference, which appears only as deferred item 5 with no specification) and **ask
6** (a proper `actor.type_id` for AI agents, so OWASP AOS can stop using `type_id: 99` with a free-text
override). Ask 6 in particular is a *tiny* change with broad benefit and independent corroboration from
outside CoSAI, which makes it unusually easy to land.

---

## 4. The pull request portfolio

Existing drafts keep their identity; new filings are lettered to avoid colliding with the numbering the
drafts already use.

### Already drafted — post as-is

**PR A — `ai_operation` profile substance.** The lead PR in
[`ocsf-pr-draft.md`](spec/ocsf-mapping/ocsf-pr-draft.md). Token usage, cost, latency, request
parameters, `ai_tool`, and the `ai_finding` profile. Additive, no new category, no new required fields.
Post the [issue draft](spec/ocsf-mapping/ocsf-issue-draft.md) first, as that document itself recommends.

**PR B — `agent_message` (object only; the class is withdrawn).** From
[`ocsf-agent-message-pr-draft.md`](spec/ocsf-mapping/ocsf-agent-message-pr-draft.md). One generic object
with a `protocol_id` discriminator covering A2A, ACP, ANP and MCP. **The `agent_communication` class in
that draft is withdrawn** — there is no `ai` category to hold it, and AITF now emits agent-to-agent
communication as API Activity `6003`. What remains fileable is the object, proposed as an addition to
the released `ai_operation` profile, sitting beside the `message_context` attribute v1.9.0 already
shipped. The first question a reviewer will ask is why `message_context` is not enough; the answer —
`protocol_id`, task lifecycle, and push-notification identity — should lead the PR body. AITF's 30
`a2a.*` gap attributes are the donor set and should be offered explicitly as the field list, which the
draft predates.

### New — enabled by v0.4

**PR C — The authority cluster: `ai_authorization`, `ai_taint`, `ai_approval`.**

The most valuable new filing and the one to write first. Three objects with complete attribute sets:
an authorization decision with its deciding authority, rule identity, policy version and obligations; a
taint label set with scope, origin and flow decision; and a human-approval record with scope-binding
result, approver verification and channel.

The argument to lead with is **why OCSF and not OpenTelemetry**: the RFC marks these out of scope for
OTel by design, and the reasoning is that an authorization decision is a security-domain fact that
belongs in the at-rest security schema, not in an observability trace. This is not AITF routing
leftovers to OCSF — it is AITF agreeing with OCSF about what OCSF is for. Say so in the first paragraph.

The `ai_approval` object carries the subtlest field in the whole set and it should be the worked
example: `identity.approval.scope_binding.result` records whether what the human approved is what the
system then executed. An approval dialog that says "allow file access" followed by a write to a
different path is the attack, and no existing OCSF object can represent it.

File `ai_authorization` and `ai_taint` together — taint decisions are authorization decisions with an
information-flow justification — and consider `ai_approval` as a separable commit, since human-in-the-loop
approval may attract a different reviewer.

**PR D — Enforcement availability and mediation coverage.**

Adds enforcement-point availability, failure mode, and mediation-coverage attributes to `ai_guardrail`
and `ai_asset`, together with the `attribute_source` marking.

This is the observability-plane-integrity row of the Appendix E.1 matrix and it answers a question no
current schema can: **did the control actually run?** A guardrail that timed out and failed open
produces telemetry almost identical to a guardrail that ran and allowed — same decision, entirely
different security posture. `security.enforcement.failure_mode` distinguishes them.

Pair it with `attribute_source` because they answer adjacent questions. A consumer that cannot tell a
verified attribute from a self-asserted one cannot safely act on any of the above, and the
[divergence table](upstream-status.md#22-what-cosai-asks-ocsf-to-include) shows AITF and Appendix E.2
currently disagree about what this field even measures — see section 5.

**PR E — `ai_content` and `ai_memory`.**

Content identity — part hashes, attachment identity, modality, extracted-text digest — and memory
configuration and provenance. Both are named in Appendix E.2 ask 3 and neither appears in any existing
draft. Both are privacy-sensitive, so both should be proposed **hash-first**, with raw content optional
and redaction flags mandatory where raw content is present. That framing is also what makes them
adoptable in regulated environments, which is the strongest adoption argument available.

**PR F — Capability change and AgBOM.**

`ai_bom` plus a capability-change activity, from the 16 `asset.capability.*` attributes. An agent that
gains a new tool between one run and the next has changed its blast radius, and there is no event for
that today in either standard. This is MUST-tier in the RFC and genuinely absent everywhere, which is a
rare combination worth stating plainly.

**PR G — MITRE ATLAS as a first-class technique reference.**

Appendix E ask 5. OCSF findings carry ATT&CK techniques; ATLAS `AML.Txxxx` has no first-class
representation, so AITF currently rides it on `compliance.control_id` with `framework=mitre_atlas` — a
workaround that should be retired. Small, self-contained, and useful well beyond AITF.

Note the deliberate overlap with the [OTel plan](upstream-pr-plan-opentelemetry.md): ATLAS is filed as a
separable commit there too, precisely because it may belong here instead. If OTel declines it, this PR
is where it lands, and nothing is lost.

**PR H — AI agent as a first-class actor type.**

Appendix E ask 6, and the cheapest filing in the entire plan. OWASP AOS's OCSF binding currently
represents an agent as `actor.type_id: 99` ("Other") with `type: "AI Agent"` — a documented workaround
by a group with no connection to CoSAI. That independent corroboration is the whole argument. File it
standalone; it should not be entangled with anything.

**PR I — Sequence ordinals and gap markers on `attestation`.**

**Substantially reduced by v1.9.0.** The release added the `record_integrity` profile, the `attestation`
object and the `prev_event` object, which between them absorb most of the
`observability.sequence.*` group: `hash` → `attestation.fingerprint`, `prev_hash` →
`prev_event.fingerprint`, `scope_id` → `chain_uid`, `signature` + `signer_key_id` →
`attestation.signatures`. See
[crosswalk §7](spec/ocsf-mapping/ocsf-agentic-crosswalk.md#7-event-sequence-continuity--the-record_integrity-profile).

What remains is five attributes and one argument: `number`, `scope_type`, `gap_detected`, `gap_size`
and `reordered`. The argument is that **a hash chain proves a log is unaltered; only an ordinal proves
it is complete**. An adversary who drops an event and re-links the chain leaves a valid `attestation`
chain behind. EU AI Act Art. 12 requires demonstrating completeness, not just integrity, so this is a
compliance argument rather than a preference. That makes it a much stronger filing than the original
ten-attribute version, and it should be *raised in the `record_integrity` discussion* rather than filed
as a standalone object.

---

## 5. Enum negotiation positions

Appendix E ask 4 proposes eleven enums. Eight now have concrete AITF value sets in all four SDKs, and
**none of the eight matches the Appendix E.2 wording exactly**. Those differences are the substance of
the negotiation and should be surfaced in the PR body rather than smoothed over — a reviewer who
discovers an undisclosed divergence trusts the rest of the filing less.

The full comparison is in [§2.2](upstream-status.md#22-what-cosai-asks-ocsf-to-include). The positions
to take:

**Concede immediately.** The `mcp_primitive` spelling (AITF singular `root`, E.2 `roots`) — defer to
whatever the MCP specification uses. The `enforcement_decision` value `modify`, which E.2 proposes and
AITF lacks; adding it costs nothing and it describes a real guardrail behaviour.

**Offer as a superset and let the working group trim.** `taint_scope` (AITF's four values against E.2's
two), `trigger_type` (eight against two), `enforcement_failure_mode` (six against two), `mcp_primitive`
(eight against six). In each case AITF's extra values came from specific attacks in RFC Appendix A, so
the defence of any individual value is evidentiary rather than aesthetic. Bring the attack references.

**Argue, because the field is genuinely different.** `attribute_source` is not a naming disagreement:
E.2 names the *producing system* (idp / pdp / enforcement-state / platform / self-asserted) while AITF
names the *assurance level* (self_asserted / verified / derived / unknown). Both are useful and they are
not the same field. Propose **both**, under distinct names.

Likewise `approval_status`. E.2 folds `bypassed` into the lifecycle enum; AITF keeps lifecycle state
(pending / resolved / expired) separate from outcome and carries `bypassed` on
`identity.approval.decision`. AITF's split is defensible — a bypassed approval is an outcome, not a
lifecycle stage — but this is a preference, not evidence, so concede if pressed.

**Bring nothing yet.** `trust_level`, `autonomy_level` (L1–L5) and `tool_trust_boundary` have **no AITF
value set at all**. For `trust_level` this is being fixed as a
[prerequisite to the OTel filing](upstream-pr-plan-opentelemetry.md#3-the-prerequisite-aitf-has-no-trust_level-value-set),
and the same definition should serve both. The other two need to be defined or dropped from the ask
before ask 4 is filed as a whole. Filing an enum proposal with three empty rows invites the question of
whether the other eight are equally speculative.

---

## 6. Sequencing

One ordering constraint is real, and one has disappeared. PR B no longer depends on an `ai` category,
because it no longer asks for a class. What remains is that PRs C through I all read more naturally once
the released `ai_operation` profile has been *extended* once successfully, because PR A establishes the
pattern — additive attributes on an existing profile, all optional, no new required fields — that the
rest reuse.

| Wave | Filing | Depends on |
|:---:|:---|:---|
| **0** | Rebase the [issue draft](spec/ocsf-mapping/ocsf-issue-draft.md) and [PR draft](spec/ocsf-mapping/ocsf-pr-draft.md) onto released v1.9.0; post the issue | — |
| **1** | **PR A** (profile substance), **PR H** (actor type) | Wave 0 rebase. PR H is independent and trivially small — file it in parallel |
| **2** | **PR C** (authority cluster), **PR G** (ATLAS) | PR A establishing the additive pattern |
| **3** | **PR D** (enforcement availability), **PR E** (content and memory) | Wave 2 |
| **4** | **PR B** (`agent_message` object), **PR F** (capability change / AgBOM) | Both independent; PR B benefits from PR A landing first, since it extends the same profile |
| **5** | **PR I** (sequence ordinals on `attestation`) | Raise in the `record_integrity` discussion rather than filing standalone |

Wave 0 is now a rebase, not an investigation. Both drafts were written against a world in which
`ai_operation` did not yet exist; they must be re-read line by line against the released profile, and
every attribute that shipped must be **deleted from the ask**. A filing that re-proposes fields already
in the schema is the fastest way to lose a reviewer's attention. **Do that before writing any JSON.**

Two process steps from
[`COSAI-OCSF-discussions.md` §4.7](COSAI-OCSF-discussions.md) remain outstanding and are worth more than
any individual PR: a **joint working session** with the #1640 authors, and **publishing the mapping
table in both repositories** so consumers have one source of truth. Section 1 above means the working
session no longer has to negotiate the category UID, so it can start from the hybrid rule and the
authority cluster instead — a much better agenda.

---

## 7. Pre-filing verification gate

Items 1, 2 and 4 were **answered on 2026-09-05** and are recorded here rather than left open.

1. ~~**Re-read #1640 and #1641.**~~ **Done, and the answer was wider than one PR.** v1.9.0
   (2026-08-03) landed `objects/ai_agent.json` (#1641), `objects/delegation.json` (#1665),
   `objects/attestation.json` + `objects/prev_event.json` + `profiles/record_integrity.json` (#1661),
   and `prompt_text` / `response_text` on `message_context` (#1674); it **extended**
   `profiles/ai_operation.json` rather than introducing it, and added `ai_agent` and
   `hosted_ai_agent_list` to `process`. #1640's object/profile portion is satisfied, but **the issue
   is still open** — its agent-layer class placement is TBD, and its delegation-scope portion is
   untouched. Anything filed against #1640 should say *"building on what landed"*, not *"closing"*.
2. ~~**Confirm the `ai` category UID.**~~ **Done — there is none.** `categories.json` stops at `uid 8`
   (*Unmanned Systems*). AITF has retired `9001`/`9002`/`9003` accordingly.
3. **Check for competing proposals.** Still open, and now more important, not less: v1.9.0 shows the
   AI-schema space is moving fast. An object someone else has already proposed is a PR to join, not to
   duplicate. Re-check immediately before each filing, not once for the whole portfolio.
4. ~~**Verify the OCSF version.**~~ **Done.** AITF now targets **v1.9.0** throughout — spec, JSON
   schema, and all four SDKs. Three asks were absorbed by the release and are struck from the portfolio:
   the `ai_agent` object, the `delegation` object's four released attributes, and the creation of the
   `ai_operation` profile.
5. **Read the contribution guide.** Confirm the JSON schema conventions, the `dictionary.json`
   requirements, and the validation commands the drafts' checklists assume.
6. **Confirm the OWASP AOS `type_id: 99` workaround still exists.** It is PR H's entire argument. If AOS
   has moved on, that PR needs a new one. Note that released `ai_agent.type_id` does **not** include
   `MCP` or `A2A` members despite their discussion on #1641 — adding those two is a small, well-evidenced
   companion ask to PR H.

## 8. Acceptance criteria

The same standard as the [OTel plan](upstream-pr-plan-opentelemetry.md#7-acceptance-criteria): a filing
succeeds when the object or enum is in the upstream schema and AITF can **delete** its own attributes.
Every AITF-defined attribute is a workaround, and the [delta table](upstream-status.md) is the list of
workarounds still outstanding.

The OCSF-specific version of partial success is worth naming, because it will be common. An object that
lands with fewer fields than proposed is a success — the remainder becomes a follow-up. An object that
lands under a different name is a success requiring a mapping note in the
[crosswalk](spec/ocsf-mapping/ocsf-agentic-crosswalk.md). An object redirected from a new class onto an
existing class with a profile is a success and specifically *the outcome the reuse-first rule predicts*,
which is why AITF adopted that rule in the first place.

The one outcome to guard against is not rejection. It is **AITF's attributes landing upstream while
AITF's own definitions drift away from them**. Every merged PR must be followed by an update to the
crosswalk and the delta table in the same week, or the workaround inventory stops being trustworthy and
the next contributor cannot tell what is still outstanding.

---

## Related documents

- [Upstream status: OpenTelemetry & OCSF](upstream-status.md) — what is missing, and why
- [Upstream PR plan: OpenTelemetry](upstream-pr-plan-opentelemetry.md) — the companion filing plan
- [OCSF ↔ CoSAI AITF alignment discussion](COSAI-OCSF-discussions.md) — divergences D1–D8 and the hybrid rule
- [AITF ↔ OCSF agentic crosswalk](spec/ocsf-mapping/ocsf-agentic-crosswalk.md) — field-level mapping
- [OCSF contribution proposal](spec/ocsf-mapping/ocsf-contribution-proposal.md), [issue draft](spec/ocsf-mapping/ocsf-issue-draft.md), [PR draft](spec/ocsf-mapping/ocsf-pr-draft.md), [`agent_message` PR draft](spec/ocsf-mapping/ocsf-agent-message-pr-draft.md)
- [RFC v0.4 Appendix C resolution log](AITF_gaps.md) — the 27 closed gaps and their grounding attacks
