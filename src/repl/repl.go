package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"rocket/src/lexer"
	"rocket/src/token"
)

func runString(s string) {
	// todo: take environment as argument
	l := lexer.New(s)
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%s (%s) | %d:%d\n", tok.Literal, tok.Type, tok.Ln, tok.Col)
	}
}

func fileMode() {
	fileName := os.Args[1]
	b, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Unable to read file", err)
	}
	str := string(b)
	runString(str)
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
		runString(str)
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