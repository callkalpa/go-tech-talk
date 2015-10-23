package main

import "fmt"

func printMap(timeZone map[string]int) {
	for code, timeDiff := range timeZone {
		fmt.Println(code, " : ", timeDiff)
	}
}

func main() {
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
	}

	// check whether a key is in the Map
	_, present := timeZone["IST"]
	fmt.Println("Availability of IST timezone: ", present)

	printMap(timeZone)
}
