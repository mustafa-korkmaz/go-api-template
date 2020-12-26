package secure

import (
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// IsPasswordSecure checks whether password is secure enough
func IsPasswordSecure(pass string) bool {

	var (
		hasMinLen  = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(pass) >= 7 {
		hasMinLen = true
	}
	for _, char := range pass {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasNumber && hasSpecial
}

// Hash hashes the password using bcrypt
func Hash(password string) string {
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPW)
}

// ValidatePassword matches hashed password text with plain password. Returns true if hash and password match.
func ValidatePassword(hashedPass, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password)) == nil
}
