package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Declare the flag options of word count
var byteCountFlag bool
var lineCountFlag bool
var wordCountFlag bool
var charCountFlag bool

func main() {
	// Parse command line arguments for flags
	flag.BoolVar(&byteCountFlag, "c", false, "print the byte counts")
	flag.BoolVar(&lineCountFlag, "l", false, "print the newline counts")
	flag.BoolVar(&wordCountFlag, "w", false, "print the word counts")
	flag.BoolVar(&charCountFlag, "m", false, "print the character counts")

	// flag.BoolVar(&byteCountFlag, "bytes", false, "print the byte counts") TODO remove this later
	flag.Parse()

	//TODO: Remove this later
	// fmt.Println(len(os.Args))

	// fmt.Println("byteCountFlag:", byteCountFlag)

	// Get the file name to be read
	fileName := os.Args[len(os.Args)-1]

	// Open the file that has been passed as an argument
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("wc:", fileName+":", "No such file or directory")
		return
	}

	// Close the file later
	defer file.Close()

	// Execute the appropriate functions based on the flags
	if byteCountFlag {
		printByteCount(file, fileName)
	}

	if lineCountFlag {
		printLineCount(file, fileName)
	}

	if wordCountFlag {
		printWordCount(file, fileName)
	}

	if charCountFlag {
		printCharCount(file, fileName)
	}
}

func printByteCount(file *os.File, fileName string) {
	// Read file as a buffer
	reader := bufio.NewReader(file)

	var byteCount int64 = 0

	for {
		// Read a chunk of 1kB of data from the file at at time
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

	fmt.Println(byteCount, fileName)
}

func printLineCount(file *os.File, fileName string) {
	// Initialize file scanner
	fileScanner := bufio.NewScanner(file)

	var lineCount int64 = 0

	// Iterate over each line in the file and increment the line count
	for fileScanner.Scan() {
		lineCount++
	}

	fmt.Println(lineCount, fileName)
}

func printWordCount(file *os.File, fileName string) {
	// Initialize file scanner
	fileScanner := bufio.NewScanner(file)

	var wordCount int64 = 0

	// Iterate over each line and get number of words
	for fileScanner.Scan() {
		line := fileScanner.Text()

		// Increment word count by number of words in the line
		wordCount += int64(len(strings.Fields(line)))
	}

	fmt.Println(wordCount, fileName)
}

func printCharCount(file *os.File, fileName string) {
	// Initialize file scanner
	fileScanner := bufio.NewScanner(file)

	var charCount int64 = 0

	// Iterate over each line and get number of characters
	for fileScanner.Scan() {
		line := fileScanner.Text()

		// Increment word count by number of words in the line
		charCount += int64(len(line))
	}

	fmt.Println(charCount, fileName)
}
