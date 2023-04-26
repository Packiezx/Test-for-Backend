package main

import (
	"fmt"
)

func IntersectArray(a []string, b []string) []string {
	result := []string{}
	for _, i := range a {
		for _, otherValue := range b {
			if i == otherValue {
				result = append(result, i)
				break
			}
		}
	}
	return result
}

func MergeArray(a []string, b []string) ([]string, []string) {
	intersect := IntersectArray(a, b)
	result1 := append([]string{}, a...)
	result2 := []string{}
	for _, i := range b {
		found := false
		for _, intersectValue := range intersect {
			if i == intersectValue {
				found = true
				break
			}
		}
		if !found {
			result1 = append(result1, i)
			result2 = append(result2, i)
		}
	}
	for _, i := range a {
		found := false
		for _, intersectValue := range intersect {
			if i == intersectValue {
				found = true
				break
			}
		}
		if !found {
			result2 = append(result2, i)
		}
	}
	return result1, result2
}

func main() {
	a := []string{"a", "b", "c"}
	b := []string{"b", "c", "d"}

	fmt.Println("ข้อมูลที่ซ้ำกัน", IntersectArray(a, b))
	fmt.Print("ข้อมูลที่รวมกัน และ ข้อมูลที่ตัดข้อมูลที่ซ้ำกัน")
	fmt.Println(MergeArray(a, b))

}
