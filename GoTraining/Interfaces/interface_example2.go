package main

import (
	"fmt"
)

// Shift interface with one method
type Shift interface {
	Shift(dx, dy, dz float64)
}

// Draw interface with one method
type Draw interface {
	Draw()
}

// Point2D struct representing a point in 2D space
type Point2D struct {
	X, Y float64
}

// Point3D struct representing a point in 3D space
type Point3D struct {
	X, Y, Z float64
}

// Implementing the Shift method for Point2D
func (p *Point2D) Shift(dx, dy, dz float64) {
	p.X += dx
	p.Y += dy
	// Point2D doesn't have a Z component, so dz is ignored
}

// Implementing the Draw method for Point2D
func (p Point2D) Draw() {
	fmt.Printf("Drawing Point2D at (X: %.2f, Y: %.2f)\n", p.X, p.Y)
}

// Implementing the Shift method for Point3D
func (p *Point3D) Shift(dx, dy, dz float64) {
	p.X += dx
	p.Y += dy
	p.Z += dz
}

// Implementing the Draw method for Point3D
func (p Point3D) Draw() {
	fmt.Printf("Drawing Point3D at (X: %.2f, Y: %.2f, Z: %.2f)\n", p.X, p.Y, p.Z)
}

func main() {
	// Creating an instance of Point2D
	p2d := Point2D{X: 1.0, Y: 2.0}
	// Creating an instance of Point3D
	p3d := Point3D{X: 3.0, Y: 4.0, Z: 5.0}

	// Using Shift and Draw methods for Point2D
	p2d.Draw()
	p2d.Shift(1.0, 1.0, 0.0)
	p2d.Draw()

	// Using Shift and Draw methods for Point3D
	p3d.Draw()
	p3d.Shift(1.0, 1.0, 1.0)
	p3d.Draw()
}
