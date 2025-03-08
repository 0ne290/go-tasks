package internal

type AssignmentProblem struct {
	costTable *SquareTable
}

func NewAssignmentProblem(costTable *SquareTable) *AssignmentProblem {
	return &AssignmentProblem{costTable}
}

func (assignmentProblem *AssignmentProblem) HungarianAlgorithm() {
	costTableCopy := assignmentProblem.costTable.Copy()

	costTableCopy.subMinimumCellFromRows()
	costTableCopy.subMinimumCellFromColumns()

	bipartiteGraph := newBipartiteGraph(costTableCopy)
	var greatestMatching = bipartiteGraph.fordFulkersonAlgorithm()

	for len(greatestMatching)/2 < costTableCopy.dimension {
		/*var minimumVertexCover = _bipartiteGraph.SearchMinimumVertexCover(greatestMatching);
		  var graphExceptMinimumVertexCover = GraphExceptMinimumVertexCover(minimumVertexCover);
		  AlphaConversion(costTable, graphExceptMinimumVertexCover);
		  _bipartiteGraph = new BipartiteGraph(costTable);
		  greatestMatching = _bipartiteGraph.FordFulkersonAlgorithm();
		  logger.AddTable(costTable);
		  logger.AddMatching(greatestMatching);*/
	}

	assignmentTable := make([][]bool, costTableCopy.dimension)
	for i := 0; i < costTableCopy.dimension; i++ {
		assignmentTable[i] = make([]bool, costTableCopy.dimension)
	}
	var minimumCost float32 = 0
	for i := 0; i < costTableCopy.dimension; i++ {
		columnIndex := greatestMatching[i*2+1] - bipartiteGraph.rightPartIndex
		minimumCost += costTableCopy.rows[i][columnIndex].value
		assignmentTable[i][columnIndex] = true
	}

	/*var result = new TheAssignmentProblemDto(assignmentTable, CostTable, minimumCost);

	  logger.Dispose();

	  return result;*/
}

/*private IDictionary<string, ISet<int>> GraphExceptMinimumVertexCover(IDictionary<string, ISet<int>> minimumVertexCover)
  {
      var graphExceptMinimumVertexCover = new Dictionary<string, ISet<int>>
      {
          { "leftNodes", _bipartiteGraph.LeftNodes },
          { "rightNodes", _bipartiteGraph.RightNodes }
      };

      graphExceptMinimumVertexCover["leftNodes"].ExceptWith(minimumVertexCover["leftNodes"]);
      graphExceptMinimumVertexCover["rightNodes"].ExceptWith(minimumVertexCover["rightNodes"]);

      return graphExceptMinimumVertexCover;
  }

  private void AlphaConversion(Table<double> costTable, IDictionary<string, ISet<int>> graphExceptMinimumVertexCover)
  {
      var selectedRows = graphExceptMinimumVertexCover["leftNodes"].ToArray();
      var selectedColumns = graphExceptMinimumVertexCover["rightNodes"].ToArray();

      for (var i = 0; i < selectedRows.Length; i++)
          selectedRows[i] -= BipartiteGraph.LeftPartIndex;
      for (var i = 0; i < selectedColumns.Length; i++)
          selectedColumns[i] -= _bipartiteGraph.RightPartIndex;

      var cells = costTable.GetCellsAtTheIntersectionOfRowsAndColumns(selectedRows, selectedColumns);

      var alphaCell = cells.Min();
      if (alphaCell == null)
          throw new ArgumentNullException(nameof(alphaCell));
      var alpha = alphaCell.Value;

      foreach (var row in selectedRows)
          foreach (var cell in costTable.Rows[row])
              cell.Value -= alpha;

      var unselectedColumns = _bipartiteGraph.RightNodes;
      unselectedColumns.ExceptWith(graphExceptMinimumVertexCover["rightNodes"]);
      foreach (var column in unselectedColumns)
          foreach (var cell in costTable.Columns[column - _bipartiteGraph.RightPartIndex])
              cell.Value += alpha;
  }
*/
