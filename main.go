package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var flag_nearest *int = flag.Int("nearest", 0, "Missing values get replaced by their nearest neighbor value at most this many cells away")
var flag_missing *float64 = flag.Float64("missing", 0, "Missing values get replaced by this value")

func main() {
	flag.Parse()

	var in io.Reader
	if flag.NArg() == 0 {
		in = os.Stdin
	} else {
		fname := flag.Args()[0]
		var err error
		in, err = os.Open(fname)
		check(err)
	}

	data := read3columns(in)
	I, J, DATA := dgrid(data[0], data[1], data[2], *flag_nearest, *flag_missing)
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
