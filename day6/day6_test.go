package day5

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func Test_part1(t *testing.T) {
	fileName := "day6.txt"

	map1, x, y, dir := readMap(t, fileName)

	mapHeight := len(map1)
	mapWidth := len(map1[0])

	visitingMap, _ := testPath(mapHeight, mapWidth, y, x, dir, map1)

	totalPositionsCount := 0
	for rowIndex := 0; rowIndex < mapHeight; rowIndex++ {
		for colIndex := 0; colIndex < mapWidth; colIndex++ {
			if visitingMap[rowIndex][colIndex] != 0 {
				totalPositionsCount++
			}
		}
	}

	if totalPositionsCount != 5162 {
		t.Fatal("expected 5162 positions but got ", totalPositionsCount)
	}

	fmt.Println(totalPositionsCount)

}

func Test_part2(t *testing.T) {
	fileName := "day6.txt"

	map1, x, y, dir := readMap(t, fileName)

	mapHeight := len(map1)
	mapWidth := len(map1[0])

	visitingMap, _ := testPath(mapHeight, mapWidth, y, x, dir, map1)

	infinityLoopsFound := 0

	count := 0
	for rowIndex, row := range visitingMap {
		for colIndex, countOfVisits := range row {
			if countOfVisits > 0 {
				count++
				fmt.Println("count", count, "infinityLoopsFound", infinityLoopsFound)
				map1[rowIndex][colIndex] = 1
				_, isInfiniteLoop := testPath(mapHeight, mapWidth, y, x, dir, map1)
				if isInfiniteLoop {
					infinityLoopsFound++
				}
				map1[rowIndex][colIndex] = 0
			}
		}
	}

	if infinityLoopsFound != 1909 {
		t.Error("too high")
	}

	fmt.Println(infinityLoopsFound)

}

func testPath(mapHeight int, mapWidth int, y int, x int, dir int, map1 [][]int) ([][]int, bool) {
	visitingMap := make([][]int, mapHeight)
	for rowIndex := 0; rowIndex < mapHeight; rowIndex++ {
		visitingMap[rowIndex] = make([]int, mapWidth)
		for colIndex := 0; colIndex < mapWidth; colIndex++ {
			visitingMap[rowIndex][colIndex] = 0
		}
	}

	for {

		visitingMap[y][x]++

		if visitingMap[y][x] > 4 {

			return nil, true
		}

		// try to move to top
		if dir == 0 {
			// already at top of map
			if y == 0 {
				break
			}
			elementOnFront := map1[y-1][x]
			if elementOnFront == 0 {
				y--
			} else if elementOnFront == 1 {
				dir++
			}
			continue
		}

		// try to move to right
		if dir == 1 {
			// already at right of the map
			if x == mapWidth-1 {
				break
			}
			elementOnFront := map1[y][x+1]
			if elementOnFront == 0 {
				x++
			} else if elementOnFront == 1 {
				dir++
			}
			continue
		}

		// try to move to bottom
		if dir == 2 {
			// already at bottom of map
			if y == mapHeight-1 {
				break
			}
			elementOnFront := map1[y+1][x]
			if elementOnFront == 0 {
				y++
			} else if elementOnFront == 1 {
				dir++
			}
			continue
		}

		// try to move to left
		if dir == 3 {
			// already at left of map
			if x == 0 {
				break
			}
			elementOnFront := map1[y][x-1]
			if elementOnFront == 0 {
				x--
			} else if elementOnFront == 1 {
				dir = 0
			}
		}
	}

	return visitingMap, false
}

func readMap(t *testing.T, fileName string) ([][]int, int, int, int) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to open %s", fileName)
	}

	lines := strings.Split(string(bytes), "\n")

	map1 := make([][]int, len(lines))

	x := 0
	y := 0
	dir := 0

	for lineIndex, line := range lines {
		map1[lineIndex] = make([]int, len(lines[lineIndex]))
		for columnIndex, byteChar := range line {
			char := string(byteChar)
			switch char {
			case ".":
				map1[lineIndex][columnIndex] = 0
				break
			case "#":
				map1[lineIndex][columnIndex] = 1
				break
			case "^":
				map1[lineIndex][columnIndex] = 0
				x = columnIndex
				y = lineIndex
				dir = 0
				break
			case ">":
				map1[lineIndex][columnIndex] = 0
				x = columnIndex
				y = lineIndex
				dir = 1
				break
			case "v":
				map1[lineIndex][columnIndex] = 0
				x = columnIndex
				y = lineIndex
				dir = 2
				break
			case "<":
				map1[lineIndex][columnIndex] = 0
				x = columnIndex
				y = lineIndex
				dir = 3
				break
			default:
				t.Fatal("Unknown character")
			}

		}
	}
	return map1, x, y, dir
}
