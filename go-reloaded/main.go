package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("valid arg: go run . input.txt output.txt")
		return
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]
	space, err := os.Open(inputfile)
	if err != nil {
		fmt.Println("err: opening file")
		return
	}
	defer space.Close()

	value, err := io.ReadAll(space)
	if err != nil {
		fmt.Println("err: reading input")
		return
	}

	result := HandleConv(punc(handlepunchnQ(handleCases(handleAtotheAn(string(value))))))

	err = os.WriteFile(outputfile, []byte(result), 0644)
	if err != nil {
		fmt.Println("err: writing output")
		return
	}

	fmt.Println("from grass to GR@CE, GO to craze!!!")
}
