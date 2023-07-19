package main

import (
	"fmt"
	"math"
	"math/rand"
)

// struct to input map values
type Outcome struct {
	Value       interface{}
	Probability float64
}

// struct used for Alias method
type ProbabilityAlias struct {
	AliasIndex    int
	RemainingProb float64
}

/*
This function is used to caluculate cumulative probabilities of events
*/
func createCumulativeProbs(outcomes []Outcome) []Outcome {
	totalProb := 0.0
	for _, o := range outcomes {
		totalProb += o.Probability
	}

	cumulativeProbs := make([]Outcome, len(outcomes))
	currentProb := 0.0

	for i, o := range outcomes {
		currentProb += o.Probability / totalProb
		if math.IsNaN(currentProb){
			currentProb = 0
		}
		cumulativeProbs[i] = Outcome{Value: o.Value, Probability: currentProb}
	}

	return cumulativeProbs
}

/*
This function will print the simulation of event that followa thw given input bias
*/
func simulateEvent(outcomes []Outcome, numOccurrences int) map[interface{}]int {
	// calculate the cumulative probabilities of events
	// Cumulative probabilities help us map the random numbers generated to their corresponding outcomes
	// based on the given probabilities. It allows us to create intervals on the range [0, 1] for each outcome
	cumulativeProbs := createCumulativeProbs(outcomes)

	// Output map
	occurrencesCount := make(map[interface{}]int)

	// Intialise map
	for _, o := range outcomes {
		occurrencesCount[o.Value] = 0
	}

	// for number of eccurences, take a random number
	// increase the count of occurence for the event whose probability is less than generated random number
	for i := 0; i < numOccurrences; i++ {
		randomNum := rand.Float64()
		for _, o := range cumulativeProbs {
			if randomNum < o.Probability {
				occurrencesCount[o.Value]++
				break
			}
		}
	}

	return occurrencesCount
}

/*
========================================================================================================
============================================== EXPLANATION =============================================
1) Convert the input probabilities into cumulative probabilities.
2) Generate random numbers between 0 and 1 to simulate the occurrence of the event.
3) Map the random numbers to their corresponding outcomes based on the cumulative probabilities.

TIME COMPLEXITY: O(n^2)
*/
func main() {
	biasedDiceOutcomes := []Outcome{
		{1, 10},
		{2, 30},
		{3, 15},
		{4, 15},
		{5, 30},
		{6, 0},
	}

	flippingCoinOutcomes := []Outcome{
		{"Head", 35},
		{"Tail", 65},
	}

	numOccurrences := 1000

	fmt.Println("\nBiased Dice:")
	fmt.Println(simulateEvent(biasedDiceOutcomes, numOccurrences))

	fmt.Println("\nFlipping Coin:")
	fmt.Println(simulateEvent(flippingCoinOutcomes, numOccurrences))
}
