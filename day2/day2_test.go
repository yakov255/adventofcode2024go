package day2

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestDay2_part1(t *testing.T) {

	fileName := "day2-numbers.txt"

	reportResults := readFile(t, fileName)

	safeReports := 0

	fmt.Printf("Found %v reports\n", len(reportResults))

	for _, numbers := range reportResults {

		fmt.Printf("checking row %v\n", numbers)

		isSafeReport := isSafe(numbers)
		if isSafeReport {
			fmt.Print("whole report is safe\n")
			safeReports++
		} else {
			fmt.Print("whole report is unsafe\n")
		}

	}

	fmt.Printf("Safe reports found: %d\n", safeReports)
}

func isSafe(numbers []int) bool {
	isSafeReport := true
	var prevValue *int
	var isUp *bool
	for _, number := range numbers {
		if prevValue == nil {
			prevValue = &number
			continue
		}
		if *prevValue < number {
			if isUp == nil {
				val := true
				isUp = &val
			}
			if *isUp && number-*prevValue <= 3 {
				*prevValue = number
			} else {
				isSafeReport = false
				break
			}
		} else if *prevValue > number {
			if isUp == nil {
				val := false
				isUp = &val
			}
			if !*isUp && *prevValue-number <= 3 {
				*prevValue = number
			} else {
				isSafeReport = false
				break
			}
		} else {
			isSafeReport = false
			break
		}
	}
	return isSafeReport
}

func TestDay2_part2(t *testing.T) {

	fileName := "day2-numbers.txt"
	reportResults := readFile(t, fileName)

	findSafeReports(reportResults)
}

func TestDay2_part2_test(t *testing.T) {

	fileName := "day2-numbers-test.txt"
	reportResults := readFile(t, fileName)

	safeReports := findSafeReports(reportResults)

	excepted := 9

	if safeReports != excepted {
		t.Errorf("Expected %d, got %d", excepted, safeReports)
	}

}

func findSafeReports(reportResults [][]int) int {
	safeReports := 0
	fmt.Printf("Found %v reports\n", len(reportResults))

	for _, numbers := range reportResults {

		fmt.Printf("checking row %v", numbers)
		if isSafe(numbers) {
			fmt.Print(" ok\n")
			safeReports++
		} else {
			fmt.Print(" fail\n")
			if isRepostSafeWithRemoving(numbers) {
				safeReports++
			}
		}
	}
	fmt.Printf("Safe reports found: %d\n", safeReports)
	return safeReports
}

func isRepostSafeWithRemoving(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		v1 := numbers[:i]
		v2 := numbers[i+1:]
		newNumbers := append(slices.Clone(v1), slices.Clone(v2)...)
		fmt.Printf("  rechecking row %v", newNumbers)

		isSafeReport := isSafe(newNumbers)
		if isSafeReport {
			fmt.Print(" ok\n")
			return true
		} else {
			fmt.Print(" fail\n")
		}
	}
	return false
}

func readFile(t testing.TB, fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("Failed to open %s", fileName)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := make([][]int, 0)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Split(line, " ")

		values := make([]int, len(parts))

		for i, v := range parts {
			values[i], err = strconv.Atoi(v)
			if err != nil {
				t.Fatalf("Failed to convert string %s to int", v)
			}
		}

		numbers = append(numbers, values)
	}
	return numbers
}
