package usecase

import (
	"calculator/internal/domain"
	"math"
)

type Calculator struct {
	listOfcalculations []domain.Calculation
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Calculate(left, right float64, op domain.Operation) (float64, error) {
	var result float64
	switch op {
	case domain.OpAdd:
		result = left + right
	case domain.OpSub:
		result = left - right
	case domain.OpMul:
		result = left * right
	case domain.OpDiv:
		if right == 0 {
			return 0.0, domain.ErrDivisionByZero
		}
		result = left / right
	case domain.OpMod:
		if right == 0 {
			return 0.0, domain.ErrDivisionByZero
		}
		result = math.Mod(left, right)
	case domain.OpPow:
		result = math.Pow(left, right)
	default:
		return 0, domain.ErrUnknownOperation
	}

	c.listOfcalculations = append(c.listOfcalculations, domain.Calculation{
		LeftOperand:  left,
		Operator:     op,
		RightOperand: right,
		Result:       result,
	})
	if len(c.listOfcalculations) > 5 {
		c.listOfcalculations = c.listOfcalculations[len(c.listOfcalculations)-5:]
	}
	return result, nil
}
func (c *Calculator) History() []domain.Calculation {
	out := make([]domain.Calculation, len(c.listOfcalculations))
	copy(out, c.listOfcalculations)
	return out
}
