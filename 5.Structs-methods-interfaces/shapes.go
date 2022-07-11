package structs

import "math"

// In golang, to implement an interface you only need to
// satisfies the interface. For example, both Circle and
// Rectangle implement an Area() method, so they satisfy
// the interface.
// In Go interface resolution is implicit. If the type you 
// pass in matches what the interface is asking for, it will compile.
type Shape interface {
	Area() float64
}

// Using structs avoid to rename the functions RectangleArea
// and RectanglePerimeter, this is a data oriented approach.
type Rectange struct {
	Width float64
	Height float64
}

func (r Rectange) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectange) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//func Area(rectangle Rectange) float64 {
//	return rectangle.Width * rectangle.Height
//}
