package utlis

import (
	"strings"
	"time"
	"unicode"
)

func GetFirstDayOfMonth(date time.Time) time.Time {
	// Get the year and month for the given date
	year, month, _ := date.Date()

	// Create a time value for the first day of the month
	firstDayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, date.Location())

	return firstDayOfMonth
}

func GetLastDayOfMonth(date time.Time) time.Time {
	// Subtract one day from the next month to get the last day of the current month
	lastDayOfMonth := GetFirstDayOfMonth(date).AddDate(0, 1, -1)

	return lastDayOfMonth
}

func ToSnakeCase(input string) string {
	var result strings.Builder
	for i, char := range RemoveSpaces(input) {
		if i > 0 && unicode.IsUpper(char) {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(char))
	}

	return result.String()
}

func RemoveSpaces(input string) string {
	return strings.ReplaceAll(input, " ", "")
}

// slice
// Remove Array by Index
func RemoveIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func SubString(text string, endIndex int) string {
	originalString := text
	if endIndex > len(originalString) {
		endIndex = len(originalString)
	}
	return originalString[:endIndex]
}
