package data

import "github.com/aminkbi/gobok/cmd/internal/validation"

type Category struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
}

func ValidateCategory(v *validation.Validator, category *Category) {
	v.Check(category.Name != "", "name", "must be provided")
	v.Check(len(category.Name) <= 500, "name", "must not be more than 500 bytes long")
	v.Check(category.Description != "", "description", "must be provided")
	v.Check(len(category.Description) <= 5000, "description", "must not be more than 5000 bytes long")

}
