package domain

type Operation string

const (
	OpAdd Operation = "+"
	OpSub Operation = "-"
	OpMul Operation = "*"
	OpDiv Operation = "/"
	OpPow Operation = "^"
	OpMod Operation = "%"
)

func ParseOperation(s string) (Operation, error) {
	switch s {
	case string(OpAdd):
		return OpAdd, nil
	case string(OpSub):
		return OpSub, nil
	case string(OpMul):
		return OpMul, nil
	case string(OpDiv):
		return OpDiv, nil
	case string(OpPow):
		return OpPow, nil
	case string(OpMod):
		return OpMod, nil
	default:
		return "", ErrUnknownOperation
	}
}
