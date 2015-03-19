package main

import "fmt"

// STRUCTS

func main() {

// Define a rectangle
rect1 := Rectangle{leftX: 0, TopY: 50, height: 10, width: 10}

// Leave off attribute names if we know the order
// rect1 := Rectangle{0, 50, 10, 10}

// We access values with the dot operator
fmt.Println("Rectangle is", rect1.width, "wide")

// Call the method area for Rectangle
fmt.Println("Area of the rectangle =", rect1.area())

}

// We can define our own types using struct
type Rectangle struct{
	leftX float64
	TopY float64
	height float64
	width float64
}

// We can define methods for our Rectangle by adding the receiver
// rect *Rectangle between func and the function name so we can
// call it with the dot operator
func (rect *Rectangle) area() float64{

	return rect.width * rect.height

}
