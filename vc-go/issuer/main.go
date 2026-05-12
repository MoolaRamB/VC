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

	issuerDID := "did:key:issuer123"

	// 2. Create credential (without proof first)
	cred := shared.Credential{
		Context:      []string{"https://www.w3.org/2018/credentials/v1"},
		ID:           "http://example.edu/credentials/1872",
		Type:         []string{"VerifiableCredential", "DegreeCredential"},
		Issuer:       issuerDID,
		IssuanceDate: time.Now().Format(time.RFC3339),
		CredentialSubject: map[string]interface{}{
			"id":     "did:key:student123",
			"degree": "B.Tech",
		},
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
		JWS:                base64.StdEncoding.EncodeToString(signature),
	}

	// 6. Final VC
	finalBytes, _ := json.MarshalIndent(cred, "", "  ")

	fmt.Println(string(finalBytes))

	// Save public key for verifier (copy manually for now)
	fmt.Println("\nPublic Key (base64):")
	fmt.Println(base64.StdEncoding.EncodeToString(pub))
}
