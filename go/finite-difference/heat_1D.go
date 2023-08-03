package main

import (
	"fmt"
	"os"
)

type heat_1D_bar struct {
	N int // number of spatial points
	M int // number of temporal points

	L     float64 // length of bar
	alpha float64 // thermal diffusivity (for the form du/dt = alpha * d2u/d2x)

	delx, delt float64 // time/space steps

	// boundary conditions
	T1, T2 float64

	u [][]float64
	t []float64
	x []float64

	// internal params
	internal struct {
		r float64

		a []float64
		b []float64
		c []float64
		d []float64
	}
}

// u0(x,L) = ...
type bar_init_fcn func(float64, float64) float64

func (b *heat_1D_bar) create(num_points, num_time int, bar_length, T1, T2, alpha float64, u0 bar_init_fcn) {
	b.N = num_points
	b.M = num_time
	b.L = bar_length
	b.delx = bar_length / float64(num_points)
	b.T1 = T1
	b.T2 = T2

	// calc delt to enforce stability
	max_delt := (0.5) * b.delx * b.delx / alpha
	b.delt = max_delt

	b.internal.r = alpha * b.delt / (b.delx * b.delx)
	b.internal.a = make([]float64, num_points+1)
	b.internal.b = make([]float64, num_points+1)
	b.internal.c = make([]float64, num_points+1)
	b.internal.d = make([]float64, num_points+1)

	// allocate memory for data
	b.t = make([]float64, num_time+1)
	for i := 0; i <= num_time; i++ {
		b.t[i] = float64(i) * b.delt
	}
	b.x = make([]float64, num_points+1)
	for i := 0; i <= num_points; i++ {
		b.x[i] = float64(i) * b.delx
	}

	b.u = make([][]float64, num_time+1)
	for m := range b.u {
		b.u[m] = make([]float64, num_points+1)
	}

	// initialize using the user supplied function
	for n := 0; n <= num_points; n++ {
		x := b.x[n]
		b.u[0][n] = u0(x, bar_length)
	}
	b.u[0][0] = T1
	b.u[0][num_points] = T2
}

func (b *heat_1D_bar) reset() {
	// initialize using the user supplied function
	//for n := 0; n <= b.N; n++ {
	//	x := b.x[n]
	//	b.u[0][n] = u0(x, b.L)
	//}

	for k := 1; k <= b.M; k++ {
		for n := 0; n <= b.N; n++ {
			b.u[k][n] = 0
		}
	}
}

func (b *heat_1D_bar) solve_FTCS() {
	fmt.Println("Solving using FTCS method.")
	fmt.Println("N =", b.N, ", M =", b.M)
	r := b.internal.r

	for k := 1; k <= b.M; k++ {
		(b.u)[k][0] = b.T1
		for n := 1; n <= (b.N - 1); n++ {
			(b.u)[k][n] = (b.u)[k-1][n] + r*(b.u[k-1][n-1]-2*b.u[k-1][n]+b.u[k-1][n+1])
		}
		b.u[k][b.N] = b.T2
	}
}

func (b *heat_1D_bar) solve_BTCS() {
	fmt.Println("Solving using BTCS method.")
	fmt.Println("N =", b.N, ", M =", b.M)
	r := b.internal.r

	// construct tridagonal coefficients - these don't change
	b.internal.a[0] = 0 // unused
	for i := 1; i < b.N; i++ {
		b.internal.a[i] = -r
	}
	b.internal.a[b.N] = 0

	b.internal.b[0] = 1
	for i := 1; i < (b.N + 1); i++ {
		b.internal.b[i] = 1 + 2*r
	}
	b.internal.b[b.N] = 1

	b.internal.c[0] = 0
	for i := 1; i < b.N; i++ {
		b.internal.c[i] = -r
	}

	// reduced c coeffiecient
	cp := make([]float64, b.N+1)
	cp[0] = b.internal.c[0] / b.internal.b[0]
	for i := 1; i < b.N; i++ {
		cp[i] = b.internal.c[i] / (b.internal.b[i] - b.internal.a[i]*cp[i-1])
	}

	for k := 1; k <= b.M; k++ {
		// d coefficient is constant in this method
		//b.internal.d[0] = b.T1
		for i := 0; i <= b.N; i++ {
			b.internal.d[i] = b.u[k-1][i]
		}
		//b.internal.d[b.N] = b.T2

		// reduced d coeffiecient (Thomas Method)
		dp := make([]float64, b.N+1)
		dp[0] = b.internal.d[0] / b.internal.b[0]
		for i := 1; i < b.N+1; i++ {
			dp[i] = (b.internal.d[i] - b.internal.a[i]*dp[i-1]) / (b.internal.b[i] - b.internal.a[i]*cp[i-1])
		}

		// substitue back using these new coefficients
		b.u[k][b.N] = dp[b.N]
		for i := (b.N - 1); i >= 0; i-- {
			b.u[k][i] = dp[i] - cp[i]*b.u[k][i+1]
		}
	}
}

func (b *heat_1D_bar) solve_CTCS() {
	fmt.Println("Solving using CTCS method.")
	fmt.Println("N =", b.N, ", M =", b.M)
	r := b.internal.r

	// construct tridagonal coefficients - these don't change
	b.internal.a[0] = 0
	for i := 1; i < b.N; i++ {
		b.internal.a[i] = -1
	}
	b.internal.a[b.N] = 0

	b.internal.b[0] = 1
	for i := 1; i < (b.N + 1); i++ {
		b.internal.b[i] = 2 * (1 + r) / r
	}
	b.internal.b[b.N] = 1

	b.internal.c[0] = 0
	for i := 1; i < b.N; i++ {
		b.internal.c[i] = -1
	}

	// reduced c coeffiecient (Thomas Method)
	cp := make([]float64, b.N+1)
	cp[0] = b.internal.c[0] / b.internal.b[0]
	for i := 1; i < b.N; i++ {
		cp[i] = b.internal.c[i] / (b.internal.b[i] - b.internal.a[i]*cp[i-1])
	}

	for k := 1; k <= b.M; k++ {
		// d coefficient
		b.internal.d[0] = b.u[k-1][0]
		for i := 1; i < b.N; i++ {
			b.internal.d[i] = b.u[k-1][i-1] + (2*(1-r)/r)*b.u[k-1][i] + b.u[k-1][i+1]
		}
		b.internal.d[b.N] = b.u[k-1][b.N]

		// reduced d coeffiecient (Thomas Method)
		dp := make([]float64, b.N+1)
		dp[0] = b.internal.d[0] / b.internal.b[0]
		for i := 1; i < b.N+1; i++ {
			dp[i] = (b.internal.d[i] - b.internal.a[i]*dp[i-1]) / (b.internal.b[i] - b.internal.a[i]*cp[i-1])
		}

		// substitue back using these new coefficients
		b.u[k][b.N] = dp[b.N]
		for i := (b.N - 1); i >= 0; i-- {
			b.u[k][i] = dp[i] - cp[i]*b.u[k][i+1]
		}
	}
}

func (b *heat_1D_bar) write_to_file(filename string) {
	fid, err := os.Create(filename)
	if fid == nil {
		panic(err)
	}
	fmt.Printf("'%s' opened...", filename)

	// close file when done
	defer func() {
		fid.Close()
		fmt.Println("File closed.")
	}()

	fid.Write([]byte(fmt.Sprintf("N = %v, M = %v, delt = %v, delx = %v\n", b.N, b.M, b.delt, b.delx)))

	for m := 0; m <= b.M; m++ {
		t := float64(m) * b.delt
		fid.Write([]byte(fmt.Sprintf("%.4e = [", t)))

		for n := 0; n < b.N; n++ {
			fid.Write([]byte(fmt.Sprintf("%.2f, ", b.u[m][n])))
		}
		fid.Write([]byte(fmt.Sprintf("%.2f]\n", b.u[m][b.N])))
	}
}
