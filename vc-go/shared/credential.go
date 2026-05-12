package shared

// type Proof struct {
// 	Type               string `json:"type"`
// 	Created            string `json:"created"`
// 	ProofPurpose       string `json:"proofPurpose"`
// 	VerificationMethod string `json:"verificationMethod"`
// 	JWS                string `json:"jws"`
// }

type Credential struct {
	Context           []string               `json:"@context"`
	ID                string                 `json:"id"`
	Type              []string               `json:"type"`
	Issuer            string                 `json:"issuer"`
	IssuanceDate      string                 `json:"issuanceDate"`
	CredentialSubject map[string]interface{} `json:"credentialSubject"`
	Proof             Proof                  `json:"proof"`
}

type AgentCapabilityCredential struct {
	Context           []string          `json:"@context"`
	ID                string            `json:"id"`
	Type              []string          `json:"type"`
	Issuer            string            `json:"issuer"`
	ValidFrom         string            `json:"validFrom"`
	ValidUntil        string            `json:"validUntil"`
	CredentialSubject CredentialSubject `json:"credentialSubject"`
	CredentialSchema  CredentialSchema  `json:"credentialSchema"`
	CredentialStatus  CredentialStatus  `json:"credentialStatus"`
	TermsOfUse        TermsOfUse        `json:"termsOfUse"`
	Proof             Proof             `json:"proof"`
}

// Credential subject
type CredentialSubject struct {
	ID                  string              `json:"id"`
	AgentProfile        AgentProfile        `json:"agentProfile"`
	Capabilities        []Capability        `json:"capabilities"`
	Restrictions        Restrictions        `json:"restrictions"`
	SandboxPolicy       SandboxPolicy       `json:"sandboxPolicy"`
	DelegationPolicy    DelegationPolicy    `json:"delegationPolicy"`
	HumanApprovalPolicy HumanApprovalPolicy `json:"humanApprovalPolicy"`
}

// Agent metadata
type AgentProfile struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Model   string `json:"model"`
}

// Capability definition
type Capability struct {
	ID          string                `json:"id"`
	Actions     []string              `json:"actions"`
	Constraints CapabilityConstraints `json:"constraints"`
}

// Capability constraints
type CapabilityConstraints struct {
	AllowedDomains        []string   `json:"allowedDomains,omitempty"`
	AllowedPaths          []string   `json:"allowedPaths,omitempty"`
	RateLimit             *RateLimit `json:"rateLimit,omitempty"`
	MaxExecutionTimeMs    int        `json:"maxExecutionTimeMs,omitempty"`
	RequiresHumanApproval bool       `json:"requiresHumanApproval"`
}

// Rate limiting
type RateLimit struct {
	RequestsPerMinute int `json:"requestsPerMinute"`
}

// Restrictions
type Restrictions struct {
	BlockedCapabilities []string      `json:"blockedCapabilities"`
	NetworkPolicy       NetworkPolicy `json:"networkPolicy"`
}

// Network policy
type NetworkPolicy struct {
	AllowExternalRequests bool     `json:"allowExternalRequests"`
	BlockedDomains        []string `json:"blockedDomains"`
}

// Sandbox policy
type SandboxPolicy struct {
	MemoryLimitMB       int    `json:"memoryLimitMB"`
	CPULimit            string `json:"cpuLimit"`
	FilesystemIsolation string `json:"filesystemIsolation"`
	EphemeralFilesystem bool   `json:"ephemeralFilesystem"`
}

// Delegation policy
type DelegationPolicy struct {
	Allowed  bool `json:"allowed"`
	MaxDepth int  `json:"maxDepth"`
}

// Human approval
type HumanApprovalPolicy struct {
	Approver string `json:"approver"`
	Mode     string `json:"mode"`
}

// Credential schema reference
type CredentialSchema struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Credential revocation / status
type CredentialStatus struct {
	ID                   string `json:"id"`
	Type                 string `json:"type"`
	StatusPurpose        string `json:"statusPurpose"`
	StatusListIndex      string `json:"statusListIndex"`
	StatusListCredential string `json:"statusListCredential"`
}

// Terms of use
type TermsOfUse struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// Linked Data Proof
type Proof struct {
	Type               string `json:"type"`
	Cryptosuite        string `json:"cryptosuite"`
	Created            string `json:"created"`
	VerificationMethod string `json:"verificationMethod"`
	ProofPurpose       string `json:"proofPurpose"`
	ProofValue         string `json:"proofValue"`
}
