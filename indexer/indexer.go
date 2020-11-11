package main

import (
	"fmt"
	"bytes"
	"os"
	"io/ioutil"
)


// this is a test to see if it is a fast way to parse a mmap file
// for lines

var i int

func Indexer(slice []byte) []byte {
	index := bytes.IndexByte(slice[i:], '\n')

	if index == -1 {
		return nil
	}
	//fmt.Println(slice[i:])

	//fmt.Println(i)
	//fmt.Println(index)

	temp := i
	i += index + 1
	// would copy be more efficent here
	return slice[temp:index+temp]
}

func main() {
	file, err := os.Open("/home/coolbry95/wrappertesting/testlist")
	if err != nil {
		fmt.Println("err opening file")
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("err reading file")
	}

	for {
		str := Indexer(contents)
		if len(str) == 0 {
			break
		}
		fmt.Println(string(str))
	}
}
