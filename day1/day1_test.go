package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

type NumberPair struct {
	LeftNumber  int
	RightNumber int
}

func TestDay1_1(t *testing.T) {

	fileWithNumbers := "day1-numbers.txt"
	var numberPairs []NumberPair = readNumbers(t, fileWithNumbers)

	var totalDistance int = 0

	for _, pair := range numberPairs {
		if pair.LeftNumber > pair.RightNumber {
			totalDistance += pair.LeftNumber - pair.RightNumber
		} else if pair.RightNumber > pair.LeftNumber {
			totalDistance += pair.RightNumber - pair.LeftNumber
		}
	}

	fmt.Println(totalDistance)

}

func TestDay1_2(t *testing.T) {

	fileWithNumbers := "day1-numbers.txt"
	var numberPairs []NumberPair = readNumbers(t, fileWithNumbers)

	var numbersCounts map[int]int = make(map[int]int)
	for _, pair := range numberPairs {
		_, exists := numbersCounts[pair.RightNumber]
		if exists {
			numbersCounts[pair.RightNumber] = numbersCounts[pair.RightNumber] + 1
		} else {
			numbersCounts[pair.RightNumber] = 1
		}
	}

	var totalSimilarity int = 0

	for _, pair := range numberPairs {
		count, exists := numbersCounts[pair.LeftNumber]
		if exists {
			totalSimilarity += pair.LeftNumber * count
		}
	}

	fmt.Println(totalSimilarity)

}

func readNumbers(t *testing.T, fileWithNumbers string) []NumberPair {
	file, err := os.Open(fileWithNumbers)
	if err != nil {
		t.Fatalf("Error opening file: %v", fileWithNumbers)
	}
	defer file.Close()

	var numberPairs []NumberPair

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' '
		})

		num1, err1 := strconv.Atoi(parts[0])

		num2, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			t.Fatalf("Error converting %v to int: %v", line, err2)
		}

		numberPairs = append(numberPairs, NumberPair{LeftNumber: num1, RightNumber: num2})

	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("Scanner error: %v", err)
	}
	return numberPairs
}
