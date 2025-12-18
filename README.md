# Smart Tool Pack (STP)

**"Transparency is the best Security."**

## ⚠️ Disclaimer (免責事項 - 必ずお読みください)
This project is a **personal hobby project** by the author.
It is **NOT** an official product of any company or organization the author belongs to.

本ツールは、作者個人の研究・趣味として作成されたものです。
**作者の所属する企業・組織とは一切関係ありません。**
本ツールの利用によって生じたいかなる損害についても、作者および所属組織は責任を負いません。
「コードを読み、自己責任で使う」ことに同意できる方のみご利用ください。

## 📌 概要 (Overview)
現場の「めんどくさい」を解消するための、Windows用業務効率化ツールセットです。
本リポジトリでは、配布パッケージに含まれる**2つのバージョン**のソースコードを管理しています。

1.  **Script Version (.bat):** 中身が即座に確認できる「ホワイトボックス」版
2.  **App Version (.exe):** Go言語で記述・ビルドされた「高速実行」版

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

## 🛠️ Go製アプリについて (About .exe)
配布パッケージに含まれる `.exe` ファイルは、本リポジトリの `go_src` ディレクトリ内にあるソースコードからビルドされています。
「.exeの中身も確認したい」「自分でビルドしたい」という方は、以下のソースを参照してください。

* **Source Location:**
    * Backup Tool: `./go_src/backup/`
    * AddDate Tool: `./go_src/date_append/`
    * Common Logic: `./go_src/common/`

* **Build Command (Example):**
```powershell
  cd go_src/backup
  go build -o ../../bin/BackupTool.exe main.go
```

## ⚙️ 設定 (Configuration)

`setting.ini` を編集することで、すべての挙動を制御できます。 (.bat版、.exe版 共通の設定ファイルです)

Ini, TOML

## 🛡️ セキュリティ

```
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
```

## 哲学 (Philosophy)

このリポジトリは、サプライチェーン攻撃やブラックボックス化する現代のソフトウェアへのアンチテーゼとして公開されています。

- **White Box Policy:** Script版のロジックはすべて平文で記述されています。 App版(.exe)も、このリポジトリのGoソースコードと100%一致することを保証します。
    
- **Audit First:** 利用者は、実行前に必ずコードを閲覧（Audit）してください。 「意図しない通信を行っていないか」「不正な削除コマンドがないか」を、自身の目で検品できます。
    

## 🚀 使い方 (Usage)

1.  このリポジトリをClone、またはDownloadしてください。
    
2.  `setting.ini` を環境に合わせて調整してください。
    
3.  対象のファイルを `.bat` または `.exe` にドラッグ＆ドロップしてください。
    

* * *

Author: \[Totototo0526\] License: MIT
