// potcut extracts plain text passwords from pot files
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	original := make(map[string]int)
	pot := make(map[string]string)

	// make the scanner for the hash file
	scanner := bufio.NewScanner(hashFile)

	// only scan the first one get the length
	// and store it in the original map
	scanner.Scan()
	temp := scanner.Text()
	length := len(temp)

	original[temp] = 1

	// read throught the hash file
	// increment each time the same hash is seen
	for scanner.Scan() {
		original[scanner.Text()] += 1
	}

	// check for errors in the scanner
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	// make a scanner for the pot file
	scan := bufio.NewScanner(potFile)
	for scan.Scan() {
		temp = scan.Text()
		pot[temp[0:length]] = temp[length+1:]
	}

	// check for errors in the scanner
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	// fine a quick way to do this
	// print out the plains
	for hash, pass := range pot {
		count := original[hash]
		for i := 0; i < count; i++ {
			fmt.Println(pass)
		}
	}
}
