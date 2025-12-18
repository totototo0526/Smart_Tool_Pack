@echo off
setlocal enabledelayedexpansion

REM ========================================================
REM  KSP スナップショット作成ツール (バックアップ版)
REM ========================================================

REM 1. 設定ファイルの読み込み
set "SETTINGS_FILE=%~dp0setting.ini"

REM デフォルト設定
set "DateFormat=yyyyMMdd"
set "Delimiter=_"
set "Position=Suffix"
set "MaxHistory=10"

if not exist "!SETTINGS_FILE!" (
    echo [ERROR] setting.ini が見つかりません。
    echo バッチファイルと同じ場所に置いてください。
    echo.
    echo [INFO] 10秒後に自動的に閉じます...
    timeout /t 10
    exit /b 1
)

REM setting.ini の解析
for /f "usebackq tokens=1,* delims==" %%A in ("!SETTINGS_FILE!") do (
    set "LINE=%%A"
    REM コメント行チェック
    if not "!LINE:~0,1!"==";" (
        if not "%%B"=="" (
            set "KEY=%%A"
            set "VAL=%%B"
            REM 簡易的な空白除去をして変数をセット
            set "!KEY: =!=!VAL: =!"
        )
    )
)

REM 強制的に時間付きフォーマットを使用（世代管理のため）
set "DateFormat=yyyyMMdd_HHmmss"

REM 2. PowerShellで日付を取得
for /f "usebackq delims=" %%A in (`powershell -NoProfile -Command "Get-Date -Format '!DateFormat!'"`) do (
    set "CURRENT_DATE=%%A"
)

REM 3. ドラッグ＆ドロップされたファイルを処理
:PROCESS_FILES
if "%~1"=="" goto :EOF

    set "TARGET_FILE=%~1"
    set "FILE_DIR=%~dp1"
    set "FILE_NAME=%~n1"
    set "FILE_EXT=%~x1"

    REM バックアップ保存先フォルダ（元の場所\_History）
    set "BACKUP_DIR=!FILE_DIR!_History"
    
    REM フォルダがなければ作成
    if not exist "!BACKUP_DIR!" (
        mkdir "!BACKUP_DIR!"
    )

    REM 新しいファイル名を作成
    if /i "!Position!"=="Prefix" (
        set "NEW_NAME=!CURRENT_DATE!!Delimiter!!FILE_NAME!!FILE_EXT!"
    ) else (
        set "NEW_NAME=!FILE_NAME!!Delimiter!!CURRENT_DATE!!FILE_EXT!"
    )

    set "DEST_PATH=!BACKUP_DIR!\!NEW_NAME!"
    
    REM コピー実行
    echo [Backup] "!TARGET_FILE!"
    echo       -> "!DEST_PATH!"
    copy "!TARGET_FILE!" "!DEST_PATH!" >nul

    REM ▼▼▼ 世代管理（古いファイルの削除）▼▼▼
    
    REM 設定が0なら何もしない
    if "!MaxHistory!"=="0" goto :SKIP_CLEANUP

    pushd "!BACKUP_DIR!"
    
    REM 検索パターン: 元ファイル名 + 区切り文字 + * + 拡張子
    REM これで "見積書_2025..." も "見積書_2026..." もまとめて新しい順に並びます
    
    set "COUNT=0"
    for /f "delims=" %%F in ('dir /b /o:-n "!FILE_NAME!!Delimiter!*!FILE_EXT!" 2^>nul') do (
        set /a COUNT+=1
        if !COUNT! gtr !MaxHistory! (
            echo [Clean] 古い履歴を削除: "%%F"
            del "%%F"
        )
    )
    
    popd

    :SKIP_CLEANUP
    REM ▲▲▲ 世代管理ここまで ▲▲▲

    echo.
    shift
    goto :PROCESS_FILES

endlocal

REM 処理が終わったらキー入力待ち（何が起きたか見えるように）
echo.
echo [INFO] 30秒後に自動的に閉じます...
timeout /t 30
