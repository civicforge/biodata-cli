package search

import (
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/civicforge/biodata-cli/internal/logging"
	"github.com/civicforge/biodata-cli/internal/model"
)

type mod struct {
	table table.Model
}

func (m mod) Init() tea.Cmd { return nil }

func (m mod) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
			// case "enter":
			// 	return m, tea.Batch(
			// 		tea.Printf("Let's go to %s!", m.table.SelectedRow()[0]),
			// 	)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m mod) View() string {
	return m.table.View()
}

func DisplayResults(result []model.IndexedFile) {
	columns := []table.Column{
		{Title: "ID", Width: 3},
		{Title: "Format", Width: 10},
		{Title: "CRS", Width: 12},
		{Title: "# Features", Width: 10},
		{Title: "Path", Width: 40},
	}

	rows := make([]table.Row, len(result))
	for i, idxf := range result {
		rows[i] = table.Row{strconv.Itoa(idxf.ID), idxf.Format, idxf.CRS, strconv.Itoa(idxf.NumFeatures), idxf.Path}
	}

	height := min(len(rows)+3, 20)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(height),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(lipgloss.Color("240")).
		Bold(true)

	s.Cell = s.Cell.
		BorderRight(true).
		BorderForeground(lipgloss.Color("238"))

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true)

	t.SetStyles(s)

	m := mod{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		logging.Error(err.Error())
		os.Exit(1)
	}

}
