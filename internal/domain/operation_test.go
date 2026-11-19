package domain

import (
	"errors"
	"testing"
)

func TestParseOperation(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    Operation
		expectError bool
	}{
		{
			name:        "Plus",
			input:       "+",
			expected:    OpAdd,
			expectError: false,
		},
		{
			name:        "Invalid commamd",
			input:       "s",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Sub",
			input:       "-",
			expected:    OpSub,
			expectError: false,
		},
		{
			name:        "Mul",
			input:       "*",
			expected:    OpMul,
			expectError: false,
		},
		{
			name:        "Div",
			input:       "/",
			expected:    OpDiv,
			expectError: false,
		},
		{
			name:        "Pow",
			input:       "^",
			expected:    OpPow,
			expectError: false,
		},
		{
			name:        "Mod",
			input:       "%",
			expected:    OpMod,
			expectError: false,
		},
		{
			name:        "Empty string",
			input:       "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Space",
			input:       " ",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Multiple characters",
			input:       "++",
			expected:    "",
			expectError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseOperation(tt.input)

			if tt.expectError && err == nil {
				t.Errorf("Ожидалась ошибка,но ее нет")
			}
			if !tt.expectError && result != tt.expected {
				t.Errorf("ождалось %v, получили %v", tt.expected, result)
			}
			if tt.expectError {
				if err == nil {
					t.Errorf("Ожидалась ошибка,но ее нет")
				} else if !errors.Is(err, ErrUnknownOperation) {
					t.Errorf("Ожидалась ErrUnknownOperation, получили %v", err)
				}
			}
		})
	}
}
