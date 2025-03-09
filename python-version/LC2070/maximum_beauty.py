class Solution:
    def maximumBeauty(self, items: list[list[int]], queries: list[int]) -> list[int]:
        # 按照price进行排序
        items.sort(key=lambda x: x[0])
        # 替换beauty的前缀最大值
        n = len(items)
        for i in range(n-1):
            items[i+1][1] = max(items[i+1][1], items[i][1])

        # 按照二分查找进行查询处理
        n = len(queries)
        ans = [0] * n
        for i in range(n):
            ans[i] = self.__query(items, queries[i])
        return ans

    def __query(self, items: list[list[int]], q: int) -> int:
        low, high = 0, len(items)
        while low < high:
            mid = low + ((high - low) >> 1)
            if items[mid][0] > q:
                high = mid
            else:
                low = mid + 1
        if low == 0:
            return 0
        return items[low-1][1]
