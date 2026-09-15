package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		firstWord := strings.Split(input, " ")[0]
		loweredFirstWord := strings.ToLower(firstWord)
		fmt.Printf("Your command was: %s\n", loweredFirstWord)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func cleanInput(text string) []string {
	cleanText := []string{}
	textSplit := strings.Split(text, " ")
	for _, word := range textSplit {
		loweredWord := strings.ToLower(word)
		if loweredWord != "" {
			cleanText = append(cleanText, loweredWord)
		}
	}
	return cleanText
}
