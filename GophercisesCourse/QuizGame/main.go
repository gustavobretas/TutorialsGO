package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	shuffle := flag.Bool("s", false, "shuffle the quiz order each time it is run")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintln("Failed to open the CSV file: ", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines, *shuffle)
	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer != problem.answer {
			fmt.Println("You gave the wrong answer.")
		} else {
			correct++
			fmt.Println("Correct!")
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string, shuffle bool) []problem {
	result := make([]problem, len(lines))
	for i, line := range lines {
		result[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(lines), func(i, j int) { result[i], result[j] = result[j], result[i] })
	}
	return result
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
