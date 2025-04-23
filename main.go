package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: md2html-go <input.md>")
		return
	}

	inputFile := os.Args[1]
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var html strings.Builder

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Check for a matching handler
		matched := false
		for prefix, handler := range handlers {
			if strings.HasPrefix(line, prefix) {
				html.WriteString(handler(line))
				matched = true
				break
			}
		}

		// Fallback to paragraph if no handler matches
		if !matched {
			html.WriteString(handleParagraph(line))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(html.String())
}
