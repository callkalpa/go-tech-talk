package main
import (
	"time"
	"net/http"
	"fmt"
)

func main() {
	urls := []string{
		"http://golang.org/",
		"http://www.wso2.com",
		"https://www.wikipedia.org/",
		"http://www.google.com"}

	channel := make(chan bool)

	start := time.Now()
	for _, url := range (urls) {
		go func() {
			start := time.Now()
			http.Get(url)
			fmt.Println(url, " : ", time.Since(start))
			channel  <- true
		}()
	}

	for i := 0; i < len(urls); i++ {
		<-channel
	}
	fmt.Println("Total time: ", time.Since(start))
}
