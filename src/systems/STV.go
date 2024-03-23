package voting_systems

import (
	"fmt"
	"math"
)

func InitiateSTV(candidates []Candidate, voters []Voter) (Candidate, []Candidate, [][]Candidate) {
	//Array of rounds of canidate changes
	CanidateRounds := [][]Candidate{}

	for i := range candidates {
		candidates[i].NumVotes = 0
	}
	totalVoters := countTotalVotes(voters)
	candidates = Round(voters, candidates)
	CanidateRounds = append(CanidateRounds, append([]Candidate(nil), candidates...))
	_, winner, candidates, CanidateRounds := SimulateSTV(candidates, totalVoters, voters, CanidateRounds)
	for i := range CanidateRounds {
		fmt.Println("Round ", i)
		fmt.Println(CanidateRounds[i])
	}
	return winner, candidates, CanidateRounds
}

func countTotalVotes(voters []Voter) int {
	totalVotes := 0
	for _, voter := range voters {
		totalVotes += voter.Number
	}
	return totalVotes
}

func SimulateSTV(candidates []Candidate, totalVoters int, voters []Voter, CanidateRounds [][]Candidate) (float64, Candidate, []Candidate, [][]Candidate) {
	cost := -1.0
	//Base case
	for i := range candidates {
		if candidates[i].NumVotes > totalVoters/2 {
			fmt.Println("Winner: ", candidates[i])
			return GetSocailCost(candidates[i], voters), candidates[i], candidates, CanidateRounds
		}
	}

	//Recursive Case
	if cost == -1.0 {
		minVotes := totalVoters // Assuming totalVoters is higher than any candidate's numVotes can be
		candidatePosition := 0
		for j, candidate := range candidates {
			if candidate.NumVotes < minVotes {
				minVotes = candidate.NumVotes
				candidatePosition = j
			}
		}

		candidates = append(candidates[:candidatePosition], candidates[candidatePosition+1:]...)
		candidates = Round(voters, candidates)
		CanidateRounds = append(CanidateRounds, append([]Candidate(nil), candidates...))

		// Adjust recursive call to capture and return the updated candidates slice
		cost, winner, candidates, CanidateRounds := SimulateSTV(candidates, totalVoters, voters, CanidateRounds)
		return cost, winner, candidates, CanidateRounds
	}

	//Return empty candidate if no winner is found
	return cost, Candidate{}, candidates, CanidateRounds
}

// Simulates a round by determining the number of votes each canidate gets
func Round(voters []Voter, canidates []Candidate) []Candidate {
	//Reset the number of votes for each candidate
	for i := 0; i < len(canidates); i++ {
		canidates[i].NumVotes = 0
	}
	//For each voter determine the closest candidate and save it
	for i := 0; i < len(voters); i++ {
		minDistance := 10000.0
		canidatePosition := 0
		for j := 0; j < len(canidates); j++ {
			distance := math.Abs(voters[i].Position - canidates[j].Position)
			if distance < minDistance {
				minDistance = distance
				canidatePosition = j
			}
		}
		//add that number of voters to the candidate
		canidates[canidatePosition].NumVotes += voters[i].Number

	}
	return canidates
}
