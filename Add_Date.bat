@echo off
setlocal enabledelayedexpansion

REM setting.ini の場所を特定（バッチファイルと同じ場所）
set "SETTINGS_FILE=%~dp0setting.ini"

REM setting.ini を読み込み
if not exist "!SETTINGS_FILE!" (
    echo setting.ini not found.

    exit /b 1
)

for /f "usebackq tokens=1,* delims==" %%A in ("!SETTINGS_FILE!") do (
    set "LINE=%%A"
    
    REM コメント行（;で始まる）を無視
    if not "!LINE:~0,1!"==";" (
        set "KEY=%%A"
        set "VAL=%%B"
        
        REM キーと値の前後空白除去（簡易版）
        REM 変数展開時の置換を利用して、タブやスペースを除去しようとすると複雑になるため
        REM 一般的な "Key = Value" の形式に対応するため、for 変数の性質を利用
        
        REM キーの空白除去
        for /f "tokens=*" %%k in ("!KEY!") do set "KEY=%%k"
        for /l %%i in (1,1,3) do if "!KEY:~-1!"==" " set "KEY=!KEY:~0,-1!"
        
        REM 値の空白除去（値がある場合のみ）
        if defined VAL (
            for /f "tokens=*" %%v in ("!VAL!") do set "VAL=%%v"
            REM 末尾の空白も除去したいが、通常は見えないので簡易的に
            
            REM 値をセット
            if defined KEY (
                set "!KEY!=!VAL!"
            )
        )
    )
)

REM 設定値の確認・デフォルト設定
if not defined DateFormat set "DateFormat=yyyyMMdd"
if not defined Delimiter set "Delimiter=_"
if not defined Position set "Position=Suffix"

REM PowerShellを使って現在の日付を取得
for /f "usebackq delims=" %%A in (`powershell -NoProfile -Command "Get-Date -Format '!DateFormat!'"`) do (
    set "CURRENT_DATE=%%A"
)

REM ドラッグ＆ドロップされたファイルを処理
:PROCESS_FILES
if "%~1"=="" goto :EOF

    set "TARGET_FILE=%~1"
    set "FILE_DIR=%~dp1"
    set "FILE_NAME=%~n1"
    set "FILE_EXT=%~x1"
    
    REM 新しいファイル名を作成
    if /i "!Position!"=="Prefix" (
        set "NEW_NAME=!CURRENT_DATE!!Delimiter!!FILE_NAME!!FILE_EXT!"
    ) else (
        set "NEW_NAME=!FILE_NAME!!Delimiter!!CURRENT_DATE!!FILE_EXT!"
    )

    set "NEW_PATH=!FILE_DIR!!NEW_NAME!"
    
    REM 同名ファイルの存在確認
    if exist "!NEW_PATH!" (
        echo Skipping "!TARGET_FILE!" ... Target "!NEW_NAME!" already exists.
    ) else (
        REM リネーム実行
        echo Renaming "!TARGET_FILE!" to "!NEW_NAME!"
        ren "!TARGET_FILE!" "!NEW_NAME!"
    )
    
    shift
    goto :PROCESS_FILES

endlocal
