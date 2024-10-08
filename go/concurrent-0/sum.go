package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// read a file from a filepath and return a slice of bytes
func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

// sum all bytes of a file
func sum(filepath string, ch chan typeSum) {
	data, _ := readFile(filepath)

	// if err != nil {
	// 	return 0, err
	// }

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}

	channel := typeSum{filepath, _sum}

	ch <- channel
}

type typeSum struct {
	filepath string
	sum      int
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	sumChannel := make(chan typeSum)

	var totalSum int64
	sums := make(map[int][]string)
	for _, path := range os.Args[1:] {
		go sum(path, sumChannel)
	}

	for i := 0; i < len(os.Args[1:]); i++ {
		outSum := <-sumChannel
		totalSum += int64(outSum.sum)
		sums[outSum.sum] = append(sums[outSum.sum], outSum.filepath)
	}

	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}

}
