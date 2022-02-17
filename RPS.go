// Rock Paper Scissors by Justin
// The can be played 5 times before ending

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var possibleoptions = []string{"Rock", "Paper", "Scissors"}

func main() {

	timesplayed := 0
	for timesplayed <= 4 {
		if timesplayed == 4 {
			color.Yellow("This will be your last turn.")
		}
		if timesplayed <= 4 {
			rockpaperscissors()

			if timesplayed == 5 {
				color.Blue("Thank you for Playing." + "\n")
			}
		}
		timesplayed++
	}
}

func rockpaperscissors() {
	// User Choice
	fmt.Println("Please choose Rock, Paper, or Scissors")
	var choiceinput = bufio.NewReader(os.Stdin)
	userchoice, _ := choiceinput.ReadString('\n')

	// Computer choice
	// Seed to randomize the computer choice
	rand.Seed(time.Now().UnixNano())
	computerchoice := possibleoptions[rand.Intn(len(possibleoptions))]

	// Prints the Choices to Screen
	color.Cyan("Your choice is: " + userchoice + "\n")
	color.Magenta("The computer choice is: " + computerchoice + "\n")

	// Check for Tie
	if strings.TrimRight(userchoice, "\n") == computerchoice {
		color.Yellow("It's a tie!" + "\n")
	}

	// Calculate Winner
	if strings.TrimRight(userchoice, "\n") == "Rock" {
		if computerchoice == "Paper" {
			color.Red("Paper covers rock...computer wins!" + "\n")
		}
		if computerchoice == "Scissors" {
			color.Green("Rock crushes scissors...user wins!" + "\n")
		}
	}

	if strings.TrimRight(userchoice, "\n") == "Paper" {
		if computerchoice == "Scissors" {
			color.Red("Scissors cuts paper...computer wins!" + "\n")
		}
		if computerchoice == "Rock" {
			color.Green("Paper covers rock...user wins!" + "\n")
		}
	}

	if strings.TrimRight(userchoice, "\n") == "Scissors" {
		if computerchoice == "Rock" {
			color.Red("Rock crushes scissors...computer wins!" + "\n")
		}
		if computerchoice == "Paper" {
			color.Green("Scissors cuts paper...user wins!" + "\n")
		}
	}
}
