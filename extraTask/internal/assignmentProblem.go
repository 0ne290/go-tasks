package internal

import (
	"cmp"
	"slices"
)

type AssignmentProblem struct {
	costTable *SquareTable
}

func NewAssignmentProblem(costTable *SquareTable) *AssignmentProblem {
	return &AssignmentProblem{costTable}
}

func (assignmentProblem *AssignmentProblem) HungarianAlgorithm() (assignmentTable [][]bool, costTable *SquareTable, minimumCost float32) {
	costTableCopy := assignmentProblem.costTable.Copy()

	costTableCopy.subMinimumCellFromRows()
	costTableCopy.subMinimumCellFromColumns()

	graph := newBipartiteGraph(costTableCopy)
	var greatestMatching = graph.fordFulkersonAlgorithm()

	for len(greatestMatching)/2 < costTableCopy.dimension {
		leftUnvisitedNodes, rightVisitedNodes := graph.searchMinimumVertexCover(greatestMatching)
		leftVisitedNodes, rightUnvisitedNodes := graph.graphExceptMinimumVertexCover(leftUnvisitedNodes, rightVisitedNodes)

		alphaConversion(costTableCopy, graph, leftVisitedNodes, rightUnvisitedNodes)
		graph = newBipartiteGraph(costTableCopy)
		greatestMatching = graph.fordFulkersonAlgorithm()
	}

	assignmentTable = make([][]bool, costTableCopy.dimension)
	for i := 0; i < costTableCopy.dimension; i++ {
		assignmentTable[i] = make([]bool, costTableCopy.dimension)
	}
	minimumCost = 0
	for i := 0; i < costTableCopy.dimension; i++ {
		columnIndex := greatestMatching[i*2+1] - graph.rightPartIndex
		minimumCost += assignmentProblem.costTable.rows[i][columnIndex].value
		assignmentTable[i][columnIndex] = true
	}

	return assignmentTable, assignmentProblem.costTable, minimumCost
}

func alphaConversion(costTable *SquareTable, graph *BipartiteGraph, leftVisitedNodes, rightUnvisitedNodes []int) {
	selectedRows := make([]int, len(leftVisitedNodes))
	copy(selectedRows, leftVisitedNodes)
	selectedColumns := make([]int, len(rightUnvisitedNodes))
	copy(selectedColumns, rightUnvisitedNodes)

	for i := 0; i < len(selectedRows); i++ {
		selectedRows[i] -= leftPartIndex
	}

	for i := 0; i < len(selectedColumns); i++ {
		selectedColumns[i] -= graph.rightPartIndex
	}

	cells := costTable.getCellsAtTheIntersectionOfRowsAndColumns(selectedRows, selectedColumns)

	alphaCell := slices.MinFunc(cells, func(a, b *cell) int {
		return cmp.Compare(a.value, b.value)
	})
	alpha := alphaCell.value

	for _, row := range selectedRows {
		for _, cell := range costTable.rows[row] {
			cell.value -= alpha
		}
	}

	unselectedColumns := make([]int, 0, 32)
	for _, node := range graph.rightNodes {
		if !slices.Contains(rightUnvisitedNodes, node) {
			unselectedColumns = append(unselectedColumns, node)
		}
	}
	for _, column := range unselectedColumns {
		for _, cell := range costTable.columns[column-graph.rightPartIndex] {
			cell.value += alpha
		}
	}
}
