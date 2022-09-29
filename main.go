package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1
	playerScore := 0
	computerScore := 0

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	//instructions:
	fmt.Println("Rock, Paper & Scissors")
	fmt.Println("__________________________________________________________________________________")
	fmt.Println("Game is played for thee rounds, and the best of three Wins a Car. Good Luck!")
	fmt.Println("__________________________________________________________________________________")
	//loop for the attepms
	for i := 1; i <= 3; i++ {
		//print the atemp
		fmt.Println()
		fmt.Println("Round: ", i)
		fmt.Println("__________________________________________________________________________________")
		fmt.Println()
		fmt.Println("Please enter rock, paper or scissors ->")

		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\r", "", -1)
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)

		//add random to ensure the for loop resets everytime
		computerValue := rand.Intn(3)

		if playerChoice == "rock" {
			playerValue = ROCK
		} else if playerChoice == "paper" {
			playerValue = PAPER
		} else if playerChoice == "scissors" {
			playerValue = SCISSORS
		} else {
			//resets invalid option to keep playing
			playerValue = -1
		}
		//print out player choise
		fmt.Println()
		//get mayuscular character
		fmt.Println("Player chose", strings.ToUpper(playerChoice))

		switch computerValue {
		case ROCK:
			fmt.Println("Computer chose ROCK")

		case PAPER:
			fmt.Println("Computer chose PAPER")

		case SCISSORS:
			fmt.Println("Computer chose SCISSORS")

		default:
		}

		fmt.Println("-------------")

		if playerValue == computerValue {
			fmt.Println("it's a Draw")
			i--
		} else {
			switch playerValue {
			case ROCK:
				if computerValue == PAPER {
					computerScore = computerWins(computerScore)
					//fmt.Println("Computer Wins")
				} else {
					playerScore = playerWins(computerScore)
					//fmt.Println("Player Wins")
				}
				break
			case PAPER:
				if computerValue == SCISSORS {
					computerScore = computerWins(computerScore)
					//fmt.Println("Computer Wins")
				} else {
					playerScore = playerWins(playerScore)
					//fmt.Println("Player Wins")
				}

			case SCISSORS:
				if computerValue == ROCK {
					computerScore = computerWins(computerScore)
					//fmt.Println("Computer Wins")
				} else {
					playerScore = playerWins(playerScore)
					//fmt.Println("Player Wins")
				}
			default:
				fmt.Println("Invalid Choise")
				i--
			}
		}
	}
	fmt.Println("Final Scrore")
	fmt.Println("----------------------")
	fmt.Printf("Player: %d /3, Computer %d /3", playerScore, computerScore)
	fmt.Println()
	if playerScore > computerScore {
		fmt.Println("Paleyer Wins game!")
	} else {
		fmt.Println("Computer Wins game!")
	}
}

func computerWins(score int) int {
	fmt.Println("Computer Wins!")
	return score + 1
}
func playerWins(score int) int {
	fmt.Println("Player Wins!")
	return score + 1
}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
