package main

import (
	"os"
	"fmt"
	"flag"
//	"math/big"
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: number2tupper [inputfile]\n")
    os.Exit(2)
}

func fileparser(filename string) {
        
// luku := big.NewInt(0)

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		usage()
	} else {
		//pass
	}
}
