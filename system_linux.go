package wotreplay

import (
	"github.com/mitchellh/go-homedir"
)

// FIXME: where directory installed?
var systemDir = "~/.wine/drive_c/Games/World_of_Tanks"

func init() {
	dir, err := homedir.Expand(systemDir)
	if err != nil {
		println(systemDir, ":", err)
		return
	}
	systemDir = dir
}
