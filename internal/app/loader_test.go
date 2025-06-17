package app

import (
	"os"
	"path/filepath"
	"testing"

)

// Вспомогательная функция для создания временного .desktop файла
func createTempDesktopFile(t *testing.T, content string) string {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.desktop")
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create .desktop file: %v", err)
	}
	return filePath
}

func TestLoadApplications_Basic(t *testing.T) {
	desktopContent := `
[Desktop Entry]
Name=TestApp
Exec=testapp
Comment=Test application
`
	file := createTempDesktopFile(t, desktopContent)
	weights := map[string]int{file: 5}

	apps := LoadApplications(weights)
	if len(apps) == 0 {
		t.Fatalf("LoadApplications returned empty list")
	}
	found := false
	for _, a := range apps {
		if a.Name == "TestApp" && a.Weight == 5 {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("TestApp with weight=5 not found in apps: %+v", apps)
	}
}

func TestLoadApplications_Deduplication(t *testing.T) {
	desktopContent := `
[Desktop Entry]
Name=DupApp
Exec=dupapp
`
	file := createTempDesktopFile(t, desktopContent)
	weights := map[string]int{file: 10}

	// Добавим тот же файл в DesktopDirs через временную директорию
	origDirs := DesktopDirs
	tmpDir := t.TempDir()
	DesktopDirs = []string{tmpDir}
	defer func() { DesktopDirs = origDirs }()

	// Копируем файл во временную директорию
	dst := filepath.Join(tmpDir, "dup.desktop")
	data, _ := os.ReadFile(file)
	os.WriteFile(dst, data, 0644)

	apps := LoadApplications(weights)
	count := 0
	for _, a := range apps {
		if a.Name == "DupApp" {
			count++
		}
	}
	if count != 1 {
		t.Errorf("expected deduplication, got %d DupApp entries", count)
	}
}

func TestLoadApplications_Sorting(t *testing.T) {
	content1 := `
[Desktop Entry]
Name=App1
Exec=app1
`
	content2 := `
[Desktop Entry]
Name=App2
Exec=app2
`
	file1 := createTempDesktopFile(t, content1)
	file2 := createTempDesktopFile(t, content2)
	weights := map[string]int{file1: 1, file2: 5}

	apps := LoadApplications(weights)
	if len(apps) < 2 {
		t.Fatalf("expected at least 2 apps, got %d", len(apps))
	}
	if apps[0].Weight < apps[1].Weight {
		t.Errorf("expected sorting by weight desc, got: %+v", apps)
	}
}

func TestLoadApplications_InvalidDesktopFile(t *testing.T) {
	weights := map[string]int{"/non/existent/file.desktop": 3}
	apps := LoadApplications(weights)
	for _, a := range apps {
		if a.ID == "/non/existent/file.desktop" {
			t.Errorf("invalid .desktop file should not be in result")
		}
	}
}
