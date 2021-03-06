package golinq

import (
	"fmt"
)

type Person struct {
	Name         string
	Age          int
	Hobbies      []string
	LuckyNumbers []int
	BookPrices   map[string]float64
}

var persons = []Person{
	{
		Name:         "Jack",
		Age:          15,
		Hobbies:      []string{"basketball", "running"},
		LuckyNumbers: []int{31, 7, 91},
		BookPrices: map[string]float64{
			"linux": 101.2,
			"apple": 58.1,
		},
	},
	{
		Name:         "Rose",
		Age:          18,
		Hobbies:      []string{"movie", "reading"},
		LuckyNumbers: []int{1, 33, 9},
		BookPrices: map[string]float64{
			"algorithm": 101.2,
			"google":    56.4,
		},
	},
	{
		Name:         "Leon",
		Age:          20,
		Hobbies:      []string{"basketball", "reading", "coding"},
		LuckyNumbers: []int{23, 3, 7},
		BookPrices: map[string]float64{
			"computer":  24.8,
			"microsoft": 87.54,
		},
	},
}

func ExampleFromSlice() {
	// Who Is The Youngest
	// without considering persons with the same age
	res := FromSlice(persons).
		OrderBy(func(i int, p Person) any { return p.Age }).
		First().Name
	fmt.Println(res)
	// Output: Jack
}

func ExampleFromMap() {
	// Who Has The Most Expensive Book
	res := FromSlice(persons).
		OrderBy(func(i int, p Person) any {
			return FromMap(p.BookPrices).OrderBy(func(s string, f float64) any {
				return f
			}).Last()
		}).ThenBy(func(i int, p Person) any {
		return FromSlice(p.LuckyNumbers).OrderByDescending(func(i, n int) any {
			return n
		}).First()
	}).
		Last().Name
	fmt.Println(res)
	// Output: Jack
}

func ExampleSelect() {
	// Whose Hobbies Contain Basketball
	res := FromSlice(persons).
		Where(func(i int, p Person) bool {
			return FromSlice(p.Hobbies).Contains(func(i int, s string) bool {
				return s == "basketball"
			})
		}).
		Select(func(i int, p Person) any { return p.Name }).
		ToSlice()
	fmt.Println(res)
	// Output: [Jack Leon]
}

func ExampleJust() {
	// Who Has Biggest Lucky Number
	res := Just(persons...).
		OrderBy(func(i int, p Person) any {
			return FromSlice(p.LuckyNumbers).OrderBy(func(i1, i2 int) any { return i2 }).Last()
		}).
		Last().Name
	fmt.Println(res)
	// Output: Jack
}

func Example() {
	// Age Sequence
	res := FromSlice(persons).
		Select(func(i int, p Person) any { return p.Age }).
		ToSlice()
	fmt.Println(res)
	// Output: [15 18 20]
}
