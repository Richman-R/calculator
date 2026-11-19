package console

import (
	"bufio"
	"calculator/internal/domain"
	"calculator/internal/usecase"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Handler struct {
	cals    *usecase.Calculator
	scanner *bufio.Scanner
}

func NewHandler(cals *usecase.Calculator) *Handler {
	return &Handler{
		cals:    cals,
		scanner: bufio.NewScanner(os.Stdin),
	}
}
func (h *Handler) Run() {
	for {
		left, err := h.readOperand("Input left operand (or 'exit'/'history'): ")
		if errors.Is(err, ErrExit) {
			return
		}
		if err != nil {
			continue
		}

		op, err := h.readOperation()
		if errors.Is(err, ErrExit) {
			return
		}
		if err != nil {
			continue
		}

		right, err := h.readOperand("Input right operand: ")
		if errors.Is(err, ErrExit) {
			return
		}
		if err != nil {
			return
		}
		result, err := h.cals.Calculate(left, right, op)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Printf("Result: %g %s %g = %g\n", left, op, right, result)
	}
}

func (h *Handler) readOperand(prompt string) (float64, error) {
	for {
		fmt.Print(prompt)
		if !h.scanner.Scan() {
			return 0.0, ErrExit
		}
		inputStr := strings.TrimSpace(h.scanner.Text())
		if inputStr == "exit" || inputStr == "q" {
			return 0, ErrExit
		}
		if inputStr == "history" {
			h.printHistory()
			continue
		}
		num, err := strconv.ParseFloat(inputStr, 64)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		return num, nil
	}
}
func (h *Handler) readOperation() (domain.Operation, error) {
	for {
		fmt.Print("Input operation (+, -, *, /, ^, %): ")
		if !h.scanner.Scan() {
			return "", ErrExit
		}
		inputStr := strings.TrimSpace(h.scanner.Text())
		if inputStr == "exit" || inputStr == "q" {
			return "", ErrExit
		}
		if inputStr == "history" {
			h.printHistory()
			continue
		}
		op, err := domain.ParseOperation(inputStr)
		if err != nil {
			fmt.Println("invalid operation")
			continue
		}
		return op, nil
	}
}
func (h *Handler) printHistory() {
	history := h.cals.History()
	if len(history) == 0 {
		fmt.Println("История пуста.")
		return
	}
	for _, c := range history {
		fmt.Printf("%g %s %g = %g\n", c.LeftOperand, c.Operator, c.RightOperand, c.Result)
	}
}
