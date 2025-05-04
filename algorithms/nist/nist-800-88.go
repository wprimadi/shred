package nist

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

// Wipe securely deletes the contents of a file using the NIST 800-88 method.
// This method sequentially writes 0x00, 0xFF, and random data.
func Wipe(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0)
	if err != nil {
		return fmt.Errorf("unable to open current file: %w", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("cannot read current file: %w", err)
	}
	size := info.Size()

	patterns := [][]byte{
		{0x00},
		{0xFF},
		nil, // acak
	}

	for i, pattern := range patterns {
		if pattern == nil {
			if err := overwriteRandom(file, size); err != nil {
				return fmt.Errorf("unable to write random data at pass %d: %w", i+1, err)
			}
		} else {
			if err := overwritePattern(file, pattern, size); err != nil {
				return fmt.Errorf("unable to write pattern data at pass %d: %w", i+1, err)
			}
		}
		if _, err := file.Seek(0, io.SeekStart); err != nil {
			return fmt.Errorf("unable to seek a file after the pass %d: %w", i+1, err)
		}
	}

	return os.Remove(filePath)
}

func overwritePattern(file *os.File, pattern []byte, size int64) error {
	buf := make([]byte, 4096)
	for i := range buf {
		buf[i] = pattern[0]
	}
	var written int64
	for written < size {
		toWrite := size - written
		if toWrite > int64(len(buf)) {
			toWrite = int64(len(buf))
		}
		n, err := file.Write(buf[:toWrite])
		if err != nil {
			return err
		}
		written += int64(n)
	}
	return nil
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
