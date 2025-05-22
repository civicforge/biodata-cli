package fileparser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/civicforge/biodata-cli/internal/logging"
	"github.com/civicforge/biodata-cli/internal/model"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

func ParseGeoJson(path string) (model.IndexedFile, error) {
	info, err := os.Stat(path)
	if err != nil {
		logging.Error(err.Error())
		return model.IndexedFile{}, err
	}

	sizeInBytes := info.Size()

	dat, err := os.ReadFile(path)
	if err != nil {
		logging.Error(err.Error())
	}

	fc := geojson.NewFeatureCollection()
	err = json.Unmarshal(dat, &fc)
	if err != nil {
		logging.Error(err.Error())
	}

	featureLength := len(fc.Features)

	// TODO! need to find a way to parse the CRS for now defaulting to 4326, most common for geojson

	fieldTypes := map[string]string{}

	for _, feature := range fc.Features {
		for key, val := range feature.Properties {
			if _, exists := fieldTypes[key]; exists {
				continue
			}

			switch val.(type) {
			case string:
				fieldTypes[key] = "string"
			case float64, float32, int, int64:
				fieldTypes[key] = "number"
			case bool:
				fieldTypes[key] = "boolean"
			default:
				fieldTypes[key] = "unknown"
			}
		}
	}

	var fields []model.FieldMetadata
	for name, typ := range fieldTypes {
		fields = append(fields, model.FieldMetadata{
			Name: name,
			Type: typ,
		})
	}

	geomTypes := map[string]struct{}{}
	var bbox orb.Bound
	first := true

	for _, feature := range fc.Features {
		if feature.Geometry != nil {
			if first {
				bbox = feature.Geometry.Bound()
				first = false
			} else {
				bbox = bbox.Union(feature.Geometry.Bound())
			}
			geomTypes[feature.Geometry.GeoJSONType()] = struct{}{}
		}
	}

	bboxArray := [4]float64{
		bbox.Min[0], // min longitude
		bbox.Min[1], // min latitude
		bbox.Max[0], // max longitude
		bbox.Max[1], // max latitude
	}

	var geomList []string
	for typ := range geomTypes {
		geomList = append(geomList, typ)
	}

	extra := map[string]string{
		"geometry_types": "[" + joinQuoted(geomList) + "]",
		"bounding_box":   fmt.Sprintf("%f %f %f %f", bboxArray[0], bboxArray[1], bboxArray[2], bboxArray[3]),
	}

	indexFile := model.IndexedFile{
		Path:         path,
		Filename:     filepath.Base(path),
		Extension:    "geojson",
		SizeBytes:    sizeInBytes,
		ModifiedTime: time.Now(),
		Format:       "geojson",
		CRS:          "EPSG:4326",
		NumFeatures:  featureLength,
		Fields:       fields,
		Extra:        extra,
	}

	return indexFile, nil
}

func joinQuoted(list []string) string {
	out := ""
	for i, v := range list {
		out += `"` + v + `"`
		if i < len(list)-1 {
			out += ","
		}
	}
	return out
}
