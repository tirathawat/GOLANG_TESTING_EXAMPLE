package utils

import (
	"net/mail"
	"regexp"
)

type helper struct{}

func NewHelper() helper {
	return helper{}
}

func (h helper) IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (h helper) IsPhoneValid(phone string) bool {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	return re.MatchString(phone)
}
