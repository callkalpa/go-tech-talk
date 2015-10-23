package main
import (
	"fmt"
)

func B() {
	panic("Panic in function B")
}


func A() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	B()
}

func main() {
	A()
}