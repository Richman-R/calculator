package domain

type Calculation struct {
	LeftOperand  float64
	Operator     Operation
	RightOperand float64
	Result       float64
}
