package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readStringsByChar(f *os.File) []string {
	var strings []string
	var buffer []byte
	var err error
	for err != io.EOF {
		chr := make([]byte, 1)
		_, err = f.Read(chr)
		if chr[0] == '\n' || err == io.EOF {
			strings = append(strings, string(buffer))
			buffer = []byte{}
		} else {
			buffer = append(buffer, chr...)
		}
	}
	return strings
}

func main() {
	fin, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fin.Close()
	
	lines := readStringsByChar(fin)
	fmt.Println(len(lines))
	
	
	/*
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	*/

}
