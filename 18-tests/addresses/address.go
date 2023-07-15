package addresses

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func AddressType(address string) string {
	validTypes := []string{"rua", "avenida", "rodovia"}

	addressLowercase := strings.ToLower(address)
	firstWord := strings.Split(addressLowercase, " ")[0]

	isValidType := false

	for _, t := range validTypes {
		if firstWord == t {
			isValidType = true
		}
	}

	if isValidType {
		return cases.Title(language.BrazilianPortuguese).String(firstWord)
	}

	return "Invalid type"
}
