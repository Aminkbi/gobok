package data

import (
	"github.com/aminkbi/gobok/cmd/internal/validation"
	"time"
)

type Book struct {
	ID          string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string   `json:"title" bson:"title"`
	Pages       int      `json:"pages" bson:"pages"`
	Categories  []string `json:"categories" bson:"categories"`
	Authors     []string `bson:"authors" json:"authors"`
	Description string   `bson:"description" json:"description"`
	Published   int      `json:"published" bson:"published"`
	ISBN        string   `json:"isbn" bson:"isbn"`
	ImageUrl    string   `json:"imageUrl" bson:"imageUrl" `
}

func ValidateBook(v *validation.Validator, book *Book) {

	v.Check(book.Title != "", "title", "must be provided")
	v.Check(len(book.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(book.Description != "", "description", "must be provided")
	v.Check(len(book.Description) <= 5000, "description", "must not be more than 5000 bytes long")
	v.Check(book.Authors != nil, "authors", "must be provided")
	v.Check(len(book.Authors) >= 1, "authors", "must contain at least 1 author")
	v.Check(len(book.Authors) <= 10, "authors", "must not contain more than 5 authors")
	v.Check(book.Published != 0, "published", "must be provided")
	v.Check(book.Published <= time.Now().Year(), "published", "must not be in the future")
	v.Check(book.Pages != 0, "pages", "must be provided")
	v.Check(book.Pages > 0, "pages", "must be a positive integer")
	v.Check(book.Categories != nil, "categories", "must be provided")
	v.Check(len(book.Categories) >= 1, "categories", "must contain at least 1 categories")
	v.Check(len(book.Categories) <= 10, "categories", "must not contain more than 5 categories")
	v.Check(len(book.ImageUrl) > 0, "imageUrl", "must be provided")
	v.Check(len(book.ISBN) > 0, "isbn", "must be provided")

}
