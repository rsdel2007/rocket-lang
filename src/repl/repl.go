package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func runString(s string) {
	// create environment
	// tokenize
	// ast
	// eval
	// return output
}

func fileMode() {
	fileName := os.Args[1]
	b, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Unable to read file", err)
	}
	str := string(b)
	fmt.Println(str) // todo: remove this
}

const PROMT = "\x1b[32m>> \x1b[0m"

// REPL persists env values withing the same repl session
func replMode() {
	fmt.Println("Welcome to rocket REPL 👋")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		str := scanner.Text()
		fmt.Println(str) // todo: remove this
		fmt.Println()
	}
}

func Start() {
	if(len(os.Args) > 1) {
		fileMode()
	} else {
		replMode()
	}
}