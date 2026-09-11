# MCP Span Conventions (MCP_ACTIVITY)

> **OTel Alignment Note:** MCP tool invocations adopt OTel GenAI tool attributes (`gen_ai.tool.{name,type,call.id,call.arguments,call.result}`)
> where OTel defines them. MCP-specific extension attributes (server lifecycle, resources, prompts, sampling) use the `mcp.*` namespace.

Status: **Normative** | CoSAI WS2 Alignment: **MCP_ACTIVITY** | OCSF Class: **API Activity (6003)** (`ai_operation` profile)

AITF defines semantic conventions for the Model Context Protocol (MCP), covering server lifecycle, tool discovery and invocation, resource access, prompt management, and sampling. This specification defines the normative field requirements aligned with CoSAI Working Stream 2 (Telemetry for AI) and mapped to applicable compliance and threat frameworks.

Key words "MUST", "SHOULD", "MAY" follow [RFC 2119](https://www.rfc-editor.org/rfc/rfc2119).

---

## Overview

MCP is a protocol that enables AI models to interact with external systems through a standardized interface. AITF provides first-class telemetry for all MCP operations:

```
MCP Server Lifecycle:
  connect -> initialize -> discover (tools/resources/prompts) -> use -> disconnect

MCP Operations:
  - Tool Discovery & Invocation
  - Resource Read & Subscribe
  - Prompt Get & Execute
  - Sampling (server-initiated LLM requests)
  - Root Management
```

---

## Cross-Cutting: Primitive & Server Identity [RFC v0.4 gap closure]

> **RFC field:** MCP Server Identity & Primitive · **Tier:** SHOULD · **§9 Tools & External Services** · **Grounding attacks:** `TA-06`, `AOC-10`, `TA-01`, `AOC-09`, `TA-12`, `TA-13`

The span names below already distinguish a tool call from a resource read, but a consumer analysing a mixed MCP stream has to parse span names to recover which protocol surface was exercised — and span names are not stable analytical keys. Two of the six primitives, `sampling` and `elicitation`, invert the direction of trust: the server drives the client. Telemetry that cannot cheaply separate those from ordinary tool traffic cannot detect a server that has begun steering the host.

`mcp.primitive` makes the surface explicit and queryable. It applies to **every** MCP span in this document.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.primitive` | string | **Conditionally Required** | Protocol surface exercised: `"tool"`, `"resource"`, `"prompt"`, `"sampling"`, `"completion"`, `"elicitation"`, `"root"`, `"logging"` | OWASP LLM06 (Excessive Agency), MITRE ATLAS [AML.T0040](https://atlas.mitre.org/techniques/AML.T0040) |
| `mcp.method.name` | string | **Conditionally Required** | JSON-RPC method as sent on the wire (`"tools/call"`, `"resources/read"`, `"sampling/createMessage"`) | MITRE ATLAS [AML.T0040](https://atlas.mitre.org/techniques/AML.T0040) |
| `mcp.request.id` | string | **Conditionally Required** | JSON-RPC request identifier, correlating a request to its response across the transport | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 (Record-Keeping) |
| `mcp.session.id` | string | **Recommended** | MCP session identifier where the transport carries one, distinct from `mcp.connection.id` | NIST AI RMF GOVERN-1.2 |
| `mcp.server.instance.id` | string | **Recommended** | Identity of the specific running server process, not just its configured name | MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010) |
| `mcp.server.identity.method` | string | **Recommended** | How server identity was established: `"none"`, `"config"`, `"tls"`, `"oauth"`, `"signature"`, `"binary_hash"` | NIST AI RMF GOVERN-1.2 |
| `mcp.server.identity.verified` | boolean | **Recommended** | Whether that identity was verified. Absent or `false` means the server is trusted purely on its configured name | NIST AI RMF GOVERN-1.2, MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010) |
| `mcp.server.trust_domain` | string | **Recommended** | Trust domain the server belongs to. A local `stdio` server and a third-party hosted one warrant different treatment | OWASP LLM06, NIST AI RMF GOVERN-1.2 |
| `mcp.server.command` | string | **Optional** | Command line used to launch a `stdio` server — the substitution surface | MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010) |
| `mcp.server.binary.hash` | string | **Optional** | Digest of the server binary or entrypoint script | MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010) |
| `mcp.capabilities.negotiated` | string[] | **Optional** | Capabilities agreed during `initialize`, as the authoritative record of what the connection is permitted to do | OWASP LLM06 (Excessive Agency) |

---

## Span: `mcp.server.connect`

Represents establishing a connection to an MCP server.

### Span Name

Format: `mcp.server.connect {mcp.server.name}`

### Span Kind

`CLIENT`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.server.name` | string | **Required** | MCP server name | OWASP LLM06 (Excessive Agency), MITRE ATLAS [AML.T0040](https://atlas.mitre.org/techniques/AML.T0040) |
| `mcp.server.transport` | string | **Required** | Transport type: `"stdio"`, `"sse"`, `"streamable_http"` | MITRE ATLAS [AML.T0040](https://atlas.mitre.org/techniques/AML.T0040) (ML Supply Chain) |
| `mcp.connection.id` | string | **Recommended** | Unique connection identifier for session correlation | NIST AI RMF GOVERN-1.2 |
| `mcp.server.version` | string | **Recommended** | MCP server version | NIST AI RMF MAP-1.1 |
| `mcp.server.url` | string | **Recommended** | Server URL (if network transport) | MITRE ATLAS AML.T0040 |
| `mcp.protocol.version` | string | **Recommended** | MCP protocol version (e.g. `"2025-03-26"`) | NIST AI RMF MAP-1.1 |

### Events

#### `mcp.server.capabilities`

Emitted after successful initialization.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.capabilities.tools` | boolean | **Recommended** | Server supports tools | OWASP LLM06 |
| `mcp.capabilities.resources` | boolean | **Recommended** | Server supports resources | — |
| `mcp.capabilities.prompts` | boolean | **Optional** | Server supports prompts | — |
| `mcp.capabilities.sampling` | boolean | **Optional** | Server supports sampling | — |
| `mcp.capabilities.roots` | boolean | **Optional** | Server supports roots | — |

---

## Span: `mcp.server.disconnect`

Represents disconnecting from an MCP server.

### Span Name

Format: `mcp.server.disconnect {mcp.server.name}`

### Span Kind

`CLIENT`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.server.name` | string | **Required** | Server name | OWASP LLM06 |
| `mcp.connection.id` | string | **Recommended** | Connection identifier | NIST AI RMF GOVERN-1.2 |

---

## Span: `mcp.tool.discover`

Represents discovering available tools from an MCP server.

### Span Name

Format: `mcp.tool.discover {mcp.server.name}`

### Span Kind

`CLIENT`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.server.name` | string | **Required** | Server name | OWASP LLM06 |
| `mcp.tool.count` | int | **Recommended** | Number of tools discovered | OWASP LLM06 |
| `mcp.tool.names` | string[] | **Recommended** | Names of discovered tools | OWASP LLM06, MITRE ATLAS AML.T0048 |
| `mcp.connection.id` | string | **Optional** | Connection identifier | NIST AI RMF GOVERN-1.2 |

---

## Span: `mcp.tool.invoke`

Represents invoking a tool on an MCP server. This is the primary MCP telemetry span.

### Span Name

Format: `execute_tool {gen_ai.tool.name}`

### Span Kind

`CLIENT`

### Normative Field Table

#### Tool Identification

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.tool.name` | string | **Required** | Tool name | OWASP LLM06 (Excessive Agency), MITRE ATLAS [AML.T0048](https://atlas.mitre.org/techniques/AML.T0048) |
| `gen_ai.tool.type` | string | **Recommended** | Tool type (value: `"extension"` for MCP tools) | OWASP LLM06 |
| `gen_ai.tool.call.id` | string | **Recommended** | Unique identifier for this tool call | NIST AI RMF GOVERN-1.2 |
| `gen_ai.tool.description` | string | **Optional** | Human-readable tool description | EU AI Act Art.13 |
| `mcp.tool.server` | string | **Required** | Source MCP server name | OWASP LLM06, MITRE ATLAS [AML.T0040](https://atlas.mitre.org/techniques/AML.T0040) |
| `mcp.connection.id` | string | **Recommended** | Connection identifier for session correlation | NIST AI RMF GOVERN-1.2 |

#### Request & Response (CoSAI WS2)

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `gen_ai.tool.call.arguments` | string | **Recommended** | Input parameters (JSON) | OWASP LLM01 (Prompt Injection via tools), MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `gen_ai.tool.call.result` | string | **Recommended** | Tool output (may be redacted for sensitive content) | OWASP LLM05 (Improper Output), OWASP LLM02 (Sensitive Info) |
| `mcp.tool.response_error` | string | **Recommended** | Error message content when tool execution fails | NIST AI RMF MEASURE-2.5 |
| `mcp.tool.is_error` | boolean | **Recommended** | Whether tool returned an error | NIST AI RMF MEASURE-2.5 |

#### Execution Metadata

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.tool.duration_ms` | double | **Recommended** | Execution duration in milliseconds | NIST AI RMF MEASURE-2.5, OWASP LLM10 |
| `mcp.server.transport` | string | **Recommended** | Transport type | MITRE ATLAS AML.T0040 |

#### Human Approval

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.tool.approval_required` | boolean | **Recommended** | Whether human approval was needed | EU AI Act Art.14 (Human Oversight) |
| `mcp.tool.approved` | boolean | **Recommended** | Whether the tool was approved (if required) | EU AI Act Art.14 |

These two fields record *that* approval happened. They do not record who approved, what they were shown, or how long they took, and an investigation into approval fatigue or consent spoofing needs all three. Where those questions matter, emit the [`identity.approval.*`](identity-spans.md#human-approval--elicitation-rfc-v04-gap-closure) event pair alongside them; `mcp.tool.approval_required` and `mcp.tool.approved` remain as the cheap in-span summary.

#### Tool Definition Digest [RFC v0.4 gap closure]

> **RFC field:** Tool Definition Digest · **Tier:** SHOULD · **§9 Tools & External Services** · **Grounding attacks:** `AOC-10`, `AOC-14`, `TA-06`, `AOC-09` · **ODIS:** `approved_software_refs` (6.1)

`gen_ai.tool.name` identifies the tool a user approved. It does not identify the tool that then ran. An MCP server can return one definition at `tools/list` and a different one on the next call — a changed description that instructs the model differently, or a widened argument schema that accepts a path it previously refused. The name is unchanged throughout, so name-based approval survives the substitution intact.

The digest is computed over the tool's declared contract **as presented at invocation time**, canonicalised with [RFC 8785 JCS](https://www.rfc-editor.org/rfc/rfc8785) over the object `{"name", "description", "inputSchema", "outputSchema"}` and prefixed `sha256:`.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.tool.definition.hash` | string | **Conditionally Required** | Digest of the contract as presented for this call | MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010), OWASP LLM06 (Excessive Agency) |
| `mcp.tool.definition.baseline_hash` | string | **Recommended** | Digest recorded when the tool was approved (ODIS `approved_software_refs`) | NIST AI RMF GOVERN-1.2, MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010) |
| `mcp.tool.definition.changed` | boolean | **Conditionally Required** | Whether `hash` differs from `baseline_hash`. `true` on a tool the user approved earlier is the rug-pull signature | MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010), OWASP LLM06 |
| `mcp.tool.definition.change_type` | string | **Recommended** | What changed: `"description"`, `"input_schema"`, `"output_schema"`, `"multiple"`, `"new_tool"` | MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010) |
| `mcp.tool.definition.description_hash` | string | **Optional** | Digest of the description alone. Description text reaches the model as instruction, so it is the injection-bearing component | OWASP LLM01 (Prompt Injection), MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051) |
| `mcp.tool.definition.schema_hash` | string | **Optional** | Digest of the argument schema alone. Schema widening expands what the tool will accept | OWASP LLM06 (Excessive Agency) |
| `mcp.tool.definition.approved` | boolean | **Recommended** | Whether *this* digest, not merely this tool name, is on the approved list | NIST AI RMF GOVERN-1.2, EU AI Act Art.12 |
| `mcp.tool.definition.approved_at` | string | **Optional** | RFC 3339 timestamp the baseline digest was approved | EU AI Act Art.12 (Record-Keeping) |
| `mcp.tool.definition.first_seen` | string | **Optional** | RFC 3339 timestamp this digest was first observed | NIST AI RMF MEASURE-2.5 |
| `mcp.tool.definition.source` | string | **Optional** | Where the definition came from for this call: `"tools_list"`, `"cache"`, `"inline"`, `"config"` | MITRE ATLAS [AML.T0010](https://atlas.mitre.org/techniques/AML.T0010) |

**Emission rule.** The digest MUST be attached to every `mcp.tool.invoke` span, not only to `mcp.tool.discover`. Recording it only at discovery time is precisely the gap a rug-pull exploits: the definition that was listed and the definition that ran are different objects, and only the invocation span can attest to the second.

#### Execution Environment

An MCP tool call's blast radius depends on where the tool body executes. Where the invocation is delegated to a sandbox — a code interpreter, a per-call container, a remote execution service — attach [`supply_chain.runtime.*`](attributes-registry.md#supply_chainruntime-rfc-v04-gap-closure) to this span, since the environment differs between spans within one process and cannot be carried as a Resource attribute.

### Events

#### `mcp.tool.input`

Emitted with tool input (can be sampled/redacted for sensitive inputs).

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.tool.input.content` | string | **Optional** | Full input content | OWASP LLM01 |

#### `mcp.tool.output`

Emitted with tool output (can be sampled/redacted).

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.tool.output.content` | string | **Optional** | Full output content | OWASP LLM05, OWASP LLM02 |
| `mcp.tool.output.type` | string | **Optional** | Output type: `"text"`, `"image"`, `"resource"` | — |

#### `mcp.tool.approval`

Emitted when human approval is requested/granted.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.tool.approval.status` | string | **Required** | `"requested"`, `"approved"`, `"denied"` | EU AI Act Art.14 |
| `mcp.tool.approval.approver` | string | **Recommended** | Who approved (if applicable) | EU AI Act Art.14, NIST AI RMF GOVERN-1.7 |

---

## Span: `mcp.resource.read`

Represents reading a resource from an MCP server.

### Span Name

Format: `mcp.resource.read {mcp.resource.uri}`

### Span Kind

`CLIENT`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.resource.uri` | string | **Required** | Resource URI | OWASP LLM02 (Sensitive Info) |
| `mcp.server.name` | string | **Required** | Server name | OWASP LLM06 |
| `mcp.connection.id` | string | **Optional** | Connection identifier | NIST AI RMF GOVERN-1.2 |
| `mcp.resource.name` | string | **Recommended** | Resource display name | — |
| `mcp.resource.mime_type` | string | **Recommended** | Content MIME type | — |
| `mcp.resource.size_bytes` | int | **Optional** | Content size in bytes | OWASP LLM10 |

---

## Span: `mcp.resource.subscribe`

Represents subscribing to resource updates from an MCP server.

### Span Name

Format: `mcp.resource.subscribe {mcp.resource.uri}`

### Span Kind

`CLIENT`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.resource.uri` | string | **Required** | Resource URI | OWASP LLM02 |
| `mcp.server.name` | string | **Required** | Server name | OWASP LLM06 |
| `mcp.connection.id` | string | **Optional** | Connection identifier | NIST AI RMF GOVERN-1.2 |

---

## Span: `mcp.prompt.get`

Represents retrieving a prompt template from an MCP server.

### Span Name

Format: `mcp.prompt.get {mcp.prompt.name}`

### Span Kind

`CLIENT`

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.prompt.name` | string | **Required** | Prompt template name | OWASP LLM01 (Prompt Injection) |
| `mcp.server.name` | string | **Required** | Server name | OWASP LLM06 |
| `mcp.prompt.arguments` | string | **Recommended** | Prompt arguments (JSON) | OWASP LLM01 |
| `mcp.prompt.description` | string | **Optional** | Prompt description | EU AI Act Art.13 |

---

## Span: `mcp.sampling.request`

Represents a server-initiated sampling (LLM) request via MCP.

### Span Name

Format: `mcp.sampling.request {mcp.server.name}`

### Span Kind

`SERVER` (server is requesting from the client)

### Normative Field Table

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.server.name` | string | **Required** | Requesting server name | OWASP LLM06, MITRE ATLAS AML.T0048 |
| `mcp.sampling.model` | string | **Required** | Requested model hint | MITRE ATLAS AML.T0044 |
| `mcp.sampling.max_tokens` | int | **Recommended** | Max tokens | OWASP LLM10 |
| `mcp.sampling.include_context` | string | **Optional** | Context scope: `"thisServer"`, `"allServers"` | OWASP LLM02 |
| `gen_ai.usage.input_tokens` | int | **Recommended** | Input tokens used | OWASP LLM10 |
| `gen_ai.usage.output_tokens` | int | **Recommended** | Output tokens used | OWASP LLM10 |

---

## Protocol Envelope Capture [RFC v0.4 gap closure]

> **RFC field:** Protocol Envelope Capture · **Tier:** MAY · **§12 Orchestration, Multi-Agent & Background Execution** · **Grounding attacks:** `AOC-09`, `AOC-12`, `TA-06`

Every field above is an *interpretation* of the wire protocol produced by whichever framework sat between the transport and the instrumentation. That interpretation is lossy by design: unrecognised fields, protocol extensions, malformed structures that the parser silently repaired, and the exact byte-level shape of an argument are all discarded before an attribute is ever set. When the question is whether a peer sent something the framework did not expect, the framework's own summary is the wrong witness.

Envelope capture preserves the raw JSON-RPC message alongside the interpreted fields. It is tiered **MAY** because the envelope contains full request and response payloads: enabling it is a deliberate, policy-gated decision, not a default. The same fields apply to A2A JSON-RPC traffic.

| Field Name | Type | Requirement | Description | Compliance |
|---|---|---|---|---|
| `mcp.envelope.captured` | boolean | **Optional** | Whether raw envelope capture was active. Distinguishes "nothing unusual on the wire" from "nobody was looking" | NIST AI RMF MEASURE-2.5 |
| `mcp.envelope.request` | string | **Optional** | Raw JSON-RPC request envelope | MITRE ATLAS [AML.T0051](https://atlas.mitre.org/techniques/AML.T0051), OWASP LLM01 (Prompt Injection) |
| `mcp.envelope.response` | string | **Optional** | Raw JSON-RPC response envelope | OWASP LLM05 (Improper Output Handling) |
| `mcp.envelope.request.hash` | string | **Optional** | Digest of the request envelope, allowing integrity comparison without retaining the payload | EU AI Act Art.12 (Record-Keeping) |
| `mcp.envelope.response.hash` | string | **Optional** | Digest of the response envelope | EU AI Act Art.12 |
| `mcp.envelope.size_bytes` | int | **Optional** | Envelope size. Sudden growth in response size against a stable request is an exfiltration or injection indicator | OWASP LLM10 (Unbounded Consumption) |
| `mcp.envelope.truncated` | boolean | **Optional** | Whether the captured envelope was truncated. A truncated envelope is not evidence of absence | NIST AI RMF MEASURE-2.5 |
| `mcp.envelope.redacted` | boolean | **Optional** | Whether fields were removed before capture | GDPR Art.5 (Data Minimisation), GDPR Art.32 |
| `mcp.envelope.jsonrpc.version` | string | **Optional** | JSON-RPC version declared in the envelope | — |
| `mcp.envelope.error.code` | int | **Optional** | JSON-RPC error code, which carries protocol-level detail that a framework exception often flattens away | NIST AI RMF MEASURE-2.5 |
| `mcp.envelope.error.message` | string | **Optional** | JSON-RPC error message as sent | NIST AI RMF MEASURE-2.5 |

**Privacy note.** `mcp.envelope.request` and `mcp.envelope.response` carry unredacted content. Where they are enabled, `mcp.envelope.redacted` MUST accurately reflect whether redaction ran, since a consumer treating an unredacted envelope as redacted is a worse outcome than not capturing it at all.

---

## CoSAI WS2 Field Mapping

Cross-reference between CoSAI WS2 `MCP_ACTIVITY` field names and AITF attribute keys:

| CoSAI WS2 Field | AITF Attribute | Notes |
|---|---|---|
| `mcp.server.name` | `mcp.server.name` | Direct match |
| `mcp.tool.name` | `gen_ai.tool.name` | OTel GenAI tool attribute |
| `mcp.request.args` | `gen_ai.tool.call.arguments` | OTel GenAI tool attribute; JSON-encoded input parameters |
| `mcp.response.result` | `gen_ai.tool.call.result` | OTel GenAI tool attribute; may be redacted |
| `mcp.response.error` | `mcp.tool.response_error` | New in CoSAI WS2 alignment |
| `mcp.transport` | `mcp.server.transport` | On server/tool spans |
| `mcp.connection.id` | `mcp.connection.id` | New in CoSAI WS2 alignment |
| — | `mcp.primitive`, `mcp.method.name`, `mcp.request.id` | New in RFC v0.4 gap closure; no CoSAI WS2 counterpart yet |
| — | `mcp.tool.definition.*` | New in RFC v0.4 gap closure |
| — | `mcp.envelope.*` | New in RFC v0.4 gap closure; policy-gated raw capture |

---

## Security Considerations

MCP tool invocations should be monitored for:

1. **Unauthorized file access** -- Tools accessing files outside allowed paths
2. **Command injection** -- Malicious input parameters
3. **Data exfiltration** -- Tools sending sensitive data to external systems
4. **Privilege escalation** -- Tools operating beyond granted permissions

AITF's Security Processor automatically flags suspicious MCP tool invocations based on configurable policies.

---

## Example: MCP Tool Discovery and Invocation

```
Span: mcp.server.connect filesystem
  mcp.server.transport: "stdio"
  mcp.protocol.version: "2025-03-26"
  mcp.connection.id: "conn-fs-abc123"
  |
  +- Event: mcp.server.capabilities
  |    mcp.capabilities.tools: true
  |    mcp.capabilities.resources: true
  |
  +- Span: mcp.tool.discover filesystem
  |    mcp.tool.count: 5
  |    mcp.tool.names: ["read_file", "write_file", "list_dir", "search", "move_file"]
  |
  +- Span: execute_tool read_file
  |    gen_ai.tool.name: "read_file"
  |    gen_ai.tool.type: "extension"
  |    mcp.tool.server: "filesystem"
  |    gen_ai.tool.call.arguments: "{\"path\":\"/data/config.yaml\"}"
  |    mcp.tool.is_error: false
  |    mcp.tool.response_error: ""
  |    mcp.tool.duration_ms: 12.5
  |    mcp.connection.id: "conn-fs-abc123"
  |    Events:
  |      mcp.tool.input: {content: "{\"path\":\"/data/config.yaml\"}"}
  |      mcp.tool.output: {content: "server:\n  port: 8080\n...", type: "text"}
  |
  +- Span: execute_tool write_file
       gen_ai.tool.name: "write_file"
       gen_ai.tool.type: "extension"
       mcp.tool.server: "filesystem"
       mcp.primitive: "tool"
       mcp.method.name: "tools/call"
       mcp.request.id: "42"
       mcp.tool.approval_required: true
       mcp.tool.approved: true
       mcp.tool.duration_ms: 25.0
       mcp.connection.id: "conn-fs-abc123"
       mcp.tool.definition.hash: "sha256:9f2c...d41a"
       mcp.tool.definition.baseline_hash: "sha256:0b77...e83c"
       mcp.tool.definition.changed: true
       mcp.tool.definition.change_type: "input_schema"
       mcp.tool.definition.approved: false
       Events:
         mcp.tool.approval: {status: "approved", approver: "user@example.com"}
         mcp.tool.input: {content: "{\"path\":\"/data/output.txt\",\"content\":\"...\"}"}
         mcp.tool.output: {content: "File written successfully", type: "text"}
```

The `write_file` span carries the finding. The tool name is the one the user approved and `mcp.tool.approved` is `true`, so every name-based control passed. But `mcp.tool.definition.changed` is `true` against a schema change, and `mcp.tool.definition.approved` is `false` — the contract that executed is not the contract that was approved. Without the digest, this span is indistinguishable from a routine approved write.
