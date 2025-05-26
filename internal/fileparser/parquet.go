package fileparser

import (
	"os"
	"path/filepath"

	"github.com/civicforge/biodata-cli/internal/logging"
	"github.com/civicforge/biodata-cli/internal/model"
	"github.com/parquet-go/parquet-go"
)

func ParseParquet(path string) (model.IndexedFile, error) {
	info, err := os.Stat(path)
	if err != nil {
		logging.Error(err.Error())
		return model.IndexedFile{}, err
	}

	sizeInBytes := info.Size()
	modTime := info.ModTime()

	f, err := os.Open(path)
	if err != nil {
		return model.IndexedFile{}, err
	}
	defer f.Close()

	pf, err := parquet.OpenFile(f, sizeInBytes)
	if err != nil {
		return model.IndexedFile{}, err
	}

	schema := pf.Schema()
	var fields []model.FieldMetadata
	for _, field := range schema.Fields() {
		fields = append(fields, model.FieldMetadata{
			Name: field.Name(),
			Type: field.Type().String(),
		})
	}

	indexFile := model.IndexedFile{
		Path:         path,
		Filename:     filepath.Base(path),
		Extension:    filepath.Ext(path)[1:], // remove dot
		SizeBytes:    sizeInBytes,
		ModifiedTime: modTime,
		Format:       "parquet",
		CRS:          "N/A", // Not spatial, safe to leave as-is or blank
		NumFeatures:  0,     // Could read 1 row group and count rows if needed
		Fields:       fields,
		Warnings:     []string{},
		Extra:        map[string]string{},
	}

	return indexFile, nil
}
