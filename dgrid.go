package main

import (
	"math"
)

// 3 columns contain x,y and data values.
// Sorted, unique x and y indices are returned, which can serve as an "X-Y meshdom" for plotting,
// The data matrix coresponds to x, y: data[i][j] belongs to co-ordinate x[i] y[j]
// if nearest > 0, missing values are  replaced by a nearest neighbor at most 'nearest' cells away.
func dgrid(Icol, Jcol, D []float64, nearest int, missing float64) (i, j []float64, data [][]float64) {

	// (1) Construct a sorted set of unique i,j indices (floats).
	// This is the "meshdom", in matlab terms.
	setI := MakeSet()
	for _, i := range Icol {
		setI.Add(i)
	}
	I := setI.ToArray()

	setJ := MakeSet()
	for _, j := range Jcol {
		setJ.Add(j)
	}
	J := setJ.ToArray()

	//debug("I", I)
	//debug("J", J)
	//debug("D", D)

	const SENTINEL = -123.45678901234567 // quick and dirty hack

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

	//debug("matrix", matrix)

	// (3.5)
	// Missing data gets replaced by nearest value
	if nearest > 0 {
		DELTA := nearest // do not look further than DELTA neighbors 
		for i := range I {
			for j := range J {
				if matrix[I[i]][J[j]] == SENTINEL {

					debug("missing: ", i, j)
					minDst := float64(math.Inf(1))
					nearestV := float64(0)
					for i_ := imax(0, i-DELTA); i_ < imin(len(I), i+DELTA); i_++ {
						for j_ := imax(0, j-DELTA); j_ < imin(len(J), j+DELTA); j_++ {
							if matrix[I[i_]][J[j_]] != SENTINEL {

								dst := sqr(i-i_) + sqr(j-j_) // search for the nearest cell
								if dst < minDst {
									minDst = dst
									nearestV = matrix[I[i_]][J[j_]]
								}

							}
						}
					}
					matrix[I[i]][J[j]] = nearestV
				}
			}
		}
	}

	data = make([][]float64, len(I))
	for i := range I {
		data[i] = make([]float64, len(J))
		for j := range J {
			data[i][j] = matrix[I[i]][J[j]]
			if data[i][j] == SENTINEL {
				data[i][j] = missing
			}
		}
	}
	i = I
	j = J
	return

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

func sqr(x int) float64 {
	return float64(x) * float64(x)
}

//	if octave_format {
//		fmt.Print("x=[")
//		for i := range I {
//			if i != 0 {
//				fmt.Print(", ")
//			}
//			fmt.Print(I[i])
//		}
//		fmt.Println("];")
//
//		fmt.Print("y=[")
//		for i := range J {
//			if i != 0 {
//				fmt.Print(", ")
//			}
//			fmt.Print(J[i])
//		}
//		fmt.Println("];")
//	}

//(4) Print the matrix
//if octave_format {
//	fmt.Print("data=reshape([")
//}
//for ind_i, i := range I {
//	for ind_j, j := range J {
//		fmt.Print(matrix[i][j], "\t")
//		if octave_format && !(ind_i == len(I)-1 && ind_j == len(J)-1) {
//			fmt.Print(",")
//		}
//	}
//	if !octave_format {
//		fmt.Println()
//	}
//}
//if octave_format {
//	fmt.Println("], ", len(J), ", ", len(I), ");")
//}

// Like Matrix() but outputs octave/matlab commands that set
// x, y and data for use in surf(x,y,data)
//func Meshdom(x, y, data []float64) {
//	matrix(x, y, data, true)
//}

// Assuming columns i,j contain matrix indices,
// coutput column data in a correspondig 2D grid.
// Missing values become 0.
//func Matrix(x, y, data []float64) {
//	matrix(x, y, data, false)
//}
