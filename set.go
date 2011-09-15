package main

type Empty struct{}

type Set map[float64]Empty

func MakeSet() Set {
	return Set(make(map[float64]Empty))
}

func (s Set) Add(x float64) {
	if _, ok := s[x]; !ok {
		s[x] = Empty{}
	}
}

func (s Set) ToArray() []float64 {
	array := make([]float64, len(s))
	i := 0
	for val, _ := range s {
		array[i] = val
		i++
	}
	Float64Array(array).Sort()
	return array
}
