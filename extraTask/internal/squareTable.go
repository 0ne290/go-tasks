package internal

import (
	"cmp"
	"slices"
)

type cell struct {
	value float32
}

type line = []*cell

type SquareTable struct {
	dimension int
	rows      []line
	columns   []line
}

func NewSquareTable(dimension int) *SquareTable {
	rows := createRows(dimension)
	columns := createColumns(rows)

	return &SquareTable{dimension, rows, columns}
}

func (source *SquareTable) Copy() *SquareTable {
	dimension := source.dimension
	copy := NewSquareTable(dimension)

	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			copy.rows[i][j].value = source.rows[i][j].value
		}
	}

	return copy
}

func (table *SquareTable) SetValue(value float32, row, column int) {
	table.rows[row][column].value = value
}

func createRows(dimension int) []line {
	rows := make([]line, dimension)
	for i := 0; i < dimension; i++ {
		rows[i] = make(line, dimension)
		for j := 0; j < dimension; j++ {
			rows[i][j] = &cell{}
		}
	}

	return rows
}

func createColumns(rows []line) []line {
	dimension := len(rows)

	columns := make([]line, dimension)
	for i := 0; i < dimension; i++ {
		columns[i] = make(line, dimension)
		for j := 0; j < dimension; j++ {
			columns[i][j] = rows[j][i]
		}
	}

	return columns
}

func (table *SquareTable) subMinimumCellFromRows() {
	for i := 0; i < table.dimension; i++ {
		row := table.rows[i]
		rowMinimumCell := slices.MinFunc(row, func(a, b *cell) int {
			return cmp.Compare(a.value, b.value)
		})
		rowMinimum := rowMinimumCell.value
		for j := 0; j < table.dimension; j++ {
			row[j].value -= rowMinimum
		}
	}
}

func (table *SquareTable) subMinimumCellFromColumns() {
	for i := 0; i < table.dimension; i++ {
		column := table.columns[i]
		columnMinimumCell := slices.MinFunc(column, func(a, b *cell) int {
			return cmp.Compare(a.value, b.value)
		})
		columnMinimum := columnMinimumCell.value
		for j := 0; j < table.dimension; j++ {
			column[j].value -= columnMinimum
		}
	}
}