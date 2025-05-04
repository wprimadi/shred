package dod_ece

import (
	"crypto/rand"
	"errors"
	"fmt"
	"os"
)

var (
	// ErrInvalidFile indicates the target is not a regular file.
	ErrInvalidFile = errors.New("target is not a regular file")
)

// Wipe securely deletes the file at the given path using the DoD 5220.22-M (ECE) 7-pass method.
func Wipe(filePath string) error {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("failed to stat file: %w", err)
	}
	if fileInfo.IsDir() {
		return ErrInvalidFile
	}

	f, err := os.OpenFile(filePath, os.O_WRONLY, 0)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	size := fileInfo.Size()
	passes := []func(*os.File, int64) error{
		func(f *os.File, s int64) error { return overwriteWithByte(f, s, 0xF6) },
		func(f *os.File, s int64) error { return overwriteWithByte(f, s, 0x00) },
		func(f *os.File, s int64) error { return overwriteWithByte(f, s, 0xFF) },
		overwriteWithRandom,
		func(f *os.File, s int64) error { return overwriteWithByte(f, s, 0x00) },
		func(f *os.File, s int64) error { return overwriteWithByte(f, s, 0xFF) },
		overwriteWithRandom,
	}

	for i, pass := range passes {
		if err := pass(f, size); err != nil {
			return fmt.Errorf("pass %d failed: %w", i+1, err)
		}
	}

	return os.Remove(filePath)
}

func overwriteWithByte(f *os.File, size int64, b byte) error {
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}
	buf := make([]byte, 4096)
	for i := range buf {
		buf[i] = b
	}
	written := int64(0)
	for written < size {
		toWrite := min(int64(len(buf)), size-written)
		if _, err := f.Write(buf[:toWrite]); err != nil {
			return err
		}
		written += toWrite
	}
	return f.Sync()
}

func overwriteWithRandom(f *os.File, size int64) error {
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}
	buf := make([]byte, 4096)
	written := int64(0)
	for written < size {
		toWrite := min(int64(len(buf)), size-written)
		if _, err := rand.Read(buf[:toWrite]); err != nil {
			return err
		}
		if _, err := f.Write(buf[:toWrite]); err != nil {
			return err
		}
		written += toWrite
	}
	return f.Sync()
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
