package index

import (
	"encoding/json"
	"os"

	"github.com/civicforge/biodata-cli/internal/model"
)

type Index struct {
	IndexedFiles []model.IndexedFile
}

func SaveIndexToJson(index Index, path string) error {
	jsonIdx, err := json.Marshal(index)

	writeTo := path + ".temp_store.json"

	err = os.WriteFile(writeTo, jsonIdx, 0666)
	if err != nil {
		return err
	}

	return nil
}
