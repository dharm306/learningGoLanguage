package utility

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func UcFirst(input string) string {
	return cases.Title(language.English).String(input)
}
