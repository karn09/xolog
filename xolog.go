package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"xolog/scanner"
)

var (
	hadError bool
)

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Xolog REPL")
	fmt.Println("------------------")
	fmt.Print("> ")

	for scanner.Scan() {
		run(scanner.Text())
		fmt.Print("> ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func runFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	buf := bytes.Buffer{}
	buf.ReadFrom(file)
	content := buf.String()

	run(content)
	if hadError {
		os.Exit(65)
	}

}

func run(src string) {
	s := scanner.NewScanner(src)
	tokens := s.ScanTokens()
	for _, token := range tokens {
		fmt.Println(token)
	}
	if s.HadError {
		hadError = true
	}
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: xolog [script]")
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}
