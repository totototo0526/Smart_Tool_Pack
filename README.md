# Smart Tool Pack

This repository contains a set of useful Windows tools written in Go.

## Included Tools

### 1. Date Append Tool
Adds a timestamp to the filename of dropped files.
- **Usage**: Drag and drop files onto `日付付与_go.exe`.
- **Settings**: Configurable via `setting.ini` (shared).

### 2. Backup Tool
Creates a backup of the file in a `_History` subdirectory with a timestamp.
- **Usage**: Drag and drop files onto `作業バックアップ_go.exe`.
- **Features**: Auto-cleans old backups based on `setting.ini`.

## Build Instructions

1. Install [Go](https://go.dev/dl/).
2. Open a terminal in `go_src` directory.
3. Run the build commands:

```bash
# Windows GUI build (no console window)
go build -ldflags -H=windowsgui -o ../日付付与_go.exe ./date_append
go build -ldflags -H=windowsgui -o ../作業バックアップ_go.exe ./backup
```

## Configuration (setting.ini)

- **DateFormat**: Timestamp format (e.g. `yyyyMMdd`).
- **Delimiter**: Separator char (e.g. `_`).
- **Position**: `Prefix` or `Suffix`.
- **MaxHistory**: Number of backups to keep (0 = unlimited).
