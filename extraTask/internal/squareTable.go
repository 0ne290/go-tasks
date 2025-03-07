package internal

type cell struct {
	value float32
}

type line = []*cell

type squareTable struct {
	dimension int
	rows      []line
	columns   []line
}

func NewSquareTable(dimension int) *squareTable {
	rows := createRows(dimension)
	columns := createColumns(rows)

	return &squareTable{dimension, rows, columns}
}

func (table* squareTable) SetValue(value float32, row, column int) {
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