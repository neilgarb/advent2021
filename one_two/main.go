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
	f, err := os.Open(filepath.Join(wd, "one_two/input"))
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var prev []int
	var sum int
	var cnt int
	for scanner.Scan() {
		l := scanner.Text()
		i, err := strconv.Atoi(strings.TrimSpace(l))
		if err != nil {
			return err
		}
		if len(prev) < 3 {
			prev = append(prev, i)
			sum += i
			continue
		}
		if sum-prev[0]+i > sum {
			cnt++
		}
		sum = sum-prev[0]+i
		prev = append(prev[1:], i)
	}

	fmt.Println(cnt)

	return scanner.Err()
}
