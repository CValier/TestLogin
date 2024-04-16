package utils

import (
	"regexp"
	"strings"

	"github.com/CValier/PruebaGO/internal/pkg/entity"
)

// Validate password must contain a minimum of 6 characters and a maximum of 12 and contain at least one uppercase letter, one lowercase letter, one special character (@ $ or &) and a number.
func ValidatePassword(password string) bool {
	if len(password) < 6 || len(password) > 12 {
		return false
	}
	var hasUpper, hasLower, hasSpecial, hasNumber bool
	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case strings.ContainsAny(string(char), "@$&"):
			hasSpecial = true
		case '0' <= char && char <= '9':
			hasNumber = true
		}
	}
	return hasUpper && hasLower && hasSpecial && hasNumber
}

// Validate email format
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// Validate phone number format
func ValidatePhoneNumber(phoneNumber string) bool {
	phoneNumberRegex := regexp.MustCompile(`^\d{10}$`)
	return phoneNumberRegex.MatchString(phoneNumber)
}

// Validate all requested fields
func ValidateFields(user *entity.User) (bool, string) {

	if user.User == "" {
		return false, "'User' field is missing."
	} else if user.Email == "" {
		return false, "'Email' field is missing."
	} else if user.Password == "" {
		return false, "'Password' field is missing."
	} else if user.PhoneNumber == "" {
		return false, "'PhoneNumber' field is missing."
	}

	return true, ""
}
