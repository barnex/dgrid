package main

import (
	"flag"
	"os"
	"fmt"
)

func main() {
	flag.Parse()
	fname := flag.Args()[0]
	in, err := os.Open(fname)
	check(err)
	data := read3columns(in)
	I, J, DATA := dgrid(data[0], data[1], data[2])
	for i := range I {
		for j := range J {
			fmt.Println(I[i], J[j], DATA[i][j])
		}
		fmt.Println()
	}
}


func debug(msg ...interface{}) {
	fmt.Fprintln(os.Stderr, msg...)
}
