package logs

// Application identifies the application emitting the given log.
/*
Implement the `Application` function that takes a log line and returns the application that emitted the log line.

To identify which application emitted a given log line, search the log line for a specific character as specified by the following table:
*/
import (
	"fmt"
	"unicode/utf8"
)

func Application(log string) string {
	// found := false
	fmt.Println(log)
	for _, r := range log {
		fmt.Println(r)
		switch r {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
/*
```go
log := "please replace '👎' with '👍'"

Replace(log, '👎', '👍')
// => please replace '👍' with '👍'"
```
*/
func Replace(log string, oldRune, newRune rune) string {
	newString := ""
	for _, r := range log {
		switch r {
		case oldRune:
			newString += string(newRune)
		default:
			newString += string(r)
		}
	}
	return newString
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
/*
WithinLimit("hello❗", 6)
// => true
*/
func WithinLimit(log string, limit int) bool {

	return utf8.RuneCountInString(log) <= limit

}
