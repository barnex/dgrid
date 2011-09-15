package main

import (
	"io"
	"io/ioutil"	
	"strings"
	"strconv"
	"os"
)

func ReadArray(in io.Reader) (data [][]float64){
	bytes, err := ioutil.ReadAll(in)
	Check(err)
	str := string(bytes)
	lines := strings.Split(str, "\n")
	for _,l := range lines{
		var numbers []float64
		words1 := strings.Split(l, "\t")
		for _,w1 := range words1{
				words := strings.Split(w1, " ")
				numbers = make([]float64, len(words))
				for i, w := range words{
						numbers[i]	= atof(w)
				}
		}
		data = append(data, numbers)
	}
	return
}


func Check(err os.Error){
	if err != nil{panic(err)}
}

func atof(str string) float64{
	f, err := strconv.Atof64(str)
	Check(err)
	return f
}
