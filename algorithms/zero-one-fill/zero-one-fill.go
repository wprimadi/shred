package zeroonefill

import (
	"fmt"
	"io"
	"os"
)

// Mode represents the type of byte fill pattern.
type Mode int

const (
	ZeroFill Mode = iota
	OneFill
)

// Wipe securely deletes the contents of a file by overwriting it with either zeros or ones.
func Wipe(filePath string, mode Mode, passes int) error {
	if passes < 1 {
		passes = 1
	}

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

	var pattern byte
	switch mode {
	case ZeroFill:
		pattern = 0x00
	case OneFill:
		pattern = 0xFF
	default:
		return fmt.Errorf("invalid fill mode")
	}

	for i := 0; i < passes; i++ {
		if err := overwritePattern(file, pattern, size); err != nil {
			return fmt.Errorf("failed to overwrite pass %d: %w", i+1, err)
		}
		if _, err := file.Seek(0, io.SeekStart); err != nil {
			return fmt.Errorf("failed to seek after pass %d: %w", i+1, err)
		}
	}

	return os.Remove(filePath)
}

func overwritePattern(file *os.File, pattern byte, size int64) error {
	buf := make([]byte, 4096)
	for i := range buf {
		buf[i] = pattern
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
