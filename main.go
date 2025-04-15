package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type LiveCells map[Coord]struct{}

type Coord [2]int64

var neighbors = []Coord{
	{-1, -1}, {0, -1},
	{1, -1}, {-1, 0},
	{1, 0}, {-1, 1},
	{0, 1}, {1, 1},
}

func main() {
	lc := parseInput()

	for range 10 {
		lc = lc.Next()
	}

	fmt.Println("#Life 1.06")
	for cell := range lc {
		fmt.Printf("%s\n", cell)
	}
}

func (lc LiveCells) Next() LiveCells {
	neighborCount := make(map[Coord]int)

	var neighbor Coord
	for cell := range lc {
		for _, delta := range neighbors {
			neighbor[0] = cell[0] + delta[0]
			neighbor[1] = cell[1] + delta[1]
			neighborCount[neighbor]++
		}
	}

	lc2 := make(LiveCells)
	for cell, count := range neighborCount {
		_, alive := lc[cell]
		if count == 3 || (count == 2 && alive) {
			lc2[cell] = struct{}{}
		}
	}

	return lc2
}

func (c Coord) String() string {
	return fmt.Sprintf("%d %d", c[0], c[1])
}

func parseInput() LiveCells {
	lc := make(LiveCells)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Fatalf("invalid input: %s", line)
		}
		x, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatalf("invalid x coordinate: %s", parts[0])
		}
		y, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatalf("invalid y coordinate: %s", parts[1])
		}
		lc[Coord{x, y}] = struct{}{}
	}

	return lc
}
