// Package services for library service functiosn
package services

import (
	"fmt"

	"github.com/Ghaby-X/library_manager/models"
)

type LibraryManagerInterface interface {
	AddBook(models.Book)
	RemoveBook(int)
	BorrowBook(int, int) error
	ReturnBook(int, int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(int) []models.Book
}

func IDGenerator() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

type LibraryManager struct {
	Books      map[int]*models.Book
	Members    map[int]*models.Member
	GenerateID func() int
}

func NewLibraryManager() *LibraryManager {
	idGenerator := IDGenerator()
	lm := &LibraryManager{
		Books:      make(map[int]*models.Book),
		Members:    make(map[int]*models.Member),
		GenerateID: idGenerator,
	}
	return lm
}

func (lm *LibraryManager) AddBook(book models.Book) {
	lm.Books[book.ID] = &book
}

func (lm *LibraryManager) RemoveBook(bookID int) {
	delete(lm.Books, bookID)
}

func (lm *LibraryManager) BorrowBook(bookID int, memberID int) error {
	book, prs := lm.Books[bookID]
	if !prs {
		return fmt.Errorf("could not find book with ID: %d", bookID)
	}

	_, prs = lm.Members[memberID]
	if !prs {
		return fmt.Errorf("could not find member with ID: %d", memberID)
	}

	// copy book to member
	lm.Members[memberID].BorrowedBooks[bookID] = *book

	// update book status
	lm.Books[bookID].Status = "Borrowed"

	return nil
}

func (lm *LibraryManager) ReturnBook(bookID int, memberID int) error {
	book, prs := lm.Books[bookID]
	if !prs {
		return fmt.Errorf("could not find book with ID: %d", bookID)
	}

	_, prs = lm.Members[memberID]
	if !prs {
		return fmt.Errorf("could not find member with ID: %d", memberID)
	}

	// removeBook from member
	book.Status = "Available"

	// update book status
	return nil
}

func (lm *LibraryManager) ListAvailableBooks() []models.Book {
	books := []models.Book{}
	for _, v := range lm.Books {
		books = append(books, *v)
	}

	return books
}

func (lm *LibraryManager) ListBorrowedBooks(memberID int) []models.Book {
	return lm.Members[memberID].BorrowedBooks
}
