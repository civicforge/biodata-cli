package show

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/civicforge/biodata-cli/internal/model"
)

func showPretty(idx model.IndexedFile) {
	labelStyle := lipgloss.NewStyle().Bold(true).MarginRight(3)
	valueStyle := lipgloss.NewStyle()

	var b strings.Builder

	writeLine := func(label, value string) {
		b.WriteString(labelStyle.Render(label) + valueStyle.Render(value) + "\n")
	}

	writeLine("ID:", fmt.Sprintf("%d", idx.ID))
	writeLine("Path:", idx.Path)
	writeLine("Filename:", idx.Filename)
	writeLine("Extension:", idx.Extension)
	writeLine("Size (bytes):", fmt.Sprintf("%d", idx.SizeBytes))
	writeLine("Modified:", idx.ModifiedTime.Format(time.RFC3339))
	writeLine("Format:", idx.Format)
	writeLine("CRS:", idx.CRS)
	writeLine("Num Features:", fmt.Sprintf("%d", idx.NumFeatures))
	b.WriteString("\n")

	if len(idx.Fields) > 0 {
		b.WriteString(labelStyle.Render("Fields:") + "\n")
		columns := []table.Column{
			{Title: "Field Name", Width: 40},
			{Title: "Type", Width: 15},
		}
		var rows []table.Row
		for _, f := range idx.Fields {
			rows = append(rows, table.Row{f.Name, f.Type})
		}

		height := len(rows) + 2
		if height == 0 {
			height = 1
		}

		t := table.New(table.WithColumns(columns), table.WithRows(rows), table.WithHeight(height))
		styles := table.DefaultStyles()
		styles.Selected = lipgloss.NewStyle()
		t.SetStyles(styles)
		b.WriteString(t.View() + "\n\n")
	}

	if len(idx.Warnings) > 0 {
		b.WriteString(labelStyle.Render("Warnings:") + "\n")
		for _, w := range idx.Warnings {
			b.WriteString("- " + valueStyle.Render(w) + "\n")
		}
	}

	if len(idx.Extra) > 0 {
		b.WriteString(labelStyle.Render("Extra:") + "\n")
		for k, v := range idx.Extra {
			b.WriteString(fmt.Sprintf("%s: %s\n", labelStyle.Render(k), valueStyle.Render(v)))
		}
	}

	fmt.Println(b.String())
}
