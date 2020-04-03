package wotreplay

import (
	"os"
	"path/filepath"
	"testing"
)

func testParseFile(t *testing.T, filename string) error {
	//t.Log("Testing File", filename)

	_, err := ParseFile(filename)
	if err != nil {
		t.Errorf("Parse Error: %v in %v\n", err, filename)
	}
	return nil
}

func TestParseFiles(t *testing.T) {
	tested := 0
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil || !IsReplayFile(path) {
			return err
		}
		tested++
		return testParseFile(t, path)
	}
	walk := func(dir string) {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Log("Skip Directory", err)
			return
		}
		if err := filepath.Walk(dir, walkFunc); err != nil {
			t.Error(err)
		}

	}

	walk(replayDir())
	walk("testdata/replays")

	t.Logf("tested %d replay files\n", tested)
}
