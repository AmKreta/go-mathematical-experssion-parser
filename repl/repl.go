package repl

import (
	"bufio"
	"fmt"
	"os"
)

func Repl() {
	for {
		fmt.Print(">>> enter expression: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}
		result := eval(input)
		fmt.Println("result: ", result)
		fmt.Println()
	}
}
