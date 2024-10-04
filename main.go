package main

import "fmt"

func main() {
	//numbers := []int{1, 2, 3}
	//for _, number := range numbers {
	//	fmt.Println(number)
	//}

	//fmt.Println(welcome("Гика"))
	a := "Гей машина"
	b := "Dddddddddd"
	fmt.Println(len(a), len(b))
	arunes := []rune(a)
	brunes := []rune(b)
	fmt.Println(len(arunes))
	fmt.Println(len(brunes))
	//fmt.Println(task(15))
}
func task(maxNumber int) int {
	var number = 1
	var result int
	for number <= maxNumber {
		result = result + number
		number++
	}
	return result
}

func minAge(ages []int) int {
	min := ages[0]
	for _, age := range ages {
		if age < min {
			min = age
		}
	}
	return min
}

type Square struct {
	SideA int
}

func area(s Square) int {
	return s.SideA * s.SideA
}

func welcome(name string) string {
	return fmt.Sprintf("Добро пожаловать, %s в клуб", name)
}
