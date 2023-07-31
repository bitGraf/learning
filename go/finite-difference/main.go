package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	const N int = 100            // number of spatial points
	const M int = 200            // number of temporal points
	const L float64 = 1          // size of bar
	var alpha = math.Sqrt(111.0) // thermal diffusivity
	const delx float64 = L / float64(N)
	max_delt := delx * delx / (2 * alpha * alpha)
	var delt float64 = max_delt

	fmt.Println("N =", N, ", M =", M)
	fmt.Println("delX =", delx, "delT =", delt)
	fmt.Println("Stable if r <= 0.5")
	fmt.Println("r =", (alpha*alpha)*delt/(delx*delx))

	// Allocate a 2-D slice of slices
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
		//u[0][n] = (T1 + T2) / 2.0
		u[0][n] = (T1 + T2)
	}

	// Solve using Forward-Difference in time, Central-Difference in space
	r := (alpha * alpha) * (delt) / (delx * delx)
	solve_FTCS(&u, T1, T2, r)

	write_to_file(u, N, M, delt, delx, "heat_1D_FTCS.txt")

	// solve using Backward-Difference in time, Central-Difference in space
	//delt = 0.00001
	//r = (alpha * alpha) * (delt) / (delx * delx)
	fmt.Println("r =", r)
	solve_BTCS(&u, T1, T2, r)

	write_to_file(u, N, M, delt, delx, "heat_1D_BTCS.txt")

	// solve using Central-Difference in time, Central-Difference in space
	// Crank-Nicholson method
	solve_CTCS(&u, T1, T2, r)

	write_to_file(u, N, M, delt, delx, "heat_1D_CTCS.txt")
}

func solve_FTCS(u *[][]float64, T1, T2, r float64) {
	M := len(*u) - 1
	N := len((*u)[0]) - 1
	//fmt.Println("N =", N, ", M =", M)

	//print_at_time(N, u[0][:])
	for k := 1; k <= M; k++ {

		(*u)[k][0] = T1
		for n := 1; n <= (N - 1); n++ {
			(*u)[k][n] = (*u)[k-1][n] + r*((*u)[k-1][n-1]-2*(*u)[k-1][n]+(*u)[k-1][n+1])
		}
		(*u)[k][N] = T2

		//print_at_time(N, u[k][:])
	}
}

func solve_BTCS(u *[][]float64, T1, T2, r float64) {
	M := len(*u) - 1
	N := len((*u)[0]) - 1
	fmt.Println("N =", N, ", M =", M)

	//print_at_time(N, u[0][:])
	for k := 1; k <= M; k++ {

		// construct tridagonal coefficients
		n := N + 1
		a := make([]float64, n) // zero-initialized
		b := make([]float64, n)
		c := make([]float64, n)
		d := (*u)[k-1]
		//fmt.Println("d =", d)

		a[0] = 0 // unused
		for i := 1; i < N; i++ {
			a[i] = -r
		}
		a[N] = 0

		//fmt.Println("a =", a)

		b[0] = 1
		for i := 1; i < n; i++ {
			b[i] = 1 + 2*r
		}
		b[N] = 1
		//fmt.Println("b =", b)

		c[0] = 0
		for i := 1; i < (n - 1); i++ {
			c[i] = -r
		}
		//fmt.Println("c =", c)

		// find reduced coeffiecients (Thomas Method)
		cp := make([]float64, n)
		cp[0] = c[0] / b[0]
		for i := 1; i < (n - 1); i++ {
			cp[i] = c[i] / (b[i] - a[i]*cp[i-1])
		}
		//fmt.Println("cp =", cp)

		dp := make([]float64, n)
		dp[0] = d[0] / b[0]
		for i := 1; i < n; i++ {
			dp[i] = (d[i] - a[i]*dp[i-1]) / (b[i] - a[i]*cp[i-1])
		}
		//fmt.Println("dp =", dp)

		// substitue back using these new coefficients
		(*u)[k][N] = dp[N]
		for i := (n - 2); i >= 0; i-- {
			(*u)[k][i] = dp[i] - cp[i]*(*u)[k][i+1]
		}
		//fmt.Println("x =", (*u)[k])

		//print_at_time(N, (*u)[k])

		//break
	}
}

func solve_CTCS(u *[][]float64, T1, T2, r float64) {
	M := len(*u) - 1
	N := len((*u)[0]) - 1
	fmt.Println("N =", N, ", M =", M)

	//print_at_time(N, u[0][:])
	for k := 1; k <= M; k++ {

		// construct tridagonal coefficients
		n := N + 1
		a := make([]float64, n)
		b := make([]float64, n)
		c := make([]float64, n)
		d := make([]float64, n)

		// constant vector
		//d := (*u)[k-1]
		d[0] = (*u)[k-1][0]
		for i := 1; i < N; i++ {
			d[i] = (*u)[k-1][i-1] + (2*(1-r)/r)*(*u)[k-1][i] + (*u)[k-1][i+1]
		}
		d[N] = (*u)[k-1][N]
		//fmt.Println("d =", d)

		a[0] = 0 // unused
		for i := 1; i < N; i++ {
			a[i] = -1
		}
		a[N] = 0

		//fmt.Println("a =", a)

		b[0] = 1
		for i := 1; i < n; i++ {
			b[i] = 2 * (1 + r) / r
		}
		b[N] = 1
		//fmt.Println("b =", b)

		c[0] = 0
		for i := 1; i < (n - 1); i++ {
			c[i] = -1
		}
		//fmt.Println("c =", c)

		// find reduced coeffiecients (Thomas Method)
		cp := make([]float64, n)
		cp[0] = c[0] / b[0]
		for i := 1; i < (n - 1); i++ {
			cp[i] = c[i] / (b[i] - a[i]*cp[i-1])
		}
		//fmt.Println("cp =", cp)

		dp := make([]float64, n)
		dp[0] = d[0] / b[0]
		for i := 1; i < n; i++ {
			dp[i] = (d[i] - a[i]*dp[i-1]) / (b[i] - a[i]*cp[i-1])
		}
		//fmt.Println("dp =", dp)

		// substitue back using these new coefficients
		(*u)[k][N] = dp[N]
		for i := (n - 2); i >= 0; i-- {
			(*u)[k][i] = dp[i] - cp[i]*(*u)[k][i+1]
		}
		//fmt.Println("x =", (*u)[k])

		//print_at_time(N, (*u)[k])

		//break
	}
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
