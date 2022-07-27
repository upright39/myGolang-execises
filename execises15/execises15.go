package execises15

import "fmt"

type ElectionResult struct {
	Name  string
	Votes int
}

//NewVoteCount returns a new vote counte with a given number of inital votes
func NewVoteCounter(initialVote int) *int {

	return &initialVote

}

//VoteCount extracts the number of votes from a counter
func VoteCount(Counter *int) int {
	if Counter == nil {
		return 0
	}
	return *Counter
}

//IncrementVoteCount increments the value in a vote counter

func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

//NewElectionResult creates a new election result
func NewElectionResult(canditdateName string, votes int) *ElectionResult {
	return &ElectionResult{
		Name:  canditdateName,
		Votes: votes,
	}
}

//DesplayResult creats a message with the result to be displayed

func DesplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

//DecrementVotesOfCandidate decrement by one the vote count of a candidate in a map

func DecrementVotesOfCandidate(results map[string]int, candidate string) {

	results[candidate] -= 1

}
