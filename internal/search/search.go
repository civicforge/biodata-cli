package search

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/civicforge/biodata-cli/internal/index"
	"github.com/civicforge/biodata-cli/internal/logging"
	"github.com/civicforge/biodata-cli/internal/model"
	"github.com/spf13/cobra"
)

type FilterOptions struct {
	Format string
	CRS    string
	Path   string
	Name   string
	Field  string
}

func (f FilterOptions) Matches(file model.IndexedFile) bool {
	if f.Format != "" && !strings.EqualFold(file.Format, f.Format) {
		return false
	}
	if f.CRS != "" && !strings.EqualFold(file.CRS, f.CRS) {
		return false
	}
	if f.Path != "" && !strings.Contains(file.Path, f.Path) {
		return false
	}
	if f.Name != "" && !strings.Contains(file.Filename, f.Name) {
		return false
	}
	if f.Field != "" && !hasField(file.Fields, f.Field) {
		return false
	}
	return true
}

func hasField(fields []model.FieldMetadata, name string) bool {
	for _, f := range fields {
		if strings.EqualFold(f.Name, name) {
			return true
		}
	}
	return false
}

func Search(cmd *cobra.Command, args []string) {
	store, err := os.ReadFile(".temp_store.json")
	if err != nil {
		logging.Error(err.Error())
	}

	var indexes index.Index
	err = json.Unmarshal(store, &indexes)
	if err != nil {
		logging.Error(err.Error())
	}

	filters := FilterOptions{}
	filters.Format, _ = cmd.Flags().GetString("format")
	filters.CRS, _ = cmd.Flags().GetString("crs")
	filters.Path, _ = cmd.Flags().GetString("path")
	filters.Name, _ = cmd.Flags().GetString("name")
	filters.Field, _ = cmd.Flags().GetString("field")

	for _, idx := range indexes.IndexedFiles {
		if filters.Matches(idx) {
			data, err := json.MarshalIndent(idx, "", "  ")
			if err != nil {
				fmt.Printf("Error marshaling struct: %v\n", err)
				continue
			}
			fmt.Println(string(data))
		}
	}
}
