package main

import (
	"fmt"
	"math"
)

func main() {
	//var list1 = []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	//var list2 = []string{"KFC", "Shogun", "Burger King"}
	var list1 = []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	var list2 = []string{"KFC", "Burger King", "Tapioca Express", "Shogun"}
	fmt.Println(findRestaurant(list1, list2))
}

func findRestaurant(list1 []string, list2 []string) (ans []string) {
	index := make(map[string]int, len(list1))
	for i, s := range list1 {
		index[s] = i
	}

	indexSum := math.MaxInt32
	for i, s := range list2 {
		if j, ok := index[s]; ok {
			if i+j < indexSum {
				indexSum = i + j
				ans = []string{s}
			} else if i+j == indexSum {
				ans = append(ans, s)
			}
		}
	}
	return
}
