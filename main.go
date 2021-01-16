package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter a letter : ")
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	if len(line) == 1 {
		// print line[0]
		fmt.Println(line[0])
	} else if len(line) == 0 {
		// print no line
		fmt.Println("Enter any alphabet.")
	} else {
		// print line is large
		fmt.Println("Enter only single alphabet.")
	}
}
