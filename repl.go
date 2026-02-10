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
		
		command, exists := commands[clean[0]]
		if exists {
			err:= command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func commandExit () error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	//Header
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, directive := range commands {
		fmt.Printf("%s: %s\n", directive.name, directive.description)
	}
	return nil
}

func cleanInput(text string) []string {
	clean := strings.ToLower(strings.TrimSpace(text))
	words := strings.Fields(clean)
	return words
}
	
type cliCommand struct {
	name string
	description string
	callback func() error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
}
