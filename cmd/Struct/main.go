package main

import (
	"fmt"
	"sort"
)

func main() {
	//students := []Student{
	//	{Name: "Иван", Surname: "Иванов", Age: 20, Class: "10-A"},
	//	{Name: "Мария", Surname: "Петрова", Age: 22, Class: "11-B"},s
	//	{Name: "Петр", Surname: "Сидоров", Age: 21, Class: "10-C"},
	//}
	//
	//PrintStudents(students)
	//student := createStudent()
	//fmt.Printf("Студент: %s, Возраст: %d, Курс: %s\n", student.Name, student.Age, student.Class)
	s := []int{46, 96, 8, 14, -18, -258}
	sort.Ints(s)
	fmt.Println(s)

	a := 28
	b := 123
	c := a + b
	fmt.Println(c)
	var new1 int
	for i := 0; i < b; i++ {
		if i == a {
			new1 = i
			break
		}
	}
	fmt.Println(new1)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	fmt.Println(getName(numbers))

	//	myMap := map[string]int{
	//		"один": 1,
	//		"два":  2,
	//		"три":  3,
	//	}
	//	value := myMap["три"]
	//	fmt.Println(value)
	//
}
func getName(numbers []int) map[int]string {
	myMap := make(map[int]string)
	for _, number := range numbers {
		if number <= 2 || number == 12 {
			myMap[number] = "Зима"
		}
		if number >= 3 && number < 6 {
			myMap[number] = "Весна"
		}
		if number >= 6 && number < 9 {
			myMap[number] = "Лето"
		}
		if number >= 9 && number < 12 {
			myMap[number] = "Осень"
		}
	}
	return myMap
}

// создать структуру Student, придумать свойтсва этой структуры с разными типами данных
// написать функцию, которая принимает слайс студентов, и печатает набор свойств в консоль

//написать функцию, которая ничего не принимает и возвращает студента(создать объект студента)
//создать функцию, используя предыдущую и вывести какие то параметры в консоль

//написать функцию которая принимает слайсл студентов и возвращает тех, кому больше 20 лет(сделать новый слайс из подходящих параметров)

//МАПЫ
//сделать мапу, где в качестве ключа строка, в качестве значения число
//написать функцию, которая заполнит мапу валютами, ключ -это валюту(например доллар, евро, фунт) а в качестве значения текущий  курс
//вывести в консоль длину мапы

type Student struct {
	Name    string
	Age     int
	Surname string
	Class   string
}

func PrintStudents(students []Student) {
	for _, student := range students {
		fmt.Printf("Имя: %s, Фамилия: %s, Возраст: %d, Класс: %s\n",
			student.Name, student.Surname, student.Age, student.Class)
	}
}
func createStudent() Student {
	return Student{
		Name:    "Георгий",
		Age:     34,
		Surname: "Абхазава",
		Class:   "6С",
	}
}

func studentAgeTwenty(students []Student) []Student {
	agetwenty := []Student{}
	for _, student := range students {
		if student.Age >= 20 {
			agetwenty = append(agetwenty, student)
		}
	}
	return agetwenty
}

//удаление элемента из слайса
//a := []string{"A", "B", "C"}
//a = a[:len(a)-1]

func returnTwo() (int, string) {
	return 28, "Голова"
}
