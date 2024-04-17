package main

import (
	"fmt"
	"regexp"
	"strings"
)

var cardNumRegex = regexp.MustCompile(`^([456]\d{3}-?\d{4}-?\d{4}-?\d{4})$`)

func isValid(cardNum string) bool {

	if !cardNumRegex.MatchString(cardNum) {
		return false
	}

	lastDigit := ' '
	lastDigitCount := 1
	for _, letter := range cardNum {
		if !strings.ContainsRune("0123456789", letter) {
			continue
		}

		if letter == lastDigit {
			lastDigitCount++
		} else {
			lastDigitCount = 1
			lastDigit = letter
		}

		if lastDigitCount >= 4 {
			return false
		}
	}

	return true
}

func main() {
	cardNum := []string{
		"4123456789123456",
		"5123-4567-8912-3456",
		"61234-567-8912-3456",
		"4123356789123456",
		"5133-3367-8912-3456",
		"5123 - 3567 - 8912 - 3456",
	}

	for i := 0; i < len(cardNum); i++ {
		if isValid(cardNum[i]) {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
	}
}
