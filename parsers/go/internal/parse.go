package internal

import (
	"strings"
)

type separator uint8

const (
	prefix                      = "$vars:"
	suffix                      = '$'
	escape                      = '\\'
	pairSeparator     separator = ';'
	keyValueSeparator separator = '='
)

func ParseSponsorTextVariables(text string) map[string]string {
	start := strings.Index(text, prefix)
	if start == -1 {
		return nil
	}

	variables := make(map[string]string)
	var key, value strings.Builder
	var escaped bool
	next := keyValueSeparator
	for i := start + len(prefix); i < len(text); i++ {
		// Stop parsing on end marker
		if !escaped && text[i] == suffix {
			break
		}

		// Detect and skip escape characters (unless they are escaped themselves)
		if !escaped && text[i] == escape {
			escaped = true
			continue
		}

		// Skip (irrelevant) whitespaces
		if !escaped && isSkippableWhitespace(text, i, next, value.String()) {
			continue
		}

		marker := !escaped && text[i] == uint8(next)
		if marker && next == keyValueSeparator {
			// Found colon after key, add key to result and switch to reading value
			variables[key.String()] = ""
			next = pairSeparator
		} else if marker && next == pairSeparator {
			// Found semicolon after value, save value, start new key and switch to reading key
			variables[key.String()] = value.String()
			key.Reset()
			value.Reset()
			next = keyValueSeparator
		} else if next == keyValueSeparator {
			// Currently reading key, append char to key
			key.WriteByte(text[i])
		} else {
			// Currently reading value, append char to value
			value.WriteByte(text[i])
		}

		// Reset escaped flag
		escaped = false
	}

	// Add current key-value pair if key was complete (and thus already added to the map)
	// No need to add an empty value, as we already initialized the key with an empty string
	if _, exists := variables[key.String()]; exists && value.Len() > 0 {
		variables[key.String()] = value.String()
	}

	return variables
}

func isSkippableWhitespace(text string, i int, next separator, value string) bool {
	// Bounds check
	if i < 0 || i >= len(text) {
		return false
	}

	// Don't skip any non-whitespace characters
	if text[i] != ' ' {
		return false
	}

	// Skip key whitespace (keys do not permit whitespaces)
	if next == keyValueSeparator {
		return true
	}

	// Safety check, next should only ever be pairSeparator here
	if next != pairSeparator {
		return false
	}

	// Skip leading value whitespace (value empty = nothing's been added yet)
	if value == "" {
		return true
	}

	// Skip global trailing whitespace (current whitespace is last character or next character is suffix)
	if len(text) == i+1 || len(text) > i+1 && text[i+1] == suffix {
		return true
	}

	// Skip consecutive value whitespace (next character is also a whitespace)
	if len(text) > i+1 && text[i+1] == ' ' {
		return true
	}

	// Skip trailing value whitespace (next character is key-value pair separator)
	if len(text) > i+1 && text[i+1] == uint8(pairSeparator) {
		return true
	}

	return false
}
