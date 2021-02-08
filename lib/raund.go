package lib

import (
	"math/rand"
	"time"
)

// 重み一覧
var weights = [...]int{
	5,  // SSR
	10, // SR
	35, // R
	50, // N
}

func WeightedNumber(times int) [len(weights)]int {

	// 境界値を格納する配列
	border := make([]int, len(weights))

	// 境界値を作成
	for i := 0; i < len(weights); i++ {
		if i == 0 {
			border[i] = weights[i]
		} else {
			border[i] = border[i-1] + weights[i]
		}
	}

	var result [len(weights)]int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < times; i++ {
		draw := rand.Intn(99) + 1 // 1 ~ 100で乱数を生成
		for j, b := range border {
			// 境界値を超えていたら、ランクごとに回数を保存
			if draw <= b {
				result[j]++
				break
			}
		}
	}
	return result
}
