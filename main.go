package main

import 
"fmt"


// Ask the user for 5 numbers and calculate
// the average of those 5 numbers

func main (){

	var number [5] int
	 // declaring the array will consist of 5


	for i := i < 5; i++  {
		fmt.Println("Enter 5 numbers: ")
		fmt.Scan(&number[i])
	}

	var sum float32 = 0

	for i := 0; i < 5; i++ {
		sum += float32(number[i])
	}
	var mean float32

	mean = sum / 5

	fmt.Println("The average is,", mean)
/*
	for i := 0; i < 5; i++ {
		fmt.Println("number", i+1, "->", number[i])

		number[0] = 15
	}

	*/

	}








}