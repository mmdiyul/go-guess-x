package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	epochTime := strconv.FormatInt(now.UnixMilli(), 10)
	splitEpoch := strings.Split(epochTime, "")
	xValue := splitEpoch[11]
	splitEpoch[11] = "X"
	quest := strings.Join(splitEpoch, "")
	totalFailed := 0
	fmt.Println("Hi, You have 3 attempts!")
	fmt.Printf("Current time: %v\n", now.Local())
	guess(xValue, quest, totalFailed)
}

func guess(x string, quest string, totalFailed int) {
	if totalFailed != 0 {
		fmt.Printf("Total failed: %d times\n", totalFailed)
	}
	fmt.Println("Quest: " + quest)
	var guessValue string
	fmt.Print("Guess X value (X is number): ")
	fmt.Scanf("%s", &guessValue)
	if guessValue != x {
		if totalFailed+1 == 3 {
			fmt.Println("Game over!")
			return
		} else {
			fmt.Println("Try again")
		}
		guess(x, quest, totalFailed+1)
	} else {
		fmt.Println("Congrats! You win!")
	}
}
