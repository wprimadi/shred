package ce

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"
)

// Wipe simulates Cryptographic Erase by encrypting the file with a random key and then deleting the key.
// In real scenarios, this assumes the encryption key is securely discarded, rendering the data unreadable.
func Wipe(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat file: %w", err)
	}
	size := info.Size()

	// Generate pseudo encryption key (not used to actually encrypt for this simulation)
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}

	// Derive pseudo nonce/hash to simulate unique overwrite
	hash := sha256.Sum256(key)
	hexHash := hash[:]

	// Overwrite file with pseudo encrypted-like data
	buf := make([]byte, 4096)
	var written int64
	for written < size {
		toWrite := size - written
		if toWrite > int64(len(buf)) {
			toWrite = int64(len(buf))
		}
		copy(buf[:toWrite], hexHash[:toWrite])
		n, err := file.Write(buf[:toWrite])
		if err != nil {
			return fmt.Errorf("failed to write simulated encrypted data: %w", err)
		}
		written += int64(n)
	}

	// Simulate deletion of encryption key (not stored)
	for i := range key {
		key[i] = 0
	}

	return os.Remove(filePath)
}
