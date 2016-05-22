package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0

	for _, v := range a {
		total += v
	}

	c <- total
	// send total to c the channel
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c) // slice from 0 to the length /2 => first half
	go sum(a[len(a)/2:], c) // sum of second half

	x, y := <-c, <-c // receive from channel c

	fmt.Println(x, y, x+y)
}
