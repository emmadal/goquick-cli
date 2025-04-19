package ui

import (
	"fmt"
	"time"
)

// Spinner represents an animated CLI progress indicator
type Spinner struct {
	message string
	active  bool
	done    chan bool
}

// NewSpinner returns a new Spinner with a given message
func NewSpinner(message string) *Spinner {
	return &Spinner{
		message: message,
		active:  false,
		done:    make(chan bool),
	}
}

// Start begins the spinner animation in a separate goroutine
func (s *Spinner) Start() {
	if s.active {
		return
	}

	s.active = true

	go func() {
		frames := []string{"|", "/", "-", "\\"}
		ticker := time.NewTicker(100 * time.Millisecond)
		i := 0

		fmt.Printf("%s ", s.message)

		for {
			select {
			case <-ticker.C:
				fmt.Printf("\r%s %s", s.message, frames[i])
				i = (i + 1) % len(frames)
			case <-s.done:
				ticker.Stop()
				fmt.Print("\r")
				return
			}
		}
	}()
}

// Stop ends the spinner animation and prints a "done" message
func (s *Spinner) Stop() {
	if !s.active {
		return
	}

	s.active = false
	s.done <- true
	fmt.Printf("\r%s... Done!\n", s.message)
}
