package main

import (
	"fmt"  // For input and output operations
	"math" // For advanced mathematical operations
)

// Main function
func main() {
	var num1, num2 float64
	var operator string

	// Display welcome message
	fmt.Println("Simple Calculator")
	fmt.Println("==================")

	// Get the first number from the user
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	// Get the operator from the user
	fmt.Print("Enter the operator (+, -, *, /, sqrt): ")
	fmt.Scanln(&operator)

	if operator == "sqrt" {
		// Calculate square root of the first number
		fmt.Printf("Result: %.2f\n", math.Sqrt(num1))
	} else {
		// Get the second number from the user
		fmt.Print("Enter the second number: ")
		fmt.Scanln(&num2)

		// Perform the calculation based on the operator
		switch operator {
		case "+":
			fmt.Printf("Result: %.2f\n", num1+num2)
		case "-":
			fmt.Printf("Result: %.2f\n", num1-num2)
		case "*":
			fmt.Printf("Result: %.2f\n", num1*num2)
		case "/":
			if num2 != 0 { // Check for division by zero
				fmt.Printf("Result: %.2f\n", num1/num2)
			} else {
				fmt.Println("Error: Division by zero is not allowed!")
			}
		default:
			fmt.Println("Error: Invalid operator!")
		}
	}

	// Exit message
	fmt.Println("Thank you for using the calculator!")
}
