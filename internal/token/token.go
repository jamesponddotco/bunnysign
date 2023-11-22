// Package token provides functionality for generating secure tokens used for
// signing URLs.
package token

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

// Generate creates a secure token based on the provided base string. The token
// is generated using SHA256 hashing and a slightly modified version of the Base64 encoding.
func Generate(base string) string {
	var (
		hash  = sha256.Sum256([]byte(base))
		token = base64.StdEncoding.EncodeToString(hash[:])
	)

	// Replace characters to match the format expected by Bunny.net.
	token = strings.ReplaceAll(token, "+", "-")
	token = strings.ReplaceAll(token, "/", "_")
	token = strings.ReplaceAll(token, "=", "")

	return token
}
