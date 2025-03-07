package main

import (
	"bufio"
	"fmt"
	"github.com/0ne290/go-tasks/task7/internal"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdinReader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter numbers separated by commas: ")
	source, err := stdinReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	source = strings.TrimSpace(source)

	binarySearchTree := internal.NewBinarySearchTree()

	splitedSource := strings.Split(source, ",")
	for _, text := range splitedSource {
		number, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}

		binarySearchTree.Add(number)
	}

	nlr := make([]string, 0, len(splitedSource))
	binarySearchTree.NodeLeftRight(func(value int) {
		nlr = append(nlr, strconv.Itoa(value))
	})

	lnr := make([]string, 0, len(splitedSource))
	binarySearchTree.LeftNodeRight(func(value int) {
		lnr = append(lnr, strconv.Itoa(value))
	})

	lrn := make([]string, 0, len(splitedSource))
	binarySearchTree.LeftRightNode(func(value int) {
		lrn = append(lrn, strconv.Itoa(value))
	})

	fmt.Printf("NLR: %s.\n", strings.Join(nlr, ", "))
	fmt.Printf("LNR: %s.\n", strings.Join(lnr, ", "))
	fmt.Printf("LRN: %s.", strings.Join(lrn, ", "))
}
