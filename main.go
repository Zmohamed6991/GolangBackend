package main

import (
	"fmt"
	//"strconv"
)
	

func main() {

	/* myAge := 27
	var myName string = "Zagel"
	
		fmt.Printf("my name is %v\n", myName)
		fmt.Printf("i am %d years old\n", myAge)

	 var hobby string 
	 
	 var holiday string

	 fmt.Println("Please enter your favourite hobby")

	 fmt.Scan(&hobby)

	 fmt.Println(hobby, "is a very cool hobby")



	 fmt.Println("where do you like to go on holiday?")

	 fmt.Scan(&holiday)

	 */

/*
	 var bankBalance int
	 const housePrice int = 1900

	 for { 


	 fmt.Println("What is your bank balance?")
	 fmt.Scan(&bankBalance)

	 //fmt.Println("please enter house price:")
	 //fmt.Scan(&housePrice)

	 if  bankBalance > housePrice {
	 fmt.Println("congratulations you can buy it")

	 } else if housePrice > bankBalance {
		fmt.Println("Insufficient funds")

	 } else {
		fmt.Println("wait")
	 }

	}

	*/
/*
	var name string
	var Surname string

	fmt.Println("What is your name?")
	fmt.Scan(&name)

	fmt.Println("What is your surname?")
	fmt.Scan(&Surname)

	fmt.Printf("Hi %v %v, thank you for your details", name, Surname)
	

 SIMPLE CALCULATOR

	Create a simple calculator in Go that can perform basic arithmetic operations (addition, subtraction, multiplication, and division)
 	based on user input. This challenge will help you practice using variables, printing and scanning user input, and 
 	conditional statements.

	Requirements

	- Prompt the user to enter two numbers.
	- Prompt the user to choose an arithmetic operation (+, -, *, /).
	- Perform the chosen operation on the two numbers.
	- Print the result of the operation.
	*/

	var firstNumber int
	var secondNumber int
	var selectOperation string

	fmt.Println("Hi, I am a simple calculator")

	fmt.Println("Type your first number:")
	fmt.Scan(&firstNumber)

	fmt.Println("Type your second number:")
	fmt.Scan(&secondNumber)

	fmt.Println("Choose your your operation? e.g +, -, *, /")
	fmt.Scan(&selectOperation)

	
	// conditions 

	switch selectOperation {
		case "+":
			fmt.Println(firstNumber + secondNumber)
		case "-": 
			fmt.Println(firstNumber - secondNumber)
		case "*": 
			fmt.Println(firstNumber * secondNumber)
		case "/":
			fmt.Println(firstNumber / secondNumber)
		default:
			fmt.Println("try again")
		}


	 



	













		






		



		/*

		var price float32 = 20.5

		budget := 3000.50

		var exit bool = false

		isRich := true
		var name string = "Zagel"

		surname := "Mohamed"

		*/



	// how to print golang variable?
	}