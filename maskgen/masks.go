// maskgen generates masks for use with hashcat
package main

// this version uses a table instead of a switch

import (
	"fmt"
	"io"
	"os"
)

// ascii table that defines a characters type
var ascii_table = [256]byte{
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	115, 115, 115, 115, 115, 115, 115, 115,
	115, 115, 115, 115, 115, 115, 115, 115,
	100, 100, 100, 100, 100, 100, 100, 100,
	100, 100, 115, 115, 115, 115, 115, 115,
	115, 117, 117, 117, 117, 117, 117, 117,
	117, 117, 117, 117, 117, 117, 117, 117,
	117, 117, 117, 117, 117, 117, 117, 117,
	117, 117, 117, 115, 115, 115, 115, 115,
	115, 108, 108, 108, 108, 108, 108, 108,
	108, 108, 108, 108, 108, 108, 108, 108,
	108, 108, 108, 108, 108, 108, 108, 108,
	108, 108, 108, 115, 115, 115, 115, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98,
}

func main() {

	// check the commandline
	if len(os.Args) == 1 {
		fmt.Println("no file specified")
		fmt.Fprintf(os.Stderr, "usage: %s <wordlist>\n", os.Args[0])
		os.Exit(-1)
	}

	// open up the file
	in, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	defer in.Close()

	// i is going to keep track of where we are in the buffer
	var i int
	// c hold the current character
	var c byte

	// buffer fo rthe mask and for the data
	mask := make([]byte, 8192, 8192)
	buf := make([]byte, 8192, 8192)

	for {
		// read in a chunk from the file
		n, err := in.Read(buf)
		if err != nil && err != io.EOF {
			return
		}

		// if the chunk read was less than the size of buf
		// then check to see if there is a \n at the end
		// if not add it
		if n < 8192 {
			if buf[n] != 10 {
				buf[n] = 10
			}
		}

		// loop through the buf
		for j := 0; j < n; j++ {

			// if we are the size of mask buffer then flush it
			if i >= 8192-1 {
				// only write up to i becuase it will be 8181
				// if not it will index error
				os.Stdout.Write(mask[:i])

				// move i back to the beginning of the mask buffer
				i = 0
			}

			// get the character type from the table
			c = ascii_table[buf[j]]

			// check if it is binary or unicode
			if c == 98 || c > 191 {
				// check to see if it is a \n
				// if so add a \n to the mask
				if c == 10 {
					mask[i] = 10
					i += 1
				}
				continue
			}

			// add the question mark and character type
			mask[i] = 63
			mask[i+1] = c
			i += 2

		}

		// if we didnt read in a full buffer
		if n < 8192 {
			// write up to what is filled
			os.Stdout.Write(mask[:i])
			return
		}
	}
}
