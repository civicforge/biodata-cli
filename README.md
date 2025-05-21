# BioData CLI

**BioData CLI** is a command-line tool for indexing, searching, and managing
local spatial and biodiversity datasets. It helps researchers, conservationists,
and data professionals make sense of fragmented files, without needing complex
GIS tools.

This project is built to support practical, high-impact work with clean, fast
tooling that works entirely offline.

---

## Why BioData CLI?

Working with spatial data is often frustrating:

- Datasets are scattered across folders, drives, and formats.
- Valuable metadata (geometry type, CRS, column structure) is hidden or missing.
- Most tooling requires heavyweight GIS software or technical expertise.

**BioData CLI** simplifies this. It's a lightweight CLI that gives you clarity
and control over your spatial files.

---

## Key Features

- **Index folders of spatial/tabular datasets** (`.geojson`, `.shp`, `.csv`, `.parquet`)
- **Extract metadata**: geometry type, CRS, row count, bounding box, column structure
- **Search** datasets by format, projection, geometry type, filename, or tags
- **Tag** datasets with custom labels for easy discovery and grouping
- **Export** the full dataset index to JSON, CSV, or SQLite
- **Run offline**, with no dependencies on cloud services

---

## Use Cases

- Quickly inventory biodiversity datasets from field teams
- Search for all files containing polygon geometries in EPSG:4326
- Tag files by project, region, or species type for future reuse
- Validate CRS and structure before loading into a database or model
- Share indexed summaries with collaborators or auditors

---

## Installation

Coming soon: prebuilt binaries for macOS, Linux, and Windows.

For now, clone the repo and build from source:

```bash
git clone https://github.com/yourusername/biodata-cli.git
cd biodata-cli
make build
````

---

## Usage

### Index a folder

```bash
biodata index ./datasets
```

### Search for GeoJSON files in WGS84

```bash
biodata search --format geojson --crs EPSG:4326
```

### Show detailed metadata for a dataset

```bash
biodata show 3
```

### Tag a dataset

```bash
biodata tag 3 --add "species" --add "africa"
```

### Export the index

```bash
biodata export --format json > export.json
```

---

## Output Example

```text
ID  Format   Geometry   CRS         Rows   Path
──  ───────  ─────────  ──────────  ─────  ─────────────────────────────────────
1   GeoJSON  Point      EPSG:4326   124    datasets/species_sightings.geojson
2   Shapefile Polygon   EPSG:3857   320    datasets/protected_areas.shp
3   CSV      -          -           88     datasets/poaching_reports.csv
```

---

## Roadmap

| Feature                | Status     |
| ---------------------- | ---------- |
| Folder indexing        | Planned    |
| Format/CRS search      | Planned    |
| Detailed metadata view | Planned    |
| Tagging                | Planned    |
| Export (JSON/CSV)      | Planned    |
| SQLite metadata store  | Planned    |
| CRS transformation     | Planned    |
| Raster file support    | Planned    |
| Plugin system          | Planned    |

---

## Frequently Asked Questions

**Q: What file types are supported?**
A: `.geojson`, `.shp`, `.csv` (with lat/lon), and `.parquet` with spatial data
are supported in v0.1. More to come.

**Q: Does it modify my data?**
A: No. All operations are read-only unless you explicitly request metadata export.

**Q: Is it offline-capable?**
A: Fully. You can run everything locally with no internet connection.

**Q: Can I use it for large datasets?**
A: Yes. It's designed to scale to thousands of files and millions of rows.

---

## Who Is This For?

BioData CLI is designed for:

- Conservation NGOs
- Academic researchers in ecology, geography, and biology
- Data engineers managing spatial pipelines
- Environmental journalists
- Open data contributors

If your work touches spatial files and you need to manage them with confidence
and speed, this tool is for you.

---

## Credits

Created and maintained by [Mathew Bravo](https://github.com/MathewBravo).
BioData CLI was inspired by the daily frustrations of working with unindexed,
poorly documented spatial data in critical domains like conservation and public research.

If you're using this tool in the field, let me know. I’d love to hear how it's helping.
