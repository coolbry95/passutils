// potcut extracts plain text passwords from pot files
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/coolbry95/bitset"
	"github.com/dchest/siphash"
)

func main() {

	// check the commandline
	if len(os.Args) != 3 {
		log.Printf("usage: %s original pot\n", os.Args[0])
		os.Exit(-1)
	}

	// open the hash file for reading
	hashFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Println(err.Error())
	}
	defer hashFile.Close()

	// open the pot file for reading
	potFile, err := os.Open(os.Args[2])
	if err != nil {
		log.Println(err.Error())
	}
	defer potFile.Close()

	// make the original and pot map
	original := bitset.NewBitset(500)

	// make the scanner for the hash file
	scanner := bufio.NewScanner(hashFile)

	// only scan the first one get the length
	// and store it in the original map
	scanner.Scan()
	temp := scanner.Bytes()

	// the size of the hash is to big for the bitset
	// this
	// https://github.com/Workiva/go-datastructures/blob/master/hashmap/fastinteger/hashmap.go#L47
	fmt.Println(siphash.Hash(9, 10, temp))
	original.Set(siphash.Hash(9, 10, temp) % 64 )

	// read throught the hash file
	// increment each time the same hash is seen
	for scanner.Scan() {
		original.Set(siphash.Hash(1, 0, scanner.Bytes()) % 64)
	}

	// check for errors in the scanner
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	// make a scanner for the pot file
	scan := bufio.NewScanner(potFile)
	for scan.Scan() {
		temp := scan.Bytes()
		if original.Get(siphash.Hash(1, 0, temp)) {
			continue
		}
		//fmt.Println(string(temp))
	}

	// check for errors in the scanner
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
