package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var possibleoptions = []string{"Rock", "Paper", "Scissors"}

func main() {
	username()
	rockpaperscissors()
}

func username() {
	// Seed to randomize the computer choice
	rand.Seed(time.Now().UnixNano())

	// Instructs the user to enter their name
	fmt.Println("Enter Your Name")

	// Welcome Text
	welcometext := "\n welcome to Rock, Paper, Scissors"

	// Reads the User's Name from Console
	nameinput := bufio.NewReader(os.Stdin)

	// Sets the name variable based on user input
	name, _ := nameinput.ReadString('\n')

	// Appends the name to the welcome text and prints to console
	name += welcometext
	fmt.Println(name)
}

func rockpaperscissors() {
	// User Choice
	fmt.Println("Please choose Rock, Paper, or Scissors")
	var choiceinput = bufio.NewReader(os.Stdin)
	userchoice, _ := choiceinput.ReadString('\n')

	// Computer choice
	computerchoice := possibleoptions[rand.Intn(len(possibleoptions))]

	// Prints the Choices to Screen
	fmt.Println("Your choice is: " + userchoice)
	fmt.Println("The computer choice is: " + computerchoice)

	// Check for Tie
	if strings.TrimRight(userchoice, "\n") == computerchoice {
		fmt.Println("It's a tie!")
	}

	// Calculate Winner
	if strings.TrimRight(userchoice, "\n") == "Rock" {
		if computerchoice == "Paper" {
			fmt.Println("Paper covers rock...computer wins!")
		}
		if computerchoice == "Scissors" {
			fmt.Println("Rock crushes scissors...user wins!")
		}
	}

	if strings.TrimRight(userchoice, "\n") == "Paper" {
		if computerchoice == "Scissors" {
			fmt.Println("Scissors cuts paper...computer wins!")
		}
		if computerchoice == "Rock" {
			fmt.Println("Paper covers rock...user wins!")
		}
	}

	if strings.TrimRight(userchoice, "\n") == "Scissors" {
		if computerchoice == "Rock" {
			fmt.Println("Rock crushes scissors...computer wins!")
		}
		if computerchoice == "Paper" {
			fmt.Println("Scissors cuts paper...user wins!")
		}
	}
}
