package main

// this is just a test
// this runs very slow for a reason I have not
// tried to figure out yet

import (
	"fmt"
	"io"
	"os"

	"sync"
)
var wg sync.WaitGroup

func main() {

	// check the commandline
	if len(os.Args) == 1 {
		fmt.Println("no file specified")
		fmt.Fprintf(os.Stderr, "usage: %s <wordlist>\n", os.Args[0])
		os.Exit(-1)
	}

	info, err := os.Stat(os.Args[1])
	if info.IsDir() {
		fmt.Println("Cannot use directory")
		os.Exit(-1)
	} else if err != nil {
		fmt.Println("could not open file for reading")
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	// open up the file
	in, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	defer in.Close()

	stuff := make(chan []byte)
	out := make(chan []byte)
	quit := make(chan struct{})


	go Out(out, quit)
	for i := 0; i < 5; i++ {

		go Masks(stuff, out, quit)
		wg.Add(1)

	}
	produce(stuff, in)

	wg.Wait()
}

func Out(buf chan []byte, quit chan struct{}) {

		fmt.Println("here")
		for {
			select {
			case stuff := <-buf:
				os.Stdout.Write(stuff)
			case <-quit:
				break
			}
		}
}

func produce(out chan []byte, in io.Reader) {
	// read in a chunk from the file
	buf := make([]byte, 8192, 8192)
		fmt.Println("here")
		for {
			n, err := in.Read(buf)
			fmt.Println(n)
			if err != nil && err != io.EOF {
				close(out)
				return
			}
			out <- buf
		}
}

func Masks(in, out chan []byte, quit chan struct{}) {
		fmt.Println("here")
	//go func() {
	var i int
	mask := make([]byte, 8192, 8192)
	buf := make([]byte, 8192, 8192)
		for {
			buf = <-in
			n := len(buf)
			// if the chunk read was less than the size of buf
			// then check to see if there is a \n at the end
			// if not add it
			if n < 8192 {
				if buf[n] != 10 {
					buf[n] = 10
				}
			}

			// loop through buf
			for j := 0; j < n; j++ {

				// if we are the size of mask buffer then flush it
				if i >= 8192-1 {
					// only write up to i becuase it will be 8181
					// if not it will index error
					//os.Stdout.Write(mask[:i])
					out <- mask[:i]

					// move i back to the beginning of the mask buffer
					i = 0
				}

				// switch through for what character class it is in
				switch {
				// \n so end of the line
				// add a \n to the mask buffer
				case buf[j] == 10:
					mask[i] = 10
					i += 1
				// digit
				case buf[j] >= 47 && buf[j] <= 57:
					mask[i] = 63
					mask[i+1] = 100
					i += 2
				// lower
				case buf[j] >= 97 && buf[j] <= 122:
					mask[i] = 63
					mask[i+1] = 108
					i += 2
				// upper
				case buf[j] >= 65 && buf[j] <= 90:
					mask[i] = 63
					mask[i+1] = 117
					i += 2
				// special
				case (buf[j] >= 32 && buf[j] <= 47) || (buf[j] >= 58 && buf[j] <= 64) || (buf[j] >= 91 && buf[j] <= 96) || (buf[j] >= 123 && buf[j] <= 126):
					mask[i] = 63
					mask[i+1] = 115
					i += 2
				// binary
				// greater than 191 is unicode
				case (buf[j] >= 1 && buf[j] <= 31) || buf[j] == 127 || buf[j] > 191:
					continue
				}
			}

			// if we didnt read in a full buffer
			if n < 8192 {
				// write up to what is filled
				//os.Stdout.Write(mask[:i])
				out <- mask[:i]
				close(quit)
				return
			}
		}
	//}()
	wg.Done()
}
