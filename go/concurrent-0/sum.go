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
func sum(ch chan typeSum) {
	channel := <-ch
	data, _ := readFile(channel.filepath)

	// if err != nil {
	// 	return 0, err
	// }

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}

	channel.sum = _sum
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

	join := make(chan int)

	sumChannel := make(chan typeSum)

	var totalSum int64
	sums := make(map[int][]string)
	for _, path := range os.Args[1:] {
		inSum := typeSum{path, 0}
		sumChannel <- inSum
		go sum(sumChannel)
		outSum := <-sumChannel
		totalSum += int64(outSum.sum)

		sums[outSum.sum] = append(sums[outSum.sum], path)
	}

	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}

	<-join

}
