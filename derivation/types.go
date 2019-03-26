package derivation

import "golang.org/x/crypto/ed25519"

// PrivateKey is a ed25519 elliptic curve point.
type PrivateKey = ed25519.PrivateKey

// PublicKey is the computed public keypair of the PrivateKey.
type PublicKey = ed25519.PublicKey
