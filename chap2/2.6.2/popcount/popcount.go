package popcount

import "fmt"

// pc[i] は i のポピュレーションカウンタ
var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//PopCount は x のポピュレーションカウント (1 が設定されているビット数) を返す
func PopCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var pcOfx byte
	for i := 0; i < 8; i++ {
		pcOfx += pc[byte(x>>(uint(i)*8))]
	}
	return int(pcOfx)
}

func PopCount3(x uint64) int {
	var count int
	for i := 0; i < 64; i++ {
		count += int((x >> uint64(i)) & uint64(1))
	}
	return count
}

func PopCount4(x uint64) int {
	var count int

	for x > 0 {
		x &= (x - uint64(1))
		fmt.Print(x, "\n")
		count++
	}

	return count
}
