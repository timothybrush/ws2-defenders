# RAG Span Conventions (RAG_CONTEXT)

Status: **Normative** | CoSAI WS2 Alignment: **RAG_CONTEXT** | OCSF Class: **Datastore Activity (6005)** (`ai_operation` profile)

AITF defines semantic conventions for Retrieval-Augmented Generation (RAG) pipelines, covering query processing, vector retrieval, document scoring, reranking, and context quality evaluation. This specification defines the normative field requirements aligned with CoSAI Working Stream 2 (Telemetry for AI) and mapped to applicable compliance and threat frameworks.

Key words "MUST", "SHOULD", "MAY" follow [RFC 2119](https://www.rfc-editor.org/rfc/rfc2119).

---

## Overview

RAG pipelines combine retrieval from external knowledge bases with generative AI to produce grounded, context-aware responses. AITF provides telemetry across the full RAG lifecycle:

```
RAG Pipeline:
  query -> embed -> retrieve -> [rerank] -> augment -> generate -> [evaluate]

Spans:
  rag.pipeline      (root)
    gen_ai.retrieval.query.text       (query embedding)
    rag.retrieve    (vector search)
    rag.rerank      (optional reranking)
    gen_ai.inference      (generation with context)
    rag.evaluate    (optional quality evaluation)
```

---

## Span: `rag.pipeline`

Represents a complete RAG pipeline execution.

### Span Name

Format: `rag.pipeline {rag.pipeline.name}`

### Span Kind

`INTERNAL`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `rag.pipeline.name` | string | **Required** | Pipeline name/identifier | NIST AI RMF MAP-1.1 |
| `rag.pipeline.stage` | string | **Required** | Current stage: `"retrieve"`, `"rerank"`, `"generate"`, `"evaluate"` | NIST AI RMF MEASURE-2.5 |
| `gen_ai.retrieval.query.text` | string | **Required** | User query text | OWASP LLM01 (Prompt Injection), MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |

---

## Span: `gen_ai.retrieval.query.text`

Represents query processing and embedding generation for retrieval.

### Span Name

Format: `rag.query {rag.pipeline.name}`

### Span Kind

`INTERNAL`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.retrieval.query.text` | string | **Required** | User query text | OWASP LLM01, MITRE ATLAS AML.T0051 |
| `rag.query.embedding_model` | string | **Recommended** | Embedding model used | MITRE ATLAS [AML.T0044](https://atlas.mitre.org/techniques/AML.T0044), EU AI Act Art.13 |
| `rag.query.embedding_dimensions` | int | **Optional** | Embedding vector dimensions | — |

---

## Span: `rag.retrieve`

Represents the retrieval/vector search phase of a RAG pipeline.

### Span Name

Format: `rag.retrieve {gen_ai.data_source.id}`

### Span Kind

`CLIENT`

### Normative Field Table

#### Database & Query (CoSAI WS2)

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.data_source.id` | string | **Required** | Vector database name (e.g. `"pinecone"`, `"chromadb"`, `"weaviate"`, `"pgvector"`) | OWASP LLM08 (Vector/Embedding Weaknesses), MITRE ATLAS [AML.T0043](https://atlas.mitre.org/techniques/AML.T0043) |
| `gen_ai.retrieval.query.text` | string | **Required** | Query text used for retrieval | OWASP LLM01, MITRE ATLAS AML.T0051 |
| `rag.retrieve.index` | string | **Recommended** | Index/collection name | NIST AI RMF MAP-1.5 |
| `rag.retrieve.top_k` | int | **Recommended** | Number of results requested | NIST AI RMF MEASURE-2.5 |
| `rag.retrieve.filter` | string | **Optional** | Metadata filter (JSON) | — |

#### Retrieval Results (CoSAI WS2)

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `rag.retrieve.results_count` | int | **Required** | Actual number of documents returned | NIST AI RMF MEASURE-2.5, OWASP LLM08 |
| `rag.retrieval.docs` | string | **Recommended** | JSON array of retrieved document summaries (id, score, snippet) | OWASP LLM08, MITRE ATLAS [AML.T0043](https://atlas.mitre.org/techniques/AML.T0043) |
| `rag.retrieve.min_score` | double | **Recommended** | Minimum similarity score among results | OWASP LLM08, NIST AI RMF MEASURE-2.5 |
| `rag.retrieve.max_score` | double | **Recommended** | Maximum similarity score among results | OWASP LLM08, NIST AI RMF MEASURE-2.5 |

#### Per-Document Fields (CoSAI WS2)

These fields are emitted as span events (`rag.doc.retrieved`) for each retrieved document, or as JSON within `rag.retrieval.docs`:

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `rag.doc.id` | string | **Recommended** | Document/chunk identifier | NIST AI RMF MAP-1.5, EU AI Act Art.12 (Record-Keeping) |
| `rag.doc.score` | double | **Recommended** | Similarity/relevance score for this document (0.0–1.0) | OWASP LLM08, NIST AI RMF MEASURE-2.5 |
| `rag.doc.provenance` | string | **Recommended** | Document source/origin (e.g. URL, collection, upload ID) | OWASP LLM09 (Misinformation), EU AI Act Art.13 (Transparency) |

---

## Span: `rag.rerank`

Represents the optional reranking phase of a RAG pipeline.

### Span Name

Format: `rag.rerank {rag.rerank.model}`

### Span Kind

`CLIENT`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `rag.rerank.model` | string | **Required** | Reranking model name (e.g. `"cross-encoder/ms-marco"`) | MITRE ATLAS AML.T0044, EU AI Act Art.13 |
| `rag.rerank.input_count` | int | **Required** | Number of documents before reranking | NIST AI RMF MEASURE-2.5 |
| `rag.rerank.output_count` | int | **Required** | Number of documents after reranking | NIST AI RMF MEASURE-2.5 |

---

## Span: `rag.evaluate`

Represents quality evaluation of the RAG pipeline output.

### Span Name

Format: `rag.evaluate {rag.pipeline.name}`

### Span Kind

`INTERNAL`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `rag.quality.context_relevance` | double | **Recommended** | Context relevance score (0.0–1.0) | NIST AI RMF MEASURE-2.5 |
| `rag.quality.answer_relevance` | double | **Recommended** | Answer relevance score (0.0–1.0) | NIST AI RMF MEASURE-2.5 |
| `rag.quality.faithfulness` | double | **Recommended** | Answer faithfulness to retrieved context (0.0–1.0) | OWASP LLM09 (Misinformation), NIST AI RMF MEASURE-2.5 |
| `rag.quality.groundedness` | double | **Recommended** | How grounded the answer is in source material (0.0–1.0) | OWASP LLM09, EU AI Act Art.13 |

### Citation & Source Attribution [RFC v0.4 gap closure]

> **RFC field:** Citations / Source Attribution · **Tier:** MUST · **§7 Output Handling, Egress & Refusals** · **Grounding attacks:** `TA-01`, `TA-02`, `IR-03`, `AOC-03`

The quality scores above are model-graded judgements about the answer. They do not record *which sources the answer claimed*, and they cannot be checked against the retrieval that actually happened. A citation the model invented, and a citation the model lifted from injected text inside a retrieved document, both score well on faithfulness while pointing at something the pipeline never returned.

The check that matters is therefore a join: every citation the output asserts is resolved against the document IDs returned by the `rag.retrieve` span in the same trace. Citations that do not resolve are the signal.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `rag.citation.count` | int | **Required** | Number of citations asserted in the output | EU AI Act Art.13 (Transparency) |
| `rag.citation.ids` | string[] | **Required** | Document identifiers the output claims to have drawn on | EU AI Act Art.13, OWASP LLM09 (Misinformation) |
| `rag.citation.retrieval_span_id` | string | **Required** | Span ID of the `rag.retrieve` these citations are resolved against. Without it the resolution is unverifiable | EU AI Act Art.12 (Record-Keeping) |
| `rag.citation.resolved_count` | int | **Required** | Citations matching a document actually returned by that retrieval | OWASP LLM09, NIST AI RMF MEASURE-2.5 |
| `rag.citation.unresolved_count` | int | **Required** | Citations with no matching retrieved document | OWASP LLM09, MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `rag.citation.unresolved_ids` | string[] | **Recommended** | The specific citations that failed to resolve | OWASP LLM09, MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `rag.citation.fabricated` | boolean | **Recommended** | Whether any citation is unresolvable. Convenience flag over `unresolved_count > 0` | OWASP LLM09 (Misinformation) |
| `rag.citation.sources` | string[] | **Recommended** | Source URLs or file identities as presented to the user, which may differ from the internal document IDs | EU AI Act Art.13 (Transparency) |
| `rag.citation.coverage_ratio` | double | **Recommended** | Fraction of retrieved documents actually cited (0.0–1.0) | NIST AI RMF MEASURE-2.5 |
| `rag.citation.uncited_content_ratio` | double | **Recommended** | Fraction of output content carrying no citation. High values with high `faithfulness` indicate the grader saw less of the answer than the user did | OWASP LLM09, NIST AI RMF MEASURE-2.5 |
| `rag.citation.verified` | boolean | **Recommended** | Whether resolution was performed against logged retrieval rather than asserted by the model. Absent or `false` is self-asserted | EU AI Act Art.12, NIST AI RMF GOVERN-1.2 |

**Emission rule.** These fields describe the *output*, so they belong on `rag.evaluate` or on the `gen_ai.inference` span that produced the answer. They MUST NOT be placed on `rag.retrieve`, which by construction cannot know what the model later claimed. Where evaluation is asynchronous, `rag.citation.retrieval_span_id` is what reconnects the two.

---

## Declared Knowledge-Source Configuration [RFC v0.4 gap closure]

> **RFC field:** Declared Knowledge-Source Configuration · **Tier:** MAY · **§11 Retrieval & Content (RAG)** · **Grounding attacks:** `TA-09`, `IR-03`, `TA-02`

The spans above record retrievals that happened. They give no denominator: a retrieval reaching a collection the agent was never configured to query looks identical to a legitimate one. The declared source set supplies that denominator, and because it is configuration rather than hot-path behaviour it is captured once at agent-registration or index-build time and carried on `asset.*` inventory records, not on every `rag.retrieve` span.

The full field list is in the registry under [`rag.source.*`](attributes-registry.md#ragsource-rfc-v04-gap-closure); it covers declared identity (`name`, `description`, `index.name`, `index.namespace`, `schema`), governance (`owner`, `classification`, `trust_level`, `write_access`, `ingestion.method`), and the declared search contract (`search.top_k`, `search.filters`, `search.scoring`, `search.min_score`, `search.reranker`).

One field does belong on the retrieval span: `rag.source.undeclared_access`, set when the retrieval reached a source outside `rag.source.declared`.

---

## CoSAI WS2 Field Mapping

Cross-reference between CoSAI WS2 `RAG_CONTEXT` field names and AITF attribute keys:

| CoSAI WS2 Field | AITF Attribute | Notes |
|---|---|---|
| `rag.database.name` | `gen_ai.data_source.id` | Vector DB identifier |
| `rag.query.text` | `gen_ai.retrieval.query.text` | User query text |
| `rag.retrieval.docs` | `rag.retrieval.docs` | New in CoSAI WS2 alignment; JSON array |
| `rag.doc.id` | `rag.doc.id` | New in CoSAI WS2 alignment |
| `rag.doc.score` | `rag.doc.score` | New in CoSAI WS2 alignment |
| `rag.doc.provenance` | `rag.doc.provenance` | New in CoSAI WS2 alignment |
| — | `rag.citation.*` | New in RFC v0.4 gap closure; no CoSAI WS2 counterpart yet |
| — | `rag.source.*` | New in RFC v0.4 gap closure; declared configuration, no CoSAI WS2 counterpart yet |

---

## Example: RAG Pipeline with Reranking

```
Span: rag.pipeline knowledge-base
  rag.pipeline.name: "knowledge-base"
  rag.pipeline.stage: "generate"
  gen_ai.retrieval.query.text: "What are the OWASP LLM Top 10 risks?"
  |
  +- Span: rag.query knowledge-base
  |    gen_ai.retrieval.query.text: "What are the OWASP LLM Top 10 risks?"
  |    rag.query.embedding_model: "text-embedding-3-small"
  |    rag.query.embedding_dimensions: 1536
  |
  +- Span: rag.retrieve pinecone
  |    gen_ai.data_source.id: "pinecone"
  |    rag.retrieve.index: "security-docs"
  |    rag.retrieve.top_k: 10
  |    rag.retrieve.results_count: 10
  |    rag.retrieve.min_score: 0.72
  |    rag.retrieve.max_score: 0.96
  |    rag.retrieval.docs: "[{\"id\":\"doc-001\",\"score\":0.96,\"provenance\":\"owasp.org/llm-top-10\"}]"
  |    Events:
  |      rag.doc.retrieved: {rag.doc.id: "doc-001", rag.doc.score: 0.96, rag.doc.provenance: "owasp.org/llm-top-10"}
  |      rag.doc.retrieved: {rag.doc.id: "doc-002", rag.doc.score: 0.91, rag.doc.provenance: "atlas.mitre.org"}
  |
  +- Span: rag.rerank cross-encoder/ms-marco
  |    rag.rerank.model: "cross-encoder/ms-marco"
  |    rag.rerank.input_count: 10
  |    rag.rerank.output_count: 5
  |
  +- Span: chat gpt-4o
  |    gen_ai.provider.name: "openai"
  |    gen_ai.usage.input_tokens: 2500
  |    gen_ai.usage.output_tokens: 800
  |
  +- Span: rag.evaluate knowledge-base
       rag.quality.context_relevance: 0.92
       rag.quality.answer_relevance: 0.88
       rag.quality.faithfulness: 0.95
       rag.quality.groundedness: 0.93
       rag.citation.count: 3
       rag.citation.ids: ["doc-001", "doc-002", "doc-417"]
       rag.citation.retrieval_span_id: "a3f1c09b4e2d7856"
       rag.citation.resolved_count: 2
       rag.citation.unresolved_count: 1
       rag.citation.unresolved_ids: ["doc-417"]
       rag.citation.fabricated: true
       rag.citation.verified: true
```

Note the shape of the finding in that example. Every quality score is high — the grader is satisfied the answer follows from the context it was shown. But `doc-417` was never returned by the `rag.retrieve` span, so the answer cites a source the pipeline did not supply. The quality scores alone would not have surfaced this; only the resolution against `rag.citation.retrieval_span_id` does.
