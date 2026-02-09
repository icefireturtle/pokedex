package main
import "strings"

func cleanInput(text string) []string {
	clean := strings.ToLower(strings.TrimSpace(text))
	words := strings.Fields(clean)
	return words
}
	
