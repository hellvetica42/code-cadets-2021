package main

import "log"

func print(s []int) {
	log.Printf("elements: %v, len=%v, cap=%v\n", s, len(s), cap(s))
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	print(array[:])

	slice := array[0:3]
	print(slice)

	slice = append(slice, 6)
	print(slice)
	print(array[:])

	emptySlice := make([]int, 5)
	print(emptySlice)

	emptySlice = append(emptySlice, 1)
	print(emptySlice)

	emptySliceWithExtraCap := make([]int, 5, 10)
	print(emptySliceWithExtraCap)

	//////////////////////////////////////////

	sliceToIterate := []int{50, 60, 70, 80}

	for idx, val := range sliceToIterate {
		log.Printf("%v: %v\n", idx, val)
	}

	for _, val := range sliceToIterate {
		log.Printf("%v\n", val)
	}

	for val := range sliceToIterate {
		log.Printf("%v\n", val)
	}
}
