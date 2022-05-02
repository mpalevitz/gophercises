package main

// Import necessary pacakges.
import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"flag"	
)

// Function 'main' prints questions from CSV column 1 and matches the user input
// to column 2, then prints out how many answers are correct.
func main() {

	// Define 'file' flag, default value of 'problems.csv', and flag description.
	filePtr := flag.String("file", "problems.csv", "path to CSV file")
	flag.Parse() // Parse given flags.

	f, err := os.Open(*filePtr) // Open CSV and error handle.
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f) // csvReader is the read CSV file.
	
	scoreCounter := 0 // scoreCounter stores the correct answer tally.

	gameLength := 0 // gameLength stores the total questions tally.

	fmt.Println("Type your answer and press 'enter' for each question.\nYour score will be displayed at the end.")

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

		// Get user input and store answer.
		var answer string
		fmt.Scanln(&answer)

		gameLength++
		// Compare input and answer. +1 if correct.
		if answer == rec[1] {
			scoreCounter++
		}

	}
	// Print the final score.
	fmt.Printf("Score: %v/%v correct.\n", scoreCounter, gameLength)
}
