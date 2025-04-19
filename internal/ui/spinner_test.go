package ui

import (
	"testing"
	"time"
)

func TestSpinner_StartAndStop(t *testing.T) {
	s := NewSpinner("Testing spinner")
	s.Start()

	time.Sleep(300 * time.Millisecond)

	s.Stop()
}
