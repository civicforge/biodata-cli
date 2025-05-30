package model

import (
	"time"
)

type IndexedFile struct {
	ID           int               // Unique ID
	Path         string            // Full file path
	Filename     string            // Base file name
	Extension    string            // File extension (e.g., csv, shp)
	SizeBytes    int64             // File size
	ModifiedTime time.Time         // Last modified
	Format       string            // Standardized format name
	CRS          string            // Spatial reference system (if applicable)
	NumFeatures  int               // Optional: how many features (for spatial)
	Fields       []FieldMetadata   // Column info (name/type), if parsable
	Warnings     []string          // Any non-critical issues encountered
	Extra        map[string]string // Format-specific extras (optional)
}

type FieldMetadata struct {
	Name string
	Type string
}
