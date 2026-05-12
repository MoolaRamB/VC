package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"vc-go/shared"
)

func main() {
	// 1. Generate keypair (in real world, persist this)
	pub, priv, _ := ed25519.GenerateKey(nil)

	issuerDID := "did:key:issuerEnsurity"

	// 2. Create credential (without proof first)
	// cred := shared.Credential{
	// 	Context:      []string{"https://www.w3.org/2018/credentials/v1", "https://agentvc.org/contexts/agent-security-v1.jsonld"},
	// 	ID:           "urn:uuid:8fa21c52-6e21-4f45-8f57-7f1e2b31f001",
	// 	Type:         []string{"VerifiableCredential", "AgentCapabilityCredential"},
	// 	Issuer:       issuerDID,
	// 	IssuanceDate: time.Now().Format(time.RFC3339),
	// 	CredentialSubject: map[string]interface{}{
	// 		"id":     "did:key:student123",
	// 		"degree": "B.Tech",
	// 	},
	// }

	cred := shared.AgentCapabilityCredential{
	Context: []string{
		"https://www.w3.org/ns/credentials/v2",
		"https://agentvc.org/contexts/agent-security-v1.jsonld",
	},

	ID: "urn:uuid:8fa21c52-6e21-4f45-8f57-7f1e2b31f001",

	Type: []string{
		"VerifiableCredential",
		"AgentCapabilityCredential",
	},

	Issuer: issuerDID,

	ValidFrom:  time.Now().UTC().Format(time.RFC3339),
	ValidUntil: time.Now().UTC().Add(24 * time.Hour).Format(time.RFC3339),

	CredentialSubject: shared.CredentialSubject{
		ID: "did:key:z6MkResearchAgent01",

		AgentProfile: shared.AgentProfile{
			Name:    "ResearchAgent",
			Version: "2.0",
			Model:   "gpt-agent-class",
		},

		Capabilities: []shared.Capability{
			{
				ID: "urn:cap:web-search",

				Actions: []string{
					"search",
					"summarize",
				},

				Constraints: shared.CapabilityConstraints{
					AllowedDomains: []string{
						"*.wikipedia.org",
						"*.arxiv.org",
					},

					RateLimit: &shared.RateLimit{
						RequestsPerMinute: 30,
					},

					MaxExecutionTimeMs:    10000,
					RequiresHumanApproval: false,
				},
			},

			{
				ID: "urn:cap:filesystem-read",

				Actions: []string{
					"read",
				},

				Constraints: shared.CapabilityConstraints{
					AllowedPaths: []string{
						"/workspace/docs/*",
					},

					RequiresHumanApproval: true,
				},
			},
		},

		Restrictions: shared.Restrictions{
			BlockedCapabilities: []string{
				"shell.execute",
				"payment.transfer",
				"system.shutdown",
			},

			NetworkPolicy: shared.NetworkPolicy{
				AllowExternalRequests: true,

				BlockedDomains: []string{
					"*.onion",
					"*.torrent",
				},
			},
		},

		SandboxPolicy: shared.SandboxPolicy{
			MemoryLimitMB:       512,
			CPULimit:            "1-core",
			FilesystemIsolation: "containerized",
			EphemeralFilesystem: true,
		},

		DelegationPolicy: shared.DelegationPolicy{
			Allowed:  true,
			MaxDepth: 1,
		},

		HumanApprovalPolicy: shared.HumanApprovalPolicy{
			Approver: "did:key:z6MkAliceOperator",
			Mode:     "interactive",
		},
	},

	CredentialSchema: shared.CredentialSchema{
		ID:   "https://agentvc.org/schemas/agent-capability-credential-v1.json",
		Type: "JsonSchema",
	},

	CredentialStatus: shared.CredentialStatus{
		ID:                   "https://security-layer.example/status/8fa21",
		Type:                 "BitstringStatusListEntry",
		StatusPurpose:        "revocation",
		StatusListIndex:      "42",
		StatusListCredential: "https://security-layer.example/status-list/2026",
	},

	TermsOfUse: shared.TermsOfUse{
		Type: "IssuerPolicy",
		ID:   "https://security-layer.example/policies/agent-usage",
	},

	Proof: shared.Proof{},
}

	// 3. Serialize WITHOUT proof
	credBytes, _ := json.Marshal(cred)

	// 4. Sign
	signature := ed25519.Sign(priv, credBytes)

	// 5. Attach proof
	cred.Proof = shared.Proof{
		Type:               "Ed25519Signature2020",
		Created:            time.Now().Format(time.RFC3339),
		ProofPurpose:       "assertionMethod",
		VerificationMethod: issuerDID + "#key-1",
		ProofValue:         base64.StdEncoding.EncodeToString(signature),
	}

	// 6. Final VC
	finalBytes, _ := json.MarshalIndent(cred, "", "  ")

	fmt.Println(string(finalBytes))

	// Save public key for verifier (copy manually for now)
	fmt.Println("\nPublic Key (base64):")
	fmt.Println(base64.StdEncoding.EncodeToString(pub))
}
