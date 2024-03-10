package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// Declare the flag options of word count
var byteCountFlag bool

func main() {
	// Parse command line arguments for flags
	flag.BoolVar(&byteCountFlag, "c", false, "print the byte counts")
	flag.BoolVar(&byteCountFlag, "bytes", false, "print the byte counts")
	flag.Parse()

	fmt.Println(len(os.Args))

	fmt.Println("byteCountFlag:", byteCountFlag)

	// Get the file name to be read
	fileName := os.Args[len(os.Args)-1]

	// Open the file that has been passed as an argument
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("wc:", fileName+":", "No such file or directory")
		return
	}

	defer file.Close()

	// Read file as a buffer
	reader := bufio.NewReader(file)

	var byteCount int64 = 0

	for {
		// Read a chunk of 1kB of data from the file
		chunk := make([]byte, 1024)
		bytesRead, err := reader.Read(chunk)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error:", err)
			return
		}

		// If no bytes were read, we've reached the end of the file
		if bytesRead == 0 {
			break
		}

		// Increment the byte count by the number of bytes read
		byteCount += int64(bytesRead)
	}

	// // Check which flags have been set
	// for index, value := range os.Args {
	// 	// Print file name
	// 	fmt.Println("File", index, ":", value)
	// }
	fmt.Println(byteCount, fileName)
}
