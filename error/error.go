package error

import (
	"fmt"
	"os"
)

func Error(line int, message string) {
	report(line, "", message)
	os.Exit(65)
}

func report(line int, where string, message string) {
	fmt.Printf("\n[line %d] Error %s : %s\n", line, where, message)
}
