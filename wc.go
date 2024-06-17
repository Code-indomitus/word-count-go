package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
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

	flag.Parse()

	// Get the file name to be read
	filePath := os.Args[len(os.Args)-1]

	// Check if the content is to be read from standard input
	readFromStdin := false
	if len(os.Args) == 1 || (len(os.Args) > 1 && strings.HasPrefix(filePath, "-")) {
		readFromStdin = true
	}

	if readFromStdin {
		processStdIn()
	} else {
		// Open the file that has been passed as an argument
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("wc:", filePath+":", "No such file or directory")
			return
		}
		// Close the file later
		defer file.Close()

		// Process the file
		processFile(file, filePath)
	}
}

func processFile(file *os.File, filePath string) {
	// Execute the appropriate functions based on the flags
	if byteCountFlag {
		fmt.Println(getByteCountFile(file), filePath)
	}

	if lineCountFlag {
		fmt.Println(getLineCountFile(file), filePath)
	}

	if wordCountFlag {
		fmt.Println(getWordCountFile(file), filePath)
	}

	if charCountFlag {
		fmt.Println(getCharCountFile(file), filePath)
	}

	// If no flags are provided, print all
	if len(os.Args) == 2 {
		fmt.Println(getLineCountFile(file), getWordCountFile(file), getByteCountFile(file), filePath)
	}
}

func processStdIn() {
	// Execute the appropriate functions based on the flags
	if lineCountFlag {
		fmt.Println(getLineCountStdIn())
	}

	//TODO: Implement the rest of the flags for standard input
}

func getByteCountFile(file *os.File) int64 {
	resetFilePointer(file)

	// Read file as a buffer
	reader := bufio.NewReader(file)

	var byteCount int64 = 0
	chunk := make([]byte, 1024)

	for {
		// Read a chunk of 1kB of data from the file at at time
		bytesRead, err := reader.Read(chunk)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error:", err)
			return -1
		}

		// If no bytes were read, we've reached the end of the file
		if bytesRead == 0 {
			break
		}

		// Increment the byte count by the number of bytes read
		byteCount += int64(bytesRead)
	}

	return byteCount
}

func getLineCountFile(file *os.File) int64 {
	resetFilePointer(file)

	// Initialize file scanner
	fileScanner := bufio.NewScanner(file)

	var lineCount int64 = 0

	// Iterate over each line in the file and increment the line count
	for fileScanner.Scan() {
		lineCount++
	}

	return lineCount
}

func getWordCountFile(file *os.File) int64 {
	resetFilePointer(file)
	// Initialize file scanner
	fileScanner := bufio.NewScanner(file)

	var wordCount int64 = 0

	// Iterate over each line and get number of words
	for fileScanner.Scan() {
		line := fileScanner.Text()

		// Increment word count by number of words in the line
		wordCount += int64(len(strings.Fields(line)))
	}

	return wordCount
}

func getCharCountFile(file *os.File) int64 {
	resetFilePointer(file)
	// Initialize file scanner
	fileScanner := bufio.NewScanner(file)

	var charCount int64 = 0

	// Iterate over each line and get number of characters
	for fileScanner.Scan() {
		line := fileScanner.Text()

		// Increment word count by number of words in the line
		charCount += int64(utf8.RuneCountInString(line))

		// // Add the newline character
		charCount++
	}

	// Remove the last newline character
	charCount--

	return charCount
}

func resetFilePointer(file *os.File) {
	// Reset the file pointer to the beginning of the file
	_, err := file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func getLineCountStdIn() int64 {
	var lineCount int64 = 0
	scanner := bufio.NewScanner(os.Stdin)

	// Read input line by line
	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}
