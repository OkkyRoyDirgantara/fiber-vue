package utils

import "net/mail"

func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
