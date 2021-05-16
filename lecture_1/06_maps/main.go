package main

import "log"

func main() {
	m := map[string]int{
		"A": 5,
		"B": 10,
		"C": 15,
	}

	val := m["B"]
	log.Println(val)

	// Zero value
	val = m["D"]
	log.Println(val)

	val, found := m["D"]
	if found {
		log.Printf("Value found %v\n", val)
	} else {
		log.Println("Value not found")
	}

	for key, value := range m {
		log.Printf("%v: %v\n", key, value)
	}

	// There are no sets, use a map as a set
	// `struct{}` is a special "empty struct" which uses no memory whatsoever
	set := map[string]struct{}{
		"A": {},
		"B": {},
	}

	log.Println(set)
}
