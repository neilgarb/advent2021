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
	f, err := os.Open(filepath.Join(wd, "five_two/input"))
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var tot int
	plane := make(map[point]int)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		p1, p2 := parseLine(line)
		p := p1
		for {
			plane[p]++
			if plane[p] == 2 {
				tot++
			}
			if p == p2 {
				break
			}
			if p2.x>p1.x {
				p.x++
			} else if p2.x < p1.x {
				p.x--
			}
			if p2.y>p1.y {
				p.y++
			} else if p2.y < p1.y {
				p.y--
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