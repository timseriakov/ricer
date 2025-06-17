// Минимальный main.go для запуска ricer

package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"timseriakov/ricer/internal/app"
	"timseriakov/ricer/ui"
)

func main() {
	// Загрузка весов и закреплённых приложений
	weights := app.LoadAppWeights()
	pinned := app.LoadPinnedApps()

	// Загрузка приложений
	apps := app.LoadApplications(weights)

	// Инициализация UI-модели
	model := ui.NewModel(apps, weights, pinned)

	// Запуск bubbletea-программы
	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
