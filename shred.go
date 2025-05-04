package shred

import (
	"fmt"
	"os"
	"strings"

	ce "github.com/wprimadi/shred/algorithms/cryptographic-erase"
	"github.com/wprimadi/shred/algorithms/dod"
	dod_ece "github.com/wprimadi/shred/algorithms/dod-ece"
	"github.com/wprimadi/shred/algorithms/gutmann"
	"github.com/wprimadi/shred/algorithms/nist"
	"github.com/wprimadi/shred/algorithms/random"
	zeroonefill "github.com/wprimadi/shred/algorithms/zero-one-fill"
)

func SecureDelete(filePath, method string) error {
	method = strings.ToLower(method)

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("unable to check the file: %w", err)
	}
	if fileInfo.IsDir() {
		return fmt.Errorf("file %s is a directory, please choose a file", filePath)
	}

	switch method {
	case "gutmann":
		return gutmann.Wipe(filePath)
	case "dod":
		return dod.Wipe(filePath)
	case "dod-ece":
		return dod_ece.Wipe(filePath)
	case "nist":
		return nist.Wipe(filePath)
	case "random":
		return random.Wipe(filePath)
	case "zero-fill":
		return zeroonefill.Wipe(filePath, zeroonefill.ZeroFill, 3)
	case "one-fill":
		return zeroonefill.Wipe(filePath, zeroonefill.OneFill, 3)
	case "cryptographic-erase":
		return ce.Wipe(filePath)
	default:
		return fmt.Errorf("unknown deletion method: '%s'", method)
	}
}
