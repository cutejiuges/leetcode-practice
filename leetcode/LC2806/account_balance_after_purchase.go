package LC2806

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/25 下午10:56
 * @FilePath: account_balance_after_purchase
 * @Description:
 */

func accountBalanceAfterPurchase(purchaseAmount int) int {
	amt := purchaseAmount % 10
	if amt < 5 {
		purchaseAmount -= amt
	} else {
		purchaseAmount -= amt - 10
	}
	return 100 - purchaseAmount
}
