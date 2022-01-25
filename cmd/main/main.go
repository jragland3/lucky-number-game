package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"github.com/jragland3/gobootcamp/060_lucky_number/pkg/handlers/returnResult"
)

const (
	usage = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers. Your mission is to guess one of those numbers.

The greater your number is, the harder it gets

Wanna Play?`
)
var maxTurns int = 5

func isVerbose(s string) (bool, int) {
	if s == "-v" {
		return true, 1
	} else {
		return false, 0
	}
}

func main() {
	
	rand.Seed(time.Now().UnixNano())
	var doubleGuess bool

	// Check for verbosity flag
	args := os.Args[1:]
	verbose, argShift := isVerbose(args[0]) // argShift shifts the args in the case of the -v flag

	// Checks number of arguments
	if len(args) != 1 + argShift && len(args) != 2 + argShift {
		fmt.Printf(usage, maxTurns)
		return
	} else if len(args) == 2 + argShift {
		// If 2 numbers are provided a special version of the program is run that attempts the game twice in one run
		fmt.Printf("👻 You found the Double Guess feature. The second number will be tried after the first.\n")
		doubleGuess = true
	}

	guess, err := strconv.Atoi(args[0 + argShift])
	if err != nil {
		fmt.Println("Not a number")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number")
		return
	}


	// Prints results (success/failure) of the guess attempts
	handlers.ReturnResult(guess, maxTurns, args, verbose, doubleGuess)
}
