package main

import (
	"io"
	"io/ioutil"
	"strings"
	"strconv"
	"os"
	"fmt"
)

func ReadArray(in io.Reader) (data [][]float64) {
	bytes, err := ioutil.ReadAll(in)
	Check(err)
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
	fmt.Println("numbers:", data)
	return
}

func Check(err os.Error) {
	if err != nil {
		panic(err)
	}
}

func atof(str string) float64 {
	f, err := strconv.Atof64(str)
	Check(err)
	return f
}
