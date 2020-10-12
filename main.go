package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor   int
	choices  []string
	selected map[int]struct{}
}

func initialize() (tea.Model, tea.Cmd) {
	return model{
		choices:  []string{"Go", "C", "Rust", "Kotlin"},
		selected: make(map[int]struct{}),
	}, nil
}

func update(msg tea.Msg, mdl tea.Model) (tea.Model, tea.Cmd) {
	m := mdl.(model)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			if m.cursor < len(m.choices) {
				m.cursor++
			}
		case "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "enter":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func view(mdl tea.Model) string {
	m := mdl.(model)

	s := "What languages do you like?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q or etc to quit\n"

	return s
}

func main() {
	p := tea.NewProgram(initialize, update, view)
	if err := p.Start(); err != nil {
		fmt.Printf("there's been an error: %s\n", err.Error())
		os.Exit(1)
	}
}
