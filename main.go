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
 
	// Create a file for writing
	f, _ := os.Create("file2.txt")
 
	// Create a writer
	w := bufio.NewWriter(f)

		w.WriteString(text)
 
	// for i, s := range lines {
	// 	fmt.Println(i)
	// 	w.WriteString(strconv.Itoa(i) + " " + s + "\n")
	// }
 
	// Very important to invoke after writing a large number of lines
	w.Flush()
}