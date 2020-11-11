package main

import (
	"io/ioutil"
	"os"
	"testing"
	"fmt"
)

func BenchmarkIndexer(b *testing.B) {

	file, err := os.Open("/home/coolbry95/wrappertesting/testlist")
	if err != nil {
		fmt.Println("err opening file")
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("err reading file")
	}

	for n := 0; n < b.N; n++ {
		for {
		str := Indexer(contents)
		if len(str) == 0 {
			break
		}
	}
	}
}
