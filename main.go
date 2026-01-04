package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		firstWord := cleaned[0]                         // Nimmt das ERSTE Element des Slice
		fmt.Printf("Your command was: %s\n", firstWord) // Gibt nur das erste Wort aus
	}
}
