package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	fmt.Println("LeetCode solutions!")
	v := Vertex{3, 4}
	v.X = 100
	fmt.Println(v.X)
}
