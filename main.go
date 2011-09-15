package main

import(
			"flag"
			"os"
	)

func main() {
		fname := flag.Arg(0)
	in,err := os.Open(fname)
	Check(err)
	data := ReadArray(in)
	Matrix( data[0], data[1], data[2])
}
