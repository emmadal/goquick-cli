package internal

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	RedColor   = "\033[31m"
	ResetColor = "\033[0m"
	BlackBold  = "\033[30;1m"
	GreenBold  = "\033[32;1m"
)

func Spinner(message string, maxBytes int64) {
	symbols := []string{"-", "\\", "|", "/"}
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	// Ensure we clear the line when done and reset color
	defer fmt.Fprintf(os.Stdout, "\r%s\r", strings.Repeat(" ", len(message)+1))

	for i := 0; i < int(maxBytes); i++ {
		for _, symbol := range symbols {
			// Print the spinner symbol in red
			fmt.Fprintf(os.Stdout, "%s%s%s\r%s", RedColor, symbol, ResetColor, message)
			<-ticker.C
		}
	}
}
