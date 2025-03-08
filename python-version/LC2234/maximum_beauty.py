class Solution:
    def maximumBeauty(self, flowers: list[int], new_flowers: int, target: int, full: int, partial: int) -> int:
        n = len(flowers)

        # 先假设全部种满，看看能剩下多少花
        left_flowers = new_flowers - target * n
        for i in range(n):  # 把原来有的花补上去
            flowers[i] = min(flowers[i], target)
            left_flowers += flowers[i]

        # 如果本来全都是满的，不需要种，直接返回
        if left_flowers == new_flowers:
            return full * n

        # 如果可以种满，从两种策略取其大，其一全部种满，其二留一个剩一朵，别的种满
        if left_flowers >= 0:
            return max(full * n, (n - 1) * full + (target - 1) * partial)

        flowers.sort()
        # 如果种不满，枚举能种满的后缀，多出来的补到前面去，选出最大的
        ans, pre_sum, j = 0, 0, 0
        # 枚举能种满的后缀，从1开始，0上面已经算过了
        for i in range(1, n + 1):
            # 撤销在i-1位置种的花
            left_flowers += target - flowers[i-1]
            if left_flowers < 0:  # 如果i位置后面还是种不满，继续撤销
                continue

            # 后面能种满的时候往前面种，种到flowers[j]
            while j < i and flowers[j] * j <= left_flowers + pre_sum:
                pre_sum += flowers[j]
                j += 1

            # 计算当前的美丽值
            avg = (left_flowers + pre_sum) // j
            beauty_value = avg * partial + full * (n-i)
            ans = max(ans, beauty_value)
        return ans
