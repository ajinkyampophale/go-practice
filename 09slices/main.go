package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruits = []string{"Mango", "Banana", "Peach"}

	fruits = append(fruits, "Watermelon", "Grapes")
	fmt.Println("Fruits", fruits)

	fruits = append(fruits[:3])
	fmt.Println("Fruits", fruits)

	fruits = append(fruits[1:3])
	fmt.Println("Fruits", fruits)

	fruits = append(fruits[1:])
	fmt.Println("Fruits", fruits)

	var ranks = make([]int, 4)

	ranks[0] = 233
	ranks[1] = 985
	ranks[2] = 345
	ranks[3] = 867
	// ranks[4] = 867

	ranks = append(ranks, 128, 652, 781)
	fmt.Println("Ranks", ranks)

	sort.Ints(ranks)
	fmt.Println("Sorted Ranks", ranks)
	fmt.Println("Is Ranks Sorted", sort.IntsAreSorted(ranks))

	// remove a particular index
	var index = 2;
	var courses = []string{"Reactjs", "Nodejs", "MongoDB", "MySQL", "Redis"}
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses", courses)
}