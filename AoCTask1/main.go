package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers, err := readLines("test.txt")
	if err != nil {
		panic(err)
	}

	increased := 0
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] < numbers[i+1] {
			increased++
		}
	}
	fmt.Println("Task1: " + fmt.Sprint(increased))

	increased = 0
	for i := 0; i < len(numbers)-3; i++ {
		previousSum := numbers[i] + numbers[i+1] + numbers[i+2]
		currentSum := numbers[i+1] + numbers[i+2] + numbers[i+3]
		if currentSum > previousSum {
			increased++
		}
	}
	fmt.Println("Task2: " + fmt.Sprint(increased))
}

func readLines(fileName string) ([]int, error) {
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

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, num)
	}

	return lines, scanner.Err()
}
