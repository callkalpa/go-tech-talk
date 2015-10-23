package main
import "os"

func processFile(fileName string) (result string, characterCount int) {
	file, err := os.Open(fileName)

	if err != nil {
		// handle error
	}

	// do some file processing

	defer file.Close()

	content := "Sample content"
	count := len(result)

	return content, count
}
