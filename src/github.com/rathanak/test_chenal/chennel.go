package main

var c = make(chan int, 10)
var a string

func f() {
	a = "Hello, world"
	c <- 0
}
func main() {
	go f()
	<-c
	print(a)
	print(c)
}
