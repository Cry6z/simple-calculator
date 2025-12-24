package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var operations = map[string]func(float64, float64) (float64, error){
	"tambah": func(a, b float64) (float64, error) { return a + b, nil },
	"kurang": func(a, b float64) (float64, error) { return a - b, nil },
	"kali": func(a, b float64) (float64, error) { return a * b, nil },
	"bagi": func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	},
}

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("usage: %s <tambah|kurang|kali|bagi> <number> <number>", os.Args[0])
	}

	opName := os.Args[1]
	left, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatalf("nomor pertama tidak valid: %v", err)
	}

	right, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatalf("nomor kedua tidak valid: %v", err)
	}

	op, ok := operations[opName]
	if !ok {
		log.Fatalf("tidak diketahui %q. gunakan tambah, kurang, kali, or bagi", opName)
	}

	result, err := op(left, right)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%g %s %g = %g\n", left, symbol(opName), right, result)
}

func symbol(op string) string {
	switch op {
	case "tambah":
		return "+"
	case "kurang":
		return "-"
	case "kali":
		return "×"
	case "bagi":
		return "÷"
	default:
		return op
	}
}
