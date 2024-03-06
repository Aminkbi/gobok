package data

import (
	"github.com/aminkbi/gobok/cmd/internal/validation"
)

type Author struct {
	Name         string   `bson:"name" json:"name"`
	Biography    string   `bson:"biography" json:"biography"`
	BooksWritten []string `bson:"books_written" json:"booksWritten"`
}

func ValidateAuthor(v *validation.Validator, author *Author) {
	v.Check(author.Name != "", "name", "must be provided")
	v.Check(len(author.Name) <= 500, "name", "must not be more than 500 bytes long")
	v.Check(author.Biography != "", "biography", "must be provided")
	v.Check(len(author.Biography) <= 5000, "biography", "must not be more than 5000 bytes long")
}
