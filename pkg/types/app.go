package types

// App — структура приложения для лаунчера.
type App struct {
	ID          string // путь к .desktop (уникальный идентификатор)
	Name        string
	Description string
	Command     string
	Weight      int
	Pinned      bool
}
