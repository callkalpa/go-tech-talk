package main

func doSomething(channel chan int) {
	// do something here
	channel <- 1
}

func main() {
	channel := make(chan int)
	go doSomething(channel)

	// wait until doSomething() is completed
	<-channel

	// continue with the execution of main()
}