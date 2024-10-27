package query

import (
	"regexp"
	"strings"

	"github.com/aslrousta/persian"
)

var wordRegexp = regexp.MustCompile(`[\p{L}\p{N}\x{200C}]+`)

// FullText converts a query string to a query expression ready to be used in
// full-text search.
func FullText(query string) string {
	query = persian.Sanitize(query, persian.Arabic, persian.Special)
	words := wordRegexp.FindAllString(query, -1)
	if len(words) == 0 {
		return ""
	}
	return strings.Join(words, "+") + ":*"
}
