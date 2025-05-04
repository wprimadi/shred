package random

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

// Wipe securely deletes the contents of a file by overwriting it with random data.
func Wipe(filePath string) error {
	passes := 3

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

	for i := 0; i < passes; i++ {
		if err := overwriteRandom(file, size); err != nil {
			return fmt.Errorf("failed to overwrite pass %d: %w", i+1, err)
		}
		if _, err := file.Seek(0, io.SeekStart); err != nil {
			return fmt.Errorf("failed to seek after pass %d: %w", i+1, err)
		}
	}

	return os.Remove(filePath)
}

func overwriteRandom(file *os.File, size int64) error {
	buf := make([]byte, 4096)
	var written int64
	for written < size {
		toWrite := size - written
		if toWrite > int64(len(buf)) {
			toWrite = int64(len(buf))
		}
		if _, err := rand.Read(buf[:toWrite]); err != nil {
			return err
		}
		n, err := file.Write(buf[:toWrite])
		if err != nil {
			return err
		}
		written += int64(n)
	}
	return nil
}
