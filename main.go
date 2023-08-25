package main

import (
	"bufio"
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
 
	// Create a file for writing
	// Create a writer
	// for i, s := range lines {
	// 	fmt.Println(i)
	// 	w.WriteString(strconv.Itoa(i) + " " + s + "\n")
	// }
	// Very important to invoke after writing a large number of lines
	err := createFile(file_name, text)
	if err != nil {
		panic(err)
	}

}

func createFile(file_name string, text string)  error {

	f, err := os.Create(file_name)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)

	w.WriteString(text)

	w.Flush()
	f.Close()
	return nil
}