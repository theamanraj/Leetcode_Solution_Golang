package main

import "sort"

func findItinerary(tickets [][]string) []string {

	adjacency := make(map[string][]string)

	for _, pairs := range tickets {
		if _, ok := adjacency[pairs[0]]; !ok {
			adjacency[pairs[0]] = []string{pairs[1]}
		} else {
			adjacency[pairs[0]] = append(adjacency[pairs[0]], pairs[1])
			sort.Strings(adjacency[pairs[0]])
		}
	}

	current := "JFK"
	itinerary := []string{current}

	if finalItinerary, ok := findItineraryHelper(itinerary, adjacency); ok {
		return finalItinerary
	}

	return []string{}

}

func findItineraryHelper(currentItinerary []string, currentGraph map[string][]string) ([]string, bool) {

	if len(currentGraph) == 0 {
		return currentItinerary, true
	}

	lastLeg := currentItinerary[len(currentItinerary)-1]

	for _, possibleDestination := range currentGraph[lastLeg] {

		currentItinerary := append(currentItinerary, possibleDestination)

		if len(currentGraph[lastLeg]) > 1 {
			currentGraph[lastLeg] = deleteStringInSlice(currentGraph[lastLeg], possibleDestination)

		} else {
			delete(currentGraph, lastLeg)
		}

		// Trying the proposal
		if proposedItinerary, found := findItineraryHelper(currentItinerary, currentGraph); found {
			return proposedItinerary, true
		}

		// Backtracking: Putting the pair back
		if _, ok := currentGraph[lastLeg]; !ok {
			currentGraph[lastLeg] = []string{possibleDestination}
		} else {
			currentGraph[lastLeg] = append(currentGraph[lastLeg], possibleDestination)
			sort.Strings(currentGraph[lastLeg])
		}

		currentItinerary = currentItinerary[:len(currentItinerary)-1]

	}

	// Backtracking
	return nil, false
}

func deleteStringInSlice(stringSlice []string, toDelete string) []string {
	for i, val := range stringSlice {
		if val == toDelete {
			if i == len(stringSlice)-1 {
				return stringSlice[:len(stringSlice)-1]
			}
			return append(stringSlice[0:i], stringSlice[i+1:len(stringSlice)]...)
		}
	}
	return []string{}
}