// Ссылка на онлайн-сервис решения задачи о назначениях венгерским алгоритмом, чтобы вы могли проверить
// правильность работы моей программы: https://math.semestr.ru/nazn/venger.php.

package main

import (
	"github.com/0ne290/go-tasks/extraTask/internal"
	"fmt"
)

// 1    4    6    3
// 9    7    10   9
// 4    5    11   7
// 8    7    8    5
func createTestCostTable() *internal.SquareTable {
	costTable := internal.NewSquareTable(4)

	costTable.SetValue(1, 0, 0)
	costTable.SetValue(4, 0, 1)
	costTable.SetValue(6, 0, 2)
	costTable.SetValue(3, 0, 3)

	costTable.SetValue(9, 1, 0)
	costTable.SetValue(7, 1, 1)
	costTable.SetValue(10, 1, 2)
	costTable.SetValue(9, 1, 3)

	costTable.SetValue(4, 2, 0)
	costTable.SetValue(5, 2, 1)
	costTable.SetValue(11, 2, 2)
	costTable.SetValue(7, 2, 3)

	costTable.SetValue(8, 3, 0)
	costTable.SetValue(7, 3, 1)
	costTable.SetValue(8, 3, 2)
	costTable.SetValue(5, 3, 3)

	return costTable
}

func main() {
	originalCostTable := createTestCostTable()
		
		assignmentProblem := internal.NewAssignmentProblem(originalCostTable)
	    assignmentTable, costTable, minimumCost := assignmentProblem.HungarianAlgorithm();

	    dimension := costTable.GetDimension()
	    fmt.Println("Матрица стоимостей:");
	    for i := 0; i < dimension; i++ {
		    for j := 0; j < dimension; j++ {
				fmt.Printf("%f ", costTable.GetValue(i, j));
			}
			    
		    fmt.Println();
	    }
	    
	    fmt.Println();
	    fmt.Printf("Стоимость самой выгодной совокупности назначений равна %f\n", minimumCost);
	    fmt.Println();

	    dimension = len(assignmentTable)
	    fmt.Println("Матрица назначений:");
	    for i := 0; i < dimension; i++ {
		    for j := 0; j < dimension; j++ {
				fmt.Printf("%t ", assignmentTable[i][j]);
			}
			    
		    fmt.Println();
	    }
}