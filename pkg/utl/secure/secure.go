package secure

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// IsPasswordSecure checks whether password is secure enough
func IsPasswordSecure(pass string) bool {

	//min 8 chars, upper, lower, digit
	matched, _ := regexp.Match(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$`, []byte(pass))
	return matched
}

// Hash hashes the password using bcrypt
func Hash(password string) string {
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPW)
}

// HashMatchesPassword matches hash with password. Returns true if hash and password match.
func HashMatchesPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
