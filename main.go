package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()
	fname := flag.Args()[0]
	in, err := os.Open(fname)
	Check(err)
	data := ReadArray(in)
	Matrix(data[0], data[1], data[2])
}
