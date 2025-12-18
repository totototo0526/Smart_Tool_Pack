# KSP Smart Tool Pack (KSP-STP)

**"Transparency is the best Security."**

## 📌 概要 (Overview)
現場の「めんどくさい」を解消するための、Windows用業務効率化ツールセットです。
あえてモダンなGUIアプリではなく、**枯れた技術（Batchfile）**をベースに構築しています。

なぜか？
それは、**「中身が見える（監査可能である）こと」こそが、業務ツールにおける最強のセキュリティ機能だから**です。

## 📦 収録ツール (Tools)

### 1. 作業バックアップ (Smart Backup)
ドラッグ＆ドロップされたファイルを、日時付きでバックアップします。
「あ、上書きしちゃった！」という事故を物理的に防ぎます。

* **Logic:** 対象ファイルのコピーを `_History` フォルダに生成。
* **Feature:** 世代管理機能搭載。設定した数（デフォルト10）を超えた古いバックアップは自動削除されます。

### 2. 日付付与 (Add Date)
ファイル名に自動で「現在日時」を付与してリネームします。
提出用ファイルの作成や、ログファイルの整理に最適です。

* **Logic:** `FileName.ext` -> `FileName_20251218.ext` (Suffix/Prefix切替可)

## ⚙️ 設定 (Configuration)
`setting.ini` を編集することで、すべての挙動を制御できます。
コンパイル不要で、現場の運用に合わせて即座にポリシー変更が可能です。

```ini
[基本設定]
; 日付をつける位置 (Prefix=頭 / Suffix=お尻)
Position=Suffix

; 日付のフォーマット (PowerShell準拠)
DateFormat=yyyyMMdd

; 区切り文字
Delimiter=_

[詳細設定]
; バックアップの世代管理 (0=無制限)
MaxHistory=10

---

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
