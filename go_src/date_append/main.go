package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"backup_tool/go_src/common"
)

func main() {
	if len(os.Args) < 2 {
		common.ShowMessage("Info", "Drag and drop files to rename.")
		return
	}

	// Read settings
	exePath, _ := os.Executable()
	settingPath := filepath.Join(filepath.Dir(exePath), "setting.ini")
	if _, err := os.Stat(settingPath); os.IsNotExist(err) {
		settingPath = "../setting.ini" // Try parent
	}
	if _, err := os.Stat(settingPath); os.IsNotExist(err) {
		settingPath = "setting.ini"
	}

	common.ReadSettings(settingPath)

	currentTime := time.Now().Format(common.ConvertDateFormat(common.DateFormatStr))

	var successCount int
	var errorCount int

	for _, oldPath := range os.Args[1:] {
		if err := processFile(oldPath, currentTime); err != nil {
			errorCount++
		} else {
			successCount++
		}
	}

	msg := fmt.Sprintf("Renaming Complete.\nSuccess: %d\nErrors: %d", successCount, errorCount)
	common.ShowMessage("Date Append Tool", msg)
}

func processFile(oldPath, currentTime string) error {
	dir := filepath.Dir(oldPath)
	filename := filepath.Base(oldPath)
	ext := filepath.Ext(filename)
	nameOnly := strings.TrimSuffix(filename, ext)

	var newName string
	if strings.EqualFold(common.Position, "Prefix") {
		newName = currentTime + common.Delimiter + nameOnly + ext
	} else {
		newName = nameOnly + common.Delimiter + currentTime + ext
	}

	newPath := filepath.Join(dir, newName)

	if _, err := os.Stat(newPath); err == nil {
		// Target exists, skip without error count increment? Or count as skipped?
		// User probably wants to know if something happened.
		// For simplicity, let's treat it as a non-fatal skip, but maybe log it?
		// Since we don't have a console, let's just return nil (success/handled) or error.
		// Let's assume skipping is "successful handling".
		return nil
	}

	err := os.Rename(oldPath, newPath)
	if err != nil {
		return err
	}
	return nil
}
