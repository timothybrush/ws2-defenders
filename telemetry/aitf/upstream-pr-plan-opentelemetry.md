# Upstream PR plan: OpenTelemetry GenAI semantic conventions

How AITF v0.4 gets proposed to OpenTelemetry — staged as filable pull requests, in the order they
should be filed, with the acceptance criteria each must meet.

Companion to the [upstream status report](upstream-status.md), which establishes *what* is missing.
This document covers *how it gets filed*. The nine asks in
[§1.3](upstream-status.md#13-the-nine-asks-in-priority-order) are the input; seven pull requests are the
output, plus one non-schema guidance filing that does not derive from the nine.

> **Source constraint, carried forward.** Every coverage claim behind this plan comes from **RFC v0.4
> Appendix D**. The live OpenTelemetry registry was **not** re-read while writing it. Appendix D's own
> verification note applies with full force: **re-verify against the live registry before filing
> anything.** Section 6 below is the gate that makes that non-optional.

> **Repository mechanics are asserted, not verified.** The GenAI conventions moved to
> **`open-telemetry/semantic-conventions-genai`**, a repository this plan has not inspected. The
> workflow described in section 2 — YAML model files, generated markdown, a changelog fragment per
> change — is the established `open-telemetry/semantic-conventions` process and is assumed to carry
> over. Confirm it against the target repository's `CONTRIBUTING.md` before writing any YAML. Where
> this plan says "model file", read "whatever that repository's source-of-truth format turns out to
> be."

---

## 1. What "match 0.4" means here, and what it does not

RFC v0.4 asks OpenTelemetry for nine things. It deliberately does **not** ask for two whole clusters —
identity and delegation, and policy enforcement — with a single narrow exception for correlation. That
restraint is the plan's strongest asset in review, and it should be stated in the first paragraph of
every pull request rather than left for a maintainer to discover.

The argument that carries these filings is not that OpenTelemetry has a security gap. It is that every
one of the nine asks is an **extension to a structure OTel already has**: a field on a message part
that already exists, an enum on an evaluation event that already exists, an identifier level between
two identifier levels that already exist. Nothing here proposes a new signal type, a new namespace, or
a new subsystem. An ask that cannot be phrased that way is not on this list — it went to
[OCSF](upstream-pr-plan-ocsf.md) instead.

Two further constraints shape every PR below. Content capture is **hash-first**: AITF's donor
attributes carry digests, not payloads, and the proposals inherit that property, which matters because
content capture is opt-in and privacy-gated in OTel already. And the **`Status: Development`** label on
every GenAI convention is the reason to move now — nothing proposed here would be a breaking change to
a stable convention, and that window will not stay open.

---

## 2. Filing mechanics

OpenTelemetry semantic conventions are not merged from prose. Each change is a YAML model file edit
that regenerates the published markdown, plus a changelog fragment. A pull request that edits the
generated markdown by hand will be asked to start over, so the per-PR file lists below name the model
file as the deliverable and treat the docs as output.

The project also expects an **issue before a pull request** for anything that adds attributes. The
sequencing in section 5 reflects that: one umbrella issue introduces the whole nine-ask set and asks
maintainers to rule on scope, then individual pull requests land against that issue rather than each
arriving cold.

Two review realities are worth planning around. Attribute additions to GenAI conventions are reviewed
by a working group that meets on a fixed cadence, so wall-clock time is dominated by review cycles
rather than by authoring — which is an argument for filing the cheap, high-consensus PRs first to
establish credibility. And every new attribute needs a stated **requirement level** and a **stability
level**; proposing everything as `opt_in` or `recommended` at `development` stability removes the most
common reason a semconv PR stalls.

---

## 3. The prerequisite: AITF has no `trust_level` value set

The highest-value ask on the list is the one AITF cannot currently support with a concrete proposal,
and this has to be fixed before filing rather than papered over.

Ask 1 is `gen_ai.input.trust_level` — the trusted-instruction / trusted-data / untrusted-data /
adversarial-suspected distinction that separates a system prompt from retrieved web content. It is the
single field with no analogue in either upstream standard, and it is simultaneously
[OCSF ask E.2-4](upstream-status.md#22-what-cosai-asks-ocsf-to-include). AITF defines **no value set for
it**. Two attributes carry the *name* — `rag.source.trust_level` and `identity.trust.trust_level` — but
neither is that enum, and this was not one of the 27 Appendix C gaps, so the v0.4 closure did not touch
it.

Filing an enum proposal upstream while the reference implementation has no enum is a bad position to
argue from. **PR 1 is therefore gated on AITF first defining the value set**, as a `gen_ai.input.*`
attribute with an enum class in all four SDKs, matching the pattern the other eighteen enum sets
already follow. That is new AITF work, tracked here because it blocks the most important filing:

| Step | Work | Where |
|:---|:---|:---|
| 3a | Define `gen_ai.input.trust_level` in the registry with the four-value enum and its grounding attacks | [`attributes-registry.md`](spec/semantic-conventions/attributes-registry.md) |
| 3b | Add the enum class to `GenAIAttributes` in Python, then port to Go, TypeScript and Rust | [`sdk/`](sdk/) |
| 3c | Document the per-part application rule and the propagation prohibition (section 4, PR 1) | [`gen-ai-spans.md`](spec/semantic-conventions/gen-ai-spans.md) |
| 3d | Confirm ≥ 2 grounding attacks in RFC Appendix A, per the promotion rule | RFC Appendix A |

Steps 3a–3c are mechanical once 3d confirms the field earns promotion. Nothing else in this plan
depends on them.

---

## 4. The pull request portfolio

Nine asks do not make nine pull requests. Asks that touch the same model file, share a reviewer, and
rest on the same argument should arrive together; asks that are contentious should arrive alone so
they cannot sink their neighbours. That collapses the nine into seven filings.

Each entry below states its scope, the AITF attributes offered as donor material, the acceptance
criteria that would let AITF retire its own namespace for that cluster, and the objection most likely
to be raised.

### PR 1 — Input trust classification

| | |
|:---|:---|
| **Covers** | Ask 1 |
| **Proposes** | `gen_ai.input.trust_level`, an enum on message parts: `trusted_instruction` / `trusted_data` / `untrusted_data` / `adversarial_suspected` |
| **Donor** | New AITF work — see section 3. No existing AITF attribute carries this value set |
| **Blocked by** | Section 3 (AITF must define the enum first) |
| **Requirement level** | `recommended` where the producer can classify; not required |
| **Retires** | Nothing in AITF today; establishes the field upstream from the start |

This is the ask with the largest security payoff and the smallest schema footprint. Prompt injection is
definitionally a confusion between instruction and data, and no field in either upstream standard
records which is which. A single enum on a structure OTel already has — message parts — closes it.

The proposal must carry one **normative prohibition**: a trust level is a property of a specific
content part, not of a conversation, and it must never be propagated across parts or inherited by
downstream spans. Without that sentence, the field decays into a per-session label that reads "trusted"
for the whole session because the first part was a system prompt, which is worse than not having it.

**Likely objection.** That classification is the instrumentation's judgement rather than an observed
fact, and OTel conventions prefer observed facts. The answer is that the producer *always* knows the
provenance of a content part — it assembled the prompt — so this is closer to recording an input than
to inferring a verdict. Pair the enum with the `security.attribute_source` marking from PR 2 so a
consumer can tell a verified classification from a self-asserted one.

### PR 2 — Guardrail and evaluation outcomes

| | |
|:---|:---|
| **Covers** | Ask 2 |
| **Proposes** | Security-verdict semantics extending `gen_ai.evaluation.*`: guardrail identity, verdict, score, blocked flag, and a threat-technique reference accepting MITRE **ATLAS** `AML.Txxxx` alongside ATT&CK |
| **Donor** | [`security.guardrail.*`](spec/semantic-conventions/security-cross-cutting.md), `security.blocked`, `security.threat_type` |
| **Requirement level** | `recommended` on evaluation events; `opt_in` for score |
| **Retires** | AITF's guardrail-verdict attributes, if accepted in full |

OTel already has an evaluation event with a verdict shape. A guardrail is an evaluator whose verdict
has an enforcement consequence, so this extends the existing structure rather than competing with it.
The ATLAS reference is the part most likely to need a separate conversation: OTel does not model attack
techniques at all today, and this may be the ask that belongs in OCSF instead. **File the ATLAS
reference as a separable commit** so it can be dropped without losing the guardrail work.

**Likely objection.** Evaluation is a quality concern, and folding blocking into it conflates
measurement with enforcement. The answer is that the blocked flag records what the pipeline *did*,
which is observable, and that separating the two would require a second event type for something that
is structurally identical.

### PR 3 — Execution structure: turn, step and trigger

| | |
|:---|:---|
| **Covers** | Asks 5 and 6 |
| **Proposes** | `gen_ai.turn.id` / `.index` / `.parent_id`, `gen_ai.step.id` / `.parent_id`, `gen_ai.run.id`; plus `gen_ai.trigger.type` and `gen_ai.trigger.event` |
| **Donor** | [`gen-ai-spans.md`](spec/semantic-conventions/gen-ai-spans.md) — all nine attributes exist in AITF v0.4 verbatim |
| **Requirement level** | `recommended` |
| **Retires** | The whole AITF identifier-hierarchy and trigger cluster |

This is the highest-probability filing and should go first. OTel has `gen_ai.conversation.id` at the top
and span identifiers at the bottom, with nothing in between — no way to say "this is the third turn"
or "this is step 7 of the agent loop". The proposal fills in intermediate levels of a hierarchy OTel
already established, using AITF names that were chosen to slot into it.

The trigger attributes ride along because they answer a question the identifiers raise: an
observability consumer that can see a run can immediately ask what started it, and
`user_initiated` versus `scheduled` versus `agent_initiated` is a distinction OTel has no way to draw.
AITF's eight-value enum is a refinement of the RFC's user-initiated/autonomous split and should be
offered as a superset the working group can trim.

**Likely objection.** Scope. Two asks in one PR invites a request to split. Splitting is cheap here and
should be conceded immediately if asked — the identifiers stand alone fine.

### PR 4 — Memory and retrieval provenance

| | |
|:---|:---|
| **Covers** | Asks 3 and 4 |
| **Proposes** | Provenance and footprint on `gen_ai.memory.*`; retrieved-item source, provenance and integrity on the retrieval conventions |
| **Donor** | [`memory.provenance`](spec/semantic-conventions/attributes-registry.md), [`memory.config.*`](spec/semantic-conventions/attributes-registry.md#memoryconfig-rfc-v04-gap-closure), [`rag.source.*`](spec/semantic-conventions/attributes-registry.md#ragsource-rfc-v04-gap-closure), [`rag.citation.*`](spec/semantic-conventions/rag-spans.md) |
| **Requirement level** | `recommended` for provenance; `opt_in` for configuration digests |
| **Retires** | `memory.provenance` and the `rag.source.*` provenance subset |

Memory poisoning and retrieval poisoning are the same attack against two stores, and both are invisible
without knowing where a retrieved item *came from*. OTel has `retrieval.documents` but no way to record
that a document originated from a tool result rather than from a curated index. `memory.provenance` —
`conversation` / `tool_result` / `imported` — is the smallest field that makes the distinction, and it
is the one AITF enum that already matches Appendix E.2 wording, so it can be offered without
negotiation.

Keep the **configuration** attributes (`memory.config.*`, the search parameters under `rag.source.*`)
as a clearly separable second commit. They are MAY-tier in the RFC and are the most likely thing a
reviewer will call out as belonging in a resource attribute rather than a span.

### PR 5 — Tool definition digest and MCP primitive

| | |
|:---|:---|
| **Covers** | Ask 7 |
| **Proposes** | `gen_ai.tool.definition.hash` on tool definitions; `mcp.primitive` as a discriminator on MCP operations |
| **Donor** | [`mcp.tool.definition.*`](spec/semantic-conventions/mcp-spans.md), [`mcp.primitive`](spec/semantic-conventions/mcp-spans.md) |
| **Requirement level** | `recommended` |
| **Retires** | `mcp.tool.definition.hash` and `mcp.primitive` |

OTel has `tool.definitions` but no digest of them, which means a rug-pull — a server that changes a
tool's description after approval — leaves no trace. A hash is cheap, has no privacy exposure, and is
the canonical way to detect the attack.

`mcp.primitive` is the more interesting half. OTel's `mcp.method.name` partially serves, but it does not
separate the primitives that **invert the direction of control**: `sampling` lets a server drive the
client's model, and `elicitation` lets it drive the client's user. Those two values are the reason the
field is worth having, and the PR should lead with them rather than with the enum's completeness.

Offer AITF's eight values as a superset of the RFC's six and flag the one **known spelling divergence**:
AITF uses singular `root`, Appendix E.2 writes `roots`. Concede whichever the MCP specification uses.

### PR 6 — Attachment identity on content parts

| | |
|:---|:---|
| **Covers** | Ask 8 |
| **Proposes** | Attachment name, hash, size, source and extracted-text hash on content parts |
| **Donor** | `gen_ai.content.attachment.*`, `gen_ai.content.part.hash` |
| **Requirement level** | `opt_in` — this is content-adjacent |
| **Retires** | The `gen_ai.content.attachment.*` cluster |

OTel content parts carry a type but no identity, so a poisoned PDF and a benign one are
indistinguishable in telemetry. The `extracted_text_hash` attribute is the one that matters most and is
the least obvious: it digests what the model actually *read* after extraction, which is where indirect
injection lands, and it differs from the file hash whenever the extractor is the attack surface.

File this one at `opt_in` and lead with the privacy properties — hashes only, no filenames required —
because it is the proposal most likely to trigger a content-capture objection.

### PR 7 — Model provenance

| | |
|:---|:---|
| **Covers** | Ask 9 |
| **Proposes** | Model weights digest, signer identity and signature-verification result alongside the existing model attributes |
| **Donor** | [`supply_chain.*`](spec/semantic-conventions/attributes-registry.md) |
| **Requirement level** | `opt_in` |
| **Retires** | The `supply_chain.*` model-integrity subset |

Lowest priority and lowest consensus of the seven. OTel records model *name* and *provider*, which
answers "what was called" but not "was it the artifact we approved". File it last, after the earlier PRs
have established that these proposals are careful, and be prepared for it to be redirected to OCSF as a
supply-chain concern rather than an observability one — which would be a defensible outcome, not a
failure.

### PR 8 — Operational guidance (non-schema)

| | |
|:---|:---|
| **Covers** | RFC Appendix D.5, the [three operational traps](upstream-status.md#15-three-operational-traps) |
| **Proposes** | Documentation only: context propagation through MCP `params._meta` per SEP-414; a sampling caveat for security-relevant spans; hash-first content capture guidance |
| **Files** | Supplementary markdown, no model file changes |

This is not an attribute proposal and should not be filed as one. It is the observation that the
proposals above are defeated in practice by three deployment realities: trace context does not survive
an MCP hop without explicit propagation, head-based sampling drops exactly the rare spans a security
consumer needs, and naive content capture makes the whole convention set unadoptable in regulated
environments. Filing it as a guidance PR costs a reviewer nothing and makes the other seven more likely
to be adopted correctly.

---

## 5. Sequencing

Nothing here depends on anything else, so the ordering is chosen for **credibility accumulation**
rather than for dependency. File the proposals most likely to be accepted first, so that the
contentious ones arrive from a contributor with merged work behind them.

| Wave | Filing | Rationale |
|:---:|:---|:---|
| **0** | Umbrella issue introducing all nine asks | Project convention; gets a scope ruling before any YAML is written |
| **1** | PR 3 (turn/step/trigger), PR 5 (tool digest + primitive) | Highest consensus, purely additive, no privacy surface, and both fill in structures OTel already established |
| **2** | PR 4 (memory/retrieval provenance), PR 2 (guardrail outcomes) | Higher value, more discussion; PR 2's ATLAS half may split off |
| **3** | PR 1 (trust level) | Gated on section 3, and the ask that most benefits from prior merged work |
| **4** | PR 6 (attachment identity), PR 7 (model provenance) | Content-adjacent and supply-chain-adjacent; most likely to be deferred or redirected |
| **any** | PR 8 (guidance) | Independent; file whenever it is written |

The umbrella issue in wave 0 is the highest-leverage single action in this plan. It is where the
question "does OpenTelemetry want security-relevant attributes in GenAI conventions at all?" gets
answered, and a negative answer would redirect most of this work to
[OCSF](upstream-pr-plan-ocsf.md) before any of it is wasted.

---

## 6. Pre-filing verification gate

The source constraint means **nothing in this plan has been checked against the live registry**. Every
filing must clear this gate first, and a failure here changes the proposal rather than merely delaying
it:

1. **Confirm the repository.** Verify that `open-telemetry/semantic-conventions-genai` is the live home
   and that the main-registry deprecations are relocations, not removals. Appendix D says so; confirm it.
2. **Re-check each proposed attribute against the current registry.** An attribute that landed upstream
   since the RFC was written converts that PR from a proposal into an alignment note — a better outcome,
   but only if caught before filing.
3. **Confirm the stability labels.** The entire "now is the moment" argument rests on
   `Status: Development`. If any target convention has stabilized, that PR needs a migration story.
4. **Read the contribution guide.** Confirm the model-file format, the changelog fragment mechanism, and
   whether an issue is genuinely required before a PR.
5. **Check for existing proposals.** Someone may already have filed the trust-level or turn-identifier
   ask. Joining an open PR beats opening a competing one.

## 7. Acceptance criteria

A filing has succeeded when the attribute is in the upstream registry and AITF can **delete** its own.
That is the point of the whole exercise: every AITF-defined attribute is a workaround, and the
[delta table](upstream-status.md) is the list of workarounds still outstanding.

Partial success is normal and should be recorded rather than treated as failure. An ask that lands with
a different name is a full success requiring only a mapping note. An ask that lands as `opt_in` when
`recommended` was proposed is a success. An ask that is redirected to OCSF is a success for the *ask*
and a re-file for the *plan*. Only an ask that is rejected on the merits — "OpenTelemetry does not want
this field to exist" — is a real failure, and it should be recorded in the delta table with the
maintainer's reasoning, because that reasoning is what a future contributor needs.

---

## Related documents

- [Upstream status: OpenTelemetry & OCSF](upstream-status.md) — what is missing, and why
- [Upstream PR plan: OCSF](upstream-pr-plan-ocsf.md) — the companion filing plan
- [RFC v0.4 Appendix C resolution log](AITF_gaps.md) — the 27 closed gaps and their grounding attacks
- [Attributes registry](spec/semantic-conventions/attributes-registry.md) — the donor attribute definitions
