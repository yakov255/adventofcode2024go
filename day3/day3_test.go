package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	fileName := "day3.txt"
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to open %s", fileName)
	}
	regex, _ := regexp.Compile(`(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`)
	res := regex.FindAllSubmatch(bytes, -1)

	result := 0

	enabled := true

	for _, row := range res {

		if row[5] != nil {
			// don't

			enabled = false

		} else if row[4] != nil {
			// do
			enabled = true

		} else {
			if enabled {
				a, _ := strconv.Atoi(string(row[2]))
				b, _ := strconv.Atoi(string(row[3]))
				result += a * b
			}
		}

	}

	fmt.Println(result)

}
