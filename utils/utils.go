package utils

import "regexp"

func ConvertToHTML(text string) string {
	// Bold: **text**
	boldRe := regexp.MustCompile(`\*\*(.*?)\*\*`)
	text = boldRe.ReplaceAllString(text, `<strong>$1</strong>`)

	// Italic: *text*
	italicRe := regexp.MustCompile(`\*(.*?)\*`)
	text = italicRe.ReplaceAllString(text, `<em>$1</em>`)

	// Link: [text](url)
	linkRe := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	text = linkRe.ReplaceAllString(text, `<a href="$2">$1</a>`)

	return text
}
