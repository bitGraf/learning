package main

func main() {
	var bar heat_1D_bar
	bar.create(100, 500, 1.0, 300, 400, 111.0, func(x, L float64) float64 { return 0.0 })

	bar.solve_FTCS()
	bar.write_to_file("heat_1D_FTCS.txt")

	bar.reset()

	bar.solve_BTCS()
	bar.write_to_file("heat_1D_BTCS.txt")

	bar.reset()

	bar.solve_CTCS()
	bar.write_to_file("heat_1D_CTCS.txt")
}
