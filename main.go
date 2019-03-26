package main

import (
	"derivation/derivation"
	"log"
)

// Test vectors:
var (
	testPubKeyStr     = "5195ac8b6933ca20d1f09114cffc89b61c6531b7c2228d03ea84dc5b944cbe8a"
	testPrivKeyStr    = "9fde8d863a3040ff67ccc07c49b55ee4746d4db410fb18bdde7dbd7ccba4180e"
	testDerivationStr = "9a1bdc439bb8446b5a7cfbbc3279bee5777336d98ba70f5c5a6f6bbbfb07d1b0"
	testPubKey        = derivation.HexToKey(testPubKeyStr)
	testPrivKey       = derivation.HexToKey(testPrivKeyStr)
	testDerivation    = derivation.HexToKey(testDerivationStr)
)

func main() {
	log.Printf("Input:\n\nPublic key: '%s'\nPrivateKey: '%s'",
		testPubKeyStr, testPrivKeyStr)
	log.Printf("Performing derivation:")
	d := derivation.KeyDerivation(&testPubKey, &testPrivKey)

	log.Printf("--> Derived key '%x'", d)
}
