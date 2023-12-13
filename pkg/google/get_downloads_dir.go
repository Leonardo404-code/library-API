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

	switch os := runtime.GOOS; os {
	case "darwin":
		return filepath.Join(home, "Downloads"), nil
	case "linux":
		return filepath.Join(home, "Downloads"), nil
	case "windows":
		return filepath.Join(home, "Downloads"), nil
	default:
		return "", fmt.Errorf("operatin system not supported: %v", err)
	}
}
