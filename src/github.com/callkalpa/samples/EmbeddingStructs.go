package main
import "fmt"

type Parent struct {
	surname string
}

type Child struct {
	Parent
	isInfant bool
}

func (parent *Parent) toString() {
	fmt.Println(parent.surname)
}

func main() {
	parent := Parent{"some surname"}
	child := Child{parent, false}
	child.toString()
}