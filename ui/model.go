package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"timseriakov/ricer/pkg/types"
)

// Model — основная модель TUI.
type Model struct {
	List     list.Model
	Apps     []types.App
	Weights  map[string]int
	Pinned   map[string]bool
	Filtered []types.App
	Search   string
	Searching bool
}

// NewModel создает новую модель для UI.
func NewModel(apps []types.App, weights map[string]int, pinned map[string]bool) Model {
	width, height := 80, 24 // Можно доработать получение размера терминала
	l := list.New(convertToListItems(apps), list.NewDefaultDelegate(), width-5, height)
	l.Title = "Apps"
	l.SetShowHelp(true)
	return Model{
		List:     l,
		Apps:     apps,
		Weights:  weights,
		Pinned:   pinned,
		Filtered: apps,
	}
}

func convertToListItems(apps []types.App) []list.Item {
	items := make([]list.Item, len(apps))
	for i, a := range apps {
		items[i] = appListItem{App: a}
	}
	return items
}

// appListItem — обёртка для types.App, реализующая интерфейс list.Item.
type appListItem struct {
	App types.App
}

func (a appListItem) Title() string       { return a.App.Name }
func (a appListItem) Description() string { return a.App.Description }
func (a appListItem) FilterValue() string { return a.App.Name }

// Реализация tea.Model (Init, Update, View) — заглушки для примера.
// Необходимо доработать под ваши нужды.

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Здесь должна быть логика обработки событий (поиск, запуск, пин и т.д.)
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.List.View()
}
