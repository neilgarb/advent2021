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
	f, err := os.Open(filepath.Join(wd, "four_two/input"))
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var i int
	var nos []string
	var curboard board
	var boards []board
	for scanner.Scan() {
		i++
		line := strings.TrimSpace(scanner.Text())
		if i == 1 {
			nos = strings.Split(line, ",")
			continue
		}
		if line == "" {
			if len(curboard.nos) > 0 {
				curboard.marked = make([]bool, len(curboard.nos))
				boards = append(boards, curboard)
				curboard = board{}
			}
		}
		curboard.nos = append(curboard.nos, strings.Fields(line)...)
	}
	curboard.marked = make([]bool, len(curboard.nos))
	boards = append(boards, curboard)


	complete := make(map[int]bool)

	var win int

	for _, i := range nos {
		for j := range boards {
			for k := range boards[j].nos {
				if boards[j].nos[k] == i {
					boards[j].marked[k] = true
					if boards[j].done() {
						if !complete[j] {
							complete[j] = true
							win = boards[j].calc(i)
						}
					}
				}
			}
		}
	}

	 fmt.Println(win)


	return scanner.Err()
}

type board struct {
	nos []string
	marked []bool
}

func (b board) done() bool {
	for i := 0; i < 5; i ++ {
		done := true
		for j := 0; j < 5; j ++ {
			if !b.marked[i*5+j] {
				done = false
				break
			}
		}
		if done {
			return true
		}
	}

	for i := 0; i < 5; i ++ {
		done := true
		for j := 0; j < 5; j ++ {
			if !b.marked[i+5*j] {
				done = false
				break
			}
		}
		if done {
			return true
		}
	}
	return false
}

func (b board) calc(lastNoStr string) int {
	lastNo, _ := strconv.Atoi(lastNoStr)
	var tot int
	for i := range b.nos {
		if b.marked[i] {
			continue
		}
		ii, _ := strconv.Atoi(b.nos[i])
		tot += ii
	}

	return tot * lastNo
}