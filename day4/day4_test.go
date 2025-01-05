package day4

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func Test_part1(t *testing.T) {
	fileName := "day4.txt"

	bytes, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to open %s", fileName)
	}
	lines := strings.Split(string(bytes), "\n")

	totalWorldsCount := 0

	// rows
	for _, line := range lines {
		for i := 0; i < len(line)-3; i++ {
			substring := line[i : i+4]
			if substring == "XMAS" || substring == "SAMX" {
				totalWorldsCount++
			}
		}
	}
	// cols
	for rowIndex := 0; rowIndex < len(lines)-3; rowIndex++ {
		line := lines[rowIndex]
		for colIndex := 0; colIndex < len(line); colIndex++ {
			var substring = string(lines[rowIndex][colIndex]) + string(lines[rowIndex+1][colIndex]) + string(lines[rowIndex+2][colIndex]) + string(lines[rowIndex+3][colIndex])
			if substring == "XMAS" || substring == "SAMX" {
				totalWorldsCount++
			}
		}
	}
	// diagonal from top left
	for rowIndex := 0; rowIndex < len(lines)-3; rowIndex++ {
		line := lines[rowIndex]
		for colIndex := 0; colIndex < len(line)-3; colIndex++ {
			var substring = string(lines[rowIndex][colIndex]) + string(lines[rowIndex+1][colIndex+1]) + string(lines[rowIndex+2][colIndex+2]) + string(lines[rowIndex+3][colIndex+3])
			if substring == "XMAS" || substring == "SAMX" {
				totalWorldsCount++
			}
		}
	}
	// diagonal from top right
	for rowIndex := 0; rowIndex < len(lines)-3; rowIndex++ {
		line := lines[rowIndex]
		for colIndex := 3; colIndex < len(line); colIndex++ {
			var substring = string(lines[rowIndex][colIndex]) + string(lines[rowIndex+1][colIndex-1]) + string(lines[rowIndex+2][colIndex-2]) + string(lines[rowIndex+3][colIndex-3])
			if substring == "XMAS" || substring == "SAMX" {
				totalWorldsCount++
			}
		}
	}

	if totalWorldsCount != 2514 {
		t.Error("bad response")
	}

	fmt.Println(totalWorldsCount)

}

func Test_part(t *testing.T) {
	fileName := "day4.txt"

	bytes, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to open %s", fileName)
	}
	lines := strings.Split(string(bytes), "\n")

	totalXCount := 0

	// cols
	for rowIndex := 0; rowIndex < len(lines)-2; rowIndex++ {
		line := lines[rowIndex]
		for colIndex := 0; colIndex < len(line)-2; colIndex++ {

			//0 1
			// 3
			//2 4
			var parts = [5]string{
				string(lines[rowIndex][colIndex]),
				string(lines[rowIndex][colIndex+2]),
				string(lines[rowIndex+2][colIndex]),
				string(lines[rowIndex+1][colIndex+1]),
				string(lines[rowIndex+2][colIndex+2]),
			}

			variant1 := parts[0] + parts[3] + parts[4]
			variant2 := parts[1] + parts[3] + parts[2]

			if (variant1 == "MAS" || variant1 == "SAM") && (variant2 == "MAS" || variant2 == "SAM") {
				totalXCount++
			}

		}
	}

	fmt.Println(totalXCount)

	if totalXCount != 1888 {
		t.Error("bad count")
	}

}
