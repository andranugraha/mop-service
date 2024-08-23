package strings

import "net/mail"

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidCellNumber(cellNumber string) bool {
	return true
}
