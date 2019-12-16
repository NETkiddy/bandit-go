package bandit_go

import (
	"fmt"
	"math"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type ThomsonSampling struct {
	Src                     rand.Source
	Number                  int
	BanditPlayCount         map[int]int
	BanditPlayWinCount      map[int]float64
	BanditPlayLoseCount     map[int]float64
	BanditPlayReward        map[int]float64
	BanditBeta              map[int]float64
	BanditPlayRewardHistory map[int]float64
}

func NewThomsonSampling(cnt int) ThomsonSampling {
	eg := ThomsonSampling{}
	eg.Src = rand.NewSource(uint64(cnt))
	eg.Number = cnt
	eg.BanditPlayCount = make(map[int]int, 0)
	eg.BanditPlayWinCount = make(map[int]float64, 0)
	eg.BanditPlayLoseCount = make(map[int]float64, 0)
	eg.BanditBeta = make(map[int]float64, 0)
	eg.BanditPlayReward = make(map[int]float64, 0)
	eg.BanditPlayRewardHistory = make(map[int]float64, 0)
	for i := 0; i < eg.Number; i++ {
		eg.BanditPlayWinCount[i] = 0
		eg.BanditPlayLoseCount[i] = 0
		eg.BanditPlayCount[i] = 0
		eg.BanditPlayReward[i] = 0.0
		eg.BanditPlayRewardHistory[i] = 0.0
		eg.BanditBeta[i] = 0.0
	}
	return eg
}

func (e ThomsonSampling) ArmSelector() (idx int) {
	maxValue := float64(math.MinInt16)

	for i := 0; i < e.Number; i++ {
		a := math.Max(e.BanditPlayWinCount[i], 1)
		b := math.Max(e.BanditPlayLoseCount[i], 1)

		v := distuv.Beta{Alpha: a, Beta: b, Src: e.Src}.Rand()
		if v > maxValue {
			maxValue = v
			idx = i
		}
	}

	return
}

func (e ThomsonSampling) PlayBandit(idx int, cost float64) {
	e.BanditPlayCount[idx] += 1

	currPlayCount := e.BanditPlayCount[idx]
	currReward := e.BanditPlayReward[idx]
	newReword := (float64(currPlayCount-1)*currReward + cost) / float64(currPlayCount)

	e.BanditPlayReward[idx] = newReword
	fmt.Printf("PlayBandit currentReword: %v, newReword: %v\n", currReward, newReword)

	e.calcBeta(idx, newReword > currReward)

}

func (e ThomsonSampling) calcBeta(idx int, won bool) {
	if won {
		e.BanditPlayWinCount[idx] += 1
		fmt.Printf("win! [%v]\n", idx)
	} else {
		e.BanditPlayLoseCount[idx] += 1
		fmt.Printf("lose! [%v]\n", idx)
	}

	fmt.Printf("BanditPlayWinCount [%v], BanditPlayLoseCount [%v]\n", e.BanditPlayWinCount, e.BanditPlayLoseCount)
}

func (e ThomsonSampling) PlayCount() int {
	var total int
	for _, v := range e.BanditPlayCount {
		total += v
	}

	return total
}

func (e ThomsonSampling) PlayReword() float64 {
	var total float64
	for _, v := range e.BanditPlayReward {
		total += v
	}
	fmt.Printf("PlayReword: %v", e.BanditPlayReward)

	return total
}
