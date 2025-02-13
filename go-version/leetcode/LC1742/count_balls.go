package lc1742

func countBalls(lowLimit int, highLimit int) int {
	mp := make(map[int]int)
	mx := 0
	for ballNo := lowLimit; ballNo <= highLimit; ballNo++ {
		curBall, boxNo := ballNo, 0
		for curBall > 0 {
			boxNo += curBall % 10
			curBall /= 10
		}
		mp[boxNo]++
		if mp[boxNo] > mx {
			mx = mp[boxNo]
		}
	}
	return mx
}
