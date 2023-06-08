package utils_test

import (
	"testing"

	"github.com/cubexsone/fiber-backend-template/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestToDefault(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		defaultInput string
		expected     string
	}{
		{"Empty check value", "", "no value", "no value"},
		{"Empty default value", "", "", ""},
		{"Value given", "Value", "", "Value"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defaultedValue := utils.ToDefault(test.input, test.defaultInput)
			assert.Equal(t, defaultedValue, test.expected)
		})
	}
}
