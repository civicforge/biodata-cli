# Commands & Flags

## Root Command

```bash
biodata [command] [flags]
```

---

## 1. `version`

**Description:**  
Display the current version of the CLI.

**Usage:**

```bash
biodata version
```

**Output:**

```bash
BioData CLI v0.1.0
```

---

## 2. `index`

**Description:**  
Index all supported data files in a given directory, recursively.

**Usage:**

```bash
biodata index <path> [flags]
```

**Flags:**

|Flag|Description|Default|
|---|---|---|
|`--recursive, -r`|Recurse into subdirectories|true|
|`--verbose, -v`|Show all indexed files and errors|false|
|`--format`|Only include files of this format (e.g. geojson, csv)|all formats|
|`--output`|Path to save index cache or export later|in-memory only in v0.1|

---

## 3. `search`

**Description:**  
Search the in-memory index for datasets matching filter criteria.

**Usage:**

```bash
biodata search [flags]
```

**Flags:**

|Flag|Description|
|---|---|
|`--format`|Filter by format (geojson, shapefile, csv, parquet)|
|`--geometry`|Filter by geometry type (Point, Polygon, etc.)|
|`--crs`|Filter by CRS (e.g. EPSG:4326)|
|`--path`|Search by partial path match|
|`--name`|Search by filename match|
|`--tag`|Filter by one or more tags|
|`--output`|Export results to a file (CSV or JSON)|

**Example:**

```bash
biodata search --format geojson --geometry Point --crs EPSG:4326
```

---

## 4. `show`

**Description:**  
Display full metadata for a dataset in the index.

**Usage:**

```bash
biodata show <dataset-id>
```

**Flags:**

|Flag|Description|
|---|---|
|`--pretty`|Format output nicely for terminal|
|`--json`|Output metadata in JSON format|
|`--raw`|Dump all fields unformatted|

---

## 5. `tag`

**Description:**  
Add or remove tags to a dataset in the index.

**Usage:**

```bash
biodata tag <dataset-id> [flags]
```

**Flags:**

|Flag|Description|
|---|---|
|`--add`|One or more tags to add|
|`--remove`|One or more tags to remove|
|`--replace`|Replace all tags with provided|

**Examples:**

```bash
biodata tag 3 --add habitat africa
biodata tag 3 --remove outdated
```

---

## 6. `export`

**Description:**  
Export the entire index or a filtered subset to a structured file.

**Usage:**

```bash
biodata export [flags]
```

**Flags:**

|Flag|Description|
|---|---|
|`--format`|Output format: `json`, `csv`|
|`--output, -o`|Path to output file|
|`--filter`|Reuse a search expression (e.g., geojson only)|

**Example:**

```bash
biodata export --format json --output index.json
```

---

## 7. `help`

**Description:**  
Display command-specific help.

**Usage:**

```bash
biodata help <command>
```

---

## Global Flags (Available to all commands)

|Flag|Description|
|---|---|
|`--help`|Show help for any command|
|`--quiet`|Suppress non-error output|
|`--debug`|Show debug logs|
|`--version`|Show version|
