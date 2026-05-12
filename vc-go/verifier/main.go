package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"vc-go/shared"
)

func main() {
	// 🔴 Paste VC JSON here
	vcJSON := `{
  "@context": [
    "https://www.w3.org/2018/credentials/v1"
  ],
  "id": "http://example.edu/credentials/1872",
  "type": [
    "VerifiableCredential",
    "DegreeCredential"
  ],
  "issuer": "did:key:issuer123",
  "issuanceDate": "2026-05-12T17:52:36+05:30",
  "credentialSubject": {
    "degree": "B.Tech",
    "id": "did:key:student123"
  },
  "proof": {
    "type": "Ed25519Signature2020",
    "created": "2026-05-12T17:52:36+05:30",
    "proofPurpose": "assertionMethod",
    "verificationMethod": "did:key:issuer123#key-1",
    "jws": "52mJgKJX9sqZLhdXRRO5QFkpB4iYGMf8Or/K76/4a8uI3mYkK55QPESSdvtoiOAALPis7DWjrowthGTmQYZmAQ=="
  }
}`

	// 🔴 Paste issuer public key
	pubKeyBase64 := "XWl0Q5/nJUqVwINdBolZyksAQHmxi43vJ0M9Zt1OwjE="

	pubKeyBytes, _ := base64.StdEncoding.DecodeString(pubKeyBase64)
	publicKey := ed25519.PublicKey(pubKeyBytes)

	var cred shared.Credential
	json.Unmarshal([]byte(vcJSON), &cred)

	// Remove proof before verification
	proof := cred.Proof
	cred.Proof = shared.Proof{}

	unsignedBytes, _ := json.Marshal(cred)

	sigBytes, _ := base64.StdEncoding.DecodeString(proof.JWS)

	valid := ed25519.Verify(publicKey, unsignedBytes, sigBytes)

	if valid {
		fmt.Println("✅ Credential is VALID")
	} else {
		fmt.Println("❌ Credential is INVALID")
	}
}
