package main

import (
	"fmt"
	"os"
)

// The program reverses a list of integers stored in a file, and writes it to an output file.
func main() {
	fin, err := os.Open("numbers_in.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fin.Close()
	var n int
	fmt.Fscan(fin, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(fin, &a[i])
	}

	fout, err := os.Create("numbers_out.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()
	fmt.Fprintln(fout, n)
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Fprintln(fout, a[i])
	}
}
