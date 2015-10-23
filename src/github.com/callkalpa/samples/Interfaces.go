package main
import "math"

type Cup struct {
	color string; radius, length float64; sellingPrice float64
}

type SellingItem interface {
	price() float64
}

func (cup Cup) price() float64 {
	return cup.sellingPrice * 0.95
}

type Shape interface {
	volume() float64
}

func (cup Cup) volume() float64 {
	return math.Pi * math.Pow(cup.radius, 2) * cup.length
}

func main() {
	var item SellingItem
	var shape Shape

	cup := Cup{color:"", radius:2.5, length:10, sellingPrice:150}
	item = cup

	shape = cup
	_ = item
	_ = shape

}