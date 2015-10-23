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

	start := time.Now()
	for _, url := range (urls) {
		func() {
			start := time.Now()
			http.Get(url)
			fmt.Println(url, " : ", time.Since(start))

		}()
	}
	fmt.Println("Total time: ", time.Since(start))

}
