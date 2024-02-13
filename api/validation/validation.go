package validation

import "net/mail"

func EmailValidation(email string) bool {
	emailAddress, err := mail.ParseAddress(email)
	return err == nil && emailAddress.Address == email
}
