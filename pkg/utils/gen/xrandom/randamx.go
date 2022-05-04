package xrandom

import (
	"math/rand"
	"time"
)

const (
	RAND_NUM   = 0 // 纯数字
	RAND_LOWER = 1 // 小写字母
	RAND_UPPER = 2 // 大写字母
	RAND_ALL   = 3 // 数字、大小写字母
)

// GetRandom 随机字符串 0 纯数字 1 小写字母 2 大写字母 3 数字、大小写字母
func GetRandom(size int, kind int) string {
	iKind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			iKind = rand.Intn(3)
		}
		scope, base := kinds[iKind][0], kinds[iKind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
