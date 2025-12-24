package auth_utils

import (
	"errors"
	"unicode"
)

// ValidatePassword checks that password meets requirements:
// 8-60 characters, at least one lowercase, one uppercase, one digit, one special character
func ValidatePassword(password string) error {
	if len(password) < 8 || len(password) > 60 {
		return errors.New("password must be 8-60 characters")
	}

	var hasLower, hasUpper, hasNumber, hasSpecial bool
	for _, c := range password {
		switch {
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsDigit(c):
			hasNumber = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	if !hasLower || !hasUpper || !hasNumber || !hasSpecial {
		return errors.New("password must include at least one uppercase letter, one lowercase letter, one number, and one special character")
	}

	return nil
}

// ValidateUsername checks username requirements: 4-120 characters, letters, numbers, special characters allowed
func ValidateUsername(username string) error {
	if len(username) < 4 || len(username) > 120 {
		return errors.New("username must be 4-120 characters")
	}
	// optional: further character checks can be added
	return nil
}

// ValidateFullname checks fullname requirements: 4-120 characters, letters, spaces, special characters allowed
func ValidateFullname(fullname string) error {
	if len(fullname) < 4 || len(fullname) > 120 {
		return errors.New("fullname must be 4-120 characters")
	}
	// optional: further character checks can be added
	return nil
}

func ValidateUniqueID(id string) error {
	if len(id) != 20 {
		return errors.New("unique id must be exactly 20 characters")
	}

	for _, c := range id {
		if c == ' ' {
			return errors.New("unique id cannot contain spaces")
		}
		if !(unicode.IsLower(c) || unicode.IsDigit(c) || unicode.IsPunct(c) || unicode.IsSymbol(c)) {
			return errors.New("unique id must contain only lowercase letters numbers or symbols")
		}
	}
	return nil
}
