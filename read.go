package main

import (
	"io"
	"io/ioutil"
	"strings"
	"strconv"
	"os"
)

// Reads a 3-column array of ascii numbers:
//	1.0 2.0 3.0
//	4.0 5.0 6.0
//	7.0 8.0 9.0
// 	...
func read3columns(in io.Reader) (data [][]float64) {
	bytes, err := ioutil.ReadAll(in)
	check(err)
	data = make([][]float64, 3)
	str := string(bytes)
	lines := strings.Split(str, "\n")
	for _, l := range lines {
		i := 0
		words1 := strings.Split(l, "\t")
		for _, w1 := range words1 {
			words := strings.Split(w1, " ")
			for _, w := range words {
				if len(w) > 0 {
					data[i] = append(data[i], atof(w))
					i++
				}
			}
		}
		if i != 0 && i != 3 {
			panic(l)
		}
	}
	//debug("numbers:", data)
	return
}

// panics if err != nil
func check(err os.Error) {
	if err != nil {
		panic(err)
	}
}

// atof which panics on error
func atof(str string) float64 {
	f, err := strconv.Atof64(str)
	check(err)
	return f
}
