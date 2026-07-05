package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input left operand:")
	leftInput, _ := reader.ReadString('\n')
	leftInput = strings.TrimSpace(leftInput)
	leftOperand, err := strconv.ParseFloat(leftInput, 64)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	var operation string
	for {
		fmt.Println("Input operation:")
		operInput, _ := reader.ReadString('\n')
		operation = strings.TrimSpace(operInput)
		if operation == "+" || operation == "-" || operation == "*" || operation == "/" {
			break
		}
		fmt.Println("Invalid input")
	}

	var rightOperand float64
	for {
		fmt.Println("Input right operand:")
		rightInput, _ := reader.ReadString('\n')
		rightInput = strings.TrimSpace(rightInput)

		rightOperand, err = strconv.ParseFloat(rightInput, 64)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		if operation == "/" && rightOperand == 0 {
			fmt.Println("MISTAKE")
			continue
		}
		break
	}

	var result float64

	switch operation {
	case "+":
		result = leftOperand + rightOperand
	case "-":
		result = leftOperand - rightOperand
	case "*":
		result = leftOperand * rightOperand
	case "/":
		result = leftOperand / rightOperand
	}
	fmt.Printf("Result:%.3f\n", result)
}
