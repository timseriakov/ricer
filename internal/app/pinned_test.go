package app

import (
	"os"
	"path/filepath"
	"testing"
)

// Вспомогательная функция для временного pinned.json
func tempPinnedPath(t *testing.T) string {
	tmpDir := t.TempDir()
	return filepath.Join(tmpDir, "pinned.json")
}

func TestLoadPinnedApps_EmptyFile(t *testing.T) {
	path := tempPinnedPath(t)
	os.WriteFile(path, []byte("{}"), 0644)

	origPinnedFilePath := pinnedFilePath
	pinnedFilePath = func() string { return path }
	defer func() { pinnedFilePath = origPinnedFilePath }()

	pinned := LoadPinnedApps()
	if pinned == nil {
		t.Fatal("expected non-nil map")
	}
	if len(pinned) != 0 {
		t.Errorf("expected empty map, got %v", pinned)
	}
}

func TestSaveAndLoadPinnedApps(t *testing.T) {
	path := tempPinnedPath(t)

	origPinnedFilePath := pinnedFilePath
	pinnedFilePath = func() string { return path }
	defer func() { pinnedFilePath = origPinnedFilePath }()

	pinned := map[string]bool{"app1": true, "app2": false}
	SavePinnedApps(pinned)

	loaded := LoadPinnedApps()
	if len(loaded) != 2 || !loaded["app1"] || loaded["app2"] {
		t.Errorf("Save/LoadPinnedApps mismatch: got %v", loaded)
	}
}
