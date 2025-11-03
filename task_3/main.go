// Packag emain
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Ghaby-X/library_manager/controllers"
	"github.com/Ghaby-X/library_manager/models"
	"github.com/Ghaby-X/library_manager/services"
)

func seedMembers(lm *services.LibraryManager) {
	members := []models.Member{
		{ID: 1, Name: "Joyce"},
		{ID: 2, Name: "Gabs"},
		{ID: 3, Name: "Flore"},
		{ID: 4, Name: "Kaycy"},
		{ID: 5, Name: "Noah"},
		{ID: 6, Name: "Ali"},
	}

	for _, v := range members {
		lm.Members[v.ID] = &v
	}
}

func displayWelcome() {
	fmt.Println("Welcome to Library mamanger")
	fmt.Println("type 'q' to quit, 'help' for help")
	fmt.Println()
}

func displayHelp() {
	fmt.Println()
	fmt.Println("Available commands: ")
	fmt.Println("add_book;  add new book to the library")
	fmt.Println("Usage: add_book <Title> <Author>")
	fmt.Println()
	fmt.Println("remove_book;  remove book to the library")
	fmt.Println("Usage: remove_book <bookID>")
	fmt.Println()
	fmt.Println("borrow_book;  borrow book to a member")
	fmt.Println("Usage: borrow_book <bookID> <memberID>")
	fmt.Println()
	fmt.Println("return_book;  return book to library")
	fmt.Println("Usage: return_book <bookID> <memberID>")
	fmt.Println()
	fmt.Println("list_available_books;  list all available books")
	fmt.Println("Usage: list_available_books")
	fmt.Println()
	fmt.Println("list_borrowed_books;  list all borrowed books by a member")
	fmt.Println("Usage: list_borrowed_books <memberID>")
	fmt.Println()
}

func main() {
	// initializing a library
	lm := services.NewLibraryManager()
	seedMembers(lm)

	scanner := bufio.NewScanner(os.Stdin)

	displayWelcome()
	for {
		fmt.Print(">>> ")
		scanner.Scan()
		line := scanner.Text()

		if line == "help" {
			displayHelp()
			continue
		}

		if line == "q" {
			return
		}

		controllers.Run(line, lm)
		fmt.Println()
	}
}
