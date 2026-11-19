package usecase

import (
	"calculator/internal/domain"
	"errors"
	"math"
	"testing"
)

const epsilon = 1e-9

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
func TestCalculate(t *testing.T) {
	tests := []struct {
		name        string
		left        float64
		right       float64
		op          domain.Operation
		expected    float64
		expectError error
	}{
		{
			name:        "Plus",
			left:        32,
			right:       43.,
			op:          domain.OpAdd,
			expected:    75.,
			expectError: nil,
		},
		{
			name:        "Minus",
			left:        32,
			right:       2.000,
			op:          domain.OpSub,
			expected:    30.000,
			expectError: nil,
		},
		{
			name:        "Float addition",
			left:        0.1,
			right:       0.2,
			op:          domain.OpAdd,
			expected:    0.3,
			expectError: nil,
		},
		{
			name:        "Division by zero",
			left:        10.0,
			right:       0.0,
			op:          domain.OpDiv,
			expected:    0.0, // Значение не важно, так как будет ошибка
			expectError: domain.ErrDivisionByZero,
		},
		{
			name:        "Modulo by zero",
			left:        10.0,
			right:       0.0,
			op:          domain.OpMod,
			expected:    0.0,
			expectError: domain.ErrDivisionByZero,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := NewCalculator()
			result, err := calc.Calculate(tt.left, tt.right, tt.op)

			if tt.expectError != nil {
				if err == nil {
					t.Errorf("ожидалась ошибка %v, но её нет", tt.expectError)
				} else if !errors.Is(err, tt.expectError) {
					t.Errorf("ожидалась ошибка %v, получили %v", tt.expectError, err)
				}
				return // важно! если ошибка — не проверяем результат
			}

			if err != nil {
				t.Errorf("не ожидали ошибку, но получили: %v", err)
				return
			}
			if !floatEquals(result, tt.expected) {
				t.Errorf("ожидалось %v, получили %v", tt.expected, result)
			}
		})
	}
}

func TestHistory(t *testing.T) {
	cals := NewCalculator()
	
	// 1. Проверка пустоты
	if len(cals.History()) != 0 {
		t.Fatalf("ожидалось 0 элементов, получили %d", len(cals.History()))
	}
	
	cals.Calculate(1, 1, domain.OpAdd)
	cals.Calculate(10, 2, domain.OpMod)
	cals.Calculate(8, 4, domain.OpDiv)

	// 2. Проверка, что добавилось 3 элемента (исправлено !=)
	if len(cals.History()) != 3 {
		t.Fatalf("ожидалось 3 элемента, получили %d", len(cals.History()))
	}
	
	cals.Calculate(10, 323.32, domain.OpAdd)
	cals.Calculate(13, 12, domain.OpMod)
	cals.Calculate(332, 42.2, domain.OpDiv)
	cals.Calculate(1332, 3231, domain.OpSub)
	cals.Calculate(132.323, -2.31, domain.OpMul)
	
	history := cals.History()

	// 3. Проверка ограничения длины (осталось 5)
	if len(history) != 5 {
		t.Errorf("ожидалось 5 элементов в истории, получили %d", len(history))
	}

	// 4. Проверка, что это ПОСЛЕДНИЕ 5 элементов
	if history[0].LeftOperand != 10 || history[0].RightOperand != 323.32 {
		t.Errorf("ожидалось, что первый элемент истории будет 10 + 323.32, а получили %v + %v", history[0].LeftOperand, history[0].RightOperand)
	}

	// 5. Проверка самого последнего элемента (исправлены = на != и опечатка)
	if history[4].LeftOperand != 132.323 || history[4].RightOperand != -2.31 {
		t.Errorf("ожидалось, что последний элемент истории будет 132.323 * (-2.31), а получили %v * %v", history[4].LeftOperand, history[4].RightOperand)
	}
}