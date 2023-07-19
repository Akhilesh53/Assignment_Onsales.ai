//====================================================================
//====================== UNIT TEST CASE FILE =========================
//====================================================================

package main

import (
	"reflect"
	"testing"
)

func TestCreateCumulativeProbs(t *testing.T) {
	// Test case 1: Normal probabilities
	outcomes := []Outcome{
		{1, 10},
		{2, 30},
		{3, 15},
		{4, 15},
		{5, 30},
		{6, 0},
	}
	expectedCumulativeProbs := []Outcome{
		{1, 0.1},
		{2, 0.4},
		{3, 0.55},
		{4, 0.7000000000000001},
		{5, 1.0},
		{6, 1.0},
	}

	cumulativeProbs := createCumulativeProbs(outcomes)

	if !reflect.DeepEqual(cumulativeProbs, expectedCumulativeProbs) {
		t.Errorf("Expected cumulative probabilities: %v, got: %v", expectedCumulativeProbs, cumulativeProbs)
	}

	// Test case 2: Zero probabilities
	outcomes = []Outcome{
		{1, 0},
		{2, 0},
		{3, 0},
	}
	expectedCumulativeProbs = []Outcome{
		{1, 0.0},
		{2, 0.0},
		{3, 0.0},
	}

	cumulativeProbs = createCumulativeProbs(outcomes)

	if !reflect.DeepEqual(cumulativeProbs, expectedCumulativeProbs) {
		t.Errorf("Expected cumulative probabilities: %v, got: %v", expectedCumulativeProbs, cumulativeProbs)
	}
}

func TestSimulateEvent(t *testing.T) {
	outcomes := []Outcome{
		{1, 10},
		{2, 30},
		{3, 15},
		{4, 15},
		{5, 30},
		{6, 0},
	}
	numOccurrences := 1000

	result := simulateEvent(outcomes, numOccurrences)

	// Check the occurrence count of each outcome
	expectedCounts := map[interface{}]int{
		1: 117,
		2: 291,
		3: 156,
		4: 149,
		5: 287,
		6: 0,
	}

	if !reflect.DeepEqual(result, expectedCounts) {
		t.Errorf("Expected occurrence counts: %v, got: %v", expectedCounts, result)
	}
}

func TestSimulateEventWithZeroOccurrences(t *testing.T) {
	outcomes := []Outcome{
		{1, 10},
		{2, 30},
	}
	numOccurrences := 0

	result := simulateEvent(outcomes, numOccurrences)

	// Check that there are no occurrences when the number of occurrences is 0
	expectedCounts := map[interface{}]int{
		1: 0,
		2: 0,
	}

	if !reflect.DeepEqual(result, expectedCounts) {
		t.Errorf("Expected occurrence counts: %v, got: %v", expectedCounts, result)
	}
}

func TestSimulateEventWithEmptyOutcomes(t *testing.T) {
	outcomes := []Outcome{}
	numOccurrences := 100

	result := simulateEvent(outcomes, numOccurrences)

	// Check that the result is an empty map when there are no outcomes
	expectedCounts := map[interface{}]int{}

	if !reflect.DeepEqual(result, expectedCounts) {
		t.Errorf("Expected occurrence counts: %v, got: %v", expectedCounts, result)
	}
}

func TestSimulateEventDistribution(t *testing.T) {
	outcomes := []Outcome{
		{1, 10},
		{2, 30},
		{3, 15},
		{4, 15},
		{5, 30},
	}
	numOccurrences := 10000

	result := simulateEvent(outcomes, numOccurrences)

	// Calculate the expected occurrence count for each outcome based on the probabilities
	expectedCounts := map[interface{}]int{
		1: 1000,
		2: 3000,
		3: 1500,
		4: 1500,
		5: 3000,
	}

	// Check if the occurrence counts are roughly within the expected range
	for outcome, expectedCount := range expectedCounts {
		actualCount := result[outcome]
		lowerBound := int(0.95 * float64(expectedCount))
		upperBound := int(1.05 * float64(expectedCount))

		if actualCount < lowerBound || actualCount > upperBound {
			t.Errorf("Unexpected occurrence count for outcome %v. Expected: %d, Actual: %d", outcome, expectedCount, actualCount)
		}
	}
}
