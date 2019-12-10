package bandit_go

import (
	"fmt"
	"math"
)

type UCB struct {
	MaxIdx           int
	Number           int
	BanditPlayCount  map[int]int
	BanditPlayReward map[int]float64
	BanditFitness    map[int]float64
	BanditPlayRewardHistory map[int]float64
}

func (u *UCB) calcUcbFitness() {
	for idx, cnt := range u.BanditPlayCount {
		bonus := math.Sqrt(2*math.Log2(float64(u.PlayCount()))) / float64(cnt)
		avgReword := u.BanditPlayReward[idx] / float64(cnt)
		u.BanditFitness[idx] = avgReword + bonus
	}
	fmt.Printf("calcUcbFitness: %v\n", u.BanditFitness)
}

func NewUCB(cnt int) UCB {
	eg := UCB{}
	eg.MaxIdx = 0
	eg.Number = cnt
	eg.BanditPlayCount = make(map[int]int, 0)
	eg.BanditPlayReward = make(map[int]float64, 0)
	eg.BanditFitness = make(map[int]float64, 0)
	eg.BanditPlayRewardHistory = make(map[int]float64, 0)
	for i := 0; i < eg.Number; i++ {
		eg.BanditPlayCount[i] = 0
		eg.BanditPlayReward[i] = 0.0
		eg.BanditPlayRewardHistory[i] = 0.0
		eg.BanditFitness[i] = 0.0
	}

	return eg
}

func (u UCB) InitBandit() {
	for i := 0; i < u.Number; i++ {
		u.PlayBandit(i, 2.0)
	}
}

func (u UCB) ArmSelector() int {
	maxFitness := float64(math.MinInt32)
	maxIdx := 0
	for idx, f := range u.BanditFitness {
		if maxFitness < f {
			maxFitness = f
			maxIdx = idx
		}
	}

	return maxIdx
}

func (u UCB) PlayBandit(idx int, cost float64) {
	u.BanditPlayCount[idx] += 1

	currPlayCount := u.BanditPlayCount[idx]
	currReward := u.BanditPlayReward[idx]
	newReword := (float64(currPlayCount-1)*currReward + cost) / float64(currPlayCount)

	u.BanditPlayReward[idx] = newReword
	fmt.Printf("PlayBandit: %v, currentReword: %v, newReword: %v\n", idx, currReward, newReword)

	u.calcUcbFitness()

}

func (u *UCB) calcMaxRewordIdx() {
	idx := 0
	reword := float64(math.MinInt16)
	for i, v := range u.BanditPlayReward {
		if v > reword {
			idx = i
			reword = v
		}
	}

	u.MaxIdx = idx
	fmt.Printf("calcMaxRewordIdx [%v], reword [%v]\n", u.MaxIdx, reword)
}

func (u UCB) PlayCount() int {
	var total int
	for _, v := range u.BanditPlayCount {
		total += v
	}

	return total
}

func (u UCB) PlayReword() float64 {
	var total float64
	for _, v := range u.BanditPlayReward {
		total += v
	}
	fmt.Printf("PlayReword: %v", u.BanditPlayReward)

	return total
}
