package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Unit     int
	Price    int
	Currency string
	Total    int
}

func main() {
	// Read the currency conversion rates from currency.txt file
	currencyFile, err := os.Open("currency.txt")
	if err != nil {
		fmt.Println("Error opening currency file:", err)
		return
	}
	defer currencyFile.Close()

	conversionRates := make(map[string]int)
	scanner := bufio.NewScanner(currencyFile)
	// Read the first line to skip it
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(strings.TrimSpace(line), ",")
		currency := fields[0]
		rate, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("Error parsing currency rate:", err)
			return
		}
		conversionRates[currency] = rate
	}
	fmt.Println(conversionRates)

	// open the file for reading
	file, err := os.Open("pricing.txt")
	if err != nil {
		fmt.Println("error opening file", err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)

	var items []Item
	// Read the first line to skip it
	scanner.Scan()
	for scanner.Scan() {
		var currency string
		var unit, price int
		line := scanner.Text()
		fields := strings.Split(strings.TrimSpace(line), ",")
		if len(fields) >= 4 {
			currency = fields[3]
			unit, _ = strconv.Atoi(fields[1])
			price, _ = strconv.Atoi(fields[2])
		}

		total := unit * price * conversionRates[currency]

		// fmt.Println(total)
		item := Item{
			Unit:     unit,
			Price:    price,
			Currency: currency,
			Total:    total,
		}
		items = append(items, item)
	}
	fmt.Println(items)
	if _, err := os.Stat("pricing.txt"); os.IsExist(err) {
		// Attempt to delete the file
		err := os.Remove("pricing.txt")
		if err != nil {
			fmt.Println("Error deleting file:", err)
			return
		}
	}
	pricingFile, err := os.Create("pricing.txt")
	if err != nil {
		fmt.Println("Error creating pricing file:", err)
		return
	}
	defer pricingFile.Close()
	writer := bufio.NewWriter(pricingFile)
	fmt.Fprint(writer, "Item,Unit,Price,Currency,Total in IDR\n")
	fmt.Println(items)
	for i := 0; i < len(items); i++ {
		fmt.Fprintf(writer, "Item %d,%d,%d,%s,%d\n", i+1, items[i].Unit, items[i].Price, items[i].Currency, items[i].Total)
	}
	writer.Flush()
}
