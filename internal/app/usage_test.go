package app

import (
	"os"
	"path/filepath"
	"testing"
)

// Вспомогательная функция для временного usage.json
func tempUsagePath(t *testing.T) string {
	tmpDir := t.TempDir()
	return filepath.Join(tmpDir, "usage.json")
}

func TestLoadAppWeights_EmptyFile(t *testing.T) {
	path := tempUsagePath(t)
	os.WriteFile(path, []byte("{}"), 0644)

	origUsageFilePath := usageFilePath
	usageFilePath = func() string { return path }
	defer func() { usageFilePath = origUsageFilePath }()

	weights := LoadAppWeights()
	if weights == nil {
		t.Fatal("expected non-nil map")
	}
	if len(weights) != 0 {
		t.Errorf("expected empty map, got %v", weights)
	}
}

func TestSaveAndLoadAppWeights(t *testing.T) {
	path := tempUsagePath(t)

	origUsageFilePath := usageFilePath
	usageFilePath = func() string { return path }
	defer func() { usageFilePath = origUsageFilePath }()

	weights := map[string]int{"app1": 2, "app2": 5}
	SaveAppWeights(weights)

	loaded := LoadAppWeights()
	if len(loaded) != 2 || loaded["app1"] != 2 || loaded["app2"] != 5 {
		t.Errorf("Save/LoadAppWeights mismatch: got %v", loaded)
	}
}

func TestIncAppWeight(t *testing.T) {
	weights := map[string]int{"app1": 1}
	IncAppWeight(weights, "app1")
	if weights["app1"] != 2 {
		t.Errorf("expected app1 weight=2, got %d", weights["app1"])
	}
	IncAppWeight(weights, "app2")
	if weights["app2"] != 1 {
		t.Errorf("expected app2 weight=1, got %d", weights["app2"])
	}
}
