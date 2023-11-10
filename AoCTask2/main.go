package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	lines, err := readLines("test.txt")
	if err != nil {
		panic(err)
	}

	runTaskOne(lines)
	runTaskTwo(lines)
}

func runTaskOne(lines []string) {

	var horizontal int
	var depth int

	for _, value := range lines {
		command := strings.Split(value, " ")
		num, err := strconv.Atoi(command[1])
		if err != nil {
			panic(err)
		}

		switch command[0] {
		case "forward":
			horizontal += num
		case "down":
			depth += num
		case "up":
			depth -= num
		default:
			panic("Command mismatch")
		}
	}

	fmt.Println("Task1: " + fmt.Sprint(horizontal*depth))
}

func runTaskTwo(lines []string) {
	var horizontal int
	var depth int
	var aim int

	for _, value := range lines {
		command := strings.Split(value, " ")
		num, err := strconv.Atoi(command[1])
		if err != nil {
			panic(err)
		}

		switch command[0] {
		case "forward":
			horizontal += num
			depth += aim * num
		case "down":
			aim += num
		case "up":
			aim -= num
		default:
			panic("Command mismatch")
		}
	}

	fmt.Println("Task2: " + fmt.Sprint(depth*horizontal))
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
