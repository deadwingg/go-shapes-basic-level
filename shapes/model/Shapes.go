package model

import (
	"fmt"
	"github.com/google/uuid"
	"math"
)

type Shape interface {
	CalculateArea() float64
	Print()
}

type Ellipse struct {
	Id     uuid.UUID `json:"id"`
	RadioA float64   `json:"radioA"`
	RadioB float64   `json:"radioB"`
}

func NewEllipse(radioA, radioB float64) *Ellipse {
	return &Ellipse{
		Id:     uuid.New(),
		RadioA: radioA,
		RadioB: radioB,
	}
}
func (e Ellipse) CalculateArea() float64 {
	return math.Pi * e.RadioA * e.RadioB / 4.0
}

func (e Ellipse) Print() {
	fmt.Printf("Ellipse:\n\tId: %s\n\tRadius A: %f\n\tRadius B: %f\n", e.Id, e.RadioA, e.RadioB)
}

type Rectangle struct {
	Id     uuid.UUID `json:"id"`
	Length float64   `json:"length"`
	Width  float64   `json:"width"`
}

func NewRectangle(length, width float64) *Rectangle {
	return &Rectangle{
		Id:     uuid.New(),
		Length: length,
		Width:  width,
	}
}

func (r Rectangle) CalculateArea() float64 {
	return r.Width * r.Length
}

func (r Rectangle) Print() {
	fmt.Printf("Rectangle:\n\tId: %s\n\tWidth: %f\n\tLength: %f\n", r.Id, r.Width, r.Length)
}

type Triangle struct {
	Id     uuid.UUID `json:"id"`
	Base   float64   `json:"base"`
	Height float64   `json:"height"`
}

func NewTriangle(base, height float64) *Triangle {
	return &Triangle{
		Id:     uuid.New(),
		Base:   base,
		Height: height,
	}
}
func (t Triangle) CalculateArea() float64 {
	return t.Base * t.Height / 2.0
}

func (t Triangle) Print() {
	fmt.Printf("Triangle:\n\tId: %s\n\tBase: %f\n\tHeight: %f\n", t.Id, t.Base, t.Height)
}
