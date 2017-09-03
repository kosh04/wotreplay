package wotreplay

import "path/filepath"

const extension = ".wotreplay"

// IsReplayFile ask PATH has ".wotreplay" extention
func IsReplayFile(path string) bool {
	return filepath.Ext(path) == extension
}

// WoT replay directory
func replayDir() string {
	return filepath.Join(systemDir, "replays/")
}
