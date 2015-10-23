package main
import "fmt"

func main() {

	arr1 := [3] int{1, 2, 3}
	arr2 := arr1

	fmt.Println("arr1 is equal to arr2: ", arr1 == arr2) // true

	arr3 := [4] int{1, 2, 3, 4}
	fmt.Println("arr1 is equal to arr3: ", arr1 == arr3) // this is an invalid operation

}
