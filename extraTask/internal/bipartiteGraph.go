package internal

const (
	leftPartIndex int = 1
	sourceIndex       = 0
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

func newBipartiteGraph(table *squareTable) *BipartiteGraph {
	graph := &BipartiteGraph{}

	graph.preBuild(table)
	graph.buildEdges(table)
	graph.buildAdjacencyLists()

	return graph
}

func (graph *BipartiteGraph) preBuild(table *squareTable) {
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

func (graph *BipartiteGraph) buildEdges(table *squareTable) {
	graph.edges = make([][]*edge, graph.nodesCount)
	for i := 0; i < graph.nodesCount; i++ {
		graph.edges[i] = make([]*edge, graph.nodesCount)
	}

	directedGraphAdjacencyMatrix := graph.createDirectedGraphAdjacencyMatrix(table)

	for i := 0; i < graph.nodesCount; i++ {
		for j := 0; j < graph.nodesCount; j++ {
			if directedGraphAdjacencyMatrix[i][j] {
				graph.edges[i][j] = &edge{1, 1, false}
				graph.edges[j][i] = &edge{0, 0, true}
			} else if !graph.adjacencyMatrix[i][j] {
				graph.edges[i][j] = &edge{0, 0, false}
			}
		}
	}
}

func (graph *BipartiteGraph) createDirectedGraphAdjacencyMatrix(table *squareTable) [][]bool {
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
				if j != 0 && j != graph.nodesCount - 1 {
					graph.graphAdjacencyLists[i] = append(graph.graphAdjacencyLists[i], j)
				}
					
			}
		}

		graph.graphAdjacencyLists[0] = make([]int, 0)
		graph.graphAdjacencyLists[graph.nodesCount - 1] = make([]int, 0)
}

func (graph *BipartiteGraph) fordFulkersonAlgorithm() []int {
		//path := FindPathToNode(0, graph.nodesCount - 1, new NodeStack());
		var path []int
		for len(path) > 1 {
			for i := 0; i < len(path) - 1; i++ {
				graph.edges[path[i]][path[i + 1]].sendFlow(1);
				graph.edges[path[i + 1]][path[i]].receiveFlow(1);
			}
			//path = FindPathToNode(0, NumberOfNodes - 1, new NodeStack());
		}

		//var greatestMatching = GetReverseFreeEdges();
		//RestoreEdges();
		var greatestMatching []int

		return greatestMatching;
	}

	/*func findPathToNode(int startNode, int targetNode, INodeStorage nodesStorage) []int {
		path = new List<int>();

		var searchRoute = NodeSearch(startNode, targetNode, _transportNetworkAdjacencyLists, nodesStorage);
		if (searchRoute[^1] != targetNode)
			return new List<int>(0);

		var node = searchRoute[^1];
		path.Add(node);

		for (var i = searchRoute.Count - 2; i > -1; i--)
		{
			if (!_adjacencyMatrix[searchRoute[i], node])
				continue;
			node = searchRoute[i];
			path.Add(node);
		}

		path.TrimExcess();
		path.Reverse();

		return path;
	}

	private IList<int> NodeSearch(int startNode, int targetNode, IList<IList<int>> adjacencyLists, INodeStorage nodesStorage)
	{
		var visited = new bool[NumberOfNodes];

		nodesStorage.Insert(startNode);

		var visitedNodes = new List<int>();

		while (!nodesStorage.IsEmpty())
		{
			var currentNode = nodesStorage.GetFirst();

			if (visited[currentNode])
				continue;

			visited[currentNode] = true;
			visitedNodes.Add(currentNode);

			if (currentNode == targetNode)
				return visitedNodes;

			var neighbours = adjacencyLists[currentNode];
			foreach (var nodeToGo in neighbours)
			{
				if (visited[nodeToGo] || GetEdge(currentNode, nodeToGo).IsBusy())
					continue;

				nodesStorage.Insert(nodeToGo);
			}
		}

		return visitedNodes;
	}

	private IList<int> GetReverseFreeEdges()
	{
		var edges = new List<int>();

		for (var i = LeftPartIndex; i < RightPartIndex; i++)
		{
			for (var j = RightPartIndex; j < _drainIndex; j++)
			{
				if (!GetEdge(j, i).IsBusy() && GetEdge(j, i).IsReverse)
				{
					edges.Add(i);
					edges.Add(j);
				}
			}
		}

		return edges;
	}

	private void RestoreEdges()
	{
		for (var i = 0; i < NumberOfNodes; i++)
			for (var j = 0; j < NumberOfNodes; j++)
				GetEdge(i, j).ResetСapacity();
	}

	public IDictionary<string, ISet<int>> SearchMinimumVertexCover(IList<int> greatestMatching)
	{
		for (var i = 0; i < greatestMatching.Count; i += 2)
		{
			GetEdge(greatestMatching[i], greatestMatching[i + 1]).SendFlow(1);
			GetEdge(greatestMatching[i + 1], greatestMatching[i]).ReceiveFlow(1);
		}

		var leftUnvisitedNodes = LeftNodes;
		var minimumVertexCover = new Dictionary<string, ISet<int>>();

		var rightVisitedNodes = new HashSet<int>();

		for (var i = LeftPartIndex; i < RightPartIndex; i++)
		{
			if (greatestMatching.Contains(i))
				continue;
			var searchRoute = NodeSearch(i, -1, _graphAdjacencyLists, new NodeStack());
			foreach (var node in searchRoute)
			{
				leftUnvisitedNodes.Remove(node);
				if (node >= RightPartIndex)
					rightVisitedNodes.Add(node);
			}
		}

		RestoreEdges();

		minimumVertexCover.Add("leftNodes", leftUnvisitedNodes);
		minimumVertexCover.Add("rightNodes", rightVisitedNodes);

		return minimumVertexCover;
	}*/