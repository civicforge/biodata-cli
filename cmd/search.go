package cmd

import (
	"github.com/civicforge/biodata-cli/internal/search"
	"github.com/spf13/cobra"
)

func init() {
	searchCmd.PersistentFlags().StringP("format", "", "", "Filter by format (geojson, shapefile, csv, parquet)")
	searchCmd.PersistentFlags().StringP("geometry", "", "", "Filter by geometry type (Point, Polygon, etc.)")
	searchCmd.PersistentFlags().StringP("crs", "", "", "Filter by CRS (e.g. EPSG:4326)")
	searchCmd.PersistentFlags().StringP("path", "", "", "Search by partial path match")
	searchCmd.PersistentFlags().StringP("name", "", "", "Search by filename match")
	searchCmd.PersistentFlags().StringP("tag", "", "", "Filter by one or more tags")
	searchCmd.PersistentFlags().StringP("output", "", "", "Export results to a file (CSV or JSON)")
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search the in-memory index for datasets matching filter criteria.",
	Run:   search.Search,
}
