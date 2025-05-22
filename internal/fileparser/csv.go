package fileparser

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/civicforge/biodata-cli/internal/logging"
	"github.com/civicforge/biodata-cli/internal/model"
)

var latCandidates = []string{"lat", "latitude", "y", "y_coord", "lat_dd", "point_lat"}
var lonCandidates = []string{"lon", "lng", "long", "longitude", "x", "x_coord", "lon_dd", "point_lon"}

func ParseCSV(path string) (model.IndexedFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return model.IndexedFile{}, err
	}
	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return model.IndexedFile{}, err
	}

	latCol, longCol, ok := detectLatLongFromHeader(headers)
	if !ok {
		logging.Error("Could not find a lat or long col")
		return model.IndexedFile{}, nil
	}

	fmt.Println(latCol, longCol)

	return model.IndexedFile{}, nil
}

func detectLatLongFromHeader(headers []string) (latCol, lonCol string, ok bool) {
	for _, h := range headers {
		lh := strings.ToLower(strings.TrimSpace(h))

		for _, latKey := range latCandidates {
			if strings.Contains(lh, latKey) {
				latCol = h
				break
			}
		}

		for _, lonKey := range lonCandidates {
			if strings.Contains(lh, lonKey) {
				lonCol = h
				break
			}
		}

		// Exit early if both found
		if latCol != "" && lonCol != "" {
			return latCol, lonCol, true
		}
	}

	return "", "", false
}
