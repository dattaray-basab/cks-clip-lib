package main

import (
	"bufio"
	"path/filepath"
	// "fmt"
	"os"
	// "strconv"
)
 
func main() {
	// lines := [2]string{"我想吃！", "I want to eat!"}

	text := `Hi, 0 我想吃！
	How are you?
my name is basab`

	file_name := "file2.txt"
 
	err := createFile(file_name, text)
	if err != nil {
		panic(err)
	}

}

func createFile(fileName string, text string)  error {

	filePath := filepath.Join(fileName) 

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)

	w.WriteString(text)

	w.Flush()
	f.Close()
	return nil
}