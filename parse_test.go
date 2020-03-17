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
	if err := filepath.Walk(replayDir(), walkFunc); err != nil {
		t.Error(err)
	}
	if err := filepath.Walk("testdata/replays", walkFunc); err != nil {
		t.Error(err)
	}
	t.Logf("tested %d replay files\n", tested)
}
