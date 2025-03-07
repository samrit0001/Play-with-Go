package main

import "fmt"

type shape interface {
	area() int
	circum() int
}

type circle struct {
	radius int
}

func (c circle) area() int {
	return c.radius * c.radius
}

func (c circle) circum() int {
	return 2 * 3 * c.radius
}

type square struct {
	side int
}

func (s square) area() int {
	return s.side * s.side
}

func (s square) circum() int {
	return 4 * s.side
}

// This is good but it will work only with circle or square
// func PrintInfo(c circle) {
// 	fmt.Printf("area is %d", c.area())
// 	fmt.Printf("circum is %d", c.circum())
// }

//use interface to achive same
func PrintInfo(s shape) {
	fmt.Printf("area is %d\n", s.area())
	fmt.Printf("circum is %d\n", s.circum())
}

func main() {
	var circle1 = circle{radius: 10}

	square1 := square{side: 5}

	fmt.Println(circle1.area())
	fmt.Println(square1.area())
	fmt.Println(circle1.circum())
	fmt.Println(square1.circum())

	PrintInfo(circle1)
	PrintInfo(square1)
}
