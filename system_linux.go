package wotreplay

// FIXME: where directory installed?
var systemDir = "~/.wine/drive_c/Games/World_of_Tanks"

func init() {
	dir, err := expandHomeDir(systemDir)
	if err != nil {
		println(systemDir, ":", err)
		return
	}
	systemDir = dir
}
