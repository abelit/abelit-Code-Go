package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// CountLines is a function to count lines from file
func CountLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)

	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return 0, err
	}

	return lines, nil
}

func main() {
	f, err := os.Open("./hello.txt")

	if err != nil {
		fmt.Println("open file err, ", err)
		return
	}

	defer f.Close()

	lines, err := CountLines(f)

	if err != nil {
		fmt.Println("count lines from file err, ", err)
		return
	}

	fmt.Printf("File has %d lines in all.", lines)
}
