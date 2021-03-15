package main

import (
	"fmt"
	"os"
)

func createFile(file string) *os.File {
	fmt.Print("creating")
	f, err := os.Create(file)
	if err != nil {
		fmt.Print("create file error")
		os.Exit(0)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Print("writing")
	fmt.Fprintf(f, "data")
}

func closeFile(f *os.File) {
	fmt.Print("closing")
	f.Close()
}

func main() {
	file := "./tmp.txt"
	f := createFile(file)
	defer closeFile(f)
	writeFile(f)
}
