package infra

import (
	"os"
	"path/filepath"
	"testing"
)

type storageTestStruct struct {
	A int
	B string
}

func TestSaveAndLoadJSON(t *testing.T) {
	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, "test.json")

	original := storageTestStruct{A: 42, B: "hello"}
	SaveJSON(file, original)

	var loaded storageTestStruct
	LoadJSON(file, &loaded)

	if loaded != original {
		t.Errorf("Loaded struct %+v does not match original %+v", loaded, original)
	}
}

func TestLoadJSON_FileNotExist(t *testing.T) {
	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, "no_such_file.json")

	var loaded storageTestStruct
	LoadJSON(file, &loaded) // не должно паниковать

	// loaded должен быть zero value
	if loaded != (storageTestStruct{}) {
		t.Errorf("Expected zero value, got %+v", loaded)
	}
}

func TestLoadJSON_BadJSON(t *testing.T) {
	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, "bad.json")

	os.WriteFile(file, []byte("{not json"), 0644)

	var loaded storageTestStruct
	LoadJSON(file, &loaded) // не должно паниковать

	// loaded должен быть zero value
	if loaded != (storageTestStruct{}) {
		t.Errorf("Expected zero value for bad JSON, got %+v", loaded)
	}
}
