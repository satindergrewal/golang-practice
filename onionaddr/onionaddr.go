package main

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base32"
	"fmt"
	"strings"

	"golang.org/x/crypto/sha3"
)

func main() {
	// Generate key pair
	// r := strings.NewReader("")
	// fmt.Println(len("testtesttesttesttesttesttesttest"))
	fmt.Println(ed25519.NewKeyFromSeed([]byte("testtesttesttesttesttesttesttest")))
	publicKey, _, _ := ed25519.GenerateKey(nil)

	// checksum = H(".onion checksum" || pubkey || version)
	var checksumBytes bytes.Buffer
	checksumBytes.Write([]byte(".onion checksum"))
	checksumBytes.Write([]byte(publicKey))
	checksumBytes.Write([]byte{0x03})
	checksum := sha3.Sum256(checksumBytes.Bytes())

	// onion_address = base32(pubkey || checksum || version)
	var onionAddressBytes bytes.Buffer
	onionAddressBytes.Write([]byte(publicKey))
	onionAddressBytes.Write([]byte(checksum[:2]))
	onionAddressBytes.Write([]byte{0x03})
	onionAddress := base32.StdEncoding.EncodeToString(onionAddressBytes.Bytes())

	fmt.Println("publicKey:", publicKey)
	fmt.Println("onionAddress:", strings.ToLower(onionAddress)+".onion")
}
