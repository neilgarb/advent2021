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
	f, err := os.Open(filepath.Join(wd, "two_two/input"))
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var horizontal, depth, aim int
	for scanner.Scan() {
		horizontal, depth, aim = applyCommand(scanner.Text(), horizontal, depth, aim)
		fmt.Println(horizontal, depth, horizontal*depth)
	}


	return scanner.Err()
}

func applyCommand(text string, horizontal, depth, aim int) (int, int, int) {
	parts := strings.SplitN(text, " ", 2)
	if len(parts) < 2 {
		return horizontal, depth, aim
	}
	val, err := strconv.Atoi(parts[1])
	if err != nil {
		return horizontal, depth, aim
	}
	switch parts[0] {
	case "up":
		aim -= val
	case "down":
		aim += val
	case "forward":
		horizontal += val
		depth += aim*val
	}
	return horizontal, depth, aim
}
