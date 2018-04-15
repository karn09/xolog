package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		run(scanner.Text())
	}
}

func run(src string) {
	scanner := bufio.NewScanner(strings.NewReader(src))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Printf("\n[line %d] Error %s : %s\n", line, where, message)
	hadError = true
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
