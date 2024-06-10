package main

import 
	"fmt"

func main () {

	/* creating a project that will ask for book data to the user. 
	for each book we should ask for the title, author, and publication year
	Now, depending on the year we should say the century of the book.




	*/

	var title string
	var author string
	var publicationYear int

	var exit bool = false 

	for exit == false {

		fmt.Println("Please enter the book title: ")
		fmt.Scan(&title)

		if title == "exit" || title == "Exit" {
			exit = true
			continue
		}

		fmt.Println("Please enter the book author: ")
		fmt.Scan(&author)

		fmt.Println("Please the book publication year: ")
		fmt.Scan(&publicationYear)
		 
		if publicationYear != 1635 && publicationYear != 1724 && publicationYear != 1820 && publicationYear != 1920 && publicationYear != 2020 {
			exit = true 
			continue
		}

		switch publicationYear {
		case 1635:
			fmt.Println("Your book belongs to the 17th century.")
		case 1724:
			fmt.Println("Your book belongs to the 18th century.")
		case 1820:
			fmt.Println("Your book belongs to the 19th century")
		case 1920:
			fmt.Println("Your book belons to the 20th century.")
		case 2020: 
			fmt.Println("Your book belongs to the 21st century")

		}


	}	






}