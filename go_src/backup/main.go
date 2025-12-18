package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"backup_tool/go_src/common"
)

func main() {
	if len(os.Args) < 2 {
		common.ShowMessage("Info", "Drag and drop files to backup.")
		return
	}

	// Read settings
	exePath, _ := os.Executable()
	settingPath := filepath.Join(filepath.Dir(exePath), "setting.ini")
	if _, err := os.Stat(settingPath); os.IsNotExist(err) {
		settingPath = "../setting.ini" // Try parent if running from subfolder in dev
	}
	// Fallback check for dev environment
	if _, err := os.Stat(settingPath); os.IsNotExist(err) {
		settingPath = "setting.ini"
	}

	common.ReadSettings(settingPath)

	// Force timestamp format for backup consistency
	goTimeFormat := common.ConvertDateFormat("yyyyMMdd_HHmmss")
	currentTime := time.Now().Format(goTimeFormat)

	var successCount int
	var errorCount int

	for _, srcPath := range os.Args[1:] {
		if err := processFile(srcPath, currentTime); err != nil {
			errorCount++
		} else {
			successCount++
		}
	}

	msg := fmt.Sprintf("Backup Complete.\nSuccess: %d\nErrors: %d", successCount, errorCount)
	common.ShowMessage("Backup Tool", msg)
}

func processFile(srcPath string, timestamp string) error {
	dir := filepath.Dir(srcPath)
	filename := filepath.Base(srcPath)
	ext := filepath.Ext(filename)
	nameOnly := strings.TrimSuffix(filename, ext)

	backupDir := filepath.Join(dir, "_History")
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		common.ShowError("Error", fmt.Sprintf("Error creating directory: %v", err))
		return err
	}

	// Always use suffix for backup files
	newName := nameOnly + common.Delimiter + timestamp + ext
	destPath := filepath.Join(backupDir, newName)

	if err := copyFile(srcPath, destPath); err != nil {
		common.ShowError("Error", fmt.Sprintf("Failed to backup %s: %v", filename, err))
		return err
	}

	cleanOldGenerations(backupDir, nameOnly, ext)
	return nil
}

func cleanOldGenerations(dir, nameOnly, ext string) {
	if common.MaxHistory <= 0 {
		return
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	var backups []os.DirEntry
	prefix := nameOnly + common.Delimiter

	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasPrefix(name, prefix) && strings.HasSuffix(name, ext) {
			backups = append(backups, e)
		}
	}

	sort.Slice(backups, func(i, j int) bool {
		return backups[i].Name() > backups[j].Name()
	})

	if len(backups) > common.MaxHistory {
		for _, e := range backups[common.MaxHistory:] {
			fullPath := filepath.Join(dir, e.Name())
			os.Remove(fullPath)
		}
	}
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
