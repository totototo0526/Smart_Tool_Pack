//go:build !windows

package common

import "fmt"

func ShowMessage(title, message string) {
	fmt.Printf("[%s] %s\n", title, message)
}

func ShowError(title, message string) {
	fmt.Printf("[ERROR][%s] %s\n", title, message)
}
