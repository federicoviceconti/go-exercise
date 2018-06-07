package main

import (
	"math/rand"
	"fmt"
)

type Maximum struct {
	value int
	occurrence int
}

func main() {
	var values []int
	max := Maximum{0, 0}

	//Init map
	for i := 0; i < rand.Intn(100); i++ {
		values = append(values, rand.Intn(30))
	}

	occurrence := map[int]int{}

	for _, val := range values {
		occurrence[val] += 1

		if occurrence[val] > max.occurrence {
			max.value = val
			max.occurrence = occurrence[val]
		}
	}


	fmt.Println("Values", values)
	fmt.Println("Maximum", max)
	fmt.Println("Occ", occurrence)
}
