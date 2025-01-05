package day5

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func Test_part1(t *testing.T) {
	fileName := "day5.txt"

	bytes, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to open %s", fileName)
	}

	parts := strings.Split(string(bytes), "\n\n")
	rulesMap := getRules(parts[0])
	documents := getDocuments(parts[1])

	sumOfMidPages := 0

	for _, pages := range documents {

		isGood := isPagesInRightOrder(pages, rulesMap)

		if isGood {
			migPage := pages[len(pages)/2]
			sumOfMidPages += migPage
		}
	}

	if sumOfMidPages != 6041 {
		t.Error("Sum of MidPages should be 6041")
	}
}

func Test_part2(t *testing.T) {
	fileName := "day5.txt"

	bytes, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to open %s", fileName)
	}

	parts := strings.Split(string(bytes), "\n\n")
	rulesMap := getRules(parts[0])
	documents := getDocuments(parts[1])

	sumOfMidPages := 0

	for index, pages := range documents {

		isGood := isPagesInRightOrder(pages, rulesMap)

		if !isGood {

			fixOrder(pages, rulesMap)

			isGood := isPagesInRightOrder(pages, rulesMap)
			if !isGood {
				t.Error("bad sort at index", index)
			}

			migPage := pages[len(pages)/2]
			sumOfMidPages += migPage
		}
	}

	fmt.Println(sumOfMidPages)

	if sumOfMidPages != 4884 {
		t.Error("Sum of MidPages should be 4884")
	}
}

func fixOrder(pages []int, rulesMap map[int][]int) {

	sort.Slice(pages, func(i, j int) bool {
		rightPage := pages[i]
		leftPage := pages[j]
		rules, exists := rulesMap[rightPage]
		if !exists {
			return false
		}
		for _, mustBefore := range rules {
			if mustBefore == leftPage {
				return true
			}
		}
		return false
	})
}

func isPagesInRightOrder(pages []int, rulesMap map[int][]int) bool {
	for pageIndex, pageNumber := range pages {
		rules, exists := rulesMap[pageNumber]
		if !exists {
			continue
		}
		if pageIndex == 0 {
			continue
		}
		for _, mustBefore := range rules {

			pagesBefore := pages[:pageIndex]
			for _, pageBefore := range pagesBefore {

				if mustBefore == pageBefore {
					return false
				}
			}
		}
	}
	return true
}

func getDocuments(string string) [][]int {
	documents := strings.Split(string, "\n")
	digitalDocuments := make([][]int, 0)
	for _, document := range documents {
		pagesNumbers := make([]int, 0)
		for _, pages := range strings.Split(document, ",") {
			v, _ := strconv.Atoi(pages)
			pagesNumbers = append(pagesNumbers, v)
		}
		digitalDocuments = append(digitalDocuments, pagesNumbers)
	}
	return digitalDocuments
}

func getRules(ee string) map[int][]int {
	rulesMap := map[int][]int{}
	rules := strings.Replace(ee, "\n", ",", -1)
	for _, rule := range strings.Split(rules, ",") {
		ruleParts := strings.Split(rule, "|")
		left, _ := strconv.Atoi(ruleParts[0])
		right, _ := strconv.Atoi(ruleParts[1])
		rulesMap[left] = append(rulesMap[left], right)
	}
	return rulesMap
}
