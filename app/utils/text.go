package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToTitleCase(data string) string {
	mustBeLowerCase := make(map[string]string)
	mustBeLowerCase["AND"] = "and"
	mustBeLowerCase["ON"] = "on"

	caser := cases.Title(language.English)

	words := strings.Fields(data)
	for i, word := range words {
		if strings.HasPrefix(word, "'") && strings.HasSuffix(word, "'") {
			words[i] = strings.ToLower(word)
		} else if newWord := mustBeLowerCase[strings.ToUpper(word)]; newWord != "" {
			words[i] = newWord
		} else {
			words[i] = caser.String(word)
		}
	}
	convertedString := strings.Join(words, " ")
	return convertedString
}

func ToUpperCase(data string) string {
	caser := cases.Upper(language.English)
	return caser.String(data)
}

func ToLowerCase(data string) string {
	caser := cases.Lower(language.English)
	return caser.String(data)
}
