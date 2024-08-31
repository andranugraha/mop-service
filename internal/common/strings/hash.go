package strings

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	// Generate a salted hash of the password using bcrypt
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// CheckPasswordHash compares a plain text password with a hashed password
func CheckPasswordHash(password, hashedPassword string) bool {
	// Compare the hashed password with the plain text password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
