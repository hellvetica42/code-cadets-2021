package main

import "log"

func main() {

	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	i := 0
	for i <= 10 {
		log.Println(i)
		i++
	}

	i = 0
	for {
		log.Println(i)
		i++
		if i == 10 {
			break
		}
	}
}
