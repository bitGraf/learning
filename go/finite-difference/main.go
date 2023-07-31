package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	const N int = 100
	const M int = 100
	const L float64 = 0.1
	var alpha = math.Sqrt(111.0)
	const delx float64 = L / float64(N)
	max_delt := delx * delx / (2 * alpha * alpha)
	var delt float64 = max_delt

	//fmt.Println("delx*delx/2 =", delx*delx/2)
	//fmt.Println("delt =", delt)
	//fmt.Println("stable: ", delt <= delx*delx/2)

	//u := [M + 1][N + 1]float64{}
	u := make([][]float64, M+1)
	for m := range u {
		u[m] = make([]float64, N+1)
	}
	var T1, T2 float64 = 300, 400

	// initialize bar to f(x) = (T1+T2)/2
	u[0][0] = T1
	u[0][N] = T2
	for n := 1; n <= (N - 1); n++ {
		u[0][n] = (T1 + T2) / 2.0
	}

	// loop through time
	h := (alpha * alpha) * (delt) / (delx * delx)
	fmt.Println("h =", h)
	print_at_time(N, u[0][:])
	for k := 1; k <= M; k++ {

		u[k][0] = T1
		for n := 1; n <= (N - 1); n++ {
			u[k][n] = u[k-1][n] + h*(u[k-1][n-1]-2*u[k-1][n]+u[k-1][n+1])
		}
		u[k][N] = T2

		print_at_time(N, u[k][:])
	}

	//fmt.Println("u =", u)

	write_to_file(u, N, M, delt, delx, "heat_1D.txt")
}

func print_at_time(N int, u []float64) {
	fmt.Printf("[%6.0f, ", u[0])
	for n := 1; n <= (N - 1); n++ {
		fmt.Printf("%6.2f, ", u[n])
	}
	fmt.Printf("%6.2f]\n", u[N])
}

func write_to_file(u [][]float64, N, M int, delt, delx float64, filename string) {
	fid, err := os.Create(filename)
	if fid == nil {
		panic(err)
	}

	// close file when done
	defer func() {
		fid.Close()
		fmt.Println("File closed.")
	}()

	fid.Write([]byte(fmt.Sprintf("N = %v, M = %v, delt = %v, delx = %v\n", N, M, delt, delx)))

	for m := 0; m <= M; m++ {
		t := float64(m) * delt
		fid.Write([]byte(fmt.Sprintf("%.4e = [", t)))

		for n := 0; n < N; n++ {
			fid.Write([]byte(fmt.Sprintf("%.2f, ", u[m][n])))
		}
		fid.Write([]byte(fmt.Sprintf("%.2f]\n", u[m][N])))
	}
}
