package graphs_test

import (
	"graph-assignment/graphs"
	"testing"
)

func GenerateGraph() *graphs.CityGraph {
	testGraph := graphs.NewGraph()
	testGraph.AddTown("A")
	testGraph.AddTown("B")
	testGraph.AddTown("C")
	testGraph.AddTown("D")
	testGraph.AddTown("E")
	testGraph.SetRoute("A", "B", 5)
	testGraph.SetRoute("B", "C", 4)
	testGraph.SetRoute("C", "D", 8)
	testGraph.SetRoute("D", "C", 8)
	testGraph.SetRoute("D", "E", 6)
	testGraph.SetRoute("A", "D", 5)
	testGraph.SetRoute("C", "E", 2)
	testGraph.SetRoute("E", "B", 3)
	testGraph.SetRoute("A", "E", 7)
	return testGraph
}

func TestDistances(t *testing.T) {
	var distance, expected uint16
	var err error
	testGraph := GenerateGraph()

	expected = 9
	distance, err = testGraph.CalculateRoute([]string{"A", "B", "C"})
	if distance != expected {
		t.Errorf("distance was incorrect, got: %d, want: %d.", distance, expected)
	}

	expected = 5
	distance, err = testGraph.CalculateRoute([]string{"A", "D"})
	if distance != expected {
		t.Errorf("distance was incorrect, got: %d, want: %d.", distance, expected)
	}

	expected = 13
	distance, err = testGraph.CalculateRoute([]string{"A", "D", "C"})
	if distance != expected {
		t.Errorf("distance was incorrect, got: %d, want: %d.", distance, expected)
	}

	expected = 22
	distance, err = testGraph.CalculateRoute([]string{"A", "E", "B", "C", "D"})
	if distance != expected {
		t.Errorf("distance was incorrect, got: %d, want: %d.", distance, expected)
	}

	expected = 0
	distance, err = testGraph.CalculateRoute([]string{"A", "E", "D"})
	if err == nil {
		t.Errorf("Expected no route to be found")
	}
}

func TestNumRoutesByStops(t *testing.T) {
	var numRoutes, expected uint16
	testGraph := GenerateGraph()

	expected = 2
	numRoutes = testGraph.GetNumRoutesByNumStops("max", "C", "C", 3)
	if numRoutes != expected {
		t.Errorf("numRoutes was incorrect, got: %d, want: %d.", numRoutes, expected)
	}

	expected = 3
	numRoutes = testGraph.GetNumRoutesByNumStops("eq", "A", "C", 4)
	if numRoutes != expected {
		t.Errorf("numRoutes was incorrect, got: %d, want: %d.", numRoutes, expected)
	}

}

func TestShortestRoute(t *testing.T) {
	var distance, expected uint16
	testGraph := GenerateGraph()

	expected = 9
	distance = testGraph.GetLengthShortestRoute("A", "C")
	if distance != expected {
		t.Errorf("distance was incorrect, got: %d, want: %d.", distance, expected)
	}

	expected = 9
	distance = testGraph.GetLengthShortestRoute("B", "B")
	if distance != expected {
		t.Errorf("distance was incorrect, got: %d, want: %d.", distance, expected)
	}

}

func TestNumRoutesByDistance(t *testing.T) {
	var numRoutes, expected uint16
	testGraph := GenerateGraph()

	expected = 7
	numRoutes = testGraph.GetNumRoutesByMaxDistance("C", "C", 30)
	if numRoutes != expected {
		t.Errorf("numRoutes was incorrect, got: %d, want: %d.", numRoutes, expected)
	}

}
