package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func inputNum(scanner *bufio.Scanner, prompt string) (float64, error) {
	for {
		fmt.Println(prompt)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return 0., err
			}
			return 0.0, io.EOF
		}
		text := strings.TrimSpace(scanner.Text())
		num, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		return num, err
	}
}
func inputOperator(scanner *bufio.Scanner, prompt string) (string, error) {

	fmt.Println(prompt)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}
	text := strings.TrimSpace(scanner.Text())
	return text, nil

}
func Operator(scanner *bufio.Scanner, prompt string) (string, error) {
	for {
		op, err := inputOperator(scanner, prompt)
		if err != nil {
			return "", err
		}
		switch op {
		case "+", "-", "*", "/", "**", "^", "%":
			return op, nil
		}
		fmt.Printf("Invalid operator %q, allowed: + - * / ** ^ %%\n", op)
	}
}
func calc(lOperand, rOperand float64, op string) (float64, error) {
	switch op {
	case "+":
		return lOperand + rOperand, nil
	case "-":
		return lOperand - rOperand, nil
	case "*":
		return lOperand * rOperand, nil
	case "/":
		switch rOperand {
		case 0:
			return 0.0, errors.New("division by zero")
		default:
			return lOperand / rOperand, nil
		}
	case "**", "^":
		return math.Pow(lOperand, rOperand), nil
	case "%":
		return math.Mod(lOperand, rOperand), nil
	default:
		return 0, fmt.Errorf("unknown operator %q", op)
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	leftOperand, err := inputNum(scanner, "Enter left number:")
	if err != nil {
		fmt.Println(err)
		return
	}
	op, err := Operator(scanner, "Enter Operator:")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	rightOperand, err := inputNum(scanner, "Enter right number:")
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := calc(leftOperand, rightOperand, op)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("%g %s %g = %g\n", leftOperand, op, rightOperand, result)
}
