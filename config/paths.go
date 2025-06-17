package config

import (
	"os"
	"path/filepath"
)

// Путь к файлу закреплённых приложений
func PinnedFilePath() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "ricer", "pinned.json")
}

// Путь к файлу хранения весов приложений
func UsageFilePath() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "ricer", "ricer-drun.json")
}
