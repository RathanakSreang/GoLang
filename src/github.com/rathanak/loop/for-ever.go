package main

import "fmt"

func main() {
	sum := 1

	for {
		sum += sum
		if sum > 1000 {
			break
		}
	}
	fmt.Println(sum)
}
