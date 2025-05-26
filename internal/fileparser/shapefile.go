package fileparser

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"github.com/civicforge/biodata-cli/internal/logging"
	"github.com/civicforge/biodata-cli/internal/model"
	"github.com/jonas-p/go-shp"
)

func ParseShapefile(path string) (model.IndexedFile, error) {
	info, err := os.Stat(path)
	if err != nil {
		logging.Error(err.Error())
		return model.IndexedFile{}, err
	}

	sizeInBytes := info.Size()

	shape, err := shp.Open(path)
	if err != nil {
		return model.IndexedFile{}, err
	}
	defer shape.Close()

	fields := shape.Fields()

	// Extract fields only once
	var fieldMetadata []model.FieldMetadata
	for _, f := range fields {
		fieldMetadata = append(fieldMetadata, model.FieldMetadata{
			Name: string(f.String()),
			Type: string(f.Fieldtype),
		})
	}

	// Track geometry types
	geometryTypes := map[string]struct{}{}
	for shape.Next() {
		_, geom := shape.Shape()
		if geom != nil {
			geomType := reflect.TypeOf(geom).Elem().Name()
			geometryTypes[geomType] = struct{}{}
		}
	}

	typeList := []string{}
	for t := range geometryTypes {
		typeList = append(typeList, t)
	}
	sort.Strings(typeList)
	extra := map[string]string{
		"geometry_types": strings.Join(typeList, ","),
	}

	indexFile := model.IndexedFile{
		Path:         path,
		Filename:     filepath.Base(path),
		Extension:    "shp",
		SizeBytes:    sizeInBytes,
		ModifiedTime: info.ModTime(),
		Format:       "shapefile",
		// once I set up zip file reading this needs to be extracted from .prj file
		CRS:         "EPSG:4326",
		NumFeatures: len(fields),
		Fields:      fieldMetadata,
		Warnings:    []string{},
		Extra:       extra,
	}

	return indexFile, nil
}
