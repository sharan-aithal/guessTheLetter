package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// get random number from "a" to "z"
	num := getRandomInt(int('a'), int('z'))
	fmt.Print("Enter a alphabet character : ")
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	if len(line) == 1 {
		// print line[0]
		fmt.Println("Your alphabet is :", string(line[0]))
		if int(line[0]) >= int('a') && int(line[0]) <= int('z') {
			fmt.Println("Entered letter is in the range")
		} else {
			fmt.Println("Enter only small letter alphabet character (a-z)")
		}
	} else if len(line) == 0 {
		// print no line
		fmt.Println("Enter any alphabet character.")
	} else {
		// print line is large
		fmt.Println("Enter only single alphabet.")
	}
}

func getRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
