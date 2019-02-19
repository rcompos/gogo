package main

import(
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func main() {
	fmt.Println("Hello")

	c := Circle{x: 1, y: 2, r: 8}
	fmt.Println("x ", c.x)
	fmt.Println("y ", c.y)
	fmt.Println("r ", c.r)

	d := Circle{2, 3, 9}
	fmt.Println("dx ", d.x, "dy ", d.y, "dr ", d.r)

	e := Circle{0, 0, 1}
	fmt.Println("ex ", e.x, "ey ", e.y, "er ", e.r)

	//fmt.Println(circleArea(&e))
	fmt.Println(e.area())

	r := Rectangle{0, 0, 1, 1}
	fmt.Println(r.area())
}

//Function
//func circleArea(c *Circle) float64 {
//	return math.Pi * c.r * c.r
//}

//Method
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
