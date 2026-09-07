// Package semconv provides AITF semantic convention constants for Go.
//
// All attribute keys used across AITF instrumentation, processors, and exporters.
// OTel GenAI attributes (gen_ai.*) are preserved for compatibility.
// AITF extensions use the aitf.* namespace.
package semconv

import "go.opentelemetry.io/otel/attribute"

// --- OTel GenAI Attributes (Preserved) ---

const (
	GenAISystemKey        = attribute.Key("gen_ai.provider.name")
	GenAIOperationNameKey = attribute.Key("gen_ai.operation.name")

	// Prompt management
	GenAIPromptVersionKey = attribute.Key("gen_ai.prompt.version")
	GenAIPromptLabelKey   = attribute.Key("gen_ai.prompt.label")

	// End-user / tagging (general; e.g. Langfuse userId / tags)
	UserIDKey = attribute.Key("user.id")
	TagsKey   = attribute.Key("tags")

	// Evaluation
	GenAIEvaluationNameKey          = attribute.Key("gen_ai.evaluation.name")
	GenAIEvaluationScoreValueKey    = attribute.Key("gen_ai.evaluation.score.value")
	GenAIEvaluationScoreLabelKey    = attribute.Key("gen_ai.evaluation.score.label")
	GenAIEvaluationExplanationKey   = attribute.Key("gen_ai.evaluation.explanation")
	GenAIEvaluationScoreDataTypeKey = attribute.Key("gen_ai.evaluation.score.data_type")
	GenAIEvaluationSourceKey        = attribute.Key("gen_ai.evaluation.source")
	GenAIEvaluationCommentKey       = attribute.Key("gen_ai.evaluation.comment")
	GenAIEvaluationDatasetItemIDKey = attribute.Key("gen_ai.evaluation.dataset.item_id")

	// Request
	GenAIRequestModelKey            = attribute.Key("gen_ai.request.model")
	GenAIRequestMaxTokensKey        = attribute.Key("gen_ai.request.max_tokens")
	GenAIRequestTemperatureKey      = attribute.Key("gen_ai.request.temperature")
	GenAIRequestTopPKey             = attribute.Key("gen_ai.request.top_p")
	GenAIRequestTopKKey             = attribute.Key("gen_ai.request.top_k")
	GenAIRequestStreamKey           = attribute.Key("gen_ai.request.stream")
	GenAIRequestToolsKey            = attribute.Key("gen_ai.request.tools")
	GenAIRequestToolChoiceKey       = attribute.Key("gen_ai.request.tool_choice")
	GenAIRequestResponseFormatKey   = attribute.Key("gen_ai.request.response_format")
	GenAIRequestFrequencyPenaltyKey = attribute.Key("gen_ai.request.frequency_penalty")
	GenAIRequestPresencePenaltyKey  = attribute.Key("gen_ai.request.presence_penalty")
	GenAIRequestSeedKey             = attribute.Key("gen_ai.request.seed")

	// Response
	GenAIResponseIDKey            = attribute.Key("gen_ai.response.id")
	GenAIResponseModelKey         = attribute.Key("gen_ai.response.model")
	GenAIResponseFinishReasonsKey = attribute.Key("gen_ai.response.finish_reasons")

	// Usage
	GenAIUsageInputTokensKey     = attribute.Key("gen_ai.usage.input_tokens")
	GenAIUsageOutputTokensKey    = attribute.Key("gen_ai.usage.output_tokens")
	GenAIUsageCachedTokensKey    = attribute.Key("gen_ai.usage.cached_tokens")
	GenAIUsageReasoningTokensKey = attribute.Key("gen_ai.usage.reasoning_tokens")

	// System prompt hash (CoSAI WS2: AI_INTERACTION)
	GenAISystemPromptHashKey = attribute.Key("gen_ai.system_prompt.hash")

	// Events
	GenAIPromptKey        = attribute.Key("gen_ai.prompt")
	GenAICompletionKey    = attribute.Key("gen_ai.completion")
	GenAIToolNameKey      = attribute.Key("gen_ai.tool.name")
	GenAIToolCallIDKey    = attribute.Key("gen_ai.tool.call_id")
	GenAIToolArgumentsKey = attribute.Key("gen_ai.tool.arguments")
	GenAIToolResultKey    = attribute.Key("gen_ai.tool.result")
)

// GenAI system values.
const (
	GenAISystemOpenAI    = "openai"
	GenAISystemAnthropic = "anthropic"
	GenAISystemBedrock   = "bedrock"
	GenAISystemAzure     = "azure"
	GenAISystemGCPVertex = "gcp_vertex"
	GenAISystemCohere    = "cohere"
	GenAISystemMistral   = "mistral"
	GenAISystemMeta      = "meta"
	GenAISystemGoogle    = "google"
)

// GenAI operation values.
const (
	GenAIOperationChat           = "chat"
	GenAIOperationTextCompletion = "text_completion"
	GenAIOperationEmbeddings     = "embeddings"
	GenAIOperationImageGen       = "image_generation"
	GenAIOperationAudio          = "audio"
)

// --- AITF Agent Attributes ---

const (
	AgentNameKey        = attribute.Key("gen_ai.agent.name")
	AgentIDKey          = attribute.Key("gen_ai.agent.id")
	AgentTypeKey        = attribute.Key("agent.type")
	AgentFrameworkKey   = attribute.Key("agent.framework")
	AgentVersionKey     = attribute.Key("agent.version")
	AgentDescriptionKey = attribute.Key("gen_ai.agent.description")

	AgentSessionIDKey        = attribute.Key("gen_ai.conversation.id")
	AgentSessionTurnCountKey = attribute.Key("agent.session.turn_count")

	// CoSAI WS2: AGENT_TRACE fields
	AgentWorkflowIDKey = attribute.Key("agent.workflow_id")
	AgentStateKey      = attribute.Key("agent.state")
	AgentScratchpadKey = attribute.Key("agent.scratchpad")
	AgentNextActionKey = attribute.Key("agent.next_action")

	AgentStepTypeKey        = attribute.Key("agent.step.type")
	AgentStepIndexKey       = attribute.Key("agent.step.index")
	AgentStepThoughtKey     = attribute.Key("agent.step.thought")
	AgentStepActionKey      = attribute.Key("agent.step.action")
	AgentStepObservationKey = attribute.Key("agent.step.observation")
	AgentStepStatusKey      = attribute.Key("agent.step.status")

	AgentDelegationTargetAgentKey   = attribute.Key("agent.delegation.target_agent")
	AgentDelegationTargetAgentIDKey = attribute.Key("agent.delegation.target_agent_id")
	AgentDelegationReasonKey        = attribute.Key("agent.delegation.reason")
	AgentDelegationStrategyKey      = attribute.Key("agent.delegation.strategy")
	AgentDelegationTaskKey          = attribute.Key("agent.delegation.task")

	AgentTeamNameKey            = attribute.Key("agent.team.name")
	AgentTeamIDKey              = attribute.Key("agent.team.id")
	AgentTeamTopologyKey        = attribute.Key("agent.team.topology")
	AgentTeamCoordinatorKey     = attribute.Key("agent.team.coordinator")
	AgentTeamConsensusMethodKey = attribute.Key("agent.team.consensus_method")
)

// Agent type values.
const (
	AgentTypeConversational = "conversational"
	AgentTypeAutonomous     = "autonomous"
	AgentTypeReactive       = "reactive"
	AgentTypeProactive      = "proactive"
)

// Agent framework values.
const (
	AgentFrameworkLangChain      = "langchain"
	AgentFrameworkLangGraph      = "langgraph"
	AgentFrameworkCrewAI         = "crewai"
	AgentFrameworkAutoGen        = "autogen"
	AgentFrameworkSemanticKernel = "semantic_kernel"
	AgentFrameworkCustom         = "custom"
)

// Agent step type values.
const (
	AgentStepPlanning      = "planning"
	AgentStepReasoning     = "reasoning"
	AgentStepToolUse       = "tool_use"
	AgentStepDelegation    = "delegation"
	AgentStepResponse      = "response"
	AgentStepReflection    = "reflection"
	AgentStepMemoryAccess  = "memory_access"
	AgentStepGuardrail     = "guardrail_check"
	AgentStepHumanInLoop   = "human_in_loop"
	AgentStepErrorRecovery = "error_recovery"
)

// Agent team topology values.
const (
	TeamTopologyHierarchical = "hierarchical"
	TeamTopologyPeer         = "peer"
	TeamTopologyPipeline     = "pipeline"
	TeamTopologyConsensus    = "consensus"
	TeamTopologyDebate       = "debate"
	TeamTopologySwarm        = "swarm"
)

// --- AITF MCP Attributes ---

const (
	MCPServerNameKey      = attribute.Key("mcp.server.name")
	MCPServerVersionKey   = attribute.Key("mcp.server.version")
	MCPServerTransportKey = attribute.Key("mcp.server.transport")
	MCPServerURLKey       = attribute.Key("mcp.server.url")
	MCPProtocolVersionKey = attribute.Key("mcp.protocol.version")

	MCPToolNameKey             = attribute.Key("gen_ai.tool.name")
	MCPToolServerKey           = attribute.Key("mcp.tool.server")
	MCPToolInputKey            = attribute.Key("gen_ai.tool.call.arguments")
	MCPToolOutputKey           = attribute.Key("gen_ai.tool.call.result")
	MCPToolIsErrorKey          = attribute.Key("mcp.tool.is_error")
	MCPToolDurationMsKey       = attribute.Key("mcp.tool.duration_ms")
	MCPToolApprovalRequiredKey = attribute.Key("mcp.tool.approval_required")
	MCPToolApprovedKey         = attribute.Key("mcp.tool.approved")
	MCPToolCountKey            = attribute.Key("mcp.tool.count")

	// CoSAI WS2: MCP_ACTIVITY fields
	MCPToolResponseErrorKey = attribute.Key("mcp.tool.response_error")
	MCPConnectionIDKey      = attribute.Key("mcp.connection.id")

	MCPResourceURIKey      = attribute.Key("mcp.resource.uri")
	MCPResourceNameKey     = attribute.Key("mcp.resource.name")
	MCPResourceMimeTypeKey = attribute.Key("mcp.resource.mime_type")
	MCPResourceSizeBytesKey = attribute.Key("mcp.resource.size_bytes")

	MCPPromptNameKey      = attribute.Key("mcp.prompt.name")
	MCPPromptArgumentsKey = attribute.Key("mcp.prompt.arguments")

	MCPSamplingModelKey          = attribute.Key("mcp.sampling.model")
	MCPSamplingMaxTokensKey      = attribute.Key("mcp.sampling.max_tokens")
	MCPSamplingIncludeContextKey = attribute.Key("mcp.sampling.include_context")
)

// MCP transport values.
const (
	MCPTransportStdio          = "stdio"
	MCPTransportSSE            = "sse"
	MCPTransportStreamableHTTP = "streamable_http"
)

// --- AITF Skill Attributes ---

const (
	SkillNameKey        = attribute.Key("skill.name")
	SkillIDKey          = attribute.Key("skill.id")
	SkillVersionKey     = attribute.Key("skill.version")
	SkillProviderKey    = attribute.Key("skill.provider")
	SkillCategoryKey    = attribute.Key("skill.category")
	SkillDescriptionKey = attribute.Key("skill.description")
	SkillInputKey       = attribute.Key("skill.input")
	SkillOutputKey      = attribute.Key("skill.output")
	SkillStatusKey      = attribute.Key("skill.status")
	SkillDurationMsKey  = attribute.Key("skill.duration_ms")
	SkillRetryCountKey  = attribute.Key("skill.retry_count")
	SkillSourceKey      = attribute.Key("skill.source")
	SkillHashKey        = attribute.Key("skill.hash")
	SkillAuthorsKey     = attribute.Key("skill.authors")
	SkillCountKey       = attribute.Key("skill.count")

	SkillComposeNameKey    = attribute.Key("skill.compose.name")
	SkillComposePatternKey = attribute.Key("skill.compose.pattern")
	SkillComposeTotalKey   = attribute.Key("skill.compose.total_skills")
)

// Skill status values.
const (
	SkillStatusSuccess = "success"
	SkillStatusError   = "error"
	SkillStatusTimeout = "timeout"
	SkillStatusDenied  = "denied"
	SkillStatusRetry   = "retry"
)

// Skill category values.
const (
	SkillCategorySearch        = "search"
	SkillCategoryCode          = "code"
	SkillCategoryData          = "data"
	SkillCategoryCommunication = "communication"
	SkillCategoryAnalysis      = "analysis"
	SkillCategoryGeneration    = "generation"
	SkillCategoryKnowledge     = "knowledge"
	SkillCategorySecurity      = "security"
	SkillCategoryIntegration   = "integration"
	SkillCategoryWorkflow      = "workflow"
)

// --- AITF RAG Attributes ---

const (
	RAGPipelineNameKey  = attribute.Key("rag.pipeline.name")
	RAGPipelineStageKey = attribute.Key("rag.pipeline.stage")
	RAGQueryKey         = attribute.Key("rag.query")

	RAGRetrieveDatabaseKey     = attribute.Key("gen_ai.data_source.id")
	RAGRetrieveIndexKey        = attribute.Key("rag.retrieve.index")
	RAGRetrieveTopKKey         = attribute.Key("rag.retrieve.top_k")
	RAGRetrieveResultsCountKey = attribute.Key("rag.retrieve.results_count")
	RAGRetrieveMinScoreKey     = attribute.Key("rag.retrieve.min_score")
	RAGRetrieveMaxScoreKey     = attribute.Key("rag.retrieve.max_score")
	RAGRetrieveFilterKey       = attribute.Key("rag.retrieve.filter")

	// CoSAI WS2: RAG_CONTEXT document-level fields
	RAGDocIDKey         = attribute.Key("rag.doc.id")
	RAGDocScoreKey      = attribute.Key("rag.doc.score")
	RAGDocProvenanceKey = attribute.Key("rag.doc.provenance")
	RAGRetrievalDocsKey = attribute.Key("rag.retrieval.docs")

	RAGRerankModelKey       = attribute.Key("rag.rerank.model")
	RAGRerankInputCountKey  = attribute.Key("rag.rerank.input_count")
	RAGRerankOutputCountKey = attribute.Key("rag.rerank.output_count")

	RAGQualityContextRelevanceKey = attribute.Key("rag.quality.context_relevance")
	RAGQualityFaithfulnessKey     = attribute.Key("rag.quality.faithfulness")
	RAGQualityGroundednessKey     = attribute.Key("rag.quality.groundedness")
)

// RAG stage values.
const (
	RAGStageRetrieve = "retrieve"
	RAGStageRerank   = "rerank"
	RAGStageGenerate = "generate"
	RAGStageEvaluate = "evaluate"
)

// --- AITF Security Attributes ---

const (
	SecurityRiskScoreKey       = attribute.Key("security.risk_score")
	SecurityRiskLevelKey       = attribute.Key("security.risk_level")
	SecurityThreatDetectedKey  = attribute.Key("security.threat_detected")
	SecurityThreatTypeKey      = attribute.Key("security.threat_type")
	SecurityOWASPCategoryKey   = attribute.Key("security.owasp_category")
	SecurityBlockedKey         = attribute.Key("security.blocked")
	SecurityDetectionMethodKey = attribute.Key("security.detection_method")
	SecurityConfidenceKey      = attribute.Key("security.confidence")

	SecurityGuardrailNameKey     = attribute.Key("security.guardrail.name")
	SecurityGuardrailTypeKey     = attribute.Key("security.guardrail.type")
	SecurityGuardrailResultKey   = attribute.Key("security.guardrail.result")
	SecurityGuardrailProviderKey = attribute.Key("security.guardrail.provider")

	SecurityPIIDetectedKey = attribute.Key("security.pii.detected")
	SecurityPIICountKey    = attribute.Key("security.pii.count")
	SecurityPIIActionKey   = attribute.Key("security.pii.action")
)

// OWASP LLM Top 10 categories.
const (
	OWASPLICM01 = "LLM01" // Prompt Injection
	OWASPLICM02 = "LLM02" // Sensitive Information Disclosure
	OWASPLICM03 = "LLM03" // Supply Chain
	OWASPLICM04 = "LLM04" // Data and Model Poisoning
	OWASPLICM05 = "LLM05" // Improper Output Handling
	OWASPLICM06 = "LLM06" // Excessive Agency
	OWASPLICM07 = "LLM07" // System Prompt Leakage
	OWASPLICM08 = "LLM08" // Vector and Embedding Weaknesses
	OWASPLICM09 = "LLM09" // Misinformation
	OWASPLICM10 = "LLM10" // Unbounded Consumption
)

// Security risk level values.
const (
	RiskLevelCritical = "critical"
	RiskLevelHigh     = "high"
	RiskLevelMedium   = "medium"
	RiskLevelLow      = "low"
	RiskLevelInfo     = "info"
)

// Threat type values.
const (
	ThreatTypePromptInjection = "prompt_injection"
	ThreatTypeJailbreak       = "jailbreak"
	ThreatTypeDataExfil       = "data_exfiltration"
	ThreatTypeSystemPromptLeak = "system_prompt_leak"
	ThreatTypeModelTheft      = "model_theft"
)

// --- AITF Cost Attributes ---

const (
	CostInputCostKey  = attribute.Key("cost.input_cost")
	CostOutputCostKey = attribute.Key("cost.output_cost")
	CostTotalCostKey  = attribute.Key("cost.total_cost")
	CostCurrencyKey   = attribute.Key("cost.currency")

	CostPricingInputPer1MKey  = attribute.Key("cost.model_pricing.input_per_1m")
	CostPricingOutputPer1MKey = attribute.Key("cost.model_pricing.output_per_1m")

	CostBudgetLimitKey     = attribute.Key("cost.budget.limit")
	CostBudgetUsedKey      = attribute.Key("cost.budget.used")
	CostBudgetRemainingKey = attribute.Key("cost.budget.remaining")

	CostAttributionUserKey    = attribute.Key("cost.attribution.user")
	CostAttributionTeamKey    = attribute.Key("cost.attribution.team")
	CostAttributionProjectKey = attribute.Key("cost.attribution.project")
)

// --- AITF Compliance Attributes ---

const (
	ComplianceNISTControlsKey  = attribute.Key("compliance.nist_ai_rmf.controls")
	ComplianceMITRETechniquesKey = attribute.Key("compliance.mitre_atlas.techniques")
	ComplianceISOControlsKey   = attribute.Key("compliance.iso_42001.controls")
	ComplianceEUArticlesKey    = attribute.Key("compliance.eu_ai_act.articles")
	ComplianceSOC2ControlsKey  = attribute.Key("compliance.soc2.controls")
	ComplianceGDPRArticlesKey  = attribute.Key("compliance.gdpr.articles")
	ComplianceCCPASectionsKey  = attribute.Key("compliance.ccpa.sections")
	ComplianceCSAAICMControlsKey = attribute.Key("compliance.csa_aicm.controls")
)

// --- AITF Latency Attributes ---

const (
	LatencyTotalMsKey            = attribute.Key("latency.total_ms")
	LatencyTimeToFirstTokenMsKey = attribute.Key("latency.time_to_first_token_ms")
	LatencyTokensPerSecondKey    = attribute.Key("latency.tokens_per_second")
	LatencyQueueTimeMsKey        = attribute.Key("latency.queue_time_ms")
	LatencyInferenceTimeMsKey    = attribute.Key("latency.inference_time_ms")
)

// --- AITF Memory Attributes ---

const (
	MemoryOperationKey  = attribute.Key("memory.operation")
	MemoryStoreKey      = attribute.Key("memory.store")
	MemoryKeyKey        = attribute.Key("memory.key")
	MemoryTTLSecondsKey = attribute.Key("memory.ttl_seconds")
	MemoryHitKey        = attribute.Key("memory.hit")
	MemoryProvenanceKey = attribute.Key("memory.provenance")
)

// --- AITF Quality Attributes ---

const (
	QualityHallucinationScoreKey = attribute.Key("quality.hallucination_score")
	QualityConfidenceKey         = attribute.Key("quality.confidence")
	QualityFactualityKey         = attribute.Key("quality.factuality")
	QualityToxicityScoreKey      = attribute.Key("quality.toxicity_score")
	QualityFeedbackRatingKey     = attribute.Key("quality.feedback.rating")
	QualityFeedbackThumbsKey     = attribute.Key("quality.feedback.thumbs")
)

// --- AITF Model Operations Attributes ---

const (
	// Training attributes
	ModelOpsTrainingRunIDKey          = attribute.Key("model_ops.training.run_id")
	ModelOpsTrainingTypeKey           = attribute.Key("model_ops.training.type")
	ModelOpsTrainingBaseModelKey      = attribute.Key("model_ops.training.base_model")
	ModelOpsTrainingFrameworkKey      = attribute.Key("model_ops.training.framework")
	ModelOpsTrainingDatasetIDKey      = attribute.Key("model_ops.training.dataset.id")
	ModelOpsTrainingDatasetVersionKey = attribute.Key("model_ops.training.dataset.version")
	ModelOpsTrainingDatasetSizeKey    = attribute.Key("model_ops.training.dataset.size")
	ModelOpsTrainingHyperparamsKey    = attribute.Key("model_ops.training.hyperparameters")
	ModelOpsTrainingEpochsKey         = attribute.Key("model_ops.training.epochs")
	ModelOpsTrainingLossFinalKey      = attribute.Key("model_ops.training.loss_final")
	ModelOpsTrainingValLossFinalKey   = attribute.Key("model_ops.training.val_loss_final")
	ModelOpsTrainingGPUTypeKey        = attribute.Key("model_ops.training.compute.gpu_type")
	ModelOpsTrainingGPUCountKey       = attribute.Key("model_ops.training.compute.gpu_count")
	ModelOpsTrainingGPUHoursKey       = attribute.Key("model_ops.training.compute.gpu_hours")
	ModelOpsTrainingOutputModelIDKey  = attribute.Key("model_ops.training.output_model.id")
	ModelOpsTrainingOutputModelHashKey = attribute.Key("model_ops.training.output_model.hash")
	ModelOpsTrainingCodeCommitKey     = attribute.Key("model_ops.training.code_commit")
	ModelOpsTrainingExperimentIDKey   = attribute.Key("model_ops.training.experiment.id")
	ModelOpsTrainingExperimentNameKey = attribute.Key("model_ops.training.experiment.name")
	ModelOpsTrainingStatusKey         = attribute.Key("model_ops.training.status")

	// Evaluation attributes
	ModelOpsEvalRunIDKey              = attribute.Key("model_ops.evaluation.run_id")
	ModelOpsEvalModelIDKey            = attribute.Key("model_ops.evaluation.model_id")
	ModelOpsEvalTypeKey               = attribute.Key("model_ops.evaluation.type")
	ModelOpsEvalDatasetIDKey          = attribute.Key("model_ops.evaluation.dataset.id")
	ModelOpsEvalDatasetSizeKey        = attribute.Key("model_ops.evaluation.dataset.size")
	ModelOpsEvalMetricsKey            = attribute.Key("model_ops.evaluation.metrics")
	ModelOpsEvalJudgeModelKey         = attribute.Key("model_ops.evaluation.judge_model")
	ModelOpsEvalBaselineModelKey      = attribute.Key("model_ops.evaluation.baseline_model")
	ModelOpsEvalRegressionDetectedKey = attribute.Key("model_ops.evaluation.regression_detected")
	ModelOpsEvalPassKey               = attribute.Key("model_ops.evaluation.pass")

	// Registry attributes
	ModelOpsRegistryOperationKey        = attribute.Key("model_ops.registry.operation")
	ModelOpsRegistryModelIDKey          = attribute.Key("model_ops.registry.model_id")
	ModelOpsRegistryModelVersionKey     = attribute.Key("model_ops.registry.model_version")
	ModelOpsRegistryModelAliasKey       = attribute.Key("model_ops.registry.model_alias")
	ModelOpsRegistryStageKey            = attribute.Key("model_ops.registry.stage")
	ModelOpsRegistryOwnerKey            = attribute.Key("model_ops.registry.owner")
	ModelOpsRegistryTrainingRunIDKey    = attribute.Key("model_ops.registry.lineage.training_run_id")
	ModelOpsRegistryParentModelIDKey    = attribute.Key("model_ops.registry.lineage.parent_model_id")

	// Deployment attributes
	ModelOpsDeploymentIDKey            = attribute.Key("model_ops.deployment.id")
	ModelOpsDeploymentModelIDKey       = attribute.Key("model_ops.deployment.model_id")
	ModelOpsDeploymentStrategyKey      = attribute.Key("model_ops.deployment.strategy")
	ModelOpsDeploymentEnvironmentKey   = attribute.Key("model_ops.deployment.environment")
	ModelOpsDeploymentEndpointKey      = attribute.Key("model_ops.deployment.endpoint")
	ModelOpsDeploymentInfraProviderKey = attribute.Key("model_ops.deployment.infrastructure.provider")
	ModelOpsDeploymentInfraGPUTypeKey  = attribute.Key("model_ops.deployment.infrastructure.gpu_type")
	ModelOpsDeploymentInfraReplicasKey = attribute.Key("model_ops.deployment.infrastructure.replicas")
	ModelOpsDeploymentCanaryPctKey     = attribute.Key("model_ops.deployment.canary_percent")
	ModelOpsDeploymentStatusKey        = attribute.Key("model_ops.deployment.status")
	ModelOpsDeploymentHealthStatusKey  = attribute.Key("model_ops.deployment.health_check.status")
	ModelOpsDeploymentHealthLatencyKey = attribute.Key("model_ops.deployment.health_check.latency_ms")

	// Serving attributes
	ModelOpsServingOperationKey         = attribute.Key("model_ops.serving.operation")
	ModelOpsServingRouteSelectedKey     = attribute.Key("model_ops.serving.route.selected_model")
	ModelOpsServingRouteReasonKey       = attribute.Key("model_ops.serving.route.reason")
	ModelOpsServingRouteCandidatesKey   = attribute.Key("model_ops.serving.route.candidates")
	ModelOpsServingFallbackChainKey     = attribute.Key("model_ops.serving.fallback.chain")
	ModelOpsServingFallbackDepthKey     = attribute.Key("model_ops.serving.fallback.depth")
	ModelOpsServingFallbackTriggerKey   = attribute.Key("model_ops.serving.fallback.trigger")
	ModelOpsServingFallbackOriginalKey  = attribute.Key("model_ops.serving.fallback.original_model")
	ModelOpsServingFallbackFinalKey     = attribute.Key("model_ops.serving.fallback.final_model")
	ModelOpsServingCacheHitKey          = attribute.Key("model_ops.serving.cache.hit")
	ModelOpsServingCacheTypeKey         = attribute.Key("model_ops.serving.cache.type")
	ModelOpsServingCacheSimilarityKey   = attribute.Key("model_ops.serving.cache.similarity_score")
	ModelOpsServingCacheCostSavedKey    = attribute.Key("model_ops.serving.cache.cost_saved_usd")

	// Monitoring attributes
	ModelOpsMonitorCheckTypeKey      = attribute.Key("model_ops.monitoring.check_type")
	ModelOpsMonitorModelIDKey        = attribute.Key("model_ops.monitoring.model_id")
	ModelOpsMonitorResultKey         = attribute.Key("model_ops.monitoring.result")
	ModelOpsMonitorMetricNameKey     = attribute.Key("model_ops.monitoring.metric_name")
	ModelOpsMonitorMetricValueKey    = attribute.Key("model_ops.monitoring.metric_value")
	ModelOpsMonitorBaselineValueKey  = attribute.Key("model_ops.monitoring.baseline_value")
	ModelOpsMonitorDriftScoreKey     = attribute.Key("model_ops.monitoring.drift_score")
	ModelOpsMonitorDriftTypeKey      = attribute.Key("model_ops.monitoring.drift_type")
	ModelOpsMonitorActionTriggeredKey = attribute.Key("model_ops.monitoring.action_triggered")

	// Prompt lifecycle attributes
	ModelOpsPromptNameKey        = attribute.Key("model_ops.prompt.name")
	ModelOpsPromptOperationKey   = attribute.Key("model_ops.prompt.operation")
	ModelOpsPromptVersionKey     = attribute.Key("model_ops.prompt.version")
	ModelOpsPromptContentHashKey = attribute.Key("model_ops.prompt.content_hash")
	ModelOpsPromptLabelKey       = attribute.Key("model_ops.prompt.label")
	ModelOpsPromptModelTargetKey = attribute.Key("model_ops.prompt.model_target")
	ModelOpsPromptEvalScoreKey   = attribute.Key("model_ops.prompt.evaluation.score")
	ModelOpsPromptEvalPassKey    = attribute.Key("model_ops.prompt.evaluation.pass")
	ModelOpsPromptABTestIDKey    = attribute.Key("model_ops.prompt.a_b_test.id")
	ModelOpsPromptABTestVariantKey = attribute.Key("model_ops.prompt.a_b_test.variant")
)

// --- AITF Drift Detection Attributes ---

const (
	DriftModelIDKey          = attribute.Key("drift.model_id")
	DriftTypeKey             = attribute.Key("drift.type")
	DriftScoreKey            = attribute.Key("drift.score")
	DriftResultKey           = attribute.Key("drift.result")
	DriftDetectionMethodKey  = attribute.Key("drift.detection_method")
	DriftBaselineMetricKey   = attribute.Key("drift.baseline_metric")
	DriftCurrentMetricKey    = attribute.Key("drift.current_metric")
	DriftMetricNameKey       = attribute.Key("drift.metric_name")
	DriftThresholdKey        = attribute.Key("drift.threshold")
	DriftPValueKey           = attribute.Key("drift.p_value")
	DriftRefDatasetKey       = attribute.Key("drift.reference_dataset")
	DriftRefPeriodKey        = attribute.Key("drift.reference_period")
	DriftSampleSizeKey       = attribute.Key("drift.sample_size")
	DriftAffectedSegmentsKey = attribute.Key("drift.affected_segments")
	DriftFeatureNameKey      = attribute.Key("drift.feature_name")
	DriftFeatureImportanceKey = attribute.Key("drift.feature_importance")
	DriftActionTriggeredKey  = attribute.Key("drift.action_triggered")

	// Baseline attributes
	DriftBaselineOperationKey  = attribute.Key("drift.baseline.operation")
	DriftBaselineIDKey         = attribute.Key("drift.baseline.id")
	DriftBaselineDatasetKey    = attribute.Key("drift.baseline.dataset")
	DriftBaselineSampleSizeKey = attribute.Key("drift.baseline.sample_size")
	DriftBaselinePeriodKey     = attribute.Key("drift.baseline.period")
	DriftBaselineMetricsKey    = attribute.Key("drift.baseline.metrics")
	DriftBaselineFeaturesKey   = attribute.Key("drift.baseline.features")
	DriftBaselinePreviousIDKey = attribute.Key("drift.baseline.previous_id")

	// Investigation attributes
	DriftInvestTriggerIDKey       = attribute.Key("drift.investigation.trigger_id")
	DriftInvestRootCauseKey       = attribute.Key("drift.investigation.root_cause")
	DriftInvestRootCauseCatKey    = attribute.Key("drift.investigation.root_cause_category")
	DriftInvestAffectedSegmentsKey = attribute.Key("drift.investigation.affected_segments")
	DriftInvestAffectedUsersKey   = attribute.Key("drift.investigation.affected_users_estimate")
	DriftInvestBlastRadiusKey     = attribute.Key("drift.investigation.blast_radius")
	DriftInvestSeverityKey        = attribute.Key("drift.investigation.severity")
	DriftInvestRecommendationKey  = attribute.Key("drift.investigation.recommendation")

	// Remediation attributes
	DriftRemediationActionKey      = attribute.Key("drift.remediation.action")
	DriftRemediationTriggerIDKey   = attribute.Key("drift.remediation.trigger_id")
	DriftRemediationAutomatedKey   = attribute.Key("drift.remediation.automated")
	DriftRemediationInitiatedByKey = attribute.Key("drift.remediation.initiated_by")
	DriftRemediationStatusKey      = attribute.Key("drift.remediation.status")
	DriftRemediationRollbackToKey  = attribute.Key("drift.remediation.rollback_to")
	DriftRemediationRetrainKey     = attribute.Key("drift.remediation.retrain_dataset")
	DriftRemediationValidPassedKey = attribute.Key("drift.remediation.validation_passed")
)

// --- AITF Identity Attributes ---

const (
	// Core identity attributes
	IdentityAgentIDKey       = attribute.Key("identity.agent_id")
	IdentityAgentNameKey     = attribute.Key("identity.agent_name")
	IdentityTypeKey          = attribute.Key("identity.type")
	IdentityProviderKey      = attribute.Key("identity.provider")
	IdentityOwnerKey         = attribute.Key("identity.owner")
	IdentityOwnerTypeKey     = attribute.Key("identity.owner_type")
	IdentityCredentialTypeKey = attribute.Key("identity.credential_type")
	IdentityCredentialIDKey  = attribute.Key("identity.credential_id")
	IdentityStatusKey        = attribute.Key("identity.status")
	IdentityPreviousStatusKey = attribute.Key("identity.previous_status")
	IdentityScopeKey         = attribute.Key("identity.scope")
	IdentityExpiresAtKey     = attribute.Key("identity.expires_at")
	IdentityTTLSecondsKey    = attribute.Key("identity.ttl_seconds")
	IdentityAutoRotateKey    = attribute.Key("identity.auto_rotate")
	IdentityRotationIntervalKey = attribute.Key("identity.rotation_interval_seconds")

	// Lifecycle
	IdentityLifecycleOpKey = attribute.Key("identity.lifecycle.operation")

	// Authentication attributes
	IdentityAuthMethodKey         = attribute.Key("identity.auth.method")
	IdentityAuthResultKey         = attribute.Key("identity.auth.result")
	IdentityAuthProviderKey       = attribute.Key("identity.auth.provider")
	IdentityAuthTargetServiceKey  = attribute.Key("identity.auth.target_service")
	IdentityAuthFailureReasonKey  = attribute.Key("identity.auth.failure_reason")
	IdentityAuthTokenTypeKey      = attribute.Key("identity.auth.token_type")
	IdentityAuthScopeRequestedKey = attribute.Key("identity.auth.scope_requested")
	IdentityAuthScopeGrantedKey   = attribute.Key("identity.auth.scope_granted")
	IdentityAuthContinuousKey     = attribute.Key("identity.auth.continuous")
	IdentityAuthPKCEUsedKey       = attribute.Key("identity.auth.pkce_used")
	IdentityAuthDPoPUsedKey       = attribute.Key("identity.auth.dpop_used")

	// Authorization attributes
	IdentityAuthzDecisionKey     = attribute.Key("identity.authz.decision")
	IdentityAuthzResourceKey     = attribute.Key("identity.authz.resource")
	IdentityAuthzActionKey       = attribute.Key("identity.authz.action")
	IdentityAuthzPolicyEngineKey = attribute.Key("identity.authz.policy_engine")
	IdentityAuthzPolicyIDKey     = attribute.Key("identity.authz.policy_id")
	IdentityAuthzDenyReasonKey   = attribute.Key("identity.authz.deny_reason")
	IdentityAuthzRiskScoreKey    = attribute.Key("identity.authz.risk_score")
	IdentityAuthzJEAKey          = attribute.Key("identity.authz.jea")
	IdentityAuthzTimeLimitedKey  = attribute.Key("identity.authz.time_limited")
	IdentityAuthzExpiresAtKey    = attribute.Key("identity.authz.expires_at")

	// Delegation attributes
	IdentityDelegDelegatorKey      = attribute.Key("identity.delegation.delegator")
	IdentityDelegDelegatorIDKey    = attribute.Key("identity.delegation.delegator_id")
	IdentityDelegDelegateeKey      = attribute.Key("identity.delegation.delegatee")
	IdentityDelegDelegateeIDKey    = attribute.Key("identity.delegation.delegatee_id")
	IdentityDelegTypeKey           = attribute.Key("identity.delegation.type")
	IdentityDelegChainKey          = attribute.Key("identity.delegation.chain")
	IdentityDelegChainDepthKey     = attribute.Key("identity.delegation.chain_depth")
	IdentityDelegScopeDelegatedKey = attribute.Key("identity.delegation.scope_delegated")
	IdentityDelegScopeAttenuatedKey = attribute.Key("identity.delegation.scope_attenuated")
	IdentityDelegResultKey         = attribute.Key("identity.delegation.result")
	IdentityDelegProofTypeKey      = attribute.Key("identity.delegation.proof_type")
	IdentityDelegTTLSecondsKey     = attribute.Key("identity.delegation.ttl_seconds")

	// Trust attributes
	IdentityTrustOperationKey  = attribute.Key("identity.trust.operation")
	IdentityTrustPeerAgentKey  = attribute.Key("identity.trust.peer_agent")
	IdentityTrustPeerAgentIDKey = attribute.Key("identity.trust.peer_agent_id")
	IdentityTrustResultKey     = attribute.Key("identity.trust.result")
	IdentityTrustMethodKey     = attribute.Key("identity.trust.method")
	IdentityTrustDomainKey     = attribute.Key("identity.trust.trust_domain")
	IdentityTrustPeerDomainKey = attribute.Key("identity.trust.peer_trust_domain")
	IdentityTrustCrossDomainKey = attribute.Key("identity.trust.cross_domain")
	IdentityTrustLevelKey      = attribute.Key("identity.trust.trust_level")
	IdentityTrustProtocolKey   = attribute.Key("identity.trust.protocol")

	// Session attributes
	IdentitySessionIDKey              = attribute.Key("identity.session.id")
	IdentitySessionOperationKey       = attribute.Key("identity.session.operation")
	IdentitySessionScopeKey           = attribute.Key("identity.session.scope")
	IdentitySessionExpiresAtKey       = attribute.Key("identity.session.expires_at")
	IdentitySessionActionsCountKey    = attribute.Key("identity.session.actions_count")
	IdentitySessionTerminationReasonKey = attribute.Key("identity.session.termination_reason")
)

// --- AITF Asset Inventory Attributes ---

const (
	// Core asset attributes
	AssetIDKey                = attribute.Key("asset.id")
	AssetNameKey              = attribute.Key("asset.name")
	AssetTypeKey              = attribute.Key("asset.type")
	AssetVersionKey           = attribute.Key("asset.version")
	AssetHashKey              = attribute.Key("asset.hash")
	AssetOwnerKey             = attribute.Key("asset.owner")
	AssetOwnerTypeKey         = attribute.Key("asset.owner_type")
	AssetDeployEnvKey         = attribute.Key("asset.deployment_environment")
	AssetRiskClassificationKey = attribute.Key("asset.risk_classification")
	AssetSourceRepoKey        = attribute.Key("asset.source_repository")
	AssetTagsKey              = attribute.Key("asset.tags")

	// Discovery attributes
	AssetDiscoveryScopeKey       = attribute.Key("asset.discovery.scope")
	AssetDiscoveryMethodKey      = attribute.Key("asset.discovery.method")
	AssetDiscoveryAssetsFoundKey = attribute.Key("asset.discovery.assets_found")
	AssetDiscoveryNewAssetsKey   = attribute.Key("asset.discovery.new_assets")
	AssetDiscoveryShadowKey      = attribute.Key("asset.discovery.shadow_assets")
	AssetDiscoveryStatusKey      = attribute.Key("asset.discovery.status")

	// Audit attributes
	AssetAuditTypeKey            = attribute.Key("asset.audit.type")
	AssetAuditResultKey          = attribute.Key("asset.audit.result")
	AssetAuditAuditorKey         = attribute.Key("asset.audit.auditor")
	AssetAuditFrameworkKey       = attribute.Key("asset.audit.framework")
	AssetAuditFindingsKey        = attribute.Key("asset.audit.findings")
	AssetAuditNextDueKey         = attribute.Key("asset.audit.next_audit_due")
	AssetAuditRiskScoreKey       = attribute.Key("asset.audit.risk_score")
	AssetAuditIntegrityKey       = attribute.Key("asset.audit.integrity_verified")
	AssetAuditComplianceKey      = attribute.Key("asset.audit.compliance_status")

	// Classification attributes
	AssetClassFrameworkKey        = attribute.Key("asset.classification.framework")
	AssetClassPreviousKey         = attribute.Key("asset.classification.previous")
	AssetClassReasonKey           = attribute.Key("asset.classification.reason")
	AssetClassAssessorKey         = attribute.Key("asset.classification.assessor")
	AssetClassUseCaseKey          = attribute.Key("asset.classification.use_case")
	AssetClassAutonomousDecisionKey = attribute.Key("asset.classification.autonomous_decision")

	// Decommission attributes
	AssetDecommissionReasonKey      = attribute.Key("asset.decommission.reason")
	AssetDecommissionReplacementKey = attribute.Key("asset.decommission.replacement_id")
	AssetDecommissionApprovedByKey  = attribute.Key("asset.decommission.approved_by")
)

// --- AITF Memory Security Attributes ---

const (
	MemorySecurityContentHashKey    = attribute.Key("memory.security.content_hash")
	MemorySecurityContentSizeKey    = attribute.Key("memory.security.content_size")
	MemorySecurityIntegrityHashKey  = attribute.Key("memory.security.integrity_hash")
	MemorySecurityPoisoningScoreKey = attribute.Key("memory.security.poisoning_score")
	MemorySecurityCrossSessionKey   = attribute.Key("memory.security.cross_session")
)

// --- AITF A2A (Agent-to-Agent Protocol) Attributes ---

const (
	A2AAgentNameKey            = attribute.Key("a2a.agent.name")
	A2AAgentURLKey             = attribute.Key("a2a.agent.url")
	A2AAgentVersionKey         = attribute.Key("a2a.agent.version")
	A2AAgentProviderOrgKey     = attribute.Key("a2a.agent.provider.organization")
	A2AAgentSkillsKey          = attribute.Key("a2a.agent.skills")
	A2AAgentCapStreamingKey    = attribute.Key("a2a.agent.capabilities.streaming")
	A2AAgentCapPushKey         = attribute.Key("a2a.agent.capabilities.push_notifications")
	A2AProtocolVersionKey      = attribute.Key("a2a.protocol.version")
	A2ATransportKey            = attribute.Key("a2a.transport")

	A2ATaskIDKey               = attribute.Key("a2a.task.id")
	A2ATaskContextIDKey        = attribute.Key("a2a.task.context_id")
	A2ATaskStateKey            = attribute.Key("a2a.task.state")
	A2ATaskPreviousStateKey    = attribute.Key("a2a.task.previous_state")
	A2ATaskArtifactsCountKey   = attribute.Key("a2a.task.artifacts_count")
	A2ATaskHistoryLengthKey    = attribute.Key("a2a.task.history_length")

	A2AMessageIDKey            = attribute.Key("a2a.message.id")
	A2AMessageRoleKey          = attribute.Key("a2a.message.role")
	A2AMessagePartsCountKey    = attribute.Key("a2a.message.parts_count")
	A2AMessagePartTypesKey     = attribute.Key("a2a.message.part_types")

	A2AMethodKey               = attribute.Key("a2a.method")
	A2AInteractionModeKey      = attribute.Key("a2a.interaction_mode")
	A2AJSONRPCRequestIDKey     = attribute.Key("a2a.jsonrpc.request_id")
	A2AJSONRPCErrorCodeKey     = attribute.Key("a2a.jsonrpc.error_code")
	A2AJSONRPCErrorMessageKey  = attribute.Key("a2a.jsonrpc.error_message")

	A2AArtifactIDKey           = attribute.Key("a2a.artifact.id")
	A2AArtifactNameKey         = attribute.Key("a2a.artifact.name")
	A2AArtifactPartsCountKey   = attribute.Key("a2a.artifact.parts_count")

	A2AStreamEventTypeKey      = attribute.Key("a2a.stream.event_type")
	A2AStreamIsFinalKey        = attribute.Key("a2a.stream.is_final")
	A2AStreamEventsCountKey    = attribute.Key("a2a.stream.events_count")

	A2APushURLKey              = attribute.Key("a2a.push.url")
)

// A2A task state values.
const (
	A2ATaskStateSubmitted     = "submitted"
	A2ATaskStateWorking       = "working"
	A2ATaskStateInputRequired = "input-required"
	A2ATaskStateCompleted     = "completed"
	A2ATaskStateCanceled      = "canceled"
	A2ATaskStateFailed        = "failed"
	A2ATaskStateRejected      = "rejected"
	A2ATaskStateAuthRequired  = "auth-required"
)

// A2A transport values.
const (
	A2ATransportJSONRPC  = "jsonrpc"
	A2ATransportGRPC     = "grpc"
	A2ATransportHTTPJSON = "http_json"
)

// --- AITF ACP (Agent Communication Protocol) Attributes ---

const (
	ACPAgentNameKey              = attribute.Key("acp.agent.name")
	ACPAgentDescriptionKey       = attribute.Key("acp.agent.description")
	ACPAgentInputContentTypesKey = attribute.Key("acp.agent.input_content_types")
	ACPAgentOutputContentTypesKey = attribute.Key("acp.agent.output_content_types")
	ACPAgentFrameworkKey         = attribute.Key("acp.agent.framework")
	ACPAgentSuccessRateKey       = attribute.Key("acp.agent.status.success_rate")
	ACPAgentAvgRunTimeKey        = attribute.Key("acp.agent.status.avg_run_time_seconds")

	ACPRunIDKey                  = attribute.Key("acp.run.id")
	ACPRunAgentNameKey           = attribute.Key("acp.run.agent_name")
	ACPRunSessionIDKey           = attribute.Key("acp.run.session_id")
	ACPRunModeKey                = attribute.Key("acp.run.mode")
	ACPRunStatusKey              = attribute.Key("acp.run.status")
	ACPRunPreviousStatusKey      = attribute.Key("acp.run.previous_status")
	ACPRunErrorCodeKey           = attribute.Key("acp.run.error.code")
	ACPRunErrorMessageKey        = attribute.Key("acp.run.error.message")
	ACPRunCreatedAtKey           = attribute.Key("acp.run.created_at")
	ACPRunFinishedAtKey          = attribute.Key("acp.run.finished_at")
	ACPRunDurationMsKey          = attribute.Key("acp.run.duration_ms")

	ACPMessageRoleKey            = attribute.Key("acp.message.role")
	ACPMessagePartsCountKey      = attribute.Key("acp.message.parts_count")
	ACPMessageContentTypesKey    = attribute.Key("acp.message.content_types")
	ACPMessageHasCitationsKey    = attribute.Key("acp.message.has_citations")
	ACPMessageHasTrajectoryKey   = attribute.Key("acp.message.has_trajectory")

	ACPAwaitActiveKey            = attribute.Key("acp.await.active")
	ACPAwaitCountKey             = attribute.Key("acp.await.count")
	ACPAwaitDurationMsKey        = attribute.Key("acp.await.duration_ms")

	ACPInputMessageCountKey      = attribute.Key("acp.input.message_count")
	ACPOutputMessageCountKey     = attribute.Key("acp.output.message_count")

	ACPOperationKey              = attribute.Key("acp.operation")
	ACPHTTPMethodKey             = attribute.Key("acp.http.method")
	ACPHTTPStatusCodeKey         = attribute.Key("acp.http.status_code")
	ACPHTTPURLKey                = attribute.Key("acp.http.url")

	ACPTrajectoryToolNameKey     = attribute.Key("acp.trajectory.tool_name")
	ACPTrajectoryMessageKey      = attribute.Key("acp.trajectory.message")
)

// ACP run status values.
const (
	ACPRunStatusCreated    = "created"
	ACPRunStatusInProgress = "in-progress"
	ACPRunStatusAwaiting   = "awaiting"
	ACPRunStatusCancelling = "cancelling"
	ACPRunStatusCancelled  = "cancelled"
	ACPRunStatusCompleted  = "completed"
	ACPRunStatusFailed     = "failed"
)

// ACP run mode values.
const (
	ACPRunModeSync  = "sync"
	ACPRunModeAsync = "async"
	ACPRunModeStream = "stream"
)

// --- AITF ANP (Agent Network Protocol) Attributes ---
//
// ANP is a decentralized, DID-based agent-to-agent protocol with meta-protocol
// negotiation and encrypted peer channels. https://agentnetworkprotocol.com
const (
	ANPProtocolVersionKey = attribute.Key("anp.protocol.version")
	ANPTransportKey       = attribute.Key("anp.transport") // http, ws

	// DID identity
	ANPDIDKey     = attribute.Key("anp.did")      // this agent's DID
	ANPPeerDIDKey = attribute.Key("anp.peer.did") // peer agent's DID

	// Meta-protocol negotiation
	ANPMetaProtocolNameKey       = attribute.Key("anp.meta_protocol.name")
	ANPMetaProtocolVersionKey    = attribute.Key("anp.meta_protocol.version")
	ANPMetaProtocolNegotiatedKey = attribute.Key("anp.meta_protocol.negotiated")

	// Message
	ANPMessageIDKey         = attribute.Key("anp.message.id")
	ANPMessageTypeKey       = attribute.Key("anp.message.type")
	ANPMessageRoleKey       = attribute.Key("anp.message.role")
	ANPMessagePartsCountKey = attribute.Key("anp.message.parts_count")

	// Encrypted channel
	ANPEncryptedKey  = attribute.Key("anp.encrypted")
	ANPEncryptionKey = attribute.Key("anp.encryption") // e.g. ecdhe

	// Trust / domains
	ANPTrustDomainKey     = attribute.Key("anp.trust.domain")
	ANPPeerTrustDomainKey = attribute.Key("anp.trust.peer_domain")
	ANPCrossDomainKey     = attribute.Key("anp.trust.cross_domain")

	// Errors
	ANPErrorCodeKey    = attribute.Key("anp.error.code")
	ANPErrorMessageKey = attribute.Key("anp.error.message")
)

// --- AITF Canonical Agent Communication Attributes ---
//
// A single, protocol-agnostic namespace that A2A / ACP / ANP (and future
// agentic protocols) normalize onto. The protocol discriminator carries the
// wire protocol; protocol-specific detail stays in the per-protocol namespaces
// (a2a.* / acp.* / anp.*). This is the AITF analogue of OCSF's "one generic
// class + protocol id" pattern and the source for the OCSF agent_message object.
const (
	AgentCommProtocolKey        = attribute.Key("agent.comm.protocol") // a2a | acp | anp | mcp | custom
	AgentCommProtocolVersionKey = attribute.Key("agent.comm.protocol_version")
	AgentCommDirectionKey       = attribute.Key("agent.comm.direction") // request | response | stream | notification
	AgentCommRoleKey            = attribute.Key("agent.comm.role")      // client | server
	AgentCommOperationKey       = attribute.Key("agent.comm.operation")
	AgentCommUnitIDKey          = attribute.Key("agent.comm.unit_id")   // normalized task/run/message id
	AgentCommUnitTypeKey        = attribute.Key("agent.comm.unit_type") // task | run | message
	AgentCommStatusKey          = attribute.Key("agent.comm.status")    // canonical lifecycle status
	AgentCommPreviousStatusKey  = attribute.Key("agent.comm.previous_status")
	AgentCommSrcAgentIDKey      = attribute.Key("agent.comm.src_agent_id")
	AgentCommSrcAgentNameKey    = attribute.Key("agent.comm.src_agent_name")
	AgentCommPeerAgentIDKey     = attribute.Key("agent.comm.peer_agent_id")
	AgentCommPeerAgentNameKey   = attribute.Key("agent.comm.peer_agent_name")
	AgentCommPeerDIDKey         = attribute.Key("agent.comm.peer_did")
	AgentCommPartsCountKey      = attribute.Key("agent.comm.parts_count")
	AgentCommPartTypesKey       = attribute.Key("agent.comm.part_types")
	AgentCommArtifactsCountKey  = attribute.Key("agent.comm.artifacts_count")
	AgentCommTransportKey       = attribute.Key("agent.comm.transport")
	AgentCommEndpointKey        = attribute.Key("agent.comm.endpoint")
	AgentCommPeerEndpointKey    = attribute.Key("agent.comm.peer_endpoint")
	AgentCommTrustDomainKey     = attribute.Key("agent.comm.trust_domain")
	AgentCommPeerTrustDomainKey = attribute.Key("agent.comm.peer_trust_domain")
	AgentCommCrossDomainKey     = attribute.Key("agent.comm.cross_domain")
	AgentCommErrorCodeKey       = attribute.Key("agent.comm.error_code")
	AgentCommErrorMessageKey    = attribute.Key("agent.comm.error_message")
	AgentCommDurationMsKey      = attribute.Key("agent.comm.duration_ms")
)

// Canonical agent-communication lifecycle status values.
const (
	AgentCommStatusSubmitted     = "submitted"
	AgentCommStatusWorking       = "working"
	AgentCommStatusInputRequired = "input_required"
	AgentCommStatusCompleted     = "completed"
	AgentCommStatusFailed        = "failed"
	AgentCommStatusCanceling     = "canceling"
	AgentCommStatusCanceled      = "canceled"
)

// Canonical agent-communication protocol values.
const (
	AgentCommProtocolA2A    = "a2a"
	AgentCommProtocolACP    = "acp"
	AgentCommProtocolANP    = "anp"
	AgentCommProtocolMCP    = "mcp"
	AgentCommProtocolCustom = "custom"
)

// --- AITF Agentic Log Attributes (Table 10.1 minimal fields) ---

const (
	// EventID: A unique identifier for the specific log entry
	AgenticLogEventIDKey = attribute.Key("agentic_log.event_id")

	// Timestamp: ISO 8601 formatted timestamp with millisecond precision
	AgenticLogTimestampKey = attribute.Key("agentic_log.timestamp")

	// AgentID: The unique, cryptographically verifiable identity of the agent
	AgenticLogAgentIDKey = attribute.Key("agentic_log.agent_id")

	// SessionID: A unique ID for the agent's current operational session
	AgenticLogSessionIDKey = attribute.Key("agentic_log.session_id")

	// GoalID: An identifier for the high-level goal the agent is pursuing
	AgenticLogGoalIDKey = attribute.Key("agentic_log.goal_id")

	// SubTaskID: The specific, immediate task the agent is performing
	AgenticLogSubTaskIDKey = attribute.Key("agentic_log.sub_task_id")

	// ToolUsed: The specific tool, function, or API being invoked
	AgenticLogToolUsedKey = attribute.Key("agentic_log.tool_used")

	// ToolParameters: Sanitized log of parameters (PII/credentials redacted)
	AgenticLogToolParametersKey = attribute.Key("agentic_log.tool_parameters")

	// Outcome: The result of the action (success, failure, error code)
	AgenticLogOutcomeKey = attribute.Key("agentic_log.outcome")

	// ConfidenceScore: Agent's own assessment of how likely the action succeeds
	AgenticLogConfidenceScoreKey = attribute.Key("agentic_log.confidence_score")

	// AnomalyScore: Score indicating how unusual this action is
	AgenticLogAnomalyScoreKey = attribute.Key("agentic_log.anomaly_score")

	// PolicyEvaluation: Record of a check against a security policy engine
	AgenticLogPolicyEvaluationKey = attribute.Key("agentic_log.policy_evaluation")
)

// Agentic log outcome values.
const (
	AgenticLogOutcomeSuccess = "SUCCESS"
	AgenticLogOutcomeFailure = "FAILURE"
	AgenticLogOutcomeError   = "ERROR"
	AgenticLogOutcomeDenied  = "DENIED"
	AgenticLogOutcomeTimeout = "TIMEOUT"
	AgenticLogOutcomePartial = "PARTIAL"
)

// Agentic log policy evaluation result values.
const (
	AgenticLogPolicyPass = "PASS"
	AgenticLogPolicyFail = "FAIL"
	AgenticLogPolicyWarn = "WARN"
	AgenticLogPolicySkip = "SKIP"
)

// --- Anthropic Claude Compliance API (Activity Feed) Attributes ---
//
// Normalizes records from GET /v1/compliance/activities so Claude Enterprise
// audit activity can be carried as AITF telemetry and mapped to OCSF.
// See https://platform.claude.com/docs/en/manage-claude/compliance-api
const (
	// Activity envelope
	ClaudeComplianceActivityID       = "claude.compliance.activity.id"
	ClaudeComplianceActivityType     = "claude.compliance.activity.type"
	ClaudeComplianceActivityCategory = "claude.compliance.activity.category" // derived: auth/account/content/...
	ClaudeComplianceCreatedAt        = "claude.compliance.activity.created_at"
	ClaudeComplianceOrganizationID   = "claude.compliance.organization.id"
	ClaudeComplianceOrganizationUUID = "claude.compliance.organization.uuid"

	// Actor (discriminated union)
	ClaudeComplianceActorType            = "claude.compliance.actor.type" // user_actor, api_actor, admin_api_key_actor, ...
	ClaudeComplianceActorEmail           = "claude.compliance.actor.email_address"
	ClaudeComplianceActorUserID          = "claude.compliance.actor.user_id"
	ClaudeComplianceActorIP              = "claude.compliance.actor.ip_address"
	ClaudeComplianceActorUserAgent       = "claude.compliance.actor.user_agent"
	ClaudeComplianceActorAPIKeyID        = "claude.compliance.actor.api_key_id"
	ClaudeComplianceActorAdminAPIKeyID   = "claude.compliance.actor.admin_api_key_id"
	ClaudeComplianceActorDirectoryID     = "claude.compliance.actor.directory_id"
	ClaudeComplianceActorIdpConnectionType = "claude.compliance.actor.idp_connection_type"

	// Type-specific resource identifiers
	ClaudeComplianceChatID       = "claude.compliance.chat.id"
	ClaudeComplianceProjectID    = "claude.compliance.project.id"
	ClaudeComplianceFileID       = "claude.compliance.file.id"
	ClaudeComplianceFilename     = "claude.compliance.file.name"
	ClaudeComplianceTargetUserID = "claude.compliance.target.user_id"
)

// ===========================================================================
// RFC v0.4 gap closure
//
// Attributes added to close the 27 Appendix C gaps against CoSAI RFC v0.4.
// Mirrors sdk/python/src/aitf/semantic_conventions/attributes.py.
// ===========================================================================

// --- GenAIAttributes [RFC v0.4 gap closure] ---

// Identifier hierarchy [RFC v0.4 gap closure]
const (
	GenAITurnIDKey       = attribute.Key("gen_ai.turn.id")
	GenAITurnIndexKey    = attribute.Key("gen_ai.turn.index")
	GenAITurnParentIDKey = attribute.Key("gen_ai.turn.parent_id")
	GenAIStepIDKey       = attribute.Key("gen_ai.step.id")
	GenAIStepParentIDKey = attribute.Key("gen_ai.step.parent_id")
	GenAIRunIDKey        = attribute.Key("gen_ai.run.id")
)

// Trigger provenance [RFC v0.4 gap closure]
const (
	GenAITriggerTypeKey            = attribute.Key("gen_ai.trigger.type")
	GenAITriggerEventKey           = attribute.Key("gen_ai.trigger.event")
	GenAITriggerEventIDKey         = attribute.Key("gen_ai.trigger.event.id")
	GenAITriggerSourceKey          = attribute.Key("gen_ai.trigger.source")
	GenAITriggerSourcePrincipalKey = attribute.Key("gen_ai.trigger.source.principal")
	GenAITriggerReceivedAtKey      = attribute.Key("gen_ai.trigger.received_at")
	GenAITriggerHumanInLoopKey     = attribute.Key("gen_ai.trigger.human_in_loop")
)

// Content modality & attachment identity [RFC v0.4 gap closure]
const (
	GenAIContentPartIndexKey                   = attribute.Key("gen_ai.content.part.index")
	GenAIContentPartTypeKey                    = attribute.Key("gen_ai.content.part.type")
	GenAIContentPartMimeTypeKey                = attribute.Key("gen_ai.content.part.mime_type")
	GenAIContentPartSizeBytesKey               = attribute.Key("gen_ai.content.part.size_bytes")
	GenAIContentPartHashKey                    = attribute.Key("gen_ai.content.part.hash")
	GenAIContentModalitiesKey                  = attribute.Key("gen_ai.content.modalities")
	GenAIContentAttachmentCountKey             = attribute.Key("gen_ai.content.attachment.count")
	GenAIContentAttachmentNameKey              = attribute.Key("gen_ai.content.attachment.name")
	GenAIContentAttachmentHashKey              = attribute.Key("gen_ai.content.attachment.hash")
	GenAIContentAttachmentSizeBytesKey         = attribute.Key("gen_ai.content.attachment.size_bytes")
	GenAIContentAttachmentSourceKey            = attribute.Key("gen_ai.content.attachment.source")
	GenAIContentAttachmentExtractedTextHashKey = attribute.Key("gen_ai.content.attachment.extracted_text_hash")
	GenAIContentTotalSizeBytesKey              = attribute.Key("gen_ai.content.total_size_bytes")
)

// Backend / route restriction decision [RFC v0.4 gap closure]
const (
	GenAIRouteCandidatesKey        = attribute.Key("gen_ai.route.candidates")
	GenAIRouteCandidatesCountKey   = attribute.Key("gen_ai.route.candidates.count")
	GenAIRouteSelectedKey          = attribute.Key("gen_ai.route.selected")
	GenAIRouteDecisionKey          = attribute.Key("gen_ai.route.decision")
	GenAIRouteConstraintTypeKey    = attribute.Key("gen_ai.route.constraint.type")
	GenAIRouteConstraintValueKey   = attribute.Key("gen_ai.route.constraint.value")
	GenAIRouteConstraintSourceKey  = attribute.Key("gen_ai.route.constraint.source")
	GenAIRouteExcludedKey          = attribute.Key("gen_ai.route.excluded")
	GenAIRouteNoCandidateActionKey = attribute.Key("gen_ai.route.no_candidate_action")
	GenAIRoutePolicyIDKey          = attribute.Key("gen_ai.route.policy_id")
)

// Organization / tenant ID [RFC v0.4 gap closure]
const (
	GenAIConversationTenantIDKey = attribute.Key("gen_ai.conversation.tenant.id")
	GenAIUserTenantIDKey         = attribute.Key("user.tenant.id")
)

// Values for gen_ai.trigger.type (RFC v0.4 gap closure).
const (
	GenAITriggertypeUserInitiated  = "user_initiated"
	GenAITriggertypeScheduled      = "scheduled"
	GenAITriggertypeWebhook        = "webhook"
	GenAITriggertypeEventDriven    = "event_driven"
	GenAITriggertypeAgentInitiated = "agent_initiated"
	GenAITriggertypeSystem         = "system"
	GenAITriggertypeRetry          = "retry"
	GenAITriggertypeUnknown        = "unknown"
)

// Values for gen_ai.route.decision (RFC v0.4 gap closure).
const (
	GenAIRoutedecisionAllowed    = "allowed"
	GenAIRoutedecisionRestricted = "restricted"
	GenAIRoutedecisionDenied     = "denied"
	GenAIRoutedecisionFallback   = "fallback"
	GenAIRoutedecisionNoPolicy   = "no_policy"
)

// --- AgentAttributes [RFC v0.4 gap closure] ---

// Peer agent card / descriptor [RFC v0.4 gap closure]
const (
	AgentPeerIDKey                    = attribute.Key("gen_ai.agent.peer.id")
	AgentPeerNameKey                  = attribute.Key("gen_ai.agent.peer.name")
	AgentPeerURLKey                   = attribute.Key("gen_ai.agent.peer.url")
	AgentPeerVersionKey               = attribute.Key("gen_ai.agent.peer.version")
	AgentPeerProviderKey              = attribute.Key("gen_ai.agent.peer.provider")
	AgentPeerSkillsKey                = attribute.Key("gen_ai.agent.peer.skills")
	AgentPeerProtocolKey              = attribute.Key("gen_ai.agent.peer.protocol")
	AgentPeerCardHashKey              = attribute.Key("gen_ai.agent.peer.card.hash")
	AgentPeerCardBaselineHashKey      = attribute.Key("gen_ai.agent.peer.card.baseline_hash")
	AgentPeerCardChangedKey           = attribute.Key("gen_ai.agent.peer.card.changed")
	AgentPeerCardChangeFieldsKey      = attribute.Key("gen_ai.agent.peer.card.change_fields")
	AgentPeerCardFirstSeenKey         = attribute.Key("gen_ai.agent.peer.card.first_seen")
	AgentPeerVerificationMethodKey    = attribute.Key("gen_ai.agent.peer.verification.method")
	AgentPeerVerificationResultKey    = attribute.Key("gen_ai.agent.peer.verification.result")
	AgentPeerVerificationAuthorityKey = attribute.Key("gen_ai.agent.peer.verification.authority")
	AgentPeerApprovedKey              = attribute.Key("gen_ai.agent.peer.approved")
)

// Organization / tenant ID [RFC v0.4 gap closure]
const (
	AgentTenantIDKey = attribute.Key("gen_ai.agent.tenant.id")
)

// Values for gen_ai.agent.peer.verification.result (RFC v0.4 gap closure).
const (
	AgentPeerverificationresultVerified     = "verified"
	AgentPeerverificationresultUnverified   = "unverified"
	AgentPeerverificationresultFailed       = "failed"
	AgentPeerverificationresultRefused      = "refused"
	AgentPeerverificationresultNotAttempted = "not_attempted"
)

// --- MCPAttributes [RFC v0.4 gap closure] ---

// Tool definition digest [RFC v0.4 gap closure]
const (
	MCPToolDefinitionHashKey            = attribute.Key("mcp.tool.definition.hash")
	MCPToolDefinitionBaselineHashKey    = attribute.Key("mcp.tool.definition.baseline_hash")
	MCPToolDefinitionChangedKey         = attribute.Key("mcp.tool.definition.changed")
	MCPToolDefinitionChangeTypeKey      = attribute.Key("mcp.tool.definition.change_type")
	MCPToolDefinitionDescriptionHashKey = attribute.Key("mcp.tool.definition.description_hash")
	MCPToolDefinitionSchemaHashKey      = attribute.Key("mcp.tool.definition.schema_hash")
	MCPToolDefinitionApprovedKey        = attribute.Key("mcp.tool.definition.approved")
	MCPToolDefinitionApprovedAtKey      = attribute.Key("mcp.tool.definition.approved_at")
	MCPToolDefinitionFirstSeenKey       = attribute.Key("mcp.tool.definition.first_seen")
	MCPToolDefinitionSourceKey          = attribute.Key("mcp.tool.definition.source")
)

// Server identity & primitive [RFC v0.4 gap closure]
const (
	MCPPrimitiveKey              = attribute.Key("mcp.primitive")
	MCPMethodNameKey             = attribute.Key("mcp.method.name")
	MCPRequestIDKey              = attribute.Key("mcp.request.id")
	MCPSessionIDKey              = attribute.Key("mcp.session.id")
	MCPServerInstanceIDKey       = attribute.Key("mcp.server.instance.id")
	MCPServerIdentityMethodKey   = attribute.Key("mcp.server.identity.method")
	MCPServerIdentityVerifiedKey = attribute.Key("mcp.server.identity.verified")
	MCPServerTrustDomainKey      = attribute.Key("mcp.server.trust_domain")
	MCPServerCommandKey          = attribute.Key("mcp.server.command")
	MCPServerBinaryHashKey       = attribute.Key("mcp.server.binary.hash")
	MCPCapabilitiesNegotiatedKey = attribute.Key("mcp.capabilities.negotiated")
)

// Protocol envelope capture [RFC v0.4 gap closure]
const (
	MCPEnvelopeCapturedKey       = attribute.Key("mcp.envelope.captured")
	MCPEnvelopeRequestKey        = attribute.Key("mcp.envelope.request")
	MCPEnvelopeResponseKey       = attribute.Key("mcp.envelope.response")
	MCPEnvelopeRequestHashKey    = attribute.Key("mcp.envelope.request.hash")
	MCPEnvelopeResponseHashKey   = attribute.Key("mcp.envelope.response.hash")
	MCPEnvelopeSizeBytesKey      = attribute.Key("mcp.envelope.size_bytes")
	MCPEnvelopeTruncatedKey      = attribute.Key("mcp.envelope.truncated")
	MCPEnvelopeRedactedKey       = attribute.Key("mcp.envelope.redacted")
	MCPEnvelopeJsonrpcVersionKey = attribute.Key("mcp.envelope.jsonrpc.version")
	MCPEnvelopeErrorCodeKey      = attribute.Key("mcp.envelope.error.code")
	MCPEnvelopeErrorMessageKey   = attribute.Key("mcp.envelope.error.message")
)

// Values for mcp.primitive (RFC v0.4 gap closure).
const (
	MCPPrimitiveTool        = "tool"
	MCPPrimitiveResource    = "resource"
	MCPPrimitivePrompt      = "prompt"
	MCPPrimitiveSampling    = "sampling"
	MCPPrimitiveCompletion  = "completion"
	MCPPrimitiveElicitation = "elicitation"
	MCPPrimitiveRoot        = "root"
	MCPPrimitiveLogging     = "logging"
)

// --- RAGAttributes [RFC v0.4 gap closure] ---

// Citations / source attribution [RFC v0.4 gap closure]
const (
	RAGCitationCountKey               = attribute.Key("rag.citation.count")
	RAGCitationIDSKey                 = attribute.Key("rag.citation.ids")
	RAGCitationSourcesKey             = attribute.Key("rag.citation.sources")
	RAGCitationResolvedCountKey       = attribute.Key("rag.citation.resolved_count")
	RAGCitationUnresolvedCountKey     = attribute.Key("rag.citation.unresolved_count")
	RAGCitationUnresolvedIDSKey       = attribute.Key("rag.citation.unresolved_ids")
	RAGCitationFabricatedKey          = attribute.Key("rag.citation.fabricated")
	RAGCitationRetrievalSpanIDKey     = attribute.Key("rag.citation.retrieval_span_id")
	RAGCitationCoverageRatioKey       = attribute.Key("rag.citation.coverage_ratio")
	RAGCitationUncitedContentRatioKey = attribute.Key("rag.citation.uncited_content_ratio")
	RAGCitationFormatKey              = attribute.Key("rag.citation.format")
	RAGCitationVerifiedKey            = attribute.Key("rag.citation.verified")
)

// Declared knowledge-source configuration [RFC v0.4 gap closure]
const (
	RAGSourceDeclaredKey         = attribute.Key("rag.source.declared")
	RAGSourceDeclaredCountKey    = attribute.Key("rag.source.declared_count")
	RAGSourceNameKey             = attribute.Key("rag.source.name")
	RAGSourceDescriptionKey      = attribute.Key("rag.source.description")
	RAGSourceTypeKey             = attribute.Key("rag.source.type")
	RAGSourceIndexNameKey        = attribute.Key("rag.source.index.name")
	RAGSourceIndexNamespaceKey   = attribute.Key("rag.source.index.namespace")
	RAGSourceSchemaKey           = attribute.Key("rag.source.schema")
	RAGSourceSchemaHashKey       = attribute.Key("rag.source.schema.hash")
	RAGSourceTrustLevelKey       = attribute.Key("rag.source.trust_level")
	RAGSourceClassificationKey   = attribute.Key("rag.source.classification")
	RAGSourceOwnerKey            = attribute.Key("rag.source.owner")
	RAGSourceWriteAccessKey      = attribute.Key("rag.source.write_access")
	RAGSourceIngestionMethodKey  = attribute.Key("rag.source.ingestion.method")
	RAGSourceLastIndexedKey      = attribute.Key("rag.source.last_indexed")
	RAGSourceIndexHashKey        = attribute.Key("rag.source.index.hash")
	RAGSourceSearchTopKKey       = attribute.Key("rag.source.search.top_k")
	RAGSourceSearchFiltersKey    = attribute.Key("rag.source.search.filters")
	RAGSourceSearchScoringKey    = attribute.Key("rag.source.search.scoring")
	RAGSourceSearchMinScoreKey   = attribute.Key("rag.source.search.min_score")
	RAGSourceSearchRerankerKey   = attribute.Key("rag.source.search.reranker")
	RAGSourceUndeclaredAccessKey = attribute.Key("rag.source.undeclared_access")
)

// Organization / tenant ID [RFC v0.4 gap closure]
const (
	RAGDataSourceTenantIDKey = attribute.Key("gen_ai.data_source.tenant.id")
)

// --- SecurityAttributes [RFC v0.4 gap closure] ---

// Guardrail modification record [RFC v0.4 gap closure]
const (
	SecurityGuardrailActionKey                       = attribute.Key("security.guardrail.action")
	SecurityGuardrailModifiedKey                     = attribute.Key("security.guardrail.modified")
	SecurityGuardrailModificationTypeKey             = attribute.Key("security.guardrail.modification.type")
	SecurityGuardrailModificationSideKey             = attribute.Key("security.guardrail.modification.side")
	SecurityGuardrailModificationEnforcementPointKey = attribute.Key("security.guardrail.modification.enforcement_point")
	SecurityGuardrailModificationHashBeforeKey       = attribute.Key("security.guardrail.modification.hash_before")
	SecurityGuardrailModificationHashAfterKey        = attribute.Key("security.guardrail.modification.hash_after")
	SecurityGuardrailModificationCountKey            = attribute.Key("security.guardrail.modification.count")
	SecurityGuardrailModificationBytesRemovedKey     = attribute.Key("security.guardrail.modification.bytes_removed")
	SecurityGuardrailModificationRedactionMapKey     = attribute.Key("security.guardrail.modification.redaction_map")
	SecurityGuardrailModificationReasonKey           = attribute.Key("security.guardrail.modification.reason")
)

// Encoded / obfuscated payload indicator [RFC v0.4 gap closure]
const (
	SecurityObfuscationDetectedKey             = attribute.Key("security.obfuscation.detected")
	SecurityObfuscationEncodingsKey            = attribute.Key("security.obfuscation.encodings")
	SecurityObfuscationDepthKey                = attribute.Key("security.obfuscation.depth")
	SecurityObfuscationDecodedFormKey          = attribute.Key("security.obfuscation.decoded_form")
	SecurityObfuscationDecodedHashKey          = attribute.Key("security.obfuscation.decoded_hash")
	SecurityObfuscationDecodedLengthKey        = attribute.Key("security.obfuscation.decoded_length")
	SecurityObfuscationSourceFieldKey          = attribute.Key("security.obfuscation.source_field")
	SecurityObfuscationInspectedAfterDecodeKey = attribute.Key("security.obfuscation.inspected_after_decode")
)

// Authorization decision record [RFC v0.4 gap closure]
const (
	SecurityAuthorizationDecisionKey             = attribute.Key("security.authorization.decision")
	SecurityAuthorizationDecisionIDKey           = attribute.Key("security.authorization.decision_id")
	SecurityAuthorizationOperationKey            = attribute.Key("security.authorization.operation")
	SecurityAuthorizationReasonKey               = attribute.Key("security.authorization.reason")
	SecurityAuthorizationReasonCodeKey           = attribute.Key("security.authorization.reason_code")
	SecurityAuthorizationAuthorityTypeKey        = attribute.Key("security.authorization.authority.type")
	SecurityAuthorizationAuthorityIDKey          = attribute.Key("security.authorization.authority.id")
	SecurityAuthorizationAuthorityEngineKey      = attribute.Key("security.authorization.authority.engine")
	SecurityAuthorizationRuleIDKey               = attribute.Key("security.authorization.rule_id")
	SecurityAuthorizationPolicyVersionKey        = attribute.Key("security.authorization.policy_version")
	SecurityAuthorizationObligationsKey          = attribute.Key("security.authorization.obligations")
	SecurityAuthorizationObligationsFulfilledKey = attribute.Key("security.authorization.obligations_fulfilled")
	SecurityAuthorizationPrincipalKey            = attribute.Key("security.authorization.principal")
	SecurityAuthorizationResourceKey             = attribute.Key("security.authorization.resource")
	SecurityAuthorizationLatencyMsKey            = attribute.Key("security.authorization.latency_ms")
)

// Session taint labels [RFC v0.4 gap closure]
const (
	SecurityTaintLabelsKey                 = attribute.Key("security.taint.labels")
	SecurityTaintScopeKey                  = attribute.Key("security.taint.scope")
	SecurityTaintLabelCountKey             = attribute.Key("security.taint.label_count")
	SecurityTaintAppliedByKey              = attribute.Key("security.taint.applied_by")
	SecurityTaintOriginKey                 = attribute.Key("security.taint.origin")
	SecurityTaintOriginSpanIDKey           = attribute.Key("security.taint.origin_span_id")
	SecurityTaintFlowDecisionKey           = attribute.Key("security.taint.flow.decision")
	SecurityTaintDeniedKey                 = attribute.Key("security.taint.denied")
	SecurityTaintDeniedLabelsKey           = attribute.Key("security.taint.denied_labels")
	SecurityTaintDeclassifiedByKey         = attribute.Key("security.taint.declassified_by")
	SecurityTaintDeclassificationReasonKey = attribute.Key("security.taint.declassification_reason")
)

// Enforcement-point availability [RFC v0.4 gap closure]
const (
	SecurityEnforcementPointNameKey     = attribute.Key("security.enforcement.point.name")
	SecurityEnforcementPointTypeKey     = attribute.Key("security.enforcement.point.type")
	SecurityEnforcementPointVersionKey  = attribute.Key("security.enforcement.point.version")
	SecurityEnforcementReachedKey       = attribute.Key("security.enforcement.reached")
	SecurityEnforcementLatencyMsKey     = attribute.Key("security.enforcement.latency_ms")
	SecurityEnforcementTimeoutMsKey     = attribute.Key("security.enforcement.timeout_ms")
	SecurityEnforcementFailureModeKey   = attribute.Key("security.enforcement.failure_mode")
	SecurityEnforcementFailureReasonKey = attribute.Key("security.enforcement.failure_reason")
	SecurityEnforcementActionTakenKey   = attribute.Key("security.enforcement.action_taken")
	SecurityEnforcementDegradedModeKey  = attribute.Key("security.enforcement.degraded_mode")
)

// Attribute source / trusted-provenance marking [RFC v0.4 gap closure]
const (
	SecurityAttributeSourceDefaultKey          = attribute.Key("security.attribute_source.default")
	SecurityAttributeSourceMapKey              = attribute.Key("security.attribute_source.map")
	SecurityAttributeSourceSelfAssertedKey     = attribute.Key("security.attribute_source.self_asserted")
	SecurityAttributeSourceVerifiedKey         = attribute.Key("security.attribute_source.verified")
	SecurityAttributeSourceAuthorityIDKey      = attribute.Key("security.attribute_source.authority.id")
	SecurityAttributeSourceVerificationTimeKey = attribute.Key("security.attribute_source.verification_time")
)

// Mediation coverage & bypass path [RFC v0.4 gap closure]
const (
	SecurityMediationMediatedKey           = attribute.Key("security.mediation.mediated")
	SecurityMediationPlacementKey          = attribute.Key("security.mediation.placement")
	SecurityMediationReferenceMonitorIDKey = attribute.Key("security.mediation.reference_monitor.id")
	SecurityMediationBypassAvailableKey    = attribute.Key("security.mediation.bypass_available")
	SecurityMediationBypassPathsKey        = attribute.Key("security.mediation.bypass_paths")
	SecurityMediationCoverageRatioKey      = attribute.Key("security.mediation.coverage_ratio")
	SecurityMediationCapabilityKey         = attribute.Key("security.mediation.capability")
	SecurityMediationAssessedAtKey         = attribute.Key("security.mediation.assessed_at")
)

// Tenant-crossing detection [RFC v0.4 gap closure]
const (
	SecurityTenantCrossingDetectedKey = attribute.Key("security.tenant.crossing_detected")
	SecurityTenantCrossingTypeKey     = attribute.Key("security.tenant.crossing_type")
	SecurityTenantExpectedKey         = attribute.Key("security.tenant.expected")
)

// Values for security.authorization.decision (RFC v0.4 gap closure).
const (
	SecurityAuthorizationdecisionAllow         = "allow"
	SecurityAuthorizationdecisionDeny          = "deny"
	SecurityAuthorizationdecisionChallenge     = "challenge"
	SecurityAuthorizationdecisionNotApplicable = "not_applicable"
	SecurityAuthorizationdecisionError         = "error"
)

// Values for security.taint.flow.decision (RFC v0.4 gap closure).
const (
	SecurityTaintflowdecisionAllowed      = "allowed"
	SecurityTaintflowdecisionDenied       = "denied"
	SecurityTaintflowdecisionDeclassified = "declassified"
	SecurityTaintflowdecisionNotEvaluated = "not_evaluated"
)

// Values for security.taint.scope (RFC v0.4 gap closure).
const (
	SecurityTaintscopeSession = "session"
	SecurityTaintscopeTurn    = "turn"
	SecurityTaintscopeMessage = "message"
	SecurityTaintscopeRun     = "run"
)

// Values for security.enforcement.failure_mode (RFC v0.4 gap closure).
const (
	SecurityEnforcementfailuremodeNone        = "none"
	SecurityEnforcementfailuremodeFailOpen    = "fail_open"
	SecurityEnforcementfailuremodeFailClosed  = "fail_closed"
	SecurityEnforcementfailuremodeTimeout     = "timeout"
	SecurityEnforcementfailuremodeUnreachable = "unreachable"
	SecurityEnforcementfailuremodeDegraded    = "degraded"
)

// Values for security.attribute_source.* (RFC v0.4 gap closure).
const (
	SecurityAttributesourceSelfAsserted = "self_asserted"
	SecurityAttributesourceVerified     = "verified"
	SecurityAttributesourceDerived      = "derived"
	SecurityAttributesourceUnknown      = "unknown"
)

// --- ComplianceAttributes [RFC v0.4 gap closure] ---

// Authorization decision record [RFC v0.4 gap closure]
const (
	ComplianceFrameworkKey = attribute.Key("compliance.framework")
	ComplianceControlIDKey = attribute.Key("compliance.control_id")
)

// --- SupplyChainAttributes [RFC v0.4 gap closure] ---

// Execution environment / sandbox [RFC v0.4 gap closure]
const (
	SupplyChainRuntimeSandboxModeKey              = attribute.Key("supply_chain.runtime.sandbox.mode")
	SupplyChainRuntimeSandboxProviderKey          = attribute.Key("supply_chain.runtime.sandbox.provider")
	SupplyChainRuntimeLanguageNameKey             = attribute.Key("supply_chain.runtime.language.name")
	SupplyChainRuntimeLanguageVersionKey          = attribute.Key("supply_chain.runtime.language.version")
	SupplyChainRuntimeOSTypeKey                   = attribute.Key("supply_chain.runtime.os.type")
	SupplyChainRuntimeOSVersionKey                = attribute.Key("supply_chain.runtime.os.version")
	SupplyChainRuntimeArchitectureKey             = attribute.Key("supply_chain.runtime.architecture")
	SupplyChainRuntimeInstanceIDKey               = attribute.Key("supply_chain.runtime.instance.id")
	SupplyChainRuntimeImageDigestKey              = attribute.Key("supply_chain.runtime.image.digest")
	SupplyChainRuntimeImageRefKey                 = attribute.Key("supply_chain.runtime.image.ref")
	SupplyChainRuntimePrivilegedKey               = attribute.Key("supply_chain.runtime.privileged")
	SupplyChainRuntimeUserKey                     = attribute.Key("supply_chain.runtime.user")
	SupplyChainRuntimeCapabilitiesKey             = attribute.Key("supply_chain.runtime.capabilities")
	SupplyChainRuntimeNetworkEgressPolicyKey      = attribute.Key("supply_chain.runtime.network.egress_policy")
	SupplyChainRuntimeNetworkEgressAllowlistKey   = attribute.Key("supply_chain.runtime.network.egress_allowlist")
	SupplyChainRuntimeNetworkNamespaceKey         = attribute.Key("supply_chain.runtime.network.namespace")
	SupplyChainRuntimeFilesystemModeKey           = attribute.Key("supply_chain.runtime.filesystem.mode")
	SupplyChainRuntimeFilesystemMountsKey         = attribute.Key("supply_chain.runtime.filesystem.mounts")
	SupplyChainRuntimeResourceCPULimitKey         = attribute.Key("supply_chain.runtime.resource.cpu_limit")
	SupplyChainRuntimeResourceMemoryLimitBytesKey = attribute.Key("supply_chain.runtime.resource.memory_limit_bytes")
	SupplyChainRuntimeResourceTimeoutMsKey        = attribute.Key("supply_chain.runtime.resource.timeout_ms")
	SupplyChainRuntimeSecretsExposedKey           = attribute.Key("supply_chain.runtime.secrets.exposed")
	SupplyChainRuntimeSecretsCountKey             = attribute.Key("supply_chain.runtime.secrets.count")
	SupplyChainRuntimeAttestationMethodKey        = attribute.Key("supply_chain.runtime.attestation.method")
	SupplyChainRuntimeAttestationVerifiedKey      = attribute.Key("supply_chain.runtime.attestation.verified")
	SupplyChainRuntimeEscapeDetectedKey           = attribute.Key("supply_chain.runtime.escape_detected")
	SupplyChainRuntimeEscapeIndicatorKey          = attribute.Key("supply_chain.runtime.escape_indicator")
)

// --- MemoryAttributes [RFC v0.4 gap closure] ---

// Declared memory configuration [RFC v0.4 gap closure]
const (
	MemoryConfigEnabledKey           = attribute.Key("memory.config.enabled")
	MemoryConfigNameKey              = attribute.Key("memory.config.name")
	MemoryConfigTypesKey             = attribute.Key("memory.config.types")
	MemoryConfigBackendKey           = attribute.Key("memory.config.backend")
	MemoryConfigScopeKey             = attribute.Key("memory.config.scope")
	MemoryConfigPersistenceKey       = attribute.Key("memory.config.persistence")
	MemoryConfigRetentionSecondsKey  = attribute.Key("memory.config.retention_seconds")
	MemoryConfigMaxEntriesKey        = attribute.Key("memory.config.max_entries")
	MemoryConfigMaxBytesKey          = attribute.Key("memory.config.max_bytes")
	MemoryConfigRetrievalTopKKey     = attribute.Key("memory.config.retrieval.top_k")
	MemoryConfigRetrievalScoringKey  = attribute.Key("memory.config.retrieval.scoring")
	MemoryConfigRetrievalMinScoreKey = attribute.Key("memory.config.retrieval.min_score")
	MemoryConfigWritePrincipalsKey   = attribute.Key("memory.config.write_principals")
	MemoryConfigAgentWritableKey     = attribute.Key("memory.config.agent_writable")
	MemoryConfigWriteReviewKey       = attribute.Key("memory.config.write_review")
	MemoryConfigCrossSessionKey      = attribute.Key("memory.config.cross_session")
	MemoryConfigCrossTenantKey       = attribute.Key("memory.config.cross_tenant")
	MemoryConfigClassificationKey    = attribute.Key("memory.config.classification")
	MemoryConfigEncryptionAtRestKey  = attribute.Key("memory.config.encryption_at_rest")
	MemoryConfigProfileRefKey        = attribute.Key("memory.config.profile_ref")
	MemoryConfigHashKey              = attribute.Key("memory.config.hash")
)

// Values for memory.config.scope (RFC v0.4 gap closure).
const (
	MemoryConfigscopeTurn    = "turn"
	MemoryConfigscopeSession = "session"
	MemoryConfigscopeUser    = "user"
	MemoryConfigscopeAgent   = "agent"
	MemoryConfigscopeTenant  = "tenant"
	MemoryConfigscopeGlobal  = "global"
)

// Values for memory.config.write_review (RFC v0.4 gap closure).
const (
	MemoryConfigwritereviewNone      = "none"
	MemoryConfigwritereviewPolicy    = "policy"
	MemoryConfigwritereviewGuardrail = "guardrail"
	MemoryConfigwritereviewHuman     = "human"
)

// --- IdentityAttributes [RFC v0.4 gap closure] ---

// Trust-domain crossing & delegation depth [RFC v0.4 gap closure]
const (
	IdentityBoundaryCrossedKey               = attribute.Key("identity.boundary.crossed")
	IdentityBoundarySourceDomainKey          = attribute.Key("identity.boundary.source_domain")
	IdentityBoundaryTargetDomainKey          = attribute.Key("identity.boundary.target_domain")
	IdentityBoundaryCrossingTypeKey          = attribute.Key("identity.boundary.crossing_type")
	IdentityBoundaryDecisionKey              = attribute.Key("identity.boundary.decision")
	IdentityBoundaryDecisionReasonKey        = attribute.Key("identity.boundary.decision_reason")
	IdentityBoundaryPolicyRefKey             = attribute.Key("identity.boundary.policy_ref")
	IdentityBoundaryCrossingsCountKey        = attribute.Key("identity.boundary.crossings_count")
	IdentityBoundaryDomainsTraversedKey      = attribute.Key("identity.boundary.domains_traversed")
	IdentityDelegationDepthKey               = attribute.Key("identity.delegation.depth")
	IdentityDelegationMaxDepthKey            = attribute.Key("identity.delegation.max_depth")
	IdentityDelegationDepthExceededKey       = attribute.Key("identity.delegation.depth_exceeded")
	IdentityDelegationRootPrincipalKey       = attribute.Key("identity.delegation.root_principal")
	IdentityDelegationRootPrincipalTypeKey   = attribute.Key("identity.delegation.root_principal_type")
	IdentityDelegationRootAuthenticatedAtKey = attribute.Key("identity.delegation.root_authenticated_at")
)

// Credential minting & scope-narrowing check [RFC v0.4 gap closure]
const (
	IdentityCredentialMintOperationKey           = attribute.Key("identity.credential.mint.operation")
	IdentityCredentialMintGrantTypeKey           = attribute.Key("identity.credential.mint.grant_type")
	IdentityCredentialMintSubjectClassKey        = attribute.Key("identity.credential.mint.subject_class")
	IdentityCredentialMintParentIDKey            = attribute.Key("identity.credential.mint.parent_id")
	IdentityCredentialMintChildIDKey             = attribute.Key("identity.credential.mint.child_id")
	IdentityCredentialMintIssuerKey              = attribute.Key("identity.credential.mint.issuer")
	IdentityCredentialMintSubjectKey             = attribute.Key("identity.credential.mint.subject")
	IdentityCredentialMintAudienceKey            = attribute.Key("identity.credential.mint.audience")
	IdentityCredentialScopeParentKey             = attribute.Key("identity.credential.scope.parent")
	IdentityCredentialScopeChildKey              = attribute.Key("identity.credential.scope.child")
	IdentityCredentialScopeRequestedKey          = attribute.Key("identity.credential.scope.requested")
	IdentityCredentialScopeNarrowedKey           = attribute.Key("identity.credential.scope.narrowed")
	IdentityCredentialScopeVerifiedKey           = attribute.Key("identity.credential.scope.verified")
	IdentityCredentialScopeForwardedUnchangedKey = attribute.Key("identity.credential.scope.forwarded_unchanged")
	IdentityCredentialScopeAddedKey              = attribute.Key("identity.credential.scope.added")
	IdentityCredentialScopeRemovedKey            = attribute.Key("identity.credential.scope.removed")
	IdentityCredentialScopeEscalationKey         = attribute.Key("identity.credential.scope.escalation")
	IdentityCredentialMintParentTTLSecondsKey    = attribute.Key("identity.credential.mint.parent_ttl_seconds")
	IdentityCredentialMintChildTTLSecondsKey     = attribute.Key("identity.credential.mint.child_ttl_seconds")
	IdentityCredentialMintResourceIndicatorsKey  = attribute.Key("identity.credential.mint.resource_indicators")
	IdentityCredentialMintConstraintsKey         = attribute.Key("identity.credential.mint.constraints")
	IdentityCredentialMintResultKey              = attribute.Key("identity.credential.mint.result")
	IdentityCredentialMintDenialReasonKey        = attribute.Key("identity.credential.mint.denial_reason")
	IdentityCredentialMintOnBehalfOfKey          = attribute.Key("identity.credential.mint.on_behalf_of")
)

// Human approval / elicitation [RFC v0.4 gap closure]
const (
	IdentityApprovalIDKey                     = attribute.Key("identity.approval.id")
	IdentityApprovalRequiredKey               = attribute.Key("identity.approval.required")
	IdentityApprovalStatusKey                 = attribute.Key("identity.approval.status")
	IdentityApprovalTriggerKey                = attribute.Key("identity.approval.trigger")
	IdentityApprovalOperationKey              = attribute.Key("identity.approval.operation")
	IdentityApprovalPromptHashKey             = attribute.Key("identity.approval.prompt_hash")
	IdentityApprovalOperationHashKey          = attribute.Key("identity.approval.operation_hash")
	IdentityApprovalScopeBindingResultKey     = attribute.Key("identity.approval.scope_binding.result")
	IdentityApprovalScopeBindingDivergenceKey = attribute.Key("identity.approval.scope_binding.divergence")
	IdentityApprovalDecisionKey               = attribute.Key("identity.approval.decision")
	IdentityApprovalApproverKey               = attribute.Key("identity.approval.approver")
	IdentityApprovalApproverTypeKey           = attribute.Key("identity.approval.approver_type")
	IdentityApprovalApproverVerifiedKey       = attribute.Key("identity.approval.approver_verified")
	IdentityApprovalAuthMethodKey             = attribute.Key("identity.approval.auth_method")
	IdentityApprovalRequestedAtKey            = attribute.Key("identity.approval.requested_at")
	IdentityApprovalDecidedAtKey              = attribute.Key("identity.approval.decided_at")
	IdentityApprovalLatencyMsKey              = attribute.Key("identity.approval.latency_ms")
	IdentityApprovalTimeoutMsKey              = attribute.Key("identity.approval.timeout_ms")
	IdentityApprovalTimeoutActionKey          = attribute.Key("identity.approval.timeout_action")
	IdentityApprovalChannelKey                = attribute.Key("identity.approval.channel")
	IdentityApprovalElicitationSchemaHashKey  = attribute.Key("identity.approval.elicitation.schema_hash")
	IdentityApprovalElicitationFieldsKey      = attribute.Key("identity.approval.elicitation.fields")
	IdentityApprovalScopeKey                  = attribute.Key("identity.approval.scope")
	IdentityApprovalRememberedKey             = attribute.Key("identity.approval.remembered")
	IdentityApprovalBypassReasonKey           = attribute.Key("identity.approval.bypass_reason")
	IdentityApprovalPriorDenialsKey           = attribute.Key("identity.approval.prior_denials")
)

// Values for identity.approval.decision (RFC v0.4 gap closure).
const (
	IdentityApprovaldecisionApproved     = "approved"
	IdentityApprovaldecisionDenied       = "denied"
	IdentityApprovaldecisionTimeout      = "timeout"
	IdentityApprovaldecisionCancelled    = "cancelled"
	IdentityApprovaldecisionBypassed     = "bypassed"
	IdentityApprovaldecisionAutoApproved = "auto_approved"
)

// Values for identity.approval.scope_binding.result (RFC v0.4 gap closure).
const (
	IdentityApprovalscopebindingMatch      = "match"
	IdentityApprovalscopebindingMismatch   = "mismatch"
	IdentityApprovalscopebindingNotChecked = "not_checked"
)

// Values for identity.approval.status (RFC v0.4 gap closure).
const (
	IdentityApprovalstatusPending  = "pending"
	IdentityApprovalstatusResolved = "resolved"
	IdentityApprovalstatusExpired  = "expired"
)

// Values for identity.credential.mint.result (RFC v0.4 gap closure).
const (
	IdentityCredentialmintresultIssued = "issued"
	IdentityCredentialmintresultDenied = "denied"
	IdentityCredentialmintresultError  = "error"
)

// --- AssetInventoryAttributes [RFC v0.4 gap closure] ---

// Organization / tenant ID [RFC v0.4 gap closure]
const (
	AssetInventoryTenantIDKey                = attribute.Key("asset.tenant.id")
	AssetInventoryOrganizationIDKey          = attribute.Key("asset.organization.id")
	AssetInventoryOrganizationNameKey        = attribute.Key("asset.organization.name")
	AssetInventoryTenantTierKey              = attribute.Key("asset.tenant.tier")
	AssetInventoryTenantIsolationBoundaryKey = attribute.Key("asset.tenant.isolation_boundary")
	AssetInventoryTenantDataResidencyKey     = attribute.Key("asset.tenant.data_residency")
)

// Capability-set change event [RFC v0.4 gap closure]
const (
	AssetInventoryCapabilityChangeTypeKey      = attribute.Key("asset.capability.change_type")
	AssetInventoryCapabilitySetHashKey         = attribute.Key("asset.capability.set.hash")
	AssetInventoryCapabilitySetPreviousHashKey = attribute.Key("asset.capability.set.previous_hash")
	AssetInventoryCapabilitySetSizeKey         = attribute.Key("asset.capability.set.size")
	AssetInventoryCapabilityAddedKey           = attribute.Key("asset.capability.added")
	AssetInventoryCapabilityRemovedKey         = attribute.Key("asset.capability.removed")
	AssetInventoryCapabilityModifiedKey        = attribute.Key("asset.capability.modified")
	AssetInventoryCapabilityCategoryKey        = attribute.Key("asset.capability.category")
	AssetInventoryCapabilityRiskDeltaKey       = attribute.Key("asset.capability.risk_delta")
	AssetInventoryCapabilityPrivilegedAddedKey = attribute.Key("asset.capability.privileged_added")
	AssetInventoryCapabilityChangeSourceKey    = attribute.Key("asset.capability.change_source")
	AssetInventoryCapabilityChangeActorKey     = attribute.Key("asset.capability.change_actor")
	AssetInventoryCapabilityApprovedKey        = attribute.Key("asset.capability.approved")
	AssetInventoryCapabilityApprovalRefKey     = attribute.Key("asset.capability.approval_ref")
	AssetInventoryCapabilityDetectedAtKey      = attribute.Key("asset.capability.detected_at")
	AssetInventoryCapabilityDetectionMethodKey = attribute.Key("asset.capability.detection_method")
)

// Instrumentation coverage / hook attestation [RFC v0.4 gap closure]
const (
	AssetInventoryInstrumentationEnabledKey              = attribute.Key("asset.instrumentation.enabled")
	AssetInventoryInstrumentationVersionKey              = attribute.Key("asset.instrumentation.version")
	AssetInventoryInstrumentationSdkKey                  = attribute.Key("asset.instrumentation.sdk")
	AssetInventoryInstrumentationHooksDeclaredKey        = attribute.Key("asset.instrumentation.hooks.declared")
	AssetInventoryInstrumentationHooksActiveKey          = attribute.Key("asset.instrumentation.hooks.active")
	AssetInventoryInstrumentationHooksMissingKey         = attribute.Key("asset.instrumentation.hooks.missing")
	AssetInventoryInstrumentationCoverageRatioKey        = attribute.Key("asset.instrumentation.coverage_ratio")
	AssetInventoryInstrumentationUninstrumentedPathsKey  = attribute.Key("asset.instrumentation.uninstrumented_paths")
	AssetInventoryInstrumentationAttestationMethodKey    = attribute.Key("asset.instrumentation.attestation.method")
	AssetInventoryInstrumentationAttestationVerifiedKey  = attribute.Key("asset.instrumentation.attestation.verified")
	AssetInventoryInstrumentationAttestationSignatureKey = attribute.Key("asset.instrumentation.attestation.signature")
	AssetInventoryInstrumentationAttestationAtKey        = attribute.Key("asset.instrumentation.attestation.at")
	AssetInventoryInstrumentationTamperDetectedKey       = attribute.Key("asset.instrumentation.tamper_detected")
	AssetInventoryInstrumentationTamperIndicatorKey      = attribute.Key("asset.instrumentation.tamper_indicator")
	AssetInventoryInstrumentationExporterConfiguredKey   = attribute.Key("asset.instrumentation.exporter.configured")
	AssetInventoryInstrumentationExporterReachableKey    = attribute.Key("asset.instrumentation.exporter.reachable")
	AssetInventoryInstrumentationDroppedSpansKey         = attribute.Key("asset.instrumentation.dropped_spans")
)

// Values for asset.capability.change_type (RFC v0.4 gap closure).
const (
	AssetInventoryCapabilitychangetypeAdded    = "added"
	AssetInventoryCapabilitychangetypeRemoved  = "removed"
	AssetInventoryCapabilitychangetypeModified = "modified"
	AssetInventoryCapabilitychangetypeReplaced = "replaced"
)

// Values for asset.capability.change_source (RFC v0.4 gap closure).
const (
	AssetInventoryCapabilitychangesourceDeployment       = "deployment"
	AssetInventoryCapabilitychangesourceConfiguration    = "configuration"
	AssetInventoryCapabilitychangesourceRuntimeDiscovery = "runtime_discovery"
	AssetInventoryCapabilitychangesourceServerPush       = "server_push"
	AssetInventoryCapabilitychangesourceOperator         = "operator"
	AssetInventoryCapabilitychangesourceUnknown          = "unknown"
)

// --- A2AAttributes [RFC v0.4 gap closure] ---

// Task lifecycle event [RFC v0.4 gap closure]
const (
	A2ATaskLifecycleEventKey               = attribute.Key("a2a.task.lifecycle.event")
	A2ATaskLifecycleFromStateKey           = attribute.Key("a2a.task.lifecycle.from_state")
	A2ATaskLifecycleToStateKey             = attribute.Key("a2a.task.lifecycle.to_state")
	A2ATaskLifecycleTransitionValidKey     = attribute.Key("a2a.task.lifecycle.transition_valid")
	A2ATaskLifecycleActorKey               = attribute.Key("a2a.task.lifecycle.actor")
	A2ATaskLifecycleActorTypeKey           = attribute.Key("a2a.task.lifecycle.actor_type")
	A2ATaskLifecycleAtKey                  = attribute.Key("a2a.task.lifecycle.at")
	A2ATaskLifecycleAgeMsKey               = attribute.Key("a2a.task.lifecycle.age_ms")
	A2ATaskLifecycleTerminalKey            = attribute.Key("a2a.task.lifecycle.terminal")
	A2ATaskLifecycleExpectedTerminalByKey  = attribute.Key("a2a.task.lifecycle.expected_terminal_by")
	A2ATaskLifecycleOrphanedKey            = attribute.Key("a2a.task.lifecycle.orphaned")
	A2ATaskLifecycleDelegatedScopeKey      = attribute.Key("a2a.task.lifecycle.delegated_scope")
	A2ATaskLifecycleDelegationExpiresAtKey = attribute.Key("a2a.task.lifecycle.delegation_expires_at")
	A2ATaskLifecycleInitiatingTraceIDKey   = attribute.Key("a2a.task.lifecycle.initiating_trace_id")
	A2ATaskLifecycleInitiatingSpanIDKey    = attribute.Key("a2a.task.lifecycle.initiating_span_id")
	A2ATaskLifecycleRootPrincipalKey       = attribute.Key("a2a.task.lifecycle.root_principal")
	A2ATaskLifecycleFailureReasonKey       = attribute.Key("a2a.task.lifecycle.failure_reason")
	A2ATaskLifecycleCancelRequestedByKey   = attribute.Key("a2a.task.lifecycle.cancel_requested_by")
	A2ATaskLifecycleInputRequiredReasonKey = attribute.Key("a2a.task.lifecycle.input_required_reason")
	A2ATaskLifecycleTransitionCountKey     = attribute.Key("a2a.task.lifecycle.transition_count")
	A2ATaskLifecyclePollCountKey           = attribute.Key("a2a.task.lifecycle.poll_count")
	A2ATaskLifecycleResubscribeCountKey    = attribute.Key("a2a.task.lifecycle.resubscribe_count")
	A2ATaskLifecycleSubscriberKey          = attribute.Key("a2a.task.lifecycle.subscriber")
)

// Push notification configuration [RFC v0.4 gap closure]
const (
	A2APushConfigURLKey           = attribute.Key("a2a.push.config.url")
	A2APushConfigURLHashKey       = attribute.Key("a2a.push.config.url_hash")
	A2APushConfigChangedKey       = attribute.Key("a2a.push.config.changed")
	A2APushConfigAuthenticatedKey = attribute.Key("a2a.push.config.authenticated")
	A2APushConfigSchemeKey        = attribute.Key("a2a.push.config.scheme")
	A2APushConfigInAllowlistKey   = attribute.Key("a2a.push.config.in_allowlist")
	A2APushConfigSetByKey         = attribute.Key("a2a.push.config.set_by")
)

// Values for a2a.task.lifecycle.event (RFC v0.4 gap closure).
const (
	A2ATasklifecycleeventSubmitted     = "submitted"
	A2ATasklifecycleeventAccepted      = "accepted"
	A2ATasklifecycleeventWorking       = "working"
	A2ATasklifecycleeventInputRequired = "input_required"
	A2ATasklifecycleeventAuthRequired  = "auth_required"
	A2ATasklifecycleeventCompleted     = "completed"
	A2ATasklifecycleeventFailed        = "failed"
	A2ATasklifecycleeventCanceled      = "canceled"
	A2ATasklifecycleeventRejected      = "rejected"
	A2ATasklifecycleeventExpired       = "expired"
	A2ATasklifecycleeventOrphaned      = "orphaned"
	A2ATasklifecycleeventUnknown       = "unknown"
)

// --- ObservabilityAttributes [RFC v0.4 gap closure] ---

// Event sequence continuity [RFC v0.4 gap closure]
const (
	ObservabilitySequenceNumberKey           = attribute.Key("observability.sequence.number")
	ObservabilitySequenceScopeIDKey          = attribute.Key("observability.sequence.scope_id")
	ObservabilitySequenceScopeTypeKey        = attribute.Key("observability.sequence.scope_type")
	ObservabilitySequencePrevHashKey         = attribute.Key("observability.sequence.prev_hash")
	ObservabilitySequenceHashKey             = attribute.Key("observability.sequence.hash")
	ObservabilitySequenceSignatureKey        = attribute.Key("observability.sequence.signature")
	ObservabilitySequenceSignerKeyIDKey      = attribute.Key("observability.sequence.signer_key_id")
	ObservabilitySequenceGapDetectedKey      = attribute.Key("observability.sequence.gap_detected")
	ObservabilitySequenceGapSizeKey          = attribute.Key("observability.sequence.gap_size")
	ObservabilitySequenceReorderedKey        = attribute.Key("observability.sequence.reordered")
	ObservabilitySamplingDecisionKey         = attribute.Key("observability.sampling.decision")
	ObservabilitySamplingSecurityRelevantKey = attribute.Key("observability.sampling.security_relevant")
)
