package main

import (
	"fmt"
	"log"
)

// Shape ...
type Shape interface {
	Draw()
}

// Rectangle ...
type Rectangle struct{}

// Draw of rectangle
func (r Rectangle) Draw() {
	fmt.Println("Inside Rectangle::draw() method.")
}

// Square ...
type Square struct{}

// Draw of square
func (s Square) Draw() {
	fmt.Println("Inside Square::draw() method.")
}

// Circle ...
type Circle struct{}

// Draw of circle
func (c Circle) Draw() {
	fmt.Println("Inside Circle::draw() method.")
}

// ShapeFactory ...
type ShapeFactory struct{}

// getShape is shape factory method
func (f ShapeFactory) getShape(ShapeType string) (Shape, error) {
	if ShapeType == "" {
		return nil, fmt.Errorf("params error")
	}
	if ShapeType == "Rectangle" {
		return Rectangle{}, nil
	} else if ShapeType == "Square" {
		return Square{}, nil
	} else if ShapeType == "Circle" {
		return Circle{}, nil
	}
	return nil, fmt.Errorf("%s Shape is not found", ShapeType)
}

func main() {
	f := ShapeFactory{}
	obj, err := f.getShape("Circle")
	if err != nil {
		log.Fatalln(err)
	}
	obj.Draw()
}
