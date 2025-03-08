package internal

import (
	"slices"
)

const (
	leftPartIndex int = 1
	sourceIndex   int = 0
)

type edge struct {
	capacity, weight int
	isReverse        bool
}

func newEdge(capacity int, isReverse bool) *edge {
	return &edge{capacity, capacity, isReverse}
}

func (edge *edge) sendFlow(flow int) {
	edge.capacity -= flow
}

func (edge *edge) receiveFlow(flow int) {
	edge.capacity += flow
}

func (edge *edge) resetСapacity() {
	edge.capacity = edge.weight
}

func (edge *edge) isBusy() bool {
	return edge.capacity < 1
}

type BipartiteGraph struct {
	rightPartIndex                 int
	nodesCount                     int
	adjacencyMatrix                [][]bool
	transportNetworkAdjacencyLists [][]int
	graphAdjacencyLists            [][]int
	edges                          [][]*edge
	leftNodes                      []int
	rightNodes                     []int
	drainIndex                     int
}

func newBipartiteGraph(table *SquareTable) *BipartiteGraph {
	graph := &BipartiteGraph{}

	graph.preBuild(table)
	graph.buildEdges(table)
	graph.buildAdjacencyLists()

	return graph
}

func (graph *BipartiteGraph) preBuild(table *SquareTable) {
	graph.nodesCount = table.dimension*2 + 2
	graph.adjacencyMatrix = make([][]bool, graph.nodesCount)
	for i := 0; i < graph.nodesCount; i++ {
		graph.adjacencyMatrix[i] = make([]bool, graph.nodesCount)
	}

	graph.rightPartIndex = table.dimension + 1
	graph.drainIndex = graph.nodesCount - 1

	graph.leftNodes = make([]int, 0, 32)
	for i := leftPartIndex; i < graph.rightPartIndex; i++ {
		graph.leftNodes = append(graph.leftNodes, i)
	}

	graph.rightNodes = make([]int, 0, 32)
	for i := graph.rightPartIndex; i < graph.drainIndex; i++ {
		graph.rightNodes = append(graph.rightNodes, i)
	}
}

func (graph *BipartiteGraph) buildEdges(table *SquareTable) {
	graph.edges = make([][]*edge, graph.nodesCount)
	for i := 0; i < graph.nodesCount; i++ {
		graph.edges[i] = make([]*edge, graph.nodesCount)
	}

	directedGraphAdjacencyMatrix := graph.createDirectedGraphAdjacencyMatrix(table)

	for i := 0; i < graph.nodesCount; i++ {
		for j := 0; j < graph.nodesCount; j++ {
			if directedGraphAdjacencyMatrix[i][j] {
				graph.edges[i][j] = newEdge(1, false)
				graph.edges[j][i] = newEdge(0, true)
			} else if !graph.adjacencyMatrix[i][j] {
				graph.edges[i][j] = &edge{0, 0, false}
			}
		}
	}
}

func (graph *BipartiteGraph) createDirectedGraphAdjacencyMatrix(table *SquareTable) [][]bool {
	directedGraphAdjacencyMatrix := make([][]bool, graph.nodesCount)
	for i := 0; i < graph.nodesCount; i++ {
		directedGraphAdjacencyMatrix[i] = make([]bool, graph.nodesCount)
	}

	for i := graph.rightPartIndex; i < graph.drainIndex; i++ {
		directedGraphAdjacencyMatrix[i][graph.drainIndex] = true
		graph.adjacencyMatrix[i][graph.drainIndex] = true
		graph.adjacencyMatrix[graph.drainIndex][i] = true
	}

	k := 0
	for i := leftPartIndex; i < graph.rightPartIndex; i++ {
		directedGraphAdjacencyMatrix[sourceIndex][i] = true
		graph.adjacencyMatrix[sourceIndex][i] = true
		graph.adjacencyMatrix[i][sourceIndex] = true
		for j := graph.rightPartIndex; j < graph.drainIndex; j++ {
			directedGraphAdjacencyMatrix[i][j] = (table.rows[k][j-graph.rightPartIndex].value) == 0
			graph.adjacencyMatrix[i][j] = (table.rows[k][j-graph.rightPartIndex].value) == 0
			graph.adjacencyMatrix[j][i] = (table.rows[k][j-graph.rightPartIndex].value) == 0
		}
		k++
	}

	return directedGraphAdjacencyMatrix
}

func (graph *BipartiteGraph) buildAdjacencyLists() {
	graph.transportNetworkAdjacencyLists = make([][]int, graph.nodesCount)
	graph.graphAdjacencyLists = make([][]int, graph.nodesCount)

	for i := 0; i < graph.nodesCount; i++ {
		graph.transportNetworkAdjacencyLists[i] = make([]int, 0, graph.nodesCount)
		graph.graphAdjacencyLists[i] = make([]int, 0, graph.nodesCount)
		for j := 0; j < graph.nodesCount; j++ {
			if !graph.adjacencyMatrix[i][j] {
				continue
			}

			graph.transportNetworkAdjacencyLists[i] = append(graph.transportNetworkAdjacencyLists[i], j)
			if j != 0 && j != graph.nodesCount-1 {
				graph.graphAdjacencyLists[i] = append(graph.graphAdjacencyLists[i], j)
			}

		}
	}

	graph.graphAdjacencyLists[0] = make([]int, 0)
	graph.graphAdjacencyLists[graph.nodesCount-1] = make([]int, 0)
}

func (graph *BipartiteGraph) fordFulkersonAlgorithm() []int {
	path := graph.findPathToNode(0, graph.nodesCount-1)
	for len(path) > 1 {
		for i := 0; i < len(path)-1; i++ {
			graph.edges[path[i]][path[i+1]].sendFlow(1)
			graph.edges[path[i+1]][path[i]].receiveFlow(1)
		}
		path = graph.findPathToNode(0, graph.nodesCount-1)
	}

	var greatestMatching = graph.getReverseFreeEdges()
	graph.restoreEdges()

	return greatestMatching
}

func (graph *BipartiteGraph) findPathToNode(startNode, targetNode int) []int {
	path := make([]int, 0, 32)

	var searchRoute = graph.nodeSearch(startNode, targetNode)
	if searchRoute[len(searchRoute)-1] != targetNode {
		return make([]int, 0)
	}

	node := searchRoute[len(searchRoute)-1]
	path = append(path, node)

	for i := len(searchRoute) - 2; i > -1; i-- {
		if !graph.adjacencyMatrix[searchRoute[i]][node] {
			continue
		}

		node = searchRoute[i]
		path = append(path, node)
	}

	reverse(path)

	return path
}

func reverse(numbers []int) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

func (graph *BipartiteGraph) nodeSearch(startNode, targetNode int) []int {
	visited := make([]bool, graph.nodesCount)
	visitedNodes := make([]int, 0, 32)
	nodesStack := NewStack[int]()
	nodesStack.Push(startNode)

	for !nodesStack.IsEmpty() {
		currentNode, _ := nodesStack.Pop()

		if visited[currentNode] {
			continue
		}

		visited[currentNode] = true
		visitedNodes = append(visitedNodes, currentNode)

		if currentNode == targetNode {
			return visitedNodes
		}

		var neighbours = graph.transportNetworkAdjacencyLists[currentNode]
		for _, nodeToGo := range neighbours {
			if visited[nodeToGo] || graph.edges[currentNode][nodeToGo].isBusy() {
				continue
			}
			nodesStack.Push(nodeToGo)
		}
	}

	return visitedNodes
}

func (graph *BipartiteGraph) getReverseFreeEdges() []int {
	var edges = make([]int, 0, 32)

	for i := leftPartIndex; i < graph.rightPartIndex; i++ {
		for j := graph.rightPartIndex; j < graph.drainIndex; j++ {
			if !graph.edges[j][i].isBusy() && graph.edges[j][i].isReverse {
				edges = append(edges, i)
				edges = append(edges, j)
			}
		}
	}

	return edges
}

func (graph *BipartiteGraph) restoreEdges() {
	for i := 0; i < graph.nodesCount; i++ {
		for j := 0; j < graph.nodesCount; j++ {
			graph.edges[i][j].resetСapacity()
		}
	}
}

func (graph *BipartiteGraph) searchMinimumVertexCover(greatestMatching []int) (leftUnvisitedNodes, rightVisitedNodes []int) {
	for i := 0; i < len(greatestMatching); i += 2 {
		graph.edges[greatestMatching[i]][greatestMatching[i+1]].sendFlow(1)
		graph.edges[greatestMatching[i+1]][greatestMatching[i]].receiveFlow(1)
	}

	leftUnvisitedNodes = make([]int, len(graph.leftNodes))
	copy(leftUnvisitedNodes, graph.leftNodes)
	rightVisitedNodes = make([]int, 0, 32)

	for i := leftPartIndex; i < graph.rightPartIndex; i++ {
		if slices.Contains(greatestMatching, i) {
			continue
		}

		searchRoute := graph.nodeSearch(i, -1)
		for _, node := range searchRoute {
			leftUnvisitedNodes = slices.DeleteFunc(leftUnvisitedNodes, func(a int) bool {
				return a == node
			})
			if node >= graph.rightPartIndex {
				rightVisitedNodes = append(rightVisitedNodes, node)
			}
		}
	}

	graph.restoreEdges()

	return
}

func (graph *BipartiteGraph) graphExceptMinimumVertexCover(leftUnvisitedNodes, rightVisitedNodes []int) (leftVisitedNodes, rightUnvisitedNodes []int) {
	leftVisitedNodes = make([]int, 0, 32)
	rightUnvisitedNodes = make([]int, 0, 32)

	for _, node := range graph.leftNodes {
		if !slices.Contains(leftUnvisitedNodes, node) {
			leftVisitedNodes = append(leftVisitedNodes, node)
		}
	}
	for _, node := range graph.rightNodes {
		if !slices.Contains(rightVisitedNodes, node) {
			rightUnvisitedNodes = append(rightUnvisitedNodes, node)
		}
	}

	return
}
