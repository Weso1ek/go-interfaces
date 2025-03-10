package main

import "fmt"

type shape interface {
	area() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) area() float64 {
	return 0.5 * t.height * t.base
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func printAres(s shape) {
	fmt.Println(s.area())
}
