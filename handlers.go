package main

import (
	"strings"
)

type handlerFunc func(string) string

func applyBold(line string) string {
	for strings.Contains(line, "*") {
		start := strings.Index(line, "*")
		end := strings.Index(line[start+1:], "*")
		if start == -1 || end == -1 {
			break
		}
		end += start + 1
		boldText := "<b>" + line[start+1:end] + "</b>"
		line = line[:start] + boldText + line[end+1:]
	}
	return line
}

func handleParagraph(line string) string {
	line = applyBold(strings.TrimSpace(line))
	return "<p>" + line
}

func handleHeading(line string) string {
	line = applyBold(strings.TrimSpace(line[2:]))
	return "<h1>" + line + "</h1>"
}

func handleHeading2(line string) string {
	line = applyBold(strings.TrimSpace(line[3:]))
	return "<h2>" + line + "</h2>"
}

var handlers = map[string]handlerFunc{
	"##": handleHeading2,
	"#":  handleHeading,
}
