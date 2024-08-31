package workbench

type dGraph struct { // directed graph
	edges map[int][]int
}

func NewGraph() *dGraph {
	return &dGraph{
		edges: make(map[int][]int),
	}
}

// adds a directed edge to the graph from v -> w
func (g dGraph) AddEdge(v, w int) {
	g.edges[v] = append(g.edges[v], w)
}

func (g *dGraph) BFS(start int) []int {
	visited := NewSet()
	queue := NewQueue()

	visited.Add(start)
	results := []int{}

	for queue.Size() > 0 {
		_, _ = queue.Dequeue()

	}

	return results
}
