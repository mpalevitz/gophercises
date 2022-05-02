package main

// Import necessary pacakges.
import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"flag"
	"time"
)


var scoreCounter int = 0 // scoreCounter holds the player score.
var gameLength int = 0 // gameLength holds the amount of questions asked.

// Function gameRules displays the instructions at the start of the program and blocks 
// the rest of the program from running until user input. 
func gameRules(start chan bool, timer *int) {
	fmt.Printf("Questions will display one at a time. Type your answer and press 'enter' to register it. You will have %v seconds to finish, press 'enter' to begin.\n", *timer)
	fmt.Scanln()
	start <- true

}

// Function gameTime sleeps for a determined amount of time then prints the current
// game score and exits the program.
func gameTime(secondsPointer *int) {
	time.Sleep(time.Duration(*secondsPointer) * time.Second)
	fmt.Println("Your time is up!")
	fmt.Printf("Score: %v/%v correct.\n", scoreCounter, gameLength)
	os.Exit(0)
}

// Function 'main' prints questions from CSV column 1 and matches the user input
// to column 2, then prints out how many answers are correct.
func gamePlay(filePointer *string) {

	f, err := os.Open(*filePointer) // Open CSV and error handle.
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f) // csvReader is the read CSV file.



	// Loop through CSV lines.
	for {
		rec, err := csvReader.Read()
		// CSV line read error handling.
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		
		// Print column 1 question for that row.
		fmt.Printf("Question #%v: %v = \n", gameLength+1, rec[0]) 
		
		gameLength++

		// Get user input and store answer.
		var answer string
		fmt.Scanln(&answer)

		// Compare input and answer. +1 if correct.
		if answer == rec[1] {
			scoreCounter++
		}

	}
	// Print the final score.
	fmt.Printf("Score: %v/%v correct.\n", scoreCounter, gameLength)
}

func main() {
	// Define 'file' flag, default value of 'problems.csv', and flag description.
	// Define 'timer' flag, default value of 30 seconds, and flag description
	filePtr := flag.String("file", "problems.csv", "path to CSV file")
	timerPtr := flag.Int("timer", 30, "set game timer in seconds")
	flag.Parse() // Parse given flags.
	
	start := make (chan bool) // Define channel for gameRules
	
	go gameRules(start, timerPtr)
	<-start // Receive from gameRules channel to continue program.

	// Run timer concurrently with answering questions.
	go gameTime(timerPtr)
	gamePlay(filePtr)
}