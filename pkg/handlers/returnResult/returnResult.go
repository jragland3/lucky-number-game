package handlers

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/jragland3/gobootcamp/060_lucky_number/pkg/handlers/randomMsgs"
)

// Used to compare the guess to the random numbers generated and print the results
func ReturnResult(guess, maxTurns int, args []string, verbose, doubleGuess bool) {
	// Randomize messages
	win, lose := handlers.SetRandomMsg()

	if doubleGuess {
		fmt.Printf("First Round: ")
	}

	for turn := maxTurns + guess/4; turn > 0; turn-- {
		var n int
		if guess < 10 {
			n = rand.Intn(11)
		} else {
			n = rand.Intn(guess + 1)
		}

		if verbose {
			fmt.Printf("%d ", n)
		}
		
		// A special message displays if the user wins on the first attempt of the first guess
		if n == guess && turn == maxTurns + guess/4 {
			fmt.Printf("🙌🙌🙌 You done went and won on your first try. Have an NFT: 🎂 🙌🙌🙌\n\n")
			break
		} else if n == guess {
			fmt.Println(win)
			break
		} else if n != guess && (turn - 1) == 0 {
			fmt.Println(lose)
		}
	}

	// The results will display for the second guess if doubleGuess == true
	if doubleGuess {
		var (
			guess2 int
			err error
		)
		if verbose {
			guess2, err = strconv.Atoi(args[2])
		} else {
			guess2, err = strconv.Atoi(args[1])
		}
		if err != nil {
			fmt.Println("Not a number")
			return
		}
		
		win, lose := handlers.SetRandomMsg()

		fmt.Printf("Second Round: ")
		
		for turn := maxTurns + guess2/4; turn > 0; turn-- {
			var n int
			if guess2 < 10 {
				n = rand.Intn(11)
			} else {
				n = rand.Intn(guess2 + 1)
			}

			if verbose {
				fmt.Printf("%d ", n)
			}
			
			if n == guess2 {
				fmt.Println(win)
				break
			} else if n != guess2 && (turn - 1) == 0 {
				fmt.Println(lose)
			}
		}
	}
}
