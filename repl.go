package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
)
func startREPL() {
scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() 
		clean := cleanInput(scanner.Text())
		if len(clean) == 0 {
			continue
		}
		firstWord := clean[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func cleanInput(text string) []string {
	clean := strings.ToLower(strings.TrimSpace(text))
	words := strings.Fields(clean)
	return words
}
	
