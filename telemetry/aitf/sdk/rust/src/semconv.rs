//! AITF semantic convention constants.
//!
//! Attribute key string constants used across AITF instrumentation, processors,
//! and exporters. OTel GenAI attributes (`gen_ai.*`) are preserved for
//! compatibility; AITF extensions use dedicated namespaces. Values mirror the
//! Go (`semconv/attributes.go`) and Python (`semantic_conventions/attributes.py`)
//! SDKs exactly.

/// OTel GenAI attributes (preserved) plus AITF GenAI extensions.
pub mod gen_ai {
    pub const SYSTEM: &str = "gen_ai.provider.name";
    pub const PROVIDER_NAME: &str = "gen_ai.provider.name";
    pub const OPERATION_NAME: &str = "gen_ai.operation.name";

    // Prompt management
    pub const PROMPT_VERSION: &str = "gen_ai.prompt.version";
    pub const PROMPT_LABEL: &str = "gen_ai.prompt.label";

    // End-user / tagging (general; e.g. Langfuse userId / tags)
    pub const USER_ID: &str = "user.id";
    pub const TAGS: &str = "tags";

    // Evaluation
    pub const EVALUATION_NAME: &str = "gen_ai.evaluation.name";
    pub const EVALUATION_SCORE_VALUE: &str = "gen_ai.evaluation.score.value";
    pub const EVALUATION_SCORE_LABEL: &str = "gen_ai.evaluation.score.label";
    pub const EVALUATION_EXPLANATION: &str = "gen_ai.evaluation.explanation";
    pub const EVALUATION_SCORE_DATA_TYPE: &str = "gen_ai.evaluation.score.data_type";
    pub const EVALUATION_SOURCE: &str = "gen_ai.evaluation.source";
    pub const EVALUATION_COMMENT: &str = "gen_ai.evaluation.comment";
    pub const EVALUATION_DATASET_ITEM_ID: &str = "gen_ai.evaluation.dataset.item_id";

    // Request
    pub const REQUEST_MODEL: &str = "gen_ai.request.model";
    pub const REQUEST_MAX_TOKENS: &str = "gen_ai.request.max_tokens";
    pub const REQUEST_TEMPERATURE: &str = "gen_ai.request.temperature";
    pub const REQUEST_TOP_P: &str = "gen_ai.request.top_p";
    pub const REQUEST_TOP_K: &str = "gen_ai.request.top_k";
    pub const REQUEST_STREAM: &str = "gen_ai.request.stream";
    pub const REQUEST_TOOLS: &str = "gen_ai.request.tools";
    pub const REQUEST_TOOL_CHOICE: &str = "gen_ai.request.tool_choice";
    pub const REQUEST_RESPONSE_FORMAT: &str = "gen_ai.request.response_format";
    pub const REQUEST_FREQUENCY_PENALTY: &str = "gen_ai.request.frequency_penalty";
    pub const REQUEST_PRESENCE_PENALTY: &str = "gen_ai.request.presence_penalty";
    pub const REQUEST_SEED: &str = "gen_ai.request.seed";

    // Response
    pub const RESPONSE_ID: &str = "gen_ai.response.id";
    pub const RESPONSE_MODEL: &str = "gen_ai.response.model";
    pub const RESPONSE_FINISH_REASONS: &str = "gen_ai.response.finish_reasons";

    // Usage
    pub const USAGE_INPUT_TOKENS: &str = "gen_ai.usage.input_tokens";
    pub const USAGE_OUTPUT_TOKENS: &str = "gen_ai.usage.output_tokens";
    pub const USAGE_CACHED_TOKENS: &str = "gen_ai.usage.cached_tokens";
    pub const USAGE_REASONING_TOKENS: &str = "gen_ai.usage.reasoning_tokens";

    // System prompt hash (CoSAI WS2: AI_INTERACTION)
    pub const SYSTEM_PROMPT_HASH: &str = "gen_ai.system_prompt.hash";

    // Events
    pub const PROMPT: &str = "gen_ai.prompt";
    pub const COMPLETION: &str = "gen_ai.completion";
    pub const TOOL_NAME: &str = "gen_ai.tool.name";
    pub const TOOL_CALL_ID: &str = "gen_ai.tool.call_id";
    pub const TOOL_ARGUMENTS: &str = "gen_ai.tool.arguments";
    pub const TOOL_RESULT: &str = "gen_ai.tool.result";
    pub const TOOL_CALL_ARGUMENTS: &str = "gen_ai.tool.call.arguments";
    pub const TOOL_CALL_RESULT: &str = "gen_ai.tool.call.result";

    // Data source / RAG
    pub const DATA_SOURCE_ID: &str = "gen_ai.data_source.id";

    // Identifier hierarchy [RFC v0.4 gap closure]
    pub const TURN_ID: &str = "gen_ai.turn.id";
    pub const TURN_INDEX: &str = "gen_ai.turn.index";
    pub const TURN_PARENT_ID: &str = "gen_ai.turn.parent_id";
    pub const STEP_ID: &str = "gen_ai.step.id";
    pub const STEP_PARENT_ID: &str = "gen_ai.step.parent_id";
    pub const RUN_ID: &str = "gen_ai.run.id";

    // Trigger provenance [RFC v0.4 gap closure]
    pub const TRIGGER_TYPE: &str = "gen_ai.trigger.type";
    pub const TRIGGER_EVENT: &str = "gen_ai.trigger.event";
    pub const TRIGGER_EVENT_ID: &str = "gen_ai.trigger.event.id";
    pub const TRIGGER_SOURCE: &str = "gen_ai.trigger.source";
    pub const TRIGGER_SOURCE_PRINCIPAL: &str = "gen_ai.trigger.source.principal";
    pub const TRIGGER_RECEIVED_AT: &str = "gen_ai.trigger.received_at";
    pub const TRIGGER_HUMAN_IN_LOOP: &str = "gen_ai.trigger.human_in_loop";

    // Content modality & attachment identity [RFC v0.4 gap closure]
    pub const CONTENT_PART_INDEX: &str = "gen_ai.content.part.index";
    pub const CONTENT_PART_TYPE: &str = "gen_ai.content.part.type";
    pub const CONTENT_PART_MIME_TYPE: &str = "gen_ai.content.part.mime_type";
    pub const CONTENT_PART_SIZE_BYTES: &str = "gen_ai.content.part.size_bytes";
    pub const CONTENT_PART_HASH: &str = "gen_ai.content.part.hash";
    pub const CONTENT_MODALITIES: &str = "gen_ai.content.modalities";
    pub const CONTENT_ATTACHMENT_COUNT: &str = "gen_ai.content.attachment.count";
    pub const CONTENT_ATTACHMENT_NAME: &str = "gen_ai.content.attachment.name";
    pub const CONTENT_ATTACHMENT_HASH: &str = "gen_ai.content.attachment.hash";
    pub const CONTENT_ATTACHMENT_SIZE_BYTES: &str = "gen_ai.content.attachment.size_bytes";
    pub const CONTENT_ATTACHMENT_SOURCE: &str = "gen_ai.content.attachment.source";
    pub const CONTENT_ATTACHMENT_EXTRACTED_TEXT_HASH: &str = "gen_ai.content.attachment.extracted_text_hash";
    pub const CONTENT_TOTAL_SIZE_BYTES: &str = "gen_ai.content.total_size_bytes";

    // Backend / route restriction decision [RFC v0.4 gap closure]
    pub const ROUTE_CANDIDATES: &str = "gen_ai.route.candidates";
    pub const ROUTE_CANDIDATES_COUNT: &str = "gen_ai.route.candidates.count";
    pub const ROUTE_SELECTED: &str = "gen_ai.route.selected";
    pub const ROUTE_DECISION: &str = "gen_ai.route.decision";
    pub const ROUTE_CONSTRAINT_TYPE: &str = "gen_ai.route.constraint.type";
    pub const ROUTE_CONSTRAINT_VALUE: &str = "gen_ai.route.constraint.value";
    pub const ROUTE_CONSTRAINT_SOURCE: &str = "gen_ai.route.constraint.source";
    pub const ROUTE_EXCLUDED: &str = "gen_ai.route.excluded";
    pub const ROUTE_NO_CANDIDATE_ACTION: &str = "gen_ai.route.no_candidate_action";
    pub const ROUTE_POLICY_ID: &str = "gen_ai.route.policy_id";

    // Organization / tenant ID [RFC v0.4 gap closure]
    pub const CONVERSATION_TENANT_ID: &str = "gen_ai.conversation.tenant.id";
    pub const USER_TENANT_ID: &str = "user.tenant.id";

    /// Values for gen_ai.trigger.type (RFC v0.4 gap closure).
    pub mod trigger_type {
        pub const USER_INITIATED: &str = "user_initiated";
        pub const SCHEDULED: &str = "scheduled";
        pub const WEBHOOK: &str = "webhook";
        pub const EVENT_DRIVEN: &str = "event_driven";
        pub const AGENT_INITIATED: &str = "agent_initiated";
        pub const SYSTEM: &str = "system";
        pub const RETRY: &str = "retry";
        pub const UNKNOWN: &str = "unknown";
    }

    /// Values for gen_ai.route.decision (RFC v0.4 gap closure).
    pub mod route_decision {
        pub const ALLOWED: &str = "allowed";
        pub const RESTRICTED: &str = "restricted";
        pub const DENIED: &str = "denied";
        pub const FALLBACK: &str = "fallback";
        pub const NO_POLICY: &str = "no_policy";
    }
}

/// AITF agent attributes.
pub mod agent {
    pub const NAME: &str = "gen_ai.agent.name";
    pub const ID: &str = "gen_ai.agent.id";
    pub const TYPE: &str = "agent.type";
    pub const FRAMEWORK: &str = "agent.framework";
    pub const VERSION: &str = "agent.version";
    pub const DESCRIPTION: &str = "gen_ai.agent.description";

    pub const SESSION_ID: &str = "gen_ai.conversation.id";
    pub const CONVERSATION_ID: &str = "gen_ai.conversation.id";
    pub const SESSION_TURN_COUNT: &str = "agent.session.turn_count";

    pub const WORKFLOW_ID: &str = "agent.workflow_id";
    pub const STATE: &str = "agent.state";
    pub const SCRATCHPAD: &str = "agent.scratchpad";
    pub const NEXT_ACTION: &str = "agent.next_action";

    pub const STEP_TYPE: &str = "agent.step.type";
    pub const STEP_INDEX: &str = "agent.step.index";
    pub const STEP_THOUGHT: &str = "agent.step.thought";
    pub const STEP_ACTION: &str = "agent.step.action";
    pub const STEP_OBSERVATION: &str = "agent.step.observation";
    pub const STEP_STATUS: &str = "agent.step.status";

    pub const DELEGATION_TARGET_AGENT: &str = "agent.delegation.target_agent";
    pub const DELEGATION_TARGET_AGENT_ID: &str = "agent.delegation.target_agent_id";
    pub const DELEGATION_REASON: &str = "agent.delegation.reason";
    pub const DELEGATION_STRATEGY: &str = "agent.delegation.strategy";
    pub const DELEGATION_TASK: &str = "agent.delegation.task";

    pub const TEAM_NAME: &str = "agent.team.name";
    pub const TEAM_ID: &str = "agent.team.id";
    pub const TEAM_TOPOLOGY: &str = "agent.team.topology";
    pub const TEAM_COORDINATOR: &str = "agent.team.coordinator";
    pub const TEAM_CONSENSUS_METHOD: &str = "agent.team.consensus_method";

    // Peer agent card / descriptor [RFC v0.4 gap closure]
    pub const PEER_ID: &str = "gen_ai.agent.peer.id";
    pub const PEER_NAME: &str = "gen_ai.agent.peer.name";
    pub const PEER_URL: &str = "gen_ai.agent.peer.url";
    pub const PEER_VERSION: &str = "gen_ai.agent.peer.version";
    pub const PEER_PROVIDER: &str = "gen_ai.agent.peer.provider";
    pub const PEER_SKILLS: &str = "gen_ai.agent.peer.skills";
    pub const PEER_PROTOCOL: &str = "gen_ai.agent.peer.protocol";
    pub const PEER_CARD_HASH: &str = "gen_ai.agent.peer.card.hash";
    pub const PEER_CARD_BASELINE_HASH: &str = "gen_ai.agent.peer.card.baseline_hash";
    pub const PEER_CARD_CHANGED: &str = "gen_ai.agent.peer.card.changed";
    pub const PEER_CARD_CHANGE_FIELDS: &str = "gen_ai.agent.peer.card.change_fields";
    pub const PEER_CARD_FIRST_SEEN: &str = "gen_ai.agent.peer.card.first_seen";
    pub const PEER_VERIFICATION_METHOD: &str = "gen_ai.agent.peer.verification.method";
    pub const PEER_VERIFICATION_RESULT: &str = "gen_ai.agent.peer.verification.result";
    pub const PEER_VERIFICATION_AUTHORITY: &str = "gen_ai.agent.peer.verification.authority";
    pub const PEER_APPROVED: &str = "gen_ai.agent.peer.approved";

    // Organization / tenant ID [RFC v0.4 gap closure]
    pub const TENANT_ID: &str = "gen_ai.agent.tenant.id";

    /// Values for gen_ai.agent.peer.verification.result (RFC v0.4 gap closure).
    pub mod peer_verification_result {
        pub const VERIFIED: &str = "verified";
        pub const UNVERIFIED: &str = "unverified";
        pub const FAILED: &str = "failed";
        pub const REFUSED: &str = "refused";
        pub const NOT_ATTEMPTED: &str = "not_attempted";
    }
}

/// AITF MCP attributes.
pub mod mcp {
    pub const SERVER_NAME: &str = "mcp.server.name";
    pub const SERVER_VERSION: &str = "mcp.server.version";
    pub const SERVER_TRANSPORT: &str = "mcp.server.transport";
    pub const SERVER_URL: &str = "mcp.server.url";
    pub const PROTOCOL_VERSION: &str = "mcp.protocol.version";

    pub const TOOL_NAME: &str = "gen_ai.tool.name";
    pub const TOOL_SERVER: &str = "mcp.tool.server";
    pub const TOOL_INPUT: &str = "gen_ai.tool.call.arguments";
    pub const TOOL_OUTPUT: &str = "gen_ai.tool.call.result";
    pub const TOOL_IS_ERROR: &str = "mcp.tool.is_error";
    pub const TOOL_DURATION_MS: &str = "mcp.tool.duration_ms";
    pub const TOOL_APPROVAL_REQUIRED: &str = "mcp.tool.approval_required";
    pub const TOOL_APPROVED: &str = "mcp.tool.approved";
    pub const TOOL_COUNT: &str = "mcp.tool.count";

    pub const RESOURCE_URI: &str = "mcp.resource.uri";
    pub const RESOURCE_NAME: &str = "mcp.resource.name";
    pub const RESOURCE_MIME_TYPE: &str = "mcp.resource.mime_type";

    // Tool definition digest [RFC v0.4 gap closure]
    pub const TOOL_DEFINITION_HASH: &str = "mcp.tool.definition.hash";
    pub const TOOL_DEFINITION_BASELINE_HASH: &str = "mcp.tool.definition.baseline_hash";
    pub const TOOL_DEFINITION_CHANGED: &str = "mcp.tool.definition.changed";
    pub const TOOL_DEFINITION_CHANGE_TYPE: &str = "mcp.tool.definition.change_type";
    pub const TOOL_DEFINITION_DESCRIPTION_HASH: &str = "mcp.tool.definition.description_hash";
    pub const TOOL_DEFINITION_SCHEMA_HASH: &str = "mcp.tool.definition.schema_hash";
    pub const TOOL_DEFINITION_APPROVED: &str = "mcp.tool.definition.approved";
    pub const TOOL_DEFINITION_APPROVED_AT: &str = "mcp.tool.definition.approved_at";
    pub const TOOL_DEFINITION_FIRST_SEEN: &str = "mcp.tool.definition.first_seen";
    pub const TOOL_DEFINITION_SOURCE: &str = "mcp.tool.definition.source";

    // Server identity & primitive [RFC v0.4 gap closure]
    pub const PRIMITIVE: &str = "mcp.primitive";
    pub const METHOD_NAME: &str = "mcp.method.name";
    pub const REQUEST_ID: &str = "mcp.request.id";
    pub const SESSION_ID: &str = "mcp.session.id";
    pub const SERVER_INSTANCE_ID: &str = "mcp.server.instance.id";
    pub const SERVER_IDENTITY_METHOD: &str = "mcp.server.identity.method";
    pub const SERVER_IDENTITY_VERIFIED: &str = "mcp.server.identity.verified";
    pub const SERVER_TRUST_DOMAIN: &str = "mcp.server.trust_domain";
    pub const SERVER_COMMAND: &str = "mcp.server.command";
    pub const SERVER_BINARY_HASH: &str = "mcp.server.binary.hash";
    pub const CAPABILITIES_NEGOTIATED: &str = "mcp.capabilities.negotiated";

    // Protocol envelope capture [RFC v0.4 gap closure]
    pub const ENVELOPE_CAPTURED: &str = "mcp.envelope.captured";
    pub const ENVELOPE_REQUEST: &str = "mcp.envelope.request";
    pub const ENVELOPE_RESPONSE: &str = "mcp.envelope.response";
    pub const ENVELOPE_REQUEST_HASH: &str = "mcp.envelope.request.hash";
    pub const ENVELOPE_RESPONSE_HASH: &str = "mcp.envelope.response.hash";
    pub const ENVELOPE_SIZE_BYTES: &str = "mcp.envelope.size_bytes";
    pub const ENVELOPE_TRUNCATED: &str = "mcp.envelope.truncated";
    pub const ENVELOPE_REDACTED: &str = "mcp.envelope.redacted";
    pub const ENVELOPE_JSONRPC_VERSION: &str = "mcp.envelope.jsonrpc.version";
    pub const ENVELOPE_ERROR_CODE: &str = "mcp.envelope.error.code";
    pub const ENVELOPE_ERROR_MESSAGE: &str = "mcp.envelope.error.message";

    /// Values for mcp.primitive (RFC v0.4 gap closure).
    pub mod primitive {
        pub const TOOL: &str = "tool";
        pub const RESOURCE: &str = "resource";
        pub const PROMPT: &str = "prompt";
        pub const SAMPLING: &str = "sampling";
        pub const COMPLETION: &str = "completion";
        pub const ELICITATION: &str = "elicitation";
        pub const ROOT: &str = "root";
        pub const LOGGING: &str = "logging";
    }
}

/// AITF skill attributes.
pub mod skill {
    pub const NAME: &str = "skill.name";
    pub const ID: &str = "skill.id";
    pub const VERSION: &str = "skill.version";
    pub const PROVIDER: &str = "skill.provider";
    pub const CATEGORY: &str = "skill.category";
    pub const DESCRIPTION: &str = "skill.description";
    pub const INPUT: &str = "skill.input";
    pub const OUTPUT: &str = "skill.output";
    pub const STATUS: &str = "skill.status";
    pub const DURATION_MS: &str = "skill.duration_ms";
}

/// AITF RAG attributes.
pub mod rag {
    pub const PIPELINE_NAME: &str = "rag.pipeline.name";
    pub const PIPELINE_STAGE: &str = "rag.pipeline.stage";
    pub const QUERY: &str = "rag.query";

    pub const RETRIEVE_DATABASE: &str = "gen_ai.data_source.id";
    pub const RETRIEVE_INDEX: &str = "rag.retrieve.index";
    pub const RETRIEVE_TOP_K: &str = "rag.retrieve.top_k";
    pub const RETRIEVE_RESULTS_COUNT: &str = "rag.retrieve.results_count";
    pub const RETRIEVE_MIN_SCORE: &str = "rag.retrieve.min_score";
    pub const RETRIEVE_MAX_SCORE: &str = "rag.retrieve.max_score";
    pub const RETRIEVE_FILTER: &str = "rag.retrieve.filter";

    // Citations / source attribution [RFC v0.4 gap closure]
    pub const CITATION_COUNT: &str = "rag.citation.count";
    pub const CITATION_IDS: &str = "rag.citation.ids";
    pub const CITATION_SOURCES: &str = "rag.citation.sources";
    pub const CITATION_RESOLVED_COUNT: &str = "rag.citation.resolved_count";
    pub const CITATION_UNRESOLVED_COUNT: &str = "rag.citation.unresolved_count";
    pub const CITATION_UNRESOLVED_IDS: &str = "rag.citation.unresolved_ids";
    pub const CITATION_FABRICATED: &str = "rag.citation.fabricated";
    pub const CITATION_RETRIEVAL_SPAN_ID: &str = "rag.citation.retrieval_span_id";
    pub const CITATION_COVERAGE_RATIO: &str = "rag.citation.coverage_ratio";
    pub const CITATION_UNCITED_CONTENT_RATIO: &str = "rag.citation.uncited_content_ratio";
    pub const CITATION_FORMAT: &str = "rag.citation.format";
    pub const CITATION_VERIFIED: &str = "rag.citation.verified";

    // Declared knowledge-source configuration [RFC v0.4 gap closure]
    pub const SOURCE_DECLARED: &str = "rag.source.declared";
    pub const SOURCE_DECLARED_COUNT: &str = "rag.source.declared_count";
    pub const SOURCE_NAME: &str = "rag.source.name";
    pub const SOURCE_DESCRIPTION: &str = "rag.source.description";
    pub const SOURCE_TYPE: &str = "rag.source.type";
    pub const SOURCE_INDEX_NAME: &str = "rag.source.index.name";
    pub const SOURCE_INDEX_NAMESPACE: &str = "rag.source.index.namespace";
    pub const SOURCE_SCHEMA: &str = "rag.source.schema";
    pub const SOURCE_SCHEMA_HASH: &str = "rag.source.schema.hash";
    pub const SOURCE_TRUST_LEVEL: &str = "rag.source.trust_level";
    pub const SOURCE_CLASSIFICATION: &str = "rag.source.classification";
    pub const SOURCE_OWNER: &str = "rag.source.owner";
    pub const SOURCE_WRITE_ACCESS: &str = "rag.source.write_access";
    pub const SOURCE_INGESTION_METHOD: &str = "rag.source.ingestion.method";
    pub const SOURCE_LAST_INDEXED: &str = "rag.source.last_indexed";
    pub const SOURCE_INDEX_HASH: &str = "rag.source.index.hash";
    pub const SOURCE_SEARCH_TOP_K: &str = "rag.source.search.top_k";
    pub const SOURCE_SEARCH_FILTERS: &str = "rag.source.search.filters";
    pub const SOURCE_SEARCH_SCORING: &str = "rag.source.search.scoring";
    pub const SOURCE_SEARCH_MIN_SCORE: &str = "rag.source.search.min_score";
    pub const SOURCE_SEARCH_RERANKER: &str = "rag.source.search.reranker";
    pub const SOURCE_UNDECLARED_ACCESS: &str = "rag.source.undeclared_access";

    // Organization / tenant ID [RFC v0.4 gap closure]
    pub const DATA_SOURCE_TENANT_ID: &str = "gen_ai.data_source.tenant.id";
}

/// AITF security attributes.
pub mod security {
    pub const RISK_SCORE: &str = "security.risk_score";
    pub const RISK_LEVEL: &str = "security.risk_level";
    pub const THREAT_DETECTED: &str = "security.threat_detected";
    pub const THREAT_TYPE: &str = "security.threat_type";
    pub const OWASP_CATEGORY: &str = "security.owasp_category";
    pub const BLOCKED: &str = "security.blocked";
    pub const DETECTION_METHOD: &str = "security.detection_method";
    pub const CONFIDENCE: &str = "security.confidence";

    // Guardrail modification record [RFC v0.4 gap closure]
    pub const GUARDRAIL_ACTION: &str = "security.guardrail.action";
    pub const GUARDRAIL_MODIFIED: &str = "security.guardrail.modified";
    pub const GUARDRAIL_MODIFICATION_TYPE: &str = "security.guardrail.modification.type";
    pub const GUARDRAIL_MODIFICATION_SIDE: &str = "security.guardrail.modification.side";
    pub const GUARDRAIL_MODIFICATION_ENFORCEMENT_POINT: &str = "security.guardrail.modification.enforcement_point";
    pub const GUARDRAIL_MODIFICATION_HASH_BEFORE: &str = "security.guardrail.modification.hash_before";
    pub const GUARDRAIL_MODIFICATION_HASH_AFTER: &str = "security.guardrail.modification.hash_after";
    pub const GUARDRAIL_MODIFICATION_COUNT: &str = "security.guardrail.modification.count";
    pub const GUARDRAIL_MODIFICATION_BYTES_REMOVED: &str = "security.guardrail.modification.bytes_removed";
    pub const GUARDRAIL_MODIFICATION_REDACTION_MAP: &str = "security.guardrail.modification.redaction_map";
    pub const GUARDRAIL_MODIFICATION_REASON: &str = "security.guardrail.modification.reason";

    // Encoded / obfuscated payload indicator [RFC v0.4 gap closure]
    pub const OBFUSCATION_DETECTED: &str = "security.obfuscation.detected";
    pub const OBFUSCATION_ENCODINGS: &str = "security.obfuscation.encodings";
    pub const OBFUSCATION_DEPTH: &str = "security.obfuscation.depth";
    pub const OBFUSCATION_DECODED_FORM: &str = "security.obfuscation.decoded_form";
    pub const OBFUSCATION_DECODED_HASH: &str = "security.obfuscation.decoded_hash";
    pub const OBFUSCATION_DECODED_LENGTH: &str = "security.obfuscation.decoded_length";
    pub const OBFUSCATION_SOURCE_FIELD: &str = "security.obfuscation.source_field";
    pub const OBFUSCATION_INSPECTED_AFTER_DECODE: &str = "security.obfuscation.inspected_after_decode";

    // Authorization decision record [RFC v0.4 gap closure]
    pub const AUTHORIZATION_DECISION: &str = "security.authorization.decision";
    pub const AUTHORIZATION_DECISION_ID: &str = "security.authorization.decision_id";
    pub const AUTHORIZATION_OPERATION: &str = "security.authorization.operation";
    pub const AUTHORIZATION_REASON: &str = "security.authorization.reason";
    pub const AUTHORIZATION_REASON_CODE: &str = "security.authorization.reason_code";
    pub const AUTHORIZATION_AUTHORITY_TYPE: &str = "security.authorization.authority.type";
    pub const AUTHORIZATION_AUTHORITY_ID: &str = "security.authorization.authority.id";
    pub const AUTHORIZATION_AUTHORITY_ENGINE: &str = "security.authorization.authority.engine";
    pub const AUTHORIZATION_RULE_ID: &str = "security.authorization.rule_id";
    pub const AUTHORIZATION_POLICY_VERSION: &str = "security.authorization.policy_version";
    pub const AUTHORIZATION_OBLIGATIONS: &str = "security.authorization.obligations";
    pub const AUTHORIZATION_OBLIGATIONS_FULFILLED: &str = "security.authorization.obligations_fulfilled";
    pub const AUTHORIZATION_PRINCIPAL: &str = "security.authorization.principal";
    pub const AUTHORIZATION_RESOURCE: &str = "security.authorization.resource";
    pub const AUTHORIZATION_LATENCY_MS: &str = "security.authorization.latency_ms";

    // Session taint labels [RFC v0.4 gap closure]
    pub const TAINT_LABELS: &str = "security.taint.labels";
    pub const TAINT_SCOPE: &str = "security.taint.scope";
    pub const TAINT_LABEL_COUNT: &str = "security.taint.label_count";
    pub const TAINT_APPLIED_BY: &str = "security.taint.applied_by";
    pub const TAINT_ORIGIN: &str = "security.taint.origin";
    pub const TAINT_ORIGIN_SPAN_ID: &str = "security.taint.origin_span_id";
    pub const TAINT_FLOW_DECISION: &str = "security.taint.flow.decision";
    pub const TAINT_DENIED: &str = "security.taint.denied";
    pub const TAINT_DENIED_LABELS: &str = "security.taint.denied_labels";
    pub const TAINT_DECLASSIFIED_BY: &str = "security.taint.declassified_by";
    pub const TAINT_DECLASSIFICATION_REASON: &str = "security.taint.declassification_reason";

    // Enforcement-point availability [RFC v0.4 gap closure]
    pub const ENFORCEMENT_POINT_NAME: &str = "security.enforcement.point.name";
    pub const ENFORCEMENT_POINT_TYPE: &str = "security.enforcement.point.type";
    pub const ENFORCEMENT_POINT_VERSION: &str = "security.enforcement.point.version";
    pub const ENFORCEMENT_REACHED: &str = "security.enforcement.reached";
    pub const ENFORCEMENT_LATENCY_MS: &str = "security.enforcement.latency_ms";
    pub const ENFORCEMENT_TIMEOUT_MS: &str = "security.enforcement.timeout_ms";
    pub const ENFORCEMENT_FAILURE_MODE: &str = "security.enforcement.failure_mode";
    pub const ENFORCEMENT_FAILURE_REASON: &str = "security.enforcement.failure_reason";
    pub const ENFORCEMENT_ACTION_TAKEN: &str = "security.enforcement.action_taken";
    pub const ENFORCEMENT_DEGRADED_MODE: &str = "security.enforcement.degraded_mode";

    // Attribute source / trusted-provenance marking [RFC v0.4 gap closure]
    pub const ATTRIBUTE_SOURCE_DEFAULT: &str = "security.attribute_source.default";
    pub const ATTRIBUTE_SOURCE_MAP: &str = "security.attribute_source.map";
    pub const ATTRIBUTE_SOURCE_SELF_ASSERTED: &str = "security.attribute_source.self_asserted";
    pub const ATTRIBUTE_SOURCE_VERIFIED: &str = "security.attribute_source.verified";
    pub const ATTRIBUTE_SOURCE_AUTHORITY_ID: &str = "security.attribute_source.authority.id";
    pub const ATTRIBUTE_SOURCE_VERIFICATION_TIME: &str = "security.attribute_source.verification_time";

    // Mediation coverage & bypass path [RFC v0.4 gap closure]
    pub const MEDIATION_MEDIATED: &str = "security.mediation.mediated";
    pub const MEDIATION_PLACEMENT: &str = "security.mediation.placement";
    pub const MEDIATION_REFERENCE_MONITOR_ID: &str = "security.mediation.reference_monitor.id";
    pub const MEDIATION_BYPASS_AVAILABLE: &str = "security.mediation.bypass_available";
    pub const MEDIATION_BYPASS_PATHS: &str = "security.mediation.bypass_paths";
    pub const MEDIATION_COVERAGE_RATIO: &str = "security.mediation.coverage_ratio";
    pub const MEDIATION_CAPABILITY: &str = "security.mediation.capability";
    pub const MEDIATION_ASSESSED_AT: &str = "security.mediation.assessed_at";

    // Tenant-crossing detection [RFC v0.4 gap closure]
    pub const TENANT_CROSSING_DETECTED: &str = "security.tenant.crossing_detected";
    pub const TENANT_CROSSING_TYPE: &str = "security.tenant.crossing_type";
    pub const TENANT_EXPECTED: &str = "security.tenant.expected";

    /// Values for security.authorization.decision (RFC v0.4 gap closure).
    pub mod authorization_decision {
        pub const ALLOW: &str = "allow";
        pub const DENY: &str = "deny";
        pub const CHALLENGE: &str = "challenge";
        pub const NOT_APPLICABLE: &str = "not_applicable";
        pub const ERROR: &str = "error";
    }

    /// Values for security.taint.flow.decision (RFC v0.4 gap closure).
    pub mod taint_flow_decision {
        pub const ALLOWED: &str = "allowed";
        pub const DENIED: &str = "denied";
        pub const DECLASSIFIED: &str = "declassified";
        pub const NOT_EVALUATED: &str = "not_evaluated";
    }

    /// Values for security.taint.scope (RFC v0.4 gap closure).
    pub mod taint_scope {
        pub const SESSION: &str = "session";
        pub const TURN: &str = "turn";
        pub const MESSAGE: &str = "message";
        pub const RUN: &str = "run";
    }

    /// Values for security.enforcement.failure_mode (RFC v0.4 gap closure).
    pub mod enforcement_failure_mode {
        pub const NONE: &str = "none";
        pub const FAIL_OPEN: &str = "fail_open";
        pub const FAIL_CLOSED: &str = "fail_closed";
        pub const TIMEOUT: &str = "timeout";
        pub const UNREACHABLE: &str = "unreachable";
        pub const DEGRADED: &str = "degraded";
    }

    /// Values for security.attribute_source.* (RFC v0.4 gap closure).
    pub mod attribute_source {
        pub const SELF_ASSERTED: &str = "self_asserted";
        pub const VERIFIED: &str = "verified";
        pub const DERIVED: &str = "derived";
        pub const UNKNOWN: &str = "unknown";
    }
}

/// AITF supply-chain attributes.
pub mod supply_chain {
    pub const MODEL_SOURCE: &str = "supply_chain.model_source";
    pub const MODEL_HASH: &str = "supply_chain.model_hash";
    pub const MODEL_LICENSE: &str = "supply_chain.model_license";
    pub const MODEL_SIGNED: &str = "supply_chain.model_signed";
    pub const MODEL_SIGNER: &str = "supply_chain.model_signer";
    pub const AI_BOM_ID: &str = "supply_chain.ai_bom_id";
    pub const AI_BOM_COMPONENTS: &str = "supply_chain.ai_bom_components";

    // Execution environment / sandbox [RFC v0.4 gap closure]
    pub const RUNTIME_SANDBOX_MODE: &str = "supply_chain.runtime.sandbox.mode";
    pub const RUNTIME_SANDBOX_PROVIDER: &str = "supply_chain.runtime.sandbox.provider";
    pub const RUNTIME_LANGUAGE_NAME: &str = "supply_chain.runtime.language.name";
    pub const RUNTIME_LANGUAGE_VERSION: &str = "supply_chain.runtime.language.version";
    pub const RUNTIME_OS_TYPE: &str = "supply_chain.runtime.os.type";
    pub const RUNTIME_OS_VERSION: &str = "supply_chain.runtime.os.version";
    pub const RUNTIME_ARCHITECTURE: &str = "supply_chain.runtime.architecture";
    pub const RUNTIME_INSTANCE_ID: &str = "supply_chain.runtime.instance.id";
    pub const RUNTIME_IMAGE_DIGEST: &str = "supply_chain.runtime.image.digest";
    pub const RUNTIME_IMAGE_REF: &str = "supply_chain.runtime.image.ref";
    pub const RUNTIME_PRIVILEGED: &str = "supply_chain.runtime.privileged";
    pub const RUNTIME_USER: &str = "supply_chain.runtime.user";
    pub const RUNTIME_CAPABILITIES: &str = "supply_chain.runtime.capabilities";
    pub const RUNTIME_NETWORK_EGRESS_POLICY: &str = "supply_chain.runtime.network.egress_policy";
    pub const RUNTIME_NETWORK_EGRESS_ALLOWLIST: &str = "supply_chain.runtime.network.egress_allowlist";
    pub const RUNTIME_NETWORK_NAMESPACE: &str = "supply_chain.runtime.network.namespace";
    pub const RUNTIME_FILESYSTEM_MODE: &str = "supply_chain.runtime.filesystem.mode";
    pub const RUNTIME_FILESYSTEM_MOUNTS: &str = "supply_chain.runtime.filesystem.mounts";
    pub const RUNTIME_RESOURCE_CPU_LIMIT: &str = "supply_chain.runtime.resource.cpu_limit";
    pub const RUNTIME_RESOURCE_MEMORY_LIMIT_BYTES: &str = "supply_chain.runtime.resource.memory_limit_bytes";
    pub const RUNTIME_RESOURCE_TIMEOUT_MS: &str = "supply_chain.runtime.resource.timeout_ms";
    pub const RUNTIME_SECRETS_EXPOSED: &str = "supply_chain.runtime.secrets.exposed";
    pub const RUNTIME_SECRETS_COUNT: &str = "supply_chain.runtime.secrets.count";
    pub const RUNTIME_ATTESTATION_METHOD: &str = "supply_chain.runtime.attestation.method";
    pub const RUNTIME_ATTESTATION_VERIFIED: &str = "supply_chain.runtime.attestation.verified";
    pub const RUNTIME_ESCAPE_DETECTED: &str = "supply_chain.runtime.escape_detected";
    pub const RUNTIME_ESCAPE_INDICATOR: &str = "supply_chain.runtime.escape_indicator";
}

/// AITF compliance framework attributes.
pub mod compliance {
    pub const FRAMEWORKS: &str = "compliance.frameworks";
    pub const NIST_CONTROLS: &str = "compliance.nist_ai_rmf.controls";
    pub const MITRE_TECHNIQUES: &str = "compliance.mitre_atlas.techniques";
    pub const ISO_CONTROLS: &str = "compliance.iso_42001.controls";
    pub const EU_ARTICLES: &str = "compliance.eu_ai_act.articles";
    pub const SOC2_CONTROLS: &str = "compliance.soc2.controls";
    pub const GDPR_ARTICLES: &str = "compliance.gdpr.articles";
    pub const CCPA_SECTIONS: &str = "compliance.ccpa.sections";
    pub const CSA_AICM_CONTROLS: &str = "compliance.csa_aicm.controls";

    // Authorization decision record [RFC v0.4 gap closure]
    pub const FRAMEWORK: &str = "compliance.framework";
    pub const CONTROL_ID: &str = "compliance.control_id";
}

/// AITF latency attributes.
pub mod latency {
    pub const TOTAL_MS: &str = "latency.total_ms";
    pub const TIME_TO_FIRST_TOKEN_MS: &str = "latency.time_to_first_token_ms";
    pub const TOKENS_PER_SECOND: &str = "latency.tokens_per_second";
    pub const QUEUE_TIME_MS: &str = "latency.queue_time_ms";
    pub const INFERENCE_TIME_MS: &str = "latency.inference_time_ms";
}

/// AITF cost attributes.
pub mod cost {
    pub const INPUT_COST: &str = "cost.input_cost";
    pub const OUTPUT_COST: &str = "cost.output_cost";
    pub const TOTAL_COST: &str = "cost.total_cost";
    pub const CURRENCY: &str = "cost.currency";
}

/// AITF quality attributes.
pub mod quality {
    pub const HALLUCINATION_SCORE: &str = "quality.hallucination_score";
    pub const CONFIDENCE: &str = "quality.confidence";
    pub const FACTUALITY: &str = "quality.factuality";
    pub const TOXICITY_SCORE: &str = "quality.toxicity_score";
    pub const FEEDBACK_RATING: &str = "quality.feedback.rating";
    pub const FEEDBACK_THUMBS: &str = "quality.feedback.thumbs";
}

/// AITF model-operations attributes (subset used by the mapper).
pub mod model_ops {
    pub const TRAINING_RUN_ID: &str = "model_ops.training.run_id";
    pub const TRAINING_TYPE: &str = "model_ops.training.type";
    pub const TRAINING_BASE_MODEL: &str = "model_ops.training.base_model";
    pub const TRAINING_DATASET_ID: &str = "model_ops.training.dataset.id";
    pub const TRAINING_EPOCHS: &str = "model_ops.training.epochs";
    pub const TRAINING_LOSS_FINAL: &str = "model_ops.training.loss_final";
    pub const TRAINING_OUTPUT_MODEL_ID: &str = "model_ops.training.output_model.id";
    pub const TRAINING_STATUS: &str = "model_ops.training.status";

    pub const EVALUATION_RUN_ID: &str = "model_ops.evaluation.run_id";
    pub const EVALUATION_MODEL_ID: &str = "model_ops.evaluation.model_id";
    pub const EVALUATION_TYPE: &str = "model_ops.evaluation.type";
    pub const EVALUATION_METRICS: &str = "model_ops.evaluation.metrics";
    pub const EVALUATION_PASS: &str = "model_ops.evaluation.pass";

    pub const REGISTRY_MODEL_ID: &str = "model_ops.registry.model_id";

    pub const DEPLOYMENT_ID: &str = "model_ops.deployment.id";
    pub const DEPLOYMENT_MODEL_ID: &str = "model_ops.deployment.model_id";
    pub const DEPLOYMENT_STRATEGY: &str = "model_ops.deployment.strategy";
    pub const DEPLOYMENT_ENVIRONMENT: &str = "model_ops.deployment.environment";
    pub const DEPLOYMENT_ENDPOINT: &str = "model_ops.deployment.endpoint";
    pub const DEPLOYMENT_STATUS: &str = "model_ops.deployment.status";

    pub const SERVING_ROUTE_SELECTED_MODEL: &str = "model_ops.serving.route.selected_model";
    pub const SERVING_FALLBACK_CHAIN: &str = "model_ops.serving.fallback.chain";
    pub const SERVING_CACHE_HIT: &str = "model_ops.serving.cache.hit";

    pub const MONITORING_CHECK_TYPE: &str = "model_ops.monitoring.check_type";
    pub const MONITORING_DRIFT_SCORE: &str = "model_ops.monitoring.drift_score";
    pub const MONITORING_DRIFT_TYPE: &str = "model_ops.monitoring.drift_type";
    pub const MONITORING_ACTION_TRIGGERED: &str = "model_ops.monitoring.action_triggered";
}

/// AITF drift-detection attributes (subset used by the mapper).
pub mod drift {
    pub const MODEL_ID: &str = "drift.model_id";
    pub const TYPE: &str = "drift.type";
    pub const SCORE: &str = "drift.score";
    pub const ACTION_TRIGGERED: &str = "drift.action_triggered";
}

/// AITF asset-inventory attributes.
pub mod asset_inventory {
    pub const ID: &str = "asset.id";
    pub const NAME: &str = "asset.name";
    pub const TYPE: &str = "asset.type";
    pub const VERSION: &str = "asset.version";
    pub const OWNER: &str = "asset.owner";
    pub const DEPLOYMENT_ENVIRONMENT: &str = "asset.deployment_environment";
    pub const RISK_CLASSIFICATION: &str = "asset.risk_classification";

    pub const DISCOVERY_SCOPE: &str = "asset.discovery.scope";
    pub const DISCOVERY_METHOD: &str = "asset.discovery.method";
    pub const DISCOVERY_ASSETS_FOUND: &str = "asset.discovery.assets_found";
    pub const DISCOVERY_NEW_ASSETS: &str = "asset.discovery.new_assets";
    pub const DISCOVERY_SHADOW_ASSETS: &str = "asset.discovery.shadow_assets";

    pub const AUDIT_TYPE: &str = "asset.audit.type";
    pub const AUDIT_RESULT: &str = "asset.audit.result";
    pub const AUDIT_FRAMEWORK: &str = "asset.audit.framework";
    pub const AUDIT_FINDINGS: &str = "asset.audit.findings";

    pub const CLASSIFICATION_FRAMEWORK: &str = "asset.classification.framework";
    pub const CLASSIFICATION_PREVIOUS: &str = "asset.classification.previous";
    pub const CLASSIFICATION_REASON: &str = "asset.classification.reason";

    // Organization / tenant ID [RFC v0.4 gap closure]
    pub const TENANT_ID: &str = "asset.tenant.id";
    pub const ORGANIZATION_ID: &str = "asset.organization.id";
    pub const ORGANIZATION_NAME: &str = "asset.organization.name";
    pub const TENANT_TIER: &str = "asset.tenant.tier";
    pub const TENANT_ISOLATION_BOUNDARY: &str = "asset.tenant.isolation_boundary";
    pub const TENANT_DATA_RESIDENCY: &str = "asset.tenant.data_residency";

    // Capability-set change event [RFC v0.4 gap closure]
    pub const CAPABILITY_CHANGE_TYPE: &str = "asset.capability.change_type";
    pub const CAPABILITY_SET_HASH: &str = "asset.capability.set.hash";
    pub const CAPABILITY_SET_PREVIOUS_HASH: &str = "asset.capability.set.previous_hash";
    pub const CAPABILITY_SET_SIZE: &str = "asset.capability.set.size";
    pub const CAPABILITY_ADDED: &str = "asset.capability.added";
    pub const CAPABILITY_REMOVED: &str = "asset.capability.removed";
    pub const CAPABILITY_MODIFIED: &str = "asset.capability.modified";
    pub const CAPABILITY_CATEGORY: &str = "asset.capability.category";
    pub const CAPABILITY_RISK_DELTA: &str = "asset.capability.risk_delta";
    pub const CAPABILITY_PRIVILEGED_ADDED: &str = "asset.capability.privileged_added";
    pub const CAPABILITY_CHANGE_SOURCE: &str = "asset.capability.change_source";
    pub const CAPABILITY_CHANGE_ACTOR: &str = "asset.capability.change_actor";
    pub const CAPABILITY_APPROVED: &str = "asset.capability.approved";
    pub const CAPABILITY_APPROVAL_REF: &str = "asset.capability.approval_ref";
    pub const CAPABILITY_DETECTED_AT: &str = "asset.capability.detected_at";
    pub const CAPABILITY_DETECTION_METHOD: &str = "asset.capability.detection_method";

    // Instrumentation coverage / hook attestation [RFC v0.4 gap closure]
    pub const INSTRUMENTATION_ENABLED: &str = "asset.instrumentation.enabled";
    pub const INSTRUMENTATION_VERSION: &str = "asset.instrumentation.version";
    pub const INSTRUMENTATION_SDK: &str = "asset.instrumentation.sdk";
    pub const INSTRUMENTATION_HOOKS_DECLARED: &str = "asset.instrumentation.hooks.declared";
    pub const INSTRUMENTATION_HOOKS_ACTIVE: &str = "asset.instrumentation.hooks.active";
    pub const INSTRUMENTATION_HOOKS_MISSING: &str = "asset.instrumentation.hooks.missing";
    pub const INSTRUMENTATION_COVERAGE_RATIO: &str = "asset.instrumentation.coverage_ratio";
    pub const INSTRUMENTATION_UNINSTRUMENTED_PATHS: &str = "asset.instrumentation.uninstrumented_paths";
    pub const INSTRUMENTATION_ATTESTATION_METHOD: &str = "asset.instrumentation.attestation.method";
    pub const INSTRUMENTATION_ATTESTATION_VERIFIED: &str = "asset.instrumentation.attestation.verified";
    pub const INSTRUMENTATION_ATTESTATION_SIGNATURE: &str = "asset.instrumentation.attestation.signature";
    pub const INSTRUMENTATION_ATTESTATION_AT: &str = "asset.instrumentation.attestation.at";
    pub const INSTRUMENTATION_TAMPER_DETECTED: &str = "asset.instrumentation.tamper_detected";
    pub const INSTRUMENTATION_TAMPER_INDICATOR: &str = "asset.instrumentation.tamper_indicator";
    pub const INSTRUMENTATION_EXPORTER_CONFIGURED: &str = "asset.instrumentation.exporter.configured";
    pub const INSTRUMENTATION_EXPORTER_REACHABLE: &str = "asset.instrumentation.exporter.reachable";
    pub const INSTRUMENTATION_DROPPED_SPANS: &str = "asset.instrumentation.dropped_spans";

    /// Values for asset.capability.change_type (RFC v0.4 gap closure).
    pub mod capability_change_type {
        pub const ADDED: &str = "added";
        pub const REMOVED: &str = "removed";
        pub const MODIFIED: &str = "modified";
        pub const REPLACED: &str = "replaced";
    }

    /// Values for asset.capability.change_source (RFC v0.4 gap closure).
    pub mod capability_change_source {
        pub const DEPLOYMENT: &str = "deployment";
        pub const CONFIGURATION: &str = "configuration";
        pub const RUNTIME_DISCOVERY: &str = "runtime_discovery";
        pub const SERVER_PUSH: &str = "server_push";
        pub const OPERATOR: &str = "operator";
        pub const UNKNOWN: &str = "unknown";
    }
}

/// AITF identity attributes (subset used by the mapper / crosswalk).
pub mod identity {
    pub const AGENT_ID: &str = "identity.agent_id";
    pub const AGENT_NAME: &str = "identity.agent_name";
    pub const TYPE: &str = "identity.type";
    pub const PROVIDER: &str = "identity.provider";
    pub const CREDENTIAL_TYPE: &str = "identity.credential_type";
    pub const SCOPE: &str = "identity.scope";

    pub const AUTH_METHOD: &str = "identity.auth.method";
    pub const AUTH_RESULT: &str = "identity.auth.result";

    pub const DELEGATION_DELEGATOR: &str = "identity.delegation.delegator";
    pub const DELEGATION_DELEGATOR_ID: &str = "identity.delegation.delegator_id";
    pub const DELEGATION_DELEGATEE: &str = "identity.delegation.delegatee";
    pub const DELEGATION_DELEGATEE_ID: &str = "identity.delegation.delegatee_id";
    pub const DELEGATION_TYPE: &str = "identity.delegation.type";
    pub const DELEGATION_CHAIN: &str = "identity.delegation.chain";
    pub const DELEGATION_SCOPE_DELEGATED: &str = "identity.delegation.scope_delegated";
    pub const DELEGATION_PROOF_TYPE: &str = "identity.delegation.proof_type";
    pub const DELEGATION_TTL_SECONDS: &str = "identity.delegation.ttl_seconds";

    // Trust-domain crossing & delegation depth [RFC v0.4 gap closure]
    pub const BOUNDARY_CROSSED: &str = "identity.boundary.crossed";
    pub const BOUNDARY_SOURCE_DOMAIN: &str = "identity.boundary.source_domain";
    pub const BOUNDARY_TARGET_DOMAIN: &str = "identity.boundary.target_domain";
    pub const BOUNDARY_CROSSING_TYPE: &str = "identity.boundary.crossing_type";
    pub const BOUNDARY_DECISION: &str = "identity.boundary.decision";
    pub const BOUNDARY_DECISION_REASON: &str = "identity.boundary.decision_reason";
    pub const BOUNDARY_POLICY_REF: &str = "identity.boundary.policy_ref";
    pub const BOUNDARY_CROSSINGS_COUNT: &str = "identity.boundary.crossings_count";
    pub const BOUNDARY_DOMAINS_TRAVERSED: &str = "identity.boundary.domains_traversed";
    pub const DELEGATION_DEPTH: &str = "identity.delegation.depth";
    pub const DELEGATION_MAX_DEPTH: &str = "identity.delegation.max_depth";
    pub const DELEGATION_DEPTH_EXCEEDED: &str = "identity.delegation.depth_exceeded";
    pub const DELEGATION_ROOT_PRINCIPAL: &str = "identity.delegation.root_principal";
    pub const DELEGATION_ROOT_PRINCIPAL_TYPE: &str = "identity.delegation.root_principal_type";
    pub const DELEGATION_ROOT_AUTHENTICATED_AT: &str = "identity.delegation.root_authenticated_at";

    // Credential minting & scope-narrowing check [RFC v0.4 gap closure]
    pub const CREDENTIAL_MINT_OPERATION: &str = "identity.credential.mint.operation";
    pub const CREDENTIAL_MINT_GRANT_TYPE: &str = "identity.credential.mint.grant_type";
    pub const CREDENTIAL_MINT_SUBJECT_CLASS: &str = "identity.credential.mint.subject_class";
    pub const CREDENTIAL_MINT_PARENT_ID: &str = "identity.credential.mint.parent_id";
    pub const CREDENTIAL_MINT_CHILD_ID: &str = "identity.credential.mint.child_id";
    pub const CREDENTIAL_MINT_ISSUER: &str = "identity.credential.mint.issuer";
    pub const CREDENTIAL_MINT_SUBJECT: &str = "identity.credential.mint.subject";
    pub const CREDENTIAL_MINT_AUDIENCE: &str = "identity.credential.mint.audience";
    pub const CREDENTIAL_SCOPE_PARENT: &str = "identity.credential.scope.parent";
    pub const CREDENTIAL_SCOPE_CHILD: &str = "identity.credential.scope.child";
    pub const CREDENTIAL_SCOPE_REQUESTED: &str = "identity.credential.scope.requested";
    pub const CREDENTIAL_SCOPE_NARROWED: &str = "identity.credential.scope.narrowed";
    pub const CREDENTIAL_SCOPE_VERIFIED: &str = "identity.credential.scope.verified";
    pub const CREDENTIAL_SCOPE_FORWARDED_UNCHANGED: &str = "identity.credential.scope.forwarded_unchanged";
    pub const CREDENTIAL_SCOPE_ADDED: &str = "identity.credential.scope.added";
    pub const CREDENTIAL_SCOPE_REMOVED: &str = "identity.credential.scope.removed";
    pub const CREDENTIAL_SCOPE_ESCALATION: &str = "identity.credential.scope.escalation";
    pub const CREDENTIAL_MINT_PARENT_TTL_SECONDS: &str = "identity.credential.mint.parent_ttl_seconds";
    pub const CREDENTIAL_MINT_CHILD_TTL_SECONDS: &str = "identity.credential.mint.child_ttl_seconds";
    pub const CREDENTIAL_MINT_RESOURCE_INDICATORS: &str = "identity.credential.mint.resource_indicators";
    pub const CREDENTIAL_MINT_CONSTRAINTS: &str = "identity.credential.mint.constraints";
    pub const CREDENTIAL_MINT_RESULT: &str = "identity.credential.mint.result";
    pub const CREDENTIAL_MINT_DENIAL_REASON: &str = "identity.credential.mint.denial_reason";
    pub const CREDENTIAL_MINT_ON_BEHALF_OF: &str = "identity.credential.mint.on_behalf_of";

    // Human approval / elicitation [RFC v0.4 gap closure]
    pub const APPROVAL_ID: &str = "identity.approval.id";
    pub const APPROVAL_REQUIRED: &str = "identity.approval.required";
    pub const APPROVAL_STATUS: &str = "identity.approval.status";
    pub const APPROVAL_TRIGGER: &str = "identity.approval.trigger";
    pub const APPROVAL_OPERATION: &str = "identity.approval.operation";
    pub const APPROVAL_PROMPT_HASH: &str = "identity.approval.prompt_hash";
    pub const APPROVAL_OPERATION_HASH: &str = "identity.approval.operation_hash";
    pub const APPROVAL_SCOPE_BINDING_RESULT: &str = "identity.approval.scope_binding.result";
    pub const APPROVAL_SCOPE_BINDING_DIVERGENCE: &str = "identity.approval.scope_binding.divergence";
    pub const APPROVAL_DECISION: &str = "identity.approval.decision";
    pub const APPROVAL_APPROVER: &str = "identity.approval.approver";
    pub const APPROVAL_APPROVER_TYPE: &str = "identity.approval.approver_type";
    pub const APPROVAL_APPROVER_VERIFIED: &str = "identity.approval.approver_verified";
    pub const APPROVAL_AUTH_METHOD: &str = "identity.approval.auth_method";
    pub const APPROVAL_REQUESTED_AT: &str = "identity.approval.requested_at";
    pub const APPROVAL_DECIDED_AT: &str = "identity.approval.decided_at";
    pub const APPROVAL_LATENCY_MS: &str = "identity.approval.latency_ms";
    pub const APPROVAL_TIMEOUT_MS: &str = "identity.approval.timeout_ms";
    pub const APPROVAL_TIMEOUT_ACTION: &str = "identity.approval.timeout_action";
    pub const APPROVAL_CHANNEL: &str = "identity.approval.channel";
    pub const APPROVAL_ELICITATION_SCHEMA_HASH: &str = "identity.approval.elicitation.schema_hash";
    pub const APPROVAL_ELICITATION_FIELDS: &str = "identity.approval.elicitation.fields";
    pub const APPROVAL_SCOPE: &str = "identity.approval.scope";
    pub const APPROVAL_REMEMBERED: &str = "identity.approval.remembered";
    pub const APPROVAL_BYPASS_REASON: &str = "identity.approval.bypass_reason";
    pub const APPROVAL_PRIOR_DENIALS: &str = "identity.approval.prior_denials";

    /// Values for identity.approval.decision (RFC v0.4 gap closure).
    pub mod approval_decision {
        pub const APPROVED: &str = "approved";
        pub const DENIED: &str = "denied";
        pub const TIMEOUT: &str = "timeout";
        pub const CANCELLED: &str = "cancelled";
        pub const BYPASSED: &str = "bypassed";
        pub const AUTO_APPROVED: &str = "auto_approved";
    }

    /// Values for identity.approval.scope_binding.result (RFC v0.4 gap closure).
    pub mod approval_scope_binding {
        pub const MATCH: &str = "match";
        pub const MISMATCH: &str = "mismatch";
        pub const NOT_CHECKED: &str = "not_checked";
    }

    /// Values for identity.approval.status (RFC v0.4 gap closure).
    pub mod approval_status {
        pub const PENDING: &str = "pending";
        pub const RESOLVED: &str = "resolved";
        pub const EXPIRED: &str = "expired";
    }

    /// Values for identity.credential.mint.result (RFC v0.4 gap closure).
    pub mod credential_mint_result {
        pub const ISSUED: &str = "issued";
        pub const DENIED: &str = "denied";
        pub const ERROR: &str = "error";
    }
}

/// AITF A2A (Agent-to-Agent Protocol) attributes.
pub mod a2a {
    pub const AGENT_NAME: &str = "a2a.agent.name";
    pub const AGENT_URL: &str = "a2a.agent.url";
    pub const AGENT_VERSION: &str = "a2a.agent.version";
    pub const PROTOCOL_VERSION: &str = "a2a.protocol.version";
    pub const TRANSPORT: &str = "a2a.transport";

    pub const TASK_ID: &str = "a2a.task.id";
    pub const TASK_STATE: &str = "a2a.task.state";
    pub const TASK_PREVIOUS_STATE: &str = "a2a.task.previous_state";
    pub const TASK_ARTIFACTS_COUNT: &str = "a2a.task.artifacts_count";

    pub const MESSAGE_ID: &str = "a2a.message.id";
    pub const MESSAGE_PARTS_COUNT: &str = "a2a.message.parts_count";
    pub const MESSAGE_PART_TYPES: &str = "a2a.message.part_types";

    pub const METHOD: &str = "a2a.method";
    pub const INTERACTION_MODE: &str = "a2a.interaction_mode";
    pub const JSONRPC_ERROR_CODE: &str = "a2a.jsonrpc.error_code";
    pub const JSONRPC_ERROR_MESSAGE: &str = "a2a.jsonrpc.error_message";

    // Task lifecycle event [RFC v0.4 gap closure]
    pub const TASK_LIFECYCLE_EVENT: &str = "a2a.task.lifecycle.event";
    pub const TASK_LIFECYCLE_FROM_STATE: &str = "a2a.task.lifecycle.from_state";
    pub const TASK_LIFECYCLE_TO_STATE: &str = "a2a.task.lifecycle.to_state";
    pub const TASK_LIFECYCLE_TRANSITION_VALID: &str = "a2a.task.lifecycle.transition_valid";
    pub const TASK_LIFECYCLE_ACTOR: &str = "a2a.task.lifecycle.actor";
    pub const TASK_LIFECYCLE_ACTOR_TYPE: &str = "a2a.task.lifecycle.actor_type";
    pub const TASK_LIFECYCLE_AT: &str = "a2a.task.lifecycle.at";
    pub const TASK_LIFECYCLE_AGE_MS: &str = "a2a.task.lifecycle.age_ms";
    pub const TASK_LIFECYCLE_TERMINAL: &str = "a2a.task.lifecycle.terminal";
    pub const TASK_LIFECYCLE_EXPECTED_TERMINAL_BY: &str = "a2a.task.lifecycle.expected_terminal_by";
    pub const TASK_LIFECYCLE_ORPHANED: &str = "a2a.task.lifecycle.orphaned";
    pub const TASK_LIFECYCLE_DELEGATED_SCOPE: &str = "a2a.task.lifecycle.delegated_scope";
    pub const TASK_LIFECYCLE_DELEGATION_EXPIRES_AT: &str = "a2a.task.lifecycle.delegation_expires_at";
    pub const TASK_LIFECYCLE_INITIATING_TRACE_ID: &str = "a2a.task.lifecycle.initiating_trace_id";
    pub const TASK_LIFECYCLE_INITIATING_SPAN_ID: &str = "a2a.task.lifecycle.initiating_span_id";
    pub const TASK_LIFECYCLE_ROOT_PRINCIPAL: &str = "a2a.task.lifecycle.root_principal";
    pub const TASK_LIFECYCLE_FAILURE_REASON: &str = "a2a.task.lifecycle.failure_reason";
    pub const TASK_LIFECYCLE_CANCEL_REQUESTED_BY: &str = "a2a.task.lifecycle.cancel_requested_by";
    pub const TASK_LIFECYCLE_INPUT_REQUIRED_REASON: &str = "a2a.task.lifecycle.input_required_reason";
    pub const TASK_LIFECYCLE_TRANSITION_COUNT: &str = "a2a.task.lifecycle.transition_count";
    pub const TASK_LIFECYCLE_POLL_COUNT: &str = "a2a.task.lifecycle.poll_count";
    pub const TASK_LIFECYCLE_RESUBSCRIBE_COUNT: &str = "a2a.task.lifecycle.resubscribe_count";
    pub const TASK_LIFECYCLE_SUBSCRIBER: &str = "a2a.task.lifecycle.subscriber";

    // Push notification configuration [RFC v0.4 gap closure]
    pub const PUSH_CONFIG_URL: &str = "a2a.push.config.url";
    pub const PUSH_CONFIG_URL_HASH: &str = "a2a.push.config.url_hash";
    pub const PUSH_CONFIG_CHANGED: &str = "a2a.push.config.changed";
    pub const PUSH_CONFIG_AUTHENTICATED: &str = "a2a.push.config.authenticated";
    pub const PUSH_CONFIG_SCHEME: &str = "a2a.push.config.scheme";
    pub const PUSH_CONFIG_IN_ALLOWLIST: &str = "a2a.push.config.in_allowlist";
    pub const PUSH_CONFIG_SET_BY: &str = "a2a.push.config.set_by";

    /// Values for a2a.task.lifecycle.event (RFC v0.4 gap closure).
    pub mod task_lifecycle_event {
        pub const SUBMITTED: &str = "submitted";
        pub const ACCEPTED: &str = "accepted";
        pub const WORKING: &str = "working";
        pub const INPUT_REQUIRED: &str = "input_required";
        pub const AUTH_REQUIRED: &str = "auth_required";
        pub const COMPLETED: &str = "completed";
        pub const FAILED: &str = "failed";
        pub const CANCELED: &str = "canceled";
        pub const REJECTED: &str = "rejected";
        pub const EXPIRED: &str = "expired";
        pub const ORPHANED: &str = "orphaned";
        pub const UNKNOWN: &str = "unknown";
    }
}

/// AITF ACP (Agent Communication Protocol) attributes.
pub mod acp {
    pub const AGENT_NAME: &str = "acp.agent.name";
    pub const RUN_ID: &str = "acp.run.id";
    pub const RUN_MODE: &str = "acp.run.mode";
    pub const RUN_STATUS: &str = "acp.run.status";
    pub const RUN_PREVIOUS_STATUS: &str = "acp.run.previous_status";
    pub const RUN_ERROR_CODE: &str = "acp.run.error.code";
    pub const RUN_ERROR_MESSAGE: &str = "acp.run.error.message";
    pub const RUN_DURATION_MS: &str = "acp.run.duration_ms";

    pub const MESSAGE_PARTS_COUNT: &str = "acp.message.parts_count";
    pub const MESSAGE_CONTENT_TYPES: &str = "acp.message.content_types";

    pub const OPERATION: &str = "acp.operation";
    pub const HTTP_URL: &str = "acp.http.url";
}

/// AITF ANP (Agent Network Protocol) attributes.
pub mod anp {
    pub const PROTOCOL_VERSION: &str = "anp.protocol.version";
    pub const TRANSPORT: &str = "anp.transport";

    pub const DID: &str = "anp.did";
    pub const PEER_DID: &str = "anp.peer.did";

    pub const META_PROTOCOL_NAME: &str = "anp.meta_protocol.name";

    pub const MESSAGE_ID: &str = "anp.message.id";
    pub const MESSAGE_TYPE: &str = "anp.message.type";
    pub const MESSAGE_PARTS_COUNT: &str = "anp.message.parts_count";

    pub const TRUST_DOMAIN: &str = "anp.trust.domain";
    pub const PEER_TRUST_DOMAIN: &str = "anp.trust.peer_domain";
    pub const CROSS_DOMAIN: &str = "anp.trust.cross_domain";

    pub const ERROR_CODE: &str = "anp.error.code";
    pub const ERROR_MESSAGE: &str = "anp.error.message";
}

/// AITF canonical agent-communication attributes — a single, protocol-agnostic
/// namespace that A2A / ACP / ANP (and future protocols) normalize onto.
pub mod agent_comm {
    pub const PROTOCOL: &str = "agent.comm.protocol";
    pub const PROTOCOL_VERSION: &str = "agent.comm.protocol_version";
    pub const DIRECTION: &str = "agent.comm.direction";
    pub const ROLE: &str = "agent.comm.role";
    pub const OPERATION: &str = "agent.comm.operation";
    pub const UNIT_ID: &str = "agent.comm.unit_id";
    pub const UNIT_TYPE: &str = "agent.comm.unit_type";
    pub const STATUS: &str = "agent.comm.status";
    pub const PREVIOUS_STATUS: &str = "agent.comm.previous_status";
    pub const SRC_AGENT_ID: &str = "agent.comm.src_agent_id";
    pub const SRC_AGENT_NAME: &str = "agent.comm.src_agent_name";
    pub const PEER_AGENT_ID: &str = "agent.comm.peer_agent_id";
    pub const PEER_AGENT_NAME: &str = "agent.comm.peer_agent_name";
    pub const PEER_DID: &str = "agent.comm.peer_did";
    pub const PARTS_COUNT: &str = "agent.comm.parts_count";
    pub const PART_TYPES: &str = "agent.comm.part_types";
    pub const ARTIFACTS_COUNT: &str = "agent.comm.artifacts_count";
    pub const TRANSPORT: &str = "agent.comm.transport";
    pub const ENDPOINT: &str = "agent.comm.endpoint";
    pub const PEER_ENDPOINT: &str = "agent.comm.peer_endpoint";
    pub const TRUST_DOMAIN: &str = "agent.comm.trust_domain";
    pub const PEER_TRUST_DOMAIN: &str = "agent.comm.peer_trust_domain";
    pub const CROSS_DOMAIN: &str = "agent.comm.cross_domain";
    pub const ERROR_CODE: &str = "agent.comm.error_code";
    pub const ERROR_MESSAGE: &str = "agent.comm.error_message";
    pub const DURATION_MS: &str = "agent.comm.duration_ms";

    // Canonical lifecycle status values.
    pub const STATUS_SUBMITTED: &str = "submitted";
    pub const STATUS_WORKING: &str = "working";
    pub const STATUS_INPUT_REQUIRED: &str = "input_required";
    pub const STATUS_COMPLETED: &str = "completed";
    pub const STATUS_FAILED: &str = "failed";
    pub const STATUS_CANCELING: &str = "canceling";
    pub const STATUS_CANCELED: &str = "canceled";

    // Canonical protocol values.
    pub const PROTOCOL_A2A: &str = "a2a";
    pub const PROTOCOL_ACP: &str = "acp";
    pub const PROTOCOL_ANP: &str = "anp";
    pub const PROTOCOL_MCP: &str = "mcp";
    pub const PROTOCOL_CUSTOM: &str = "custom";
}

/// AITF metric-name constants (ported from Go `semconv/metrics.go`).
pub mod metrics {
    // OTel GenAI metrics (preserved)
    pub const GEN_AI_TOKEN_USAGE: &str = "gen_ai.client.token.usage";
    pub const GEN_AI_OPERATION_DURATION: &str = "gen_ai.client.operation.duration";

    // Inference metrics
    pub const INFERENCE_REQUESTS: &str = "inference.requests";
    pub const INFERENCE_ERRORS: &str = "inference.errors";
    pub const INFERENCE_TTFT: &str = "inference.time_to_first_token";
    pub const INFERENCE_TPS: &str = "inference.tokens_per_second";

    // Agent metrics
    pub const AGENT_SESSIONS: &str = "agent.sessions";
    pub const AGENT_STEPS: &str = "agent.steps";
    pub const AGENT_SESSION_DURATION: &str = "agent.session.duration";
    pub const AGENT_DELEGATIONS: &str = "agent.delegations";

    // MCP metrics
    pub const MCP_TOOL_INVOCATIONS: &str = "mcp.tool.invocations";
    pub const MCP_TOOL_DURATION: &str = "mcp.tool.duration";
    pub const MCP_SERVER_CONNECTIONS: &str = "mcp.server.connections";
    pub const MCP_TOOL_APPROVALS: &str = "mcp.tool.approvals";

    // Skill metrics
    pub const SKILL_INVOCATIONS: &str = "skill.invocations";
    pub const SKILL_DURATION: &str = "skill.duration";

    // Cost metrics
    pub const COST_TOTAL: &str = "cost.total";
    pub const COST_BUDGET_UTILIZATION: &str = "cost.budget.utilization";

    // Security metrics
    pub const SECURITY_THREATS: &str = "security.threats_detected";
    pub const SECURITY_BLOCKED: &str = "security.requests_blocked";
    pub const SECURITY_PII: &str = "security.pii_detected";
    pub const SECURITY_GUARDRAILS: &str = "security.guardrail.checks";

    // RAG metrics
    pub const RAG_RETRIEVALS: &str = "rag.retrievals";
    pub const RAG_RETRIEVAL_DURATION: &str = "rag.retrieval.duration";

    // Quality metrics
    pub const QUALITY_HALLUCINATION: &str = "quality.hallucination";
    pub const QUALITY_USER_RATING: &str = "quality.user_rating";
}

/// Anthropic Claude Compliance API (Activity Feed) attributes.
pub mod claude_compliance {
    pub const ACTIVITY_ID: &str = "claude.compliance.activity.id";
    pub const ACTIVITY_TYPE: &str = "claude.compliance.activity.type";
    pub const ACTIVITY_CATEGORY: &str = "claude.compliance.activity.category";
    pub const CREATED_AT: &str = "claude.compliance.activity.created_at";
    pub const ORGANIZATION_ID: &str = "claude.compliance.organization.id";
    pub const ORGANIZATION_UUID: &str = "claude.compliance.organization.uuid";

    pub const ACTOR_TYPE: &str = "claude.compliance.actor.type";
    pub const ACTOR_EMAIL: &str = "claude.compliance.actor.email_address";
    pub const ACTOR_USER_ID: &str = "claude.compliance.actor.user_id";
    pub const ACTOR_IP: &str = "claude.compliance.actor.ip_address";
    pub const ACTOR_USER_AGENT: &str = "claude.compliance.actor.user_agent";
    pub const ACTOR_API_KEY_ID: &str = "claude.compliance.actor.api_key_id";
    pub const ACTOR_ADMIN_API_KEY_ID: &str = "claude.compliance.actor.admin_api_key_id";
    pub const ACTOR_DIRECTORY_ID: &str = "claude.compliance.actor.directory_id";

    pub const CHAT_ID: &str = "claude.compliance.chat.id";
    pub const PROJECT_ID: &str = "claude.compliance.project.id";
    pub const FILE_ID: &str = "claude.compliance.file.id";
    pub const FILENAME: &str = "claude.compliance.file.name";
    pub const TARGET_USER_ID: &str = "claude.compliance.target.user_id";
}

/// AITF memory attributes (RFC v0.4 gap closure).
pub mod memory {
    // Declared memory configuration [RFC v0.4 gap closure]
    pub const CONFIG_ENABLED: &str = "memory.config.enabled";
    pub const CONFIG_NAME: &str = "memory.config.name";
    pub const CONFIG_TYPES: &str = "memory.config.types";
    pub const CONFIG_BACKEND: &str = "memory.config.backend";
    pub const CONFIG_SCOPE: &str = "memory.config.scope";
    pub const CONFIG_PERSISTENCE: &str = "memory.config.persistence";
    pub const CONFIG_RETENTION_SECONDS: &str = "memory.config.retention_seconds";
    pub const CONFIG_MAX_ENTRIES: &str = "memory.config.max_entries";
    pub const CONFIG_MAX_BYTES: &str = "memory.config.max_bytes";
    pub const CONFIG_RETRIEVAL_TOP_K: &str = "memory.config.retrieval.top_k";
    pub const CONFIG_RETRIEVAL_SCORING: &str = "memory.config.retrieval.scoring";
    pub const CONFIG_RETRIEVAL_MIN_SCORE: &str = "memory.config.retrieval.min_score";
    pub const CONFIG_WRITE_PRINCIPALS: &str = "memory.config.write_principals";
    pub const CONFIG_AGENT_WRITABLE: &str = "memory.config.agent_writable";
    pub const CONFIG_WRITE_REVIEW: &str = "memory.config.write_review";
    pub const CONFIG_CROSS_SESSION: &str = "memory.config.cross_session";
    pub const CONFIG_CROSS_TENANT: &str = "memory.config.cross_tenant";
    pub const CONFIG_CLASSIFICATION: &str = "memory.config.classification";
    pub const CONFIG_ENCRYPTION_AT_REST: &str = "memory.config.encryption_at_rest";
    pub const CONFIG_PROFILE_REF: &str = "memory.config.profile_ref";
    pub const CONFIG_HASH: &str = "memory.config.hash";

    /// Values for memory.config.scope (RFC v0.4 gap closure).
    pub mod config_scope {
        pub const TURN: &str = "turn";
        pub const SESSION: &str = "session";
        pub const USER: &str = "user";
        pub const AGENT: &str = "agent";
        pub const TENANT: &str = "tenant";
        pub const GLOBAL: &str = "global";
    }

    /// Values for memory.config.write_review (RFC v0.4 gap closure).
    pub mod config_write_review {
        pub const NONE: &str = "none";
        pub const POLICY: &str = "policy";
        pub const GUARDRAIL: &str = "guardrail";
        pub const HUMAN: &str = "human";
    }
}

/// AITF observability meta-telemetry attributes (RFC v0.4 gap closure).
///
/// Attributes describing the telemetry pipeline itself: sequence continuity,
/// tamper evidence, and sampling decisions. Used to establish whether the
/// absence of a signal is evidence of absence or evidence of a blind spot.
pub mod observability {
    // Event sequence continuity [RFC v0.4 gap closure]
    pub const SEQUENCE_NUMBER: &str = "observability.sequence.number";
    pub const SEQUENCE_SCOPE_ID: &str = "observability.sequence.scope_id";
    pub const SEQUENCE_SCOPE_TYPE: &str = "observability.sequence.scope_type";
    pub const SEQUENCE_PREV_HASH: &str = "observability.sequence.prev_hash";
    pub const SEQUENCE_HASH: &str = "observability.sequence.hash";
    pub const SEQUENCE_SIGNATURE: &str = "observability.sequence.signature";
    pub const SEQUENCE_SIGNER_KEY_ID: &str = "observability.sequence.signer_key_id";
    pub const SEQUENCE_GAP_DETECTED: &str = "observability.sequence.gap_detected";
    pub const SEQUENCE_GAP_SIZE: &str = "observability.sequence.gap_size";
    pub const SEQUENCE_REORDERED: &str = "observability.sequence.reordered";
    pub const SAMPLING_DECISION: &str = "observability.sampling.decision";
    pub const SAMPLING_SECURITY_RELEVANT: &str = "observability.sampling.security_relevant";
}
