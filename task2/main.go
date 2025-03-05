package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/0ne290/go-tasks/task2/internal"
)

func main() {
	hasher := internal.NewSha512Hasher()
	stdinReader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter source text: ")
	source, err := stdinReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	source = strings.TrimSpace(source)

	salt := internal.Salt(source, "Any label123", 34, time.Now())
	hash := hasher.Hash(salt)

	fmt.Printf("\nSource text: %s", source)
	fmt.Printf("\nSalt result: %s", salt)
	fmt.Printf("\nHash result: %s", hash)
}