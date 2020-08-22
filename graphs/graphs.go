package graphs

import (
	"errors"
	"math"
)

type Town struct {
	Name   string
	Routes map[string]uint16
}

type CityGraph struct {
	Towns map[string]*Town
}

func NewGraph() *CityGraph {
	return &CityGraph{
		Towns: make(map[string]*Town),
	}
}

func (g *CityGraph) AddTown(name string) {
	existingTown := g.Towns[name]
	if existingTown != nil {
		// When town already exists
		return
	}

	newTown := &Town{
		Name:   name,
		Routes: make(map[string]uint16),
	}
	g.Towns[name] = newTown
	return
}

func (g *CityGraph) SetRoute(town1, town2 string, distance uint16) {
	g.Towns[town1].Routes[town2] = distance
}

func (g *CityGraph) generateAdjacencyMatrix() ([][]uint16, map[string]uint8) {
	labelIndices := make(map[string]uint8, len(g.Towns))
	var index uint8 = 0
	for k := range g.Towns {
		labelIndices[k] = index
		index++
	}

	matrix := make([][]uint16, len(g.Towns))
	for i := 0; i < len(g.Towns); i++ {
		matrix[i] = make([]uint16, len(g.Towns))
	}

	for townLabel, currentTown := range g.Towns {
		for adjacentTown := range currentTown.Routes {
			matrix[labelIndices[townLabel]][labelIndices[adjacentTown]] = 1
		}
	}

	return matrix, labelIndices
}

func (g *CityGraph) calculateRouteHelper(towns []string, distance uint16, current *Town) (uint16, error) {
	if current == nil {
		return 0, errors.New("Error trying to start from a town that does not exist in the city graph")
	}

	if len(towns) == 0 {
		return distance, nil
	}

	if nextTown, ok := current.Routes[towns[0]]; ok {
		return g.calculateRouteHelper(towns[1:], distance+nextTown, g.Towns[towns[0]])
	}

	return 0, errors.New("NO SUCH ROUTE")
}

func (g *CityGraph) CalculateRoute(towns []string) (uint16, error) {
	return g.calculateRouteHelper(towns[1:], 0, g.Towns[towns[0]])
}

func (g *CityGraph) GetNumRoutesByNumStops(operator, start, end string, numStops uint16) uint16 {
	var accumulator uint16 = 0
	matrix, labelIndices := g.generateAdjacencyMatrix()
	multipliedMatrix := matrix

	// currentStops := 1 because the adjacency matrix already represents 1 stop.
	// also assuming that 0 stops means we should return 0
	for currentStops := 1; currentStops <= int(numStops); currentStops++ {
		if operator == "max" || (operator == "eq" && currentStops == int(numStops)) {
			accumulator += multipliedMatrix[labelIndices[start]][labelIndices[end]]
		}
		if currentStops != int(numStops) {
			multipliedMatrix = matrixMultiplication(multipliedMatrix, matrix)
		}
	}
	return accumulator
}

func (g *CityGraph) GetLengthShortestRoute(start, end string) uint16 {
	// This is just a slight modification on dijkstra's algorithm
	// returns 65535 (the largest unsigned 16-bit int) if no route exists,
	// which is kind of like returning infinity

	var currentTown *Town
	visited := make(map[string]bool, len(g.Towns))
	distances := make(map[string]uint16, len(g.Towns))
	for k := range g.Towns {
		distances[k] = math.MaxUint16
	}

	// Do some work with the origin node, but do not add it to the visited nodes set
	currentTown = g.Towns[start]
	if currentTown == nil || g.Towns[end] == nil {
		// Early return for invalid/missing town names
		return math.MaxUint16
	}

	currentTownName := ""
	for townName, distance := range currentTown.Routes {
		distances[townName] = distance
		if currentTownName == "" || distance < distances[currentTownName] {
			// This is to get the first closest node from the origin
			currentTownName = townName
		}
	}

	for currentTownName != end {
		visited[currentTownName] = true
		currentTown = g.Towns[currentTownName]

		// Set new min distances
		for adjacent, distance := range currentTown.Routes {
			distances[adjacent] = min(distances[currentTownName]+distance, distances[adjacent])
		}

		// Get next unvisited town with min distance from origin
		// we default currentTownName to end so that in the event we reach a dead end, we don't try and continue on a different path
		currentTownName = end
		for townName, distance := range distances {
			if _, ok := visited[townName]; !ok && distance < distances[currentTownName] {
				currentTownName = townName
			}
		}
	}
	return distances[end]
}

func (g *CityGraph) GetNumRoutesByMaxDistance(start, end string, maxDistance uint16) uint16 {
	return g.getNumRoutesByMaxDistanceHelper(start, end, maxDistance, 0, 0)
}

func (g *CityGraph) getNumRoutesByMaxDistanceHelper(current, end string, maxDistance, currentDistance, numRoutes uint16) uint16 {
	if currentDistance >= maxDistance {
		return numRoutes
	}

	if current == end && currentDistance > 0 {
		numRoutes++
	}

	if currentTown, ok := g.Towns[current]; ok {
		// This is to avoid nil pointer errors
		for nextTown, distance := range currentTown.Routes {
			numRoutes += g.getNumRoutesByMaxDistanceHelper(nextTown, end, maxDistance, currentDistance+distance, 0)
		}
	}

	return numRoutes

}

//  Non-struct-based helpers
func min(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}

func matrixMultiplication(m1, m2 [][]uint16) [][]uint16 {
	// Since we are only using this to multiply square matrices, won't bother with error handling
	// eg: if two matrices cannot be multiplied

	newMatrix := make([][]uint16, len(m1))
	for i := 0; i < len(m1); i++ {
		newMatrix[i] = make([]uint16, len(m1))
	}

	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m1); j++ {
			var total uint16 = 0
			for k := 0; k < len(m1); k++ {
				total = total + m1[i][k]*m2[k][j]
			}
			newMatrix[i][j] = total
		}
	}
	return newMatrix
}
