dpackage main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	startGame()
}

func startGame() {
	solved := false
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hello, what is your name?: ")

	userName, _ := reader.ReadString('\n')
	userName = strings.Replace(userName, "\n", "", -1)
	fmt.Println("---------------------")
	fmt.Println(fmt.Sprintf("Hello %s, nice to see you ; lets play a game!)", userName))
	fmt.Print(fmt.Sprintf("%s, tell me a random number: ", userName))

	number, _ := reader.ReadString('\n')
	number = strings.Replace(number, "\n", "", -1)
	rangeRandNumber, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Random number is not a valid character")
	}

	randomNumber := rand.Intn(rangeRandNumber)

	fmt.Println("")
	fmt.Println(fmt.Sprintf("Guess the number between 0 and %d", rangeRandNumber))

	for !solved {
		fmt.Println("")
		fmt.Println("---------------------")
		fmt.Print("Enter youre number: ")
		guessedNumberStr, _ := reader.ReadString('\n')
		guessedNumberStr = strings.Replace(guessedNumberStr, "\n", "", -1)
		guessedNumberInt, err := strconv.Atoi(guessedNumberStr)
		if err != nil {
			fmt.Println("Guessed number is not a valid character")
		}

		if guessedNumberInt < randomNumber {
			fmt.Println(fmt.Sprintf("Your number %d is to low! Try again!", guessedNumberInt))
		} else if guessedNumberInt > randomNumber {
			fmt.Println(fmt.Sprintf("Your number %d is to high! Try again!", guessedNumberInt))
		} else {
			solved = true
			fmt.Println("YOU WIN :)")
		}
	}
}
