package main

import (
	"bufio"
	"fmt"

	graphs "graph-assignment/graphs"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var graph *graphs.CityGraph
	outputNo := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputParams := strings.Split(scanner.Text(), " ")

		switch inputParams[0] {
		case "graph":
			graph = graphs.NewGraph()
			for _, graphInput := range inputParams[1:] {
				cityA := graphInput[0:1]
				cityB := graphInput[1:2]
				distance, _ := strconv.Atoi(graphInput[2:3])

				graph.AddTown(cityA)
				graph.AddTown(cityB)
				graph.SetRoute(cityA, cityB, uint16(distance))
			}
			outputNo = 0

		case "distance":
			distance, err := graph.CalculateRoute(inputParams[1:])

			if err != nil {
				fmt.Printf("Output #%d: %s\n", outputNo, err.Error())
			} else {
				fmt.Printf("Output #%d: %d\n", outputNo, distance)
			}

		case "numRoutesByStops":
			if len(inputParams) < 5 {
				fmt.Printf("Output #%d: %s\n", outputNo, "Missing params in input - Expected: 5")
			} else {
				operator := inputParams[1]
				start := inputParams[2]
				end := inputParams[3]
				numStops, _ := strconv.Atoi(inputParams[4])

				fmt.Printf("Output #%d: %d\n", outputNo, graph.GetNumRoutesByNumStops(operator, start, end, uint16(numStops)))
			}

		case "shortestRoute":
			if len(inputParams) < 3 {
				fmt.Printf("Output #%d: %s\n", outputNo, "Missing params in input - Expected: 3")
			} else {
				start := inputParams[1]
				end := inputParams[2]

				fmt.Printf("Output #%d: %d\n", outputNo, graph.GetLengthShortestRoute(start, end))
			}
		case "numRoutesByDistance":
			if len(inputParams) < 4 {
				fmt.Printf("Output #%d: %s\n", outputNo, "Missing params in input - Expected: 4")
			} else {
				start := inputParams[1]
				end := inputParams[2]
				maxDistance, _ := strconv.Atoi(inputParams[3])

				fmt.Printf("Output #%d: %d\n", outputNo, graph.GetNumRoutesByMaxDistance(start, end, uint16(maxDistance)))
			}
		}

		outputNo++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
