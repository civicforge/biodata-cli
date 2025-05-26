package indexer

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"

	"github.com/civicforge/biodata-cli/internal/fileparser"
	"github.com/civicforge/biodata-cli/internal/logging"
	"github.com/spf13/cobra"
)

var supportedFormats = map[string]struct{}{
	"csv":     {},
	"geojson": {},
	"shp":     {},
	"parquet": {},
	"gpkg":    {},
	"json":    {},
	"kml":     {},

	// Later
	// "tif":     {},
	// "tiff":    {},
	// "xlsx":    {},
	// "zip":     {},
}

var hasFormat bool
var selectedFormats []string

func Index(cmd *cobra.Command, args []string) {
	inputPath := args[0]
	verbose, _ := cmd.Flags().GetBool("verbose")
	// recursive, _ := cmd.Flags().GetBool("recursive")
	formats, _ := cmd.Flags().GetStringArray("format")

	if verbose {
		logging.EnableDebug()
	}

	logging.Info("Input path: " + inputPath)

	if len(formats) > 0 {
		for _, format := range formats {
			ok, _ := isSupportedFormat(format)
			if !ok {
				logging.Error("Invalid format passed as flag: " + format)
				return
			}
		}
		selectedFormats = formats
		hasFormat = true
	}

	err := filepath.WalkDir(inputPath, indexFile)
	if err != nil {
		logging.Error(err.Error())
		return
	}
}

func isSupportedFormat(ext string) (bool, error) {
	if ext[0] == '.' {
		ext = ext[1:]
	}

	_, ok := supportedFormats[ext]
	if ok == true {
		output := fmt.Sprintf("%s found within supported formats", ext)
		logging.Debug(output)
	}
	return ok, nil
}

func indexFile(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if d.IsDir() {
		return nil
	}

	ext := filepath.Ext(path)
	if ext != "" && ext[0] == '.' {
		ext = ext[1:]
	}

	if _, ok := supportedFormats[ext]; !ok {
		return nil
	}

	if hasFormat && !isInSelectedFormats(ext) {
		return nil
	}

	switch ext {
	case "geojson":
		idxf, err := fileparser.ParseGeoJson(path)
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", idxf)
	case "csv":
		idxf, err := fileparser.ParseCSV(path)
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", idxf)
	case "shp":
		idxf, err := fileparser.ParseShapefile(path)
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", idxf)
	default:
		logging.Error("Parsing not implemented for supported type: " + ext)
	}

	return nil
}

func isInSelectedFormats(ext string) bool {
	return slices.Contains(selectedFormats, ext)
}
