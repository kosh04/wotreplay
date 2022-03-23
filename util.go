package wotreplay

import (
	"os"
	"path/filepath"
	"strings"
)

// Expand ~tilde as home directory
func expandHomeDir(path string) (string, error) {
	if !strings.HasPrefix(path, "~/") {
		return path, nil
	}
	home, err := os.UserHomeDir()
	return filepath.Join(home, path[2:]), err
}
