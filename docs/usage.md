# Usage

Welcome to **BioData CLI** — your lightweight command-line tool for making
sense of messy spatial datasets.

This page walks you through how to use `biodata` in real-world workflows, what
outputs to expect, and how to start exploring your spatial data more confidently.

---

## What Is It For?

BioData CLI is purpose-built for anyone working with geospatial files who needs to:

* Rapidly understand *what data you have*
* Find relevant datasets without opening them in GIS software
* Attach meaningful tags to files you want to track or share
* Export summaries for dashboards, reports, or team review

Supported formats include:

* GeoJSON
* Shapefiles
* CSV (with geometry column)
* Parquet (with geometry column)

---

## Basic Workflow

Here's the typical flow of using BioData CLI:

### 1. Index your data

```bash
biodata index ./my-data-folder
```

This scans the directory and builds an in-memory index of all supported files,
reading geometry, CRS, format, and metadata (row counts, columns, etc.).

**Example output:**

```bash
Indexed 26 files from ./my-data-folder
```

---

### 2. Search for what you need

```bash
biodata search --geometry Polygon --crs EPSG:4326
```

This lets you filter the index to just datasets with a certain geometry type or projection.

**Example output:**

```bash
ID: 2
Path: ./my-data-folder/wetlands.geojson
Format: geojson
Geometry: Polygon
CRS: EPSG:4326
Tags: 
```

---

### 3. Tag key datasets

```bash
biodata tag 2 --add habitat high_priority
```

Use tags to track which datasets are ready for export, require review, or match
a theme.

**Example output:**

```bash
Dataset 2 updated. Tags: habitat, high_priority
```

---

### 4. Export what matters

```bash
biodata export --format json --output index.json
```

This creates a structured summary (JSON or CSV) that can be used in dashboards,
notebooks, or shared with your team.

**Example output:**

```bash
Exported 3 dataset(s) to index.json
```

---

## Typical Use Cases

* **Before fieldwork:** Get a sense of which spatial datasets you have and what
projections they’re in.
* **During analysis:** Tag datasets you're actively using and export summaries
for notebooks.
* **Before sharing:** Export just the cleaned/confirmed layers to pass on to collaborators.

---

## Output Snapshot

Here’s what metadata for a dataset looks like when viewed via `biodata show`:

```bash
ID: 4
Path: ./data/regions.shp
Format: shapefile
Geometry: Polygon
CRS: EPSG:3857
Rows: 72
Columns: name, population, geometry
Tags: review, region_layer
```

---

## Quickstart Summary

```bash
biodata index ./data
biodata search --format geojson --geometry Point
biodata tag 5 --add species observation
biodata export --format csv --output summary.csv
```

---

## Want More Control?

See the full list of supported [Commands & Flags](commands.md).

You’ll find:

* All CLI commands with descriptions
* Every flag and its default behavior
* Examples for more complex tasks

---

Start small. Grow confident. Use your data.
