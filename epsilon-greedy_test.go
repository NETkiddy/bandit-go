package bandit_go

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"testing"
)

func TestEpsilonGreedy(t *testing.T) {
	costs := []float64{3.0, 1.0, 5.0, 3.0, 3.0, 3.0, 8.0, 3.0, 7.0, 3.0, 3.0, 3.0, 9.0, 3.0, 7.0, 3.0,3.0, 1.0, 5.0, 3.0, 3.0, 3.0, 8.0, 3.0, 7.0, 3.0, 3.0, 3.0, 9.0, 3.0, 7.0, 3.0,3.0, 1.0, 5.0, 3.0, 3.0, 3.0, 8.0, 3.0, 7.0, 3.0, 3.0, 3.0, 9.0, 3.0, 7.0, 3.0,3.0, 1.0, 5.0, 3.0, 3.0, 3.0, 8.0, 3.0, 7.0, 3.0, 3.0, 3.0, 9.0, 3.0, 7.0, 3.0,3.0, 1.0, 5.0, 3.0, 3.0, 3.0, 8.0, 3.0, 7.0, 3.0, 3.0, 3.0, 9.0, 3.0, 7.0, 3.0,3.0, 1.0, 5.0, 3.0, 3.0, 3.0, 8.0, 3.0, 7.0, 3.0, 3.0, 3.0, 9.0, 3.0, 7.0, 3.0,3.0, 1.0, 5.0, 3.0, 3.0, 3.0, 8.0, 3.0, 7.0, 3.0, 3.0, 3.0, 9.0, 3.0, 7.0, 3.0,3.0, 1.0, 5.0, 3.0, 3.0, 3.0, 8.0, 3.0, 7.0, 3.0, 3.0, 3.0, 9.0, 3.0, 7.0, 3.0,3.0, 1.0, 5.0, 3.0, 3.0, 3.0, 8.0, 3.0, 7.0, 3.0, 3.0, 3.0, 9.0, 3.0, 7.0, 3.0,3.0, 1.0, 5.0, 3.0, 3.0, 3.0, 8.0, 3.0, 7.0, 3.0, 3.0, 3.0, 9.0, 3.0, 7.0, 3.0}
	eg := NewEpsilonGreedy(10)

	for i := 0; i < len(costs); i++ {
		nextIdx := eg.ArmSelector()
		fmt.Printf("ArmSelector [%v]\n", nextIdx)
		eg.PlayBandit(nextIdx, costs[i])

		eg.BanditPlayRewardHistory[i] = eg.PlayReword()

		fmt.Printf("Round [%v], current total reword:%v\n", i, eg.PlayReword())
	}

	//============================================
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	err = plotutil.AddLinePoints(p,
		//"count", genPts(eg.BanditPlayCount),
		"reword", genPts(eg.BanditPlayRewardHistory))

	if err != nil {
		panic(err)
	}
	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "TestEpsilonGreedy.png"); err != nil {
		panic(err)
	}

}
