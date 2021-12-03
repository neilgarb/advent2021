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
	f, err := os.Open(filepath.Join(wd, "three_two/input"))
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var input []string
	var length int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		input = append(input, line)
		length = len(line)
	}

	ox := get(input,length, 0, '1')
	scrub := get(input, length, 0, '0')

	fmt.Println(ox, scrub, ox*scrub)

	return scanner.Err()
}

func get(input []string, length, pos int, c rune) int {
	if len(input) == 1 {
		return strtoint(input[0])
	}

	var ones int
	for _, s := range input {
		if s[pos] == '1' {
			ones++
		}
	}
	zeroes := len(input) - ones

	var newinput []string
	var want rune
	if c == '1' {
		if ones >= zeroes {
			want = '1'
		} else {
			want = '0'
		}
	} else {
		if ones < zeroes {
			want = '1'
		} else {
			want = '0'
		}
	}
	for _, s := range input {
		if rune(s[pos]) == want {
			newinput = append(newinput, s)
		}
	}

	return get(newinput, length, pos+1, c)
}

func strtoint(s string) int {
	i, _ := strconv.ParseInt(s, 2, 64)
	return int(i)
}
