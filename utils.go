package bandit_go

import "gonum.org/v1/plot/plotter"

func genPts(data map[int]float64) plotter.XYs {
	pts := make(plotter.XYs, len(data))
	for i := 0; i < len(data); i++ {
		pts[i].X = float64(i)
		pts[i].Y = data[i]
	}

	return pts
}
