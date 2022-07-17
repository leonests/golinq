[![Build](https://github.com/leonests/golinq/workflows/CI/badge.svg)](https://github.com/leonests/golinq/actions?query=workflow)
[![Coverage](https://codecov.io/gh/leonests/golinq/branch/main/graphs/badge.svg?branch=main)](https://codecov.io/gh/leonests/golinq)
[![Go Report](https://goreportcard.com/badge/github.com/leonests/golinq)](https://goreportcard.com/report/github.com/leonests/golinq)
[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
------
# Golinq

Golinq provides support for language integrated query (LINQ) in Go, like C# LINQ.

* Generic LINQ, at least Go 1.18
* Support slice, array, map, channel, string and custom collections
* Easy to use, flexible to control 
* Also support non-generic version 

## Installation
When used with Go modules, use the following import path:

    go get github.com/leonests/golinq

## Quickstart

**Example 1: Find who has the most expensive book**

```go
import . "github.com/leonests/golinq"

type Person struct {
	Name         string
	Age          int
	Hobbies      []string
	LuckyNumbers []int
	BookPrices   map[string]float64
}
...

var persons []Person

res := FromSlice(persons).OrderBy(func(i int, p Person) any {
		return FromMap(p.BookPrices).OrderBy(func(s string, f float64) any {
			return f
		}).Last()
	}).ThenBy(func(i int, p Person) any { // if there are duplicates, ranking by lucky number
		return FromSlice(p.LuckyNumbers).OrderByDescending(func(i, n int) any {
			return n
		}).First()
	}).Last().Name
```

**Example 2: Find who has a hobby of basketball**

```go
import . "github.com/leonests/golinq"

type Person struct {
	Name         string
	Age          int
	Hobbies      []string
	LuckyNumbers []int
	BookPrices   map[string]float64
}
...

var persons []Person

res := FromSlice(persons).Where(func(i int, p Person) bool {
		return FromSlice(p.Hobbies).Contains(func(i int, hobby string) bool {
			return hobby == "basketball"
		})
	}).Select(func(i int, p Person) any { 
		return p.Name 
	}).ToSlice()
```