package strings

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Generate a salted hash of the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares a plain text password with a hashed password
func CheckPasswordHash(password, hashedPassword string) bool {
	// Compare the hashed password with the plain text password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
