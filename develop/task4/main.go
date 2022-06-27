package main

import (
	"fmt"
	"reflect"
)

// findInterseciton is a function for find intersection of two sets
func findIntersection(set1, set2 set) set {
	intersection := set{data: make(map[rune]void)}
	for key := range set1.data {
		if set2.checkContain(key) {
			intersection.put(key)
		}
	}
	return intersection
}

// makeSet is a function for create a new rune set on base of given string
func makeSet(str string) set {
	newSet := set{data: make(map[rune]void)}
	for _, v := range str {
		newSet.put(v)
	}
	return newSet
}

func findAnograms(arr *[]string) *map[string][]string {
	anograms := make(map[string][]string)
	for _, s := range *arr {
		checkSet := makeSet(s)
		flag := false
		for key := range anograms {
			if reflect.DeepEqual(makeSet(key), checkSet) {
				anograms[key] = append(anograms[key], s)
				flag = true
				break
			}
		}
		if !flag {
			anograms[s] = make([]string, 0)
		}

	}
	return &anograms
}

func main() {
	arr := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	result := findAnograms(&arr)
	fmt.Printf("%v\n", *result)
}
