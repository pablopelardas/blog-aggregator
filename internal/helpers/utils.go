package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomString(length int) string {
    bytes := make([]byte, length/2) // because hex encoding doubles the length
    if _, err := rand.Read(bytes); err != nil {
        panic(err) // handle this error appropriately in your real code
    }
    return hex.EncodeToString(bytes)
}