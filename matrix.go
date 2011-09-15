package main

import (
	"fmt"
	"os"
	"math"
)

// Like Matrix() but outputs octave/matlab commands that set
// x, y and data for use in surf(x,y,data)
func Meshdom(x, y, data []float64) {
	matrix(x, y, data, true)
}

// Assuming columns i,j contain matrix indices,
// coutput column data in a correspondig 2D grid.
// Missing values become 0.
func Matrix(x, y, data []float64) {
	matrix(x, y, data, false)
}

func matrix(Icol, Jcol, D []float64, octave_format bool) {

	// (1) Construct a sorted set of unique i,j indices (floats).
	// This is the "meshdom", in matlab terms.
	setI := MakeSet()
	for _,i := range Icol {
		setI.Add(i)
	}
	I := setI.ToArray()

	setJ := MakeSet()
	for _,j := range Jcol {
		setJ.Add(j)
	}
	J := setJ.ToArray()

	fmt.Println("I", I)
	fmt.Println("J", J)

	if octave_format {
		fmt.Print("x=[")
		for i := range I {
			if i != 0 {
				fmt.Print(", ")
			}
			fmt.Print(I[i])
		}
		fmt.Println("];")

		fmt.Print("y=[")
		for i := range J {
			if i != 0 {
				fmt.Print(", ")
			}
			fmt.Print(J[i])
		}
		fmt.Println("];")
	}

	var SENTINEL float64 = -123.456789 // quick and dirty hack

	// (2) Make the "outer product" of the two index sets,
	// spanning a matrix that can be index with each possible i,j pair
	// (even those not present in the input, their data will be 0.)
	matrix := make(map[float64]map[float64]float64)
	for i := range I {
		for j := range J {
			if matrix[I[i]] == nil {
				matrix[I[i]] = make(map[float64]float64)
			}
			matrix[I[i]][J[j]] = SENTINEL
		}
	}

	// (3) Loop over the i indices in the output and add the corrsponing data
	// to the corresponding i,j position of the matrix. (j, data on the same line as i)
	// Missing pairs keep 0. as data.
	for i := range Icol {
		matrix[Icol[i]][Jcol[i]] = D[i]
	}

	// (3.5)
	// Missing data gets replaced by nearest value
	DELTA := 5 // do not look further than DELTA neighbors 
	for i := range I {
		for j := range J {
			if matrix[I[i]][J[j]] == SENTINEL {

				fmt.Fprintln(os.Stderr, "missing: ", i, j)
				minDst := float64(math.Inf(1))
				nearest := float64(0)
				for i_ := imax(0, i-DELTA); i_ < imin(len(I), i+DELTA); i_++ {
					for j_ := imax(0, j-DELTA); j_ < imin(len(J), j+DELTA); j_++ {
						if matrix[I[i_]][J[j_]] != SENTINEL {

							dst := sqr(I[i]-I[i_]) + sqr(J[j]-J[j_])
							if dst < minDst {
								minDst = dst
								nearest = matrix[I[i_]][J[j_]]
							}

						}
					}
				}
				matrix[I[i]][J[j]] = nearest

			}
		}
	}

	//(4) Print the matrix
	if octave_format {
		fmt.Print("data=reshape([")
	}
	for ind_i, i := range I {
		for ind_j, j := range J {
			fmt.Print(matrix[i][j], "\t")
			if octave_format && !(ind_i == len(I)-1 && ind_j == len(J)-1) {
				fmt.Print(",")
			}
		}
		if !octave_format {
			fmt.Println()
		}
	}
	if octave_format {
		fmt.Println("], ", len(J), ", ", len(I), ");")
	}
}

func imin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func imax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sqr(x float64) float64 {
	return x * x
}
