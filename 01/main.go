package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partTwo() {

}

func partOne() {
	// open the file
	file, err := os.Open("./01/input.txt")

	if err != nil {
		fmt.Println("error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers1 []int
	var numbers2 []int

	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Fields(line)
		number1, _ := strconv.Atoi(words[0])
		number2, _ := strconv.Atoi(words[1])

		numbers1 = append(numbers1, number1)
		numbers2 = append(numbers2, number2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error scanning file", err)
	}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	fmt.Println("sorting")

	sort.Ints(numbers1)
	sort.Ints(numbers2)

	fmt.Println("sorted numbers1", numbers1)
	fmt.Println("sorted numbers2", numbers2)

	var sum float64
	sum = 0

	for idx, num1 := range numbers1 {
		num2 := numbers2[idx]

		diff := math.Abs(float64(num1 - num2))
		sum += diff
	}

	fmt.Printf("result %f", sum)
}

func main() {
	partOne()
}