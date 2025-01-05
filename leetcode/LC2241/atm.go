package LC2241

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/1/5 下午8:21
 * @FilePath: atm
 * @Description:
 */

type ATM struct {
	cnt, value []int //分别表示每种面额的数量和面额
}

func Constructor() ATM {
	return ATM{
		cnt:   make([]int, 5),
		value: []int{20, 50, 100, 200, 500},
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < 5; i++ {
		this.cnt[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	ans := make([]int, 5)

	for i := 4; i >= 0; i-- {
		ans[i] = min(this.cnt[i], int(amount/this.value[i]))
		amount -= ans[i] * this.value[i]
	}

	if amount > 0 {
		return []int{-1}
	}
	for i := 0; i < 5; i++ {
		this.cnt[i] -= ans[i]
	}
	return ans
}
