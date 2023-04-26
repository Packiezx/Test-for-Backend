package main

import (
	"fmt"
)

func IntersectArray(a []interface{}, b []interface{}) []interface{} {
	var result []interface{}

	for _, i := range a {
		for _, y := range b {
			if i == y {
				result = append(result, i)
				break
			}
		}
	}

	return result
}

func MergeArray(a []interface{}, b []interface{}) []interface{} {
	var result []interface{}

	for _, i := range a {
		found := false
		for _, y := range result {
			if i == y {
				found = true
				break
			}
		}
		if !found {
			result = append(result, i)
		}
	}

	for _, i := range b {
		found := false
		for _, y := range result {
			if i == y {
				found = true
				break
			}
		}
		if !found {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	a := []interface{}{"a", "b", "c"}
	b := []interface{}{"b", "c", "d"}

	intersect := IntersectArray(a, b)
	fmt.Println(MergeArray(a, b))
	fmt.Println(intersect)

}
