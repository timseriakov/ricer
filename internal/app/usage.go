package app

import (
	"timseriakov/ricer/infra"
	"timseriakov/ricer/config"
)

// Позволяет подменять путь к файлу usage.json для тестов
var usageFilePath = config.UsageFilePath

// LoadAppWeights загружает веса приложений из JSON.
func LoadAppWeights() map[string]int {
	var weights map[string]int
	infra.LoadJSON(usageFilePath(), &weights)
	if weights == nil {
		weights = make(map[string]int)
	}
	return weights
}

// SaveAppWeights сохраняет веса приложений в JSON.
func SaveAppWeights(weights map[string]int) {
	infra.SaveJSON(usageFilePath(), weights)
}

// IncAppWeight увеличивает вес приложения по ID.
func IncAppWeight(weights map[string]int, appID string) {
	weights[appID] = weights[appID] + 1
}
