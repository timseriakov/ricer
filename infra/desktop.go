package infra

import (
	"bufio"
	"os"
	"strings"

	"timseriakov/ricer/pkg/types"
)

// ParseDesktopFile парсит .desktop файл и возвращает структуру App.
func ParseDesktopFile(path string, weight int) (types.App, error) {
	file, err := os.Open(path)
	if err != nil {
		return types.App{}, err
	}
	defer file.Close()

	var name, execCmd, description string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "Name=") {
			name = strings.TrimPrefix(line, "Name=")
		} else if strings.HasPrefix(line, "Exec=") {
			execCmd = strings.Trim(execCmd, `"`)
			execCmd = strings.TrimPrefix(line, "Exec=")
			execCmd = strings.Fields(execCmd)[0]
			execCmd = strings.ReplaceAll(execCmd, "%U", "")
			execCmd = strings.ReplaceAll(execCmd, "%F", "")
		} else if strings.HasPrefix(line, "Comment=") {
			description = strings.TrimPrefix(line, "Comment=")
		}

		if name != "" && execCmd != "" {
			break
		}
	}
	if description == "" {
		description = execCmd
	}

	if err := scanner.Err(); err != nil {
		return types.App{}, err
	}

	return types.App{
		ID:          path,
		Name:        name,
		Description: description,
		Command:     execCmd,
		Weight:      weight,
	}, nil
}
