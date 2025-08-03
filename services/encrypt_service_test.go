package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptCaesarShift(t *testing.T) {
	svc := &EncryptServiceImpl{}

	tests := []struct {
		name     string
		input    string
		shift    int
		expected string
	}{
		{"Simple shift", "abc123", 3, "def456"},
		{"Wrap around z", "xyz789", 5, "cde234"},
		{"With hyphen", "hello-world", 1, "ifmmp-xpsme"},
		{"Negative shift", "def456", -3, "abc123"},
		{"Zero shift", "sample", 0, "sample"},
		{"Mixed case", "HelloWorld", 1, "IfmmpXpsme"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := svc.EncryptCaesarShift(tt.input, tt.shift)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
