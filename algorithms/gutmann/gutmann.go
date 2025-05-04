package gutmann

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
)

// ErrInvalidFile is returned when the provided file is invalid.
var ErrInvalidFile = errors.New("invalid file")

// GutmannPasses defines the full 35-pass Gutmann method pattern.
var GutmannPasses = [][]byte{
	{0x55}, {0xAA}, {0x92, 0x49, 0x24}, {0x49, 0x24, 0x92}, {0x24, 0x92, 0x49},
	{0x00}, {0x11}, {0x22}, {0x33}, {0x44}, {0x55}, {0x66}, {0x77}, {0x88}, {0x99}, {0xAA}, {0xBB}, {0xCC}, {0xDD}, {0xEE}, {0xFF},
	{0x92, 0x49, 0x24}, {0x49, 0x24, 0x92}, {0x24, 0x92, 0x49},
	{0x6D, 0xB6, 0xDB}, {0xB6, 0xDB, 0x6D}, {0xDB, 0x6D, 0xB6},
	{0x00}, {0xFF},
	nil, nil, nil, nil, nil, // 5 passes with random data
}

// Wipe securely overwrites a file using the Gutmann method.
func Wipe(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil || info.IsDir() {
		return ErrInvalidFile
	}

	size := info.Size()
	buffer := make([]byte, size)

	for i, pattern := range GutmannPasses {
		if pattern == nil {
			if _, err := rand.Read(buffer); err != nil {
				return fmt.Errorf("failed to generate random data: %w", err)
			}
		} else {
			fillPattern(buffer, pattern)
		}

		if _, err := file.Seek(0, io.SeekStart); err != nil {
			return fmt.Errorf("seek failed on pass %d: %w", i, err)
		}
		if _, err := file.Write(buffer); err != nil {
			return fmt.Errorf("write failed on pass %d: %w", i, err)
		}
		if err := file.Sync(); err != nil {
			return fmt.Errorf("sync failed on pass %d: %w", i, err)
		}
	}

	return nil
}

// fillPattern fills the buffer with the repeating pattern.
func fillPattern(buf []byte, pattern []byte) {
	for i := range buf {
		buf[i] = pattern[i%len(pattern)]
	}
}
