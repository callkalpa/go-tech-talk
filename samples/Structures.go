package main
import "fmt"

type Point struct {
	x, y int
}

func (point Point) toString() {
	fmt.Println("X: ", point.x, ", Y: ", point.y)
}

func (point Point) scale(scale int) {
	point.x *= scale; point.y *= scale
}

func (point *Point) scale2(scale int) {
	point.x *= scale; point.y *= scale
}

func main() {
	point := Point{3, 4}
	point.toString()

	point.scale(2)
	point.toString()

	pointPointer := &point
	pointPointer.scale2(5)
	point.toString()
}