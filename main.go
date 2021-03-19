package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("open file ./example.txt")
	file, openErr := os.Open("./example.txt")
	if openErr != nil {
		panic(openErr)
	}

	fmt.Println("new scanner for file")
	scanner := bufio.NewScanner(file)

	linesCount, countErr := lineCounter(bufio.NewReader(file))
	if countErr != nil {
		panic(countErr)
	}
	fmt.Println(fmt.Sprintf("count lines in file: %d", linesCount))

	fmt.Println("reset file offset")
	newOffeset, seekErr := file.Seek(0, 0)
	if seekErr != nil || newOffeset != 0 {
		panic(seekErr)
	}

	fmt.Println("read file:")
	totalCounter := 0
	emptyCounter := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			emptyCounter++
		}
		fmt.Println(line)
		totalCounter++
	}
	if linesCount-totalCounter == 1 {
		emptyCounter++
	}

	fmt.Println(fmt.Sprintf("%d read line, %d empty", totalCounter, emptyCounter))
}

func lineCounter(reader io.Reader) (int, error) {
	buffer := make([]byte, 32*1024)
	count := 1 // PLEASE NOTE: starting from 1 instead of 0, because in a file there is always at least 1 line, even if empty
	newLine := []byte{'\n'}
	for {
		bytesCount, err := reader.Read(buffer)
		count += bytes.Count(buffer[:bytesCount], newLine)
		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}
