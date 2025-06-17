package app

import (
	"timseriakov/ricer/infra"
	"timseriakov/ricer/config"
)

var pinnedFilePath = config.PinnedFilePath

// LoadPinnedApps загружает map[string]bool из pinned.json.
func LoadPinnedApps() map[string]bool {
	var pinned map[string]bool
	infra.LoadJSON(pinnedFilePath(), &pinned)
	if pinned == nil {
		pinned = make(map[string]bool)
	}
	return pinned
}

// SavePinnedApps сохраняет map[string]bool в pinned.json.
func SavePinnedApps(pinned map[string]bool) {
	infra.SaveJSON(pinnedFilePath(), pinned)
}
