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
	found := false

	for i := 0; i < 10; i++ {
		fmt.Print(" Enter a alphabet character : ")
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
		}

		if len(line) == 1 {
			if int(line[0]) >= int('a') && int(line[0]) <= int('z') {
				if int(line[0]) == num {
					fmt.Println("Yeah!!! You are guessed that letter correctly in", i+1, "attempt.")
					found = true // user already guessed the answer
					break
				} else if int(line[0]) > num {
					fmt.Println("Your guess is too high.")
				} else if int(line[0]) < num {
					fmt.Println("Your guess is too low.")
				}
			} else {
				// if user entered other than letters from a to z
				fmt.Println("Enter only small letter alphabet character (a-z).")
				continue
			}
		} else if len(line) == 0 {
			// if user not provided input or given empty
			fmt.Println("Enter any alphabet character.")
			continue // continue to ask input
		} else {
			// if user entered more than one letter
			fmt.Println("Enter only single alphabet.")
			continue // continue asking user input
		}
	}
	// check user guess
	if !found {
		fmt.Println("You are not guess it. The actual letter is \"" + string(rune(num)) + "\".")
	}
}

func getRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
