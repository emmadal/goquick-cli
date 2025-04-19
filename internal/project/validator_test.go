package project

import (
	"testing"
)

func TestValidateProjectName(t *testing.T) {
	v := NewValidator()

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid_name", "myproject", false},
		{"empty_name", "", true},
		{"invalid_chars", "my project!", true},
		{"existing_dir", ".", true}, // this will pass if run in any folder with go.mod
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.ValidateProjectName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateProjectName(%q) = %v, wantErr = %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateControllerName(t *testing.T) {
	v := NewValidator()

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid_lowercase", "user", false},
		{"empty", "", true},
		{"invalid_chars", "user-controller!", true},
		{"camelCase", "UserController", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.ValidateControllerName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateControllerName(%q) = %v, wantErr = %v", tt.input, err, tt.wantErr)
			}
		})
	}
}
