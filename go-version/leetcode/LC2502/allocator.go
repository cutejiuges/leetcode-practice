package LC2502

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/2/25 下午11:57
 * @FilePath: allocator
 * @Description:
 */

type Allocator struct {
	n      int
	memory []int
}

func Constructor(n int) Allocator {
	return Allocator{
		n:      n,
		memory: make([]int, n),
	}
}

func (this *Allocator) Allocate(size int, mID int) int {
	cnt := 0
	for i := 0; i < this.n; i++ {
		if this.memory[i] != 0 {
			cnt = 0
		} else {
			cnt++
			if cnt == size {
				for j := i - cnt + 1; j <= i; j++ {
					this.memory[j] = mID
				}
				return i - cnt + 1
			}
		}
	}
	return -1
}

func (this *Allocator) FreeMemory(mID int) int {
	cnt := 0
	for i := 0; i < this.n; i++ {
		if this.memory[i] == mID {
			this.memory[i] = 0
			cnt++
		}
	}
	return cnt
}
