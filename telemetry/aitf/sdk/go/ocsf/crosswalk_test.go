package ocsf

import (
	"testing"

	"github.com/girdav01/AITF/sdk/go/semconv"
)

func TestNormalizeAgentTypeID(t *testing.T) {
	cases := []struct {
		framework string
		want      int
	}{
		{"langchain", AgentTypeIDLangChain},
		{"langgraph", AgentTypeIDLangChain},
		{"crewai", AgentTypeIDCrewAI},
		{"autogen", AgentTypeIDAutoGen},
		{"native", AgentTypeIDNative},
		{"semantic_kernel", AgentTypeIDOther},
		{"custom", AgentTypeIDOther},
		{"", AgentTypeIDUnknown},
	}
	for _, c := range cases {
		if got := NormalizeAgentTypeID(c.framework); got != c.want {
			t.Errorf("NormalizeAgentTypeID(%q) = %d, want %d", c.framework, got, c.want)
		}
	}
}

func TestBuildAIAgentCrewAI(t *testing.T) {
	attrs := map[string]interface{}{
		string(semconv.AgentIDKey):        "agent-123",
		string(semconv.AgentNameKey):      "researcher",
		string(semconv.AgentFrameworkKey): "crewai",
		string(semconv.GenAIRequestModelKey): "gpt-4o",
	}
	agent := BuildAIAgent(attrs)
	if agent == nil {
		t.Fatal("BuildAIAgent returned nil")
	}
	if agent.UID != "agent-123" {
		t.Errorf("UID = %q, want agent-123", agent.UID)
	}
	if agent.TypeID != AgentTypeIDCrewAI {
		t.Errorf("TypeID = %d, want %d", agent.TypeID, AgentTypeIDCrewAI)
	}
	if agent.Type != "CrewAI" {
		t.Errorf("Type = %q, want CrewAI", agent.Type)
	}
	if agent.AIModel != "gpt-4o" {
		t.Errorf("AIModel = %q, want gpt-4o", agent.AIModel)
	}
}

func TestBuildAIAgentNilWithoutIdentity(t *testing.T) {
	if agent := BuildAIAgent(map[string]interface{}{}); agent != nil {
		t.Errorf("BuildAIAgent on empty attrs = %+v, want nil", agent)
	}
}

func TestBuildDelegation(t *testing.T) {
	attrs := map[string]interface{}{
		string(semconv.IdentityDelegDelegateeIDKey): "delegatee-id",
		string(semconv.IdentityDelegDelegatorIDKey): "delegator-id",
		string(semconv.IdentityDelegTypeKey):        "on_behalf_of",
		string(semconv.IdentityProviderKey):         "issuer-x",
		string(semconv.IdentityDelegScopeDelegatedKey): []string{"read", "write"},
		string(semconv.IdentityDelegTTLSecondsKey):  int64(3600),
	}
	d := BuildDelegation(attrs)
	if d == nil {
		t.Fatal("BuildDelegation returned nil")
	}
	if d.UID != "delegatee-id" {
		t.Errorf("UID = %q, want delegatee-id", d.UID)
	}
	if d.ParentUID != "delegator-id" {
		t.Errorf("ParentUID = %q, want delegator-id", d.ParentUID)
	}
	if d.IssuerUID != "issuer-x" {
		t.Errorf("IssuerUID = %q, want issuer-x", d.IssuerUID)
	}
	if d.Type != "on_behalf_of" {
		t.Errorf("Type = %q, want on_behalf_of", d.Type)
	}
	if len(d.Scope) != 2 || d.Scope[0] != "read" || d.Scope[1] != "write" {
		t.Errorf("Scope = %v, want [read write]", d.Scope)
	}
	if d.TTLSeconds == nil || *d.TTLSeconds != 3600 {
		t.Errorf("TTLSeconds = %v, want 3600", d.TTLSeconds)
	}
}

func TestBuildDelegationNilWithoutContext(t *testing.T) {
	if d := BuildDelegation(map[string]interface{}{}); d != nil {
		t.Errorf("BuildDelegation on empty attrs = %+v, want nil", d)
	}
}

func TestBuildDelegationLineage(t *testing.T) {
	attrs := map[string]interface{}{
		string(semconv.IdentityDelegChainKey): []string{"root", "mid", "leaf"},
	}
	lineage := BuildDelegationLineage(attrs)
	if lineage == nil {
		t.Fatal("BuildDelegationLineage returned nil")
	}
	if len(lineage.Nodes) != 3 {
		t.Fatalf("len(Nodes) = %d, want 3", len(lineage.Nodes))
	}
	if lineage.Nodes[0].ParentUID != "" || lineage.Nodes[0].Depth != 0 {
		t.Errorf("node[0] = %+v, want empty parent and depth 0", lineage.Nodes[0])
	}
	if lineage.Nodes[2].ParentUID != "mid" || lineage.Nodes[2].Depth != 2 {
		t.Errorf("node[2] = %+v, want parent mid and depth 2", lineage.Nodes[2])
	}
	if lineage.Nodes[1].AgentUID != "mid" {
		t.Errorf("node[1].AgentUID = %q, want mid", lineage.Nodes[1].AgentUID)
	}
}

func TestCrosswalkTables(t *testing.T) {
	if got := OCSFAgentActivityCrosswalk[3]; got != "Update" { // activity 3 -> Step Execute
		t.Errorf("OCSFAgentActivityCrosswalk[3] = %q, want Update", got)
	}
	if got := OCSFAgentActivityCrosswalk[1]; got != "Spawn" {
		t.Errorf("OCSFAgentActivityCrosswalk[1] = %q, want Spawn", got)
	}
	if got := OCSFDelegationActivityCrosswalk["grant"]; got != "Create" {
		t.Errorf("OCSFDelegationActivityCrosswalk[grant] = %q, want Create", got)
	}

	entry, ok := OCSFClassCrosswalk["agent_activity"]
	if !ok {
		t.Fatal("OCSFClassCrosswalk missing agent_activity")
	}
	// AITF <= 0.4 routed this to agent_activity (9001) in a proposed "ai"
	// category. OCSF categories.json stops at uid 8, so it now reuses API
	// Activity with the ai_operation profile.
	if entry.OCSFClass != "api_activity" {
		t.Errorf("agent_activity OCSFClass = %q, want api_activity", entry.OCSFClass)
	}
	if entry.OCSFClassUID != ClassUIDAPIActivity || ClassUIDAPIActivity != 6003 {
		t.Errorf("agent_activity class_uid = %d, want 6003", entry.OCSFClassUID)
	}
	if entry.OCSFCategoryUID != OCSFCategoryUIDApplication {
		t.Errorf("agent_activity category = %d, want 6", entry.OCSFCategoryUID)
	}
	if d := OCSFClassCrosswalk["delegation_activity"]; d.OCSFClassUID != ClassUIDAuthorizeSession || d.OCSFClass != "authorize_session" {
		t.Errorf("delegation_activity = %d/%q, want 3003/authorize_session", d.OCSFClassUID, d.OCSFClass)
	}
	// No row may target the unratified ai category (uid 9).
	for name, e := range OCSFClassCrosswalk {
		if e.OCSFCategoryUID >= 9 || e.OCSFClassUID >= 9000 {
			t.Errorf("%s targets unratified category/class %d/%d", name, e.OCSFCategoryUID, e.OCSFClassUID)
		}
	}
	// Legacy UIDs stay decodable for telemetry recorded by AITF <= 0.4.
	if LegacyAIClassUIDs[9001] != ClassUIDAPIActivity || LegacyAIClassUIDs[9002] != ClassUIDAuthorizeSession || LegacyAIClassUIDs[9003] != ClassUIDAPIActivity {
		t.Errorf("LegacyAIClassUIDs = %v, want 9001/9003->6003, 9002->3003", LegacyAIClassUIDs)
	}

	// Inference and tool execution intentionally share API Activity (6003).
	inf := OCSFClassCrosswalk["model_inference"]
	tool := OCSFClassCrosswalk["tool_execution"]
	if inf.OCSFClassUID != 6003 || tool.OCSFClassUID != 6003 {
		t.Errorf("model_inference=%d, tool_execution=%d, both want 6003", inf.OCSFClassUID, tool.OCSFClassUID)
	}
	if inf.OCSFCategoryUID != 6 || tool.OCSFCategoryUID != 6 {
		t.Errorf("model_inference cat=%d, tool_execution cat=%d, both want 6", inf.OCSFCategoryUID, tool.OCSFCategoryUID)
	}
}
