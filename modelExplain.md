"credentialStatus": {
  "id": "https://issuer.com/status/42",
  "type": "BitstringStatusListEntry",
  "statusPurpose": "revocation",
  <!-- this is the id of the cred if the the issuer issued a lot of VC's and to for checking it easily this index comes into the picture -->
  "statusListIndex": "42",
  "statusListCredential": "https://issuer.com/status-list/2026"
}

BitstringStatusListEntry is a W3C Verifiable Credentials mechanism used for:

revocation
suspension
status checking

of credentials without revealing all credential details publicly.

It comes from the VC Status List specification.


The Problem It Solves

Suppose you issue 10 million credentials.

How do verifiers check:

whether a credential was revoked, suspended, expired early

without:

querying the issuer every time, exposing user identity, maintaining huge databases

The solution:

Status Lists
Basic Idea

Instead of storing revocation status per credential individually:

you create:

one giant compressed bitstring

Example:

001001000000010000...

Each bit corresponds to:

one credential
one index
Meaning
Bit	        Meaning
0	        valid
1	        revoked/suspended
Example

Suppose:

Index 42 = 1

Then credential #42 is revoked.