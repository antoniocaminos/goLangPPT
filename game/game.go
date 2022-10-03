package game

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

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 0
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
			g.DisplayChan <- ""
		}
	}
}

// use select to process input in channels
// print to screen
// Keep  track of round number

// clearScreen clears the screen
func (g *Game) ClearScreen() {
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
func (g *Game) PrintIntro() {
	//instructions:
	g.DisplayChan <- fmt.Sprintf(`
Rock, Paper & Scissors
__________________________________________________________________________________
Game is played for thee rounds, and the best of three Wins a Car. Good Luck!
`)
	<-g.DisplayChan
}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1
	g.DisplayChan <- fmt.Sprintf(`
	Round %d
______________________
`, g.Round.RoundNumber)
	<-g.DisplayChan

	fmt.Print("Please enter rock, paper, scissors -> ")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\r", "", -1)
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}
	g.DisplayChan <- ""
	<-g.DisplayChan

	g.DisplayChan <- fmt.Sprintf("Player Chose %s", strings.ToUpper(playerChoice))
	<-g.DisplayChan

	switch computerValue {
	case ROCK:
		g.DisplayChan <- "Computer Chose ROCK"
		<-g.DisplayChan
		break
	case PAPER:
		g.DisplayChan <- "Computer Chose PAPAER"
		<-g.DisplayChan
		break
	case SCISSORS:
		g.DisplayChan <- "Computer Chose SCISSORS"
		<-g.DisplayChan
		break
	default:

	}
	//fmt.Println("-------------")

	if playerValue == computerValue {
		g.DisplayChan <- "it's a Draw!"
		<-g.DisplayChan
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		default:
			g.DisplayChan <- "invalid Choise"
			<-g.DisplayChan
			return false
		}
	}
	return true
}
func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer Wins!"
	<-g.DisplayChan
}
func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player Wins!"
	<-g.DisplayChan
}
func (g *Game) PrintSummary() {
	fmt.Println("Final Score:")
	fmt.Println("-----------------------------------------")
	fmt.Printf("Player %d/3, Computer %d/3", g.Round.PlayerScore, g.Round.ComputerScore)
	fmt.Println()
	if g.Round.PlayerScore > g.Round.ComputerScore {
		fmt.Println("human wins game!")
	} else {
		fmt.Println("Computer wins the Car!")
	}
}
