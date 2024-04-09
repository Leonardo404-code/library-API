package google

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func getDownloadsDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error in get user home path: %v", err)
	}

	os := runtime.GOOS
	osSupported := []string{"darwin", "linux", "windows"}

	if isSupported(osSupported, os) {
		return filepath.Join(home, "Downloads"), nil
	} else {
		return "", fmt.Errorf("operatin system not supported: %v", err)
	}
}

func isSupported(osSupported []string, currentOs string) bool {
	for _, os := range osSupported {
		if currentOs == os {
			return true
		}
	}

	return false
}
