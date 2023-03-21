package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var start, end int
	fmt.Print("Input start number <1-100>: ")
	fmt.Scan(&start)
	fmt.Print("Input end number <1-100>: ")
	fmt.Scan(&end)

	num1 := rand.Intn(end-start+1) + start
	num2 := rand.Intn(end-start+1) + start
	num3 := rand.Intn(end-start+1) + start
	num4 := rand.Intn(end-start+1) + start

	fmt.Println("Random Number 1\tRandom Number 2\tResult")
	fmt.Printf("%d\t\t\t%d\t\t\t%s\n", num1, num2, checkNumTie(num1, num2))
	fmt.Printf("%d\t\t\t%d\t\t\t%s\n", num3, num4, checkNumTie(num3, num4))

}

func checkNumTie(num1, num2 int) string {
	if num1 == num2 {
		return "Tie"
	}
	return "No Tie"
}
