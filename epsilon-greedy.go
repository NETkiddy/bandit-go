package bandit_go

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	EPSILON = 0.8
)

type EpsilonGreedy struct {
	MaxIdx                  int
	Number                  int
	BanditPlayCount         map[int]float64
	BanditPlayReward        map[int]float64
	BanditPlayRewardHistory map[int]float64
}

func NewEpsilonGreedy(cnt int) EpsilonGreedy {
	eg := EpsilonGreedy{}
	eg.MaxIdx = 0
	eg.Number = cnt
	eg.BanditPlayCount = make(map[int]float64, 0)
	eg.BanditPlayReward = make(map[int]float64, 0)
	eg.BanditPlayRewardHistory = make(map[int]float64, 0)
	for i := 0; i < eg.Number; i++ {
		eg.BanditPlayCount[i] = 0
		eg.BanditPlayReward[i] = 0.0
		eg.BanditPlayRewardHistory[i] = 0.0
	}

	return eg
}

func (e EpsilonGreedy) ArmSelector() int {
	rnd := rand.Float64()
	fmt.Printf("ArmSelector rnd: %v\n", rnd)
	//Explore
	if rnd < EPSILON {
		return rand.Intn(e.Number)
	}

	//Exploit
	return e.MaxIdx
}

func (e EpsilonGreedy) PlayBandit(idx int, cost float64) {
	e.BanditPlayCount[idx] += 1

	currPlayCount := e.BanditPlayCount[idx]
	currReward := e.BanditPlayReward[idx]
	newReword := (float64(currPlayCount-1)*currReward + cost) / float64(currPlayCount)

	e.BanditPlayReward[idx] = newReword
	fmt.Printf("PlayBandit currentReword: %v, newReword: %v\n", currReward, newReword)

	e.calcMaxRewordIdx()

}

func (e EpsilonGreedy) calcMaxRewordIdx() {
	idx := 0
	reword := float64(math.MinInt16)
	for i, v := range e.BanditPlayReward {
		if v > reword {
			idx = i
			reword = v
		}
	}

	e.MaxIdx = idx
	fmt.Printf("calcMaxRewordIdx [%v], reword [%v]\n", e.MaxIdx, reword)
}

func (e EpsilonGreedy) PlayCount() float64 {
	var total float64
	for _, v := range e.BanditPlayCount {
		total += v
	}

	return total
}
func (e EpsilonGreedy) PlayReword() float64 {
	var total float64
	for _, v := range e.BanditPlayReward {
		total += v
	}
	fmt.Printf("PlayReword: %v", e.BanditPlayReward)

	return total
}
