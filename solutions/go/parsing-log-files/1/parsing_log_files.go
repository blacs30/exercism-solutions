package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[ERR\]|^\[INF\]|^\[DBG\]|^\[TRC\]`)
	return re.MatchString(text)
}

/*
A new team has joined the organization, and you find their log files are using a strange separator for "fields".
Instead of something sensible like a colon ":" they use a string such as "<--->" or "<=>" (because it's prettier) in fact any string that has a first character of "<" and a last character of ">" and any combination of the following characters "~", "\*", "=" and "-" in between.

Implement the `SplitLogLine` function that takes a line and returns an array of strings each of which contains a field.

```go
SplitLogLine("section 1<*>section 2<~~~>section 3")
// => []string{"section 1", "section 2", "section 3"},

```

	"~", "\*", "=" and "-"
*/
func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

/*
```go

	lines := []string{
	    `[INF] passWord`, // contains 'password' but not surrounded by quotation marks
	    `"passWord"`,  // count this one
	    `[INF] User saw error message "Unexpected Error" on page load.`, // does not contain 'password'
	    `[INF] The message "Please reset your password" was ignored by the user`, // count this one
	}

// => 2
```
*/
func CountQuotedPasswords(lines []string) int {
	var matchCounter int
	for _, line := range lines {
		line = strings.ToLower(line)
		re := regexp.MustCompile(`".*password.*"`)
		foundMatches := re.FindAllString(line, -1)
		matchCounter += len(foundMatches)
	}
	return matchCounter
}

/*
```go
RemoveEndOfLineText("[INF] end-of-line23033 Network Failure end-of-line27")
// => "[INF]  Network Failure "
```
*/
func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\w*[\d]+`)
	return re.ReplaceAllString(text, "")
}

/*

Implement a function `TagWithUserName` that processes log lines:

- Lines that do not contain the string `"User "` remain unchanged.
- For lines that contain the string `"User "`, prefix the line with `[USR]` followed by the user name.

For example:

```go
result := TagWithUserName([]string{
    "[WRN] User James123 has exceeded storage space.",
	"[WRN] Host down. User   Michelle4 lost connection.",
	"[INF] Users can login again after 23:00.",
	"[DBG] We need to check that user names are at least 6 chars long.",

})
// => []string {
//  "[USR] James123 [WRN] User James123 has exceeded storage space.",
//  "[USR] Michelle4 [WRN] Host down. User   Michelle4 lost connection.",
//  "[INF] Users can login again after 23:00.",
//  "[DBG] We need to check that user names are at least 6 chars long."
// }
```

You can assume that:

- User names are followed by at least one whitespace character in the log.
- There is at most one occurrence of the string `"User "` in each line.
- User names are non-empty strings that do not contain whitespace.

*/

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`.*User\s+([a-zA-Z0-9]+).*`)
	newText := []string{}
	for _, line := range lines {
		user := re.FindStringSubmatch(line)
		if len(user) > 1 {
			newText = append(newText, fmt.Sprintf("[USR] %s %s", user[1], line))
		} else {
			newText = append(newText, line)
		}
	}
	return newText
}
