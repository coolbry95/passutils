// rule applies rules to wordlists
// it handles UTF-8 rules such as appending $世
package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"github.com/coolbry95/passutils/ruleprocessor/rules"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("%s rules list", os.Args[0])
		os.Exit(-1)
	}

	rulefile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	wordlist, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
	}

	scanner := bufio.NewScanner(rulefile)
	r := make([][]string, 0, 100)
	for scanner.Scan() {
		temp := rules.ParseLine(scanner.Text())
		if temp != nil {
			r = append(r, temp)
		}
	}

	p := make(chan string, runtime.NumCPU())
	out := make(chan string, runtime.NumCPU())


	go func() {
		for print := range out {
			if print == "" {
				continue
			}
			fmt.Println(print)
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for pass := range p {
				for _, ru := range r {
					out <- rules.ApplyRules(ru, pass)
				}
			}
		}()
	}

	scanning := bufio.NewScanner(wordlist)
	for scanning.Scan() {
		p <- scanning.Text()

	}

	close(p)
	close(out)
}
