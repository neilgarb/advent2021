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
	f, err := os.Open(filepath.Join(wd, "one_one/input"))
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var prev int
	var cnt int
	for scanner.Scan() {
		l := scanner.Text()
		i, err := strconv.Atoi(strings.TrimSpace(l))
		if err != nil {
			return err
		}
		if prev > 0 && i > prev {
			cnt++
		}
		prev = i
	}

	fmt.Println(cnt)

	return scanner.Err()
}
