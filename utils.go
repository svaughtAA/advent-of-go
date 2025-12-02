package main

import "fmt"

const (
	colorReset = "\033[0m"

	colorRed   = "\033[91m"
	colorGreen = "\033[92m"
)

func PrintSuccess(message string) {
	fmt.Printf("%s%s%s\n", colorGreen, message, colorReset)
}

func PrintError(message string) {
	fmt.Printf("%s%s%s\n", colorRed, message, colorReset)
}
