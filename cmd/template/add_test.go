package template

import (
	"testing"
)

func TestAddCommand_Placeholder(t *testing.T) {
	if addcCmd == nil {
		t.Fatal("addcCmd is nil")
	}
}
