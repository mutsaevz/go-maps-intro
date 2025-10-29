package main

import (
	"fmt"
	"strings"
)

func getVisits(m map[string]int, day string) (int, bool) {
	v, ok := m[day]

	return v, ok
}

func main() {
	m := make(map[string]int)

	m = map[string]int{
		"Mon": 15,
		"Tue": 22,
		"Wen": 20,
		"The": 18,
	}

	delete(m, "Mon")
	v, ok := m["Mo"]
	fmt.Println(m, len(m), v, ok)

	fmt.Println(getVisits(map[string]int{"Mon": 15, "Tue": 22, "Wen": 20, "The": 18}, "The"))

	m1 := map[string]string{
		"Грозный":     "Интукод",
		"Урус-мартан": "Автомобильный рынок",
		"Закан-юрт":   "Природа",
		"Рим":         "Колизей",
		"Париж":       "Эйфелева Башня",
		"Мекка":       "Кааба",
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("-------")
		for k, v := range m1 {
			fmt.Println(k, v)

		}
	}
	var m2 = make(map[string]int)
	m2["visits"] = 1

	// fmt.Println(m2)

	type Student struct {
		Name   string
		Active bool
	}
	m3 := map[string]Student{
		"st1": {Name: "Zubayr", Active: true},
		"st2": {Name: "Adam", Active: false},
		"st3": {Name: "Rasul", Active: true},
	}
	fmt.Println(m3["st3"])

	m4 := m3["st3"]

	m3["st3"] = Student{Name: m4.Name, Active: false}

	fmt.Println(m3["st3"])

	a := map[string]int{"Mon": 10}

	b := a
	b["Mon"] = 99
	fmt.Println(a, b)
	fmt.Println("trger", clone(map[string]int{"Mon": 10,"Tue": 15,}))


	fmt.Println(WordCount("go go gopher go"))

	phones := map[string]string{}

	Add(phones, "qwcqerc", "345252334")
	Add(phones, "qwwcewceqcevd", "476486")
	Add(phones, "jytmmtym", "456262")
	for i, v := range phones {
		fmt.Println(i, v)
	}
	Get(phones, "qwcqerc")
	Remove(phones, "jytmmtym")

	fmt.Println(len(phones))
}

func Add(phones map[string]string, name, number string) {
	phones[name] = number
}
func Get(phones map[string]string, name string) (string, bool) {
	v, ok := phones[name]

	return v, ok
}
func Remove(phones map[string]string, name string) {
	delete(phones, name)
}

func clone(src map[string]int) map[string]int {
	a := make(map[string]int)

	a = src

	for i, v := range src{
		a[i] = v
	}
	return a
}
func WordCount(s string) map[string]int {
	s1 := strings.Fields(s)
	// var count int
	m := make(map[string]int)

	for _, v := range s1 {
		// count++
		m[v]++
	}

	return m
}
