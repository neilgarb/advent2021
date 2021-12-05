package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if err := do(); err != nil {
		panic(err)
	}
}

func do() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	f, err := os.Open(filepath.Join(wd, "five_one/input"))
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	plane := make(map[point]int)
	var tot int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		point1, point2 := parseLine(line)
		if point1.x == point2.x {
			if point2.y < point1.y {
				point1, point2 = point2, point1
			}
			for y := point1.y; y <= point2.y; y ++ {
				plane[point{point1.x, y}]++
				if plane[point{point1.x, y}] == 2 {
					tot ++
				}
			}
		} else if point1.y == point2.y {
			if point2.x < point1.x {
				point1, point2 = point2, point1
			}
			for x := point1.x; x <= point2.x; x ++ {
				plane[point{x, point1.y}]++
				if plane[point{x, point1.y}] == 2 {
					tot ++
				}
			}
		}
	}

	fmt.Println(tot)

	return scanner.Err()
}

type point struct {
	x, y int
}

func parseLine(line string) (point, point) {
	parts := strings.SplitN(line, " -> ", 2)
	if len(parts) != 2 {
		panic("not enough points")
	}
	return parsePoint(parts[0]), parsePoint(parts[1])
}

func parsePoint(p string) point {
	parts := strings.SplitN(p, ",", 2)
	if len(parts) != 2 {
		panic("not enough coords")
	}
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return point{x, y}
}