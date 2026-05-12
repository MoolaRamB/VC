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
	// 🔴 Paste VC JSON here
	vcJSON := `{
  "@context": [
    "https://www.w3.org/ns/credentials/v2",
    "https://agentvc.org/contexts/agent-security-v1.jsonld"
  ],
  "id": "urn:uuid:8fa21c52-6e21-4f45-8f57-7f1e2b31f001",
  "type": [
    "VerifiableCredential",
    "AgentCapabilityCredential"
  ],
  "issuer": "did:key:issuerEnsurity",
  "validFrom": "2026-05-12T13:06:25Z",
  "validUntil": "2026-05-12T13:07:25Z",
  "credentialSubject": {
    "id": "did:key:z6MkResearchAgent01",
    "agentProfile": {
      "name": "ResearchAgent",
      "version": "2.0",
      "model": "gpt-agent-class"
    },
    "capabilities": [
      {
        "id": "urn:cap:web-search",
        "actions": [
          "search",
          "summarize"
        ],
        "constraints": {
          "allowedDomains": [
            "*.wikipedia.org",
            "*.arxiv.org"
          ],
          "rateLimit": {
            "requestsPerMinute": 30
          },
          "maxExecutionTimeMs": 10000,
          "requiresHumanApproval": false
        }
      },
      {
        "id": "urn:cap:filesystem-read",
        "actions": [
          "read"
        ],
        "constraints": {
          "allowedPaths": [
            "/workspace/docs/*"
          ],
          "requiresHumanApproval": true
        }
      }
    ],
    "restrictions": {
      "blockedCapabilities": [
        "shell.execute",
        "payment.transfer",
        "system.shutdown"
      ],
      "networkPolicy": {
        "allowExternalRequests": true,
        "blockedDomains": [
          "*.onion",
          "*.torrent"
        ]
      }
    },
    "sandboxPolicy": {
      "memoryLimitMB": 512,
      "cpuLimit": "1-core",
      "filesystemIsolation": "containerized",
      "ephemeralFilesystem": true
    },
    "delegationPolicy": {
      "allowed": true,
      "maxDepth": 1
    },
    "humanApprovalPolicy": {
      "approver": "did:key:z6MkAliceOperator",
      "mode": "interactive"
    }
  },
  "credentialSchema": {
    "id": "https://agentvc.org/schemas/agent-capability-credential-v1.json",
    "type": "JsonSchema"
  },
  "credentialStatus": {
    "id": "https://security-layer.example/status/8fa21",
    "type": "BitstringStatusListEntry",
    "statusPurpose": "revocation",
    "statusListIndex": "42",
    "statusListCredential": "https://security-layer.example/status-list/2026"
  },
  "termsOfUse": {
    "type": "IssuerPolicy",
    "id": "https://security-layer.example/policies/agent-usage"
  },
  "proof": {
    "type": "Ed25519Signature2020",
    "created": "2026-05-12T18:36:25+05:30",
    "verificationMethod": "did:key:issuerEnsurity#key-1",
    "proofPurpose": "assertionMethod",
    "proofValue": "Q8hikoBQ/JcB2JsEn3CgkqUx4ikBIHg3FwQh/+T8fSRPEG7x13fFmfg4AdWh9AtrDrcprqZmDLuOOZAAqEyFCg=="
  }
}`

	//  Paste issuer public key
	pubKeyBase64 := "ejXIuQGQvz8wCmt3poeuSB9R6mkPVGPawSqLCpScpFY="

	pubKeyBytes, _ := base64.StdEncoding.DecodeString(pubKeyBase64)
	publicKey := ed25519.PublicKey(pubKeyBytes)

	var cred shared.AgentCapabilityCredential
	json.Unmarshal([]byte(vcJSON), &cred)

	// Check validity period
	now := time.Now().UTC()
	validFrom, _ := time.Parse(time.RFC3339, cred.ValidFrom)
	validUntil, _ := time.Parse(time.RFC3339, cred.ValidUntil)

	if now.Before(validFrom) {
		fmt.Println("❌ Credential is not yet valid")
		return
	}
	if now.After(validUntil) {
		fmt.Println("❌ Credential has expired")
		return
	}

	// Remove proof before verification
	proof := cred.Proof
	cred.Proof = shared.Proof{}

	unsignedBytes, _ := json.Marshal(cred)

	sigBytes, _ := base64.StdEncoding.DecodeString(proof.ProofValue)

	valid := ed25519.Verify(publicKey, unsignedBytes, sigBytes)

	if valid {
		fmt.Println("✅ Credential is VALID")
	} else {
		fmt.Println("❌ Credential is INVALID")
	}
}
