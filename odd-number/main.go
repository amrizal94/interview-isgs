package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input numbers (comma separated)")

	input, _ := reader.ReadString('\n')

	numbers := strings.Split(input, ",")

	odds := make([]int, 0)

	for _, v := range numbers {
		num, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			fmt.Println("error parsing number", v)
			continue
		}

		if num%2 != 0 {
			odds = append(odds, num)
		}
	}

	sort.Ints(odds)

	file, err := os.Create("odds.txt")
	if err != nil {
		fmt.Println("error creating file", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprint(writer, "Odd Number: ")

	for i, v := range odds {
		fmt.Fprint(writer, v)
		if i != len(odds)-1 {
			fmt.Fprint(writer, ",")
		}
	}
	fmt.Fprintln(writer)
	writer.Flush()

	fmt.Println("Odd numbers writter to file")

}
