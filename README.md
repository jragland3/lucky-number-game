# gobootcamp: Lucky Number Game

The Lucky number game accepts a number and compares it to 5 psuedo-random numbers. If the generated number matches the user-entered number, you win!

use:
* `go run ./cmd/main.go <integer>`
* `-v` before the integer at runtime to enable verbosity (shows each number generated)

### [COMPLETED] EXERCISE 1: First Turn Winner
If the player wins on the first turn, then display
a special bonus message.

RESTRICTION
The first picked random number by the computer should
match the player's guess.

EXAMPLE SCENARIO
1. The player guesses 6
2. The computer picks a random number and it happens to be 6
3. Terminate the game and print the bonus message
TEST2
 
### Satisfied in ./cmd/main
 
---------------------------------------------------------
### [COMPLETED] EXERCISE 2: Random Messages

Display a few different won and lost messages "randomly".

#### HINTS
1. You can use a switch statement to do that.
2. Set its condition to the random number generator.
3. I would use a short switch.

#### EXAMPLES
The Player wins: then randomly print one of these:

go run main.go 5
   YOU WON
go run main.go 5
   YOU'RE AWESOME

The Player loses: then randomly print one of these:

go run main.go 5
   LOSER!
go run main.go 5
   YOU LOST. TRY AGAIN?

---------------------------------------------------------
#### [COMPLETED] EXERCISE 3: Double Guesses

Let the player guess two numbers instead of one.

#### HINT:
Generate random numbers using the greatest number
among the guessed numbers.

#### EXAMPLES
go run main.go 5 6
Player wins if the random number is either 5 or 6.

---------------------------------------------------------
#### [COMPLETED] EXERCISE 4: Verbose Mode

When the player runs the game like this:

go run main.go -v 4

Display each generated random number:

   1 3 4 🎉  YOU WIN!

In this example, computer picks 1, 3, and 4. And the player wins.

#### HINT
You need to get and interpret the command-line arguments.
---------------------------------------------------------

### [COMPLETED] EXERCISE 5: Enough Picks

If the player's guess number is below 10;
then make sure that the computer generates a random
number between 0 and 10.

However, if the guess number is above 10; then the
computer should generate the numbers
between 0 and the guess number.

#### WHY?
This way the game will be harder.

Because, in the current version of the game, if
the player's guess number is for example 3; then the
computer picks a random number between 0 and 3.

So, currently a player can easily win the game.

#### EXAMPLE
Suppose that the player runs the game like this:
  go run main.go 9

Or like this:
  go run main.go 2

  Then the computer should pick a random number
  between 0-10.

Or, if the player runs it like this:
  go run main.go 15

  Then the computer should pick a random number
  between 0-15.
--------------------------------------------------------
### [COMPLETED] EXERCISE 6: Dynamic Difficulty

Current game picks only 5 numbers (5 turns).

Make sure that the game adjust its own difficulty
depending on the guess number.

#### RESTRICTION
Do not make the game too easy. Only adjust the
difficulty if the guess is above 10.

#### EXPECTED OUTPUT
Suppose that the player runs the game like this:
  go run main.go 5

  Then the computer should pick 5 random numbers.

Or, if the player runs it like this:
  go run main.go 25

  Then the computer may pick 11 random numbers
  instead.

Or, if the player runs it like this:
  go run main.go 100

  Then the computer may pick 30 random numbers
  instead.

As you can see, greater guess number causes the
game to increase the game turns, which in turn
adjust the game's difficulty dynamically.

Because, greater guess number makes it harder to win.
But, this way, game's difficulty will be dynamic.
It will adjust its own difficulty depending on the
guess number.
# gobootcamp
