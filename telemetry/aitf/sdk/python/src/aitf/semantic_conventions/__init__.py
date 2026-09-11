"""AITF Semantic Conventions.

Defines all attribute constants, metric names, and span names used by AITF.
Extends OpenTelemetry GenAI semantic conventions with AITF-specific namespaces.
"""

from aitf.semantic_conventions.attributes import (
    A2AAttributes,
    ACPAttributes,
    AgentAttributes,
    AgentCommAttributes,
    AgenticLogAttributes,
    AIBOMAttributes,
    ANPAttributes,
    AssetInventoryAttributes,
    ClaudeComplianceAttributes,
    ComplianceAttributes,
    CostAttributes,
    DriftDetectionAttributes,
    GenAIAttributes,
    IdentityAttributes,
    LatencyAttributes,
    MCPAttributes,
    MemoryAttributes,
    MemorySecurityAttributes,
    ModelOpsAttributes,
    ObservabilityAttributes,
    QualityAttributes,
    RAGAttributes,
    SecurityAttributes,
    SkillAttributes,
    SupplyChainAttributes,
)
from aitf.semantic_conventions.metrics import AITFMetrics
from aitf.semantic_conventions.resource import AITFResource

__all__ = [
    "A2AAttributes",
    "ACPAttributes",
    "AIBOMAttributes",
    "AITFMetrics",
    "AITFResource",
    "ANPAttributes",
    "AgentAttributes",
    "AgentCommAttributes",
    "AgenticLogAttributes",
    "AssetInventoryAttributes",
    "ClaudeComplianceAttributes",
    "ComplianceAttributes",
    "CostAttributes",
    "DriftDetectionAttributes",
    "GenAIAttributes",
    "IdentityAttributes",
    "LatencyAttributes",
    "MCPAttributes",
    "MemoryAttributes",
    "MemorySecurityAttributes",
    "ModelOpsAttributes",
    "ObservabilityAttributes",
    "QualityAttributes",
    "RAGAttributes",
    "SecurityAttributes",
    "SkillAttributes",
    "SupplyChainAttributes",
]
