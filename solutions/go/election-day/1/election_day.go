package electionday

import "fmt"

/*
Create a function `NewVoteCounter` that accepts the number of initial votes for a candidate and returns a pointer referring to an `int`, initialized with the given number of initial votes.
*/
// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter != nil {
		return *counter
	} else {
		return 0
	}
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	if counter != nil {
		*counter = *counter + increment
	}
}

/*
Create a function `NewElectionResult` that receives the name of a candidate and their number of votes and
returns a new election result.
*/

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{Name: candidateName, Votes: votes}
}

/*
Create a function `DisplayResult` that will receive an `*ElectionResult` as an argument and will return a string with the message to display.


```go
var result *ElectionResult
result = &ElectionResult{
    Name: "John",
    Votes: 32,
}

DisplayResult(result)
// => John (32)
*/
// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	_, ok := results[candidate]
	if ok != true {
		return
	}
	results[candidate] -= 1
}
