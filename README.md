# Word Count Command-Line Tool

This command-line tool is designed to count the number of bytes, lines, words, and characters in a given text file or standard input. The tool is implemented in Go and replicates the functionality of the wc command found in Unix-like systems.

## Features
- Count the number of bytes in a file or standard input.
- Count the number of lines in a file or standard input.
- Count the number of words in a file or standard input.
- Count the number of characters in a file or standard input.

## Installation
To use this tool, you need to have Go installed on your system. You can download and install Go from [golang.org](https://go.dev/).

1. Clone this repository:
    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. Build the executable:
    ```sh
    go build -o wc.go
    ```

3. The executable `wc` will be generated in the current directory.

## Usage
```sh
./wc.exe [OPTION]... [FILE]
```

### Options
- -c: Print the byte counts.
- -l: Print the newline counts.
- -w: Print the word counts.
- -m: Print the character counts.

If no option is provided, the tool will display the line count, word count, and byte count by default.

## Examples
Count bytes, lines, words, and characters in a file
```sh
./wc.exe -c -l -w -m example.txt

# Output:
# 123 456 789 101 example.txt
```

Count lines in a file
```sh
./wc.exe -l example.txt

# Output:
# 123 example.txt
```

Count words in a file
```sh
./wc.exe -w example.txt

# Output:
# 456 example.txt
```

Count characters in a file
```sh
./wc.exe -m example.txt

# Output:
# 789 example.txt
```

Count lines from standard input
```sh
echo "Hello, world!" | ./wc.exe -l

# Output:
# 1
```

Default behavior (line count, word count, and byte count)
```sh
./wc.exe example.txt

# Output:
# 123 456 789 example.txt
```
