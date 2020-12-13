package main

import (
	"fmt"
	"math"
)

type Person struct {
	Name  string
	State string
}

func heapSortByName(people []*Person) {
	maxIdx := len(people) - 1
	for idx := int(math.Floor(float64(len(people))/2)) - 1; idx >= 0; idx-- {
		makeMaxHeapByName(people, idx, maxIdx)
	}

	people[0], people[maxIdx] = people[maxIdx], people[0]
	maxIdx--
	for maxIdx > 0 {
		makeMaxHeapByName(people, 0, maxIdx)
		people[0], people[maxIdx] = people[maxIdx], people[0]
		maxIdx--
	}
}

func makeMaxHeapByName(people []*Person, idx, maxIdx int) {
	largestIdx := idx

	leftChildNodeIdx := idx*2 + 1
	rightChildNodeIdx := idx*2 + 2

	if leftChildNodeIdx <= maxIdx && people[leftChildNodeIdx].Name > people[largestIdx].Name {
		largestIdx = leftChildNodeIdx
	}
	if rightChildNodeIdx <= maxIdx && people[rightChildNodeIdx].Name > people[largestIdx].Name {
		largestIdx = rightChildNodeIdx
	}

	if largestIdx != idx {
		people[idx], people[largestIdx] = people[largestIdx], people[idx]
		makeMaxHeapByName(people, largestIdx, maxIdx)
	}
}

func heapSortByState(people []*Person) {
	maxIdx := len(people) - 1
	for idx := int(math.Floor(float64(len(people))/2)) - 1; idx >= 0; idx-- {
		makeMaxHeapByState(people, idx, maxIdx)
	}

	people[0], people[maxIdx] = people[maxIdx], people[0]
	maxIdx--
	for maxIdx > 0 {
		makeMaxHeapByState(people, 0, maxIdx)
		people[0], people[maxIdx] = people[maxIdx], people[0]
		maxIdx--
	}
}

func makeMaxHeapByState(people []*Person, idx, maxIdx int) {
	largestIdx := idx

	leftChildNodeIdx := idx*2 + 1
	rightChildNodeIdx := idx*2 + 2

	if leftChildNodeIdx <= maxIdx && people[leftChildNodeIdx].State > people[largestIdx].State {
		largestIdx = leftChildNodeIdx
	}
	if rightChildNodeIdx <= maxIdx && people[rightChildNodeIdx].State > people[largestIdx].State {
		largestIdx = rightChildNodeIdx
	}

	if largestIdx != idx {
		people[idx], people[largestIdx] = people[largestIdx], people[idx]
		makeMaxHeapByState(people, largestIdx, maxIdx)
	}
}

func mergeSortByState(people []*Person) {
	destination := make([]*Person, len(people))
	copy(destination, people)
	mergeSortSubByState(destination, people, 0, len(people))
}

func mergeSortSubByState(people, destination []*Person, startIdx, endIdx int) {
	if endIdx-startIdx < 2 {
		return
	}

	midIdx := int(math.Floor(float64(startIdx+endIdx) / 2))

	mergeSortSubByState(destination, people, startIdx, midIdx)
	mergeSortSubByState(destination, people, midIdx, endIdx)

	i := startIdx
	j := midIdx
	for k := startIdx; k < endIdx; k++ {
		if j >= endIdx || (i < midIdx && people[i].State <= people[j].State) {
			destination[k] = people[i]
			i++
		} else {
			destination[k] = people[j]
			j++
		}
	}
}

func printPeople(people []*Person) {
	for _, person := range people {
		fmt.Printf("Name: %s, State: %s\n", person.Name, person.State)
	}
}

func main() {
	people := []*Person{
		&Person{Name: "Thomas", State: "New Hampshire"},
		&Person{Name: "Annie", State: "New Hampshire"},
		&Person{Name: "Chris", State: "Maine"},
		&Person{Name: "Jackie", State: "Vermont"},
		&Person{Name: "Marie", State: "New Hampshire"},
		&Person{Name: "Daniel", State: "Maine"},
		&Person{Name: "Vivian", State: "Maine"},
		&Person{Name: "Michelle", State: "Maine"},
	}

	fmt.Println("\n--- Unsorted ----")
	printPeople(people)
	fmt.Println("\n--- Sorted by Name (heap sort: unstable) ----")
	heapSortByName(people)
	printPeople(people)
	fmt.Println("\n--- Sorted by State (heap sort: unstable) ----")
	heapSortByState(people)
	printPeople(people)

	fmt.Println("\n************************************")
	fmt.Println("Now trying with a stable sort on state")
	fmt.Println("************************************")

	people = []*Person{
		&Person{Name: "Thomas", State: "New Hampshire"},
		&Person{Name: "Annie", State: "New Hampshire"},
		&Person{Name: "Chris", State: "Maine"},
		&Person{Name: "Jackie", State: "Vermont"},
		&Person{Name: "Marie", State: "New Hampshire"},
		&Person{Name: "Daniel", State: "Maine"},
		&Person{Name: "Vivian", State: "Maine"},
		&Person{Name: "Michelle", State: "Maine"},
	}

	fmt.Println("\n--- Unsorted ----")
	printPeople(people)
	fmt.Println("\n--- Sorted by Name (heap sort: unstable) ----")
	heapSortByName(people)
	printPeople(people)
	fmt.Println("\n--- Sorted by State (merge sort: stable) ----")
	mergeSortByState(people)
	printPeople(people)
}
