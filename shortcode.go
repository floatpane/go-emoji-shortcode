// Package shortcode maps emoji shortcodes like :smile: to Unicode emoji
// characters and provides prefix-based suggestion lookup.
package shortcode

import (
	"sort"
	"strings"
	"unicode"
)

// Match is a shortcode lookup result.
type Match struct {
	Code  string // shortcode without colons, e.g. "smile"
	Emoji string // Unicode emoji, e.g. "😄"
}

// Lookup returns the exact emoji for a shortcode (without surrounding colons).
func Lookup(code string) (string, bool) {
	code = normalize(code)
	if code == "" {
		return "", false
	}
	em, ok := byCode[code]
	return em, ok
}

// Suggest returns fuzzy-matched shortcodes for the given prefix. The prefix is
// treated as the start of a shortcode, e.g. "sm" matches "smile" and "smirk".
// Results are sorted by exactness, then length, then alphabetically.
func Suggest(prefix string) []Match {
	prefix = normalize(prefix)
	if prefix == "" {
		return nil
	}
	var matches []Match
	for code, em := range byCode {
		if strings.HasPrefix(code, prefix) {
			matches = append(matches, Match{Code: code, Emoji: em})
		}
	}
	if len(matches) == 0 {
		return nil
	}
	sort.Slice(matches, func(i, j int) bool {
		a, b := matches[i].Code, matches[j].Code
		// Exact match first.
		if a == prefix && b != prefix {
			return true
		}
		if b == prefix && a != prefix {
			return false
		}
		// Then shorter, then alphabetical.
		if len(a) != len(b) {
			return len(a) < len(b)
		}
		return a < b
	})
	return matches
}

// All returns every shortcode/emoji pair, sorted by shortcode.
func All() []Match {
	codes := make([]string, 0, len(byCode))
	for code := range byCode {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	result := make([]Match, 0, len(codes))
	for _, code := range codes {
		result = append(result, Match{Code: code, Emoji: byCode[code]})
	}
	return result
}

func normalize(code string) string {
	code = strings.TrimSpace(strings.ToLower(code))
	code = strings.Trim(code, ":")
	var b strings.Builder
	for _, r := range code {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
