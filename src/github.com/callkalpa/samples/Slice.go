package main
import "fmt"

func PrintSlice(slice []string) {
	for index, value := range (slice) {
		fmt.Println(index, " : ", value)
	}
}

func main() {
	slice := []string{"A", "B", "C"}
	PrintSlice(slice)

	fmt.Println("Second element of the slice : ", slice[1:2])

	slice = append(slice, "D", "E", "F")
	PrintSlice(slice)
}