package main

import (
	"fmt"
	"strconv"
	"math"
	"time"
)

func main() {
	fmt.Println("Variables Episode")
	fmt.Println("----------------------------")

	// simple variable
	var num = 2 // implicite type
	var num2 int = 3 // explicite type

	// initiate var
	var num3 int // by default the value num3 is 0
	// instantiate var
	num3 = 4 // if you want changing the value of num3 with new value
	
	// convert int to string with strconv.Itoa.
	// if you want conver string to int using strconv.Atoi.
	fmt.Println("Result is " + strconv.Itoa(num + num2 + num3))

	// multiple instantiate vars with singleline
	var varA, varB string
	// varA = "Irfan"
	// varB = "Irawan Sukirman"
	// or
	varA, varB = "Irfan ", "Irawan Sukirman"
	fmt.Println(varA + varB)

	// creation and assign with same line
	result := 10 // same with var result = 9, short hand syntax
	fmt.Println(result)

	fmt.Println("Const Episode")
	fmt.Println("----------------------------")

	// simple const
	const BASE_URL = "https://www.google.com"
	// BASE_URL = "" // error because const can't re-assign
	fmt.Println(BASE_URL)

	fmt.Println("Loops Episode")
	fmt.Println("----------------------------")

	// only have one loop type is for
	for index := 0; index < 10; index++ {
		fmt.Println("Loop ke " + strconv.Itoa(index))
	}

	// index := 0
	// unlimited loop
	// for { index = index + 1 
		// fmt.Println("Loop ke " + strconv.Itoa(index))
	// }

	fmt.Println("Functions Episode")
	fmt.Println("----------------------------")

	x, y := 5, 5
	results := add(x, y)
	resultss := multiplication(x, y)
	result1, result2 := multiProcess(x, y)
	result3, result4 := multiProcess2(x, y)
	fmt.Println(results)
	fmt.Println(resultss)
	fmt.Println(result1, result2)
	fmt.Println(result3, result4)

	fmt.Println("Variable Scope Function Package Level Episode")
	fmt.Println("----------------------------")

	funcLevelVar()

	fmt.Println("Exported Names Episode")
	fmt.Println("----------------------------")
	// if local packages name function using pascalCase. let say getBookList().
	// if you want call a function from other package let say B file the function name should use CamelCase. Exp: GetBookList()

	fmt.Println("Golang Math Episode")
	fmt.Println("----------------------------")
	var angka float64 = 12
	hasil := math.Sqrt(angka)
	pembulatan := math.Round((hasil))
	fmt.Printf("Hasilnya adalah %g thanks", pembulatan)

	fmt.Println("If Else Switch Episode")
	fmt.Println("----------------------------")

	state := 5
	if state > 3 {
		fmt.Println("Lebih besar")
	} else {
		fmt.Println("Lebih Kecil")
	}

	// switch

	switch state {
	case 3 : fmt.Println("Lebih kecil")
	default: fmt.Println("Lebih Besar")
	}

	fmt.Println("Time Episode")
	fmt.Println("----------------------------")

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0: fmt.Println("Today.")
	case today + 1: fmt.Println("Tomorrow.")
	case today + 2: fmt.Println("In two days")
	default: fmt.Println("Too far away")
	}
}

var aaa = 10 // this is package level var (global var)

func funcLevelVar() {
	var aaa = 10 // this is function level var (local var)
	fmt.Println(aaa)
}

func add(x int, y int) int  {
	out := x + y
	return out
}

func multiplication(x, y int) int  {
	out := x * y
	return out
}

func multiProcess(x, y int) (int, int)  {
	out1 := x + y
	out2 := x * y
	return out1, out2
}

// or with other style

func multiProcess2(x, y int) (out1, out2 int) {
	out1 = x + y
	out2 = x * y
	return
}