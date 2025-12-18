package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	// Default settings
	DateFormatStr = "yyyyMMdd_HHmmss"
	Delimiter     = "_"
	Position      = "Suffix"
	MaxHistory    = 10
)

func ReadSettings(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "[") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		switch key {
		case "DateFormat":
			DateFormatStr = val
		case "Delimiter":
			Delimiter = val
		case "Position":
			Position = val
		case "MaxHistory":
			if n, err := strconv.Atoi(val); err == nil {
				MaxHistory = n
			}
		}
	}
}

func ConvertDateFormat(format string) string {
	r := strings.NewReplacer(
		"yyyy", "2006",
		"MM", "01",
		"dd", "02",
		"HH", "15",
		"mm", "04",
		"ss", "05",
	)
	return r.Replace(format)
}
