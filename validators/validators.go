// Package validators impls custom validators
package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidateCoolTitle returns true when the field value contains the word "cool".
func ValidateCoolTitle(fl validator.FieldLevel) bool {
	return strings.Contains(fl.Field().String(), "cool")
}