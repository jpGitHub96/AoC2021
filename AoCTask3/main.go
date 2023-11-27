package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}

	solveTask1(lines)
	solveTask2(lines)
}

func solveTask1(lines []string) {
	var mostCommonBits string
	var mostUncommonBits string
	for i := 0; i < len(lines[0]); i++ {
		var count int
		for _, s := range lines {
			if string(s[i]) == "1" {
				count++
			}
		}
		if count > len(lines)/2 {
			mostCommonBits += "1"
			mostUncommonBits += "0"
		} else {
			mostCommonBits += "0"
			mostUncommonBits += "1"
		}
	}

	fmt.Println("Task1: " + fmt.Sprint(getDecimal(mostCommonBits)*getDecimal(mostUncommonBits)))
}

func solveTask2(lines []string) {
	oxyRating := task2Recur(lines, 0, func(c int, length float64) bool {
		if float64(c) >= length {
			return true
		} else {
			return false
		}
	})

	co2Rating := task2Recur(lines, 0, func(c int, length float64) bool {
		if float64(c) < length {
			return true
		} else {
			return false
		}
	})

	fmt.Println("Task2: " + fmt.Sprint(getDecimal(oxyRating)*getDecimal(co2Rating)))
}

func task2Recur(rating []string, depth int, f func(int, float64) bool) string {
	if len(rating) == 1 {
		return rating[0]
	}

	var count int
	var leadingOne []string
	var leadingZero []string
	for _, s := range rating {
		if string(s[depth]) == "1" {
			count++
			leadingOne = append(leadingOne, s)
		} else {
			leadingZero = append(leadingZero, s)
		}
	}
	var length = float64(len(rating)) / 2.0
	if f(count, length) {
		return task2Recur(leadingOne, depth+1, f)
	} else {
		return task2Recur(leadingZero, depth+1, f)
	}
}

func getDecimal(bits string) int64 {
	decimal, err := strconv.ParseInt(bits, 2, 64)
	if err != nil {
		panic(err)
	}
	return decimal
}

func readLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
