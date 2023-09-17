package common

import (
	"bufio"
	"os"
)

var GetWordsFromFile = func(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Scan all words from the file.
	var words []string
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return words, nil
}
