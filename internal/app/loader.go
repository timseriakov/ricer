package app

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"timseriakov/ricer/infra"
	"timseriakov/ricer/pkg/types"
)

// DesktopDirs — стандартные директории для поиска .desktop файлов.
var DesktopDirs = []string{
	"/usr/share/applications/",
	"~/.local/share/applications/",
	"~/.nix-profile/share/applications",
	"/run/current-system/sw/share/applications",
}

// LoadApplications загружает, фильтрует, сортирует и устраняет дубликаты приложений.
func LoadApplications(appsWeights map[string]int) []types.App {
	var apps []types.App

	// Сначала добавляем приложения с весами (часто используемые)
	for file, weight := range appsWeights {
		parsedApp, err := infra.ParseDesktopFile(file, weight)
		if err == nil && parsedApp.Name != "" {
			apps = append(apps, parsedApp)
		}
	}

	// Затем ищем все остальные .desktop файлы
	seen := make(map[string]struct{})
	for _, a := range apps {
		// Дедупликация по имени приложения, а не по пути к файлу
		seen[a.Name] = struct{}{}
	}

	for _, dir := range DesktopDirs {
		home, _ := os.UserHomeDir()
		expandedDir := strings.Replace(dir, "~", home, 1)
		files, err := filepath.Glob(filepath.Join(expandedDir, "*.desktop"))
		if err != nil {
			continue
		}

		for _, file := range files {
			parsedApp, err := infra.ParseDesktopFile(file, 0)
			if err != nil || parsedApp.Name == "" {
				continue
			}
			if _, ok := seen[parsedApp.Name]; ok {
				continue
			}
			apps = append(apps, parsedApp)
			seen[parsedApp.Name] = struct{}{}
		}
	}

	// Сортировка по весу (частоте использования)
	sort.Slice(apps, func(i, j int) bool {
		return apps[i].Weight > apps[j].Weight
	})

	return apps
}
