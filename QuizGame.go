package main

import ( 
	"fmt"
	"flag"
	"os"
	"encoding/csv"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer")
	flag.Parse()
	
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	q := len(lines)
	problems := parseLines(lines)
	right := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, p.q)
		fmt.Printf("A) %s\n", p.a1)
		fmt.Printf("B) %s\n", p.a2)
		fmt.Printf("C) %s\n", p.a3)
		fmt.Printf("D) %s\n", p.a4)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if strings.ToUpper(answer) == p.f {
			fmt.Println("Correct!\n\n")
			right ++
		} else {
			fmt.Printf("Incorrect, the correct answer is %s\n\n", p.f)
			}
	}
	result(right, q)
}

func result(x, q int) {
	y := float64(x)
	z := float64(q)
	var score float64
	score = (y / z) * 100
	fmt.Printf("Test completed, your score is : %.0f%%\n", score)
	
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a1: line[1],
			a2: line[2],
			a3: line[3],
			a4: line[4],
			f: line[5],
		}
	}
	return ret
}

type problem struct {
	q string
	a1 string
	a2 string
	a3 string
	a4 string
	f string
}

func exit(msg string) {
		fmt.Println(msg)
		os.Exit(1)
	}




