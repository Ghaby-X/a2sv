
// Package controllers handle input
package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ghaby-X/library_manager/models"
	"github.com/Ghaby-X/library_manager/services"
)

type ControllerFunctions func(args []string, lm *services.LibraryManager)

func Run(command string, lm *services.LibraryManager) {
	availableCommands := map[string]ControllerFunctions{
		"add_book":             AddBook,
		"remove_book":          RemoveBook,
		"borrow_book":          BorrowBook,
		"return_book":          ReturnBook,
		"list_available_books": ListAvailableBooks,
		"list_borrowed_books":  ListBorrowedBooks,
		"reserve_book":         ReserveBook,
	}

	commandSlice := strings.Split(command, " ")
	// ensure commands is greater than 0
	if len(commandSlice) < 1 {
		return
	}

	controllerFunction, prs := availableCommands[commandSlice[0]]

	// checks if command is present in available commands
	if !prs {
		fmt.Println("Invalid command: ", commandSlice[0])
		return
	}

	controllerFunction(commandSlice, lm)
}

func AddBook(args []string, lm *services.LibraryManager) {
	// requires book title and author
	if len(args) != 3 {
		fmt.Println("Usage: add_book <Title> <Author>")
		return
	}

	newBook := models.Book{
		ID:     lm.GenerateID(),
		Title:  args[1],
		Author: args[2],
		Status: "Available",
	}

	lm.AddBook(newBook)
}

func RemoveBook(args []string, lm *services.LibraryManager) {
	// takes command and bookID
	if len(args) != 2 {
		fmt.Println("Usage: remove_book <bookID>")
		return
	}

	// typecasting bookid
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("failed to parse book ID: ", args[1])
		return
	}

	lm.RemoveBook(id)
}

func BorrowBook(args []string, lm *services.LibraryManager) {
	if len(args) != 3 {
		fmt.Println("Usage: borrow_book <bookID> <memberID>")
		return
	}

	memberID, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("failed to parse memberID: ", memberID)
		return
	}
	bookID, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("failed to parse bookID: ", bookID)
		return
	}

	lm.BorrowBook(bookID, memberID)
}

func ReturnBook(args []string, lm *services.LibraryManager) {
	if len(args) != 3 {
		fmt.Println("Usage: borrow_book <bookID> <memberID>")
		return
	}

	memberID, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("failed to parse memberID: ", memberID)
		return
	}
	bookID, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("failed to parse bookID: ", bookID)
		return
	}

	lm.ReturnBook(bookID, memberID)
}

func ReserveBook(args []string, lm *services.LibraryManager) {
	if len(args) != 3 {
		fmt.Println("Usage: reserve_book <bookID> <memberID>")
		return
	}

	memberID, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("failed to parse memberID: ", memberID)
		return
	}
	bookID, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("failed to parse bookID: ", bookID)
		return
	}

	err = lm.ReserveBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	}
}

func ListAvailableBooks(args []string, lm *services.LibraryManager) {
	Books := lm.ListAvailableBooks()

	if len(Books) < 1 {
		fmt.Println("No available books")
		return
	}

	for _, v := range Books {
		if v.Status == "Available" {
			printBook(&v)
		}
	}
}

func ListBorrowedBooks(args []string, lm *services.LibraryManager) {
	memberID, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("failed to convert memberID")
		return
	}

	Books := lm.ListBorrowedBooks(memberID)

	if len(Books) < 1 {
		fmt.Println("No Borrowed books")
		return
	}

	for _, v := range Books {
		if v.Status == "Borrowed" {
			printBook(&v)
		}
	}
}

func printBook(book *models.Book) {
	fmt.Println("ID: ", book.ID, " Title: ", book.Title, " Author: ", book.Author)
}

