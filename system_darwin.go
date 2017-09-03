package wotreplay

import (
	"github.com/mitchellh/go-homedir"
)

var systemDir = "~/Library/Application Support/World of Tanks/Bottles/worldoftanks/drive_c/Games/World_of_Tanks"

func init() {
	dir, err := homedir.Expand(systemDir)
	if err != nil {
		println(systemDir, ":", err)
		return
	}
	systemDir = dir
}
